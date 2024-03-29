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

package model_test

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"testing"

	"github.com/donyori/gocorenlp/internal/pbtest"
	"github.com/donyori/gocorenlp/model"
	"github.com/donyori/gocorenlp/model/v4.5.6-eb50467fa8e3/pb"
)

func TestResponseBodyDecoder_OneResponse(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(RosesResp)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	dec := model.NewResponseBodyDecoder(bytes.NewReader(respBody))
	if dec == nil {
		t.Fatal("got nil decoder")
	}
	doc := new(pb.Document)
	err = dec.Decode(doc)
	if err != nil {
		t.Fatal(err)
	}
	err = pbtest.CheckRosesAreRedDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
	err = dec.Decode(new(pb.Document))
	if !errors.Is(err, io.EOF) {
		t.Errorf("got %v; want io.EOF", err)
	}
}

func TestResponseBodyDecoder_TwoResponses(t *testing.T) {
	respBody, err := base64.StdEncoding.DecodeString(RosesResp)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}
	b := make([]byte, len(respBody)*2)
	copy(b[copy(b, respBody):], respBody)
	dec := model.NewResponseBodyDecoder(bytes.NewReader(b))
	if dec == nil {
		t.Fatal("got nil decoder")
	}
	for i := range 2 {
		doc := new(pb.Document)
		err = dec.Decode(doc)
		if err != nil {
			t.Fatalf("decode Response %d: %v", i+1, err)
		}
		err = pbtest.CheckRosesAreRedDocument(doc)
		if err != nil {
			t.Fatalf("decode Response %d: %v", i+1, err)
		}
	}
	err = dec.Decode(new(pb.Document))
	if !errors.Is(err, io.EOF) {
		t.Errorf("got %v; want io.EOF", err)
	}
}

func TestResponseBodyDecoder_DifferentResponses(t *testing.T) {
	const NumRepeat int = 3
	data, _, err := MakeDifferentResponsesData(NumRepeat)
	if err != nil {
		t.Fatal("failed to decode standard base64 encoded response:", err)
	}

	dec := model.NewResponseBodyDecoder(bytes.NewReader(data))
	if dec == nil {
		t.Fatal("got nil decoder")
	}
	for i := range NumRepeat {
		doc := new(pb.Document)
		err = dec.Decode(doc)
		if err != nil {
			t.Fatalf("Round %d: decode Response 1: %v", i+1, err)
		}
		err = pbtest.CheckRosesAreRedDocument(doc)
		if err != nil {
			t.Fatalf("Round %d: decode Response 1: %v", i+1, err)
		}

		doc = new(pb.Document)
		err = dec.Decode(doc)
		if err != nil {
			t.Fatalf("Round %d: decode Response 2: %v", i+1, err)
		} else if text := doc.GetText(); text != YesterdayIsHistory {
			t.Fatalf("Round %d: decode Response 2: got text %q; want %q",
				i+1, text, YesterdayIsHistory)
		}

		doc = new(pb.Document)
		err = dec.Decode(doc)
		if err != nil {
			t.Fatalf("Round %d: decode Response 3: %v", i+1, err)
		}
		err = pbtest.CheckRosesAreRedDocument(doc)
		if err != nil {
			t.Fatalf("Round %d: decode Response 3: %v", i+1, err)
		}

		doc = new(pb.Document)
		err = dec.Decode(doc)
		if err != nil {
			t.Fatalf("Round %d: decode Response 4: %v", i+1, err)
		} else if text := doc.GetText(); text != YesterdayIsHistory {
			t.Fatalf("Round %d: decode Response 4: got text %q; want %q",
				i+1, text, YesterdayIsHistory)
		}
	}
	err = dec.Decode(new(pb.Document))
	if !errors.Is(err, io.EOF) {
		t.Errorf("got %v; want io.EOF", err)
	}
}

func TestResponseBodyDecoder_NilReader(t *testing.T) {
	dec := model.NewResponseBodyDecoder(nil)
	if dec == nil {
		t.Fatal("got nil decoder")
	}
	err := dec.Decode(new(pb.Document))
	if !errors.Is(err, io.EOF) {
		t.Errorf("got %v; want io.EOF", err)
	}
}
