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

const testText = "The quick brown fox jumped over the lazy dog."

// testAnnotateFunc encapsulates common code for testing
// the method Annotate of *clientImpl.
func testAnnotateFunc(t *testing.T, f func() *clientImpl) {
	testAnnotateMethodsFunc(t, func(t *testing.T, annotators string) *pb.Document {
		c := f()
		reader := strings.NewReader(testText)
		doc := new(pb.Document)
		if err := c.Annotate(reader, annotators, doc); err != nil {
			t.Error(err)
			return nil
		}
		return doc
	})
}

// testAnnotateStringFunc encapsulates common code for testing
// the method AnnotateString of *clientImpl.
func testAnnotateStringFunc(t *testing.T, f func() *clientImpl) {
	testAnnotateMethodsFunc(t, func(t *testing.T, annotators string) *pb.Document {
		c := f()
		doc := new(pb.Document)
		if err := c.AnnotateString(testText, annotators, doc); err != nil {
			t.Error(err)
			return nil
		}
		return doc
	})
}

// testAnnotateMethodsFunc encapsulates common code for testing
// the methods Annotate and AnnotateString of *clientImpl.
func testAnnotateMethodsFunc(t *testing.T, f func(t *testing.T, annotators string) *pb.Document) {
	testCases := []struct {
		name       string
		annotators string
	}{
		{"specify annotators", "tokenize,ssplit,pos"},
		{"omit annotators", ""},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			doc := f(t, tc.annotators)
			if doc != nil {
				testCheckAnnotation(t, doc)
			}
		})
	}
}

// testCheckAnnotation checks the result of annotation to the text:
//  The quick brown fox jumped over the lazy dog.
//
// It checks the document text, sentence split, token word,
// content before token, content after token, and token part-of-speech tag.
func testCheckAnnotation(t *testing.T, doc *pb.Document) {
	const nTokens = 10
	wordArray := [nTokens]string{"The", "quick", "brown", "fox", "jumped", "over", "the", "lazy", "dog", "."}
	beforeArray := [nTokens]string{"", " ", " ", " ", " ", " ", " ", " ", " "}
	afterArray := [nTokens]string{" ", " ", " ", " ", " ", " ", " ", " "}
	posArray := [nTokens]string{"DT", "JJ", "JJ", "NN", "VBD", "IN", "DT", "JJ", "NN", "."}

	if txt := doc.GetText(); txt != testText {
		t.Errorf("got doc text %q; want %q", txt, testText)
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