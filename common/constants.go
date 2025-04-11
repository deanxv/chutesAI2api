package common

import "time"

var StartTime = time.Now().Unix() // unit: second
var Version = "v1.0.4"            // this hard coding will be replaced automatically when building, no need to manually change

type ModelInfo struct {
	Model     string
	Id        string
	ChuteName string
	MaxTokens int
}

var ModelRegistry = map[string]ModelInfo{
	"deepseek-r1":                {"deepseek-ai/DeepSeek-R1", "de510462-c319-543b-9c67-00bcf807d2a7", "chutes-deepseek-ai-deepseek-r1", 100000},
	"deepseek-v3-0324":           {"deepseek-ai/DeepSeek-V3-0324", "154ad01c-a431-5744-83c8-651215124360", "chutes-deepseek-ai-deepseek-v3-0324", 100000},
	"deepseek-v3":                {"deepseek-ai/DeepSeek-V3", "fa4181ad-0cf1-5531-afdb-bbffbf2ff945", "chutes-deepseek-ai-deepseek-v3", 100000},
	"qwq-32b":                    {"Qwen/QwQ-32B", "2291db94-1463-5bb3-af2b-72c8d254ee9c", "chutes-qwen-qwq-32", 100000},
	"qwen2.5-72b-instruct":       {"Qwen/Qwen2.5-72B-Instruct", "62cc0462-8983-5ef1-8859-92ccf726e235", "chutes-qwen-qwen2-5-72b-instruct", 100000},
	"qwen2.5-coder-32b-instruct": {"Qwen/Qwen2.5-Coder-32B-Instruct", "2cdc73dd-eec4-5e6b-883a-7736bbe12fd8", "chutes-qwen-qwen2-5-coder-32b-instruct", 100000},
	"gemma-3-27b-it":             {"unsloth/gemma-3-27b-it", "16f41e4f-f2ca-5580-a0a6-46727ae4c212", "chutes-unsloth-gemma-3-27b-it", 100000},
	"olympiccoder-32b":           {"open-r1/OlympicCoder-32B", "649c41b3-25d4-59f9-9124-4d834363ec93", "chutes-open-r1-olympiccoder-32b", 100000},
	"reka-flash-3":               {"RekaAI/reka-flash-3", "893971c6-0f01-54fc-8cbe-420d24a20b54", "chutes-rekaai-reka-flash-3", 100000},
	"ui-tars-72b-dpo":            {"bytedance-research/UI-TARS-72B-DPO", "4d4b0ccf-88a5-5991-a741-df4e99dbf7a2", "chutes-bytedance-research-ui-tars-72b-dpo", 100000},
}

var ImageModelRegistry = map[string]ModelInfo{
	"juggernautxl":            {"JuggernautXL", "d8a1a355-d555-5ff0-ac89-166a36cce3ad", "", 0},
	"realistic-vision-v51":    {"stablediffusionapi/realistic-vision-v51", "f3e5e5c5-4271-52bb-a323-baa4cf71755e", "", 0},
	"dreamshaper-xl-v2-turbo": {"Lykon/dreamshaper-xl-v2-turbo", "4aeda079-94d8-5ddf-9236-bb84e23b3751", "", 0},
	"playground-v2.5":         {"Playground-v2.5", "a76dd2c1-2954-532c-9da8-f1796eabad80", "", 0},
	"dreamshaper-xl-1-0":      {"Lykon/dreamshaper-xl-1-0", "7f3dc7d2-cc14-5575-88cb-3d39c99abc3f", "", 0},
	"omnigen-v1":              {"Shitao/OmniGen-v1", "bd63c106-91b9-5937-a52f-bfc7753df770", "", 0},
	"animepasteldream":        {"AnimePastelDream", "791ea6b9-04b6-5f75-9f42-b30259553277", "", 0},
	"psychedelictrees":        {"PsychedelicTrees", "9f064d15-3bd0-5512-a21d-4ace7a4240e4", "", 0},
	"orphic-lora":             {"orphic-lora", "2b3d74b5-fa25-551c-9ea0-240949530eb7", "", 0},
	"constshaper":             {"diagonalge/ConstShaper", "d70bde34-0691-5eb5-af08-bf358b4e24ed", "", 0},
	"flux.1-dev":              {"FLUX.1-dev", "2afe988d-be44-553f-9c85-3caa3d8c0f97", "", 0},
	"flex.1-alpha":            {"ostris/Flex.1-alpha", "c11a7c12-d278-5d30-b5e2-d37d4d16b9da", "", 0},
	"flux.1-schnell":          {"FLUX.1-schnell", "a292d47b-8f0f-5662-b2b0-6f0ebba48031", "", 0},
}

// 通过 model 名称查询的方法
func GetModelInfo(modelName string) (ModelInfo, bool) {
	info, exists := ModelRegistry[modelName]
	return info, exists
}

func GetImageModelInfo(modelName string) (ModelInfo, bool) {
	info, exists := ImageModelRegistry[modelName]
	return info, exists
}

func GetModelList() []string {
	var modelList []string
	for k := range ModelRegistry {
		modelList = append(modelList, k)
	}
	return modelList
}

func GetImageModelList() []string {
	var modelList []string
	for k := range ImageModelRegistry {
		modelList = append(modelList, k)
	}
	return modelList
}
