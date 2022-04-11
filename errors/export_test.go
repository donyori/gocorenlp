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

package errors

import "testing"

// TestExportStdErrors is trivial.
// It ensures that none of the four functions exported from
// the standard package errors are missing.
func TestExportStdErrors(t *testing.T) {
	errorMsg := "probably not exported from standard package errors"
	err1 := New("test error 1")
	if err1 == nil {
		t.Fatal(`New("test error 1") - got nil;`, errorMsg)
	}
	err2 := &testError{Msg: "test error 2"}
	if err := Unwrap(err1); err != nil {
		t.Errorf("Unwrap(err1) - got %v; %s", err, errorMsg)
	}
	if Is(err1, err2) {
		t.Error("Is(err1, err2) - got true;", errorMsg)
	}
	if As(err1, &err2) {
		t.Error("As(err1, err2) - got true;", errorMsg)
	}
}

type testError struct {
	Msg string
}

func (e *testError) Error() string {
	if e == nil {
		return ""
	}
	return e.Msg
}
