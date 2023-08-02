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

package model_test

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"

	"github.com/donyori/gocorenlp/model"
	"github.com/donyori/gocorenlp/model/v4.5.3-5250f9faf9f1/pb"
)

func ExampleDecodeResponseBody() {
	// A standard base64 (as defined in RFC 4648) encoded response body.
	// It carries the annotation of the text:
	//
	// Roses are red.
	//   Violets are blue.
	// Sugar is sweet.
	//   And so are you.
	//
	// with annotators "tokenize,ssplit,pos".
	const RespBodyBase64 = `
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
	// Here, we retrieve the response body from the base64 string.
	// You can retrieve it from where you saved it.
	b, err := base64.StdEncoding.DecodeString(RespBodyBase64)
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

func ExampleResponseBodyDecoder() {
	// A standard base64 (as defined in RFC 4648) encoded response body.
	// It carries the annotation of the text:
	//
	//	Roses are red.
	//	  Violets are blue.
	//	Sugar is sweet.
	//	  And so are you.
	//
	// with annotators "tokenize,ssplit,pos".
	const RosesRespBodyBase64 = `
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

	// A standard base64 (as defined in RFC 4648) encoded response body.
	// It carries the annotation of the text:
	//	Yesterday is history. Tomorrow is a mystery. Today is a gift. That’s why we call it 'The Present'
	// with annotators "tokenize,ssplit,pos".
	const YesterdayRespBodyBase64 = `
5AkKY1llc3RlcmRheSBpcyBoaXN0b3J5LiBUb21vcnJvdyBpcyBhIG15c3Rlcnku
IFRvZGF5IGlzIGEgZ2lmdC4gVGhhdOKAmXMgd2h5IHdlIGNhbGwgaXQgJ1RoZSBQ
cmVzZW50JxLRAQo6CglZZXN0ZXJkYXkSAk5OGglZZXN0ZXJkYXkqADIBIDoJWWVz
dGVyZGF5WABgCYgBAJABAagBALACAAonCgJpcxIDVkJaGgJpcyoBIDIBIDoCaXNY
CmAMiAEBkAECqAEAsAIACjQKB2hpc3RvcnkSAk5OGgdoaXN0b3J5KgEgMgA6B2hp
c3RvcnlYDWAUiAECkAEDqAEAsAIACiEKAS4SAS4aAS4qADIBIDoBLlgUYBWIAQOQ
AQSoAQCwAgAQABgEIAAoADAVmAMAsAMAiAQAEvQBCjgKCFRvbW9ycm93EgJOThoI
VG9tb3Jyb3cqASAyASA6CFRvbW9ycm93WBZgHogBBJABBagBALACAAonCgJpcxID
VkJaGgJpcyoBIDIBIDoCaXNYH2AhiAEFkAEGqAEAsAIACiMKAWESAkRUGgFhKgEg
MgEgOgFhWCJgI4gBBpABB6gBALACAAo0CgdteXN0ZXJ5EgJOThoHbXlzdGVyeSoB
IDIAOgdteXN0ZXJ5WCRgK4gBB5ABCKgBALACAAohCgEuEgEuGgEuKgAyASA6AS5Y
K2AsiAEIkAEJqAEAsAIAEAQYCSABKBYwLJgDALADAIgEABLiAQovCgVUb2RheRIC
Tk4aBVRvZGF5KgEgMgEgOgVUb2RheVgtYDKIAQmQAQqoAQCwAgAKJwoCaXMSA1ZC
WhoCaXMqASAyASA6AmlzWDNgNYgBCpABC6gBALACAAojCgFhEgJEVBoBYSoBIDIB
IDoBYVg2YDeIAQuQAQyoAQCwAgAKKwoEZ2lmdBICTk4aBGdpZnQqASAyADoEZ2lm
dFg4YDyIAQyQAQ2oAQCwAgAKIQoBLhIBLhoBLioAMgEgOgEuWDxgPYgBDZABDqgB
ALACABAJGA4gAigtMD2YAwCwAwCIBAASwwMKKwoEVGhhdBICSU4aBFRoYXQqASAy
ADoEVGhhdFg+YEKIAQ6QAQ+oAQCwAgAKLAoE4oCZcxIDUE9TGgTigJlzKgAyASA6
BOKAmXNYQmBEiAEPkAEQqAEAsAIACioKA3doeRIDV1JCGgN3aHkqASAyASA6A3do
eVhFYEiIARCQARGoAQCwAgAKJwoCd2USA1BSUBoCd2UqASAyASA6AndlWElgS4gB
EZABEqgBALACAAotCgRjYWxsEgNWQlAaBGNhbGwqASAyASA6BGNhbGxYTGBQiAES
kAETqAEAsAIACicKAml0EgNQUlAaAml0KgEgMgEgOgJpdFhRYFOIAROQARSoAQCw
AgAKIgoBJxICYGAaAScqASAyADoBJ1hUYFWIARSQARWoAQCwAgAKKAoDVGhlEgJE
VBoDVGhlKgAyASA6A1RoZVhVYFiIARWQARaoAQCwAgAKNQoHUHJlc2VudBIDTk5Q
GgdQcmVzZW50KgEgMgA6B1ByZXNlbnRYWWBgiAEWkAEXqAEAsAIACiEKAScSAicn
GgEnKgAyADoBJ1hgYGGIAReQARioAQCwAgAQDhgYIAMoPjBhmAMAsAMAiAQAWABo
AHgAgAEA
`

	// Retrieve the Stanford CoreNLP server response body.
	//
	// Here, we retrieve the response body from the base64 string.
	// You can retrieve it from where you saved it.
	roses, err := base64.StdEncoding.DecodeString(RosesRespBodyBase64)
	if err != nil {
		panic(err) // handle error
	}
	yesterday, err := base64.StdEncoding.DecodeString(YesterdayRespBodyBase64)
	if err != nil {
		panic(err) // handle error
	}

	// Concatenate two annotation results as the input stream.
	b := make([]byte, len(roses)+len(yesterday))
	copy(b[copy(b, roses):], yesterday)
	input := bytes.NewReader(b)

	// Create a ResponseBodyDecoder on it.
	dec := model.NewResponseBodyDecoder(input)

	// Decode the annotation results until EOF.
	for {
		// Specify the document model.
		// Depending on your CoreNLP version, use the appropriate model.
		// See the documentation for this package for details.
		doc := new(pb.Document)

		// Decode the response body and place the result in doc.
		err = dec.Decode(doc)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err) // handle error
		}

		// Work with doc.
		//
		// Here, we print the original text.
		// And then print the tokens in the first sentence
		// into a table along with their part-of-speech tags.
		fmt.Println("Original text:")
		fmt.Println(doc.GetText())
		sentences := doc.GetSentence()
		if len(sentences) == 0 {
			panic("doc.GetSentence returned an empty slice") // should not happen
		}
		fmt.Println("First sentence tokens:")
		tokens := sentences[0].GetToken()
		fmt.Println("+-----------+------+")
		fmt.Println("| Word      | POS  |")
		fmt.Println("+-----------+------+")
		for _, token := range tokens {
			fmt.Printf("| %-10s| %-5s|\n", token.GetWord(), token.GetPos())
		}
		fmt.Println("+-----------+------+")
	}

	// Output:
	// Original text:
	//
	// Roses are red.
	//   Violets are blue.
	// Sugar is sweet.
	//   And so are you.
	//
	// First sentence tokens:
	// +-----------+------+
	// | Word      | POS  |
	// +-----------+------+
	// | Roses     | NNPS |
	// | are       | VBP  |
	// | red       | JJ   |
	// | .         | .    |
	// +-----------+------+
	// Original text:
	// Yesterday is history. Tomorrow is a mystery. Today is a gift. That’s why we call it 'The Present'
	// First sentence tokens:
	// +-----------+------+
	// | Word      | POS  |
	// +-----------+------+
	// | Yesterday | NN   |
	// | is        | VBZ  |
	// | history   | NN   |
	// | .         | .    |
	// +-----------+------+
}
