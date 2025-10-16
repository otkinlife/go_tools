package ai_tools

import (
	"fmt"
	"time"

	"github.com/otkinlife/go_tools/http_tools"
)

// AIClient represents the AI API client
type AIClient struct {
	config *AIConfig
}

// NewAIClient creates a new AI client with the given configuration
func NewAIClient(config *AIConfig) *AIClient {
	if config.BaseURL == "" {
		config.BaseURL = "https://api.openai.com/v1"
	}
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}
	return &AIClient{config: config}
}

// ChatCompletion sends a chat completion request to the AI API
func (c *AIClient) ChatCompletion(request *ChatRequest) (*ChatResponse, error) {
	// Create HTTP client
	client, err := http_tools.NewReqClient("POST", c.config.BaseURL+"/chat/completions")
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP client: %w", err)
	}
	defer client.Close()

	// Set headers
	headers := map[string]string{
		"Authorization": "Bearer " + c.config.APIKey,
		"Content-Type":  "application/json",
	}
	client.SetHeaders(headers)

	// Set timeout
	client.SetTimeout(c.config.Timeout)

	// Set JSON body
	if err := client.SetJson(request); err != nil {
		return nil, fmt.Errorf("failed to set JSON body: %w", err)
	}

	// Send request
	if err := client.Send(); err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	// Check status code
	statusCode := client.GetHttpCode()
	if statusCode != 200 {
		var errorResp ErrorResponse
		if err := client.LoadBody(&errorResp); err != nil {
			return nil, fmt.Errorf("API request failed with status %d: %s", statusCode, client.GetBodyString())
		}
		return nil, fmt.Errorf("API error (%d): %s", statusCode, errorResp.Error.Message)
	}

	// Parse response
	var response ChatResponse
	if err := client.LoadBody(&response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &response, nil
}

// SimpleChat sends a simple chat message and returns the response content
func (c *AIClient) SimpleChat(model, message string) (string, error) {
	request := &ChatRequest{
		Model: model,
		Messages: []ChatMessage{
			{Role: "user", Content: message},
		},
	}

	response, err := c.ChatCompletion(request)
	if err != nil {
		return "", err
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no response choices returned")
	}

	return response.Choices[0].Message.Content, nil
}

// ChatWithMessages sends a chat request with multiple messages
func (c *AIClient) ChatWithMessages(model string, messages []ChatMessage, options ...ChatOption) (*ChatResponse, error) {
	request := &ChatRequest{
		Model:    model,
		Messages: messages,
	}

	// Apply options
	for _, option := range options {
		option(request)
	}

	return c.ChatCompletion(request)
}

// ChatOption defines functional options for chat requests
type ChatOption func(*ChatRequest)

// WithMaxTokens sets the maximum tokens for the response
func WithMaxTokens(maxTokens int) ChatOption {
	return func(r *ChatRequest) {
		r.MaxTokens = &maxTokens
	}
}

// WithTemperature sets the temperature for the response
func WithTemperature(temperature float64) ChatOption {
	return func(r *ChatRequest) {
		r.Temperature = &temperature
	}
}

// WithTopP sets the top_p parameter for nucleus sampling
func WithTopP(topP float64) ChatOption {
	return func(r *ChatRequest) {
		r.TopP = &topP
	}
}

// WithStop sets stop sequences
func WithStop(stop ...string) ChatOption {
	return func(r *ChatRequest) {
		r.Stop = stop
	}
}

// WithUser sets the user identifier
func WithUser(user string) ChatOption {
	return func(r *ChatRequest) {
		r.User = user
	}
}

// WithStream enables streaming response
func WithStream(stream bool) ChatOption {
	return func(r *ChatRequest) {
		r.Stream = stream
	}
}

// WithPresencePenalty sets the presence penalty
func WithPresencePenalty(penalty float64) ChatOption {
	return func(r *ChatRequest) {
		r.PresencePenalty = &penalty
	}
}

// WithFrequencyPenalty sets the frequency penalty
func WithFrequencyPenalty(penalty float64) ChatOption {
	return func(r *ChatRequest) {
		r.FrequencyPenalty = &penalty
	}
}
