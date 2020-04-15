package utils

import (
	"fmt"
	"os"
)

func ExitWithError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, err.Error() + "\n")
	os.Exit(1)
}

func Bool2int(b bool) int {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}