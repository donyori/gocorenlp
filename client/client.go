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
	"os"
	"path/filepath"
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
	// the client's default annotators will be used.
	// If the client's annotators are also not specified,
	// the server's default annotators will be used.
	//
	// The annotators are separated by commas (,) in the string without spaces.
	// For example:
	//  tokenize,ssplit,pos,depparse
	//
	// outDoc must be a non-nil pointer to an auto-generated Document
	// structure, for example:
	//
	//  import "github.com/donyori/gocorenlp/model/v4.4.0-e90f30f13c40/pb"
	//  ...
	//  outDoc := new(pb.Document)
	//  err := Annotate(reader, "tokenize,ssplit,pos", outDoc)
	//  ...
	//
	// If outDoc is nil or not a pointer to Document,
	// a runtime error will occur.
	Annotate(reader io.Reader, annotators string, outDoc proto.Message) error

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
	//  tokenize,ssplit,pos,depparse
	//
	// outDoc must be a non-nil pointer to an auto-generated Document
	// structure, for example:
	//
	//  import "github.com/donyori/gocorenlp/model/v4.4.0-e90f30f13c40/pb"
	//  ...
	//  outDoc := new(pb.Document)
	//  err := AnnotateString("Hello world!", "tokenize,ssplit,pos", outDoc)
	//  ...
	//
	// If outDoc is nil or not a pointer to Document,
	// a runtime error will occur.
	AnnotateString(text string, annotators string, outDoc proto.Message) error

	// Shutdown finds the shutdown key and then sends a shutdown request
	// to stop the target server.
	//
	// It returns nil if the server has been stopped successfully.
	Shutdown() error

	// ShutdownWithKey sends a shutdown request with the specified key
	// to stop the target server.
	//
	// It returns nil if the server has been stopped successfully.
	ShutdownWithKey(key string) error

	// private prevents others from implementing this interface,
	// so future additions to it will not violate compatibility.
	private()
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
	// the liveness and readiness server on.
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
	// Only valid when Username is not empty.
	//
	// Default: "" (empty)
	Password string

	// Charset is the character encoding of the request
	// set in the Content-Type header.
	//
	// Default: utf-8
	Charset string

	// Annotators are the default annotators with the annotation request.
	// If no annotators are specified with the annotation request,
	// these annotators will be used.
	//
	// The annotators are separated by commas (,) in the string without spaces.
	// For example:
	//  tokenize,ssplit,pos,depparse
	//
	// Default: "" (empty, no annotator is specified by default)
	Annotators string

	// ServerId is the value of the option -server_id used
	// when starting the target server.
	//
	// If the server is started without that option, leave it empty.
	//
	// Default: "" (empty)
	ServerId string
}

