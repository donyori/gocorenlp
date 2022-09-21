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

package client_test

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/donyori/gocorenlp/client"
	"github.com/donyori/gocorenlp/errors"
	"github.com/donyori/gocorenlp/model"
	"github.com/donyori/gocorenlp/model/v4.5.0-45b47e245c36/pb"
)

const DefaultPort uint16 = 9000
const FlagNameTestFuncShutdown = "testfuncshutdown"

var RunShutdownTest bool

const Text = "The quick brown fox jumped over the lazy dog."

var ParseFlagOnce sync.Once

func init() {
	flag.BoolVar(&RunShutdownTest, FlagNameTestFuncShutdown, false, "Set if want to test Shutdown.")
}

func TestLive(t *testing.T) {
	err := client.Live()
	if CheckIsServerListeningOnPort(DefaultPort) {
		if err != nil {
			t.Error(err)
		}
	} else if err == nil {
		t.Error("got nil error but server is offline")
	} else if !errors.IsConnectionError(err) {
		t.Error("got non-nil error but not connection error:", err)
	}
}

func TestReady(t *testing.T) {
	err := client.Ready()
	if CheckIsServerListeningOnPort(DefaultPort) {
		if err != nil {
			t.Error(err)
		}
	} else if err == nil {
		t.Error("got nil error but server is offline")
	} else if !errors.IsConnectionError(err) {
		t.Error("got non-nil error but not connection error:", err)
	}
}

func TestAnnotate(t *testing.T) {
	SkipIfDefaultServerOffline(t)
	AnnotateFunctionsFunc(t, func(annotators string) *pb.Document {
		doc := new(pb.Document)
		if err := client.Annotate(strings.NewReader(Text), annotators, doc); err != nil {
			t.Error(err)
			return nil
		}
		return doc
	})
}

func TestAnnotateString(t *testing.T) {
	SkipIfDefaultServerOffline(t)
	AnnotateFunctionsFunc(t, func(annotators string) *pb.Document {
		doc := new(pb.Document)
		if err := client.AnnotateString(Text, annotators, doc); err != nil {
			t.Error(err)
			return nil
		}
		return doc
	})
}

func TestAnnotateRaw(t *testing.T) {
	SkipIfDefaultServerOffline(t)
	AnnotateFunctionsFunc(t, func(annotators string) *pb.Document {
		var b bytes.Buffer
		written, err := client.AnnotateRaw(strings.NewReader(Text), annotators, &b)
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
}

func TestAnnotateStringRaw(t *testing.T) {
	SkipIfDefaultServerOffline(t)
	AnnotateFunctionsFunc(t, func(annotators string) *pb.Document {
		var b bytes.Buffer
		written, err := client.AnnotateStringRaw(Text, annotators, &b)
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
}

func TestShutdown(t *testing.T) {
	ParseFlag()
	if !RunShutdownTest {
		t.Skip("skip this test because flag -" + FlagNameTestFuncShutdown + " is not set")
	}
	SkipIfDefaultServerOffline(t)
	if err := client.Shutdown(); err != nil {
		t.Error(err)
	}
}

// ParseFlag calls flag.Parse to parse the command line flags.
//
// It only takes effect once.
func ParseFlag() {
	ParseFlagOnce.Do(func() {
		flag.Parse()
	})
}

// SkipIfDefaultServerOffline skips the test if the default server is offline.
func SkipIfDefaultServerOffline(tb testing.TB) {
	if !CheckIsServerListeningOnPort(DefaultPort) {
		tb.Skipf("server 127.0.0.1:%d is offline; skip this test", DefaultPort)
	}
}

// AnnotateFunctionsFunc encapsulates common code for testing the functions
// Annotate, AnnotateString, AnnotateRaw, and AnnotateStringRaw.
func AnnotateFunctionsFunc(t *testing.T, f func(annotators string) *pb.Document) {
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
