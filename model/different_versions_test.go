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

	pb360 "github.com/donyori/gocorenlp/model/v3.6.0-29765338a2e8/pb"
	pb440 "github.com/donyori/gocorenlp/model/v4.4.0-e90f30f13c40/pb"
)

// This test is to ensure that importing ProtoBuf models of
// different versions at the same time will not cause conflicts.

func TestDifferentVersions(t *testing.T) {
	var _ *pb360.Document
	var _ *pb440.Document
}
