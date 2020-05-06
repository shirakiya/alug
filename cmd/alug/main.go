package main

import (
	"fmt"
	"os"

	"github.com/shirakiya/alug/internal/alug"
)

// Version shows build version. Its format is Semver.
var Version = "unset"

// AppName is this application name.
const AppName = "alug"

func main() {
	err := alug.GetConsoleLoginURL(AppName, Version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
