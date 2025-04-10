package main

import (
	"fmt"
	"os"

	"github.com/rneatherway/gh-slack/cmd/gh-slack/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

// In the function where commands are added (location may vary):
rootCmd.AddCommand(
    // ...existing commands...
    commands.ReadChannelCmd(),
)
