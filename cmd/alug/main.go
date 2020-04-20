package main

import (
	"fmt"
	"os"

	"github.com/shirakiya/alug/internal/alug"
)

const (
	// AppName is this application name.
	AppName = "alug"

	// Version shows build version. (Semver)
	Version = "unset"
)

func main() {
	err := alug.GetConsoleLoginURL()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
