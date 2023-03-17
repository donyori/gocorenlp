// gocorenlp.  A Go (Golang) client for Stanford CoreNLP server.
// Copyright (C) 2022-2023  Yuan Gao
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
	"github.com/donyori/gocorenlp/model/v4.2.0-3ad83fc2e42e/pb"
)

func TestDecodeBase64Resp(t *testing.T) {
	err := pbtest.CheckDocumentFromBase64(pbtest.RosesAreRedRespV420, new(pb.Document))
	if err != nil {
		t.Error(err)
	}
}
