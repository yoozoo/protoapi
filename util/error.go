package util

import (
	"fmt"
	"os"
)

// Die prints error and exit
func Die(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}
