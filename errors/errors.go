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

package errors

import (
	stderrors "errors"
	"io/fs"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

// IsTimeoutError reports whether the specified error is caused by a timeout.
func IsTimeoutError(err error) bool {
	var t interface{ Timeout() bool }
	return stderrors.As(err, &t) && t.Timeout()
}

// IsFileError reports whether the specified error is caused by
// operating a file or file path.
func IsFileError(err error) bool {
	var e *fs.PathError
	return stderrors.As(err, &e)
}

// IsConnectionError reports whether the specified error is encountered
// during sending an HTTP request and receiving the response.
func IsConnectionError(err error) bool {
	var e *url.Error
	return stderrors.As(err, &e)
}

// IsUnacceptableResponseError reports whether the specified error
// is caused by an unacceptable HTTP response.
func IsUnacceptableResponseError(err error) bool {
	var e *UnacceptableResponseError
	return stderrors.As(err, &e)
}

// IsProtoBufError reports whether the specified error is
// encountered during reading or writing ProtoBuf messages.
func IsProtoBufError(err error) bool {
	var e *ProtoBufError
	return stderrors.As(err, &e)
}

// UnacceptableResponseError records an unacceptable HTTP response,
// including unexpected response status, errors reading the response body,
// and unexpected response body.
type UnacceptableResponseError struct {
	StatusCode int    // StatusCode is the status code of the response, set only if unexpected.
	Status     string // Status is the status of the response, set only if unexpected.
	ReadError  error  // ReadError is the error reading the response body.
	Body       string // Body is the body of the response.

	// WantBody is the expected body of the response,
	// set only if the response status is acceptable and
	// no error occurs when reading the response body.
	WantBody string
}

func (e *UnacceptableResponseError) Error() string {
	if e == nil {
		return ""
	}
	body, wantBody := shortenN(e.Body, 40), shortenN(e.WantBody, 40)
	var b strings.Builder
	if e.StatusCode != 0 || len(e.Status) > 0 {
		if len(e.Status) > 0 {
			b.WriteString("got response status ")
			b.WriteString(e.Status)
		} else {
			b.WriteString("got response status code ")
			b.WriteString(strconv.Itoa(e.StatusCode))
		}
		if e.ReadError != nil {
			b.WriteString("; error reading body: ")
			b.WriteString(e.ReadError.Error())
		} else if len(body) > 0 {
			b.WriteString("; body ")
			b.WriteString(strconv.Quote(body))
		}
	} else if e.ReadError != nil {
		b.WriteString("error reading body: ")
		b.WriteString(e.ReadError.Error())
	} else if len(body) > 0 {
		b.WriteString("got response body ")
		b.WriteString(strconv.Quote(body))
		if len(wantBody) > 0 {
			b.WriteString("; want ")
			b.WriteString(strconv.Quote(wantBody))
		}
	} else if len(wantBody) > 0 {
		b.WriteString("got no response body; want ")
		b.WriteString(strconv.Quote(wantBody))
	}
	return b.String()
}

// Unwrap returns the error reading the response body.
func (e *UnacceptableResponseError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.ReadError
}

// ProtoBufError records an error reading or writing
// ProtoBuf messages and the operation that caused it.
type ProtoBufError struct {
	Op   string // Op is the name of the operation that caused the error.
	Type string // Type is the type name of the operand.
	Err  error  // Err is the error reading or writing ProtoBuf messages.
}

// NewProtoBufError creates a new *ProtoBufError.
//
// It records the specified operation name (Op) and error (err).
//
// It records the type name according to v.
// If v is nil, it records the type name as "<nil>".
// If v is an interface or a pointer, and v is non-nil, it repeatedly replaces
// v with the element of v (i.e., the value that the interface v contains or
// that the pointer v points to) until v is neither an interface nor a pointer
// or v is nil.
// Otherwise, it records the type name of v.
func NewProtoBufError(Op string, v interface{}, err error) *ProtoBufError {
	typeName := "<nil>"
	if v != nil {
		value := reflect.ValueOf(v)
		for k := value.Kind(); (k == reflect.Interface || k == reflect.Pointer) && !value.IsNil(); k = value.Kind() {
			value = value.Elem()
		}
		t := value.Type()
		pkg := t.PkgPath()
		name := t.Name()
		if len(pkg) > 0 {
			if len(name) > 0 {
				typeName = pkg + "." + name
			} else {
				// This case may not happen.
				name = t.String()
				if len(name) == 0 {
					name = "unknown"
				}
				firstDotIdx := strings.IndexByte(name, '.')
				if firstDotIdx >= 0 {
					lastSlashIdx := strings.LastIndexByte(pkg, '/')
					pkgBase := pkg
					if lastSlashIdx >= 0 {
						pkgBase = pkg[lastSlashIdx+1:]
					}
					// Use strings.HasPrefix to fit pkgBase like mypkg.v1.
					if strings.HasPrefix(pkgBase, name[:firstDotIdx]) {
						pkg = pkg[:lastSlashIdx+1]
					}
					typeName = pkg + name
				} else {
					typeName = pkg + "." + name
				}
			}
		} else if len(name) > 0 {
			// This case may not happen.
			typeName = name
		} else {
			typeName = t.String()
		}
		if len(typeName) == 0 {
			typeName = "unknown"
		}
	}
	return &ProtoBufError{
		Op:   Op,
		Type: typeName,
		Err:  err,
	}
}

func (e *ProtoBufError) Error() string {
	if e == nil {
		return ""
	}
	return e.Op + " on " + e.Type + ": " + e.Err.Error()
}

func (e *ProtoBufError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}

// shortenN cuts the specified string if it is too long;
// otherwise, shortenN returns s itself.
//
// It preserves at least the first n bytes.
// The returned string is at most (n + 10) bytes.
//
// If the returned string is different from s,
// it must have the suffix "...".
func shortenN(s string, n int) string {
	if len(s) <= n+10 {
		return s
	}
	// Try to preserve a complete word.
	end := n + 7 // 7 = 10 (at most 10 more bytes) - 3 ("...")
	idx := strings.LastIndexFunc(s[n:end], unicode.IsSpace)
	if idx < 0 {
		return s[:end] + "..."
	}
	return s[:n+idx] + "..."
}
