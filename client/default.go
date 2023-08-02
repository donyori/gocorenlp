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

// defaultClient is a client with default settings.
var defaultClient = newClientImpl(nil)

// Live is a wrapper around Client.Live with a default client.
// The default client connects to 127.0.0.1:9000,
// using all default settings (see Options for details).
//
// Live sends a status request to the liveness endpoint (/live) and
// reports any error encountered to check whether the target server
// is online.
//
// It returns nil if the server is online.
func Live() error {
	return gogoerrors.AutoWrap(defaultClient.Live())
}

// Ready is a wrapper around Client.Ready with a default client.
// The default client connects to 127.0.0.1:9000,
// using all default settings (see Options for details).
//
// Ready sends a status request to the readiness endpoint (/ready) and
// reports any error encountered to check whether the target server
// is ready to accept connections.
//
// It returns nil if the server is ready to accept connections.
func Ready() error {
	return gogoerrors.AutoWrap(defaultClient.Ready())
}

// Annotate is a wrapper around Client.Annotate with a default client.
// The default client connects to 127.0.0.1:9000,
// using all default settings (see Options for details).
//
// Annotate sends an annotation request with the specified annotators
// to annotate the data read from the specified reader.
// The annotation result is represented as
// a CoreNLP document and stored in outDoc.
//
// If no annotators are specified,
// the server's default annotators are used.
//
// The annotators are separated by commas (,) in the string without spaces.
// For example:
//
//	"tokenize,ssplit,pos,depparse"
//
// outDoc must be a non-nil pointer to an auto-generated Document
// structure, for example:
//
//	import "github.com/donyori/gocorenlp/model/v4.5.3-5250f9faf9f1/pb"
//	...
//	outDoc := new(pb.Document)
//	err := Annotate(input, "tokenize,ssplit,pos", outDoc)
//	...
//
// If outDoc is nil or not a pointer to Document, a runtime error occurs.
func Annotate(input io.Reader, annotators string, outDoc proto.Message) error {
	return gogoerrors.AutoWrap(
		defaultClient.Annotate(input, annotators, outDoc))
}

// AnnotateString is a wrapper around
// Client.AnnotateString with a default client.
// The default client connects to 127.0.0.1:9000,
// using all default settings (see Options for details).
//
// AnnotateString sends an annotation request with
// the specified text and annotators.
// The annotation result is represented as
// a CoreNLP document and stored in outDoc.
//
// If no annotators are specified,
// the server's default annotators are used.
//
// The annotators are separated by commas (,) in the string without spaces.
// For example:
//
//	"tokenize,ssplit,pos,depparse"
//
// outDoc must be a non-nil pointer to an auto-generated Document
// structure, for example:
//
//	import "github.com/donyori/gocorenlp/model/v4.5.3-5250f9faf9f1/pb"
//	...
//	outDoc := new(pb.Document)
//	err := AnnotateString("Hello world!", "tokenize,ssplit,pos", outDoc)
//	...
//
// If outDoc is nil or not a pointer to Document, a runtime error occurs.
func AnnotateString(text, annotators string, outDoc proto.Message) error {
	return gogoerrors.AutoWrap(
		defaultClient.AnnotateString(text, annotators, outDoc))
}

// AnnotateRaw is a wrapper around Client.AnnotateRaw with a default client.
// The default client connects to 127.0.0.1:9000,
// using all default settings (see Options for details).
//
// AnnotateRaw sends an annotation request with the specified annotators
// to annotate the data read from the specified reader.
// Then AnnotateRaw writes the response body to the specified writer
// without parsing. The user can parse it later using the function
// github.com/donyori/gocorenlp/model.DecodeResponseBody.
//
// If no annotators are specified,
// the server's default annotators are used.
//
// The annotators are separated by commas (,) in the string without spaces.
// For example:
//
//	"tokenize,ssplit,pos,depparse"
//
// It returns the number of bytes written and any error encountered.
func AnnotateRaw(input io.Reader, annotators string, output io.Writer) (
	written int64, err error) {
	written, err = defaultClient.AnnotateRaw(input, annotators, output)
	return written, gogoerrors.AutoWrap(err)
}

// AnnotateStringRaw is a wrapper around
// Client.AnnotateStringRaw with a default client.
// The default client connects to 127.0.0.1:9000,
// using all default settings (see Options for details).
//
// AnnotateStringRaw sends an annotation request with
// the specified text and annotators.
// Then AnnotateStringRaw writes the response body to
// the specified writer without parsing.
// The user can parse it later using the function
// github.com/donyori/gocorenlp/model.DecodeResponseBody.
//
// If no annotators are specified,
// the server's default annotators are used.
//
// The annotators are separated by commas (,) in the string without spaces.
// For example:
//
//	"tokenize,ssplit,pos,depparse"
//
// It returns the number of bytes written and any error encountered.
func AnnotateStringRaw(text, annotators string, output io.Writer) (
	written int64, err error) {
	written, err = defaultClient.AnnotateStringRaw(text, annotators, output)
	return written, gogoerrors.AutoWrap(err)
}

// Shutdown is a wrapper around Client.ShutdownLocal with a default client.
// The default client connects to 127.0.0.1:9000,
// using all default settings (see Options for details).
//
// Shutdown finds the shutdown key and then sends
// a shutdown request to stop the target server.
//
// It returns nil if the server has been shut down successfully.
func Shutdown() error {
	return gogoerrors.AutoWrap(defaultClient.ShutdownLocal())
}
