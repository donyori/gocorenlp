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

package pb

import (
	pb360 "github.com/donyori/gocorenlp/model/v3.6.0-29765338a2e8/pb"
	pb400 "github.com/donyori/gocorenlp/model/v4.0.0-2b3dd38abe00/pb"
	pb410 "github.com/donyori/gocorenlp/model/v4.1.0-a1427196ba6e/pb"
	pb420 "github.com/donyori/gocorenlp/model/v4.2.0-3ad83fc2e42e/pb"
	pb421 "github.com/donyori/gocorenlp/model/v4.2.1-d8d09b2c81a5/pb"
	pb430 "github.com/donyori/gocorenlp/model/v4.3.0-f885cd198767/pb"
	pb440 "github.com/donyori/gocorenlp/model/v4.4.0-e90f30f13c40/pb"
)

type (
	Doc360 = pb360.Document // Doc360 is an alias for the document structure for Stanford CoreNLP 3.6.0.
	Doc400 = pb400.Document // Doc400 is an alias for the document structure for Stanford CoreNLP 4.0.0.
	Doc410 = pb410.Document // Doc410 is an alias for the document structure for Stanford CoreNLP 4.1.0.
	Doc420 = pb420.Document // Doc420 is an alias for the document structure for Stanford CoreNLP 4.2.0.
	Doc421 = pb421.Document // Doc421 is an alias for the document structure for Stanford CoreNLP 4.2.1.
	Doc422 = pb421.Document // Doc422 is an alias for the document structure for Stanford CoreNLP 4.2.2.
	Doc430 = pb430.Document // Doc430 is an alias for the document structure for Stanford CoreNLP 4.3.0.
	Doc431 = pb430.Document // Doc431 is an alias for the document structure for Stanford CoreNLP 4.3.1.
	Doc432 = pb430.Document // Doc432 is an alias for the document structure for Stanford CoreNLP 4.3.2.
	Doc440 = pb440.Document // Doc440 is an alias for the document structure for Stanford CoreNLP 4.4.0.
)

type (
	Sentence360 = pb360.Sentence // Sentence360 is an alias for the sentence structure for Stanford CoreNLP 3.6.0.
	Sentence400 = pb400.Sentence // Sentence400 is an alias for the sentence structure for Stanford CoreNLP 4.0.0.
	Sentence410 = pb410.Sentence // Sentence410 is an alias for the sentence structure for Stanford CoreNLP 4.1.0.
	Sentence420 = pb420.Sentence // Sentence420 is an alias for the sentence structure for Stanford CoreNLP 4.2.0.
	Sentence421 = pb421.Sentence // Sentence421 is an alias for the sentence structure for Stanford CoreNLP 4.2.1.
	Sentence422 = pb421.Sentence // Sentence422 is an alias for the sentence structure for Stanford CoreNLP 4.2.2.
	Sentence430 = pb430.Sentence // Sentence430 is an alias for the sentence structure for Stanford CoreNLP 4.3.0.
	Sentence431 = pb430.Sentence // Sentence431 is an alias for the sentence structure for Stanford CoreNLP 4.3.1.
	Sentence432 = pb430.Sentence // Sentence432 is an alias for the sentence structure for Stanford CoreNLP 4.3.2.
	Sentence440 = pb440.Sentence // Sentence440 is an alias for the sentence structure for Stanford CoreNLP 4.4.0.
)

type (
	Token360 = pb360.Token // Token360 is an alias for the token structure for Stanford CoreNLP 3.6.0.
	Token400 = pb400.Token // Token400 is an alias for the token structure for Stanford CoreNLP 4.0.0.
	Token410 = pb410.Token // Token410 is an alias for the token structure for Stanford CoreNLP 4.1.0.
	Token420 = pb420.Token // Token420 is an alias for the token structure for Stanford CoreNLP 4.2.0.
	Token421 = pb421.Token // Token421 is an alias for the token structure for Stanford CoreNLP 4.2.1.
	Token422 = pb421.Token // Token422 is an alias for the token structure for Stanford CoreNLP 4.2.2.
	Token430 = pb430.Token // Token430 is an alias for the token structure for Stanford CoreNLP 4.3.0.
	Token431 = pb430.Token // Token431 is an alias for the token structure for Stanford CoreNLP 4.3.1.
	Token432 = pb430.Token // Token432 is an alias for the token structure for Stanford CoreNLP 4.3.2.
	Token440 = pb440.Token // Token440 is an alias for the token structure for Stanford CoreNLP 4.4.0.
)
