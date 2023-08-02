// gocorenlp.  A Go (Golang) client for Stanford CoreNLP server.
// Copyright (C) 2022-2023  Yuan Gao
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

// NumRosesAreRedSentence is the expected number of sentences for RosesAreRed.
const NumRosesAreRedSentence = 4

// Expected annotation results for RosesAreRed.
var (
	NumRosesAreRedSentenceTokenList   = [NumRosesAreRedSentence]int{4, 4, 4, 5}
	RosesAreRedSentenceTokenWordLists = [NumRosesAreRedSentence][]string{
		{"Roses", "are", "red", "."},
		{"Violets", "are", "blue", "."},
		{"Sugar", "is", "sweet", "."},
		{"And", "so", "are", "you", "."},
	}
	RosesAreRedSentenceTokenGapLists = [NumRosesAreRedSentence][]string{
		{"\n", " ", " ", "", "\n  "},
		{"\n  ", " ", " ", "", "\n"},
		{"\n", " ", " ", "", "\n  "},
		{"\n  ", " ", " ", " ", "", "\n"},
	}
	RosesAreRedSentenceTokenPosLists = [NumRosesAreRedSentence][]string{
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

// CheckRosesAreRedText examines whether the text returned by getter.GetText
// is the same as RosesAreRed.
//
// It reports an error if getter is nil or the text returned by
// getter.GetText is different from RosesAreRed.
func CheckRosesAreRedText(getter TextGetter) error {
	if getter == nil {
		return gogoerrors.AutoNew("text getter is nil")
	}
	if text := getter.GetText(); text != RosesAreRed {
		return gogoerrors.AutoNew(fmt.Sprintf(
			"got %q; want %q", text, RosesAreRed))
	}
	return nil
}

// CheckRosesAreRedDocument examines whether
// the annotation in doc is that of RosesAreRed.
func CheckRosesAreRedDocument(doc Document) (err error) {
	if err = CheckRosesAreRedText(doc); err != nil {
		return gogoerrors.AutoWrap(err)
	}
	defer func() {
		if r := recover(); r != nil {
			if rErr, ok := r.(error); ok {
				err = gogoerrors.AutoWrapSkip(
					fmt.Errorf("panic: %w", rErr), 1)
			} else {
				err = gogoerrors.AutoNewCustom(
					fmt.Sprintf("panic: %v", r), -1, 1)
			}
		}
	}()
	return gogoerrors.AutoWrap(
		checkRosesAreRedSentenceSlice(reflect.ValueOf(doc)))
}

// DecodeBase64ToPb decodes the standard base64 (as defined in RFC 4648) encoded CoreNLP server
// response body into a ProtoBuf message.
func DecodeBase64ToPb(base64String string, msg proto.Message) error {
	b, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return gogoerrors.AutoWrap(err)
	}
	return gogoerrors.AutoWrap(model.DecodeResponseBody(b, msg))
}

// CheckRosesAreRedDocumentFromBase64 decodes base64String to doc
// and then examines whether the annotation in doc is that of RosesAreRed.
func CheckRosesAreRedDocumentFromBase64(base64String string, doc Document) (
	err error) {
	if err = DecodeBase64ToPb(base64String, doc); err != nil {
		return gogoerrors.AutoWrap(err)
	}
	return gogoerrors.AutoWrap(CheckRosesAreRedDocument(doc))
}

// checkRosesAreRedSentenceSlice is a sub procedure of
// the function CheckRosesAreRedDocumentFromBase64.
func checkRosesAreRedSentenceSlice(docV reflect.Value) error {
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
	if n != NumRosesAreRedSentence {
		return gogoerrors.AutoNew(fmt.Sprintf(
			"got %d sentence(s); want %d",
			n,
			NumRosesAreRedSentence,
		))
	}
	for i := 0; i < n; i++ {
		if err := checkRosesAreRedTokensOfSentence(sentSliceV, i); err != nil {
			return gogoerrors.AutoWrap(err)
		}
	}
	return nil
}

// checkRosesAreRedTokensOfSentence is a sub procedure of
// the function checkRosesAreRedSentenceSlice.
func checkRosesAreRedTokensOfSentence(sentSliceV reflect.Value, sentIdx int) error {
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
	if n != NumRosesAreRedSentenceTokenList[sentIdx] {
		return gogoerrors.AutoNew(fmt.Sprintf(
			"got %d token(s) in sentence#%d; want %d",
			n,
			sentIdx,
			NumRosesAreRedSentenceTokenList[sentIdx],
		))
	}
	for i := 0; i < n; i++ {
		if err := checkRosesAreRedToken(tokenSliceV, sentIdx, i); err != nil {
			return gogoerrors.AutoWrap(err)
		}
	}
	return nil
}

// checkRosesAreRedToken is a sub procedure of
// the function checkRosesAreRedTokensOfSentence.
func checkRosesAreRedToken(
	tokenSliceV reflect.Value,
	sentIdx int,
	tokenIdx int,
) error {
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
		err := checkStringMethod(
			tokenV,
			c.method,
			c.want,
			c.comparePrefixN,
			sentIdx,
			tokenIdx,
		)
		if err != nil {
			return gogoerrors.AutoWrap(err)
		}
	}
	return nil
}

// checkStringMethod is a sub procedure of the function checkRosesAreRedToken.
func checkStringMethod(
	tokenV reflect.Value,
	methodName string,
	want string,
	comparePrefixN int,
	sentIdx int,
	tokenIdx int,
) error {
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
