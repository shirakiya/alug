package main

import (
	"fmt"
	"os"

	"github.com/shirakiya/alug/internal/alug"
)

func main() {
	err := alug.GetConsoleLoginURL()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
