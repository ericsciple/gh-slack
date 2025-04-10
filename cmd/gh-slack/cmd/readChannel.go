package cmd

import (
	"fmt"
	"strings"

	"github.com/ericsciple/gh-slack/internal/api"
	"github.com/ericsciple/gh-slack/internal/markdown"
	"github.com/spf13/cobra"
)

func ReadChannelCmd() *cobra.Command {
	var verbose bool

	cmd := &cobra.Command{
		Use:   "read-channel <CHANNEL_URL>",
		Short: "Reads a Slack channel and outputs the messages as markdown",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			url := args[0]

			// Extract channel ID from URL
			// Expected format: https://{team}.slack.com/archives/{channelID}
			parts := strings.Split(url, "/")
			if len(parts) < 5 {
				return fmt.Errorf("invalid Slack channel URL: %s", url)
			}
			
			channelID := parts[len(parts)-1]
			
			if verbose {
				fmt.Printf("Channel ID: %s\n", channelID)
			}

			client, err := api.NewClient(verbose)
			if err != nil {
				return err
			}

			// Get channel messages
			messages, err := client.GetChannelHistory(channelID)
			if err != nil {
				return err
			}

			if verbose {
				fmt.Printf("Fetched %d messages\n", len(messages))
			}

			// Convert to markdown
			md := markdown.ConvertMessages(messages, client)
			fmt.Println(md)

			return nil
		},
	}

	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	return cmd
}
