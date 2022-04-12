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

// Run the following tests with a Stanford CoreNLP 4.4.0 server running and
// (both the main server and the status server) listening on 127.0.0.1:9000.
// The server should use its default language model.

func TestNew_Basic(t *testing.T) {
	c, err := New(nil)
	if err != nil {
		testLogErrorChain(t, err)
		return
	}
	if c == nil {
		t.Error("got nil client")
	}
}

func TestClientImpl_Live_Basic(t *testing.T) {
	c := testNewBasicClientImpl()
	if err := c.Live(); err != nil {
		testLogErrorChain(t, err)
	}
}

func TestClientImpl_Ready_Basic(t *testing.T) {
	c := testNewBasicClientImpl()
	if err := c.Ready(); err != nil {
		testLogErrorChain(t, err)
	}
}

func TestClientImpl_Annotate_Basic(t *testing.T) {
	testAnnotateFunc(t, testNewBasicClientImpl)
}

func TestClientImpl_AnnotateString_Basic(t *testing.T) {
	testAnnotateStringFunc(t, testNewBasicClientImpl)
}

func TestClientImpl_AnnotateRaw_Basic(t *testing.T) {
	testAnnotateRawFunc(t, testNewBasicClientImpl)
}

// testNewBasicClientImpl creates a *clientImpl
// connecting to 127.0.0.1:9000,
// with no userinfo, no timeout,
// annotators="tokenize,ssplit,pos",
// and contentType="application/x-www-form-urlencoded; charset=utf-8".
func testNewBasicClientImpl() *clientImpl {
	return newClientImpl(&Options{Annotators: "tokenize,ssplit,pos"})
}
