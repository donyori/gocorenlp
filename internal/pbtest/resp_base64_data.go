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

package pbtest

// RosesAreRedRespV360 is the standard base64 (as defined in RFC 4648) encoded
// response of annotating RosesAreRed with the server default annotators by
// Stanford CoreNLP 3.6.0.
const RosesAreRedRespV360 = `
nx4KRgpSb3NlcyBhcmUgcmVkLgogIFZpb2xldHMgYXJlIGJsdWUuClN1Z2FyIGlz
IHN3ZWV0LgogIEFuZCBzbyBhcmUgeW91LgoSgQcKUQoFUm9zZXMSBE5OUFMaBVJv
c2VzKgEKMgEgOgVSb3Nlc0IBT1IFUm9zZXNYAWAGaAByBFBFUjCIAQCQAQGoAQDi
AQ4IABABGAIgAygEMAU4BgpHCgNhcmUSA1ZCUBoDYXJlKgEgMgEgOgNhcmVCAU9S
AmJlWAdgCmgAcgRQRVIwiAEBkAECqAEA4gEOCAAQARgCIAMoBDAFOAYKRgoDcmVk
EgJKShoDcmVkKgEgMgA6A3JlZEIBT1IDcmVkWAtgDmgAcgRQRVIwiAECkAEDqAEA
4gEOCAAQARgCIAMoBDAFOAYKPwoBLhIBLhoBLioAMgMKICA6AS5CAU9SAS5YDmAP
aAByBFBFUjCIAQOQAQSoAQDiAQ4IABABGAIgAygEMAU4BhAAGAQgACgBMA9CWAoE
CAAQAQoECAAQAgoECAAQAwoECAAQBBITCAMQARoFbnN1YmogACgAMAA4CRIRCAMQ
AhoDY29wIAAoADAAOAkSEwgDEAQaBXB1bmN0IAAoADAAOAkaAQNKPQoECAAQAQoE
CAAQAgoECAAQAxITCAMQARoFbnN1YmogACgAMAA4CRIRCAMQAhoDY29wIAAoADAA
OAkaAQNSWAoECAAQAQoECAAQAgoECAAQAwoECAAQBBITCAMQARoFbnN1YmogACgA
MAA4CRIRCAMQAhoDY29wIAAoADAAOAkSEwgDEAQaBXB1bmN0IAAoADAAOAkaAQNY
AXJfCgVSb3NlcxIDYXJlGgNyZWQhAAAAAAAA8D8oADABOAJCPQoECAAQAQoECAAQ
AgoECAAQAxITCAMQARoFbnN1YmogACgAMAA4CRIRCAMQAhoDY29wIAAoADAAOAka
AQN6EwgACAEIAhACGAEhAAAAAAAA8D+YAwCwAwHCA9oBCAASBlBST1BFUhoGUExV
UkFMIgdVTktOT1dOKglJTkFOSU1BVEUyB1VOS05PV044AEgBUABaBXJvc2VzYgFP
aP////8PcP///////////wF4AIABAIgBAJABAJgBAaABAKgBALABALgBAMABAMgB
ANABANgBAeABAegBAPIBCgj/////DxAAIAD6AQwI/////w8Q/////w+CAggI////
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
D4ICBAgBEACSAwQIARAAkgMECAEQAZIDBAgBEAKSAwQIARADmgMECAEQAMgDARLR
BwpJCgNBbmQSAkNDGgNBbmQqAwogIDIBIDoDQW5kQgFPUgNhbmRYNmA5aAByBFBF
UjCIAQyQAQ2oAQDiAQ4IABABGAIgAygEMAU4BgpDCgJzbxICUkIaAnNvKgEgMgEg
OgJzb0IBT1ICc29YOmA8aAByBFBFUjCIAQ2QAQ6oAQDiAQ4IABABGAIgAygEMAU4
BgpHCgNhcmUSA1ZCUBoDYXJlKgEgMgEgOgNhcmVCAU9SAmJlWD1gQGgAcgRQRVIw
iAEOkAEPqAEA4gEOCAAQARgCIAMoBDAFOAYKSgoDeW91EgNQUlAaA3lvdSoBIDIA
OgN5b3VCAU9SA3lvdVhBYERoAHIEUEVSMIgBD5ABEKgBALgBA+IBDggAEAEYAiAD
KAQwBTgGCj0KAS4SAS4aAS4qADIBCjoBLkIBT1IBLlhEYEVoAHIEUEVSMIgBEJAB
EagBAOIBDggAEAEYAiADKAQwBTgGEAwYESADKDYwRUJzCgQIAxABCgQIAxACCgQI
AxADCgQIAxAECgQIAxAFEhAIAxABGgJjYyAAKAAwADgJEhQIAxACGgZhZHZtb2Qg
ACgAMAA4CRITCAMQBBoFbnN1YmogACgAMAA4CRITCAMQBRoFcHVuY3QgACgAMAA4
CRoBA0pYCgQIAxABCgQIAxACCgQIAxADCgQIAxAEEhAIAxABGgJjYyAAKAAwADgJ
EhQIAxACGgZhZHZtb2QgACgAMAA4CRITCAMQBBoFbnN1YmogACgAMAA4CRoBA1Jz
CgQIAxABCgQIAxACCgQIAxADCgQIAxAECgQIAxAFEhAIAxABGgJjYyAAKAAwADgJ
EhQIAxACGgZhZHZtb2QgACgAMAA4CRITCAMQBBoFbnN1YmogACgAMAA4CRITCAMQ
BRoFcHVuY3QgACgAMAA4CRoBA1gDehEIAggDEAIYASEAAAAAAADwP3oVCAAIAQgC
CAMQAhgBIQAAAAAAAPA/ehMIAAgCCAMQAhgBIQAAAAAAAPA/ehMIAQgCCAMQAhgB
IQAAAAAAAPA/mAMAsAMBwgO8AQgDEgpQUk9OT01JTkFMGgdVTktOT1dOIgdVTktO
T1dOKgdBTklNQVRFMgNZT1U4A0gEUANaA3lvdWIBT2j/////D3D///////////8B
eAOAAQOIAQOQAQCYAQOgAQGoAQCwAQC4AQDAAQDIAQDQAQDYAQHgAQHoAQDyAQYI
AhADIAD6AQYIAhACIACCAgQIAhADkgMECAIQAJIDBAgCEAGSAwQIAhACkgMECAIQ
A5IDBAgCEASaAwQIAhADyAMBGm0IAxIyCAISBlBST1BFUhoIU0lOR1VMQVIiB05F
VVRSQUwqCUlOQU5JTUFURTAAOAFIAFACWAESMwgDEgpQUk9OT01JTkFMGgdVTktO
T1dOIgdVTktOT1dOKgdBTklNQVRFMAM4BEgDUANYARgA
`

