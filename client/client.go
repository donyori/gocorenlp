// gocorenlp.  A Go (Golang) client for Stanford CoreNLP server.
// Copyright (C) 2022-2024  Yuan Gao
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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	gogoerrors "github.com/donyori/gogo/errors"
	"google.golang.org/protobuf/proto"

	"github.com/donyori/gocorenlp/errors"
	"github.com/donyori/gocorenlp/model"
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
	//  import "github.com/donyori/gocorenlp/model/v4.5.6-eb50467fa8e3/pb"
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
	//  import "github.com/donyori/gocorenlp/model/v4.5.6-eb50467fa8e3/pb"
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
	AnnotateRaw(input io.Reader, annotators string, output io.Writer) (
		written int64, err error)

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
	AnnotateStringRaw(text, annotators string, output io.Writer) (
		written int64, err error)

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

// clientImpl is an implementation of the interface Client.
type clientImpl struct {
	c           http.Client
	host        string
	statusHost  string
	userinfo    *url.Userinfo
	annotators  string
	contentType string
	serverID    string
}

// newClientImpl creates a new clientImpl and
// sets its fields according to the specified options opt.
func newClientImpl(opt *Options) *clientImpl {
	if opt == nil {
		opt = new(Options)
	}
	c := &clientImpl{serverID: strings.TrimSpace(opt.ServerID)}
	c.host, c.statusHost = opt.GetHosts()
	username := strings.TrimSpace(opt.Username)
	if len(username) > 0 {
		password := strings.TrimSpace(opt.Password)
		if len(password) > 0 {
			c.userinfo = url.UserPassword(username, password)
		} else {
			c.userinfo = url.User(username)
		}
	}
	if opt.ClientTimeout > 0 {
		c.c.Timeout = opt.ClientTimeout
	}
	if len(opt.Annotators) > 0 {
		c.annotators = strings.Join(strings.Fields(opt.Annotators), "") // drop white space
	}
	charset := strings.TrimSpace(opt.Charset)
	if len(charset) == 0 {
		charset = "utf-8"
	}
	// According to the Stanford CoreNLP online documentation
	// <https://stanfordnlp.github.io/CoreNLP/corenlp-server.html#annotate-with-corenlp->,
	// the post data should be percent-encoded.
	// Thus, set the Content-Type header to "application/x-www-form-urlencoded".
	c.contentType = "application/x-www-form-urlencoded; charset=" + charset
	return c
}

func (c *clientImpl) Live() error {
	liveUrl := &url.URL{
		Scheme: "http",
		User:   c.userinfo,
		Host:   c.statusHost,
		Path:   "live",
	}
	resp, err := c.c.Get(liveUrl.String())
	if err != nil {
		return gogoerrors.AutoWrap(err)
	}
	defer func(c io.Closer) {
		_ = c.Close() // ignore error
	}(resp.Body)
	err = checkResponse(resp, "live")
	return gogoerrors.AutoWrap(err)
}

func (c *clientImpl) Ready() error {
	readyUrl := &url.URL{
		Scheme: "http",
		User:   c.userinfo,
		Host:   c.statusHost,
		Path:   "ready",
	}
	resp, err := c.c.Get(readyUrl.String())
	if err != nil {
		return gogoerrors.AutoWrap(err)
	}
	defer func(c io.Closer) {
		_ = c.Close() // ignore error
	}(resp.Body)
	err = checkResponse(resp, "ready")
	return gogoerrors.AutoWrap(err)
}

func (c *clientImpl) Annotate(
	input io.Reader,
	annotators string,
	outDoc proto.Message,
) error {
	var b bytes.Buffer
	if _, err := c.AnnotateRaw(input, annotators, &b); err != nil {
		return gogoerrors.AutoWrap(err)
	}
	// Parse ProtoBuf message.
	return gogoerrors.AutoWrap(model.DecodeResponseBody(b.Bytes(), outDoc))
}

func (c *clientImpl) AnnotateString(
	text, annotators string,
	outDoc proto.Message,
) error {
	return gogoerrors.AutoWrap(c.Annotate(
		strings.NewReader(text), annotators, outDoc))
}

