// gocorenlp.  A Go (Golang) client for Stanford CoreNLP server.
// Copyright (C) 2022  Yuan Gao
//
// This file is part of gocorenlp.
//
// gocorenlp is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package pbtest

import (
	"encoding/base64"
	"fmt"
	"reflect"

	gogoerrors "github.com/donyori/gogo/errors"
	"google.golang.org/protobuf/proto"

	"github.com/donyori/gocorenlp/model"
)

// RosesAreRed is the text used as annotation input in model testing.
const RosesAreRed = `
Roses are red.
  Violets are blue.
Sugar is sweet.
  And so are you.
`

// RosesAreRedSentenceNum is the expected number of sentences for RosesAreRed.
const RosesAreRedSentenceNum = 4

// Expected annotation results for RosesAreRed.
var (
	RosesAreRedSentenceTokenNumList   = [RosesAreRedSentenceNum]int{4, 4, 4, 5}
	RosesAreRedSentenceTokenWordLists = [RosesAreRedSentenceNum][]string{
		{"Roses", "are", "red", "."},
		{"Violets", "are", "blue", "."},
		{"Sugar", "is", "sweet", "."},
		{"And", "so", "are", "you", "."},
	}
	RosesAreRedSentenceTokenGapLists = [RosesAreRedSentenceNum][]string{
		{"\n", " ", " ", "", "\n  "},
		{"\n  ", " ", " ", "", "\n"},
		{"\n", " ", " ", "", "\n  "},
		{"\n  ", " ", " ", " ", "", "\n"},
	}
	RosesAreRedSentenceTokenPosLists = [RosesAreRedSentenceNum][]string{
		{"NNPS", "VBP", "JJ", "."},
		{"NNS", "VBP", "JJ", "."},
		{"NNP", "VBZ", "JJ", "."},
		{"CC", "RB", "VBP", "PRP", "."},
	}
)

// TextGetter wraps the method GetText.
type TextGetter interface {
	GetText() string
}

// Document combines interfaces proto.Message and TextGetter,
// and a method GetDocID.
//
// The client should only pass the document structure
// to this type of parameter.
// Functions in this package may call the document's
// GetSentence method through reflection.
type Document interface {
	proto.Message
	TextGetter
	GetDocID() string
}

// DecodeBase64ToPb decodes the base64 encoded CoreNLP server
// response body into a ProtoBuf message.
func DecodeBase64ToPb(base64String string, msg proto.Message) error {
	b, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return gogoerrors.AutoWrap(err)
	}
	return gogoerrors.AutoWrap(model.DecodeResponseBody(b, msg))
}

// CheckText examines whether the text returned by getter.GetText
// is the same as RosesAreRed.
//
// It reports an error if getter is nil or the text returned by
// getter.GetText is different from RosesAreRed.
func CheckText(getter TextGetter) error {
	if getter == nil {
		return gogoerrors.AutoNew("text getter is nil")
	}
	if text := getter.GetText(); text != RosesAreRed {
		return gogoerrors.AutoNew(fmt.Sprintf("got %q; want %q", text, RosesAreRed))
	}
	return nil
}

// CheckDocumentFromBase64 decodes base64String to doc
// and then checks the annotation results in doc.
func CheckDocumentFromBase64(base64String string, doc Document) (err error) {
	if err = DecodeBase64ToPb(base64String, doc); err != nil {
		return gogoerrors.AutoWrap(err)
	}
	if err = CheckText(doc); err != nil {
		return gogoerrors.AutoWrap(err)
	}
	defer func() {
		if r := recover(); r != nil {
			if rErr, ok := r.(error); ok {
				err = gogoerrors.AutoWrapSkip(fmt.Errorf("panic: %v", rErr), 1)
			} else {
				err = gogoerrors.AutoNewCustom(fmt.Sprintf("panic: %v", r), gogoerrors.PrependFullFuncName, 1)
			}
		}
	}()
	return gogoerrors.AutoWrap(checkSentenceSlice(reflect.ValueOf(doc)))
}

// checkSentenceSlice is a sub procedure of
// the function CheckDocumentFromBase64.
func checkSentenceSlice(docV reflect.Value) error {
	if kind := docV.Kind(); kind != reflect.Interface && kind != reflect.Pointer {
		return gogoerrors.AutoNew("doc is neither an interface nor a pointer")
	}
	m := docV.MethodByName("GetSentence")
	if !m.IsValid() {
		return gogoerrors.AutoNew("doc does not have method GetSentence")
	}
	retV := m.Call(nil)
	if n := len(retV); n != 1 {
		return gogoerrors.AutoNew(fmt.Sprintf(
			"doc.GetSentence returned %d values; want 1",
			n,
		))
	}
	sentSliceV := retV[0]
	if kind := sentSliceV.Kind(); kind != reflect.Slice {
		return gogoerrors.AutoNew(fmt.Sprintf(
			"doc.GetSentence returned a %v; want a slice",
			kind,
		))
	}
	if kind := sentSliceV.Type().Elem().Kind(); kind != reflect.Pointer {
		return gogoerrors.AutoNew(fmt.Sprintf(
			"doc.GetSentence returned a slice of %v; want a slice of pointer",
			kind,
		))
	}
	n := sentSliceV.Len()
	if n != RosesAreRedSentenceNum {
		return gogoerrors.AutoNew(fmt.Sprintf(
			"got %d sentence(s); want %d",
			n,
			RosesAreRedSentenceNum,
		))
	}
	for i := 0; i < n; i++ {
		if err := checkTokensOfSentence(sentSliceV, i); err != nil {
			return gogoerrors.AutoWrap(err)
		}
	}
	return nil
}

