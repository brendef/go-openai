package openai

type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choices []struct {
	Message Message `json:"message"`
	Text    string  `json:"text"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}

type ChatResponse struct {
	Choices Choices `json:"choices"`
	Usage   Usage   `json:"usage"`
	Error   Error   `json:"error"`
}

type Error struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}
