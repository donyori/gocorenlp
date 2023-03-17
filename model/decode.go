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
// <https://developers.google.com/protocol-buffers/docs/encoding>.
func DecodeMessage(b []byte, msg proto.Message) error {
	if err := proto.Unmarshal(b, msg); err != nil {
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
// The specified message msg must be a non-nil pointer to a ProtoBuf message.
//
// If the returned error is non-nil, it has an underlying error of type
// *github.com/donyori/gocorenlp/errors.ProtoBufError.
// The function github.com/donyori/gocorenlp/errors.IsProtoBufError
// returns true for this error.
func DecodeResponseBody(b []byte, msg proto.Message) error {
	v, n := protowire.ConsumeBytes(b)
	var err error
	if n >= 0 {
		err = DecodeMessage(v, msg)
	} else {
		err = errors.NewProtoBufError(
			"google.golang.org/protobuf/encoding/protowire.ConsumeBytes",
			b,
			protowire.ParseError(n),
		)
	}
	return gogoerrors.AutoWrap(err)
}
