package chutes_api

import (
	"bytes"
	"chutesai2api/common/config"
	logger "chutesai2api/common/loggger"
	"fmt"
	"github.com/deanxv/CycleTLS/cycletls"
	"github.com/gin-gonic/gin"
	"math/rand"
	"mime/multipart"
)

const (
	baseURL       = "https://chutes.ai"
	chatEndpoint  = baseURL + "/app/api/chat"
	imageEndpoint = baseURL + "/app/api/invoke-image"
)

func MakeStreamChatRequest(c *gin.Context, client cycletls.CycleTLS, modelId string, jsonData []byte, cookie string) (<-chan cycletls.SSEResponse, error) {

	options := cycletls.Options{
		Timeout: 10 * 60 * 60,
		Body:    string(jsonData),
		Proxy:   config.ProxyUrl, // 在每个请求中设置代理
		Method:  "POST",
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "text/event-stream",
			"Origin":       baseURL,
			"Referer":      baseURL + "/app/chute/" + modelId,
			//"Cookie":       cookie,
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome",
		},
	}

	logger.Debug(c.Request.Context(), fmt.Sprintf("cookie: %v", cookie))

	sseChan, err := client.DoSSE(chatEndpoint, options, "POST")
	if err != nil {
		logger.Errorf(c, "Failed to make stream request: %v", err)
		return nil, fmt.Errorf("Failed to make stream request: %v", err)
	}
	return sseChan, nil
}

type MakeImageReq struct {
	Prompt            string `json:"prompt" form:"prompt" binding:"required"`
	GuidanceScale     string `json:"guidance_scale" form:"guidance_scale"`
	Width             string `json:"width" form:"width"`
	Height            string `json:"height" form:"height"`
	NumInferenceSteps string `json:"num_inference_steps" form:"num_inference_steps"`
}

func MakeImageRequest(c *gin.Context, client cycletls.CycleTLS, requestData MakeImageReq, modelId string) (*cycletls.Response, error) {

	// 创建multipart请求体
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	boundary := "----WebKitFormBoundary" + generateRandomString(16)

	// 设置boundary，与原请求一致
	writer.SetBoundary(boundary)

	// 按顺序添加表单字段
	fields := []struct {
		name  string
		value string
	}{
		{"prompt", requestData.Prompt},
		{"cord", "/generate"},
		{"method", "POST"},
		{"guidance_scale", requestData.GuidanceScale},
		{"width", requestData.Width},
		{"height", requestData.Height},
		{"num_inference_steps", requestData.NumInferenceSteps},
	}

	for _, field := range fields {
		part, err := writer.CreateFormField(field.name)
		if err != nil {
			logger.Errorf(c, "Failed to create form field: %v", err)
			return nil, err
		}
		part.Write([]byte(field.value))
	}
	writer.Close()

	// 配置CycleTLS选项
	options := cycletls.Options{
		Timeout: 10 * 60 * 60,
		Method:  "POST",
		URL:     "https://chutes.ai/app/api/invoke-image/" + modelId,
		Headers: map[string]string{
			"accept":             "*/*",
			"accept-language":    "zh-CN,zh;q=0.9,en;q=0.8",
			"content-type":       writer.FormDataContentType(), // 自动包含正确boundary
			"origin":             "https://chutes.ai",
			"priority":           "u=1, i",
			"referer":            "https://chutes.ai/app/chute/" + modelId,
			"sec-ch-ua":          `"Chromium";v="134", "Not:A-Brand";v="24", "Google Chrome";v="134"`,
			"sec-ch-ua-mobile":   "?0",
			"sec-ch-ua-platform": `"macOS"`,
			"sec-fetch-dest":     "empty",
			"sec-fetch-mode":     "cors",
			"sec-fetch-site":     "same-origin",
			"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36",
		},
		Body: body.String(),
	}

	// 初始化客户端并发送请求
	response, err := client.Do(options.URL, options, "POST")
	if err != nil {
		logger.Errorf(c, "Failed to make image request: %v", err)
		return nil, fmt.Errorf("Failed to make image request: %v", err)
	}

	return &response, nil
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
