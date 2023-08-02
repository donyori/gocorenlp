# gocorenlp

[![Go Report Card](https://goreportcard.com/badge/github.com/donyori/gocorenlp)](https://goreportcard.com/report/github.com/donyori/gocorenlp)
[![Go Reference](https://pkg.go.dev/badge/github.com/donyori/gocorenlp.svg)](https://pkg.go.dev/github.com/donyori/gocorenlp)

A Go (Golang) client for [Stanford CoreNLP server](https://stanfordnlp.github.io/CoreNLP/corenlp-server.html "CoreNLP Server").

## Installation

Similar to getting other Go libraries,
run the `go get` command for the latest version:

```shell
go get github.com/donyori/gocorenlp@latest
```

Or specify a particular version (for example, v0.1.0):

```shell
go get github.com/donyori/gocorenlp@v0.1.0
```

For more information, see `go get` documentation
[here](https://pkg.go.dev/cmd/go#hdr-Add_dependencies_to_current_module_and_install_them "Add dependencies to current module and install them")
and [here](https://go.dev/ref/mod#go-get "Go Modules Reference - go get").

## Usage

### 0. Prerequisites

We assume that you are familiar with the Stanford CoreNLP server.
If not, see its [official page](https://stanfordnlp.github.io/CoreNLP/corenlp-server.html "CoreNLP Server").

Before using this library, you need to have Stanford CoreNLP installed or
have a Stanford CoreNLP server available somewhere.
To install Stanford CoreNLP, see its [download page](https://stanfordnlp.github.io/CoreNLP/download.html "Download CoreNLP").

---

### 1. Start with the `client` package

The `client` package provides functionality to annotate
your human language text with the CoreNLP server.
See the [documentation of package `client`](https://pkg.go.dev/github.com/donyori/gocorenlp/client "client package") for details.

Now we assume that you have already launched a CoreNLP server
on `127.0.0.1:9000`.
(If your server is elsewhere, see [Section 2](#2-create-a-client-with-custom-settings "2. Create a client with custom settings").)

Here is a simple example of using this server to annotate the text

> *The quick brown fox jumped over the lazy dog.*

```go
package main

import (
	"fmt"

	"github.com/donyori/gocorenlp/client"
	"github.com/donyori/gocorenlp/model/v4.5.3-5250f9faf9f1/pb"
)

func main() {
	text := "The quick brown fox jumped over the lazy dog."
	annotators := "tokenize,ssplit,pos,lemma"

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

	// Print some annotation results.
	fmt.Println("Original text:", doc.GetText())
	fmt.Println("+--------+-----+--------+")
	fmt.Println("| Word   | POS | Lemma  |")
	fmt.Println("+--------+-----+--------+")
	for _, token := range doc.GetSentence()[0].GetToken() {
		fmt.Printf(
			"| %-7s| %-4s| %-7s|\n",
			token.GetWord(),
			token.GetPos(),
			token.GetLemma(),
		)
	}
	fmt.Println("+--------+-----+--------+")
}
```

It outputs:

```text
Original text: The quick brown fox jumped over the lazy dog.
+--------+-----+--------+
| Word   | POS | Lemma  |
+--------+-----+--------+
| The    | DT  | the    |
| quick  | JJ  | quick  |
| brown  | JJ  | brown  |
| fox    | NN  | fox    |
| jumped | VBD | jump   |
| over   | IN  | over   |
| the    | DT  | the    |
| lazy   | JJ  | lazy   |
| dog    | NN  | dog    |
| .      | .   | .      |
+--------+-----+--------+
```

---

### 2. Create a client with custom settings

The previous example uses the default client, which can only connect to
the CoreNLP server on `127.0.0.1:9000`.

If you want to use a CoreNLP server elsewhere, you need to create a new client.

To create a client, use function `New`.
The supported options are in struct `Options`.

Here is an example snippet:

```go
c, err := client.New(&client.Options{
	Hostname:   "localhost", // Set the hostname here. If omitted, "127.0.0.1" is used.
	Port:       8080,        // Set the port number here. If omitted, 9000 is used.
	StatusPort: 8081,        // Set the port number of the status server here. If omitted, it is the same as Port.

	Timeout:    time.Second * 15,      // Set a timeout for each request here.
	Charset:    "utf-8",               // Set the charset of your text here. If omitted, "utf-8" is used.
	Annotators: "tokenize,ssplit,pos", // Set the default annotators here.

	// Set the username and password here
	// if your server requires a basic auth.
	Username: "Alice",
	Password: "Alice's password",

	// If your server has a server ID
	// (i.e., server name, set by -server_id),
	// set it here.
	ServerID: "CoreNLPServer",
})
if err != nil {
	panic(err) // handle error
}
// Now you can work with the new client c.
```

---

### 3. Check status and stop server

You can check if the server is online (liveness) and
ready to accept connections (readiness)
using functions/methods `Live` and `Ready`:

```go
if err := client.Live(); err == nil {
	fmt.Println("Server is live.")
} else {
	fmt.Println("Server is offline.")
}

if err := client.Ready(); err == nil {
	fmt.Println("Server is ready to accept connections.")
} else {
	fmt.Println("Server is not ready.")
}
```

In addition, you can shut down the server through the client.

To shut down a local server, use function `Shutdown` or
client's method `ShutdownLocal`:

```go
err := client.Shutdown()
if err == nil {
	fmt.Println("Server has been shut down.")
} else {
	panic(err) // handle error
}
```

To shut down a remote server, you need to provide the shutdown key
and using client's method `Shutdown`.
(Don't know what the shutdown key is?
See [here](https://stanfordnlp.github.io/CoreNLP/corenlp-server.html#stopping-the-server "Stopping the Server").)

---

### 4. Cache annotation results

You can cache the annotation results for future use.
To do this, use functions/methods `AnnotateRaw` or `AnnotateStringRaw`.

Here is an example snippet to cache the annotation results in a local file:

```go
// Create a file to save the annotation result.
filename := "./annotation.ann"
f, err := os.Create(filename)
if err != nil {
	panic(err) // handle error
}
defer f.Close()

// Annotate the text with the specified annotators
// and store the result in f.
_, err = client.AnnotateStringRaw(text, annotators, f)
if err != nil {
	panic(err) // handle error
}
// Then you can use the annotation by reading it from the file.
```

`AnnotateRaw` and `AnnotateStringRaw` output data without parsing.
You can decode it into the document model by function `DecodeResponseBody`
in our package `model`:

```go
doc := new(pb.Document) // specify your document model
err := model.DecodeResponseBody(data, doc) // data is that output by AnnotateRaw or AnnotateStringRaw
if err != nil {
	panic(err) // handle error
}
```

If you put many annotation results together in a large file,
you can decode them using `ResponseBodyDecoder`:

```go
// Open your annotation result file.
filename := "./annotation.ann"
f, err := os.Open(filename)
if err != nil {
	panic(err) // handle error
}
defer f.Close()

// Create a ResponseBodyDecoder on it.
dec := model.NewResponseBodyDecoder(f)

// Decode the annotation results until EOF.
for {
	doc := new(pb.Document) // specify your document model
	err := dec.Decode(doc)
	if err != nil {
		if errors.Is(err, io.EOF) {
			break
		}
		panic(err) // handle error
	}
	// Work with doc.
}
```

---

### 5. Annotation data model

The CoreNLP server provides several forms to present annotation results,
such as JSON, XML, text, and serialized formats (including ProtoBuf).
(See [this page](https://stanfordnlp.github.io/CoreNLP/corenlp-server.html#annotate-with-corenlp- "Annotate with CoreNLP: /") for details.)

Our client asks the CoreNLP server to serialize the results in
[Protocol Buffers (ProtoBuf)](https://protobuf.dev/ "Protocol Buffers").

At the current stage, we provide the models supporting CoreNLP 3.6.0,
and 4.0.0 to 4.5.4.
These models are organized into subpackages of `model` named in the form

> github.com/donyori/gocorenlp/model/vX.Y.Z-abcdefabcdef/pb

where `vX.Y.Z-abcdefabcdef` means:

* `X.Y.Z` is the version of CoreNLP.
* `abcdefabcdef` is a 12-character prefix of the commit hash of
the retrieved ProtoBuf file in the Stanford CoreNLP project.

For example, the model for CoreNLP 4.5.0 is in

> github.com/donyori/gocorenlp/model/v4.5.0-45b47e245c36/pb

See the [documentation of package `model`](https://pkg.go.dev/github.com/donyori/gocorenlp/model "model package")
for details.

The following table shows the correspondence between
models and CoreNLP versions:

| **Model Package**            | **CoreNLP Version** |
|------------------------------|---------------------|
| model/v3.6.0-29765338a2e8/pb | 3.6.0               |
| model/v4.0.0-2b3dd38abe00/pb | 4.0.0               |
| model/v4.1.0-a1427196ba6e/pb | 4.1.0               |
| model/v4.2.0-3ad83fc2e42e/pb | 4.2.0               |
| model/v4.2.1-d8d09b2c81a5/pb | 4.2.1, 4.2.2        |
| model/v4.3.0-f885cd198767/pb | 4.3.0, 4.3.1, 4.3.2 |
| model/v4.4.0-e90f30f13c40/pb | 4.4.0               |
| model/v4.5.0-45b47e245c36/pb | 4.5.0, 4.5.1        |
| model/v4.5.2-9c3dfee5af50/pb | 4.5.2               |
| model/v4.5.3-5250f9faf9f1/pb | 4.5.3, 4.5.4        |

*(CoreNLP 4.2.1 and 4.2.2 use exactly the identical ProtoBuf model,
as do 4.3.0, 4.3.1, 4.3.2, and 4.5.0, 4.5.1, and 4.5.3, 4.5.4.)*

Our library also accepts the ProtoBuf model generated by you.
You can retrieve the CoreNLP ProtoBuf file (`edu.stanford.nlp.pipeline.CoreNLP.proto`)
from your CoreNLP `JAR` file (usually `stanford-corenlp-X.Y.Z-sources.jar`),
edit it, and compile it to `.go` file.
Then pass your `Document` struct to our API to make it work with
your CoreNLP server:

```go
doc := new(mymodel.Document) // mymodel.Document is the document struct compiled by you
err := client.AnnotateString(text, annotators, doc)
if err != nil {
	panic(err) // handle error
}
```

You can also retrieve the CoreNLP ProtoBuf file from its
[GitHub repository](https://github.com/stanfordnlp/CoreNLP/blob/main/src/edu/stanford/nlp/pipeline/CoreNLP.proto "CoreNLP.proto").

About how to compile ProtoBuf to Go, see
[this tutorial](https://protobuf.dev/getting-started/gotutorial/ "Protocol Buffer Basics: Go").

---

For more documentation about this library, see on
[*pkg.go.dev*](https://pkg.go.dev/github.com/donyori/gocorenlp "gocorenlp package").

## License

The GNU Affero General Public License 3.0 (AGPL-3.0) - [Yuan Gao](https://github.com/donyori "donyori (Yuan Gao)").
Please have a look at the [LICENSE](LICENSE "GNU Affero General Public License v3.0").

## Contact

You can contact me by email: [\<donyoridoyodoyo@outlook.com\>](mailto:donyoridoyodoyo@outlook.com "mailto:donyoridoyodoyo@outlook.com").
