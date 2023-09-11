package main

import (
	// Core
	"github.com/m4salah/xlog"

	// All official extensions
	_ "github.com/m4salah/xlog/extensions/all"
)

func main() {
	xlog.Start()
}
