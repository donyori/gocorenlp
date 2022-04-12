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

// The following tests cannot run at the same time!

const testShutdownPort uint16 = 9300
const testServerId string = "test_server"

// Run the following test with a Stanford CoreNLP 4.4.0 server running and
// (both the main server and the status server) listening on 127.0.0.1:9300.
//
// Set testEnableShutdownLocalBasicTest to true
// when you are ready to run the following test.

const testEnableShutdownLocalBasicTest = false

func TestClientImpl_ShutdownLocal_Basic(t *testing.T) {
	if !testEnableShutdownLocalBasicTest {
		return
	}
	c := newClientImpl(&Options{Port: testShutdownPort})
	if err := c.ShutdownLocal(); err != nil {
		testLogErrorChain(t, err)
	}
}

// Run the following test with a Stanford CoreNLP 4.4.0 server running,
// (both the main server and the status server) listening on 127.0.0.1:9300,
// and starting with -server_id=test_Server.
//
// Set testEnableShutdownLocalServerIdTest to true
// when you are ready to run the following test.

const testEnableShutdownLocalServerIdTest = false

func TestClientImpl_ShutdownLocal_ServerId(t *testing.T) {
	if !testEnableShutdownLocalServerIdTest {
		return
	}
	c := newClientImpl(&Options{
		Port:     testShutdownPort,
		ServerId: testServerId,
	})
	if err := c.ShutdownLocal(); err != nil {
		testLogErrorChain(t, err)
	}
}

// Run the following test with a Stanford CoreNLP 4.4.0 server running,
// with the main server listening on 127.0.0.1:9300 and
// the status server listening on 127.0.0.1:9301.
//
// Set testEnableShutdownLocalDiffStatusPortTest to true
// when you are ready to run the following test.

const testEnableShutdownLocalDiffStatusPortTest = false

func TestClientImpl_ShutdownLocal_DiffStatusPort(t *testing.T) {
	if !testEnableShutdownLocalDiffStatusPortTest {
		return
	}
	c := newClientImpl(&Options{
		Port:       testShutdownPort,
		StatusPort: testDiffStatusPortPort + 1,
	})
	if err := c.ShutdownLocal(); err != nil {
		testLogErrorChain(t, err)
	}
}
