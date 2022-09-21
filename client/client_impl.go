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

// clientImpl is an implementation of the interface Client.
type clientImpl struct {
	c           http.Client
	host        string
	statusHost  string
	userinfo    *url.Userinfo
	annotators  string
	contentType string
	serverId    string
}

// newClientImpl creates a new clientImpl and
// sets its fields according to the specified options opt.
func newClientImpl(opt *Options) *clientImpl {
	if opt == nil {
		opt = new(Options)
	}
	c := &clientImpl{serverId: strings.TrimSpace(opt.ServerId)}
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
	if opt.Timeout > 0 {
		c.c.Timeout = opt.Timeout
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

// Live sends a status request to the liveness endpoint (/live) and
// reports any error encountered to check whether the target server
// is online.
//
// It returns nil if the server is online.
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
	_, _, err = checkResponse(resp, nil, nil, "live")
	return gogoerrors.AutoWrap(err)
}

// Ready sends a status request to the readiness endpoint (/ready) and
// reports any error encountered to check whether the target server
// is ready to accept connections.
//
// It returns nil if the server is ready to accept connections.
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
	_, _, err = checkResponse(resp, nil, nil, "ready")
	return gogoerrors.AutoWrap(err)
}

// Annotate sends an annotation request with the specified annotators
// to annotate the data read from the specified reader.
// The annotation result is represented as
// a CoreNLP document and stored in outDoc.
//
// If no annotators are specified,
// the client's default annotators will be used.
// If the client's annotators are also not specified,
// the server's default annotators will be used.
//
// The annotators are separated by commas (,) in the string without spaces.
// For example:
//
//	"tokenize,ssplit,pos,depparse"
//
// outDoc must be a non-nil pointer to an auto-generated Document
// structure, for example:
//
//	import "github.com/donyori/gocorenlp/model/v4.5.0-45b47e245c36/pb"
//	...
//	outDoc := new(pb.Document)
//	err := Annotate(input, "tokenize,ssplit,pos", outDoc)
//	...
//
// If outDoc is nil or not a pointer to Document,
// a runtime error will occur.
func (c *clientImpl) Annotate(input io.Reader, annotators string, outDoc proto.Message) error {
	var b bytes.Buffer
	if _, err := c.AnnotateRaw(input, annotators, &b); err != nil {
		return gogoerrors.AutoWrap(err)
	}
	// Parse ProtoBuf message.
	return gogoerrors.AutoWrap(model.DecodeResponseBody(b.Bytes(), outDoc))
}

// AnnotateString sends an annotation request with
// the specified text and annotators.
// The annotation result is represented as
// a CoreNLP document and stored in outDoc.
//
// If no annotators are specified,
// the client's default annotators will be used.
// If the client's annotators are also not specified,
// the server's default annotators will be used.
//
// The annotators are separated by commas (,) in the string without spaces.
// For example:
//
//	"tokenize,ssplit,pos,depparse"
//
// outDoc must be a non-nil pointer to an auto-generated Document
// structure, for example:
//
//	import "github.com/donyori/gocorenlp/model/v4.5.0-45b47e245c36/pb"
//	...
//	outDoc := new(pb.Document)
//	err := AnnotateString("Hello world!", "tokenize,ssplit,pos", outDoc)
//	...
//
// If outDoc is nil or not a pointer to Document,
// a runtime error will occur.
func (c *clientImpl) AnnotateString(text, annotators string, outDoc proto.Message) error {
	return gogoerrors.AutoWrap(c.Annotate(strings.NewReader(text), annotators, outDoc))
}

// AnnotateRaw sends an annotation request with the specified annotators
// to annotate the data read from the specified reader.
// Then AnnotateRaw writes the response body to the specified writer
// without parsing. The user can parse it later using the function
// github.com/donyori/gocorenlp/model.DecodeResponseBody.
//
// If no annotators are specified,
// the client's default annotators will be used.
// If the client's annotators are also not specified,
// the server's default annotators will be used.
//
// The annotators are separated by commas (,) in the string without spaces.
// For example:
//
//	"tokenize,ssplit,pos,depparse"
//
// It returns the number of bytes written and any error encountered.
func (c *clientImpl) AnnotateRaw(input io.Reader, annotators string, output io.Writer) (written int64, err error) {
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
	_, _, err = checkResponse(resp, nil, nil, "")
	if err != nil {
		return 0, gogoerrors.AutoWrap(err)
	}
	// The return value read of checkResponse must be false
	// as its acceptBody is nil, wantBody is empty, and it reports no error
	// (i.e., the response status is acceptable).
	defer closeIgnoreError(resp.Body)
	written, err = io.Copy(output, resp.Body)
	if err != nil {
		err = gogoerrors.AutoWrap(err)
	}
	return
}

