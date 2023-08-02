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

package model_test

import (
	"encoding/base64"
	"io"
	"testing"

	"google.golang.org/protobuf/encoding/protowire"

	"github.com/donyori/gocorenlp/errors"
	"github.com/donyori/gocorenlp/internal/pbtest"
	"github.com/donyori/gocorenlp/model"
	"github.com/donyori/gocorenlp/model/v4.5.3-5250f9faf9f1/pb"
)

func TestDecodeMessage(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(pbtest.RosesAreRedRespV453)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	msgData, n := protowire.ConsumeBytes(respBody)
	if n < 0 {
		t.Fatal("failed to decode response body:", protowire.ParseError(n))
	}
	t.Run("basic", func(t *testing.T) {
		b := make([]byte, len(msgData))
		copy(b, msgData)
		doc := new(pb.Document)
		err := model.DecodeMessage(b, doc)
		if err != nil {
			t.Fatal(err)
		}
		err = pbtest.CheckRosesAreRedDocument(doc)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("append one zero byte", func(t *testing.T) {
		b := make([]byte, len(msgData)+1)
		copy(b, msgData)
		doc := new(pb.Document)
		err := model.DecodeMessage(b, doc)
		if !errors.IsProtoBufError(err) {
			t.Errorf("got %v; want a *ProtoBufError", err)
		}
	})
	t.Run("prepend one zero byte", func(t *testing.T) {
		b := make([]byte, 1+len(msgData))
		copy(b[1:], msgData)
		doc := new(pb.Document)
		err := model.DecodeMessage(b, doc)
		if !errors.IsProtoBufError(err) {
			t.Errorf("got %v; want a *ProtoBufError", err)
		}
	})
	t.Run("truncated", func(t *testing.T) {
		b := make([]byte, len(msgData)/2)
		copy(b, msgData)
		doc := new(pb.Document)
		err := model.DecodeMessage(b, doc)
		if !errors.IsProtoBufError(err) {
			t.Errorf("got %v; want a *ProtoBufError", err)
		}
	})
}

func TestDecodeResponseBody(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(pbtest.RosesAreRedRespV453)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	t.Run("basic", func(t *testing.T) {
		b := make([]byte, len(respBody))
		copy(b, respBody)
		doc := new(pb.Document)
		err := model.DecodeResponseBody(b, doc)
		if err != nil {
			t.Fatal(err)
		}
		err = pbtest.CheckRosesAreRedDocument(doc)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("append suffix", func(t *testing.T) {
		const Suffix = "Go client for Stanford CoreNLP server"
		b := make([]byte, len(respBody), len(respBody)+len(Suffix))
		copy(b, respBody)
		b = append(b, Suffix...)
		doc := new(pb.Document)
		err := model.DecodeResponseBody(b, doc)
		if err != nil {
			t.Fatal(err)
		}
		err = pbtest.CheckRosesAreRedDocument(doc)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("prepend one zero byte", func(t *testing.T) {
		b := make([]byte, 1+len(respBody))
		copy(b[1:], respBody)
		doc := new(pb.Document)
		err := model.DecodeResponseBody(b, doc)
		if !errors.IsProtoBufError(err) {
			t.Errorf("got %v; want a *ProtoBufError", err)
		}
	})
	t.Run("truncated", func(t *testing.T) {
		b := make([]byte, len(respBody)/2)
		copy(b, respBody)
		doc := new(pb.Document)
		err := model.DecodeResponseBody(b, doc)
		if !errors.IsProtoBufError(err) ||
			!errors.Is(err, io.ErrUnexpectedEOF) {
			t.Errorf("got %v; want a *ProtoBufError wrapping io.ErrUnexpectedEOF", err)
		}
	})
}

func TestConsumeResponseBody(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(pbtest.RosesAreRedRespV453)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	t.Run("basic", func(t *testing.T) {
		b := make([]byte, len(respBody))
		copy(b, respBody)
		doc := new(pb.Document)
		n, err := model.ConsumeResponseBody(b, doc)
		if err != nil {
			t.Fatal(err)
		} else if n != len(respBody) {
			t.Errorf("got n %d; want %d", n, len(respBody))
		}
		err = pbtest.CheckRosesAreRedDocument(doc)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("append suffix", func(t *testing.T) {
		const Suffix = "Go client for Stanford CoreNLP server"
		b := make([]byte, len(respBody), len(respBody)+len(Suffix))
		copy(b, respBody)
		b = append(b, Suffix...)
		doc := new(pb.Document)
		n, err := model.ConsumeResponseBody(b, doc)
		if err != nil {
			t.Fatal(err)
		} else if n != len(respBody) {
			t.Errorf("got n %d; want %d", n, len(respBody))
		}
		err = pbtest.CheckRosesAreRedDocument(doc)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("prepend one zero byte", func(t *testing.T) {
		b := make([]byte, 1+len(respBody))
		copy(b[1:], respBody)
		doc := new(pb.Document)
		n, err := model.ConsumeResponseBody(b, doc)
		// The zero byte is treated as the prefixed length,
		// representing a length of 0.
		// model.ConsumeResponseBody reads this byte (length: 1)
		// and then decodes an empty response body.
		// Therefore, n should be 1 here.
		if n != 1 {
			t.Errorf("got %d; want 1", n)
		}
		if !errors.IsProtoBufError(err) {
			t.Errorf("got %v; want a *ProtoBufError", err)
		}
	})
	t.Run("truncated", func(t *testing.T) {
		b := make([]byte, len(respBody)/2)
		copy(b, respBody)
		doc := new(pb.Document)
		n, err := model.ConsumeResponseBody(b, doc)
		// n should be -1 here to represent truncated data,
		// corresponding to io.ErrUnexpectedEOF.
		if n != -1 {
			t.Errorf("got %d; want -1", n)
		}
		if !errors.IsProtoBufError(err) ||
			!errors.Is(err, io.ErrUnexpectedEOF) {
			t.Errorf("got %v; want a *ProtoBufError wrapping io.ErrUnexpectedEOF", err)
		}
	})
	t.Run("two responses", func(t *testing.T) {
		b := make([]byte, len(respBody)*2)
		copy(b[copy(b, respBody):], respBody)
		doc := new(pb.Document)
		n1, err := model.ConsumeResponseBody(b, doc)
		if err != nil {
			t.Fatal(err)
		} else if n1 != len(respBody) {
			t.Errorf("got n %d; want %d", n1, len(respBody))
		}
		err = pbtest.CheckRosesAreRedDocument(doc)
		if err != nil {
			t.Error(err)
		}
		if n1 != len(respBody) {
			return
		}
		doc = new(pb.Document)
		n2, err := model.ConsumeResponseBody(b[n1:], doc)
		if err != nil {
			t.Fatal(err)
		} else if n2 != len(respBody) {
			t.Errorf("got n %d; want %d", n2, len(respBody))
		}
		err = pbtest.CheckRosesAreRedDocument(doc)
		if err != nil {
			t.Error(err)
		}
	})
}
