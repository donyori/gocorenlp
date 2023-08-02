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

package model

import (
	gogoerrors "github.com/donyori/gogo/errors"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"

	"github.com/donyori/gocorenlp/errors"
)

// DecodeMessage parses a ProtoBuf message from wire encoding in b
// and stores the result in msg.
//
// To parse messages from a Stanford CoreNLP server response,
// use the function DecodeResponseBody instead.
//
// The specified message msg must be a non-nil pointer to a ProtoBuf message.
//
// If the returned error is non-nil, it has an underlying error of type
// *github.com/donyori/gocorenlp/errors.ProtoBufError.
// The function github.com/donyori/gocorenlp/errors.IsProtoBufError
// returns true for this error.
//
// For more information about the wire encoding, see
// <https://protobuf.dev/programming-guides/encoding/>.
func DecodeMessage(b []byte, msg proto.Message) error {
	if msg == nil {
		return gogoerrors.AutoNew("the provided message is nil")
	}
	err := proto.Unmarshal(b, msg)
	if err != nil {
		return gogoerrors.AutoWrap(errors.NewProtoBufError(
			"google.golang.org/protobuf/proto.Unmarshal",
			msg,
			err,
		))
	}
	return nil
}

// DecodeResponseBody parses a ProtoBuf message (usually a Stanford CoreNLP
// document) from the body of a Stanford CoreNLP server response in b,
// and stores the result in msg.
//
// If there are other contents after the response body in b,
// these contents are ignored.
//
// The specified message msg must be a non-nil pointer to a ProtoBuf message.
//
// If the returned error is non-nil, it has an underlying error of type
// *github.com/donyori/gocorenlp/errors.ProtoBufError.
// The function github.com/donyori/gocorenlp/errors.IsProtoBufError
// returns true for this error.
func DecodeResponseBody(b []byte, msg proto.Message) error {
	_, err := ConsumeResponseBody(b, msg)
	return gogoerrors.AutoWrap(err)
}

// ConsumeResponseBody is like the function DecodeResponseBody,
// but it additionally returns the number of bytes decoded from b.
// It enables the caller to determine the range of the response body (b[:n]).
//
// ConsumeResponseBody may return a negative n upon an error.
// A negative n can be converted into an error value using
// google.golang.org/protobuf/encoding/protowire.ParseError.
//
// If n is negative, the returned error err must be non-nil,
// have an underlying error of type
// *github.com/donyori/gocorenlp/errors.ProtoBufError,
// and wrap the protowire.ParseError(n).
// The function github.com/donyori/gocorenlp/errors.IsProtoBufError
// returns true for this error.
func ConsumeResponseBody(b []byte, msg proto.Message) (n int, err error) {
	if msg == nil {
		return 0, gogoerrors.AutoNew("the provided message is nil")
	}
	v, n := protowire.ConsumeBytes(b)
	if n >= 0 {
		err = proto.Unmarshal(v, msg)
		if err != nil {
			err = errors.NewProtoBufError(
				"google.golang.org/protobuf/proto.Unmarshal",
				msg,
				err,
			)
		}
	} else {
		err = errors.NewProtoBufError(
			"google.golang.org/protobuf/encoding/protowire.ConsumeBytes",
			b,
			protowire.ParseError(n),
		)
	}
	return n, gogoerrors.AutoWrap(err)
}
