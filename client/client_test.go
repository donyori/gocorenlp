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

package client_test

// To test with default settings,
// launch a Stanford CoreNLP 4.5.5 server
// (both the main server and the status server) listening on 127.0.0.1:9000.
// Set the server ID (i.e., server name) to testdefault.
// The server should use its default language model.
//
// To test with different status port settings,
// launch a Stanford CoreNLP 4.5.5 server,
// with its main server listening on 127.0.0.1:9100 and
// its status server listening on 127.0.0.1:9101.
// Set the server ID (i.e., server name) to testdiffstatus.
// The server should use its default language model.
//
// To test with basic auth settings,
// launch a Stanford CoreNLP 4.5.5 server
// (both the main server and the status server) listening on 127.0.0.1:9200,
// with username="user1" and password="u1%passWORD".
// Set the server ID (i.e., server name) to testuser.
// The server should use its default language model.
//
// To test the shutdown functionality,
// launch a Stanford CoreNLP 4.5.5 server
// (both the main server and the status server) listening on 127.0.0.1:9300,
// without setting its server ID,
// and launch a Stanford CoreNLP 4.5.5 server
// (both the main server and the status server) listening on 127.0.0.1:9301,
// with server ID testshutdown.

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/donyori/gocorenlp/client"
	"github.com/donyori/gocorenlp/errors"
	"github.com/donyori/gocorenlp/model"
	"github.com/donyori/gocorenlp/model/v4.5.5-f1b929e47a57/pb"
)

const (
	DefaultIndex = iota
	DiffStatusIndex
	UserIndex

	NumNotShutdownIndex

	ShutdownNoServerIDIndex = iota - 1
	ShutdownServerIDIndex

	NumIndex
)

const (
	DefaultPortIndex = iota
	DiffStatusMainPortIndex
	DiffStatusStatusPortIndex
	UserPortIndex
	ShutdownNoServerIDPortIndex
	ShutdownServerIDPortIndex

	NumPort
)

var NotShutdownSubtestNames = [NumNotShutdownIndex]string{"default", "diff status", "user"}
var ShutdownSubtestNames = [NumIndex - NumNotShutdownIndex]string{"no server ID", "server ID"}
var ServerIDs = [NumIndex]string{"testdefault", "testdiffstatus", "testuser", "", "testshutdown"}

var ServerPorts = [NumPort]uint16{9000, 9100, 9101, 9200, 9300, 9301}

var IndexPortMap = map[int]uint16{
	DefaultIndex:            ServerPorts[DefaultPortIndex],
	DiffStatusIndex:         ServerPorts[DiffStatusMainPortIndex],
	UserIndex:               ServerPorts[UserPortIndex],
	ShutdownNoServerIDIndex: ServerPorts[ShutdownNoServerIDPortIndex],
	ShutdownServerIDIndex:   ServerPorts[ShutdownServerIDPortIndex],
}

const (
	Username = "user1"
	Password = "u1%passWORD"
)

const Text = "The quick brown fox jumped over the lazy dog."

const (
	InvalidIndexLayout = "invalid index %d; should be [0-%d]"
	SkipLayout         = "server 127.0.0.1:%d is offline; skip this test"
)

func TestNew(t *testing.T) {
	for i, name := range NotShutdownSubtestNames {
		t.Run(name, func(t *testing.T) {
			var c client.Client
			var err error
			switch i {
			case DefaultIndex:
				port := ServerPorts[DefaultPortIndex]
				if !CheckIsServerListeningOnPort(port) {
					t.Skipf(SkipLayout, port)
				}
				c, err = client.New(nil)
			case DiffStatusIndex:
				mainPort := ServerPorts[DiffStatusMainPortIndex]
				statusPort := ServerPorts[DiffStatusStatusPortIndex]
				if !CheckIsServerListeningOnPort(mainPort) {
					t.Skipf(SkipLayout, mainPort)
				}
				if !CheckIsServerListeningOnPort(statusPort) {
					t.Skipf(SkipLayout, statusPort)
				}
				c, err = client.New(&client.Options{
					Port:       mainPort,
					StatusPort: statusPort,
				})
			case UserIndex:
				port := ServerPorts[UserPortIndex]
				if !CheckIsServerListeningOnPort(port) {
					t.Skipf(SkipLayout, port)
				}
				c, err = client.New(&client.Options{
					Port:     port,
					Username: Username,
					Password: Password,
				})
			default:
				// This case should never happen.
				t.Fatalf(InvalidIndexLayout, i, NumNotShutdownIndex)
			}
			if err != nil {
				t.Fatal(err)
			}
			if c == nil {
				t.Error("got nil client")
			}
		})
	}
}

