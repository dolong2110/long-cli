package main

import (
	"os"
	"wc-tool/cmd"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		os.Exit(1)
	}
}
