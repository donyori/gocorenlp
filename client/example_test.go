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

package client_test

import (
	"fmt"
	"os"
	"time"

	"github.com/donyori/gocorenlp/client"
	"github.com/donyori/gocorenlp/model/v4.5.0-45b47e245c36/pb"
)

func Example() {
	// This example is a simple procedure of annotating the string
	//  "The quick brown fox jumped over the lazy dog."
	// with Stanford CoreNLP.
	//
	// Don't run this example because there is no CoreNLP server
	// in the example environment.

	// Before calling client.AnnotateString,
	// launch a Stanford CoreNLP server listening on 127.0.0.1:9000.

	text := "The quick brown fox jumped over the lazy dog."
	annotators := "tokenize,ssplit,pos"

	// Specify the document model.
	// Depending on your CoreNLP version, use the appropriate model.
	// See package github.com/donyori/gocorenlp/model for details.
	doc := new(pb.Document)

	// Annotate the text with the specified annotators
	// and store the result in doc.
	err := client.AnnotateString(text, annotators, doc)
	if err != nil {
		panic(err) // handle error
	}

	// Work with the document.
	fmt.Println(doc.GetText())

	// [Optional]
	// Shut down the server.
	err = client.Shutdown()
	if err != nil {
		panic(err) // handle error
	}
}

func Example_newClient() {
	// Don't run this example because there is no CoreNLP server
	// in the example environment.

	// Before creating the client with default settings,
	// launch a Stanford CoreNLP server listening on 127.0.0.1:9000.

	c, err := client.New(nil) // Create a new client with default settings.
	if err != nil {
		panic(err) // handle error
	}

	text := "The quick brown fox jumped over the lazy dog."
	annotators := "tokenize,ssplit,pos"

	// Specify the document model.
	// Depending on your CoreNLP version, use the appropriate model.
	// See package github.com/donyori/gocorenlp/model for details.
	doc := new(pb.Document)

	// Annotate the text with the specified annotators
	// and store the result in doc.
	err = c.AnnotateString(text, annotators, doc)
	if err != nil {
		panic(err) // handle error
	}

	// Work with the document.
	fmt.Println(doc.GetText())

	// [Optional]
	// Shut down the server.
	err = c.ShutdownLocal()
	if err != nil {
		panic(err) // handle error
	}
}

func Example_specifyOptions() {
	// Don't run this example because there is no CoreNLP server
	// in the example environment.

	// Before creating the client,
	// launch a Stanford CoreNLP server listening on the specified address.
	// The default address is 127.0.0.1:9000.

	c, err := client.New(&client.Options{
		Hostname:   "localhost", // Set the hostname here. If omitted, "127.0.0.1" will be used.
		Port:       8080,        // Set the port number here. If omitted, 9000 will be used.
		StatusPort: 8081,        // Set the port number of the status server here. If omitted, it will be the same as Port.

		Timeout:    time.Second * 15,      // Set a timeout for each request here.
		Charset:    "utf-8",               // Set the charset of your text here. If omitted, "utf-8" will be used.
		Annotators: "tokenize,ssplit,pos", // Set the default annotators here.

		// Set the username and password here
		// if your server requires a basic auth.
		Username: "Alice",
		Password: "Alice's password",

		// If your server has a server ID
		// (i.e., server name, set by -server_id),
		// set it here.
		ServerId: "CoreNLPServer",
	})
	if err != nil {
		panic(err) // handle error
	}

	text := "The quick brown fox jumped over the lazy dog."

	// Specify the document model.
	// Depending on your CoreNLP version, use the appropriate model.
	// See package github.com/donyori/gocorenlp/model for details.
	doc := new(pb.Document)

	// Annotate the text with the default annotators (specified in Options above)
	// and store the result in doc.
	err = c.AnnotateString(text, "", doc)
	if err != nil {
		panic(err) // handle error
	}

	// Work with the document.
	fmt.Println(doc.GetText())

	// [Optional]
	// Shut down the server.
	err = c.ShutdownLocal()
	if err != nil {
		panic(err) // handle error
	}
}

func Example_cacheAnnotation() {
	// In this example, we cache the response from the Stanford CoreNLP server
	// into a local file for future use.
	//
	// Don't run this example because there is no CoreNLP server
	// in the example environment.

	// Before calling client.AnnotateStringRaw,
	// launch a Stanford CoreNLP server listening on 127.0.0.1:9000.

	text := "The quick brown fox jumped over the lazy dog."
	annotators := "tokenize,ssplit,pos"

	// Create a file to save the annotation result.
	filename := "./annotation.ann"
	f, err := os.Create(filename)
	if err != nil {
		panic(err) // handle error
	}
	defer func(f *os.File) {
		// This error handler can usually be omitted.
		if err := f.Close(); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}(f)

	// Annotate the text with the specified annotators
	// and store the result in f.
	_, err = client.AnnotateStringRaw(text, annotators, f)
	if err != nil {
		panic(err) // handle error
	}

	// Then you can use the annotation by reading it from the file.
	//
	// Note that the data written to the file is in ProtoBuf wire encoding.
	// You can decode it using the function
	// github.com/donyori/gocorenlp/model.DecodeResponseBody.

	// [Optional]
	// Shut down the server.
	err = client.Shutdown()
	if err != nil {
		panic(err) // handle error
	}
}
