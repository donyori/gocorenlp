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

import (
	"net/netip"
	"strconv"
	"strings"
	"time"
)

// Options are the configuration for creating a new client.
type Options struct {
	// Hostname is the host (without port number) of the target server.
	//
	// Default: 127.0.0.1
	Hostname string `json:"hostname,omitempty"`

	// Port is the port of the target server.
	//
	// Default: 9000
	Port uint16 `json:"port,omitempty"`

	// StatusPort is the port of the target server to run
	// the liveness and readiness server on.
	// If zero, treat it the same as the main server.
	//
	// Default: 0
	StatusPort uint16 `json:"statusPort,omitempty"`

	// Timeout specifies a time limit for requests made by the client.
	// The timeout includes connection time, any redirects,
	// and reading the response body.
	//
	// A non-positive value means no timeout.
	//
	// Default: 0
	Timeout time.Duration `json:"timeout,omitempty"`

	// Username is the username sent with the request.
	// Set this along with Password if the target server requires basic auth.
	//
	// Default: "" (empty)
	Username string `json:"username,omitempty"`

	// Password is the password sent with the request.
	// Set this along with Username if the target server requires basic auth.
	//
	// Only valid when Username is not empty.
	//
	// Default: "" (empty)
	Password string `json:"password,omitempty"`

	// Charset is the character encoding of the request
	// set in the Content-Type header.
	//
	// Default: utf-8
	Charset string `json:"charset,omitempty"`

	// Annotators are the default annotators with the annotation request.
	// If no annotators are specified with the annotation request,
	// these annotators will be used.
	//
	// The annotators are separated by commas (,) in the string without spaces.
	// For example:
	//  tokenize,ssplit,pos,depparse
	//
	// Default: "" (empty, no annotator is specified by default)
	Annotators string `json:"annotators,omitempty"`

	// ServerId is the value of the option -server_id used
	// when starting the target server.
	//
	// If the server is started without that option, leave it empty.
	//
	// Default: "" (empty)
	ServerId string `json:"serverId,omitempty"`

	// onlyKeyedLiterals forces others to construct Options
	// only with the keyed literals, so future additions to it
	// will not violate compatibility.
	onlyKeyedLiterals struct{}
}

// makeHosts generates the hosts (including the hostname part and
// the port number part) for the main server and status server
// from the specified hostname, port, and statusPort.
func makeHosts(hostname string, port, statusPort uint16) (host, statusHost string) {
	hostname = strings.TrimSpace(hostname)
	if len(hostname) == 0 {
		hostname = "127.0.0.1"
	}
	if port == 0 {
		port = 9000
	}
	if statusPort == 0 {
		statusPort = port
	}
	if addr, err := netip.ParseAddr(hostname); err == nil {
		// hostname is an IP address.
		host = netip.AddrPortFrom(addr, port).String()
		if statusPort == port {
			statusHost = host
		} else {
			statusHost = netip.AddrPortFrom(addr, statusPort).String()
		}
	} else {
		// hostname is not an IP address, may be a domain name, or invalid.
		// This function does not validate the host.
		// So simply join the hostname and port.
		host = hostname + ":" + strconv.FormatUint(uint64(port), 10)
		if statusPort == port {
			statusHost = host
		} else {
			statusHost = hostname + ":" + strconv.FormatUint(uint64(statusPort), 10)
		}
	}
	return
}
