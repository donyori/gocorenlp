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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/netip"
	"net/url"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"

	"github.com/donyori/gogo/errors"
)

// Client is an interface representing an HTTP client
// for the Stanford CoreNLP server.
type Client interface {
	// Ready sends a status request to the readiness endpoint (/ready) and
	// reports whether the target server is ready to accept connections.
	//
	// It returns false if the server is not ready or any error is encountered.
	Ready() bool

	// Annotate sends an annotation request with specified text and annotators.
	// The annotation result is represented as a CoreNLP document and
	// written to outDoc.
	//
	// outDoc must be a non-nil pointer to an auto-generated Document
	// structure, for example:
	//
	//  import "github.com/donyori/gocorenlp/model/v4.4.0-e90f30f13c40/pb"
	//  ...
	//  outDoc := new(pb.Document)
	//  err := Annotate(outDoc, "Hello world!", "tokenize,ssplit,pos")
	//  ...
	//
	// If outDoc is nil or not a pointer to Document,
	// a runtime error will occur.
	//
	// If no annotators are specified,
	// the client's default annotators will be used.
	// If the client's annotators are also not specified,
	// the server's default annotators will be used.
	//
	// The annotators are separated by commas (,) in the string without spaces.
	// For example:
	//  tokenize,ssplit,pos,depparse
	Annotate(outDoc proto.Message, text string, annotators string) error
}

// Options are the configuration for creating a new client.
type Options struct {
	// Hostname is the host (without port number) of the target server.
	//
	// Default: 127.0.0.1
	Hostname string

	// Port is the port of the target server.
	//
	// Default: 9000
	Port uint16

	// StatusPort is the port of the target server to run
	// the liveliness and readiness server on.
	// If zero, treat it the same as the main server.
	//
	// Default: 0
	StatusPort uint16

	// Timeout specifies a time limit for requests made by the client.
	// The timeout includes connection time, any redirects,
	// and reading the response body.
	//
	// A non-positive value means no timeout.
	//
	// Default: 0
	Timeout time.Duration

	// Username is the username sent with the request.
	// Set this along with Password if the target server requires basic auth.
	//
	// Default: "" (empty)
	Username string

	// Password is the password sent with the request.
	// Set this along with Username if the target server requires basic auth.
	//
	// Default: "" (empty)
	Password string

	// Charset is the character encoding of the request
	// set in the Content-Type header.
	//
	// Default: utf-8
	Charset string

	// Annotators are the default annotators with the annotation request.
	// If no annotators are specified with the annotation request in Annotate,
	// these annotators will be used.
	//
	// The annotators are separated by commas (,) in the string without spaces.
	// For example:
	//  tokenize,ssplit,pos,depparse
	//
	// Default: "" (empty, no annotator is specified by default)
	Annotators string
}

