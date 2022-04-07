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
	"strings"
	"testing"

	"github.com/donyori/gocorenlp/model/v4.4.0-e90f30f13c40/pb"
)

// Run the following tests with a Stanford CoreNLP 4.4.0 server running and
// (both the main server and the status server) listening to 127.0.0.1:9000.

func TestNew(t *testing.T) {
	c, err := New(nil)
	if err != nil {
		t.Fatal(err)
	}
	if c == nil {
		t.Error("got nil client")
	}
}

func TestClientImpl_Live(t *testing.T) {
	c := testNewDefaultClientImpl()
	if err := c.Live(); err != nil {
		t.Error(err)
	}
}

func TestClientImpl_Ready(t *testing.T) {
	c := testNewDefaultClientImpl()
	if err := c.Ready(); err != nil {
		t.Error(err)
	}
}

func TestClientImpl_Annotate(t *testing.T) {
	const text = "The quick brown fox jumped over the lazy dog."
	reader := strings.NewReader(text)
	c := testNewDefaultClientImpl()
	doc := new(pb.Document)
	if err := c.Annotate(reader, "tokenize,ssplit,pos", doc); err != nil {
		t.Fatal(err)
	}
	testCheckAnnotation(t, doc)
}

func TestClientImpl_AnnotateString(t *testing.T) {
	const text = "The quick brown fox jumped over the lazy dog."
	c := testNewDefaultClientImpl()
	doc := new(pb.Document)
	if err := c.AnnotateString(text, "tokenize,ssplit,pos", doc); err != nil {
		t.Fatal(err)
	}
	testCheckAnnotation(t, doc)
}

// testNewDefaultClientImpl creates a Client
// connecting to 127.0.0.1:9000,
// with no userinfo, no timeout,
// annotators "tokenize,ssplit,pos",
// and contentType "application/x-www-form-urlencoded; charset=utf-8".
func testNewDefaultClientImpl() *clientImpl {
	return &clientImpl{
		host:        "127.0.0.1:9000",
		statusHost:  "127.0.0.1:9000",
		userinfo:    nil,
		annotators:  "tokenize,ssplit,pos",
		contentType: "application/x-www-form-urlencoded; charset=utf-8",
	}
}

// testCheckAnnotation checks the result of annotation to the text:
//  The quick brown fox jumped over the lazy dog.
//
// It checks the document text, sentence split, token word,
// content before token, content after token, and token part-of-speech tag.
func testCheckAnnotation(t *testing.T, doc *pb.Document) {
	const nTokens = 10
	const text = "The quick brown fox jumped over the lazy dog."
	wordArray := [nTokens]string{"The", "quick", "brown", "fox", "jumped", "over", "the", "lazy", "dog", "."}
	beforeArray := [nTokens]string{"", " ", " ", " ", " ", " ", " ", " ", " "}
	afterArray := [nTokens]string{" ", " ", " ", " ", " ", " ", " ", " "}
	posArray := [nTokens]string{"DT", "JJ", "JJ", "NN", "VBD", "IN", "DT", "JJ", "NN", "."}

	if txt := doc.GetText(); txt != text {
		t.Errorf("got doc text %q; want %q", txt, text)
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
			t.Errorf("got No.%d token.Word %q; want %q", i, w, wordArray[i])
		}
		if b := token.GetBefore(); b != beforeArray[i] {
			t.Errorf("got No.%d token.Before %q; want %q", i, b, beforeArray[i])
		}
		if a := token.GetAfter(); a != afterArray[i] {
			t.Errorf("got No.%d token.After %q; want %q", i, a, afterArray[i])
		}
		if p := token.GetPos(); p != posArray[i] {
			t.Errorf("got No.%d token.Pos %q; want %q", i, p, posArray[i])
		}
	}
}
