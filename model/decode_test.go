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
	"github.com/donyori/gocorenlp/model/v4.5.5-f1b929e47a57/pb"
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
	// with the server default annotators by Stanford CoreNLP 4.5.5.
	RosesResp = pbtest.RosesAreRedRespV453

	// YesterdayResp is the standard base64 (as defined in RFC 4648)
	// encoded response of annotating
	//
	//	Yesterday is history. Tomorrow is a mystery. Today is a gift. That’s why we call it 'The Present'
	//
	// with the server default annotators by Stanford CoreNLP 4.5.5.
	YesterdayResp = `
1VIKY1llc3RlcmRheSBpcyBoaXN0b3J5LiBUb21vcnJvdyBpcyBhIG15c3Rlcnku
IFRvZGF5IGlzIGEgZ2lmdC4gVGhhdOKAmXMgd2h5IHdlIGNhbGwgaXQgJ1RoZSBQ
cmVzZW50JxKgDAqlAQoJWWVzdGVyZGF5EgJOThoJWWVzdGVyZGF5KgAyASA6CVll
c3RlcmRheUIEREFURUoLT0ZGU0VUIFAtMURSCXllc3RlcmRheVgAYAloAHIEUEVS
MIgBAJABAZoBIhILT0ZGU0VUIFAtMUQaCVllc3RlcmRheSIEREFURSoCdDGoAQCw
AgDyAwREQVRF+gMEREFURYAEAIgEAJIECURBVEU9LTEuMApVCgJpcxIDVkJaGgJp
cyoBIDIBIDoCaXNCAU9SAmJlWApgDGgAcgRQRVIwiAEBkAECqAEAsAIA8gMBT/oD
AU+SBBRPPTAuOTk5OTk5Njg0MzU4ODMzNgpqCgdoaXN0b3J5EgJOThoHaGlzdG9y
eSoBIDIAOgdoaXN0b3J5QgFPUgdoaXN0b3J5WA1gFGgAcgRQRVIwiAECkAEDqAEA
sAIA8gMBT/oDAU+ABAGSBBRPPTAuOTk5OTUzOTc0MjA0ODk4NgpOCgEuEgEuGgEu
KgAyASA6AS5CAU9SAS5YFGAVaAByBFBFUjCIAQOQAQSoAQCwAgDyAwFP+gMBT5IE
FE89MC45OTk5OTk1NTk0NzY5NzMzEAAYBCAAKAAwFUJYCgQIABADCgQIABABCgQI
ABACCgQIABAEEhMIAxABGgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4
CRITCAMQBBoFcHVuY3QgACgAMAA4CRoBA0pYCgQIABADCgQIABABCgQIABACCgQI
ABAEEhMIAxABGgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRITCAMQ
BBoFcHVuY3QgACgAMAA4CRoBA1JYCgQIABADCgQIABABCgQIABACCgQIABAEEhMI
AxABGgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRITCAMQBBoFcHVu
Y3QgACgAMAA4CRoBA1gBigFYCgQIABADCgQIABABCgQIABACCgQIABAEEhMIAxAB
GgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRITCAMQBBoFcHVuY3Qg
ACgAMAA4CRoBA5IBWAoECAAQAwoECAAQAQoECAAQAgoECAAQBBITCAMQARoFbnN1
YmogACgAMAA4CRIRCAMQAhoDY29wIAAoADAAOAkSEwgDEAQaBXB1bmN0IAAoADAA
OAkaAQOYAwCwAwC6A1IIABAAGAEiBERBVEUqC09GRlNFVCBQLTFEMgREQVRFOiIS
C09GRlNFVCBQLTFEGglZZXN0ZXJkYXkiBERBVEUqAnQxUABYAGIJWWVzdGVyZGF5
wgOVAggAEgZQUk9QRVIaCFNJTkdVTEFSIgdVTktOT1dOKglJTkFOSU1BVEUyB1VO
S05PV044AEgBUABaCXllc3RlcmRheWIEREFURWj///////////8BcP//////////
/wF4AIABAIgBAJABAJgBAaABAKgBALABALgBAMABAMgBANABANgBAeABAegBAPIB
Dwj///////////8BEAAgAPoBFgj///////////8BEP///////////wGCAg0I////
////////ARAAkgMNCP///////////wEQAJIDDQj///////////8BEAGSAw0I////
////////ARACkgMNCP///////////wEQA5oDDQj///////////8BEADCA5QCCAES
B05PTUlOQUwaCFNJTkdVTEFSIgdORVVUUkFMKglJTkFOSU1BVEUyB1VOS05PV044
AkgDUAJaB2hpc3RvcnliAU9o////////////AXD///////////8BeAGAAQGIAQCQ
AQCYAQGgAQCoAQCwAQC4AQDAAQDIAQDQAQDYAQHgAQHoAQDyAQ8I////////////
ARACIAD6ARYI////////////ARD///////////8BggINCP///////////wEQApID
DQj///////////8BEACSAw0I////////////ARABkgMNCP///////////wEQApID
DQj///////////8BEAOaAw0I////////////ARACuAMAyAMBiAQBoAQBEv4MCp8B
CghUb21vcnJvdxICTk4aCFRvbW9ycm93KgEgMgEgOghUb21vcnJvd0IEREFURUoK
T0ZGU0VUIFAxRFIIdG9tb3Jyb3dYFmAeaAByBFBFUjCIAQSQAQWaASASCk9GRlNF
VCBQMUQaCFRvbW9ycm93IgREQVRFKgJ0MqgBALACAPIDBERBVEX6AwREQVRFgAQC
iAQBkgQJREFURT0tMS4wClUKAmlzEgNWQloaAmlzKgEgMgEgOgJpc0IBT1ICYmVY
H2AhaAByBFBFUjCIAQWQAQaoAQCwAgDyAwFP+gMBT5IEFE89MC45OTk5OTk1MTcx
MjgwMzQ3ClMKAWESAkRUGgFhKgEgMgEgOgFhQgFPUgFhWCJgI2gAcgRQRVIwiAEG
kAEHqAEAsAIA8gMBT/oDAU+ABAOSBBRPPTAuOTk5OTk5MzQ0MDIyMTAyNApqCgdt
eXN0ZXJ5EgJOThoHbXlzdGVyeSoBIDIAOgdteXN0ZXJ5QgFPUgdteXN0ZXJ5WCRg
K2gAcgRQRVIwiAEHkAEIqAEAsAIA8gMBT/oDAU+ABAOSBBRPPTAuOTk5OTc3NTIz
OTU4Mzk4OQpOCgEuEgEuGgEuKgAyASA6AS5CAU9SAS5YK2AsaAByBFBFUjCIAQiQ
AQmoAQCwAgDyAwFP+gMBT5IEFE89MC45OTk5OTk5NDc4MTA3MzY5EAQYCSABKBYw
LEJxCgQIARAECgQIARABCgQIARACCgQIARADCgQIARAFEhMIBBABGgVuc3ViaiAA
KAAwADgJEhEIBBACGgNjb3AgACgAMAA4CRIRCAQQAxoDZGV0IAAoADAAOAkSEwgE
EAUaBXB1bmN0IAAoADAAOAkaAQRKcQoECAEQBAoECAEQAQoECAEQAgoECAEQAwoE
CAEQBRITCAQQARoFbnN1YmogACgAMAA4CRIRCAQQAhoDY29wIAAoADAAOAkSEQgE
EAMaA2RldCAAKAAwADgJEhMIBBAFGgVwdW5jdCAAKAAwADgJGgEEUnEKBAgBEAQK
BAgBEAEKBAgBEAIKBAgBEAMKBAgBEAUSEwgEEAEaBW5zdWJqIAAoADAAOAkSEQgE
EAIaA2NvcCAAKAAwADgJEhEIBBADGgNkZXQgACgAMAA4CRITCAQQBRoFcHVuY3Qg
ACgAMAA4CRoBBFgBigFxCgQIARAECgQIARABCgQIARACCgQIARADCgQIARAFEhMI
BBABGgVuc3ViaiAAKAAwADgJEhEIBBACGgNjb3AgACgAMAA4CRIRCAQQAxoDZGV0
IAAoADAAOAkSEwgEEAUaBXB1bmN0IAAoADAAOAkaAQSSAXEKBAgBEAQKBAgBEAEK
BAgBEAIKBAgBEAMKBAgBEAUSEwgEEAEaBW5zdWJqIAAoADAAOAkSEQgEEAIaA2Nv
cCAAKAAwADgJEhEIBBADGgNkZXQgACgAMAA4CRITCAQQBRoFcHVuY3QgACgAMAA4
CRoBBJgDALADALoDTggBEAQYBSIEREFURSoKT0ZGU0VUIFAxRDIEREFURTogEgpP
RkZTRVQgUDFEGghUb21vcnJvdyIEREFURSoCdDJQAVgBYghUb21vcnJvd8ID3AEI
AhIGUFJPUEVSGghTSU5HVUxBUiIHVU5LTk9XTioJSU5BTklNQVRFMgdVTktOT1dO
OABIAVAAWgh0b21vcnJvd2IEREFURWj///////////8BcP///////////wF4AoAB
AogBAZABAJgBAaABAKgBALABALgBAMABAMgBANABANgBAeABAegBAPIBBggAEAAg
APoBFgj///////////8BEP///////////wGCAgQIABAAkgMECAAQAJIDBAgAEAGS
AwQIABACkgMECAAQA5IDBAgAEASaAwQIABAAwgPjAQgDEgdOT01JTkFMGghTSU5H
VUxBUiIHTkVVVFJBTCoJSU5BTklNQVRFMgdVTktOT1dOOAJIBFADWgdteXN0ZXJ5
YgFPaP///////////wFw////////////AXgDgAEDiAEBkAEAmAEBoAEAqAEAsAEA
uAEAwAEAyAEA0AEA2AEB4AEB6AEA8gEGCAAQAyAA+gEWCP///////////wEQ////
////////AYICBAgAEAOSAwQIABAAkgMECAAQAZIDBAgAEAKSAwQIABADkgMECAAQ
BJoDBAgAEAKaAwQIABADuAMCyAMBiAQBoAQBEs8MCowBCgVUb2RheRICTk4aBVRv
ZGF5KgEgMgEgOgVUb2RheUIEREFURUoIVEhJUyBQMURSBXRvZGF5WC1gMmgAcgRQ
RVIwiAEJkAEKmgEbEghUSElTIFAxRBoFVG9kYXkiBERBVEUqAnQzqAEAsAIA8gME
REFURfoDBERBVEWABASIBAKSBAlEQVRFPS0xLjAKVQoCaXMSA1ZCWhoCaXMqASAy
ASA6AmlzQgFPUgJiZVgzYDVoAHIEUEVSMIgBCpABC6gBALACAPIDAU/6AwFPkgQU
Tz0wLjk5OTk5OTMyOTAxNTA2NjIKUwoBYRICRFQaAWEqASAyASA6AWFCAU9SAWFY
NmA3aAByBFBFUjCIAQuQAQyoAQCwAgDyAwFP+gMBT4AEBZIEFE89MC45OTk5OTgz
NzQyNDI3OTI4Cl4KBGdpZnQSAk5OGgRnaWZ0KgEgMgA6BGdpZnRCAU9SBGdpZnRY
OGA8aAByBFBFUjCIAQyQAQ2oAQCwAgDyAwFP+gMBT4AEBZIEFE89MC45OTk5ODIz
MTQ2NzI2NDY2Ck4KAS4SAS4aAS4qADIBIDoBLkIBT1IBLlg8YD1oAHIEUEVSMIgB
DZABDqgBALACAPIDAU/6AwFPkgQUTz0wLjk5OTk5OTg4MzI3MDQ0MDMQCRgOIAIo
LTA9QnEKBAgCEAQKBAgCEAEKBAgCEAIKBAgCEAMKBAgCEAUSEwgEEAEaBW5zdWJq
IAAoADAAOAkSEQgEEAIaA2NvcCAAKAAwADgJEhEIBBADGgNkZXQgACgAMAA4CRIT
CAQQBRoFcHVuY3QgACgAMAA4CRoBBEpxCgQIAhAECgQIAhABCgQIAhACCgQIAhAD
CgQIAhAFEhMIBBABGgVuc3ViaiAAKAAwADgJEhEIBBACGgNjb3AgACgAMAA4CRIR
CAQQAxoDZGV0IAAoADAAOAkSEwgEEAUaBXB1bmN0IAAoADAAOAkaAQRScQoECAIQ
BAoECAIQAQoECAIQAgoECAIQAwoECAIQBRITCAQQARoFbnN1YmogACgAMAA4CRIR
CAQQAhoDY29wIAAoADAAOAkSEQgEEAMaA2RldCAAKAAwADgJEhMIBBAFGgVwdW5j
dCAAKAAwADgJGgEEWAGKAXEKBAgCEAQKBAgCEAEKBAgCEAIKBAgCEAMKBAgCEAUS
EwgEEAEaBW5zdWJqIAAoADAAOAkSEQgEEAIaA2NvcCAAKAAwADgJEhEIBBADGgNk
ZXQgACgAMAA4CRITCAQQBRoFcHVuY3QgACgAMAA4CRoBBJIBcQoECAIQBAoECAIQ
AQoECAIQAgoECAIQAwoECAIQBRITCAQQARoFbnN1YmogACgAMAA4CRIRCAQQAhoD
Y29wIAAoADAAOAkSEQgEEAMaA2RldCAAKAAwADgJEhMIBBAFGgVwdW5jdCAAKAAw
ADgJGgEEmAMAsAMAugNECAIQCRgKIgREQVRFKghUSElTIFAxRDIEREFURTobEghU
SElTIFAxRBoFVG9kYXkiBERBVEUqAnQzUAJYAmIFVG9kYXnCA9kBCAQSBlBST1BF
UhoIU0lOR1VMQVIiB05FVVRSQUwqCUlOQU5JTUFURTIHVU5LTk9XTjgASAFQAFoF
dG9kYXliBERBVEVo////////////AXD///////////8BeASAAQSIAQKQAQCYAQGg
AQCoAQCwAQC4AQDAAQDIAQDQAQDYAQHgAQHoAQDyAQYIARAAIAD6ARYI////////
////ARD///////////8BggIECAEQAJIDBAgBEACSAwQIARABkgMECAEQApIDBAgB
EAOSAwQIARAEmgMECAEQAMID4AEIBRIHTk9NSU5BTBoIU0lOR1VMQVIiB05FVVRS
QUwqCUlOQU5JTUFURTIHVU5LTk9XTjgCSARQA1oEZ2lmdGIBT2j///////////8B
cP///////////wF4BYABBYgBApABAJgBAaABAKgBALABALgBAMABAMgBANABANgB
AeABAegBAPIBBggBEAMgAPoBFgj///////////8BEP///////////wGCAgQIARAD
kgMECAEQAJIDBAgBEAGSAwQIARACkgMECAEQA5IDBAgBEASaAwQIARACmgMECAEQ
A7gDBMgDAYgEAaAEARLXFwpbCgRUaGF0EgJJThoEVGhhdCoBIDIAOgRUaGF0QgFP
UgR0aGF0WD5gQmgAcgRQRVIwiAEOkAEPqAEAsAIA8gMBT/oDAU+SBBRPPTAuOTk5
ODg3OTg0MzU4OTA2OApZCgTigJlzEgNQT1MaBOKAmXMqADIBIDoE4oCZc0IBT1IC
J3NYQmBEaAByBFBFUjCIAQ+QARCoAQCwAgDyAwFP+gMBT5IEE089MC45OTk4Njcy
MTg2MDkyMzkKWQoDd2h5EgNXUkIaA3doeSoBIDIBIDoDd2h5QgFPUgN3aHlYRWBI
aAByBFBFUjCIARCQARGoAQCwAgDyAwFP+gMBT5IEFE89MC45OTk5OTk1NDI4Mzg4
OTc4ClcKAndlEgNQUlAaAndlKgEgMgEgOgJ3ZUIBT1ICd2VYSWBLaAByBFBFUjCI
ARGQARKoAQCwAgDyAwFP+gMBT4AEBpIEE089MC45OTk5OTc2NTM0MTk5MDcKXAoE
Y2FsbBIDVkJQGgRjYWxsKgEgMgEgOgRjYWxsQgFPUgRjYWxsWExgUGgAcgRQRVIw
iAESkAETqAEAsAIA8gMBT/oDAU+SBBNPPTAuOTk5OTk5MjM0NjAwNDMyClgKAml0
EgNQUlAaAml0KgEgMgEgOgJpdEIBT1ICaXRYUWBTaAByBFBFUjCIAROQARSoAQCw
AgDyAwFP+gMBT4AEB5IEFE89MC45OTk5OTk5MDE3MzkxNDkzCk8KAScSAmBgGgEn
KgEgMgA6ASdCAU9SASdYVGBVaAByBFBFUjCIARSQARWoAQCwAgDyAwFP+gMBT5IE
FE89MC45OTk5OTk4Mjc3NDE0MzQxCloKA1RoZRICRFQaA1RoZSoAMgEgOgNUaGVC
AU9SA3RoZVhVYFhoAHIEUEVSMIgBFZABFqgBALACAPIDAU/6AwFPgAQIkgQUTz0w
Ljk5OTk3NjgzNTg2MDQ4NTgKnAEKB1ByZXNlbnQSA05OUBoHUHJlc2VudCoBIDIA
OgdQcmVzZW50QgREQVRFSgtQUkVTRU5UX1JFRlIHUHJlc2VudFhZYGBoAHIEUEVS
MIgBFpABF5oBIAoLUFJFU0VOVF9SRUYaB1ByZXNlbnQiBERBVEUqAnQ0qAEAsAIA
8gMEREFURfoDBERBVEWABAiIBAOSBAlEQVRFPS0xLjAKTgoBJxICJycaAScqADIA
OgEnQgFPUgEnWGBgYWgAcgRQRVIwiAEXkAEYqAEAsAIA8gMBT/oDAU+SBBRPPTAu
OTk5NzA2NDAyNDM0NzU1MxAOGBggAyg+MGFC9wEKBAgDEAIKBAgDEAEKBAgDEAUK
BAgDEAMKBAgDEAQKBAgDEAYKBAgDEAcKBAgDEAkKBAgDEAgKBAgDEAoSEwgCEAEa
BW5zdWJqIAAoADAAOAkSEwgCEAUaBWNjb21wIAAoADAAOAkSFAgFEAMaBmFkdm1v
ZCAAKAAwADgJEhMIBRAEGgVuc3ViaiAAKAAwADgJEhEIBRAGGgNvYmogACgAMAA4
CRITCAUQBxoFcHVuY3QgACgAMAA4CRIRCAUQCRoDb2JqIAAoADAAOAkSEwgFEAoa
BXB1bmN0IAAoADAAOAkSEQgJEAgaA2RldCAAKAAwADgJGgECSvcBCgQIAxACCgQI
AxABCgQIAxAFCgQIAxADCgQIAxAECgQIAxAGCgQIAxAHCgQIAxAJCgQIAxAICgQI
AxAKEhMIAhABGgVuc3ViaiAAKAAwADgJEhMIAhAFGgVjY29tcCAAKAAwADgJEhQI
BRADGgZhZHZtb2QgACgAMAA4CRITCAUQBBoFbnN1YmogACgAMAA4CRIRCAUQBhoD
b2JqIAAoADAAOAkSEwgFEAcaBXB1bmN0IAAoADAAOAkSEQgFEAkaA29iaiAAKAAw
ADgJEhMIBRAKGgVwdW5jdCAAKAAwADgJEhEICRAIGgNkZXQgACgAMAA4CRoBAlL3
AQoECAMQAgoECAMQAQoECAMQBQoECAMQAwoECAMQBAoECAMQBgoECAMQBwoECAMQ
CQoECAMQCAoECAMQChITCAIQARoFbnN1YmogACgAMAA4CRITCAIQBRoFY2NvbXAg
ACgAMAA4CRIUCAUQAxoGYWR2bW9kIAAoADAAOAkSEwgFEAQaBW5zdWJqIAAoADAA
OAkSEQgFEAYaA29iaiAAKAAwADgJEhMIBRAHGgVwdW5jdCAAKAAwADgJEhEIBRAJ
GgNvYmogACgAMAA4CRITCAUQChoFcHVuY3QgACgAMAA4CRIRCAkQCBoDZGV0IAAo
ADAAOAkaAQJYAYoB9wEKBAgDEAIKBAgDEAEKBAgDEAUKBAgDEAMKBAgDEAQKBAgD
EAYKBAgDEAcKBAgDEAkKBAgDEAgKBAgDEAoSEwgCEAEaBW5zdWJqIAAoADAAOAkS
EwgCEAUaBWNjb21wIAAoADAAOAkSFAgFEAMaBmFkdm1vZCAAKAAwADgJEhMIBRAE
GgVuc3ViaiAAKAAwADgJEhEIBRAGGgNvYmogACgAMAA4CRITCAUQBxoFcHVuY3Qg
ACgAMAA4CRIRCAUQCRoDb2JqIAAoADAAOAkSEwgFEAoaBXB1bmN0IAAoADAAOAkS
EQgJEAgaA2RldCAAKAAwADgJGgECkgH3AQoECAMQAgoECAMQAQoECAMQBQoECAMQ
AwoECAMQBAoECAMQBgoECAMQBwoECAMQCQoECAMQCAoECAMQChITCAIQARoFbnN1
YmogACgAMAA4CRITCAIQBRoFY2NvbXAgACgAMAA4CRIUCAUQAxoGYWR2bW9kIAAo
ADAAOAkSEwgFEAQaBW5zdWJqIAAoADAAOAkSEQgFEAYaA29iaiAAKAAwADgJEhMI
BRAHGgVwdW5jdCAAKAAwADgJEhEIBRAJGgNvYmogACgAMAA4CRITCAUQChoFcHVu
Y3QgACgAMAA4CRIRCAkQCBoDZGV0IAAoADAAOAkaAQKYAwCwAwC6A04IAxAWGBci
BERBVEUqC1BSRVNFTlRfUkVGMgREQVRFOiAKC1BSRVNFTlRfUkVGGgdQcmVzZW50
IgREQVRFKgJ0NFADWANiB1ByZXNlbnTCA+EBCAYSClBST05PTUlOQUwaBlBMVVJB
TCIHVU5LTk9XTioHQU5JTUFURTICV0U4A0gEUANaAndlYgFPaP///////////wFw
////////////AXgGgAEGiAEDkAEAmAEBoAEBqAEAsAEAuAEAwAEAyAEA0AEA2AEB
4AEB6AEA8gEGCAIQAyAA+gEGCAIQBCAAggIECAIQA5IDBAgCEACSAwQIAhABkgME
CAIQApIDBAgCEAOSAwQIAhAEkgMECAIQBZIDBAgCEAaSAwQIAhAHkgMECAIQCJID
BAgCEAmaAwQIAhADwgPlAQgHEgpQUk9OT01JTkFMGghTSU5HVUxBUiIHTkVVVFJB
TCoJSU5BTklNQVRFMgJJVDgFSAZQBVoCaXRiAU9o////////////AXD/////////
//8BeAeAAQeIAQOQAQCYAQGgAQCoAQGwAQC4AQDAAQDIAQDQAQDYAQHgAQHoAQDy
AQYIAhAFIAD6AQYIAhAEIACCAgQIAhAFkgMECAIQAJIDBAgCEAGSAwQIAhACkgME
CAIQA5IDBAgCEASSAwQIAhAFkgMECAIQBpIDBAgCEAeSAwQIAhAIkgMECAIQCZoD
BAgCEAXCA/UBCAgSBlBST1BFUhoIU0lOR1VMQVIiB1VOS05PV04qCUlOQU5JTUFU
RTIHVU5LTk9XTjgHSAlQCFoHcHJlc2VudGIEREFURWj///////////8BcP//////
/////wF4CIABCIgBA5ABAJgBAaABAKgBAbABALgBAMABAMgBANABANgBAeABAegB
APIBBggCEAggAPoBBggCEAQgAIICBAgCEAiSAwQIAhAAkgMECAIQAZIDBAgCEAKS
AwQIAhADkgMECAIQBJIDBAgCEAWSAwQIAhAGkgMECAIQB5IDBAgCEAiSAwQIAhAJ
mgMECAIQB5oDBAgCEAjIAwGIBAGgBAFKUggAEAAYASIEREFURSoLT0ZGU0VUIFAt
MUQyBERBVEU6IhILT0ZGU0VUIFAtMUQaCVllc3RlcmRheSIEREFURSoCdDFQAFgA
YglZZXN0ZXJkYXlKTggBEAQYBSIEREFURSoKT0ZGU0VUIFAxRDIEREFURTogEgpP
RkZTRVQgUDFEGghUb21vcnJvdyIEREFURSoCdDJQAVgBYghUb21vcnJvd0pECAIQ
CRgKIgREQVRFKghUSElTIFAxRDIEREFURTobEghUSElTIFAxRBoFVG9kYXkiBERB
VEUqAnQzUAJYAmIFVG9kYXlKTggDEBYYFyIEREFURSoLUFJFU0VOVF9SRUYyBERB
VEU6IAoLUFJFU0VOVF9SRUYaB1ByZXNlbnQiBERBVEUqAnQ0UANYA2IHUHJlc2Vu
dFgAaAFylQIIABIGUFJPUEVSGghTSU5HVUxBUiIHVU5LTk9XTioJSU5BTklNQVRF
MgdVTktOT1dOOABIAVAAWgl5ZXN0ZXJkYXliBERBVEVo////////////AXD/////
//////8BeACAAQCIAQCQAQCYAQGgAQCoAQCwAQC4AQDAAQDIAQDQAQDYAQHgAQHo
AQDyAQ8I////////////ARAAIAD6ARYI////////////ARD///////////8BggIN
CP///////////wEQAJIDDQj///////////8BEACSAw0I////////////ARABkgMN
CP///////////wEQApIDDQj///////////8BEAOaAw0I////////////ARAAcpQC
CAESB05PTUlOQUwaCFNJTkdVTEFSIgdORVVUUkFMKglJTkFOSU1BVEUyB1VOS05P
V044AkgDUAJaB2hpc3RvcnliAU9o////////////AXD///////////8BeAGAAQGI
AQCQAQCYAQGgAQCoAQCwAQC4AQDAAQDIAQDQAQDYAQHgAQHoAQDyAQ8I////////
////ARACIAD6ARYI////////////ARD///////////8BggINCP///////////wEQ
ApIDDQj///////////8BEACSAw0I////////////ARABkgMNCP///////////wEQ
ApIDDQj///////////8BEAOaAw0I////////////ARACuAMActwBCAISBlBST1BF
UhoIU0lOR1VMQVIiB1VOS05PV04qCUlOQU5JTUFURTIHVU5LTk9XTjgASAFQAFoI
dG9tb3Jyb3diBERBVEVo////////////AXD///////////8BeAKAAQKIAQGQAQCY
AQGgAQCoAQCwAQC4AQDAAQDIAQDQAQDYAQHgAQHoAQDyAQYIABAAIAD6ARYI////
////////ARD///////////8BggIECAAQAJIDBAgAEACSAwQIABABkgMECAAQApID
BAgAEAOSAwQIABAEmgMECAAQAHLjAQgDEgdOT01JTkFMGghTSU5HVUxBUiIHTkVV
VFJBTCoJSU5BTklNQVRFMgdVTktOT1dOOAJIBFADWgdteXN0ZXJ5YgFPaP//////
/////wFw////////////AXgDgAEDiAEBkAEAmAEBoAEAqAEAsAEAuAEAwAEAyAEA
0AEA2AEB4AEB6AEA8gEGCAAQAyAA+gEWCP///////////wEQ////////////AYIC
BAgAEAOSAwQIABAAkgMECAAQAZIDBAgAEAKSAwQIABADkgMECAAQBJoDBAgAEAKa
AwQIABADuAMCctkBCAQSBlBST1BFUhoIU0lOR1VMQVIiB05FVVRSQUwqCUlOQU5J
TUFURTIHVU5LTk9XTjgASAFQAFoFdG9kYXliBERBVEVo////////////AXD/////
//////8BeASAAQSIAQKQAQCYAQGgAQCoAQCwAQC4AQDAAQDIAQDQAQDYAQHgAQHo
AQDyAQYIARAAIAD6ARYI////////////ARD///////////8BggIECAEQAJIDBAgB
EACSAwQIARABkgMECAEQApIDBAgBEAOSAwQIARAEmgMECAEQAHLgAQgFEgdOT01J
TkFMGghTSU5HVUxBUiIHTkVVVFJBTCoJSU5BTklNQVRFMgdVTktOT1dOOAJIBFAD
WgRnaWZ0YgFPaP///////////wFw////////////AXgFgAEFiAECkAEAmAEBoAEA
qAEAsAEAuAEAwAEAyAEA0AEA2AEB4AEB6AEA8gEGCAEQAyAA+gEWCP//////////
/wEQ////////////AYICBAgBEAOSAwQIARAAkgMECAEQAZIDBAgBEAKSAwQIARAD
kgMECAEQBJoDBAgBEAKaAwQIARADuAMEcuEBCAYSClBST05PTUlOQUwaBlBMVVJB
TCIHVU5LTk9XTioHQU5JTUFURTICV0U4A0gEUANaAndlYgFPaP///////////wFw
////////////AXgGgAEGiAEDkAEAmAEBoAEBqAEAsAEAuAEAwAEAyAEA0AEA2AEB
4AEB6AEA8gEGCAIQAyAA+gEGCAIQBCAAggIECAIQA5IDBAgCEACSAwQIAhABkgME
CAIQApIDBAgCEAOSAwQIAhAEkgMECAIQBZIDBAgCEAaSAwQIAhAHkgMECAIQCJID
BAgCEAmaAwQIAhADcuUBCAcSClBST05PTUlOQUwaCFNJTkdVTEFSIgdORVVUUkFM
KglJTkFOSU1BVEUyAklUOAVIBlAFWgJpdGIBT2j///////////8BcP//////////
/wF4B4ABB4gBA5ABAJgBAaABAKgBAbABALgBAMABAMgBANABANgBAeABAegBAPIB
BggCEAUgAPoBBggCEAQgAIICBAgCEAWSAwQIAhAAkgMECAIQAZIDBAgCEAKSAwQI
AhADkgMECAIQBJIDBAgCEAWSAwQIAhAGkgMECAIQB5IDBAgCEAiSAwQIAhAJmgME
CAIQBXL1AQgIEgZQUk9QRVIaCFNJTkdVTEFSIgdVTktOT1dOKglJTkFOSU1BVEUy
B1VOS05PV044B0gJUAhaB3ByZXNlbnRiBERBVEVo////////////AXD/////////
//8BeAiAAQiIAQOQAQCYAQGgAQCoAQGwAQC4AQDAAQDIAQDQAQDYAQHgAQHoAQDy
AQYIAhAIIAD6AQYIAhAEIACCAgQIAhAIkgMECAIQAJIDBAgCEAGSAwQIAhACkgME
CAIQA5IDBAgCEASSAwQIAhAFkgMECAIQBpIDBAgCEAeSAwQIAhAIkgMECAIQCZoD
BAgCEAeaAwQIAhAIeAGAAQGIAQCIAf///////////wGIAQGIAf///////////wGI
AQKIAf///////////wGIAf///////////wGIAf///////////wGIAf//////////
/wGQAQCQAQKQAQSQAf///////////wE=
`

	// RosesShortResp is the standard base64 (as defined in RFC 4648)
	// encoded response of annotating
	//
	//	Roses are red.
	//	  Violets are blue.
	//	Sugar is sweet.
	//	  And so are you.
	//
	// with annotators "tokenize,ssplit,pos" by Stanford CoreNLP 4.5.5.
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
	// with annotators "tokenize,ssplit,pos" by Stanford CoreNLP 4.5.5.
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
