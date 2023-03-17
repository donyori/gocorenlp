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

package errors

import (
	stderrors "errors"

	gogoerrors "github.com/donyori/gogo/errors"
)

// Export functions from standard package errors and
// github.com/donyori/gogo/errors for convenience.

// New directly calls the standard errors.New.
func New(msg string) error {
	return stderrors.New(msg)
}

// Unwrap directly calls the standard errors.Unwrap.
func Unwrap(err error) error {
	return stderrors.Unwrap(err)
}

// Is directly calls the standard errors.Is.
func Is(err, target error) bool {
	return stderrors.Is(err, target)
}

// As directly calls the github.com/donyori/gogo/errors.As.
func As(err error, target any) bool {
	return gogoerrors.As(err, target)
}

// Join directly calls the standard errors.Join.
func Join(errs ...error) error {
	return stderrors.Join(errs...)
}
