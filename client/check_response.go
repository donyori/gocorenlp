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
	"io"
	"net/http"
	"strings"

	"github.com/donyori/gocorenlp/errors"
)

// checkResponse checks the specified HTTP response and reads its body.
// It checks the response by checking its status and body.
//
// acceptStatus is a function to check the response status.
// Its two arguments are resp.StatusCode and resp.Status.
// It returns a boolean value indicating whether the status is acceptable.
// If it is nil, checkResponse accepts the status if the status code is 2XX.
//
// acceptBody is a function to check the response body.
// Its argument is the response body.
// It returns a boolean value indicating whether the body is acceptable.
// If it is nil and wantBody is not empty,
// checkResponse accepts the body if it is the same as wantBody.
//
// wantBody is the expected body of the response.
// It will be stored in the returned error
// if the response body is unacceptable.
//
// checkResponse returns the response body, and reports an error
// if the response is not as expected.
// If the returned error is non-nil, it is of type
// *github.com/donyori/gocorenlp/errors.UnacceptableResponseError.
//
// It closes resp.Body.
func checkResponse(
	resp *http.Response,
	acceptStatus func(statusCode int, status string) bool,
	acceptBody func(body []byte) bool,
	wantBody string,
) (body []byte, err error) {
	defer closeIgnoreError(resp.Body)
	if acceptStatus == nil {
		acceptStatus = func(statusCode int, _ string) bool {
			return statusCode >= 200 && statusCode < 300
		}
	}
	if acceptBody == nil && len(wantBody) > 0 {
		acceptBody = func(body []byte) bool {
			return string(body) == wantBody
		}
	}
	respErr := new(errors.UnacceptableResponseError)
	statusCode, status := resp.StatusCode, resp.Status
	if !acceptStatus(statusCode, status) {
		respErr.StatusCode, respErr.Status = statusCode, status
		err = respErr
	}
	body, respErr.ReadError = io.ReadAll(resp.Body)
	if len(body) > 0 {
		respErr.Body = string(body)
	}
	if respErr.ReadError != nil {
		err = respErr
		return
	}
	if acceptBody != nil && !acceptBody(body) {
		err = respErr
		respErr.WantBody = wantBody
	}
	return
}

// makeAcceptBodyEqualTrimSpace returns a closure function that reports whether
//  strings.TrimSpace(string(body)) == strings.TrimSpace(want)
//
// The returned function can be used as the argument acceptBody of
// the function checkResponse.
func makeAcceptBodyEqualTrimSpace(want string) func(body []byte) bool {
	return func(body []byte) bool {
		return strings.TrimSpace(string(body)) == strings.TrimSpace(want)
	}
}

// closeIgnoreError closes the specified closer if it is non-nil
// and ignores any error encountered.
func closeIgnoreError(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}
