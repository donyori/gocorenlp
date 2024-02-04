// gocorenlp.  A Go (Golang) client for Stanford CoreNLP server.
// Copyright (C) 2022-2024  Yuan Gao
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
	"io"
	"testing"

	gogoerrors "github.com/donyori/gogo/errors"
	"google.golang.org/protobuf/encoding/protowire"

	"github.com/donyori/gocorenlp/errors"
	"github.com/donyori/gocorenlp/internal/pbtest"
	"github.com/donyori/gocorenlp/model"
	"github.com/donyori/gocorenlp/model/v4.5.6-eb50467fa8e3/pb"
)

func TestDecodeMessage_Basic(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(RosesResp)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	msgData, n := protowire.ConsumeBytes(respBody)
	if n < 0 {
		t.Fatal("failed to decode response body:", protowire.ParseError(n))
	}
	doc := new(pb.Document)
	err = model.DecodeMessage(msgData, doc)
	if err != nil {
		t.Fatal(err)
	}
	err = pbtest.CheckRosesAreRedDocument(doc)
	if err != nil {
		t.Error(err)
	}
}

func TestDecodeMessage_AppendOneZeroByte(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(RosesResp)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	msgData, n := protowire.ConsumeBytes(respBody)
	if n < 0 {
		t.Fatal("failed to decode response body:", protowire.ParseError(n))
	}
	b := make([]byte, len(msgData)+1)
	copy(b, msgData)
	doc := new(pb.Document)
	err = model.DecodeMessage(b, doc)
	if !errors.IsProtoBufError(err) {
		t.Errorf("got %v; want a *ProtoBufError", err)
	}
}

func TestDecodeMessage_PrependOneZeroByte(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(RosesResp)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	msgData, n := protowire.ConsumeBytes(respBody)
	if n < 0 {
		t.Fatal("failed to decode response body:", protowire.ParseError(n))
	}
	b := make([]byte, 1+len(msgData))
	copy(b[1:], msgData)
	doc := new(pb.Document)
	err = model.DecodeMessage(b, doc)
	if !errors.IsProtoBufError(err) {
		t.Errorf("got %v; want a *ProtoBufError", err)
	}
}

func TestDecodeMessage_Truncated(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(RosesResp)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	msgData, n := protowire.ConsumeBytes(respBody)
	if n < 0 {
		t.Fatal("failed to decode response body:", protowire.ParseError(n))
	}
	b := make([]byte, len(msgData)/2)
	copy(b, msgData)
	doc := new(pb.Document)
	err = model.DecodeMessage(b, doc)
	if !errors.IsProtoBufError(err) {
		t.Errorf("got %v; want a *ProtoBufError", err)
	}
}

func TestDecodeResponseBody_Basic(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(RosesResp)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	doc := new(pb.Document)
	err = model.DecodeResponseBody(respBody, doc)
	if err != nil {
		t.Fatal(err)
	}
	err = pbtest.CheckRosesAreRedDocument(doc)
	if err != nil {
		t.Error(err)
	}
}

func TestDecodeResponseBody_AppendSuffix(t *testing.T) {
	const Suffix = "Go client for Stanford CoreNLP server"
	respBody, err := base64.StdEncoding.DecodeString(RosesResp)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	b := make([]byte, len(respBody)+len(Suffix))
	copy(b[copy(b, respBody):], Suffix)
	doc := new(pb.Document)
	err = model.DecodeResponseBody(b, doc)
	if err != nil {
		t.Fatal(err)
	}
	err = pbtest.CheckRosesAreRedDocument(doc)
	if err != nil {
		t.Error(err)
	}
}

func TestDecodeResponseBody_PrependOneZeroByte(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(RosesResp)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	b := make([]byte, 1+len(respBody))
	copy(b[1:], respBody)
	doc := new(pb.Document)
	err = model.DecodeResponseBody(b, doc)
	if !errors.IsProtoBufError(err) {
		t.Errorf("got %v; want a *ProtoBufError", err)
	}
}

func TestDecodeResponseBody_Truncated(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(RosesResp)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	b := make([]byte, len(respBody)/2)
	copy(b, respBody)
	doc := new(pb.Document)
	err = model.DecodeResponseBody(b, doc)
	if !errors.IsProtoBufError(err) ||
		!errors.Is(err, io.ErrUnexpectedEOF) {
		t.Errorf("got %v; want a *ProtoBufError wrapping io.ErrUnexpectedEOF",
			err)
	}
}

func TestConsumeResponseBody_Basic(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(RosesResp)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	doc := new(pb.Document)
	n, err := model.ConsumeResponseBody(respBody, doc)
	if err != nil {
		t.Fatal(err)
	} else if n != len(respBody) {
		t.Errorf("got n %d; want %d", n, len(respBody))
	}
	err = pbtest.CheckRosesAreRedDocument(doc)
	if err != nil {
		t.Error(err)
	}
}

func TestConsumeResponseBody_AppendSuffix(t *testing.T) {
	const Suffix = "Go client for Stanford CoreNLP server"
	respBody, err := base64.StdEncoding.DecodeString(RosesResp)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	b := make([]byte, len(respBody)+len(Suffix))
	copy(b[copy(b, respBody):], Suffix)
	doc := new(pb.Document)
	n, err := model.ConsumeResponseBody(b, doc)
	if err != nil {
		t.Fatal(err)
	} else if n != len(respBody) {
		t.Errorf("got n %d; want %d", n, len(respBody))
	}
	err = pbtest.CheckRosesAreRedDocument(doc)
	if err != nil {
		t.Error(err)
	}
}

