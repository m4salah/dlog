By the end of this tutorial you'll learn: 

* Creating a new extension
* Creating a custom installation that loads the core and your extension
* Using dlog public API to manipulate the page before processing

# What are we creating

We will create a new dlog extension that adds "Hello World!" before the page text.

# Creating an extension

dlog extensions are Go modules (check extensions for more details). So make sure Go toolchain is installed on your system.

First create an empty directory and initialize a new go module in it

```shell
mkdir helloworld
cd helloworld
go mod init github.com/emad-elsaid/helloworld
go get github.com/m4salah/dlog
```

Replace the URL to your github account or any other URL where your extension will be hosted. as per Go modules convention.

# Create a custom installation

To test our extension we need a `main` package that loads dlog and your own extension. 

We'll create `cmd/dlog/dlog.go` that acts as custom installation.

```shell
mkdir -p cmd/dlog
```

Create a file under `cmd/dlog/dlog.go` that has the following content.

```go
package main

import (
	// Core
	"github.com/m4salah/dlog"

	// Extensions
	_ "github.com/emad-elsaid/helloworld"
)

func main() {
	dlog.Start()
}
```

# Create an extension

Lets make sure Go finds a `helloworld` package in your module root. it'll do nothing for now.

Create a file `helloworld.go` that contains the package name.

```go
package helloworld
```

# Run your custom installation

Now running `cmd/dlog/dlog.go` will start the dlog core with only your extension loaded. so it's a clean environment that include only the dlog core and no other extensions.

```shell
go run ./cmd/dlog/dlog.go
```

You should see output similar to the following. And navigating to [http://localhost:3000](http://localhost:3000) should drop you in the editor to create your `index.md` page.

```
2022/11/17 17:13:38  Template  (64.165µs) commands
2022/11/17 17:13:38  Template  (47.627µs) edit
2022/11/17 17:13:38  Template  (53.813µs) layout
2022/11/17 17:13:38  Template  (21.99µs) pages
2022/11/17 17:13:38  Template  (27.596µs) properties
2022/11/17 17:13:38  Template  (67.411µs) quick_commands
2022/11/17 17:13:38  Template  (116.689µs) sidebar
2022/11/17 17:13:38  Template  (80.292µs) view
2022/11/17 17:13:38 Starting server: 127.0.0.1:3000
```

/info From now on any change to any of the Go files will require restarting the dlog server


# Create your first test page

* Try opening  [http://localhost:3000](http://localhost:3000)
* Enter any text. for example: "We are creating a Hello world dlog extension."
* Click "Save" or "Ctrl+S"
* You should see your page rendered in HTML

# Define a Preprocessor

Packages add features to dlog by calling `Register*` functions in the `init` function of the page. This allow registering a group of types for dlog to use in the appropriate time. Like:

* Preprocessor
* Autocomplete
* Command

For our extension we want to add "Hello world!" before the actual page content. this is exactly what the Preprocessor is for. a function that processes the page text before rendering it to HTML. 

We will create a function that implement the [Preprocessor interface](https://pkg.go.dev/github.com/m4salah/dlog#Preprocessor). `helloworld.go` should have the following content.

```go
package helloworld

func addHelloWorld(input string) string {
	return "Hello world!\n" + input
}
```

This is a function that takes the page content as string and return the content after processing. You can manipulate the page content as you wish in this function. for us we added a line in the beginning of the page.

# The Init function

Now we'll need to register this function as a preprocessor. we'll do this by importing dlog core and use [`RegisterPreprocessor`](https://pkg.go.dev/github.com/m4salah/dlog#RegisterPreprocessor).

```go
package helloworld

import "github.com/m4salah/dlog"

func init() {
	dlog.RegisterPreprocessor(addHelloWorld)
}

func addHelloWorld(input string) string {
	return "Hello world!\n" + input
}
```

Restarting the server and refreshing your web page will show the following: 

```
Hello world!
We are creating a Hello world dlog extension.
```

# Success

Congrates, You created a new dlog extension. Now you can publish this extension to github and import it in any custom installation of dlog.

Also you may try to explore dlog package documentation to get familiar with other types and `Register` functions.

