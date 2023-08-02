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

package model

import (
	"io"
	"sync"

	gogoerrors "github.com/donyori/gogo/errors"
	"github.com/donyori/gogo/inout"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"

	"github.com/donyori/gocorenlp/errors"
)

// ResponseBodyDecoder reads and decodes
// Stanford CoreNLP server responses from an input stream.
type ResponseBodyDecoder interface {
	// Decode parses a ProtoBuf message
	// (usually a Stanford CoreNLP document)
	// from the body of a Stanford CoreNLP server response
	// read from the input stream.
	// The result is stored in msg.
	//
	// The specified message msg must be a non-nil pointer
	// to a ProtoBuf message.
	Decode(msg proto.Message) error

	// private prevents others from implementing this interface,
	// so future additions to it will not violate compatibility.
	private()
}

// responseBodyDecoder is an implementation of
// the interface ResponseBodyDecoder.
type responseBodyDecoder struct {
	r  io.Reader
	br io.ByteReader
}

// NewResponseBodyDecoder creates a new ResponseBodyDecoder that reads from r.
//
// If r does not implement io.ByteReader,
// the ResponseBodyDecoder introduces its own buffering and
// may read data from r beyond the response requested.
func NewResponseBodyDecoder(r io.Reader) ResponseBodyDecoder {
	if r == nil {
		// Replace nil reader with eofReader to avoid panic on reading.
		r = eofReader{}
	}
	if br, ok := r.(io.ByteReader); ok {
		return &responseBodyDecoder{
			r:  r,
			br: br,
		}
	}
	bufReader := inout.NewBufferedReader(r)
	return &responseBodyDecoder{
		r:  bufReader,
		br: bufReader,
	}
}

func (rbd *responseBodyDecoder) Decode(msg proto.Message) error {
	if msg == nil {
		return gogoerrors.AutoNew("the provided message is nil")
	}
	buf := decoderBufferPool.Get().([]byte)
	defer func() {
		// The buffer put to the pool may be different from
		// that gotten from the pool.
		decoderBufferPool.Put(buf)
	}()
	// Read and parse the prefixed length:
	const MaxVarintLen int = 10
	if len(buf) < MaxVarintLen {
		buf = make([]byte, MaxVarintLen)
	}
	var n int
	c := byte(0x80)
	var err error
	for n < MaxVarintLen && c >= 0x80 && err == nil {
		c, err = rbd.br.ReadByte()
		if err != nil && (n == 0 || !errors.Is(err, io.EOF)) {
			return gogoerrors.AutoWrap(err)
		}
		buf[n], n = c, n+1
	}
	sizeBytes := buf[:n]
	size, n := protowire.ConsumeVarint(sizeBytes)
	if n < 0 {
		return gogoerrors.AutoWrap(errors.NewProtoBufError(
			"google.golang.org/protobuf/encoding/protowire.ConsumeVarint",
			sizeBytes,
			protowire.ParseError(n),
		))
	}
	// Read and parse the message body:
	if uint64(len(buf)) < size {
		buf = make([]byte, size)
	}
	_, err = io.ReadFull(rbd.r, buf)
	if err != nil {
		return gogoerrors.AutoWrap(err)
	}
	err = proto.Unmarshal(buf[:size], msg)
	if err != nil {
		return gogoerrors.AutoWrap(errors.NewProtoBufError(
			"google.golang.org/protobuf/proto.Unmarshal",
			msg,
			err,
		))
	}
	return nil
}

func (rbd *responseBodyDecoder) private() {}

// decoderBufferPool is a set of temporary buffers
// to load ProtoBuf data from readers.
var decoderBufferPool = sync.Pool{
	New: func() any {
		return []byte(nil)
	},
}

// eofReader implements io.Reader and io.ByteReader.
// It always reports io.EOF.
type eofReader struct{}

func (r eofReader) Read([]byte) (n int, err error) {
	return 0, io.EOF
}

func (r eofReader) ReadByte() (byte, error) {
	return 0, io.EOF
}
