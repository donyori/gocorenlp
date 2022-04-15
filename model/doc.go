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

// Package model defines structures for Stanford CoreNLP in its subpackages
// and provides functions to parse them from ProtoBuf wire encoding.
//
// To fit different versions of CoreNLP,
// the structures are organized into subpackages named in the form:
//  github.com/donyori/gocorenlp/model/vX.Y.Z-abcdefabcdef/pb
// where X.Y.Z is the version of CoreNLP,
// abcdefabcdef is a 12-character prefix of the commit hash of
// the retrieved .proto file in the Stanford CoreNLP project,
// and pb stands for Protocol Buffers (ProtoBuf).
//
// The .proto files are edited from "edu.stanford.nlp.pipeline.CoreNLP.proto"
// in the Stanford CoreNLP project, which can be retrieved from GitHub:
// <https://github.com/stanfordnlp/CoreNLP/blob/main/src/edu/stanford/nlp/pipeline/CoreNLP.proto>.
//
// Each .proto file has a unique name to avoid ProtoBuf file name conflicts.
// (This conflict might be a bug of protoc-gen-go as the files are in different
// packages and in different directories, but the conflict still exists.)
// Although this naming style seems verbose and redundant, it works.
//
// The .pb.go files are auto-generated according to
// the corresponding .proto files. For more information,
// see <https://developers.google.com/protocol-buffers/docs/gotutorial>.
package model

// The following go:generate directives are for compiling all the
// CoreNLP ProtoBuf files (named like "corenlp_*.proto") in its subpackages.
//
// Before running it, make sure that the ProtoBuf compiler "protoc" and
// its Go plugin "protoc-gen-go" are installed and available in $PATH.
//
// To install "protoc", see <https://developers.google.com/protocol-buffers/docs/downloads>
// and follow the instructions in the README.
//
// To install "protoc-gen-go", run the following command:
//  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//
// After running these go:generate directives, files named like
// "corenlp_*.pb.go" should be generated in its subpackages.
// Each "corenlp_*.pb.go" and its corresponding "corenlp_*.proto"
// should be in the same directory.
//
// Note that do not use any filename patterns or path patterns (e.g., *.proto).
// Patterns may cause errors or unexpected behavior on different platforms.
//
//go:generate protoc --proto_path=v3.6.0-29765338a2e8/pb --go_out=. --go_opt=module=github.com/donyori/gocorenlp/model corenlp_v3_6_0_29765338a2e8.proto
//go:generate protoc --proto_path=v4.0.0-2b3dd38abe00/pb --go_out=. --go_opt=module=github.com/donyori/gocorenlp/model corenlp_v4_0_0_2b3dd38abe00.proto
//go:generate protoc --proto_path=v4.1.0-a1427196ba6e/pb --go_out=. --go_opt=module=github.com/donyori/gocorenlp/model corenlp_v4_1_0_a1427196ba6e.proto
//go:generate protoc --proto_path=v4.2.0-3ad83fc2e42e/pb --go_out=. --go_opt=module=github.com/donyori/gocorenlp/model corenlp_v4_2_0_3ad83fc2e42e.proto
//go:generate protoc --proto_path=v4.2.1-d8d09b2c81a5/pb --go_out=. --go_opt=module=github.com/donyori/gocorenlp/model corenlp_v4_2_1_d8d09b2c81a5.proto
//go:generate protoc --proto_path=v4.3.0-f885cd198767/pb --go_out=. --go_opt=module=github.com/donyori/gocorenlp/model corenlp_v4_3_0_f885cd198767.proto
//go:generate protoc --proto_path=v4.4.0-e90f30f13c40/pb --go_out=. --go_opt=module=github.com/donyori/gocorenlp/model corenlp_v4_4_0_e90f30f13c40.proto
