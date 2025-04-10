package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var readChannelCmd = &cobra.Command{
	Use:   "read-channel <CHANNEL_URL>",
	Short: "Reads a Slack channel and outputs the messages as markdown",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Simple implementation that just prints Hello World
		fmt.Println("Hello World")
		fmt.Println("Channel URL:", args[0])
		return nil
	},
}