// RosesAreRedRespV400 is the standard base64 (as defined in RFC 4648) encoded
// response of annotating RosesAreRed with the server default annotators by
// Stanford CoreNLP 4.0.0.
const RosesAreRedRespV400 = `
+zoKRgpSb3NlcyBhcmUgcmVkLgogIFZpb2xldHMgYXJlIGJsdWUuClN1Z2FyIGlz
IHN3ZWV0LgogIEFuZCBzbyBhcmUgeW91LgoS1gwKgAEKBVJvc2VzEgROTlBTGgVS
b3NlcyoBCjIBIDoFUm9zZXNCAU9SBVJvc2VzWAFgBmgAcgRQRVIweACAAQGIAQCQ
AQGoAQDiAQ4IABABGAIgAygEMAU4BrACALoCAnVw8gMBT/oDAU+ABACSBBRPPTAu
OTkwNzE0MzIxMzQ0MzYwMQpzCgNhcmUSA1ZCUBoDYXJlKgEgMgEgOgNhcmVCAU9S
AmJlWAdgCmgAcgRQRVIweAGAAQKIAQGQAQKoAQDiAQ4IABABGAIgAygEMAU4BrAC
ALoCAnVw8gMBT/oDAU+SBBRPPTAuOTk5OTk1NjQxMjMwMDY0MwpyCgNyZWQSAkpK
GgNyZWQqASAyADoDcmVkQgFPUgNyZWRYC2AOaAByBFBFUjB4AoABA4gBApABA6gB
AOIBDggAEAEYAiADKAQwBTgGsAIAugICdXDyAwFP+gMBT5IEFE89MC45OTk1MDE1
NDA0NDc5MzY0CmsKAS4SAS4aAS4qADIDCiAgOgEuQgFPUgEuWA5gD2gAcgRQRVIw
eAOAAQSIAQOQAQSoAQDiAQ4IABABGAIgAygEMAU4BrACALoCAnVw8gMBT/oDAU+S
BBRPPTAuOTk5OTk4NjIwNjM5MjI0MxAAGAQgACgBMA86pwEKlQEKJwoYCgcSBVJv
c2VzEgROTlBTKQAAAEBCQi7AEgJOUCkAAACgF341wApLChUKBRIDYXJlEgNWQlAp
AAAAQKVWzL8KJQoUCgUSA3JlZBICSkopAAAAQNF3G8ASBEFESlApAAAAIM5IHsAS
AlZQKQAAAAADsynAChEKAxIBLhIBLikAAACA63StvxIBUykAAAAAncRBwBIEUk9P
VCkAAAAAktpBwEJYCgQIABABCgQIABACCgQIABADCgQIABAEEhMIAxABGgVuc3Vi
aiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRITCAMQBBoFcHVuY3QgACgAMAA4
CRoBA0pYCgQIABABCgQIABACCgQIABADCgQIABAEEhMIAxABGgVuc3ViaiAAKAAw
ADgJEhEIAxACGgNjb3AgACgAMAA4CRITCAMQBBoFcHVuY3QgACgAMAA4CRoBA1JY
CgQIABABCgQIABACCgQIABADCgQIABAEEhMIAxABGgVuc3ViaiAAKAAwADgJEhEI
AxACGgNjb3AgACgAMAA4CRITCAMQBBoFcHVuY3QgACgAMAA4CRoBA1gBcmsKBVJv
c2VzEgNhcmUaA3JlZCEAAAAAAADwP0I9CgQIABABCgQIABACCgQIABADEhMIAxAB
GgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRoBA2oECAAQAHIECAAQ
AXoECAAQAooBWAoECAAQAQoECAAQAgoECAAQAwoECAAQBBITCAMQARoFbnN1Ymog
ACgAMAA4CRIRCAMQAhoDY29wIAAoADAAOAkSEwgDEAQaBXB1bmN0IAAoADAAOAka
AQOSAVgKBAgAEAEKBAgAEAIKBAgAEAMKBAgAEAQSEwgDEAEaBW5zdWJqIAAoADAA
OAkSEQgDEAIaA2NvcCAAKAAwADgJEhMIAxAEGgVwdW5jdCAAKAAwADgJGgED+gFb
ClMKFQoPCgcSBVJvc2VzEgROTlBTEgJOUAo3CicKDAoFEgNhcmUSA1ZCUAoTCgsK
BRIDcmVkEgJKShIEQURKUBICVlAKCAoDEgEuEgEuEgJAUxIBUxIEUk9PVJoCEwgA
CAEIAhACGAEhAAAAAAAA8D+YAwCwAwDCA4wCCAASBlBST1BFUhoGUExVUkFMIgdV
TktOT1dOKglJTkFOSU1BVEUyB1VOS05PV044AEgBUABaBXJvc2VzYgFPaP//////
/////wFw////////////AXgAgAEAiAEAkAEAmAEBoAEAqAEAsAEAuAEAwAEAyAEA
0AEA2AEB4AEB6AEB8gEPCP///////////wEQACAA+gEWCP///////////wEQ////
////////AYICDQj///////////8BEACSAw0I////////////ARAAkgMNCP//////
/////wEQAZIDDQj///////////8BEAKSAw0I////////////ARADmgMNCP//////
/////wEQAMgDAYgEAaAEAagEARKrDAqIAQoHVmlvbGV0cxIDTk5TGgdWaW9sZXRz
KgMKICAyASA6B1Zpb2xldHNCAU9SBnZpb2xldFgSYBloAHIEUEVSMHgAgAEBiAEE
kAEFqAEA4gEOCAAQARgCIAMoBDAFOAawAgC6AgJ1cPIDAU/6AwFPgAQBkgQUTz0w
Ljk5MDQzNzg4NDAxMTQ5MDUKcwoDYXJlEgNWQlAaA2FyZSoBIDIBIDoDYXJlQgFP
UgJiZVgaYB1oAHIEUEVSMHgBgAECiAEFkAEGqAEA4gEOCAAQARgCIAMoBDAFOAaw
AgC6AgJ1cPIDAU/6AwFPkgQUTz0wLjk5OTk5OTI1MjkyMjg5OTYKdgoEYmx1ZRIC
SkoaBGJsdWUqASAyADoEYmx1ZUIBT1IEYmx1ZVgeYCJoAHIEUEVSMHgCgAEDiAEG
kAEHqAEA4gEOCAAQARgCIAMoBDAFOAawAgC6AgJ1cPIDAU/6AwFPkgQUTz0wLjk5
OTU3NDM5MjI3MDM1MTgKaQoBLhIBLhoBLioAMgEKOgEuQgFPUgEuWCJgI2gAcgRQ
RVIweAOAAQSIAQeQAQioAQDiAQ4IABABGAIgAygEMAU4BrACALoCAnVw8gMBT/oD
AU+SBBRPPTAuOTk5OTk5MzY0OTIyNTc1NhAEGAggASgSMCM6qQEKlwEKKAoZCgkS
B1Zpb2xldHMSA05OUykAAABgdOcowBICTlApAAAAYFokL8AKTAoVCgUSA2FyZRID
VkJQKQAAAEClVsy/CiYKFQoGEgRibHVlEgJKSikAAADAY1AewBIEQURKUCkAAABA
sJAgwBICVlApAAAAQEwfK8AKEQoDEgEuEgEuKQAAAIDrdK2/EgFTKQAAAEB0Uz7A
EgRST09UKQAAACBefz7AQlgKBAgBEAMKBAgBEAQKBAgBEAEKBAgBEAISEwgDEAQa
BXB1bmN0IAAoADAAOAkSEwgDEAEaBW5zdWJqIAAoADAAOAkSEQgDEAIaA2NvcCAA
KAAwADgJGgEDSlgKBAgBEAMKBAgBEAQKBAgBEAEKBAgBEAISEwgDEAQaBXB1bmN0
IAAoADAAOAkSEwgDEAEaBW5zdWJqIAAoADAAOAkSEQgDEAIaA2NvcCAAKAAwADgJ
GgEDUlgKBAgBEAMKBAgBEAQKBAgBEAEKBAgBEAISEwgDEAQaBXB1bmN0IAAoADAA
OAkSEwgDEAEaBW5zdWJqIAAoADAAOAkSEQgDEAIaA2NvcCAAKAAwADgJGgEDWAJy
bgoHVmlvbGV0cxIDYXJlGgRibHVlIQAAAAAAAPA/Qj0KBAgBEAMKBAgBEAEKBAgB
EAISEwgDEAEaBW5zdWJqIAAoADAAOAkSEQgDEAIaA2NvcCAAKAAwADgJGgEDagQI
ARAAcgQIARABegQIARACigFYCgQIARADCgQIARAECgQIARABCgQIARACEhMIAxAE
GgVwdW5jdCAAKAAwADgJEhMIAxABGgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3Ag
ACgAMAA4CRoBA5IBWAoECAEQAwoECAEQBAoECAEQAQoECAEQAhITCAMQBBoFcHVu
Y3QgACgAMAA4CRITCAMQARoFbnN1YmogACgAMAA4CRIRCAMQAhoDY29wIAAoADAA
OAkaAQP6AV0KVQoWChAKCRIHVmlvbGV0cxIDTk5TEgJOUAo4CigKDAoFEgNhcmUS
A1ZCUAoUCgwKBhIEYmx1ZRICSkoSBEFESlASAlZQCggKAxIBLhIBLhICQFMSAVMS
BFJPT1SaAhMIAAgBCAIQAhgBIQAAAAAAAPA/mAMAsAMAwgPQAQgBEgdOT01JTkFM
GgZQTFVSQUwiB1VOS05PV04qCUlOQU5JTUFURTIHVU5LTk9XTjgASAFQAFoHdmlv
bGV0c2IBT2j///////////8BcP///////////wF4AYABAYgBAZABAJgBAqABAKgB
ALABALgBAMABAMgBANABANgBAeABAegBAfIBBggAEAAgAPoBFgj///////////8B
EP///////////wGCAgQIABAAkgMECAAQAJIDBAgAEAGSAwQIABACkgMECAAQA5oD
BAgAEADIAwGIBAGgBAGoBAESnAwKfwoFU3VnYXISA05OUBoFU3VnYXIqAQoyASA6
BVN1Z2FyQgFPUgVTdWdhclgkYCloAHIEUEVSMHgAgAEBiAEIkAEJqAEA4gEOCAAQ
ARgCIAMoBDAFOAawAgC6AgJ1cPIDAU/6AwFPgAQCkgQUTz0wLjk4OTM5NDQwODMy
NDYzMDIKbwoCaXMSA1ZCWhoCaXMqASAyASA6AmlzQgFPUgJiZVgqYCxoAHIEUEVS
MHgBgAECiAEJkAEKqAEA4gEOCAAQARgCIAMoBDAFOAawAgC6AgJ1cPIDAU/6AwFP
kgQTTz0wLjk5OTk5ODE1MzQxMjYyNgp6CgVzd2VldBICSkoaBXN3ZWV0KgEgMgA6
BXN3ZWV0QgFPUgVzd2VldFgtYDJoAHIEUEVSMHgCgAEDiAEKkAELqAEA4gEOCAAQ
ARgCIAMoBDAFOAawAgC6AgJ1cPIDAU/6AwFPkgQUTz0wLjk5OTk3MDIxMDc4NDk2
MDYKawoBLhIBLhoBLioAMgMKICA6AS5CAU9SAS5YMmAzaAByBFBFUjB4A4ABBIgB
C5ABDKgBAOIBDggAEAEYAiADKAQwBTgGsAIAugICdXDyAwFP+gMBT5IEFE89MC45
OTk5OTk3OTM4OTc3NzEzEAgYDCACKCQwMzqnAQqVAQomChcKBxIFU3VnYXISA05O
UCkAAABAYLomwBICTlApAAAAgLLPK8AKTAoUCgQSAmlzEgNWQlopAAAAQNTwwr8K
JwoWCgcSBXN3ZWV0EgJKSikAAAAAlI4cwBIEQURKUCkAAADgkF8fwBICVlApAAAA
QN3cKMAKEQoDEgEuEgEuKQAAAIDrdK2/EgFTKQAAAMDohzvAEgRST09UKQAAAKDS
szvAQlgKBAgCEAEKBAgCEAIKBAgCEAMKBAgCEAQSEwgDEAEaBW5zdWJqIAAoADAA
OAkSEQgDEAIaA2NvcCAAKAAwADgJEhMIAxAEGgVwdW5jdCAAKAAwADgJGgEDSlgK
BAgCEAEKBAgCEAIKBAgCEAMKBAgCEAQSEwgDEAEaBW5zdWJqIAAoADAAOAkSEQgD
EAIaA2NvcCAAKAAwADgJEhMIAxAEGgVwdW5jdCAAKAAwADgJGgEDUlgKBAgCEAEK
BAgCEAIKBAgCEAMKBAgCEAQSEwgDEAEaBW5zdWJqIAAoADAAOAkSEQgDEAIaA2Nv
cCAAKAAwADgJEhMIAxAEGgVwdW5jdCAAKAAwADgJGgEDWAJybAoFU3VnYXISAmlz
GgVzd2VldCEAAAAAAADwP0I9CgQIAhABCgQIAhACCgQIAhADEhMIAxABGgVuc3Vi
aiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRoBA2oECAIQAHIECAIQAXoECAIQ
AooBWAoECAIQAQoECAIQAgoECAIQAwoECAIQBBITCAMQARoFbnN1YmogACgAMAA4
CRIRCAMQAhoDY29wIAAoADAAOAkSEwgDEAQaBXB1bmN0IAAoADAAOAkaAQOSAVgK
BAgCEAEKBAgCEAIKBAgCEAMKBAgCEAQSEwgDEAEaBW5zdWJqIAAoADAAOAkSEQgD
EAIaA2NvcCAAKAAwADgJEhMIAxAEGgVwdW5jdCAAKAAwADgJGgED+gFbClMKFAoO
CgcSBVN1Z2FyEgNOTlASAk5QCjgKKAoLCgQSAmlzEgNWQloKFQoNCgcSBXN3ZWV0
EgJKShIEQURKUBICVlAKCAoDEgEuEgEuEgJAUxIBUxIEUk9PVJoCEwgACAEIAhAC
GAEhAAAAAAAA8D+YAwCwAwDCA88BCAISBlBST1BFUhoIU0lOR1VMQVIiB05FVVRS
QUwqCUlOQU5JTUFURTIHVU5LTk9XTjgASAFQAFoFc3VnYXJiAU9o////////////
AXD///////////8BeAKAAQKIAQKQAQCYAQKgAQCoAQCwAQC4AQDAAQDIAQDQAQDY
AQHgAQHoAQHyAQYIARAAIAD6ARYI////////////ARD///////////8BggIECAEQ
AJIDBAgBEACSAwQIARABkgMECAEQApIDBAgBEAOaAwQIARAAyAMBiAQBoAQBqAQB
EskNCnUKA0FuZBICQ0MaA0FuZCoDCiAgMgEgOgNBbmRCAU9SA2FuZFg2YDloAHIE
UEVSMHgAgAEBiAEMkAENqAEA4gEOCAAQARgCIAMoBDAFOAawAgC6AgJ1cPIDAU/6
AwFPkgQUTz0wLjk5OTk5MDU0OTQ4ODUzNTYKbwoCc28SAlJCGgJzbyoBIDIBIDoC
c29CAU9SAnNvWDpgPGgAcgRQRVIweAGAAQKIAQ2QAQ6oAQDiAQ4IABABGAIgAygE
MAU4BrACALoCAnVw8gMBT/oDAU+SBBRPPTAuOTk5OTk1NTU5NDU3NTcwNwpzCgNh
cmUSA1ZCUBoDYXJlKgEgMgEgOgNhcmVCAU9SAmJlWD1gQGgAcgRQRVIweAKAAQOI
AQ6QAQ+oAQDiAQ4IABABGAIgAygEMAU4BrACALoCAnVw8gMBT/oDAU+SBBRPPTAu
OTk5OTk5NTgyMTM2NzY4Mgp2CgN5b3USA1BSUBoDeW91KgEgMgA6A3lvdUIBT1ID
eW91WEFgRGgAcgRQRVIweAOAAQSIAQ+QARCoAQDiAQ4IABABGAIgAygEMAU4BrAC
ALoCAnVw8gMBT/oDAU+ABAOSBBRPPTAuOTk5OTkxMTA3OTU5MjQ2MwpoCgEuEgEu
GgEuKgAyAQo6AS5CAU9SAS5YRGBFaAByBFBFUjB4BIABBYgBEJABEagBAOIBDggA
EAEYAiADKAQwBTgGsAIAugICdXDyAwFP+gMBT5IEE089MC45OTk5OTkxNjA2MTUz
NTcQDBgRIAMoNjBFOrwBCqoBChQKBRIDQW5kEgJDQykAAAAAWeH3vwokChMKBBIC
c28SAlJCKQAAAMCO3xDAEgRBRFZQKQAAAOC8mBHACiQKFQoFEgNhcmUSA1ZCUCkA
AABApVbMvxICVlApAAAA4BflFMAKJAoVCgUSA3lvdRIDUFJQKQAAAODaVQHAEgJO
UCkAAABAEvUPwAoRCgMSAS4SAS4pAAAAQE3VnL8SBFNJTlYpAAAAAJl0NsASBFJP
T1QpAAAAQDMpOsBCcQoECAMQAQoECAMQAgoECAMQAwoECAMQBAoECAMQBRIQCAQQ
ARoCY2MgACgAMAA4CRIUCAQQAhoGYWR2bW9kIAAoADAAOAkSEQgEEAMaA2NvcCAA
KAAwADgJEhMIBBAFGgVwdW5jdCAAKAAwADgJGgEESnEKBAgDEAEKBAgDEAIKBAgD
EAMKBAgDEAQKBAgDEAUSEAgEEAEaAmNjIAAoADAAOAkSFAgEEAIaBmFkdm1vZCAA
KAAwADgJEhEIBBADGgNjb3AgACgAMAA4CRITCAQQBRoFcHVuY3QgACgAMAA4CRoB
BFJxCgQIAxABCgQIAxACCgQIAxADCgQIAxAECgQIAxAFEhAIBBABGgJjYyAAKAAw
ADgJEhQIBBACGgZhZHZtb2QgACgAMAA4CRIRCAQQAxoDY29wIAAoADAAOAkSEwgE
EAUaBXB1bmN0IAAoADAAOAkaAQRYA4oBcQoECAMQAQoECAMQAgoECAMQAwoECAMQ
BAoECAMQBRIQCAQQARoCY2MgACgAMAA4CRIUCAQQAhoGYWR2bW9kIAAoADAAOAkS
EQgEEAMaA2NvcCAAKAAwADgJEhMIBBAFGgVwdW5jdCAAKAAwADgJGgEEkgFxCgQI
AxABCgQIAxACCgQIAxADCgQIAxAECgQIAxAFEhAIBBABGgJjYyAAKAAwADgJEhQI
BBACGgZhZHZtb2QgACgAMAA4CRIRCAQQAxoDY29wIAAoADAAOAkSEwgEEAUaBXB1
bmN0IAAoADAAOAkaAQT6AXwKdAoLCgUSA0FuZBICQ0MKXwoSCgoKBBICc28SAlJC
EgRBRFZQCkIKLwoSCgwKBRIDYXJlEgNWQlASAlZQChIKDAoFEgN5b3USA1BSUBIC
TlASBUBTSU5WCggKAxIBLhIBLhIFQFNJTlYSBUBTSU5WEgRTSU5WEgRST09UmgIV
CAAIAQgCCAMQAxgBIQAAAAAAAPA/mAMAsAMAwgPRAQgDEgpQUk9OT01JTkFMGgdV
TktOT1dOIgdVTktOT1dOKgdBTklNQVRFMgNZT1U4A0gEUANaA3lvdWIBT2j/////
//////8BcP///////////wF4A4ABA4gBA5ABAJgBA6ABAKgBALABALgBAMABAMgB
ANABANgBAeABAegBAfIBBggCEAMgAPoBFgj///////////8BEP///////////wGC
AgQIAhADkgMECAIQAJIDBAgCEAGSAwQIAhACkgMECAIQA5IDBAgCEASaAwQIAhAD
yAMBiAQBoAQBqAQBWABoAXKMAggAEgZQUk9QRVIaBlBMVVJBTCIHVU5LTk9XTioJ
SU5BTklNQVRFMgdVTktOT1dOOABIAVAAWgVyb3Nlc2IBT2j///////////8BcP//
/////////wF4AIABAIgBAJABAJgBAaABAKgBALABALgBAMABAMgBANABANgBAeAB
AegBAfIBDwj///////////8BEAAgAPoBFgj///////////8BEP///////////wGC
Ag0I////////////ARAAkgMNCP///////////wEQAJIDDQj///////////8BEAGS
Aw0I////////////ARACkgMNCP///////////wEQA5oDDQj///////////8BEABy
0AEIARIHTk9NSU5BTBoGUExVUkFMIgdVTktOT1dOKglJTkFOSU1BVEUyB1VOS05P
V044AEgBUABaB3Zpb2xldHNiAU9o////////////AXD///////////8BeAGAAQGI
AQGQAQCYAQKgAQCoAQCwAQC4AQDAAQDIAQDQAQDYAQHgAQHoAQHyAQYIABAAIAD6
ARYI////////////ARD///////////8BggIECAAQAJIDBAgAEACSAwQIABABkgME
CAAQApIDBAgAEAOaAwQIABAAcs8BCAISBlBST1BFUhoIU0lOR1VMQVIiB05FVVRS
QUwqCUlOQU5JTUFURTIHVU5LTk9XTjgASAFQAFoFc3VnYXJiAU9o////////////
AXD///////////8BeAKAAQKIAQKQAQCYAQKgAQCoAQCwAQC4AQDAAQDIAQDQAQDY
AQHgAQHoAQHyAQYIARAAIAD6ARYI////////////ARD///////////8BggIECAEQ
AJIDBAgBEACSAwQIARABkgMECAEQApIDBAgBEAOaAwQIARAActEBCAMSClBST05P
TUlOQUwaB1VOS05PV04iB1VOS05PV04qB0FOSU1BVEUyA1lPVTgDSARQA1oDeW91
YgFPaP///////////wFw////////////AXgDgAEDiAEDkAEAmAEDoAEAqAEAsAEA
uAEAwAEAyAEA0AEA2AEB4AEB6AEB8gEGCAIQAyAA+gEWCP///////////wEQ////
////////AYICBAgCEAOSAwQIAhAAkgMECAIQAZIDBAgCEAKSAwQIAhADkgMECAIQ
BJoDBAgCEAN4AYABAYgB////////////AYgB////////////AYgB////////////
AYgB////////////AQ==
`

