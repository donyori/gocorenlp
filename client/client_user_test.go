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
// (both the main server and the status server) listening on 127.0.0.1:9200,
// with username="user1" and password="u1%password".

const testUserPort uint16 = 9200
const testUserUsername, testUserPassword string = "user1", "u1%password"

func TestNew_User(t *testing.T) {
	c, err := New(&Options{
		Port:     testUserPort,
		Username: testUserUsername,
		Password: testUserPassword,
	})
	if err != nil {
		t.Fatal(err)
	}
	if c == nil {
		t.Error("got nil client")
	}
}

func TestClientImpl_Live_User(t *testing.T) {
	c := testNewUserClientImpl()
	if err := c.Live(); err != nil {
		t.Error(err)
	}
}

func TestClientImpl_Ready_User(t *testing.T) {
	c := testNewUserClientImpl()
	if err := c.Ready(); err != nil {
		t.Error(err)
	}
}

func TestClientImpl_Annotate_User(t *testing.T) {
	testAnnotateFunc(t, testNewUserClientImpl)
}

func TestClientImpl_AnnotateString_User(t *testing.T) {
	testAnnotateStringFunc(t, testNewUserClientImpl)
}

// testNewUserClientImpl creates a Client
// connecting to 127.0.0.1:9200,
// with username "user1",
// password "u1%password",
// no timeout,
// annotators "tokenize,ssplit,pos",
// and contentType "application/x-www-form-urlencoded; charset=utf-8".
func testNewUserClientImpl() *clientImpl {
	return newClientImpl(&Options{
		Port:       testUserPort,
		Username:   testUserUsername,
		Password:   testUserPassword,
		Annotators: "tokenize,ssplit,pos",
	})
}