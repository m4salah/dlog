dlog ships with a CLI that includes the core and all official extensions. There are cases where you need custom set of extensions:

* Maybe you need only the core features without any extension
* Maybe there is an extension that you don't need or want or misbehaving
* Maybe you developed a set of extensions and you want to include them in your installations

Here is how you can build your own custom dlog with features you select.

# Creating a Go module

Create a directory for your custom installation and initialize a go module in it.

```shell
mkdir custom_dlog
cd custom_dlog
go mod init github.com/yourusername/custom_dlog
```

# Main file

Then create a file `dlog.go` for example with the following content

```go
package main

import (
	// Core
	"github.com/m4salah/dlog"

	// All official extensions
	_ "github.com/m4salah/dlog/extensions/all"
)

func main() {
	dlog.Start()
}
```

# Selecting extensions

The previous file is what dlog ships in `cmd/dlog/dlog.go` if you missed up at any point feel free to go back to it and copy it from there. 

If you want to select specific extensions you can replace `extensions/all` line with a list of extensions that you want.

All extensions are imported to [`extensions/all/all.go`](https://github.com/m4salah/dlog/blob/master/extensions/all/all.go). feel free to copy any of them as needed.

You can also import any extensions that you developed at this point.

# Running your custom dlog

Now use Go to run your custom installation 

```shell
go get github.com/m4salah/dlog
go run dlog.go
```

