package ai_tools

import "time"

// ChatMessage represents a single message in the conversation
type ChatMessage struct {
	Role    string `json:"role"`    // "system", "user", "assistant"
	Content string `json:"content"` // Message content
}

// ChatRequest represents the request payload for OpenAI Chat API
type ChatRequest struct {
	Model            string        `json:"model"`                       // Model to use (e.g., "gpt-3.5-turbo", "gpt-4")
	Messages         []ChatMessage `json:"messages"`                    // Conversation messages
	MaxTokens        *int          `json:"max_tokens,omitempty"`        // Maximum tokens to generate
	Temperature      *float64      `json:"temperature,omitempty"`       // Sampling temperature (0-2)
	TopP             *float64      `json:"top_p,omitempty"`             // Nucleus sampling parameter
	N                *int          `json:"n,omitempty"`                 // Number of completions to generate
	Stream           bool          `json:"stream,omitempty"`            // Whether to stream responses
	Stop             []string      `json:"stop,omitempty"`              // Stop sequences
	PresencePenalty  *float64      `json:"presence_penalty,omitempty"`  // Presence penalty (-2 to 2)
	FrequencyPenalty *float64      `json:"frequency_penalty,omitempty"` // Frequency penalty (-2 to 2)
	User             string        `json:"user,omitempty"`              // User identifier
}

// ChatChoice represents a single response choice
type ChatChoice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"` // "stop", "length", "content_filter"
}

// Usage represents token usage information
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChatResponse represents the response from OpenAI Chat API
type ChatResponse struct {
	ID      string       `json:"id"`
	Object  string       `json:"object"`
	Created int64        `json:"created"`
	Model   string       `json:"model"`
	Choices []ChatChoice `json:"choices"`
	Usage   Usage        `json:"usage"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Code    string `json:"code"`
	} `json:"error"`
}

// AIConfig holds configuration for AI client
type AIConfig struct {
	APIKey  string        // API key for authentication
	BaseURL string        // Base URL for the API (default: https://api.openai.com/v1)
	Timeout time.Duration // Request timeout (default: 30s)
}

// DefaultConfig returns a default configuration
func DefaultConfig(apiKey string) *AIConfig {
	return &AIConfig{
		APIKey:  apiKey,
		BaseURL: "https://api.openai.com/v1",
		Timeout: 30 * time.Second,
	}
}

// NewAPIConfig returns configuration for New-API service
func NewAPIConfig(apiKey, baseURL string) *AIConfig {
	return &AIConfig{
		APIKey:  apiKey,
		BaseURL: baseURL,
		Timeout: 30 * time.Second,
	}
}
