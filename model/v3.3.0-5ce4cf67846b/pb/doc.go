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

// Package pb provides auto-generated structures for
// the data set of Stanford CoreNLP 3.3.0.
//
// The corresponding commit hash is 5ce4cf67846b9b53aef87825c6fcb8bca608b01a.
package pb

// The following go:generate directive is for compiling the file
// "corenlp_v3_3_0_5ce4cf67846b.proto" in this directory.
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
// After running this go:generate directive, a file named
// "corenlp_v3_3_0_5ce4cf67846b.pb.go" should be generated in this directory.
//
// This command uses the "module=" output mode to verify that the "go_package"
// defined in "corenlp_v3_3_0_5ce4cf67846b.proto" is appropriate.
// But the "paths=source_relative" output mode can make the command
// more concise, as follows:
//  protoc --go_out=. --go_opt=paths=source_relative corenlp_v3_3_0_5ce4cf67846b.proto
//
// Note that do not use any filename patterns or path patterns (e.g., *.proto).
// Patterns may cause errors or unexpected behavior on different platforms.
//
//go:generate protoc --go_out=. --go_opt=module=github.com/donyori/gocorenlp/model/v3.3.0-5ce4cf67846b/pb corenlp_v3_3_0_5ce4cf67846b.proto
