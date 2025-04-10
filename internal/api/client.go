package api

import (
	"encoding/json"
	"fmt"
)

// Client represents the API client
type Client struct {
	api func(endpoint string, params map[string]string) ([]byte, error)
}

// Message represents a Slack message
type Message struct {
	Text string `json:"text"`
}

// GetChannelMessages retrieves messages from a channel
func (c *Client) GetChannelMessages(channelID string) ([]Message, error) {
	resp, err := c.api("conversations.history", map[string]string{
		"channel": channelID,
		"limit":   "100",
	})
	if err != nil {
		return nil, err
	}

	var history struct {
		Messages []Message `json:"messages"`
		OK       bool      `json:"ok"`
		Error    string    `json:"error"`
	}

	if err := json.Unmarshal(resp, &history); err != nil {
		return nil, fmt.Errorf("unable to parse API response: %v", err)
	}

	if !history.OK {
		return nil, fmt.Errorf("API error: %s", history.Error)
	}

	return history.Messages, nil
}