func TestConsumeResponseBody_PrependOneZeroByte(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(RosesResp)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	b := make([]byte, 1+len(respBody))
	copy(b[1:], respBody)
	doc := new(pb.Document)
	n, err := model.ConsumeResponseBody(b, doc)
	// The zero byte is treated as the prefixed length,
	// representing a length of 0.
	// model.ConsumeResponseBody reads this byte (length: 1)
	// and then decodes an empty response body.
	// Therefore, n should be 1 here.
	if n != 1 {
		t.Errorf("got %d; want 1", n)
	}
	if !errors.IsProtoBufError(err) {
		t.Errorf("got %v; want a *ProtoBufError", err)
	}
}

func TestConsumeResponseBody_Truncated(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(RosesResp)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	b := make([]byte, len(respBody)/2)
	copy(b, respBody)
	doc := new(pb.Document)
	n, err := model.ConsumeResponseBody(b, doc)
	// n should be -1 here to represent truncated data,
	// corresponding to io.ErrUnexpectedEOF.
	if n != -1 {
		t.Errorf("got %d; want -1", n)
	}
	if !errors.IsProtoBufError(err) ||
		!errors.Is(err, io.ErrUnexpectedEOF) {
		t.Errorf("got %v; want a *ProtoBufError wrapping io.ErrUnexpectedEOF",
			err)
	}
}

func TestConsumeResponseBody_TwoResponses(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(RosesResp)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	b := make([]byte, len(respBody)*2)
	copy(b[copy(b, respBody):], respBody)
	doc := new(pb.Document)
	n1, err := model.ConsumeResponseBody(b, doc)
	if err != nil {
		t.Fatal(err)
	} else if n1 != len(respBody) {
		t.Errorf("got n %d; want %d", n1, len(respBody))
	}
	err = pbtest.CheckRosesAreRedDocument(doc)
	if err != nil {
		t.Error(err)
	}
	if n1 != len(respBody) {
		return
	}
	doc = new(pb.Document)
	n2, err := model.ConsumeResponseBody(b[n1:], doc)
	if err != nil {
		t.Fatal(err)
	} else if n2 != len(respBody) {
		t.Errorf("got n %d; want %d", n2, len(respBody))
	}
	err = pbtest.CheckRosesAreRedDocument(doc)
	if err != nil {
		t.Error(err)
	}
}

func TestConsumeResponseBody_DifferentResponses(t *testing.T) {
	const NumRepeat int = 3
	data, lens, err := MakeDifferentResponsesData(NumRepeat)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}

	p := data[:]
	for i := 0; i < NumRepeat; i++ {
		doc := new(pb.Document)
		n, err := model.ConsumeResponseBody(p, doc)
		if err != nil {
			t.Fatalf("Round %d: %v", i+1, err)
		} else if n != lens[0] {
			t.Fatalf("Round %d: got n %d; want %d", i+1, n, lens[0])
		}
		err = pbtest.CheckRosesAreRedDocument(doc)
		if err != nil {
			t.Fatalf("Round %d: %v", i+1, err)
		}
		p = p[n:]

		doc = new(pb.Document)
		n, err = model.ConsumeResponseBody(p, doc)
		if err != nil {
			t.Fatalf("Round %d: %v", i+1, err)
		} else if n != lens[1] {
			t.Fatalf("Round %d: got n %d; want %d", i+1, n, lens[1])
		} else if text := doc.GetText(); text != YesterdayIsHistory {
			t.Fatalf("Round %d: got text %q; want %q",
				i+1, text, YesterdayIsHistory)
		}
		p = p[n:]

		doc = new(pb.Document)
		n, err = model.ConsumeResponseBody(p, doc)
		if err != nil {
			t.Fatalf("Round %d: %v", i+1, err)
		} else if n != lens[2] {
			t.Fatalf("Round %d: got n %d; want %d", i+1, n, lens[2])
		}
		err = pbtest.CheckRosesAreRedDocument(doc)
		if err != nil {
			t.Fatalf("Round %d: %v", i+1, err)
		}
		p = p[n:]

		doc = new(pb.Document)
		n, err = model.ConsumeResponseBody(p, doc)
		if err != nil {
			t.Fatalf("Round %d: %v", i+1, err)
		} else if n != lens[3] {
			t.Fatalf("Round %d: got n %d; want %d",
				i+1, n, lens[3])
		} else if text := doc.GetText(); text != YesterdayIsHistory {
			t.Fatalf("Round %d: got text %q; want %q",
				i+1, text, YesterdayIsHistory)
		}
		p = p[n:]
	}
}

