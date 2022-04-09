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

package client

import (
	"fmt"
	"io/fs"
	"net/http"
	"net/url"
	"testing"

	"github.com/donyori/gogo/errors"
	"google.golang.org/protobuf/proto"

	"github.com/donyori/gocorenlp/model/v4.4.0-e90f30f13c40/pb"
)

func TestIsTimeoutError(t *testing.T) {
	testCases := []testIsErrorTestCase{
		{},
		{err: errors.New("common error")},
		{err: &testTimeoutError{false}},
		{err: &testTimeoutError{true}, want: true},
		{err: errors.AutoWrap(&testTimeoutError{false})},
		{err: errors.AutoWrap(&testTimeoutError{true}), want: true},
		{err: testWrapError(errors.AutoWrap(&testTimeoutError{false}))},
		{err: testWrapError(errors.AutoWrap(&testTimeoutError{true})), want: true},
	}
	testIsErrorFunc(t, IsTimeoutError, testCases)
}

func TestIsFileError(t *testing.T) {
	pathErr := &fs.PathError{
		Op:   "test",
		Path: "/",
		Err:  errors.New("path error"),
	}
	testCases := []testIsErrorTestCase{
		{},
		{err: errors.New("common error")},
		{err: pathErr, want: true},
		{err: errors.AutoWrap(pathErr), want: true},
		{err: testWrapError(errors.AutoWrap(pathErr)), want: true},
	}
	testIsErrorFunc(t, IsFileError, testCases)
}

func TestIsConnectionError(t *testing.T) {
	urlErr := &url.Error{
		Op:  "test",
		URL: "https://www.example.com/index.html",
		Err: errors.New("URL error"),
	}
	testCases := []testIsErrorTestCase{
		{},
		{err: errors.New("common error")},
		{err: urlErr, want: true},
		{err: errors.AutoWrap(urlErr), want: true},
		{err: testWrapError(errors.AutoWrap(urlErr)), want: true},
	}
	testIsErrorFunc(t, IsConnectionError, testCases)
}

func TestIsUnacceptableResponseError(t *testing.T) {
	upe := &UnacceptableResponseError{
		StatusCode: http.StatusNotFound,
		Status:     "404 Not Found (test status)",
		ReadError:  nil,
		Body:       "404 Not Found (test body)",
		WantBody:   "",
	}
	testCases := []testIsErrorTestCase{
		{},
		{err: errors.New("common error")},
		{err: upe, want: true},
		{err: errors.AutoWrap(upe), want: true},
		{err: testWrapError(errors.AutoWrap(upe)), want: true},
	}
	testIsErrorFunc(t, IsUnacceptableResponseError, testCases)
}

func TestIsProtoBufError(t *testing.T) {
	pe := &ProtoBufError{
		Op:   "test",
		Type: "test_type",
		Err:  errors.New("ProtoBuf error"),
	}
	testCases := []testIsErrorTestCase{
		{},
		{err: errors.New("common error")},
		{err: pe, want: true},
		{err: errors.AutoWrap(pe), want: true},
		{err: testWrapError(errors.AutoWrap(pe)), want: true},
	}
	testIsErrorFunc(t, IsProtoBufError, testCases)
}

func TestNewProtoBufError(t *testing.T) {
	docPtr := new(pb.Document)
	var msg proto.Message = docPtr
	docType := "github.com/donyori/gocorenlp/model/v4.4.0-e90f30f13c40/pb.Document"

	var st struct {
		testTimeoutError

		name string
		i    int
		err  error
		doc  *pb.Document
	}
	anonymousStructType := "struct { client.testTimeoutError; name string; i int; err error; doc *pb.Document }"

	underlyingErr := errors.New("ProtoBuf error")
	typeWantCases := []struct {
		showName string
		v        interface{}
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
					pe := NewProtoBufError(op, tw.v, ue)
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

// testWrapError wraps the specified error using
// github.com/donyori/gogo/errors.AutoWrap.
//
// It is useful for achieving wrapping an error within
// different functions, for example:
//  testWrapError(errors.AutoWrap(err))
func testWrapError(err error) error {
	return errors.AutoWrap(err)
}

// testIsErrorTestCase combines an error and a boolean value
// for use by TestIsXxxError.
type testIsErrorTestCase struct {
	err  error
	want bool
}

// testIsErrorFunc tests the specified IsXxxError function f
// with the specified test cases.
func testIsErrorFunc(t *testing.T, f func(err error) bool, testCases []testIsErrorTestCase) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("error=%v", tc.err), func(t *testing.T) {
			if r := f(tc.err); r != tc.want {
				t.Errorf("got %t; want %t", r, tc.want)
			}
		})
	}
}

// testTimeoutError is a timeout error for use by TestIsTimeoutError.
type testTimeoutError struct {
	timeout bool
}

func (e *testTimeoutError) Error() string {
	return "timeout (test)"
}

func (e *testTimeoutError) Timeout() bool {
	return e != nil && e.timeout
}
