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

package client

import (
	"io"

	gogoerrors "github.com/donyori/gogo/errors"
	"google.golang.org/protobuf/proto"
)

// Client is an interface representing an HTTP client
// for the Stanford CoreNLP server.
type Client interface {
	// Live sends a status request to the liveness endpoint (/live) and
	// reports any error encountered to check whether the target server
	// is online.
	//
	// It returns nil if the server is online.
	Live() error

	// Ready sends a status request to the readiness endpoint (/ready) and
	// reports any error encountered to check whether the target server
	// is ready to accept connections.
	//
	// It returns nil if the server is ready to accept connections.
	Ready() error

	// Annotate sends an annotation request with the specified annotators
	// to annotate the data read from the specified reader.
	// The annotation result is represented as
	// a CoreNLP document and stored in outDoc.
	//
	// If no annotators are specified,
	// the client's default annotators are used.
	// If the client's annotators are also not specified,
	// the server's default annotators are used.
	//
	// The annotators are separated by commas (,) in the string without spaces.
	// For example:
	//  "tokenize,ssplit,pos,depparse"
	//
	// outDoc must be a non-nil pointer to an auto-generated Document
	// structure, for example:
	//
	//  import "github.com/donyori/gocorenlp/model/v4.5.0-45b47e245c36/pb"
	//  ...
	//  outDoc := new(pb.Document)
	//  err := Annotate(input, "tokenize,ssplit,pos", outDoc)
	//  ...
	//
	// If outDoc is nil or not a pointer to Document, a runtime error occurs.
	Annotate(input io.Reader, annotators string, outDoc proto.Message) error

	// AnnotateString sends an annotation request with
	// the specified text and annotators.
	// The annotation result is represented as
	// a CoreNLP document and stored in outDoc.
	//
	// If no annotators are specified,
	// the client's default annotators are used.
	// If the client's annotators are also not specified,
	// the server's default annotators are used.
	//
	// The annotators are separated by commas (,) in the string without spaces.
	// For example:
	//  "tokenize,ssplit,pos,depparse"
	//
	// outDoc must be a non-nil pointer to an auto-generated Document
	// structure, for example:
	//
	//  import "github.com/donyori/gocorenlp/model/v4.5.0-45b47e245c36/pb"
	//  ...
	//  outDoc := new(pb.Document)
	//  err := AnnotateString("Hello world!", "tokenize,ssplit,pos", outDoc)
	//  ...
	//
	// If outDoc is nil or not a pointer to Document, a runtime error occurs.
	AnnotateString(text, annotators string, outDoc proto.Message) error

	// AnnotateRaw sends an annotation request with the specified annotators
	// to annotate the data read from the specified reader.
	// Then AnnotateRaw writes the response body to the specified writer
	// without parsing. The user can parse it later using the function
	// github.com/donyori/gocorenlp/model.DecodeResponseBody.
	//
	// If no annotators are specified,
	// the client's default annotators are used.
	// If the client's annotators are also not specified,
	// the server's default annotators are used.
	//
	// The annotators are separated by commas (,) in the string without spaces.
	// For example:
	//  "tokenize,ssplit,pos,depparse"
	//
	// It returns the number of bytes written and any error encountered.
	AnnotateRaw(input io.Reader, annotators string, output io.Writer) (written int64, err error)

	// AnnotateStringRaw sends an annotation request with
	// the specified text and annotators.
	// Then AnnotateStringRaw writes the response body to
	// the specified writer without parsing.
	// The user can parse it later using the function
	// github.com/donyori/gocorenlp/model.DecodeResponseBody.
	//
	// If no annotators are specified,
	// the client's default annotators are used.
	// If the client's annotators are also not specified,
	// the server's default annotators are used.
	//
	// The annotators are separated by commas (,) in the string without spaces.
	// For example:
	//  "tokenize,ssplit,pos,depparse"
	//
	// It returns the number of bytes written and any error encountered.
	AnnotateStringRaw(text, annotators string, output io.Writer) (written int64, err error)

	// Shutdown sends a shutdown request with the specified key
	// to stop the target server.
	//
	// It returns nil if the server has been shut down successfully.
	Shutdown(key string) error

	// ShutdownLocal finds the shutdown key and then sends
	// a shutdown request to stop the target server.
	//
	// It works only if the target server is on the local.
	//
	// It returns nil if the server has been shut down successfully.
	ShutdownLocal() error

	// private prevents others from implementing this interface,
	// so future additions to it will not violate compatibility.
	private()
}

// New creates a new Client for the Stanford CoreNLP server
// with the specified options.
//
// If opt is nil, it uses default options.
//
// Before returning the client, it tests whether the target server is live.
// If the test fails, it reports an error and returns a nil client.
// Thus, make sure the server is online and set the appropriate host address
// in opt before calling this function.
func New(opt *Options) (c Client, err error) {
	t := newClientImpl(opt)
	if err := t.Live(); err != nil {
		return nil, gogoerrors.AutoWrap(err)
	}
	return t, nil
}
