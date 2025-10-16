package ai_tools

import (
	"strings"
	"testing"
	"time"
)

func TestNewAIClient(t *testing.T) {
	config := &AIConfig{
		APIKey:  "test-key",
		BaseURL: "https://api.test.com/v1",
		Timeout: 10 * time.Second,
	}

	client := NewAIClient(config)
	if client == nil {
		t.Fatal("Expected client to be created")
	}

	if client.config.APIKey != "test-key" {
		t.Errorf("Expected API key to be 'test-key', got '%s'", client.config.APIKey)
	}

	if client.config.BaseURL != "https://api.test.com/v1" {
		t.Errorf("Expected base URL to be 'https://api.test.com/v1', got '%s'", client.config.BaseURL)
	}
}

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig("test-api-key")

	if config.APIKey != "test-api-key" {
		t.Errorf("Expected API key to be 'test-api-key', got '%s'", config.APIKey)
	}

	if config.BaseURL != "https://api.openai.com/v1" {
		t.Errorf("Expected default base URL, got '%s'", config.BaseURL)
	}

	if config.Timeout != 30*time.Second {
		t.Errorf("Expected default timeout to be 30s, got %v", config.Timeout)
	}
}

func TestNewAPIConfig(t *testing.T) {
	config := NewAPIConfig("test-key", "https://new-api.example.com/v1")

	if config.APIKey != "test-key" {
		t.Errorf("Expected API key to be 'test-key', got '%s'", config.APIKey)
	}

	if config.BaseURL != "https://new-api.example.com/v1" {
		t.Errorf("Expected base URL to be 'https://new-api.example.com/v1', got '%s'", config.BaseURL)
	}
}

func TestMessageBuilder(t *testing.T) {
	builder := NewMessageBuilder()

	messages := builder.
		System("You are a helpful assistant.").
		User("Hello!").
		Assistant("Hi there! How can I help you?").
		User("What's the weather like?").
		Build()

	if len(messages) != 4 {
		t.Errorf("Expected 4 messages, got %d", len(messages))
	}

	if messages[0].Role != "system" || messages[0].Content != "You are a helpful assistant." {
		t.Errorf("First message not correct: %+v", messages[0])
	}

	if messages[1].Role != "user" || messages[1].Content != "Hello!" {
		t.Errorf("Second message not correct: %+v", messages[1])
	}

	if messages[2].Role != "assistant" || messages[2].Content != "Hi there! How can I help you?" {
		t.Errorf("Third message not correct: %+v", messages[2])
	}

	if messages[3].Role != "user" || messages[3].Content != "What's the weather like?" {
		t.Errorf("Fourth message not correct: %+v", messages[3])
	}
}

func TestMessageBuilderCount(t *testing.T) {
	builder := NewMessageBuilder()

	if builder.Count() != 0 {
		t.Errorf("Expected count to be 0, got %d", builder.Count())
	}

	builder.User("Test message")
	if builder.Count() != 1 {
		t.Errorf("Expected count to be 1, got %d", builder.Count())
	}
}

func TestMessageBuilderClear(t *testing.T) {
	builder := NewMessageBuilder()
	builder.User("Test message")

	if builder.Count() != 1 {
		t.Errorf("Expected count to be 1 before clear, got %d", builder.Count())
	}

	builder.Clear()
	if builder.Count() != 0 {
		t.Errorf("Expected count to be 0 after clear, got %d", builder.Count())
	}
}

