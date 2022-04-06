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

// Package model provides structures for CoreNLP and
// some functions to manipulate these structures.
//
// To fit different versions of CoreNLP,
// the structures are organized into subpackages named in the form:
//  github.com/donyori/gocorenlp/model/vX.Y.Z-abcdefabcdef/pb
// where vX.Y.Z is the version of CoreNLP,
// abcdefabcdef is a 12-character prefix of the commit hash of
// the retrieved .proto file in the Stanford CoreNLP project,
// and pb stands for Protocol Buffers (Protobuf).
//
// The .proto files are edited from "edu.stanford.nlp.pipeline.CoreNLP.proto"
// in the Stanford CoreNLP project, which can be retrieved from GitHub:
// <https://github.com/stanfordnlp/CoreNLP/blob/main/src/edu/stanford/nlp/pipeline/CoreNLP.proto>.
//
// The .pb.go files are auto-generated according to
// the corresponding .proto files. For more information,
// see <https://developers.google.com/protocol-buffers/docs/gotutorial>.
package model

// The following go:generate directives are for compiling
// all the "corenlp.proto" in its subpackages.
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
// After running these go:generate directives,
// files named "corenlp.pb.go" should be generated in its subpackages.
// Each "corenlp.pb.go" and its corresponding "corenlp.proto"
// should be in the same directory.

//go:generate protoc --proto_path=v3.6.0-29765338a2e8/pb --go_out=. --go_opt=module=github.com/donyori/gocorenlp/model corenlp.proto
//go:generate protoc --proto_path=v4.4.0-e90f30f13c40/pb --go_out=. --go_opt=module=github.com/donyori/gocorenlp/model corenlp.proto
