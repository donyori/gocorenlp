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

package model_test

import (
	"encoding/base64"
	"testing"

	"github.com/donyori/gocorenlp/model"
	"github.com/donyori/gocorenlp/model/v4.4.0-e90f30f13c40/pb"
)

const testText = `
Roses are red.
  Violets are blue.
Sugar is sweet.
  And so are you.
`

const testBase64ResponseBody = `
jAcKRgpSb3NlcyBhcmUgcmVkLgogIFZpb2xldHMgYXJlIGJsdWUuClN1Z2FyIGlz
IHN3ZWV0LgogIEFuZCBzbyBhcmUgeW91LgoSwQEKMQoFUm9zZXMSBE5OUFMaBVJv
c2VzKgEKMgEgOgVSb3Nlc1gBYAaIAQCQAQGoAQCwAgAKKgoDYXJlEgNWQlAaA2Fy
ZSoBIDIBIDoDYXJlWAdgCogBAZABAqgBALACAAooCgNyZWQSAkpKGgNyZWQqASAy
ADoDcmVkWAtgDogBApABA6gBALACAAojCgEuEgEuGgEuKgAyAwogIDoBLlgOYA+I
AQOQAQSoAQCwAgAQABgEIAAoATAPmAMAsAMAiAQAEskBCjgKB1Zpb2xldHMSA05O
UxoHVmlvbGV0cyoDCiAgMgEgOgdWaW9sZXRzWBJgGYgBBJABBagBALACAAoqCgNh
cmUSA1ZCUBoDYXJlKgEgMgEgOgNhcmVYGmAdiAEFkAEGqAEAsAIACisKBGJsdWUS
AkpKGgRibHVlKgEgMgA6BGJsdWVYHmAiiAEGkAEHqAEAsAIACiEKAS4SAS4aAS4q
ADIBCjoBLlgiYCOIAQeQAQioAQCwAgAQBBgIIAEoEjAjmAMAsAMAiAQAEsMBCjAK
BVN1Z2FyEgNOTlAaBVN1Z2FyKgEKMgEgOgVTdWdhclgkYCmIAQiQAQmoAQCwAgAK
JwoCaXMSA1ZCWhoCaXMqASAyASA6AmlzWCpgLIgBCZABCqgBALACAAouCgVzd2Vl
dBICSkoaBXN3ZWV0KgEgMgA6BXN3ZWV0WC1gMogBCpABC6gBALACAAojCgEuEgEu
GgEuKgAyAwogIDoBLlgyYDOIAQuQAQyoAQCwAgAQCBgMIAIoJDAzmAMAsAMAiAQA
EuIBCisKA0FuZBICQ0MaA0FuZCoDCiAgMgEgOgNBbmRYNmA5iAEMkAENqAEAsAIA
CiYKAnNvEgJSQhoCc28qASAyASA6AnNvWDpgPIgBDZABDqgBALACAAoqCgNhcmUS
A1ZCUBoDYXJlKgEgMgEgOgNhcmVYPWBAiAEOkAEPqAEAsAIACikKA3lvdRIDUFJQ
GgN5b3UqASAyADoDeW91WEFgRIgBD5ABEKgBALACAAohCgEuEgEuGgEuKgAyAQo6
AS5YRGBFiAEQkAERqAEAsAIAEAwYESADKDYwRZgDALADAIgEAFgAaAB4AIABAA==
`

func TestDecodeResponseBody(t *testing.T) {
	data, err := base64.StdEncoding.DecodeString(testBase64ResponseBody)
	if err != nil {
		t.Fatal(err)
	}
	doc := new(pb.Document)
	if err = model.DecodeResponseBody(data, doc); err != nil {
		t.Error(err)
	}
	if txt := doc.GetText(); txt != testText {
		t.Errorf("got doc text %q; want %q", txt, testText)
	}

	const nSentences = 4
	nTokens := [nSentences]int{4, 4, 4, 5}
	wordLists := [nSentences][]string{
		{"Roses", "are", "red", "."},
		{"Violets", "are", "blue", "."},
		{"Sugar", "is", "sweet", "."},
		{"And", "so", "are", "you", "."},
	}
	gapLists := [nSentences][]string{
		{"\n", " ", " ", "", "\n  "},
		{"\n  ", " ", " ", "", "\n"},
		{"\n", " ", " ", "", "\n  "},
		{"\n  ", " ", " ", " ", "", "\n"},
	}
	posLists := [nSentences][]string{
		{"NNPS", "VBP", "JJ", "."},
		{"NNS", "VBP", "JJ", "."},
		{"NNP", "VBZ", "JJ", "."},
		{"CC", "RB", "VBP", "PRP", "."},
	}

	sentences := doc.GetSentence()
	if n := len(sentences); n != nSentences {
		t.Fatalf("got %d sentence(s); want %d", n, nSentences)
	}
	for si, sentence := range sentences {
		tokens := sentence.GetToken()
		if n := len(tokens); n != nTokens[si] {
			t.Errorf("sentence %d: got %d token(s); want %d", si, n, nTokens[si])
			continue
		}
		for ti, token := range tokens {
			if w := token.GetWord(); w != wordLists[si][ti] {
				t.Errorf("Sentence %d, Token %d: got Word %q; want %q", si, ti, w, wordLists[si][ti])
			}
			if b := token.GetBefore(); b != gapLists[si][ti] {
				t.Errorf("Sentence %d, Token %d: got Before %q; want %q", si, ti, b, gapLists[si][ti])
			}
			if a := token.GetAfter(); a != gapLists[si][ti+1] {
				t.Errorf("Sentence %d, Token %d; got After %q; want %q", si, ti, a, gapLists[si][ti+1])
			}
			if p := token.GetPos(); p != posLists[si][ti] {
				t.Errorf("Sentence %d, Token %d; got POS %q; want %q", si, ti, p, posLists[si][ti])
			}
		}
	}
}