func TestClient_Live(t *testing.T) {
	for i, name := range NotShutdownSubtestNames {
		t.Run(name, func(t *testing.T) {
			online := IsServerOnline(t, i, false)
			c := NewClientForTest(t, i)
			err := c.Live()
			if online {
				if err != nil {
					t.Error(err)
				}
			} else if err == nil {
				t.Error("got nil error but server is offline")
			} else if !errors.IsConnectionError(err) {
				t.Error("got non-nil error but not connection error:", err)
			}
		})
	}
}

func TestClient_Ready(t *testing.T) {
	for i, name := range NotShutdownSubtestNames {
		t.Run(name, func(t *testing.T) {
			online := IsServerOnline(t, i, false)
			c := NewClientForTest(t, i)
			err := c.Ready()
			if online {
				if err != nil {
					t.Error(err)
				}
			} else if err == nil {
				t.Error("got nil error but server is offline")
			} else if !errors.IsConnectionError(err) {
				t.Error("got non-nil error but not connection error:", err)
			}
		})
	}
}

func TestClient_Annotate(t *testing.T) {
	for i, name := range NotShutdownSubtestNames {
		t.Run(name, func(t *testing.T) {
			SkipIfServerOffline(t, i)
			AnnotateFunc(t, func(annotators string) *pb.Document {
				c := NewClientForTest(t, i)
				doc := new(pb.Document)
				if err := c.Annotate(strings.NewReader(Text), annotators, doc); err != nil {
					t.Error(err)
					return nil
				}
				return doc
			})
		})
	}
}

func TestClient_AnnotateString(t *testing.T) {
	for i, name := range NotShutdownSubtestNames {
		t.Run(name, func(t *testing.T) {
			SkipIfServerOffline(t, i)
			AnnotateFunc(t, func(annotators string) *pb.Document {
				c := NewClientForTest(t, i)
				doc := new(pb.Document)
				if err := c.AnnotateString(Text, annotators, doc); err != nil {
					t.Error(err)
					return nil
				}
				return doc
			})
		})
	}
}

func TestClient_AnnotateRaw(t *testing.T) {
	for i, name := range NotShutdownSubtestNames {
		t.Run(name, func(t *testing.T) {
			SkipIfServerOffline(t, i)
			AnnotateFunc(t, func(annotators string) *pb.Document {
				c := NewClientForTest(t, i)
				var b bytes.Buffer
				written, err := c.AnnotateRaw(strings.NewReader(Text), annotators, &b)
				if err != nil {
					t.Error(err)
					return nil
				}
				if n := int64(b.Len()); written != n {
					t.Errorf("got written %d; want %d", written, n)
					return nil
				}
				doc := new(pb.Document)
				if err = model.DecodeResponseBody(b.Bytes(), doc); err != nil {
					t.Error(err)
					return nil
				}
				return doc
			})
		})
	}
}

func TestClient_AnnotateStringRaw(t *testing.T) {
	for i, name := range NotShutdownSubtestNames {
		t.Run(name, func(t *testing.T) {
			SkipIfServerOffline(t, i)
			AnnotateFunc(t, func(annotators string) *pb.Document {
				c := NewClientForTest(t, i)
				var b bytes.Buffer
				written, err := c.AnnotateStringRaw(Text, annotators, &b)
				if err != nil {
					t.Error(err)
					return nil
				}
				if n := int64(b.Len()); written != n {
					t.Errorf("got written %d; want %d", written, n)
					return nil
				}
				doc := new(pb.Document)
				if err = model.DecodeResponseBody(b.Bytes(), doc); err != nil {
					t.Error(err)
					return nil
				}
				return doc
			})
		})
	}
}

func TestClient_ShutdownLocal(t *testing.T) {
	for i, name := range ShutdownSubtestNames {
		t.Run(name, func(t *testing.T) {
			index := i + NumNotShutdownIndex
			SkipIfServerOffline(t, index)
			c := NewClientForTest(t, index)
			if err := c.ShutdownLocal(); err != nil {
				t.Error(err)
			}
		})
	}
}

// CheckIsServerListeningOnPort checks whether a local server
// is listening on the specified port.
func CheckIsServerListeningOnPort(port uint16) bool {
	conn, err := net.DialTimeout(
		"tcp",
		"127.0.0.1:"+strconv.FormatUint(uint64(port), 10),
		time.Millisecond*30,
	)
	if err != nil {
		return false
	}
	_ = conn.Close() // ignore error
	return true
}

