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