func (c *clientImpl) AnnotateRaw(
	input io.Reader,
	annotators string,
	output io.Writer,
) (written int64, err error) {
	// Check arguments first.
	if input == nil {
		panic(gogoerrors.AutoMsg("input reader is nil"))
	}
	if output == nil {
		panic(gogoerrors.AutoMsg("output writer is nil"))
	}

	// Make request URL.
	ann := strings.Join(strings.Fields(annotators), "") // drop white space
	if len(ann) == 0 {
		ann = c.annotators
	}
	prop := make(map[string]string, 3)
	prop["outputFormat"] = "serialized"
	prop["serializer"] = "edu.stanford.nlp.pipeline.ProtobufAnnotationSerializer"
	if len(ann) > 0 {
		prop["annotators"] = ann
	}
	propBytes, err := json.Marshal(prop)
	if err != nil {
		// This should never happen.
		return 0, gogoerrors.AutoWrap(err)
	}
	qv := url.Values{"properties": []string{string(propBytes)}}
	annUrl := &url.URL{
		Scheme:   "http",
		User:     c.userinfo,
		Host:     c.host,
		RawQuery: qv.Encode(),
	}

	// Send request and forward response body to output.
	resp, err := c.c.Post(annUrl.String(), c.contentType, input)
	if err != nil {
		return 0, gogoerrors.AutoWrap(err)
	}
	defer func(c io.Closer) {
		_ = c.Close() // ignore error
	}(resp.Body)
	err = checkResponse(resp, "")
	if err != nil {
		return 0, gogoerrors.AutoWrap(err)
	}
	written, err = io.Copy(output, resp.Body)
	return written, gogoerrors.AutoWrap(err)
}

func (c *clientImpl) AnnotateStringRaw(
	text, annotators string,
	output io.Writer,
) (written int64, err error) {
	written, err = c.AnnotateRaw(strings.NewReader(text), annotators, output)
	return written, gogoerrors.AutoWrap(err)
}

func (c *clientImpl) Shutdown(key string) error {
	qv := url.Values{"key": []string{key}}
	shutdownUrl := &url.URL{
		Scheme:   "http",
		User:     c.userinfo,
		Host:     c.host,
		Path:     "shutdown",
		RawQuery: qv.Encode(),
	}
	resp, err := c.c.Get(shutdownUrl.String())
	if err != nil {
		return gogoerrors.AutoWrap(err)
	}
	defer func(c io.Closer) {
		_ = c.Close() // ignore error
	}(resp.Body)
	err = checkResponse(resp, "Shutdown successful!")
	return gogoerrors.AutoWrap(err)
}

func (c *clientImpl) ShutdownLocal() error {
	tmpDir := os.TempDir()
	name := filepath.Join(tmpDir, "corenlp.shutdown")
	if len(c.serverID) > 0 {
		name += "." + c.serverID
	}
	key, err := os.ReadFile(name)
	if err != nil {
		return gogoerrors.AutoWrap(fmt.Errorf(
			"failed to find the key: %w", err))
	}
	return gogoerrors.AutoWrap(c.Shutdown(string(key)))
}

func (c *clientImpl) private() {}

// checkResponse checks the status and body of the specified HTTP response.
//
// wantBody is the expected body of the response.
// If wantBody is not empty, checkResponse reads the response body
// and compares it with wantBody, but does not close the response body reader.
// Otherwise, checkResponse does nothing to the response body reader,
// neither reading nor closing it.
//
// checkResponse reports an error if the status code is not 2XX,
// or the expected body is not empty and the response body is different from
// the expected (leading and trailing whitespace characters are ignored).
// If the returned error is non-nil, it is of type
// *github.com/donyori/gocorenlp/errors.UnacceptableResponseError.
func checkResponse(resp *http.Response, wantBody string) error {
	var err error
	respErr := new(errors.UnacceptableResponseError)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respErr.StatusCode, respErr.Status = resp.StatusCode, resp.Status
		err = respErr
	}
	if wantBody == "" {
		return gogoerrors.AutoWrap(err)
	}
	var body []byte
	body, respErr.ReadError = io.ReadAll(resp.Body)
	if len(body) > 0 {
		respErr.Body = string(body)
	}
	if respErr.ReadError != nil {
		err = respErr
	} else if string(bytes.TrimSpace(body)) != strings.TrimSpace(wantBody) {
		respErr.WantBody = wantBody
		err = respErr
	}
	return gogoerrors.AutoWrap(err)
}
