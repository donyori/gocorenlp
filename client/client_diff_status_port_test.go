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

package client

import "testing"

// Run the following tests with a Stanford CoreNLP 4.4.0 server running,
// with the main server listening on 127.0.0.1:9100 and
// the status server listening on 127.0.0.1:9101.
// The server should use its default language model.
//
// Set testEnableDiffStatusTest to true
// when you are ready to run the following tests.

const testEnableDiffStatusTest = false

const testDiffStatusPortPort uint16 = 9100

func TestNew_DiffStatusPort(t *testing.T) {
	if !testEnableDiffStatusTest {
		return
	}
	c, err := New(&Options{
		Port:       testDiffStatusPortPort,
		StatusPort: testDiffStatusPortPort + 1,
	})
	if err != nil {
		testLogErrorChain(t, err)
		return
	}
	if c == nil {
		t.Error("got nil client")
	}
}

func TestClientImpl_Live_DiffStatusPort(t *testing.T) {
	if !testEnableDiffStatusTest {
		return
	}
	c := testNewDiffStatusPortClientImpl()
	if err := c.Live(); err != nil {
		testLogErrorChain(t, err)
	}
}

func TestClientImpl_Ready_DiffStatusPort(t *testing.T) {
	if !testEnableDiffStatusTest {
		return
	}
	c := testNewDiffStatusPortClientImpl()
	if err := c.Ready(); err != nil {
		testLogErrorChain(t, err)
	}
}

func TestClientImpl_Annotate_DiffStatusPort(t *testing.T) {
	if !testEnableDiffStatusTest {
		return
	}
	testAnnotateFunc(t, testNewDiffStatusPortClientImpl)
}

func TestClientImpl_AnnotateString_DiffStatusPort(t *testing.T) {
	if !testEnableDiffStatusTest {
		return
	}
	testAnnotateStringFunc(t, testNewDiffStatusPortClientImpl)
}

func TestClientImpl_AnnotateRaw_DiffStatusPort(t *testing.T) {
	if !testEnableDiffStatusTest {
		return
	}
	testAnnotateRawFunc(t, testNewDiffStatusPortClientImpl)
}

// testNewDiffStatusPortClientImpl creates a *clientImpl
// connecting to the main server on 127.0.0.1:9100,
// the status server on 127.0.0.1:9101,
// with no userinfo, no timeout,
// annotators="tokenize,ssplit,pos",
// and contentType="application/x-www-form-urlencoded; charset=utf-8".
func testNewDiffStatusPortClientImpl() *clientImpl {
	return newClientImpl(&Options{
		Port:       testDiffStatusPortPort,
		StatusPort: testDiffStatusPortPort + 1,
		Annotators: "tokenize,ssplit,pos",
	})
}
