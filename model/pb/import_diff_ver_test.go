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

	"github.com/donyori/gocorenlp/model/pb"
)

// This test is to ensure that importing different versions of
// ProtoBuf models at the same time will not cause conflicts.

func TestDifferentVersions(t *testing.T) {
	var _ *pb.Doc360
	var _ *pb.Doc400
	var _ *pb.Doc410
	var _ *pb.Doc420
	var _ *pb.Doc421
	var _ *pb.Doc422
	var _ *pb.Doc430
	var _ *pb.Doc431
	var _ *pb.Doc432
	var _ *pb.Doc440
}
