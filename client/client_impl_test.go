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

// To test with default settings, launch a Stanford CoreNLP 4.4.0 server
// (both the main server and the status server) listening on 127.0.0.1:9000.
// Set the server ID (i.e., server name) to testdefault.
// The server should use its default language model.
//
// To test with different status port settings,
// launch a Stanford CoreNLP 4.4.0 server,
// with its main server listening on 127.0.0.1:9100 and
// its status server listening on 127.0.0.1:9101.
// Set the server ID (i.e., server name) to testdiffstatus.
// The server should use its default language model.
//
// To test with basic auth settings, launch a Stanford CoreNLP 4.4.0 server
// (both the main server and the status server) listening on 127.0.0.1:9200,
// with username="user1" and password="u1%passWORD".
// Set the server ID (i.e., server name) to testuser.
// The server should use its default language model.
//
// To test the shutdown functionality, launch a Stanford CoreNLP 4.4.0 server
// (both the main server and the status server) listening on 127.0.0.1:9300,
// without setting its server ID.
// And launch a Stanford CoreNLP 4.4.0 server (both the main server and
// the status server) listening on 127.0.0.1:9301, with server ID testshutdown.

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/donyori/gocorenlp/errors"
	"github.com/donyori/gocorenlp/model"
	"github.com/donyori/gocorenlp/model/v4.4.0-e90f30f13c40/pb"
)

const (
	DefaultIndex = iota
	DiffStatusIndex
	UserIndex

	NonShutdownNum
)

const (
	ShutdownNoServerIdIndex = NonShutdownNum + iota
	ShutdownServerIdIndex

	N
)

const (
	DefaultPortIndex = iota
	DiffStatusMainPortIndex
	DiffStatusStatusPortIndex
	UserPortIndex
	ShutdownNoServerIdPortIndex
	ShutdownServerIdPortIndex

	PortNum
)

var NonShutdownSubtestNames = [NonShutdownNum]string{"default", "diff status", "user"}
var ShutdownSubtestNames = [N - NonShutdownNum]string{"no server ID", "server ID"}
var ServerIds = [N]string{"testdefault", "testdiffstatus", "testuser", "", "testshutdown"}

var IsServersOnline [PortNum]bool
var ServerPorts = [PortNum]uint16{9000, 9100, 9101, 9200, 9300, 9301}

const (
	Username = "user1"
	Password = "u1%passWORD"
)

const Text = "The quick brown fox jumped over the lazy dog."

const InvalidIndexFormat = "invalid index %d; should be [0-%d]"

var CheckServerOnlineOnce sync.Once

