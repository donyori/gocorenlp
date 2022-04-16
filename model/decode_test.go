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
	"testing"

	"github.com/donyori/gocorenlp/internal/pbtest"
	"github.com/donyori/gocorenlp/model/pb"
)

func TestDecodeResponseBody(t *testing.T) {
	if err := pbtest.CheckDocumentFromBase64(Base64ResponseBody, new(pb.Doc440)); err != nil {
		t.Error(err)
	}
}

const Base64ResponseBody = `
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
