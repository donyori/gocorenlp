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

package model

import "google.golang.org/protobuf/proto"

// Document is a type constraint for document structures.
// It is for compatibility with different versions of CoreNLP.
//
// It is defined for our HTTP client.
// The user should not use this interface directly.
// Please use a specific version in the subpackages instead.
type Document interface {
	// Document must be comparable because it is a pointer to a structure.
	comparable
	proto.Message

	// No more constraints to avoid unnecessary conflicts
	// with future versions of CoreNLP.
}