func TestClientImpl_Live(t *testing.T) {
	for i, name := range NonShutdownSubtestNames {
		t.Run(name, func(t *testing.T) {
			online := IsServerOnline(t, i, false)
			c := NewClientImpl(t, i)
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

func TestClientImpl_Ready(t *testing.T) {
	for i, name := range NonShutdownSubtestNames {
		t.Run(name, func(t *testing.T) {
			online := IsServerOnline(t, i, false)
			c := NewClientImpl(t, i)
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

func TestClientImpl_Annotate(t *testing.T) {
	for i, name := range NonShutdownSubtestNames {
		t.Run(name, func(t *testing.T) {
			SkipIfServerOffline(t, i)
			AnnotateMethodsFunc(t, func(tb testing.TB, annotators string) *pb.Document {
				c := NewClientImpl(tb, i)
				doc := new(pb.Document)
				if err := c.Annotate(strings.NewReader(Text), annotators, doc); err != nil {
					tb.Error(err)
					return nil
				}
				return doc
			})
		})
	}
}

func TestClientImpl_AnnotateString(t *testing.T) {
	for i, name := range NonShutdownSubtestNames {
		t.Run(name, func(t *testing.T) {
			SkipIfServerOffline(t, i)
			AnnotateMethodsFunc(t, func(tb testing.TB, annotators string) *pb.Document {
				c := NewClientImpl(tb, i)
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

func TestClientImpl_AnnotateRaw(t *testing.T) {
	for i, name := range NonShutdownSubtestNames {
		t.Run(name, func(t *testing.T) {
			SkipIfServerOffline(t, i)
			AnnotateMethodsFunc(t, func(tb testing.TB, annotators string) *pb.Document {
				c := NewClientImpl(t, i)
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

func TestClientImpl_ShutdownLocal(t *testing.T) {
	for i, name := range ShutdownSubtestNames {
		t.Run(name, func(t *testing.T) {
			index := i + NonShutdownNum
			SkipIfServerOffline(t, index)
			c := NewClientImpl(t, index)
			if err := c.ShutdownLocal(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestMain(m *testing.M) {
	CheckAllServerOnline()
	os.Exit(m.Run())
}

// CheckAllServerOnline checks all the servers whether they are online.
//
// It only takes effect once.
func CheckAllServerOnline() {
	CheckServerOnlineOnce.Do(func() {
		for i := 0; i < PortNum; i++ {
			IsServersOnline[i] = CheckIsServerListeningOnPort(ServerPorts[i])
		}
	})
}

// CheckIsServerListeningOnPort checks whether a local server
// is listening on the specified port.
func CheckIsServerListeningOnPort(port uint16) bool {
	conn, err := net.DialTimeout(
		"tcp",
		"127.0.0.1:"+strconv.FormatUint(uint64(port), 10),
		time.Millisecond*10,
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
	CheckAllServerOnline()
	switch index {
	case DefaultIndex:
		return IsServersOnline[DefaultPortIndex]
	case DiffStatusIndex:
		if mainServer {
			return IsServersOnline[DiffStatusMainPortIndex]
		}
		return IsServersOnline[DiffStatusStatusPortIndex]
	case UserIndex:
		return IsServersOnline[UserPortIndex]
	case ShutdownNoServerIdIndex:
		return IsServersOnline[ShutdownNoServerIdPortIndex]
	case ShutdownServerIdIndex:
		return IsServersOnline[ShutdownServerIdPortIndex]
	default:
		tb.Fatalf(InvalidIndexFormat, index, N)
	}
	// unreachable
	return false
}

// SkipIfServerOffline skips the test if the server is offline.
//
// It determines the server according to the specified test index.
//
// It calls tb.Fatalf if the index is out of range.
func SkipIfServerOffline(tb testing.TB, index int) {
	CheckAllServerOnline()
	var portIdx int
	switch index {
	case DefaultIndex:
		portIdx = DefaultPortIndex
	case DiffStatusIndex:
		portIdx = DiffStatusMainPortIndex
	case UserIndex:
		portIdx = UserPortIndex
	case ShutdownNoServerIdIndex:
		portIdx = ShutdownNoServerIdPortIndex
	case ShutdownServerIdIndex:
		portIdx = ShutdownServerIdPortIndex
	default:
		tb.Fatalf(InvalidIndexFormat, index, N)
	}
	if !IsServersOnline[portIdx] {
		tb.Skipf("server 127.0.0.1:%d is offline; skip this test", ServerPorts[portIdx])
	}
}

// NewClientImpl creates a new *clientImpl according to
// the specified test index.
//
// It calls tb.Fatalf if the index is out of range.
func NewClientImpl(tb testing.TB, index int) *clientImpl {
	var opt *Options
	switch index {
	case DefaultIndex:
		opt = &Options{ServerId: ServerIds[index]}
	case DiffStatusIndex:
		opt = &Options{
			Port:       ServerPorts[DiffStatusMainPortIndex],
			StatusPort: ServerPorts[DiffStatusStatusPortIndex],
			ServerId:   ServerIds[index],
		}
	case UserIndex:
		opt = &Options{
			Port:     ServerPorts[UserPortIndex],
			Username: Username,
			Password: Password,
			ServerId: ServerIds[index],
		}
	case ShutdownNoServerIdIndex:
		opt = &Options{Port: ServerPorts[ShutdownNoServerIdPortIndex]}
	case ShutdownServerIdIndex:
		opt = &Options{
			Port:     ServerPorts[ShutdownServerIdPortIndex],
			ServerId: ServerIds[index],
		}
	default:
		tb.Fatalf(InvalidIndexFormat, index, N)
	}
	opt.Annotators = "tokenize,ssplit,pos"
	opt.Timeout = time.Millisecond * 300
	return newClientImpl(opt)
}

// AnnotateMethodsFunc encapsulates common code for testing the methods
// Annotate, AnnotateString, and AnnotateRaw of *clientImpl.
func AnnotateMethodsFunc(t *testing.T, f func(tb testing.TB, annotators string) *pb.Document) {
	annotators := []string{"", "tokenize,ssplit,pos"}
	for _, ann := range annotators {
		t.Run(fmt.Sprintf("annotator=%q", ann), func(t *testing.T) {
			doc := f(t, ann)
			if doc != nil {
				CheckAnnotation(t, doc)
			}
		})
	}
}

// CheckAnnotation checks the result of annotation to the text:
//  The quick brown fox jumped over the lazy dog.
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
