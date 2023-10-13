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

package errors_test

import (
	"fmt"
	"io/fs"
	"net/http"
	"net/url"
	"strings"
	"testing"

	gogoerrors "github.com/donyori/gogo/errors"
	"github.com/donyori/gogo/runtime"
	"google.golang.org/protobuf/proto"

	"github.com/donyori/gocorenlp/errors"
	"github.com/donyori/gocorenlp/model/v4.5.5-f1b929e47a57/pb"
)

var PackageSimpleName string

func init() {
	// Make sure that none of the five functions exported from
	// the standard package errors and
	// github.com/donyori/gogo/errors is missing.
	_ = errors.New("")
	_ = errors.Unwrap(nil)
	_ = errors.Is(nil, nil)
	var te *TimeoutError
	_ = errors.As(nil, &te)
	_ = errors.Join()

	// Dynamically get the package name to avoid inappropriate test cases,
	// especially when moving the test to another package.
	pkg, _, ok := runtime.CallerPkgFunc(0)
	if !ok {
		panic("failed to retrieve package name")
	}
	PackageSimpleName = pkg[strings.LastIndexByte(pkg, '/')+1:]
}

func TestIsTimeoutError(t *testing.T) {
	testCases := []IsErrorTestCase{
		{},
		{err: gogoerrors.New("common error")},
		{err: &TimeoutError{false}},
		{err: &TimeoutError{true}, want: true},
		{err: gogoerrors.AutoWrap(&TimeoutError{false})},
		{err: gogoerrors.AutoWrap(&TimeoutError{true}), want: true},
		{err: WrapError(gogoerrors.AutoWrap(&TimeoutError{false}))},
		{err: WrapError(gogoerrors.AutoWrap(&TimeoutError{true})), want: true},
	}
	IsErrorFunc(t, errors.IsTimeoutError, testCases)
}

func TestIsFileError(t *testing.T) {
	pathErr := &fs.PathError{
		Op:   "test",
		Path: "/",
		Err:  gogoerrors.New("path error"),
	}
	testCases := []IsErrorTestCase{
		{},
		{err: gogoerrors.New("common error")},
		{err: pathErr, want: true},
		{err: gogoerrors.AutoWrap(pathErr), want: true},
		{err: WrapError(gogoerrors.AutoWrap(pathErr)), want: true},
	}
	IsErrorFunc(t, errors.IsFileError, testCases)
}

func TestIsConnectionError(t *testing.T) {
	urlErr := &url.Error{
		Op:  "test",
		URL: "https://www.example.com/index.html",
		Err: gogoerrors.New("URL error"),
	}
	testCases := []IsErrorTestCase{
		{},
		{err: gogoerrors.New("common error")},
		{err: urlErr, want: true},
		{err: gogoerrors.AutoWrap(urlErr), want: true},
		{err: WrapError(gogoerrors.AutoWrap(urlErr)), want: true},
	}
	IsErrorFunc(t, errors.IsConnectionError, testCases)
}

func TestIsUnacceptableResponseError(t *testing.T) {
	upe := &errors.UnacceptableResponseError{
		StatusCode: http.StatusNotFound,
		Status:     "404 Not Found (test status)",
		ReadError:  nil,
		Body:       "404 Not Found (test body)",
		WantBody:   "",
	}
	testCases := []IsErrorTestCase{
		{},
		{err: gogoerrors.New("common error")},
		{err: upe, want: true},
		{err: gogoerrors.AutoWrap(upe), want: true},
		{err: WrapError(gogoerrors.AutoWrap(upe)), want: true},
	}
	IsErrorFunc(t, errors.IsUnacceptableResponseError, testCases)
}

func TestIsProtoBufError(t *testing.T) {
	pe := &errors.ProtoBufError{
		Op:   "test",
		Type: "test_type",
		Err:  gogoerrors.New("ProtoBuf error"),
	}
	testCases := []IsErrorTestCase{
		{},
		{err: gogoerrors.New("common error")},
		{err: pe, want: true},
		{err: gogoerrors.AutoWrap(pe), want: true},
		{err: WrapError(gogoerrors.AutoWrap(pe)), want: true},
	}
	IsErrorFunc(t, errors.IsProtoBufError, testCases)
}

func TestNewProtoBufError(t *testing.T) {
	docPtr := new(pb.Document)
	var msg proto.Message = docPtr
	docType := "github.com/donyori/gocorenlp/model/v4.5.5-f1b929e47a57/pb.Document"

	var st struct {
		TimeoutError

		name string
		i    int
		err  error
		doc  *pb.Document
	}
	anonymousStructType := "struct { " + PackageSimpleName +
		".TimeoutError; name string; i int; err error; doc *pb.Document }"

	underlyingErr := gogoerrors.New("ProtoBuf error")
	typeWantCases := []struct {
		showName string
		v        any
		want     string
	}{
		{"<nil>", nil, "<nil>"},
		{"byte slice", []byte{}, "[]uint8"},
		{"proto.Message", msg, docType},
		{"*proto.Message", &msg, docType},
		{"pb.Document", *docPtr, docType},
		{"*pb.Document", docPtr, docType},
		{"anonymous struct", st, anonymousStructType},
		{"*anonymous struct", &st, anonymousStructType},
	}
	for _, op := range []string{"client.testOp1", "test_op2"} {
		for _, tw := range typeWantCases {
			for _, ue := range []error{nil, underlyingErr} {
				t.Run(fmt.Sprintf("op=%s&v type=%s&err=%v", op, tw.showName, ue), func(t *testing.T) {
					pe := errors.NewProtoBufError(op, tw.v, ue)
					if pe == nil {
						t.Fatal("got nil")
					}
					if pe.Op != op {
						t.Errorf("got op %s; want %s", pe.Op, op)
					}
					if pe.Type != tw.want {
						t.Errorf("got type %s; want %s", pe.Type, tw.want)
					}
					if pe.Err != ue {
						t.Errorf("got err %v; want %v", pe.Err, ue)
					}
				})
			}
		}
	}
}

// WrapError wraps the specified error using
// github.com/donyori/gogo/errors.AutoWrap.
//
// It is useful for achieving wrapping an error within
// different functions, for example:
//
//	WrapError(errors.AutoWrap(err))
func WrapError(err error) error {
	return gogoerrors.AutoWrap(err)
}

// IsErrorTestCase combines an error and a boolean value
// for use by TestIsXxxError.
type IsErrorTestCase struct {
	err  error
	want bool
}

// IsErrorFunc tests the specified IsXxxError function f
// with the specified test cases.
func IsErrorFunc(t *testing.T, f func(err error) bool, testCases []IsErrorTestCase) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("error=%v", tc.err), func(t *testing.T) {
			if r := f(tc.err); r != tc.want {
				t.Errorf("got %t; want %t", r, tc.want)
			}
		})
	}
}

// TimeoutError is a timeout error for use by TestIsTimeoutError.
//
// It is a private type for testing unexported struct.
type TimeoutError struct {
	timeout bool
}

func (e *TimeoutError) Error() string {
	return "timeout (test)"
}

func (e *TimeoutError) Timeout() bool {
	return e != nil && e.timeout
}