// New creates a new Client for the Stanford CoreNLP server
// with specified options.
//
// If opt is nil, it will use default options.
//
// Before returning the client, it will test whether the target server is live.
// If the test fails, it will report an error and return a nil client.
// Thus, make sure the server is ready and set the appropriate host address
// in opt before calling this function.
func New(opt *Options) (c Client, err error) {
	if opt == nil {
		opt = new(Options)
	}
	t := new(clientImpl)

	// Set fields of t (type: clientImpl) according to opt.
	hostname := strings.TrimSpace(opt.Hostname)
	if len(hostname) == 0 {
		hostname = "127.0.0.1"
	}
	mp, sp := opt.Port, opt.StatusPort
	if mp == 0 {
		mp = 9000
	}
	if sp == 0 {
		sp = mp
	}
	if addr, err := netip.ParseAddr(hostname); err == nil {
		// hostname is an IP address.
		t.host = netip.AddrPortFrom(addr, mp).String()
		if sp == mp {
			t.statusHost = t.host
		} else {
			t.statusHost = netip.AddrPortFrom(addr, sp).String()
		}
	} else {
		// hostname is not an IP address, may be a domain name, or invalid.
		// This step does not validate the host.
		// So simply join the hostname and port.
		t.host = hostname + ":" + strconv.FormatUint(uint64(mp), 10)
		if sp == mp {
			t.statusHost = t.host
		} else {
			t.statusHost = hostname + ":" + strconv.FormatUint(uint64(sp), 10)
		}
	}
	username := strings.TrimSpace(opt.Username)
	if len(username) > 0 {
		password := strings.TrimSpace(opt.Password)
		if len(password) > 0 {
			t.userinfo = url.UserPassword(username, password)
		} else {
			t.userinfo = url.User(username)
		}
	}
	if opt.Timeout > 0 {
		t.c.Timeout = opt.Timeout
	}
	if len(opt.Annotators) > 0 {
		t.annotators = strings.Join(strings.Fields(opt.Annotators), "") // drop white space
	}
	charset := strings.TrimSpace(opt.Charset)
	if len(charset) == 0 {
		charset = "utf-8"
	}
	t.contentType = "text/plain; charset=" + charset

	// Send a status request to /live.
	liveUrl := &url.URL{
		Scheme: "http",
		User:   t.userinfo,
		Host:   t.statusHost,
		Path:   "live",
	}
	resp, err := t.c.Get(liveUrl.String())
	if err != nil {
		return nil, errors.AutoWrap(fmt.Errorf("server not live: %v", err))
	}
	defer func(body io.ReadCloser) {
		_ = body.Close() // ignore error
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.AutoNew("server not live: got " + resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.AutoWrap(fmt.Errorf("server not live: %v", err))
	}
	if r := strings.TrimSpace(string(data)); r != "live" {
		return nil, errors.AutoNew("got " + r + " from " + liveUrl.Redacted() + "; want live")
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
}

// Ready sends a status request to the readiness endpoint (/ready) and
// reports whether the target server is ready to accept connections.
//
// It returns false if the server is not ready or any error is encountered.
func (c *clientImpl) Ready() bool {
	readyUrl := &url.URL{
		Scheme: "http",
		User:   c.userinfo,
		Host:   c.statusHost,
		Path:   "ready",
	}
	resp, err := c.c.Get(readyUrl.String())
	if err != nil {
		return false
	}
	defer func(body io.ReadCloser) {
		_ = body.Close() // ignore error
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return false
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(data)) == "ready"
}

// Annotate sends an annotation request with specified text and annotators.
// The annotation result is represented as a CoreNLP document and
// written to outDoc.
//
// outDoc must be a non-nil pointer to an auto-generated Document
// structure, for example:
//
//  import "github.com/donyori/gocorenlp/model/v4.4.0-e90f30f13c40/pb"
//  ...
//  outDoc := new(pb.Document)
//  err := Annotate(outDoc, "Hello world!", "tokenize,ssplit,pos")
//  ...
//
// If outDoc is nil or not a pointer to Document,
// a runtime error will occur.
//
// If no annotators are specified,
// the client's default annotators will be used.
// If the client's annotators are also not specified,
// the server's default annotators will be used.
//
// The annotators are separated by commas (,) in the string without spaces.
// For example:
//  tokenize,ssplit,pos,depparse
func (c *clientImpl) Annotate(outDoc proto.Message, text string, annotators string) error {
	// Make request.
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
		return errors.AutoWrap(err)
	}
	qv := url.Values{"properties": []string{string(propBytes)}}
	annUrl := &url.URL{
		Scheme:   "http",
		User:     c.userinfo,
		Host:     c.host,
		RawQuery: qv.Encode(),
	}

	// Send request and read response.
	resp, err := c.c.Post(annUrl.String(), c.contentType, strings.NewReader(text))
	if err != nil {
		return errors.AutoWrap(err)
	}
	defer func(body io.ReadCloser) {
		_ = body.Close() // ignore error
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return errors.AutoNew("got " + resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.AutoWrap(err)
	}

	// Parse ProtoBuf message.
	v, n := protowire.ConsumeBytes(data)
	if n < 0 {
		return errors.AutoWrap(protowire.ParseError(n))
	}
	err = proto.Unmarshal(v, outDoc)
	if err != nil {
		return errors.AutoWrap(err)
	}
	return nil
}