// AnnotateStringRaw sends an annotation request with
// the specified text and annotators.
// Then AnnotateStringRaw writes the response body to
// the specified writer without parsing.
// The user can parse it later using the function
// github.com/donyori/gocorenlp/model.DecodeResponseBody.
//
// If no annotators are specified,
// the client's default annotators will be used.
// If the client's annotators are also not specified,
// the server's default annotators will be used.
//
// The annotators are separated by commas (,) in the string without spaces.
// For example:
//
//	"tokenize,ssplit,pos,depparse"
//
// It returns the number of bytes written and any error encountered.
func (c *clientImpl) AnnotateStringRaw(text, annotators string, output io.Writer) (written int64, err error) {
	written, err = c.AnnotateRaw(strings.NewReader(text), annotators, output)
	return written, gogoerrors.AutoWrap(err)
}

// Shutdown sends a shutdown request with the specified key
// to stop the target server.
//
// It returns nil if the server has been shut down successfully.
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
	wantBody := "Shutdown successful!"
	_, _, err = checkResponse(resp, nil, nil, wantBody)
	return gogoerrors.AutoWrap(err)
}

// ShutdownLocal finds the shutdown key and then sends
// a shutdown request to stop the target server.
//
// It works only if the target server is on the local.
//
// It returns nil if the server has been shut down successfully.
func (c *clientImpl) ShutdownLocal() error {
	tmpDir := os.TempDir()
	name := filepath.Join(tmpDir, "corenlp.shutdown")
	if len(c.serverId) > 0 {
		name += "." + c.serverId
	}
	key, err := os.ReadFile(name)
	if err != nil {
		return gogoerrors.AutoWrap(fmt.Errorf("failed to find the key: %v", err))
	}
	return gogoerrors.AutoWrap(c.Shutdown(string(key)))
}

func (c *clientImpl) private() {}

// checkResponse checks the status and body of the specified HTTP response.
//
// acceptStatus is a function to check the response status.
// Its two arguments are resp.StatusCode and resp.Status.
// It returns a boolean value indicating whether the status is acceptable.
// If it is nil, checkResponse accepts the status if the status code is 2XX.
//
// acceptBody is a function to check the response body.
// Its argument is the response body.
// It returns a boolean value indicating whether the body is acceptable.
// If it is nil and wantBody is not empty,
// checkResponse accepts the body if the body and wantBody are the same
// after dropping leading and trailing white space.
//
// wantBody is the expected body of the response.
// It will be stored in the returned error
// if the response body is unacceptable.
//
// If acceptBody is non-nil or wantBody is non-empty,
// or the status is unacceptable, checkResponse reads the response body
// and closes the body reader; otherwise, checkResponse does nothing to
// the body reader, neither reading nor closing it.
//
// checkResponse returns the response body (nil if not read it)
// and an indicator read to report whether it has attempted to
// read the response body.
// If read is true, the body reader is closed by checkResponse.
// It reports an error if the response is not as expected.
// If the returned error is non-nil, it is of type
// *github.com/donyori/gocorenlp/errors.UnacceptableResponseError.
func checkResponse(
	resp *http.Response,
	acceptStatus func(statusCode int, status string) bool,
	acceptBody func(body []byte) bool,
	wantBody string,
) (body []byte, read bool, err error) {
	if acceptStatus == nil {
		acceptStatus = func(statusCode int, _ string) bool {
			return statusCode >= 200 && statusCode < 300
		}
	}
	if acceptBody == nil && len(wantBody) > 0 {
		acceptBody = func(body []byte) bool {
			return strings.TrimSpace(string(body)) == strings.TrimSpace(wantBody)
		}
	}
	respErr := new(errors.UnacceptableResponseError)
	statusCode, status := resp.StatusCode, resp.Status
	if !acceptStatus(statusCode, status) {
		respErr.StatusCode, respErr.Status = statusCode, status
		err = respErr
	} else if acceptBody == nil {
		return
	}
	defer closeIgnoreError(resp.Body)
	read = true
	body, respErr.ReadError = io.ReadAll(resp.Body)
	if len(body) > 0 {
		respErr.Body = string(body)
	}
	if respErr.ReadError != nil {
		err = respErr
		return
	}
	if acceptBody != nil && !acceptBody(body) {
		err = respErr
		respErr.WantBody = wantBody
	}
	return
}

// closeIgnoreError closes the specified closer if it is non-nil
// and ignores any error encountered.
func closeIgnoreError(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}
