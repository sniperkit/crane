package main

import (
	"os"

	// internal
	"github.com/sniperkit/snk.fork.michaelsauter-crane/pkg/repack/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
