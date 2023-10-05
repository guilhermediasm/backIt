package schemas

type ChatGPTModel string
type ChatGPTModelRole string

const (
	GPT35Turbo ChatGPTModel = "gpt-3.5-turbo"

	// Deprecated: Use gpt-3.5-turbo-0613 instead, model will discontinue on 09/13/2023
	GPT35Turbo0301 ChatGPTModel = "gpt-3.5-turbo-0301"

	GPT35Turbo0613    ChatGPTModel = "gpt-3.5-turbo-0613"
	GPT35Turbo16k     ChatGPTModel = "gpt-3.5-turbo-16k"
	GPT35Turbo16k0613 ChatGPTModel = "gpt-3.5-turbo-16k-0613"
	GPT4              ChatGPTModel = "gpt-4"

	// Deprecated: Use gpt-4-0613 instead, model will discontinue on 09/13/2023
	GPT4_0314 ChatGPTModel = "gpt-4-0314"

	GPT4_0613 ChatGPTModel = "gpt-4-0613"
	GPT4_32k  ChatGPTModel = "gpt-4-32k"

	// Deprecated: Use gpt-4-32k-0613 instead, model will discontinue on 09/13/2023
	GPT4_32k_0314 ChatGPTModel = "gpt-4-32k-0314"

	GPT4_32k_0613 ChatGPTModel = "gpt-4-32k-0613"
)

const (
	ChatGPTModelRoleUser      ChatGPTModelRole = "user"
	ChatGPTModelRoleSystem    ChatGPTModelRole = "system"
	ChatGPTModelRoleAssistant ChatGPTModelRole = "assistant"
)

type ChatCompletionRequest struct {
	// (Required)
	// ID of the model to use.
	Model ChatGPTModel `json:"model"`

	// Required
	// The messages to generate chat completions for
	Messages []ChatMessage `json:"messages"`

	// (Optional - default: 1)
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.
	// We generally recommend altering this or top_p but not both.
	Temperature float64 `json:"temperature,omitempty"`

	// (Optional - default: 1)
	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	// We generally recommend altering this or temperature but not both.
	Top_P float64 `json:"top_p,omitempty"`

	// (Optional - default: 1)
	// How many chat completion choices to generate for each input message.
	N int `json:"n,omitempty"`

	// (Optional - default: infinite)
	// The maximum number of tokens allowed for the generated answer. By default,
	// the number of tokens the model can return will be (4096 - prompt tokens).
	MaxTokens int `json:"max_tokens,omitempty"`

	// (Optional - default: 0)
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far,
	// increasing the model's likelihood to talk about new topics.
	PresencePenalty float64 `json:"presence_penalty,omitempty"`

	// (Optional - default: 0)
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far,
	// decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`

	// (Optional)
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse
	User string `json:"user,omitempty"`
}

type ChatMessage struct {
	Role    ChatGPTModelRole `json:"role"`
	Content string           `json:"content"`
}

type ChatResponse struct {
	ID        string               `json:"id"`
	Object    string               `json:"object"`
	CreatedAt int64                `json:"created_at"`
	Choices   []ChatResponseChoice `json:"choices"`
	Usage     ChatResponseUsage    `json:"usage"`
}

type ChatResponseChoice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

type ChatResponseUsage struct {
	Prompt_Tokens     int `json:"prompt_tokens"`
	Completion_Tokens int `json:"completion_tokens"`
	Total_Tokens      int `json:"total_tokens"`
}
