package util

import (
	"fmt"
	"os"
)

// Die prints error and panic
func Die(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	panic(err)
}