// IsServerOnline reports whether the server is online.
//
// It determines the server according to the specified test index
// and the indicator mainServer.
// If mainServer is true, it uses the main server port;
// otherwise, it uses the status server port.
//
// It calls tb.Fatalf if the index is out of range.
func IsServerOnline(tb testing.TB, index int, mainServer bool) bool {
	if index < 0 || index > NumIndex {
		tb.Fatalf(InvalidIndexLayout, index, NumIndex)
	}
	port, ok := IndexPortMap[index]
	if !ok {
		tb.Fatalf("cannot find port with key %d in IndexPortMap", index)
	}
	if index == DiffStatusIndex && !mainServer {
		port = ServerPorts[DiffStatusStatusPortIndex]
	}
	return CheckIsServerListeningOnPort(port)
}

// SkipIfServerOffline skips the test if the server is offline.
//
// It determines the server according to the specified test index.
//
// It calls tb.Fatalf if the index is out of range.
func SkipIfServerOffline(tb testing.TB, index int) {
	if !IsServerOnline(tb, index, true) {
		tb.Skipf(SkipLayout, IndexPortMap[index])
	}
}

// NewClientForTest creates a new client.Client
// according to the specified test index.
// Unlike client.New, it does not test whether the server is live.
//
// It calls tb.Fatalf if the index is out of range.
func NewClientForTest(tb testing.TB, index int) client.Client {
	var opts *client.Options
	switch index {
	case DefaultIndex:
		opts = &client.Options{ServerID: ServerIDs[index]}
	case DiffStatusIndex:
		opts = &client.Options{
			Port:       ServerPorts[DiffStatusMainPortIndex],
			StatusPort: ServerPorts[DiffStatusStatusPortIndex],
			ServerID:   ServerIDs[index],
		}
	case UserIndex:
		opts = &client.Options{
			Port:     ServerPorts[UserPortIndex],
			Username: Username,
			Password: Password,
			ServerID: ServerIDs[index],
		}
	case ShutdownNoServerIDIndex:
		opts = &client.Options{Port: ServerPorts[ShutdownNoServerIDPortIndex]}
	case ShutdownServerIDIndex:
		opts = &client.Options{
			Port:     ServerPorts[ShutdownServerIDPortIndex],
			ServerID: ServerIDs[index],
		}
	default:
		tb.Fatalf(InvalidIndexLayout, index, NumIndex)
	}
	opts.Annotators = "tokenize,ssplit,pos"
	opts.ClientTimeout = time.Millisecond * 600
	return client.NewClientWithoutCheckingLive(opts)
}

// AnnotateFunc encapsulates common code for testing the methods Annotate,
// AnnotateString, AnnotateRaw, and AnnotateStringRaw of client.Client,
// as well as their corresponding global functions.
func AnnotateFunc(t *testing.T, f func(annotators string) *pb.Document) {
	annotators := []string{"", "tokenize,ssplit,pos"}
	for _, ann := range annotators {
		t.Run(fmt.Sprintf("annotator=%q", ann), func(t *testing.T) {
			doc := f(ann)
			if doc != nil {
				CheckAnnotation(t, doc)
			}
		})
	}
}

// CheckAnnotation checks the result of annotation to the text:
//
//	The quick brown fox jumped over the lazy dog.
//
// It checks the document text, sentence split, token word,
// content before token, content after token, and token part-of-speech tag.
func CheckAnnotation(t *testing.T, doc *pb.Document) {
	const nTokens = 10
	wordArray := [nTokens]string{"The", "quick", "brown", "fox", "jumped", "over", "the", "lazy", "dog", "."}
	gapArray := [nTokens + 1]string{"", " ", " ", " ", " ", " ", " ", " ", " "}
	posArray := [nTokens]string{"DT", "JJ", "JJ", "NN", "VBD", "IN", "DT", "JJ", "NN", "."}

	if txt := doc.GetText(); txt != Text {
		t.Errorf("got doc text %q; want %q", txt, Text)
	}

	sentences := doc.GetSentence()
	if n := len(sentences); n != 1 {
		t.Errorf("got %d sentences; want 1", n)
		return
	}

	tokens := sentences[0].GetToken()
	if n := len(tokens); n != nTokens {
		t.Errorf("got %d token(s); want %d", n, nTokens)
		return
	}
	for i, token := range tokens {
		if w := token.GetWord(); w != wordArray[i] {
			t.Errorf("Token %d: got Word %q; want %q", i, w, wordArray[i])
		}
		if b := token.GetBefore(); b != gapArray[i] {
			t.Errorf("Token %d: got Before %q; want %q", i, b, gapArray[i])
		}
		if a := token.GetAfter(); a != gapArray[i+1] {
			t.Errorf("Token %d: got After %q; want %q", i, a, gapArray[i+1])
		}
		if p := token.GetPos(); p != posArray[i] {
			t.Errorf("Token %d: got Pos %q; want %q", i, p, posArray[i])
		}
	}
}
