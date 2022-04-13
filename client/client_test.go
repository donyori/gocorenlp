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

package client_test

import (
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/donyori/gocorenlp/client"
)

const (
	DefaultIdx = iota
	DiffStatusIndex
	UserIndex

	N
)

const (
	DefaultPortIndex = iota
	DiffStatusMainPortIndex
	DiffStatusStatusPortIndex
	UserPortIndex

	PortNum
)

var SubtestNames = [N]string{"default", "diff status", "user"}

var ServerPorts = [PortNum]uint16{9000, 9100, 9101, 9200}

const (
	Username = "user1"
	Password = "u1%passWORD"
)

func TestNew(t *testing.T) {
	const skipFormat = "server 127.0.0.1:%d is offline; skip this test"
	var isServersOnline [PortNum]bool
	for i := 0; i < PortNum; i++ {
		isServersOnline[i] = IsServerListeningOnPort(ServerPorts[i])
	}
	for i, name := range SubtestNames {
		t.Run(name, func(t *testing.T) {
			var c client.Client
			var err error
			switch i {
			case DefaultIdx:
				if !isServersOnline[DefaultPortIndex] {
					t.Skipf(skipFormat, ServerPorts[DefaultPortIndex])
				}
				c, err = client.New(nil)
			case DiffStatusIndex:
				mainPort, statusPort := ServerPorts[DiffStatusMainPortIndex], ServerPorts[DiffStatusStatusPortIndex]
				if !isServersOnline[DiffStatusMainPortIndex] {
					t.Skipf(skipFormat, mainPort)
				}
				if !isServersOnline[DiffStatusStatusPortIndex] {
					t.Skipf(skipFormat, statusPort)
				}
				c, err = client.New(&client.Options{
					Port:       mainPort,
					StatusPort: statusPort,
				})
			case UserIndex:
				port := ServerPorts[UserPortIndex]
				if !isServersOnline[UserPortIndex] {
					t.Skipf(skipFormat, port)
				}
				c, err = client.New(&client.Options{
					Port:     port,
					Username: Username,
					Password: Password,
				})
			default:
				// This case should never happen.
				t.Fatalf("%d is out of range [0-%d]", i, N)
			}
			if err != nil {
				t.Fatal(err)
			}
			if c == nil {
				t.Error("got nil client")
			}
		})
	}
}

// IsServerListeningOnPort checks whether a local server
// is listening on the specified port.
func IsServerListeningOnPort(port uint16) bool {
	conn, err := net.DialTimeout(
		"tcp",
		"127.0.0.1:"+strconv.FormatUint(uint64(port), 10),
		time.Millisecond*10,
	)
	if err != nil {
		return false
	}
	_ = conn.Close() // ignore error
	return true
}