// checkTokensOfSentence is a sub procedure of the function checkSentenceSlice.
func checkTokensOfSentence(sentSliceV reflect.Value, sentIdx int) error {
	sentV := sentSliceV.Index(sentIdx)
	// The type of sentV has already been examined.
	m := sentV.MethodByName("GetToken")
	if !m.IsValid() {
		return gogoerrors.AutoNew(fmt.Sprintf(
			"sentence#%d does not have method GetToken",
			sentIdx,
		))
	}
	retV := m.Call(nil)
	if n := len(retV); n != 1 {
		return gogoerrors.AutoNew(fmt.Sprintf(
			"sentence#%d.GetToken returned %d values; want 1",
			sentIdx,
			n,
		))
	}
	tokenSliceV := retV[0]
	if kind := tokenSliceV.Kind(); kind != reflect.Slice {
		return gogoerrors.AutoNew(fmt.Sprintf(
			"sentence#%d.GetToken returned a %v; want a slice",
			sentIdx,
			kind,
		))
	}
	if kind := tokenSliceV.Type().Elem().Kind(); kind != reflect.Pointer {
		return gogoerrors.AutoNew(fmt.Sprintf(
			"sentence#%d.GetToken returned a slice of %v; want a slice of pointer",
			sentIdx,
			kind,
		))
	}
	n := tokenSliceV.Len()
	if n != RosesAreRedSentenceTokenNumList[sentIdx] {
		return gogoerrors.AutoNew(fmt.Sprintf(
			"got %d token(s) in sentence#%d; want %d",
			n,
			sentIdx,
			RosesAreRedSentenceTokenNumList[sentIdx],
		))
	}
	for i := 0; i < n; i++ {
		if err := checkToken(tokenSliceV, sentIdx, i); err != nil {
			return gogoerrors.AutoWrap(err)
		}
	}
	return nil
}

// checkToken is a sub procedure of the function checkTokensOfSentence.
func checkToken(tokenSliceV reflect.Value, sentIdx, tokenIdx int) error {
	tokenV := tokenSliceV.Index(tokenIdx)
	// The type of tokenV has already been examined.
	cases := []struct {
		method         string
		want           string
		comparePrefixN int
	}{
		{method: "GetWord", want: RosesAreRedSentenceTokenWordLists[sentIdx][tokenIdx]},
		{method: "GetBefore", want: RosesAreRedSentenceTokenGapLists[sentIdx][tokenIdx]},
		{method: "GetAfter", want: RosesAreRedSentenceTokenGapLists[sentIdx][tokenIdx+1]},
		{method: "GetPos", want: RosesAreRedSentenceTokenPosLists[sentIdx][tokenIdx], comparePrefixN: 2},
	}
	for _, c := range cases {
		if err := checkStringMethod(tokenV, c.method, c.want, c.comparePrefixN, sentIdx, tokenIdx); err != nil {
			return gogoerrors.AutoWrap(err)
		}
	}
	return nil
}

// checkStringMethod is a sub procedure of the function checkToken.
func checkStringMethod(tokenV reflect.Value, methodName, want string, comparePrefixN, sentIdx, tokenIdx int) error {
	m := tokenV.MethodByName(methodName)
	if !m.IsValid() {
		return gogoerrors.AutoNew(fmt.Sprintf(
			"token#%d in sentence#%d does not have method %s",
			tokenIdx,
			sentIdx,
			methodName,
		))
	}
	retV := m.Call(nil)
	if n := len(retV); n != 1 {
		return gogoerrors.AutoNew(fmt.Sprintf(
			"token#%d.%s in sentence#%d returned %d values; want 1",
			tokenIdx,
			methodName,
			sentIdx,
			n,
		))
	}
	strV := retV[0]
	if kind := strV.Kind(); kind != reflect.String {
		return gogoerrors.AutoNew(fmt.Sprintf(
			"token#%d.%s in sentence#%d returned a %v; want a string",
			tokenIdx,
			methodName,
			sentIdx,
			kind,
		))
	}
	str := strV.Interface().(string)
	if comparePrefixN <= 0 {
		if str != want {
			return gogoerrors.AutoNew(fmt.Sprintf(
				"token#%d.%s in sentence#%d returned %q; want %q",
				tokenIdx,
				methodName,
				sentIdx,
				str,
				want,
			))
		}
	} else {
		n := comparePrefixN
		if n > len(want) {
			n = len(want)
		}
		if n > len(str) {
			return gogoerrors.AutoNew(fmt.Sprintf(
				"token#%d.%s in sentence#%d returned %q, which is less than %d bytes",
				tokenIdx,
				methodName,
				sentIdx,
				str,
				n,
			))
		}
		if str[:n] != want[:n] {
			return gogoerrors.AutoNew(fmt.Sprintf(
				"token#%d.%s in sentence#%d returned %q; want a prefix %q",
				tokenIdx,
				methodName,
				sentIdx,
				str,
				want[:n],
			))
		}
	}
	return nil
}
