# gocorenlp

[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](https://pkg.go.dev/github.com/donyori/gocorenlp)

A Go (Golang) client for [Stanford CoreNLP server](https://stanfordnlp.github.io/CoreNLP/corenlp-server.html "CoreNLP Server").

## Installation

Similar to getting other Go libraries through `go get`:

```console
$ go get github.com/donyori/gocorenlp
```

For more information, see `go get` documentation
[here](https://pkg.go.dev/cmd/go#hdr-Add_dependencies_to_current_module_and_install_them "Add dependencies to current module and install them")
and [here](https://go.dev/ref/mod#go-get "Go Modules Reference - go get").

## Usage

### 0. Prerequisites

We assume that you are familiar with the Stanford CoreNLP server.
If not, see its [official page](https://stanfordnlp.github.io/CoreNLP/corenlp-server.html "CoreNLP Server").

Before using this package, you need to install Stanford CoreNLP.
If not installed, see its [download page](https://stanfordnlp.github.io/CoreNLP/download.html "Download CoreNLP").

---

### 1. Start with the `client` package

Now we assume that you have already launched a CoreNLP server
on `127.0.0.1:9000`.

The `client` package provides functionality to annotate
your human language text with the CoreNLP server.

Here is a simple example:

```go
package main

import (
	"fmt"

	"github.com/donyori/gocorenlp/client"
	"github.com/donyori/gocorenlp/model/v4.4.0-e90f30f13c40/pb"
)

func main() {
	text := "The quick brown fox jumped over the lazy dog."
	annotators := "tokenize,ssplit,pos,lemma"

	// Specify the document model.
	// Depending on your CoreNLP version, import the appropriate model.
	// Here we use the model corresponding to CoreNLP 4.4.0.
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
// Open a file to save the annotation result.
filename := "./annotation.ann"
f, err := os.Open(filename)
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
//
// Note that the data written to the file is in ProtoBuf wire encoding.
// You can decode it using the function
// github.com/donyori/gocorenlp/model.DecodeResponseBody.
```

---

### 5. Annotation data model

The CoreNLP server provides several forms to present annotation results,
such as JSON, XML, and text. (See [this page](https://stanfordnlp.github.io/CoreNLP/corenlp-server.html#annotate-with-corenlp- "Annotate with CoreNLP: /") for details.)

Our client asks the CoreNLP server to serialize the results in
[Protocol Buffers (ProtoBuf)](https://developers.google.com/protocol-buffers).

At the current stage, we provide the models supporting
CoreNLP 4.4.0 and CoreNLP 3.6.0.
The model for CoreNLP 4.4.0 is in `model/v4.4.0-e90f30f13c40/pb`;
that for 3.6.0 is in `model/v3.6.0-29765338a2e8/pb`.

The naming pattern vX.Y.Z-abcdefabcdef means:

* X.Y.Z is the version of CoreNLP.
* abcdefabcdef is a 12-character prefix of the commit hash of
the retrieved ProtoBuf file in the Stanford CoreNLP project.

See the documentation of package `model` for details.

<sub>
P.S. You may find that the package also contains the model for CoreNLP 3.3.0.
This version is actually useless for our package because Stanford CoreNLP
provides the web API server starting with version 3.6.0.
We leave this model here because this version is the first time
the ProtoBuf file has been included in the Stanford CoreNLP release.
Its ProtoBuf message structure is quite different from that of the current version.
</sub>

Our package also accepts the ProtoBuf model generated by you.
You can retrieve the ProtoBuf file (`.proto`) from your CoreNLP and compile it
to `.go` file. Then pass your `Document` struct to our API to make it work with
your CoreNLP server:

```go
doc := new(mymodel.Document) // mymodel.Document is the document struct compiled by you
err := client.AnnotateString(text, annotators, doc)
if err != nil {
	panic(err) // handle error
}
```

You may retrieve the CoreNLP ProtoBuf file from its [GitHub repository](https://github.com/stanfordnlp/CoreNLP/blob/main/src/edu/stanford/nlp/pipeline/CoreNLP.proto "CoreNLP.proto").

About how to compile ProtoBuf to Go, see [this tutorial](https://developers.google.com/protocol-buffers/docs/gotutorial "Protocol Buffer Basics: Go").

---

For more documentation about this package, see on [pkg.go.dev](https://pkg.go.dev/github.com/donyori/gocorenlp "gocorenlp package").

## License

The GNU Affero General Public License 3.0 (AGPL-3.0) - [Yuan Gao](https://github.com/donyori/).
Please have a look at the [LICENSE](LICENSE).

## Contact

You can contact me by email: [<donyoridoyodoyo@outlook.com>](mailto:donyoridoyodoyo@outlook.com "mailto:donyoridoyodoyo@outlook.com").