// RosesAreRedRespV410 is the standard base64 (as defined in RFC 4648) encoded
// response of annotating RosesAreRed with the server default annotators by
// Stanford CoreNLP 4.1.0.
const RosesAreRedRespV410 = `
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

// RosesAreRedRespV420 is the standard base64 (as defined in RFC 4648) encoded
// response of annotating RosesAreRed with the server default annotators by
// Stanford CoreNLP 4.2.0.
//
// It is the same as RosesAreRedRespV410.
const RosesAreRedRespV420 = RosesAreRedRespV410

// RosesAreRedRespV421 is the standard base64 (as defined in RFC 4648) encoded
// response of annotating RosesAreRed with the server default annotators by
// Stanford CoreNLP 4.2.1.
//
// It is the same as RosesAreRedRespV410.
const RosesAreRedRespV421 = RosesAreRedRespV410

// RosesAreRedRespV430 is the standard base64 (as defined in RFC 4648) encoded
// response of annotating RosesAreRed with the server default annotators by
// Stanford CoreNLP 4.3.0.
//
// It is the same as RosesAreRedRespV410.
const RosesAreRedRespV430 = RosesAreRedRespV410

// RosesAreRedRespV440 is the standard base64 (as defined in RFC 4648) encoded
// response of annotating RosesAreRed with the server default annotators
// by Stanford CoreNLP 4.4.0.
const RosesAreRedRespV440 = `
pSsKRgpSb3NlcyBhcmUgcmVkLgogIFZpb2xldHMgYXJlIGJsdWUuClN1Z2FyIGlz
IHN3ZWV0LgogIEFuZCBzbyBhcmUgeW91LgoS2wgKZQoFUm9zZXMSBE5OUFMaBVJv
c2VzKgEKMgEgOgVSb3Nlc0IBT1IFUm9zZXNYAWAGaAByBFBFUjCIAQCQAQGoAQCw
AgDyAwFP+gMBT4AEAJIEFE89MC45OTA3MTQzMjI4MTEzODgyClgKA2FyZRIDVkJQ
GgNhcmUqASAyASA6A2FyZUIBT1ICYmVYB2AKaAByBFBFUjCIAQGQAQKoAQCwAgDy
AwFP+gMBT5IEFE89MC45OTk5OTU2NDEyMzAzOTEyClcKA3JlZBICSkoaA3JlZCoB
IDIAOgNyZWRCAU9SA3JlZFgLYA5oAHIEUEVSMIgBApABA6gBALACAPIDAU/6AwFP
kgQUTz0wLjk5OTUwMTU0MDM5MzMwODYKUAoBLhIBLhoBLioAMgMKICA6AS5CAU9S
AS5YDmAPaAByBFBFUjCIAQOQAQSoAQCwAgDyAwFP+gMBT5IEFE89MC45OTk5OTg2
MjA2MzkxMzE5EAAYBCAAKAEwD0JYCgQIABADCgQIABABCgQIABACCgQIABAEEhMI
AxABGgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRITCAMQBBoFcHVu
Y3QgACgAMAA4CRoBA0pYCgQIABADCgQIABABCgQIABACCgQIABAEEhMIAxABGgVu
c3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRITCAMQBBoFcHVuY3QgACgA
MAA4CRoBA1JYCgQIABADCgQIABABCgQIABACCgQIABAEEhMIAxABGgVuc3ViaiAA
KAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRITCAMQBBoFcHVuY3QgACgAMAA4CRoB
A1gBigFYCgQIABADCgQIABABCgQIABACCgQIABAEEhMIAxABGgVuc3ViaiAAKAAw
ADgJEhEIAxACGgNjb3AgACgAMAA4CRITCAMQBBoFcHVuY3QgACgAMAA4CRoBA5IB
WAoECAAQAwoECAAQAQoECAAQAgoECAAQBBITCAMQARoFbnN1YmogACgAMAA4CRIR
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
ARADCgQIARABCgQIARACCgQIARAEEhMIAxABGgVuc3ViaiAAKAAwADgJEhEIAxAC
GgNjb3AgACgAMAA4CRITCAMQBBoFcHVuY3QgACgAMAA4CRoBA0pYCgQIARADCgQI
ARABCgQIARACCgQIARAEEhMIAxABGgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3Ag
ACgAMAA4CRITCAMQBBoFcHVuY3QgACgAMAA4CRoBA1JYCgQIARADCgQIARABCgQI
ARACCgQIARAEEhMIAxABGgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4
CRITCAMQBBoFcHVuY3QgACgAMAA4CRoBA1gCigFYCgQIARADCgQIARABCgQIARAC
CgQIARAEEhMIAxABGgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgAMAA4CRIT
CAMQBBoFcHVuY3QgACgAMAA4CRoBA5IBWAoECAEQAwoECAEQAQoECAEQAgoECAEQ
BBITCAMQARoFbnN1YmogACgAMAA4CRIRCAMQAhoDY29wIAAoADAAOAkSEwgDEAQa
BXB1bmN0IAAoADAAOAkaAQOYAwCwAwDCA9ABCAESB05PTUlOQUwaBlBMVVJBTCIH
VU5LTk9XTioJSU5BTklNQVRFMgdVTktOT1dOOABIAVAAWgd2aW9sZXRzYgFPaP//
/////////wFw////////////AXgBgAEBiAEBkAEAmAECoAEAqAEAsAEAuAEAwAEA
yAEA0AEA2AEB4AEB6AEA8gEGCAAQACAA+gEWCP///////////wEQ////////////
AYICBAgAEACSAwQIABAAkgMECAAQAZIDBAgAEAKSAwQIABADmgMECAAQAMgDAYgE
AaAEARKiCApkCgVTdWdhchIDTk5QGgVTdWdhcioBCjIBIDoFU3VnYXJCAU9SBVN1
Z2FyWCRgKWgAcgRQRVIwiAEIkAEJqAEAsAIA8gMBT/oDAU+ABAKSBBRPPTAuOTg5
Mzk0NDA4NTA5MjMyOApVCgJpcxIDVkJaGgJpcyoBIDIBIDoCaXNCAU9SAmJlWCpg
LGgAcgRQRVIwiAEJkAEKqAEAsAIA8gMBT/oDAU+SBBRPPTAuOTk5OTk4MTUzNDEy
NTI2NQpfCgVzd2VldBICSkoaBXN3ZWV0KgEgMgA6BXN3ZWV0QgFPUgVzd2VldFgt
YDJoAHIEUEVSMIgBCpABC6gBALACAPIDAU/6AwFPkgQUTz0wLjk5OTk3MDIxMDc4
Mzk4MDEKUAoBLhIBLhoBLioAMgMKICA6AS5CAU9SAS5YMmAzaAByBFBFUjCIAQuQ
AQyoAQCwAgDyAwFP+gMBT5IEFE89MC45OTk5OTk3OTM4OTc3NjQyEAgYDCACKCQw
M0JYCgQIAhADCgQIAhABCgQIAhACCgQIAhAEEhMIAxABGgVuc3ViaiAAKAAwADgJ
EhEIAxACGgNjb3AgACgAMAA4CRITCAMQBBoFcHVuY3QgACgAMAA4CRoBA0pYCgQI
AhADCgQIAhABCgQIAhACCgQIAhAEEhMIAxABGgVuc3ViaiAAKAAwADgJEhEIAxAC
GgNjb3AgACgAMAA4CRITCAMQBBoFcHVuY3QgACgAMAA4CRoBA1JYCgQIAhADCgQI
AhABCgQIAhACCgQIAhAEEhMIAxABGgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3Ag
ACgAMAA4CRITCAMQBBoFcHVuY3QgACgAMAA4CRoBA1gCigFYCgQIAhADCgQIAhAB
CgQIAhACCgQIAhAEEhMIAxABGgVuc3ViaiAAKAAwADgJEhEIAxACGgNjb3AgACgA
MAA4CRITCAMQBBoFcHVuY3QgACgAMAA4CRoBA5IBWAoECAIQAwoECAIQAQoECAIQ
AgoECAIQBBITCAMQARoFbnN1YmogACgAMAA4CRIRCAMQAhoDY29wIAAoADAAOAkS
EwgDEAQaBXB1bmN0IAAoADAAOAkaAQOYAwCwAwDCA88BCAISBlBST1BFUhoIU0lO
R1VMQVIiB05FVVRSQUwqCUlOQU5JTUFURTIHVU5LTk9XTjgASAFQAFoFc3VnYXJi
AU9o////////////AXD///////////8BeAKAAQKIAQKQAQCYAQKgAQCoAQCwAQC4
AQDAAQDIAQDQAQDYAQHgAQHoAQDyAQYIARAAIAD6ARYI////////////ARD/////
//////8BggIECAEQAJIDBAgBEACSAwQIARABkgMECAEQApIDBAgBEAOaAwQIARAA
yAMBiAQBoAQBEuoJCloKA0FuZBICQ0MaA0FuZCoDCiAgMgEgOgNBbmRCAU9SA2Fu
ZFg2YDloAHIEUEVSMIgBDJABDagBALACAPIDAU/6AwFPkgQUTz0wLjk5OTk5MDU0
OTQ4OTU0NDYKVAoCc28SAlJCGgJzbyoBIDIBIDoCc29CAU9SAnNvWDpgPGgAcgRQ
RVIwiAENkAEOqAEAsAIA8gMBT/oDAU+SBBRPPTAuOTk5OTk1NTU5NDU3NDAwMgpY
CgNhcmUSA1ZCUBoDYXJlKgEgMgEgOgNhcmVCAU9SAmJlWD1gQGgAcgRQRVIwiAEO
kAEPqAEAsAIA8gMBT/oDAU+SBBRPPTAuOTk5OTk5NTgyMTM2Nzk2NgpbCgN5b3US
A1BSUBoDeW91KgEgMgA6A3lvdUIBT1IDeW91WEFgRGgAcgRQRVIwiAEPkAEQqAEA
sAIA8gMBT/oDAU+ABAOSBBRPPTAuOTk5OTkxMTA3OTU5MDQ3MwpOCgEuEgEuGgEu
KgAyAQo6AS5CAU9SAS5YRGBFaAByBFBFUjCIARCQARGoAQCwAgDyAwFP+gMBT5IE
FE89MC45OTk5OTkxNjA2MTUzNzEyEAwYESADKDYwRUJxCgQIAxAECgQIAxABCgQI
AxACCgQIAxADCgQIAxAFEhAIBBABGgJjYyAAKAAwADgJEhQIBBACGgZhZHZtb2Qg
ACgAMAA4CRIRCAQQAxoDY29wIAAoADAAOAkSEwgEEAUaBXB1bmN0IAAoADAAOAka
AQRKcQoECAMQBAoECAMQAQoECAMQAgoECAMQAwoECAMQBRIQCAQQARoCY2MgACgA
MAA4CRIUCAQQAhoGYWR2bW9kIAAoADAAOAkSEQgEEAMaA2NvcCAAKAAwADgJEhMI
BBAFGgVwdW5jdCAAKAAwADgJGgEEUnEKBAgDEAQKBAgDEAEKBAgDEAIKBAgDEAMK
BAgDEAUSEAgEEAEaAmNjIAAoADAAOAkSFAgEEAIaBmFkdm1vZCAAKAAwADgJEhEI
BBADGgNjb3AgACgAMAA4CRITCAQQBRoFcHVuY3QgACgAMAA4CRoBBFgDigFxCgQI
AxAECgQIAxABCgQIAxACCgQIAxADCgQIAxAFEhAIBBABGgJjYyAAKAAwADgJEhQI
BBACGgZhZHZtb2QgACgAMAA4CRIRCAQQAxoDY29wIAAoADAAOAkSEwgEEAUaBXB1
bmN0IAAoADAAOAkaAQSSAXEKBAgDEAQKBAgDEAEKBAgDEAIKBAgDEAMKBAgDEAUS
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

// RosesAreRedRespV450 is the standard base64 (as defined in RFC 4648) encoded
// response of annotating RosesAreRed with the server default annotators by
// Stanford CoreNLP 4.5.0.
//
// It is the same as RosesAreRedRespV440.
const RosesAreRedRespV450 = RosesAreRedRespV440

// RosesAreRedRespV452 is the standard base64 (as defined in RFC 4648) encoded
// response of annotating RosesAreRed with the server default annotators by
// Stanford CoreNLP 4.5.2.
//
// It is the same as RosesAreRedRespV440.
const RosesAreRedRespV452 = RosesAreRedRespV440

// RosesAreRedRespV453 is the standard base64 (as defined in RFC 4648) encoded
// response of annotating RosesAreRed with the server default annotators by
// Stanford CoreNLP 4.5.3.
//
// It is the same as RosesAreRedRespV440.
const RosesAreRedRespV453 = RosesAreRedRespV440

// RosesAreRedRespV455 is the standard base64 (as defined in RFC 4648) encoded
// response of annotating RosesAreRed with the server default annotators by
// Stanford CoreNLP 4.5.5.
//
// It is the same as RosesAreRedRespV440.
const RosesAreRedRespV455 = RosesAreRedRespV440
