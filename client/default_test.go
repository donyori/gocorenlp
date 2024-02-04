// gocorenlp.  A Go (Golang) client for Stanford CoreNLP server.
// Copyright (C) 2022-2024  Yuan Gao
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
	"bytes"
	"flag"
	"strings"
	"sync"
	"testing"

	"github.com/donyori/gocorenlp/client"
	"github.com/donyori/gocorenlp/errors"
	"github.com/donyori/gocorenlp/model"
	"github.com/donyori/gocorenlp/model/v4.5.6-eb50467fa8e3/pb"
)

const DefaultPort uint16 = 9000
const FlagNameTestFuncShutdown = "testfuncshutdown"

var RunShutdownTest bool

var ParseFlagOnce sync.Once

func init() {
	flag.BoolVar(&RunShutdownTest, FlagNameTestFuncShutdown, false,
		"Set if want to test Shutdown.")
}

func TestLive(t *testing.T) {
	err := client.Live()
	if CheckIsServerListeningOnPort(DefaultPort) {
		if err != nil {
			t.Error(err)
		}
	} else if err == nil {
		t.Error("got nil error but server is offline")
	} else if !errors.IsConnectionError(err) {
		t.Error("got non-nil error but not connection error:", err)
	}
}

func TestReady(t *testing.T) {
	err := client.Ready()
	if CheckIsServerListeningOnPort(DefaultPort) {
		if err != nil {
			t.Error(err)
		}
	} else if err == nil {
		t.Error("got nil error but server is offline")
	} else if !errors.IsConnectionError(err) {
		t.Error("got non-nil error but not connection error:", err)
	}
}

func TestAnnotate(t *testing.T) {
	SkipIfDefaultServerOffline(t)
	AnnotateFunc(t, func(annotators string) *pb.Document {
		doc := new(pb.Document)
		err := client.Annotate(strings.NewReader(Text), annotators, doc)
		if err != nil {
			t.Error(err)
			return nil
		}
		return doc
	})
}

func TestAnnotateString(t *testing.T) {
	SkipIfDefaultServerOffline(t)
	AnnotateFunc(t, func(annotators string) *pb.Document {
		doc := new(pb.Document)
		err := client.AnnotateString(Text, annotators, doc)
		if err != nil {
			t.Error(err)
			return nil
		}
		return doc
	})
}

func TestAnnotateRaw(t *testing.T) {
	SkipIfDefaultServerOffline(t)
	AnnotateFunc(t, func(annotators string) *pb.Document {
		var b bytes.Buffer
		written, err := client.AnnotateRaw(
			strings.NewReader(Text), annotators, &b)
		if err != nil {
			t.Error(err)
			return nil
		}
		if n := int64(b.Len()); written != n {
			t.Errorf("got written %d; want %d", written, n)
			return nil
		}
		doc := new(pb.Document)
		err = model.DecodeResponseBody(b.Bytes(), doc)
		if err != nil {
			t.Error(err)
			return nil
		}
		return doc
	})
}

func TestAnnotateStringRaw(t *testing.T) {
	SkipIfDefaultServerOffline(t)
	AnnotateFunc(t, func(annotators string) *pb.Document {
		var b bytes.Buffer
		written, err := client.AnnotateStringRaw(Text, annotators, &b)
		if err != nil {
			t.Error(err)
			return nil
		}
		if n := int64(b.Len()); written != n {
			t.Errorf("got written %d; want %d", written, n)
			return nil
		}
		doc := new(pb.Document)
		err = model.DecodeResponseBody(b.Bytes(), doc)
		if err != nil {
			t.Error(err)
			return nil
		}
		return doc
	})
}

func TestShutdown(t *testing.T) {
	ParseFlag()
	if !RunShutdownTest {
		t.Skip("skip this test because flag -" + FlagNameTestFuncShutdown +
			" is not set")
	}
	SkipIfDefaultServerOffline(t)
	if err := client.Shutdown(); err != nil {
		t.Error(err)
	}
}

// ParseFlag calls flag.Parse to parse the command line flags.
//
// It only takes effect once.
func ParseFlag() {
	ParseFlagOnce.Do(func() {
		flag.Parse()
	})
}

// SkipIfDefaultServerOffline skips the test if the default server is offline.
func SkipIfDefaultServerOffline(tb testing.TB) {
	if !CheckIsServerListeningOnPort(DefaultPort) {
		tb.Skipf(SkipLayout, DefaultPort)
	}
}
