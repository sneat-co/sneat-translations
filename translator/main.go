package main

import (
	"fmt"
	"os"
)

func main() {
	/* To authenticate with Google Cloud, run:

	gcloud auth application-default login

	*/
	if err := translateMissing("en-UK"); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
