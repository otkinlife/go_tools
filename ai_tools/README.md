# AI Tools

AI Tools 是一个用于调用 OpenAI Chat API 和兼容服务（如 New-API）的 Go SDK。它基于项目中的 `http_tools` 包构建，提供简洁易用的接口来进行 AI 对话。

## 特性

- 🚀 简单易用的 API 接口
- 🔧 支持 OpenAI 和 New-API 服务
- 💬 支持多轮对话
- ⚙️ 丰富的配置选项
- 🛠️ 消息构建器（MessageBuilder）
- 🧪 完整的测试覆盖
- 📝 详细的使用示例

## 安装

```bash
go get github.com/otkinlife/go_tools/ai_tools
```

## 快速开始

### 基本使用

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/otkinlife/go_tools/ai_tools"
)

func main() {
    // 创建配置
    config := ai_tools.DefaultConfig("your-api-key-here")
    
    // 创建客户端
    client := ai_tools.NewAIClient(config)
    
    // 发送简单消息
    response, err := client.SimpleChat("gpt-3.5-turbo", "Hello, how are you?")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("AI Response:", response)
}
```

### 使用 New-API 服务

```go
// 配置 New-API 服务
config := ai_tools.NewAPIConfig("your-api-key", "https://your-new-api-host.com/v1")
client := ai_tools.NewAIClient(config)

response, err := client.SimpleChat("gpt-3.5-turbo", "Hello!")
if err != nil {
    log.Fatal(err)
}
fmt.Println(response)
```

## 高级使用

### 多轮对话

```go
// 使用消息构建器创建对话
builder := ai_tools.NewMessageBuilder()
messages := builder.
    System("You are a helpful coding assistant.").
    User("How do I create a REST API in Go?").
    Build()

// 发送带选项的对话
response, err := client.ChatWithMessages(
    "gpt-3.5-turbo",
    messages,
    ai_tools.WithMaxTokens(500),
    ai_tools.WithTemperature(0.7),
    ai_tools.WithTopP(0.9),
)

if err != nil {
    log.Fatal(err)
}

fmt.Println("AI Response:", response.Choices[0].Message.Content)
fmt.Printf("Tokens used: %d\\n", response.Usage.TotalTokens)
```

### 持续对话

```go
builder := ai_tools.NewMessageBuilder()
builder.System("You are a helpful assistant.")

// 第一轮对话
builder.User("What is Go programming language?")
response1, err := client.ChatWithMessages("gpt-3.5-turbo", builder.Build())
if err != nil {
    log.Fatal(err)
}

// 添加助手回复到对话历史
builder.Assistant(response1.Choices[0].Message.Content)

// 继续对话
builder.User("What are its main advantages?")
response2, err := client.ChatWithMessages("gpt-3.5-turbo", builder.Build())
if err != nil {
    log.Fatal(err)
}

fmt.Println("Final response:", response2.Choices[0].Message.Content)
```

## API 参考

### AIConfig

配置结构体，用于设置 API 密钥、基础 URL 和超时时间。

```go
type AIConfig struct {
    APIKey  string        // API 密钥
    BaseURL string        // API 基础 URL
    Timeout time.Duration // 请求超时时间
}
```

### AIClient

主要的客户端结构体。

#### 方法

- `NewAIClient(config *AIConfig) *AIClient` - 创建新的客户端
- `SimpleChat(model, message string) (string, error)` - 发送简单消息
- `ChatCompletion(request *ChatRequest) (*ChatResponse, error)` - 发送完整的聊天请求
- `ChatWithMessages(model string, messages []ChatMessage, options ...ChatOption) (*ChatResponse, error)` - 发送带消息的聊天请求

### MessageBuilder

消息构建器，用于构建对话消息。

#### 方法

- `NewMessageBuilder() *MessageBuilder` - 创建新的消息构建器
- `System(content string) *MessageBuilder` - 添加系统消息
- `User(content string) *MessageBuilder` - 添加用户消息
- `Assistant(content string) *MessageBuilder` - 添加助手消息
- `Build() []ChatMessage` - 构建消息列表
- `Clear() *MessageBuilder` - 清空消息
- `Count() int` - 获取消息数量

### 聊天选项

可以使用以下选项来自定义聊天请求：

- `WithMaxTokens(maxTokens int)` - 设置最大 token 数
- `WithTemperature(temperature float64)` - 设置温度参数 (0-2)
- `WithTopP(topP float64)` - 设置 top_p 参数
- `WithStop(stop ...string)` - 设置停止词
- `WithUser(user string)` - 设置用户标识
- `WithStream(stream bool)` - 启用流式响应
- `WithPresencePenalty(penalty float64)` - 设置存在惩罚
- `WithFrequencyPenalty(penalty float64)` - 设置频率惩罚

## 数据结构

### ChatMessage

```go
type ChatMessage struct {
    Role    string `json:"role"`    // "system", "user", "assistant"
    Content string `json:"content"` // 消息内容
}
```

### ChatRequest

```go
type ChatRequest struct {
    Model            string        `json:"model"`
    Messages         []ChatMessage `json:"messages"`
    MaxTokens        *int          `json:"max_tokens,omitempty"`
    Temperature      *float64      `json:"temperature,omitempty"`
    // ... 其他字段
}
```

### ChatResponse

```go
type ChatResponse struct {
    ID      string       `json:"id"`
    Object  string       `json:"object"`
    Created int64        `json:"created"`
    Model   string       `json:"model"`
    Choices []ChatChoice `json:"choices"`
    Usage   Usage        `json:"usage"`
}
```

## 错误处理

SDK 提供详细的错误信息，包括 HTTP 状态码和 API 错误消息：

```go
response, err := client.SimpleChat("gpt-3.5-turbo", "Hello")
if err != nil {
    fmt.Printf("Error occurred: %v\\n", err)
    // Error occurred: API error (401): Invalid API key
    return
}
```

## 自定义配置

### 设置超时时间

```go
config := &ai_tools.AIConfig{
    APIKey:  "your-api-key",
    BaseURL: "https://api.openai.com/v1",
    Timeout: 60 * time.Second, // 60秒超时
}
```

### 使用自定义端点

```go
config := ai_tools.NewAPIConfig("your-key", "https://custom-endpoint.com/v1")
```

## 测试

运行测试：

```bash
cd ai_tools
go test -v
```

## 依赖

- `github.com/otkinlife/go_tools/http_tools` - HTTP 客户端工具

## 许可证

本项目遵循与主项目相同的许可证。

## 贡献

欢迎提交 Issue 和 Pull Request！

## 更新日志

### v1.0.0
- 初始版本发布
- 支持 OpenAI Chat API
- 支持 New-API 服务
- 消息构建器
- 完整的测试覆盖