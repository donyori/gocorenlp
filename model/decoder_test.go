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
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"testing"

	"github.com/donyori/gocorenlp/internal/pbtest"
	"github.com/donyori/gocorenlp/model"
	"github.com/donyori/gocorenlp/model/v4.5.3-5250f9faf9f1/pb"
)

func TestResponseBodyDecoder(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(pbtest.RosesAreRedRespV453)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	t.Run("one response", func(t *testing.T) {
		b := make([]byte, len(respBody))
		copy(b, respBody)
		dec := model.NewResponseBodyDecoder(bytes.NewReader(b))
		if dec == nil {
			t.Fatal("got nil decoder")
		}
		doc := new(pb.Document)
		err := dec.Decode(doc)
		if err != nil {
			t.Fatal(err)
		}
		err = pbtest.CheckRosesAreRedDocument(doc)
		if err != nil {
			t.Fatal(err)
		}
		err = dec.Decode(doc)
		if !errors.Is(err, io.EOF) {
			t.Errorf("got %v; want io.EOF", err)
		}
	})
	t.Run("two responses", func(t *testing.T) {
		b := make([]byte, len(respBody)*2)
		copy(b[copy(b, respBody):], respBody)
		dec := model.NewResponseBodyDecoder(bytes.NewReader(b))
		if dec == nil {
			t.Fatal("got nil decoder")
		}
		var err error
		for i := 0; i < 2; i++ {
			doc := new(pb.Document)
			err = dec.Decode(doc)
			if err != nil {
				t.Fatalf("Decode %d: %v", i, err)
			}
			err = pbtest.CheckRosesAreRedDocument(doc)
			if err != nil {
				t.Fatalf("Decode %d: %v", i, err)
			}
		}
		err = dec.Decode(new(pb.Document))
		if !errors.Is(err, io.EOF) {
			t.Errorf("got %v; want io.EOF", err)
		}
	})
	t.Run("nil reader", func(t *testing.T) {
		dec := model.NewResponseBodyDecoder(nil)
		if dec == nil {
			t.Fatal("got nil decoder")
		}
		err := dec.Decode(new(pb.Document))
		if !errors.Is(err, io.EOF) {
			t.Errorf("got %v; want io.EOF", err)
		}
	})
}
