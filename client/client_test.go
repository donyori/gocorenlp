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
	"testing"

	"github.com/donyori/gocorenlp/model/v4.4.0-e90f30f13c40/pb"
)

// Run the following tests with a Stanford CoreNLP server running
// and listening to 127.0.0.1:9000.

func TestNew(t *testing.T) {
	c, err := New(nil)
	if err != nil {
		t.Fatal(err)
	}
	if c == nil {
		t.Error("got nil client")
	}
}

func TestClientImpl_Ready(t *testing.T) {
	c := testNewClientImpl()
	if !c.Ready() {
		t.Error("got false; want true")
	}
}

func TestClientImpl_Annotate(t *testing.T) {
	const nTokens = 10
	const text = "The quick brown fox jumped over the lazy dog."
	wordArray := [nTokens]string{"The", "quick", "brown", "fox", "jumped", "over", "the", "lazy", "dog", "."}
	beforeArray := [nTokens]string{"", " ", " ", " ", " ", " ", " ", " ", " "}
	afterArray := [nTokens]string{" ", " ", " ", " ", " ", " ", " ", " "}
	posArray := [nTokens]string{"DT", "JJ", "JJ", "NN", "VBD", "IN", "DT", "JJ", "NN", "."}

	c := testNewClientImpl()
	doc := new(pb.Document)
	if err := c.Annotate(doc, text, "tokenize,ssplit,pos"); err != nil {
		t.Fatal(err)
	}

	if txt := doc.GetText(); txt != text {
		t.Errorf("got doc text %q; want %q", txt, text)
	}

	sentences := doc.GetSentence()
	if n := len(sentences); n != 1 {
		t.Fatalf("got %d sentences; want 1", n)
	}

	tokens := sentences[0].GetToken()
	if n := len(tokens); n != nTokens {
		t.Fatalf("got %d token(s); want %d", n, nTokens)
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

// testNewClientImpl creates a Client to 127.0.0.1:9000,
// with no userinfo, no timeout,
// annotators "tokenize,ssplit,pos",
// and contentType "text/plain; charset=utf-8".
func testNewClientImpl() *clientImpl {
	return &clientImpl{
		host:        "127.0.0.1:9000",
		statusHost:  "127.0.0.1:9000",
		userinfo:    nil,
		annotators:  "tokenize,ssplit,pos",
		contentType: "text/plain; charset=utf-8",
	}
}
