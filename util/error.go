package util

import (
	"fmt"
	"os"
)

// HandleError handles error condition
func HandleError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	panic(err)
}
