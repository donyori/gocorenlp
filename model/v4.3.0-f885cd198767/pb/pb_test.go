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

package pb_test

import (
	"testing"

	"github.com/donyori/gocorenlp/internal/pbtest"
	"github.com/donyori/gocorenlp/model/v4.3.0-f885cd198767/pb"
)

func TestDecodeBase64Resp(t *testing.T) {
	if err := pbtest.CheckDocumentFromBase64(Base64ResponseBody, new(pb.Document)); err != nil {
		t.Error(err)
	}
}

// CoreNLP 4.3.0, 4.3.1, and 4.3.2 respond with the same content:

const Base64ResponseBody = `
pSsKRgpSb3NlcyBhcmUgcmVkLgogIFZpb2xldHMgYXJlIGJsdWUuClN1Z2FyIGlz
IHN3ZWV0LgogIEFuZCBzbyBhcmUgeW91LgoS2wgKZQoFUm9zZXMSBE5OUFMaBVJv
c2VzKgEKMgEgOgVSb3Nlc0IBT1IFUm9zZXNYAWAGaAByBFBFUjCIAQCQAQGoAQCw
AgDyAwFP+gMBT4AEAJIEFE89MC45OTA3MTQzMjI4MTEzODgyClgKA2FyZRIDVkJQ
GgNhcmUqASAyASA6A2FyZUIBT1ICYmVYB2AKaAByBFBFUjCIAQGQAQKoAQCwAgDy
AwFP+gMBT5IEFE89MC45OTk5OTU2NDEyMzAzOTEyClcKA3JlZBICSkoaA3JlZCoB
IDIAOgNyZWRCAU9SA3JlZFgLYA5oAHIEUEVSMIgBApABA6gBALACAPIDAU/6AwFP
kgQUTz0wLjk5OTUwMTU0MDM5MzMwODYKUAoBLhIBLhoBLioAMgMKICA6AS5CAU9S
AS5YDmAPaAByBFBFUjCIAQOQAQSoAQCwAgDyAwFP+gMBT5IEFE89MC45OTk5OTg2
MjA2MzkxMzE5EAAYBCAAKAEwD0JYCgQIABABCgQIABACCgQIABADCgQIABAEEhMI
AxABGgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRITCAMQBBoFcHVu
Y3QgACgAMAA4CRoBA0pYCgQIABABCgQIABACCgQIABADCgQIABAEEhMIAxABGgVu
c3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRITCAMQBBoFcHVuY3QgACgA
MAA4CRoBA1JYCgQIABABCgQIABACCgQIABADCgQIABAEEhMIAxABGgVuc3ViaiAA
KAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRITCAMQBBoFcHVuY3QgACgAMAA4CRoB
A1gBigFYCgQIABABCgQIABACCgQIABADCgQIABAEEhMIAxABGgVuc3ViaiAAKAAw
ADgJEhEIAxACGgNjb3AgACgAMAA4CRITCAMQBBoFcHVuY3QgACgAMAA4CRoBA5IB
WAoECAAQAQoECAAQAgoECAAQAwoECAAQBBITCAMQARoFbnN1YmogACgAMAA4CRIR
CAMQAhoDY29wIAAoADAAOAkSEwgDEAQaBXB1bmN0IAAoADAAOAkaAQOYAwCwAwDC
A4wCCAASBlBST1BFUhoGUExVUkFMIgdVTktOT1dOKglJTkFOSU1BVEUyB1VOS05P
V044AEgBUABaBXJvc2VzYgFPaP///////////wFw////////////AXgAgAEAiAEA
kAEAmAEBoAEAqAEAsAEAuAEAwAEAyAEA0AEA2AEB4AEB6AEA8gEPCP//////////
/wEQACAA+gEWCP///////////wEQ////////////AYICDQj///////////8BEACS
Aw0I////////////ARAAkgMNCP///////////wEQAZIDDQj///////////8BEAKS
Aw0I////////////ARADmgMNCP///////////wEQAMgDAYgEAaAEARKpCAptCgdW
aW9sZXRzEgNOTlMaB1Zpb2xldHMqAwogIDIBIDoHVmlvbGV0c0IBT1IGdmlvbGV0
WBJgGWgAcgRQRVIwiAEEkAEFqAEAsAIA8gMBT/oDAU+ABAGSBBRPPTAuOTkwNDM3
ODg1NTM3ODAxMQpYCgNhcmUSA1ZCUBoDYXJlKgEgMgEgOgNhcmVCAU9SAmJlWBpg
HWgAcgRQRVIwiAEFkAEGqAEAsAIA8gMBT/oDAU+SBBRPPTAuOTk5OTk5MjUyOTIy
OTIwOQpbCgRibHVlEgJKShoEYmx1ZSoBIDIAOgRibHVlQgFPUgRibHVlWB5gImgA
cgRQRVIwiAEGkAEHqAEAsAIA8gMBT/oDAU+SBBRPPTAuOTk5NTc0MzkyMjQyOTc5
MgpOCgEuEgEuGgEuKgAyAQo6AS5CAU9SAS5YImAjaAByBFBFUjCIAQeQAQioAQCw
AgDyAwFP+gMBT5IEFE89MC45OTk5OTkzNjQ5MjI1MzI5EAQYCCABKBIwI0JYCgQI
ARADCgQIARAECgQIARABCgQIARACEhMIAxAEGgVwdW5jdCAAKAAwADgJEhMIAxAB
GgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRoBA0pYCgQIARADCgQI
ARAECgQIARABCgQIARACEhMIAxAEGgVwdW5jdCAAKAAwADgJEhMIAxABGgVuc3Vi
aiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRoBA1JYCgQIARADCgQIARAECgQI
ARABCgQIARACEhMIAxAEGgVwdW5jdCAAKAAwADgJEhMIAxABGgVuc3ViaiAAKAAw
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
sAIA8gMBT/oDAU+ABAOSBBRPPTAuOTk5OTkxMTA3OTU5MDQ3MwpOCgEuEgEuGgEu
KgAyAQo6AS5CAU9SAS5YRGBFaAByBFBFUjCIARCQARGoAQCwAgDyAwFP+gMBT5IE
FE89MC45OTk5OTkxNjA2MTUzNzEyEAwYESADKDYwRUJxCgQIAxABCgQIAxACCgQI
AxADCgQIAxAECgQIAxAFEhAIBBABGgJjYyAAKAAwADgJEhQIBBACGgZhZHZtb2Qg
ACgAMAA4CRIRCAQQAxoDY29wIAAoADAAOAkSEwgEEAUaBXB1bmN0IAAoADAAOAka
AQRKcQoECAMQAQoECAMQAgoECAMQAwoECAMQBAoECAMQBRIQCAQQARoCY2MgACgA
MAA4CRIUCAQQAhoGYWR2bW9kIAAoADAAOAkSEQgEEAMaA2NvcCAAKAAwADgJEhMI
BBAFGgVwdW5jdCAAKAAwADgJGgEEUnEKBAgDEAEKBAgDEAIKBAgDEAMKBAgDEAQK
BAgDEAUSEAgEEAEaAmNjIAAoADAAOAkSFAgEEAIaBmFkdm1vZCAAKAAwADgJEhEI
BBADGgNjb3AgACgAMAA4CRITCAQQBRoFcHVuY3QgACgAMAA4CRoBBFgDigFxCgQI
AxABCgQIAxACCgQIAxADCgQIAxAECgQIAxAFEhAIBBABGgJjYyAAKAAwADgJEhQI
BBACGgZhZHZtb2QgACgAMAA4CRIRCAQQAxoDY29wIAAoADAAOAkSEwgEEAUaBXB1
bmN0IAAoADAAOAkaAQSSAXEKBAgDEAEKBAgDEAIKBAgDEAMKBAgDEAQKBAgDEAUS
EAgEEAEaAmNjIAAoADAAOAkSFAgEEAIaBmFkdm1vZCAAKAAwADgJEhEIBBADGgNj
b3AgACgAMAA4CRITCAQQBRoFcHVuY3QgACgAMAA4CRoBBJgDALADAMID0QEIAxIK
UFJPTk9NSU5BTBoHVU5LTk9XTiIHVU5LTk9XTioHQU5JTUFURTIDWU9VOANIBFAD
WgN5b3ViAU9o////////////AXD///////////8BeAOAAQOIAQOQAQCYAQOgAQCo
AQCwAQC4AQDAAQDIAQDQAQDYAQHgAQHoAQDyAQYIAhADIAD6ARYI////////////
ARD///////////8BggIECAIQA5IDBAgCEACSAwQIAhABkgMECAIQApIDBAgCEAOS
AwQIAhAEmgMECAIQA8gDAYgEAaAEAVgAaAFyjAIIABIGUFJPUEVSGgZQTFVSQUwi
B1VOS05PV04qCUlOQU5JTUFURTIHVU5LTk9XTjgASAFQAFoFcm9zZXNiAU9o////
////////AXD///////////8BeACAAQCIAQCQAQCYAQGgAQCoAQCwAQC4AQDAAQDI
AQDQAQDYAQHgAQHoAQDyAQ8I////////////ARAAIAD6ARYI////////////ARD/
//////////8BggINCP///////////wEQAJIDDQj///////////8BEACSAw0I////
////////ARABkgMNCP///////////wEQApIDDQj///////////8BEAOaAw0I////
////////ARAActABCAESB05PTUlOQUwaBlBMVVJBTCIHVU5LTk9XTioJSU5BTklN
QVRFMgdVTktOT1dOOABIAVAAWgd2aW9sZXRzYgFPaP///////////wFw////////
////AXgBgAEBiAEBkAEAmAECoAEAqAEAsAEAuAEAwAEAyAEA0AEA2AEB4AEB6AEA
8gEGCAAQACAA+gEWCP///////////wEQ////////////AYICBAgAEACSAwQIABAA
kgMECAAQAZIDBAgAEAKSAwQIABADmgMECAAQAHLPAQgCEgZQUk9QRVIaCFNJTkdV
TEFSIgdORVVUUkFMKglJTkFOSU1BVEUyB1VOS05PV044AEgBUABaBXN1Z2FyYgFP
aP///////////wFw////////////AXgCgAECiAECkAEAmAECoAEAqAEAsAEAuAEA
wAEAyAEA0AEA2AEB4AEB6AEA8gEGCAEQACAA+gEWCP///////////wEQ////////
////AYICBAgBEACSAwQIARAAkgMECAEQAZIDBAgBEAKSAwQIARADmgMECAEQAHLR
AQgDEgpQUk9OT01JTkFMGgdVTktOT1dOIgdVTktOT1dOKgdBTklNQVRFMgNZT1U4
A0gEUANaA3lvdWIBT2j///////////8BcP///////////wF4A4ABA4gBA5ABAJgB
A6ABAKgBALABALgBAMABAMgBANABANgBAeABAegBAPIBBggCEAMgAPoBFgj/////
//////8BEP///////////wGCAgQIAhADkgMECAIQAJIDBAgCEAGSAwQIAhACkgME
CAIQA5IDBAgCEASaAwQIAhADeAGAAQGIAf///////////wGIAf///////////wGI
Af///////////wGIAf///////////wE=
`
