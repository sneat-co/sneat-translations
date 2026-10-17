package main

import (
	"fmt"
	"io"
	"os"
)

var (
	translateMissingFunc           = translateMissing
	osStderr             io.Writer = os.Stderr
)

func main() {
	/* To authenticate with Google Cloud, run:

	gcloud auth application-default login

	*/
	if err := translateMissingFunc("en-UK"); err != nil {
		_, _ = fmt.Fprintln(osStderr, err)
	}
}
