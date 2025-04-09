package model

import (
	"encoding/json"
	"strings"
)

type OpenAIChatCompletionRequest struct {
	Model    string              `json:"model"`
	Stream   bool                `json:"stream"`
	Messages []OpenAIChatMessage `json:"messages"`
}

type SessionState struct {
	Models           []string `json:"models"`
	Layers           int      `json:"layers"`
	Answer           string   `json:"answer"`
	AnswerIsFinished bool     `json:"answer_is_finished"`
}
type OpenAIChatMessage struct {
	Role    string      `json:"role"`
	Content interface{} `json:"content"`
}

func (r *OpenAIChatCompletionRequest) RemoveEmptyContentMessages() *OpenAIChatCompletionRequest {
	if r == nil || len(r.Messages) == 0 {
		return r
	}

	var filteredMessages []OpenAIChatMessage
	for _, msg := range r.Messages {
		// Check if content is nil
		if msg.Content == nil {
			continue
		}

		// Check if content is an empty string
		if strContent, ok := msg.Content.(string); ok && strContent == "" {
			continue
		}

		// Check if content is an empty slice
		if sliceContent, ok := msg.Content.([]interface{}); ok && len(sliceContent) == 0 {
			continue
		}

		// If we get here, the content is not empty
		filteredMessages = append(filteredMessages, msg)
	}

	r.Messages = filteredMessages
	return r
}

func (r *OpenAIChatCompletionRequest) AddMessage(message OpenAIChatMessage) {
	r.Messages = append([]OpenAIChatMessage{message}, r.Messages...)
}

func (r *OpenAIChatCompletionRequest) PrependMessagesFromJSON(jsonString string) error {
	var newMessages []OpenAIChatMessage
	err := json.Unmarshal([]byte(jsonString), &newMessages)
	if err != nil {
		return err
	}

	// 查找最后一个 system role 的索引
	var insertIndex int
	for i := len(r.Messages) - 1; i >= 0; i-- {
		if r.Messages[i].Role == "system" {
			insertIndex = i + 1
			break
		}
	}

	// 将 newMessages 插入到找到的索引后面
	r.Messages = append(r.Messages[:insertIndex], append(newMessages, r.Messages[insertIndex:]...)...)
	return nil
}

func (r *OpenAIChatCompletionRequest) SystemMessagesProcess(model string) {
	if r.Messages == nil {
		return
	}

	if model == "deep-seek-r1" {
		for i := range r.Messages {
			if r.Messages[i].Role == "system" {
				r.Messages[i].Role = "user"
			}

		}
	}
}

func (r *OpenAIChatCompletionRequest) FilterUserMessage() {
	if r.Messages == nil {
		return
	}

	// 返回最后一个role为user的元素
	for i := len(r.Messages) - 1; i >= 0; i-- {
		if r.Messages[i].Role == "user" {
			r.Messages = r.Messages[i:]
			break
		}
	}
}

type OpenAIErrorResponse struct {
	OpenAIError OpenAIError `json:"error"`
}

type OpenAIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    string `json:"code"`
}

type OpenAIChatCompletionResponse struct {
	ID                string         `json:"id"`
	Object            string         `json:"object"`
	Created           int64          `json:"created"`
	Model             string         `json:"model"`
	Choices           []OpenAIChoice `json:"choices"`
	Usage             OpenAIUsage    `json:"usage"`
	SystemFingerprint *string        `json:"system_fingerprint"`
}

type OpenAIChoice struct {
	Index        int           `json:"index"`
	Message      OpenAIMessage `json:"message"`
	LogProbs     *string       `json:"logprobs"`
	FinishReason *string       `json:"finish_reason"`
	Delta        OpenAIDelta   `json:"delta"`
}

type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type OpenAIDelta struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type OpenAIImagesGenerationRequest struct {
	Model          string `json:"model"`
	Prompt         string `json:"prompt"`
	ResponseFormat string `json:"response_format"`
	Image          string `json:"image"`
	Size           string `json:"size"`
}

type OpenAIImagesGenerationResponse struct {
	Created int64                                 `json:"created"`
	Data    []*OpenAIImagesGenerationDataResponse `json:"data"`
}

type OpenAIImagesGenerationDataResponse struct {
	//URL           string `json:"url"`
	//RevisedPrompt string `json:"revised_prompt"`
	B64Json string `json:"b64_json"`
}

type OpenAIGPT4VImagesReq struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	ImageURL struct {
		URL string `json:"url"`
	} `json:"image_url"`
}

type GetUserContent interface {
	GetUserContent() []string
}

type OpenAIModerationRequest struct {
	Input string `json:"input"`
}

type OpenAIModerationResponse struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Results []struct {
		Flagged        bool               `json:"flagged"`
		Categories     map[string]bool    `json:"categories"`
		CategoryScores map[string]float64 `json:"category_scores"`
	} `json:"results"`
}

type OpenaiModelResponse struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	//Created time.Time `json:"created"`
	//OwnedBy string    `json:"owned_by"`
}

// ModelList represents a list of models.
type OpenaiModelListResponse struct {
	Object string                `json:"object"`
	Data   []OpenaiModelResponse `json:"data"`
}

func (r *OpenAIChatCompletionRequest) GetUserContent() []string {
	var userContent []string

	for i := len(r.Messages) - 1; i >= 0; i-- {
		if r.Messages[i].Role == "user" {
			switch contentObj := r.Messages[i].Content.(type) {
			case string:
				userContent = append(userContent, contentObj)
			}
			break
		}
	}

	return userContent
}
func (r *OpenAIChatCompletionRequest) GetPreviousMessagePair() (string, bool, error) {
	messages := r.Messages
	if len(messages) < 3 {
		return "", false, nil
	}

	if len(messages) > 0 && messages[len(messages)-1].Role != "user" {
		return "", false, nil
	}

	for i := len(messages) - 2; i > 0; i-- {
		if messages[i].Role == "assistant" {
			if messages[i-1].Role == "user" {
				// 深拷贝消息对象避免污染原始数据
				prevPair := []OpenAIChatMessage{
					messages[i-1], // 用户消息
					messages[i],   // 助手消息
				}

				jsonData, err := json.Marshal(prevPair)
				if err != nil {
					return "", false, err
				}

				// 移除JSON字符串中的转义字符
				cleaned := strings.NewReplacer(
					`\n`, "",
					`\t`, "",
					`\r`, "",
				).Replace(string(jsonData))

				return cleaned, true, nil
			}
		}
	}
	return "", false, nil
}
