package main

import (
	"fmt"
	"os"

	app "github.com/whizard/bazel-talk-2023/cmd/myapp2/internal"
)

func main() {
	err := app.Run(os.Args, os.Stdout)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
