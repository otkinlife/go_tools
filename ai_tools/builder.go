package ai_tools

// MessageBuilder helps build conversation messages
type MessageBuilder struct {
	messages []ChatMessage
}

// NewMessageBuilder creates a new message builder
func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{
		messages: make([]ChatMessage, 0),
	}
}

// System adds a system message
func (mb *MessageBuilder) System(content string) *MessageBuilder {
	mb.messages = append(mb.messages, ChatMessage{
		Role:    "system",
		Content: content,
	})
	return mb
}

// User adds a user message
func (mb *MessageBuilder) User(content string) *MessageBuilder {
	mb.messages = append(mb.messages, ChatMessage{
		Role:    "user",
		Content: content,
	})
	return mb
}

// Assistant adds an assistant message
func (mb *MessageBuilder) Assistant(content string) *MessageBuilder {
	mb.messages = append(mb.messages, ChatMessage{
		Role:    "assistant",
		Content: content,
	})
	return mb
}

// AddMessage adds a custom message
func (mb *MessageBuilder) AddMessage(role, content string) *MessageBuilder {
	mb.messages = append(mb.messages, ChatMessage{
		Role:    role,
		Content: content,
	})
	return mb
}

// Build returns the built messages
func (mb *MessageBuilder) Build() []ChatMessage {
	return mb.messages
}

// Clear clears all messages
func (mb *MessageBuilder) Clear() *MessageBuilder {
	mb.messages = make([]ChatMessage, 0)
	return mb
}

// Count returns the number of messages
func (mb *MessageBuilder) Count() int {
	return len(mb.messages)
}

// GetMessages returns a copy of current messages
func (mb *MessageBuilder) GetMessages() []ChatMessage {
	result := make([]ChatMessage, len(mb.messages))
	copy(result, mb.messages)
	return result
}
