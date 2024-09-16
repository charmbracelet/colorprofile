// This package provides a writer that can be piped into to degrade the colors
// based on the terminal capabilities and profile.
package main

import (
	"io"
	"log"
	"os"

	"github.com/charmbracelet/colorprofile"
)

func main() {
	w := colorprofile.NewWriter(os.Stdout, os.Environ())

	// Read from stdin and write to stdout
	bts, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(bts)
}
