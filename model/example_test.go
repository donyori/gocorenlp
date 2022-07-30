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
	"fmt"

	"github.com/donyori/gocorenlp/model"
	"github.com/donyori/gocorenlp/model/v4.5.0-45b47e245c36/pb"
)

func ExampleDecodeResponseBody() {
	// A base64 encoded response body.
	// It carries the annotation of the text:
	//
	// Roses are red.
	//   Violets are blue.
	// Sugar is sweet.
	//   And so are you.
	//
	// with annotators "tokenize,ssplit,pos".
	respBodyBase64 := `
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

	// Retrieve the Stanford CoreNLP server response body.
	//
	// Here, we retrieve the response body from the base64 encoding string.
	// You can retrieve it from where you saved it.
	b, err := base64.StdEncoding.DecodeString(respBodyBase64)
	if err != nil {
		panic(err) // handle error
	}

	// Specify the document model.
	// Depending on your CoreNLP version, use the appropriate model.
	// See the documentation for this package for details.
	doc := new(pb.Document)

	// Decode the response body and place the result in doc.
	err = model.DecodeResponseBody(b, doc)
	if err != nil {
		panic(err) // handle error
	}

	// Work with doc.
	//
	// Here, we print the original text.
	// And then print the tokens in the last sentence
	// into a table along with their part-of-speech tags.
	fmt.Println("Original text:")
	fmt.Println(doc.GetText())
	sentences := doc.GetSentence()
	if len(sentences) == 0 {
		panic("doc.GetSentence returned an empty slice") // should not happen
	}
	fmt.Println("Last sentence tokens:")
	tokens := sentences[len(sentences)-1].GetToken()
	fmt.Println("+------+-----+")
	fmt.Println("| Word | POS |")
	fmt.Println("+------+-----+")
	for _, token := range tokens {
		fmt.Printf("| %-5s| %-4s|\n", token.GetWord(), token.GetPos())
	}
	fmt.Println("+------+-----+")

	// Output:
	// Original text:
	//
	// Roses are red.
	//   Violets are blue.
	// Sugar is sweet.
	//   And so are you.
	//
	// Last sentence tokens:
	// +------+-----+
	// | Word | POS |
	// +------+-----+
	// | And  | CC  |
	// | so   | RB  |
	// | are  | VBP |
	// | you  | PRP |
	// | .    | .   |
	// +------+-----+
}