const (
	YesterdayIsHistory = `Yesterday is history. Tomorrow is a mystery. Today is a gift. That’s why we call it 'The Present'`

	// RosesResp is the standard base64 (as defined in RFC 4648)
	// encoded response of annotating
	//
	//	Roses are red.
	//	  Violets are blue.
	//	Sugar is sweet.
	//	  And so are you.
	//
	// with the server default annotators by Stanford CoreNLP 4.5.6.
	RosesResp = pbtest.RosesAreRedRespV456

	// YesterdayResp is the standard base64 (as defined in RFC 4648)
	// encoded response of annotating
	//
	//	Yesterday is history. Tomorrow is a mystery. Today is a gift. That’s why we call it 'The Present'
	//
	// with the server default annotators by Stanford CoreNLP 4.5.6.
	YesterdayResp = `
kVMKY1llc3RlcmRheSBpcyBoaXN0b3J5LiBUb21vcnJvdyBpcyBhIG15c3Rlcnku
IFRvZGF5IGlzIGEgZ2lmdC4gVGhhdOKAmXMgd2h5IHdlIGNhbGwgaXQgJ1RoZSBQ
cmVzZW50JxKvDAqlAQoJWWVzdGVyZGF5EgJOThoJWWVzdGVyZGF5KgAyASA6CVll
c3RlcmRheUIEREFURUoLT0ZGU0VUIFAtMURSCXllc3RlcmRheVgAYAloAHIEUEVS
MIgBAJABAZoBIhILT0ZGU0VUIFAtMUQaCVllc3RlcmRheSIEREFURSoCdDGoAQCw
AgDyAwREQVRF+gMEREFURYAEAIgEAJIECURBVEU9LTEuMApVCgJpcxIDVkJaGgJp
cyoBIDIBIDoCaXNCAU9SAmJlWApgDGgAcgRQRVIwiAEBkAECqAEAsAIA8gMBT/oD
AU+SBBRPPTAuOTk5OTk5Njg0MzU4ODMzNgpqCgdoaXN0b3J5EgJOThoHaGlzdG9y
eSoBIDIAOgdoaXN0b3J5QgFPUgdoaXN0b3J5WA1gFGgAcgRQRVIwiAECkAEDqAEA
sAIA8gMBT/oDAU+ABAGSBBRPPTAuOTk5OTUzOTc0MjA0ODk4NgpOCgEuEgEuGgEu
KgAyASA6AS5CAU9SAS5YFGAVaAByBFBFUjCIAQOQAQSoAQCwAgDyAwFP+gMBT5IE
FE89MC45OTk5OTk1NTk0NzY5NzMzEAAYBCAAKAAwFUJbCgQIABADCgQIABABCgQI
ABACCgQIABAEEhMIAxABGgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4
CRITCAMQBBoFcHVuY3QgACgAMAA4CRoBAyoBAEpbCgQIABADCgQIABABCgQIABAC
CgQIABAEEhMIAxABGgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRIT
CAMQBBoFcHVuY3QgACgAMAA4CRoBAyoBAFJbCgQIABADCgQIABABCgQIABACCgQI
ABAEEhMIAxABGgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRITCAMQ
BBoFcHVuY3QgACgAMAA4CRoBAyoBAFgBigFbCgQIABADCgQIABABCgQIABACCgQI
ABAEEhMIAxABGgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRITCAMQ
BBoFcHVuY3QgACgAMAA4CRoBAyoBAJIBWwoECAAQAwoECAAQAQoECAAQAgoECAAQ
BBITCAMQARoFbnN1YmogACgAMAA4CRIRCAMQAhoDY29wIAAoADAAOAkSEwgDEAQa
BXB1bmN0IAAoADAAOAkaAQMqAQCYAwCwAwC6A1IIABAAGAEiBERBVEUqC09GRlNF
VCBQLTFEMgREQVRFOiISC09GRlNFVCBQLTFEGglZZXN0ZXJkYXkiBERBVEUqAnQx
UABYAGIJWWVzdGVyZGF5wgOVAggAEgZQUk9QRVIaCFNJTkdVTEFSIgdVTktOT1dO
KglJTkFOSU1BVEUyB1VOS05PV044AEgBUABaCXllc3RlcmRheWIEREFURWj/////
//////8BcP///////////wF4AIABAIgBAJABAJgBAaABAKgBALABALgBAMABAMgB
ANABANgBAeABAegBAPIBDwj///////////8BEAAgAPoBFgj///////////8BEP//
/////////wGCAg0I////////////ARAAkgMNCP///////////wEQAJIDDQj/////
//////8BEAGSAw0I////////////ARACkgMNCP///////////wEQA5oDDQj/////
//////8BEADCA5QCCAESB05PTUlOQUwaCFNJTkdVTEFSIgdORVVUUkFMKglJTkFO
SU1BVEUyB1VOS05PV044AkgDUAJaB2hpc3RvcnliAU9o////////////AXD/////
//////8BeAGAAQGIAQCQAQCYAQGgAQCoAQCwAQC4AQDAAQDIAQDQAQDYAQHgAQHo
AQDyAQ8I////////////ARACIAD6ARYI////////////ARD///////////8BggIN
CP///////////wEQApIDDQj///////////8BEACSAw0I////////////ARABkgMN
CP///////////wEQApIDDQj///////////8BEAOaAw0I////////////ARACuAMA
yAMBiAQBoAQBEo0NCp8BCghUb21vcnJvdxICTk4aCFRvbW9ycm93KgEgMgEgOghU
b21vcnJvd0IEREFURUoKT0ZGU0VUIFAxRFIIdG9tb3Jyb3dYFmAeaAByBFBFUjCI
AQSQAQWaASASCk9GRlNFVCBQMUQaCFRvbW9ycm93IgREQVRFKgJ0MqgBALACAPID
BERBVEX6AwREQVRFgAQCiAQBkgQJREFURT0tMS4wClUKAmlzEgNWQloaAmlzKgEg
MgEgOgJpc0IBT1ICYmVYH2AhaAByBFBFUjCIAQWQAQaoAQCwAgDyAwFP+gMBT5IE
FE89MC45OTk5OTk1MTcxMjgwMzQ3ClMKAWESAkRUGgFhKgEgMgEgOgFhQgFPUgFh
WCJgI2gAcgRQRVIwiAEGkAEHqAEAsAIA8gMBT/oDAU+ABAOSBBRPPTAuOTk5OTk5
MzQ0MDIyMTAyNApqCgdteXN0ZXJ5EgJOThoHbXlzdGVyeSoBIDIAOgdteXN0ZXJ5
QgFPUgdteXN0ZXJ5WCRgK2gAcgRQRVIwiAEHkAEIqAEAsAIA8gMBT/oDAU+ABAOS
BBRPPTAuOTk5OTc3NTIzOTU4Mzk4OQpOCgEuEgEuGgEuKgAyASA6AS5CAU9SAS5Y
K2AsaAByBFBFUjCIAQiQAQmoAQCwAgDyAwFP+gMBT5IEFE89MC45OTk5OTk5NDc4
MTA3MzY5EAQYCSABKBYwLEJ0CgQIARAECgQIARABCgQIARACCgQIARADCgQIARAF
EhMIBBABGgVuc3ViaiAAKAAwADgJEhEIBBACGgNjb3AgACgAMAA4CRIRCAQQAxoD
ZGV0IAAoADAAOAkSEwgEEAUaBXB1bmN0IAAoADAAOAkaAQQqAQBKdAoECAEQBAoE
CAEQAQoECAEQAgoECAEQAwoECAEQBRITCAQQARoFbnN1YmogACgAMAA4CRIRCAQQ
AhoDY29wIAAoADAAOAkSEQgEEAMaA2RldCAAKAAwADgJEhMIBBAFGgVwdW5jdCAA
KAAwADgJGgEEKgEAUnQKBAgBEAQKBAgBEAEKBAgBEAIKBAgBEAMKBAgBEAUSEwgE
EAEaBW5zdWJqIAAoADAAOAkSEQgEEAIaA2NvcCAAKAAwADgJEhEIBBADGgNkZXQg
ACgAMAA4CRITCAQQBRoFcHVuY3QgACgAMAA4CRoBBCoBAFgBigF0CgQIARAECgQI
ARABCgQIARACCgQIARADCgQIARAFEhMIBBABGgVuc3ViaiAAKAAwADgJEhEIBBAC
GgNjb3AgACgAMAA4CRIRCAQQAxoDZGV0IAAoADAAOAkSEwgEEAUaBXB1bmN0IAAo
ADAAOAkaAQQqAQCSAXQKBAgBEAQKBAgBEAEKBAgBEAIKBAgBEAMKBAgBEAUSEwgE
EAEaBW5zdWJqIAAoADAAOAkSEQgEEAIaA2NvcCAAKAAwADgJEhEIBBADGgNkZXQg
ACgAMAA4CRITCAQQBRoFcHVuY3QgACgAMAA4CRoBBCoBAJgDALADALoDTggBEAQY
BSIEREFURSoKT0ZGU0VUIFAxRDIEREFURTogEgpPRkZTRVQgUDFEGghUb21vcnJv
dyIEREFURSoCdDJQAVgBYghUb21vcnJvd8ID3AEIAhIGUFJPUEVSGghTSU5HVUxB
UiIHVU5LTk9XTioJSU5BTklNQVRFMgdVTktOT1dOOABIAVAAWgh0b21vcnJvd2IE
REFURWj///////////8BcP///////////wF4AoABAogBAZABAJgBAaABAKgBALAB
ALgBAMABAMgBANABANgBAeABAegBAPIBBggAEAAgAPoBFgj///////////8BEP//
/////////wGCAgQIABAAkgMECAAQAJIDBAgAEAGSAwQIABACkgMECAAQA5IDBAgA
EASaAwQIABAAwgPjAQgDEgdOT01JTkFMGghTSU5HVUxBUiIHTkVVVFJBTCoJSU5B
TklNQVRFMgdVTktOT1dOOAJIBFADWgdteXN0ZXJ5YgFPaP///////////wFw////
////////AXgDgAEDiAEBkAEAmAEBoAEAqAEAsAEAuAEAwAEAyAEA0AEA2AEB4AEB
6AEA8gEGCAAQAyAA+gEWCP///////////wEQ////////////AYICBAgAEAOSAwQI
ABAAkgMECAAQAZIDBAgAEAKSAwQIABADkgMECAAQBJoDBAgAEAKaAwQIABADuAMC
yAMBiAQBoAQBEt4MCowBCgVUb2RheRICTk4aBVRvZGF5KgEgMgEgOgVUb2RheUIE
REFURUoIVEhJUyBQMURSBXRvZGF5WC1gMmgAcgRQRVIwiAEJkAEKmgEbEghUSElT
IFAxRBoFVG9kYXkiBERBVEUqAnQzqAEAsAIA8gMEREFURfoDBERBVEWABASIBAKS
BAlEQVRFPS0xLjAKVQoCaXMSA1ZCWhoCaXMqASAyASA6AmlzQgFPUgJiZVgzYDVo
AHIEUEVSMIgBCpABC6gBALACAPIDAU/6AwFPkgQUTz0wLjk5OTk5OTMyOTAxNTA2
NjIKUwoBYRICRFQaAWEqASAyASA6AWFCAU9SAWFYNmA3aAByBFBFUjCIAQuQAQyo
AQCwAgDyAwFP+gMBT4AEBZIEFE89MC45OTk5OTgzNzQyNDI3OTI4Cl4KBGdpZnQS
Ak5OGgRnaWZ0KgEgMgA6BGdpZnRCAU9SBGdpZnRYOGA8aAByBFBFUjCIAQyQAQ2o
AQCwAgDyAwFP+gMBT4AEBZIEFE89MC45OTk5ODIzMTQ2NzI2NDY2Ck4KAS4SAS4a
AS4qADIBIDoBLkIBT1IBLlg8YD1oAHIEUEVSMIgBDZABDqgBALACAPIDAU/6AwFP
kgQUTz0wLjk5OTk5OTg4MzI3MDQ0MDMQCRgOIAIoLTA9QnQKBAgCEAQKBAgCEAEK
BAgCEAIKBAgCEAMKBAgCEAUSEwgEEAEaBW5zdWJqIAAoADAAOAkSEQgEEAIaA2Nv
cCAAKAAwADgJEhEIBBADGgNkZXQgACgAMAA4CRITCAQQBRoFcHVuY3QgACgAMAA4
CRoBBCoBAEp0CgQIAhAECgQIAhABCgQIAhACCgQIAhADCgQIAhAFEhMIBBABGgVu
c3ViaiAAKAAwADgJEhEIBBACGgNjb3AgACgAMAA4CRIRCAQQAxoDZGV0IAAoADAA
OAkSEwgEEAUaBXB1bmN0IAAoADAAOAkaAQQqAQBSdAoECAIQBAoECAIQAQoECAIQ
AgoECAIQAwoECAIQBRITCAQQARoFbnN1YmogACgAMAA4CRIRCAQQAhoDY29wIAAo
ADAAOAkSEQgEEAMaA2RldCAAKAAwADgJEhMIBBAFGgVwdW5jdCAAKAAwADgJGgEE
KgEAWAGKAXQKBAgCEAQKBAgCEAEKBAgCEAIKBAgCEAMKBAgCEAUSEwgEEAEaBW5z
dWJqIAAoADAAOAkSEQgEEAIaA2NvcCAAKAAwADgJEhEIBBADGgNkZXQgACgAMAA4
CRITCAQQBRoFcHVuY3QgACgAMAA4CRoBBCoBAJIBdAoECAIQBAoECAIQAQoECAIQ
AgoECAIQAwoECAIQBRITCAQQARoFbnN1YmogACgAMAA4CRIRCAQQAhoDY29wIAAo
ADAAOAkSEQgEEAMaA2RldCAAKAAwADgJEhMIBBAFGgVwdW5jdCAAKAAwADgJGgEE
KgEAmAMAsAMAugNECAIQCRgKIgREQVRFKghUSElTIFAxRDIEREFURTobEghUSElT
IFAxRBoFVG9kYXkiBERBVEUqAnQzUAJYAmIFVG9kYXnCA9kBCAQSBlBST1BFUhoI
U0lOR1VMQVIiB05FVVRSQUwqCUlOQU5JTUFURTIHVU5LTk9XTjgASAFQAFoFdG9k
YXliBERBVEVo////////////AXD///////////8BeASAAQSIAQKQAQCYAQGgAQCo
AQCwAQC4AQDAAQDIAQDQAQDYAQHgAQHoAQDyAQYIARAAIAD6ARYI////////////
ARD///////////8BggIECAEQAJIDBAgBEACSAwQIARABkgMECAEQApIDBAgBEAOS
AwQIARAEmgMECAEQAMID4AEIBRIHTk9NSU5BTBoIU0lOR1VMQVIiB05FVVRSQUwq
CUlOQU5JTUFURTIHVU5LTk9XTjgCSARQA1oEZ2lmdGIBT2j///////////8BcP//
/////////wF4BYABBYgBApABAJgBAaABAKgBALABALgBAMABAMgBANABANgBAeAB
AegBAPIBBggBEAMgAPoBFgj///////////8BEP///////////wGCAgQIARADkgME
CAEQAJIDBAgBEAGSAwQIARACkgMECAEQA5IDBAgBEASaAwQIARACmgMECAEQA7gD
BMgDAYgEAaAEARLmFwpbCgRUaGF0EgJJThoEVGhhdCoBIDIAOgRUaGF0QgFPUgR0
aGF0WD5gQmgAcgRQRVIwiAEOkAEPqAEAsAIA8gMBT/oDAU+SBBRPPTAuOTk5ODg3
OTg0MzU4OTA2OApZCgTigJlzEgNQT1MaBOKAmXMqADIBIDoE4oCZc0IBT1ICJ3NY
QmBEaAByBFBFUjCIAQ+QARCoAQCwAgDyAwFP+gMBT5IEE089MC45OTk4NjcyMTg2
MDkyMzkKWQoDd2h5EgNXUkIaA3doeSoBIDIBIDoDd2h5QgFPUgN3aHlYRWBIaABy
BFBFUjCIARCQARGoAQCwAgDyAwFP+gMBT5IEFE89MC45OTk5OTk1NDI4Mzg4OTc4
ClcKAndlEgNQUlAaAndlKgEgMgEgOgJ3ZUIBT1ICd2VYSWBLaAByBFBFUjCIARGQ
ARKoAQCwAgDyAwFP+gMBT4AEBpIEE089MC45OTk5OTc2NTM0MTk5MDcKXAoEY2Fs
bBIDVkJQGgRjYWxsKgEgMgEgOgRjYWxsQgFPUgRjYWxsWExgUGgAcgRQRVIwiAES
kAETqAEAsAIA8gMBT/oDAU+SBBNPPTAuOTk5OTk5MjM0NjAwNDMyClgKAml0EgNQ
UlAaAml0KgEgMgEgOgJpdEIBT1ICaXRYUWBTaAByBFBFUjCIAROQARSoAQCwAgDy
AwFP+gMBT4AEB5IEFE89MC45OTk5OTk5MDE3MzkxNDkzCk8KAScSAmBgGgEnKgEg
MgA6ASdCAU9SASdYVGBVaAByBFBFUjCIARSQARWoAQCwAgDyAwFP+gMBT5IEFE89
MC45OTk5OTk4Mjc3NDE0MzQxCloKA1RoZRICRFQaA1RoZSoAMgEgOgNUaGVCAU9S
A3RoZVhVYFhoAHIEUEVSMIgBFZABFqgBALACAPIDAU/6AwFPgAQIkgQUTz0wLjk5
OTk3NjgzNTg2MDQ4NTgKnAEKB1ByZXNlbnQSA05OUBoHUHJlc2VudCoBIDIAOgdQ
cmVzZW50QgREQVRFSgtQUkVTRU5UX1JFRlIHUHJlc2VudFhZYGBoAHIEUEVSMIgB
FpABF5oBIAoLUFJFU0VOVF9SRUYaB1ByZXNlbnQiBERBVEUqAnQ0qAEAsAIA8gME
REFURfoDBERBVEWABAiIBAOSBAlEQVRFPS0xLjAKTgoBJxICJycaAScqADIAOgEn
QgFPUgEnWGBgYWgAcgRQRVIwiAEXkAEYqAEAsAIA8gMBT/oDAU+SBBRPPTAuOTk5
NzA2NDAyNDM0NzU1MxAOGBggAyg+MGFC+gEKBAgDEAIKBAgDEAEKBAgDEAUKBAgD
EAMKBAgDEAQKBAgDEAYKBAgDEAcKBAgDEAkKBAgDEAgKBAgDEAoSEwgCEAEaBW5z
dWJqIAAoADAAOAkSEwgCEAUaBWNjb21wIAAoADAAOAkSFAgFEAMaBmFkdm1vZCAA
KAAwADgJEhMIBRAEGgVuc3ViaiAAKAAwADgJEhEIBRAGGgNvYmogACgAMAA4CRIT
CAUQBxoFcHVuY3QgACgAMAA4CRIRCAUQCRoDb2JqIAAoADAAOAkSEwgFEAoaBXB1
bmN0IAAoADAAOAkSEQgJEAgaA2RldCAAKAAwADgJGgECKgEASvoBCgQIAxACCgQI
AxABCgQIAxAFCgQIAxADCgQIAxAECgQIAxAGCgQIAxAHCgQIAxAJCgQIAxAICgQI
AxAKEhMIAhABGgVuc3ViaiAAKAAwADgJEhMIAhAFGgVjY29tcCAAKAAwADgJEhQI
BRADGgZhZHZtb2QgACgAMAA4CRITCAUQBBoFbnN1YmogACgAMAA4CRIRCAUQBhoD
b2JqIAAoADAAOAkSEwgFEAcaBXB1bmN0IAAoADAAOAkSEQgFEAkaA29iaiAAKAAw
ADgJEhMIBRAKGgVwdW5jdCAAKAAwADgJEhEICRAIGgNkZXQgACgAMAA4CRoBAioB
AFL6AQoECAMQAgoECAMQAQoECAMQBQoECAMQAwoECAMQBAoECAMQBgoECAMQBwoE
CAMQCQoECAMQCAoECAMQChITCAIQARoFbnN1YmogACgAMAA4CRITCAIQBRoFY2Nv
bXAgACgAMAA4CRIUCAUQAxoGYWR2bW9kIAAoADAAOAkSEwgFEAQaBW5zdWJqIAAo
ADAAOAkSEQgFEAYaA29iaiAAKAAwADgJEhMIBRAHGgVwdW5jdCAAKAAwADgJEhEI
BRAJGgNvYmogACgAMAA4CRITCAUQChoFcHVuY3QgACgAMAA4CRIRCAkQCBoDZGV0
IAAoADAAOAkaAQIqAQBYAYoB+gEKBAgDEAIKBAgDEAEKBAgDEAUKBAgDEAMKBAgD
EAQKBAgDEAYKBAgDEAcKBAgDEAkKBAgDEAgKBAgDEAoSEwgCEAEaBW5zdWJqIAAo
ADAAOAkSEwgCEAUaBWNjb21wIAAoADAAOAkSFAgFEAMaBmFkdm1vZCAAKAAwADgJ
EhMIBRAEGgVuc3ViaiAAKAAwADgJEhEIBRAGGgNvYmogACgAMAA4CRITCAUQBxoF
cHVuY3QgACgAMAA4CRIRCAUQCRoDb2JqIAAoADAAOAkSEwgFEAoaBXB1bmN0IAAo
ADAAOAkSEQgJEAgaA2RldCAAKAAwADgJGgECKgEAkgH6AQoECAMQAgoECAMQAQoE
CAMQBQoECAMQAwoECAMQBAoECAMQBgoECAMQBwoECAMQCQoECAMQCAoECAMQChIT
CAIQARoFbnN1YmogACgAMAA4CRITCAIQBRoFY2NvbXAgACgAMAA4CRIUCAUQAxoG
YWR2bW9kIAAoADAAOAkSEwgFEAQaBW5zdWJqIAAoADAAOAkSEQgFEAYaA29iaiAA
KAAwADgJEhMIBRAHGgVwdW5jdCAAKAAwADgJEhEIBRAJGgNvYmogACgAMAA4CRIT
CAUQChoFcHVuY3QgACgAMAA4CRIRCAkQCBoDZGV0IAAoADAAOAkaAQIqAQCYAwCw
AwC6A04IAxAWGBciBERBVEUqC1BSRVNFTlRfUkVGMgREQVRFOiAKC1BSRVNFTlRf
UkVGGgdQcmVzZW50IgREQVRFKgJ0NFADWANiB1ByZXNlbnTCA+EBCAYSClBST05P
TUlOQUwaBlBMVVJBTCIHVU5LTk9XTioHQU5JTUFURTICV0U4A0gEUANaAndlYgFP
aP///////////wFw////////////AXgGgAEGiAEDkAEAmAEBoAEBqAEAsAEAuAEA
wAEAyAEA0AEA2AEB4AEB6AEA8gEGCAIQAyAA+gEGCAIQBCAAggIECAIQA5IDBAgC
EACSAwQIAhABkgMECAIQApIDBAgCEAOSAwQIAhAEkgMECAIQBZIDBAgCEAaSAwQI
AhAHkgMECAIQCJIDBAgCEAmaAwQIAhADwgPlAQgHEgpQUk9OT01JTkFMGghTSU5H
VUxBUiIHTkVVVFJBTCoJSU5BTklNQVRFMgJJVDgFSAZQBVoCaXRiAU9o////////
////AXD///////////8BeAeAAQeIAQOQAQCYAQGgAQCoAQGwAQC4AQDAAQDIAQDQ
AQDYAQHgAQHoAQDyAQYIAhAFIAD6AQYIAhAEIACCAgQIAhAFkgMECAIQAJIDBAgC
EAGSAwQIAhACkgMECAIQA5IDBAgCEASSAwQIAhAFkgMECAIQBpIDBAgCEAeSAwQI
AhAIkgMECAIQCZoDBAgCEAXCA/UBCAgSBlBST1BFUhoIU0lOR1VMQVIiB1VOS05P
V04qCUlOQU5JTUFURTIHVU5LTk9XTjgHSAlQCFoHcHJlc2VudGIEREFURWj/////
//////8BcP///////////wF4CIABCIgBA5ABAJgBAaABAKgBAbABALgBAMABAMgB
ANABANgBAeABAegBAPIBBggCEAggAPoBBggCEAQgAIICBAgCEAiSAwQIAhAAkgME
CAIQAZIDBAgCEAKSAwQIAhADkgMECAIQBJIDBAgCEAWSAwQIAhAGkgMECAIQB5ID
BAgCEAiSAwQIAhAJmgMECAIQB5oDBAgCEAjIAwGIBAGgBAFKUggAEAAYASIEREFU
RSoLT0ZGU0VUIFAtMUQyBERBVEU6IhILT0ZGU0VUIFAtMUQaCVllc3RlcmRheSIE
REFURSoCdDFQAFgAYglZZXN0ZXJkYXlKTggBEAQYBSIEREFURSoKT0ZGU0VUIFAx
RDIEREFURTogEgpPRkZTRVQgUDFEGghUb21vcnJvdyIEREFURSoCdDJQAVgBYghU
b21vcnJvd0pECAIQCRgKIgREQVRFKghUSElTIFAxRDIEREFURTobEghUSElTIFAx
RBoFVG9kYXkiBERBVEUqAnQzUAJYAmIFVG9kYXlKTggDEBYYFyIEREFURSoLUFJF
U0VOVF9SRUYyBERBVEU6IAoLUFJFU0VOVF9SRUYaB1ByZXNlbnQiBERBVEUqAnQ0
UANYA2IHUHJlc2VudFgAaAFylQIIABIGUFJPUEVSGghTSU5HVUxBUiIHVU5LTk9X
TioJSU5BTklNQVRFMgdVTktOT1dOOABIAVAAWgl5ZXN0ZXJkYXliBERBVEVo////
////////AXD///////////8BeACAAQCIAQCQAQCYAQGgAQCoAQCwAQC4AQDAAQDI
AQDQAQDYAQHgAQHoAQDyAQ8I////////////ARAAIAD6ARYI////////////ARD/
//////////8BggINCP///////////wEQAJIDDQj///////////8BEACSAw0I////
////////ARABkgMNCP///////////wEQApIDDQj///////////8BEAOaAw0I////
////////ARAAcpQCCAESB05PTUlOQUwaCFNJTkdVTEFSIgdORVVUUkFMKglJTkFO
SU1BVEUyB1VOS05PV044AkgDUAJaB2hpc3RvcnliAU9o////////////AXD/////
//////8BeAGAAQGIAQCQAQCYAQGgAQCoAQCwAQC4AQDAAQDIAQDQAQDYAQHgAQHo
AQDyAQ8I////////////ARACIAD6ARYI////////////ARD///////////8BggIN
CP///////////wEQApIDDQj///////////8BEACSAw0I////////////ARABkgMN
CP///////////wEQApIDDQj///////////8BEAOaAw0I////////////ARACuAMA
ctwBCAISBlBST1BFUhoIU0lOR1VMQVIiB1VOS05PV04qCUlOQU5JTUFURTIHVU5L
Tk9XTjgASAFQAFoIdG9tb3Jyb3diBERBVEVo////////////AXD///////////8B
eAKAAQKIAQGQAQCYAQGgAQCoAQCwAQC4AQDAAQDIAQDQAQDYAQHgAQHoAQDyAQYI
ABAAIAD6ARYI////////////ARD///////////8BggIECAAQAJIDBAgAEACSAwQI
ABABkgMECAAQApIDBAgAEAOSAwQIABAEmgMECAAQAHLjAQgDEgdOT01JTkFMGghT
SU5HVUxBUiIHTkVVVFJBTCoJSU5BTklNQVRFMgdVTktOT1dOOAJIBFADWgdteXN0
ZXJ5YgFPaP///////////wFw////////////AXgDgAEDiAEBkAEAmAEBoAEAqAEA
sAEAuAEAwAEAyAEA0AEA2AEB4AEB6AEA8gEGCAAQAyAA+gEWCP///////////wEQ
////////////AYICBAgAEAOSAwQIABAAkgMECAAQAZIDBAgAEAKSAwQIABADkgME
CAAQBJoDBAgAEAKaAwQIABADuAMCctkBCAQSBlBST1BFUhoIU0lOR1VMQVIiB05F
VVRSQUwqCUlOQU5JTUFURTIHVU5LTk9XTjgASAFQAFoFdG9kYXliBERBVEVo////
////////AXD///////////8BeASAAQSIAQKQAQCYAQGgAQCoAQCwAQC4AQDAAQDI
AQDQAQDYAQHgAQHoAQDyAQYIARAAIAD6ARYI////////////ARD///////////8B
ggIECAEQAJIDBAgBEACSAwQIARABkgMECAEQApIDBAgBEAOSAwQIARAEmgMECAEQ
AHLgAQgFEgdOT01JTkFMGghTSU5HVUxBUiIHTkVVVFJBTCoJSU5BTklNQVRFMgdV
TktOT1dOOAJIBFADWgRnaWZ0YgFPaP///////////wFw////////////AXgFgAEF
iAECkAEAmAEBoAEAqAEAsAEAuAEAwAEAyAEA0AEA2AEB4AEB6AEA8gEGCAEQAyAA
+gEWCP///////////wEQ////////////AYICBAgBEAOSAwQIARAAkgMECAEQAZID
BAgBEAKSAwQIARADkgMECAEQBJoDBAgBEAKaAwQIARADuAMEcuEBCAYSClBST05P
TUlOQUwaBlBMVVJBTCIHVU5LTk9XTioHQU5JTUFURTICV0U4A0gEUANaAndlYgFP
aP///////////wFw////////////AXgGgAEGiAEDkAEAmAEBoAEBqAEAsAEAuAEA
wAEAyAEA0AEA2AEB4AEB6AEA8gEGCAIQAyAA+gEGCAIQBCAAggIECAIQA5IDBAgC
EACSAwQIAhABkgMECAIQApIDBAgCEAOSAwQIAhAEkgMECAIQBZIDBAgCEAaSAwQI
AhAHkgMECAIQCJIDBAgCEAmaAwQIAhADcuUBCAcSClBST05PTUlOQUwaCFNJTkdV
TEFSIgdORVVUUkFMKglJTkFOSU1BVEUyAklUOAVIBlAFWgJpdGIBT2j/////////
//8BcP///////////wF4B4ABB4gBA5ABAJgBAaABAKgBAbABALgBAMABAMgBANAB
ANgBAeABAegBAPIBBggCEAUgAPoBBggCEAQgAIICBAgCEAWSAwQIAhAAkgMECAIQ
AZIDBAgCEAKSAwQIAhADkgMECAIQBJIDBAgCEAWSAwQIAhAGkgMECAIQB5IDBAgC
EAiSAwQIAhAJmgMECAIQBXL1AQgIEgZQUk9QRVIaCFNJTkdVTEFSIgdVTktOT1dO
KglJTkFOSU1BVEUyB1VOS05PV044B0gJUAhaB3ByZXNlbnRiBERBVEVo////////
////AXD///////////8BeAiAAQiIAQOQAQCYAQGgAQCoAQGwAQC4AQDAAQDIAQDQ
AQDYAQHgAQHoAQDyAQYIAhAIIAD6AQYIAhAEIACCAgQIAhAIkgMECAIQAJIDBAgC
EAGSAwQIAhACkgMECAIQA5IDBAgCEASSAwQIAhAFkgMECAIQBpIDBAgCEAeSAwQI
AhAIkgMECAIQCZoDBAgCEAeaAwQIAhAIeAGAAQGIAQCIAf///////////wGIAQGI
Af///////////wGIAQKIAf///////////wGIAf///////////wGIAf//////////
/wGIAf///////////wGQAQCQAQKQAQSQAf///////////wE=
`

	// RosesShortResp is the standard base64 (as defined in RFC 4648)
	// encoded response of annotating
	//
	//	Roses are red.
	//	  Violets are blue.
	//	Sugar is sweet.
	//	  And so are you.
	//
	// with annotators "tokenize,ssplit,pos" by Stanford CoreNLP 4.5.6.
	RosesShortResp = `
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

	// YesterdayShortResp is the standard base64 (as defined in RFC 4648)
	// encoded response of annotating
	//
	//	Yesterday is history. Tomorrow is a mystery. Today is a gift. That’s why we call it 'The Present'
	//
	// with annotators "tokenize,ssplit,pos" by Stanford CoreNLP 4.5.6.
	YesterdayShortResp = `
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
)

