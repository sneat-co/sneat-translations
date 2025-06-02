package main

import (
	"fmt"
	"os"
)

func main() {
	if err := translateMissing("en-UK"); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