// New creates a new Client for the Stanford CoreNLP server
// with the specified options.
//
// If opt is nil, it will use default options.
//
// Before returning the client, it will test whether the target server is live.
// If the test fails, it will report an error and return a nil client.
// Thus, make sure the server is online and set the appropriate host address
// in opt before calling this function.
func New(opt *Options) (c Client, err error) {
	t := newClientImpl(opt)
	err = t.Live()
	if err != nil {
		return nil, errors.AutoWrap(err)
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
	serverId    string
}

// newClientImpl creates a new clientImpl and
// sets its fields according to the specified options opt.
func newClientImpl(opt *Options) *clientImpl {
	if opt == nil {
		opt = new(Options)
	}
	c := &clientImpl{serverId: strings.TrimSpace(opt.ServerId)}
	c.host, c.statusHost = makeHosts(opt.Hostname, opt.Port, opt.StatusPort)
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
		return errors.AutoWrap(err)
	}
	defer closeIgnoreError(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return errors.AutoNew("got response status " + resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.AutoWrap(err)
	}
	if body := strings.TrimSpace(string(data)); body != "live" {
		return errors.AutoNew(fmt.Sprintf("got response %s; want live", body))
	}
	return nil
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
		return errors.AutoWrap(err)
	}
	defer closeIgnoreError(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return errors.AutoNew("got response status " + resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.AutoWrap(err)
	}
	if body := strings.TrimSpace(string(data)); body != "ready" {
		return errors.AutoNew(fmt.Sprintf("got response %s; want ready", body))
	}
	return nil
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
//  tokenize,ssplit,pos,depparse
//
// outDoc must be a non-nil pointer to an auto-generated Document
// structure, for example:
//
//  import "github.com/donyori/gocorenlp/model/v4.4.0-e90f30f13c40/pb"
//  ...
//  outDoc := new(pb.Document)
//  err := Annotate(reader, "tokenize,ssplit,pos", outDoc)
//  ...
//
// If outDoc is nil or not a pointer to Document,
// a runtime error will occur.
func (c *clientImpl) Annotate(reader io.Reader, annotators string, outDoc proto.Message) error {
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
	resp, err := c.c.Post(annUrl.String(), c.contentType, reader)
	if err != nil {
		return errors.AutoWrap(err)
	}
	defer closeIgnoreError(resp.Body)
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
//  tokenize,ssplit,pos,depparse
//
// outDoc must be a non-nil pointer to an auto-generated Document
// structure, for example:
//
//  import "github.com/donyori/gocorenlp/model/v4.4.0-e90f30f13c40/pb"
//  ...
//  outDoc := new(pb.Document)
//  err := AnnotateString("Hello world!", "tokenize,ssplit,pos", outDoc)
//  ...
//
// If outDoc is nil or not a pointer to Document,
// a runtime error will occur.
func (c *clientImpl) AnnotateString(text string, annotators string, outDoc proto.Message) error {
	err := errors.AutoWrap(c.Annotate(strings.NewReader(text), annotators, outDoc))
	if err != nil {
		return errors.AutoWrap(err)
	}
	return nil
}

// Shutdown finds the shutdown key and then sends a shutdown request
// to stop the target server.
//
// It returns nil if the server has been stopped successfully.
func (c *clientImpl) Shutdown() error {
	tmpDir := os.TempDir()
	name := filepath.Join(tmpDir, "corenlp.shutdown")
	if len(c.serverId) > 0 {
		name += "." + c.serverId
	}
	key, err := os.ReadFile(name)
	if err != nil {
		return errors.AutoWrap(fmt.Errorf("failed to find the key: %v", err))
	}
	err = c.ShutdownWithKey(string(key))
	if err != nil {
		return errors.AutoWrap(err)
	}
	return nil
}

// ShutdownWithKey sends a shutdown request with the specified key
// to stop the target server.
//
// It returns nil if the server has been stopped successfully.
func (c *clientImpl) ShutdownWithKey(key string) error {
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
		return errors.AutoWrap(err)
	}
	defer closeIgnoreError(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return errors.AutoNew("got response status " + resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.AutoWrap(err)
	}
	if body := strings.TrimSpace(string(data)); body != "Shutdown successful!" {
		return errors.AutoNew(fmt.Sprintf("got response %s; want Shutdown successful!", body))
	}
	return nil
}

func (c *clientImpl) private() {}

// closeIgnoreError closes the specified closer if it is non-nil
// and ignores any error encountered.
func closeIgnoreError(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}

// makeHosts generates the hosts (including the hostname part and
// the port number part) for the main server and status server
// from the specified hostname, port, and statusPort.
func makeHosts(hostname string, port, statusPort uint16) (host, statusHost string) {
	hostname = strings.TrimSpace(hostname)
	if len(hostname) == 0 {
		hostname = "127.0.0.1"
	}
	if port == 0 {
		port = 9000
	}
	if statusPort == 0 {
		statusPort = port
	}
	if addr, err := netip.ParseAddr(hostname); err == nil {
		// hostname is an IP address.
		host = netip.AddrPortFrom(addr, port).String()
		if statusPort == port {
			statusHost = host
		} else {
			statusHost = netip.AddrPortFrom(addr, statusPort).String()
		}
	} else {
		// hostname is not an IP address, may be a domain name, or invalid.
		// This function does not validate the host.
		// So simply join the hostname and port.
		host = hostname + ":" + strconv.FormatUint(uint64(port), 10)
		if statusPort == port {
			statusHost = host
		} else {
			statusHost = hostname + ":" + strconv.FormatUint(uint64(statusPort), 10)
		}
	}
	return
}