// MakeDifferentResponsesData generates bytes consisting of the decoding
// results of RosesResp, YesterdayResp, RosesShortResp, and YesterdayShortResp,
// repeating the specified number of times.
//
// It also returns the lengths of the decoding results of RosesResp,
// YesterdayResp, RosesShortResp, and YesterdayShortResp, of type [4]int.
//
// If the returned error is non-nil,
// it is caused by base64.StdEncoding.DecodeString.
func MakeDifferentResponsesData(numRepeat int) (
	data []byte, lens [4]int, err error) {
	rosesRespBody, err := base64.StdEncoding.DecodeString(RosesResp)
	lens[0] = len(rosesRespBody)
	if err != nil {
		err = gogoerrors.AutoWrap(err)
		return
	}
	yesterdayRespBody, err := base64.StdEncoding.DecodeString(YesterdayResp)
	lens[1] = len(yesterdayRespBody)
	if err != nil {
		err = gogoerrors.AutoWrap(err)
		return
	}
	rosesShortRespBody, err := base64.StdEncoding.DecodeString(RosesShortResp)
	lens[2] = len(rosesShortRespBody)
	if err != nil {
		err = gogoerrors.AutoWrap(err)
		return
	}
	yesterdayShortRespBody, err := base64.StdEncoding.DecodeString(
		YesterdayShortResp)
	lens[3] = len(yesterdayShortRespBody)
	if err != nil {
		err = gogoerrors.AutoWrap(err)
		return
	}
	data = make([]byte, (lens[0]+lens[1]+lens[2]+lens[3])*numRepeat)
	var n int
	for i := 0; i < numRepeat; i++ {
		n += copy(data[n:], rosesRespBody)
		n += copy(data[n:], yesterdayRespBody)
		n += copy(data[n:], rosesShortRespBody)
		n += copy(data[n:], yesterdayShortRespBody)
	}
	return
}