func TestChatOptions(t *testing.T) {
	request := &ChatRequest{
		Model:    "gpt-3.5-turbo",
		Messages: []ChatMessage{{Role: "user", Content: "Hello"}},
	}

	// Test WithMaxTokens
	WithMaxTokens(100)(request)
	if request.MaxTokens == nil || *request.MaxTokens != 100 {
		t.Errorf("Expected MaxTokens to be 100, got %v", request.MaxTokens)
	}

	// Test WithTemperature
	WithTemperature(0.8)(request)
	if request.Temperature == nil || *request.Temperature != 0.8 {
		t.Errorf("Expected Temperature to be 0.8, got %v", request.Temperature)
	}

	// Test WithTopP
	WithTopP(0.9)(request)
	if request.TopP == nil || *request.TopP != 0.9 {
		t.Errorf("Expected TopP to be 0.9, got %v", request.TopP)
	}

	// Test WithStop
	WithStop(".", "!", "?")(request)
	if len(request.Stop) != 3 || request.Stop[0] != "." {
		t.Errorf("Expected Stop to be ['.', '!', '?'], got %v", request.Stop)
	}

	// Test WithUser
	WithUser("user123")(request)
	if request.User != "user123" {
		t.Errorf("Expected User to be 'user123', got '%s'", request.User)
	}

	// Test WithStream
	WithStream(true)(request)
	if !request.Stream {
		t.Errorf("Expected Stream to be true, got %v", request.Stream)
	}
}

// TestRealAPICall 测试实际的 API 调用
func TestRealAPICall(t *testing.T) {
	apiKey := ""
	baseURL := ""
	model := ""

	var config *AIConfig
	config = NewAPIConfig(apiKey, baseURL)

	client := NewAIClient(config)

	t.Run("简单聊天测试", func(t *testing.T) {
		response, err := client.SimpleChat(model, "你好，请回复一个简短的问候")
		if err != nil {
			t.Fatalf("简单聊天失败: %v", err)
		}

		if response == "" {
			t.Error("响应内容为空")
		}

		t.Logf("AI 回复: %s", response)
	})

	t.Run("带选项的聊天测试", func(t *testing.T) {
		messages := NewMessageBuilder().
			System("你是一个友好的助手，请用中文回复，回复要简洁。").
			User("介绍一下 Go 语言的特点").
			Build()

		response, err := client.ChatWithMessages(
			model,
			messages,
			WithMaxTokens(100),
			WithTemperature(0.7),
		)

		if err != nil {
			t.Fatalf("带选项聊天失败: %v", err)
		}

		if len(response.Choices) == 0 {
			t.Fatal("没有返回任何选择")
		}

		content := response.Choices[0].Message.Content
		if content == "" {
			t.Error("回复内容为空")
		}

		if response.Usage.TotalTokens == 0 {
			t.Error("Token 使用量为 0")
		}

		t.Logf("AI 回复: %s", content)
		t.Logf("Token 使用: %d (提示: %d + 完成: %d)",
			response.Usage.TotalTokens,
			response.Usage.PromptTokens,
			response.Usage.CompletionTokens)
	})

	t.Run("多轮对话测试", func(t *testing.T) {
		builder := NewMessageBuilder()
		builder.System("你是一个编程助手，请简洁回答。")

		// 第一轮
		builder.User("什么是 REST API？")
		response1, err := client.ChatWithMessages(model, builder.Build(), WithMaxTokens(50))
		if err != nil {
			t.Fatalf("第一轮对话失败: %v", err)
		}

		// 添加助手回复到对话历史
		builder.Assistant(response1.Choices[0].Message.Content)

		// 第二轮
		builder.User("它有什么优点？")
		response2, err := client.ChatWithMessages(model, builder.Build(), WithMaxTokens(50))
		if err != nil {
			t.Fatalf("第二轮对话失败: %v", err)
		}

		t.Logf("第一轮 - 用户: 什么是 REST API？")
		t.Logf("第一轮 - AI: %s", response1.Choices[0].Message.Content)
		t.Logf("第二轮 - 用户: 它有什么优点？")
		t.Logf("第二轮 - AI: %s", response2.Choices[0].Message.Content)
	})

	t.Run("错误处理测试", func(t *testing.T) {
		invalidConfig := DefaultConfig("invalid-key-test")
		invalidClient := NewAIClient(invalidConfig)

		_, err := invalidClient.SimpleChat(model, "测试")
		if err == nil {
			t.Error("期望返回错误，但没有错误")
		}

		if !strings.Contains(err.Error(), "API error") {
			t.Errorf("期望包含 'API error'，实际错误: %v", err)
		}

		t.Logf("正确捕获到错误: %v", err)
	})
}
