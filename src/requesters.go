package main

import (
	"context"
	"fmt"
	"os"

	"github.com/zendev-sh/goai"
	"github.com/zendev-sh/goai/provider/anthropic"
	"github.com/zendev-sh/goai/provider/cerebras"
	"github.com/zendev-sh/goai/provider/cloudflare"
	"github.com/zendev-sh/goai/provider/cohere"
	"github.com/zendev-sh/goai/provider/deepinfra"
	"github.com/zendev-sh/goai/provider/deepseek"
	"github.com/zendev-sh/goai/provider/fireworks"
	"github.com/zendev-sh/goai/provider/google"
	"github.com/zendev-sh/goai/provider/groq"
	"github.com/zendev-sh/goai/provider/minimax"
	"github.com/zendev-sh/goai/provider/mistral"
	"github.com/zendev-sh/goai/provider/nvidia"
	"github.com/zendev-sh/goai/provider/ollama"
	"github.com/zendev-sh/goai/provider/openai"
	"github.com/zendev-sh/goai/provider/openrouter"
	"github.com/zendev-sh/goai/provider/perplexity"
	"github.com/zendev-sh/goai/provider/together"
	"github.com/zendev-sh/goai/provider/vertex"
	"github.com/zendev-sh/goai/provider/vllm"
	"github.com/zendev-sh/goai/provider/xai"
)

func anthropic_func(specfile, output, model_input string, prompt_caching bool) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "claude-sonnet-4-6-20260310"
	}
	model := anthropic.Chat(model_input, anthropic.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
		goai.WithPromptCaching(prompt_caching),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func openai_func(specfile, output, model_input string, prompt_caching bool) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "gpt-5"
	}
	model := openai.Chat(model_input, openai.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
		goai.WithPromptCaching(prompt_caching),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func google_func(specfile, output, model_input string) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "gemini-2.5-pro"
	}
	model := google.Chat(model_input, google.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func vertex_func(specfile, output, model_input string) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "gemini-2.5-pro"
	}
	model := vertex.Chat(model_input, vertex.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func cohere_func(specfile, output, model_input string) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "command-r-plus"
	}
	model := cohere.Chat(model_input, cohere.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func mistral_func(specfile, output, model_input string) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "mistral-large-latest"
	}
	model := mistral.Chat(model_input, mistral.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func xAI_func(specfile, output, model_input string) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "grok-3"
	}
	model := xai.Chat(model_input, xai.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func groq_func(specfile, output, model_input string) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "llama-3.3-70b-versatile"
	}
	model := groq.Chat(model_input, groq.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func deepseek_func(specfile, output, model_input string) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "deepseek-chat"
	}
	model := deepseek.Chat(model_input, deepseek.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func minimax_func(specfile, output, model_input string) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "MiniMax-M2"
	}
	model := minimax.Chat(model_input, minimax.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func fireworks_func(specfile, output, model_input string) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "accounts/fireworks/models/llama-v3p3-70b-instruct"
	}
	model := fireworks.Chat(model_input, fireworks.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func together_func(specfile, output, model_input string) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "meta-llama/Llama-3.3-70B-Instruct-Turbo"
	}
	model := together.Chat(model_input, together.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func deepinfra_func(specfile, output, model_input string) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "meta-llama/Llama-3.3-70B-Instruct"
	}
	model := deepinfra.Chat(model_input, deepinfra.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func openrouter_func(specfile, output, model_input string) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "meta-llama/llama-3.3-70b-instruct"
	}
	model := openrouter.Chat(model_input, openrouter.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func perplexity_func(specfile, output, model_input string) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "sonar"
	}
	model := perplexity.Chat(model_input, perplexity.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func cerebras_func(specfile, output, model_input string) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "llama-3.3-70b"
	}
	model := cerebras.Chat(model_input, cerebras.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func nvidianim_func(specfile, output, model_input string) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "nvidia/llama-3.3-70b-instruct"
	}
	model := nvidia.Chat(model_input, nvidia.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func cloudflare_func(specfile, output, model_input string) {
	api_key := config()
	if api_key == "" {
		fmt.Println("API key not set, Unable to proceed")
		return
	}
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	}
	model := cloudflare.Chat(model_input, cloudflare.WithAPIKey(api_key))

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func ollama_func(specfile, output, model_input string) {
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "llama3"
	}
	model := ollama.Chat(model_input)

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

func vLLM_func(specfile, output, model_input string) {
	file, err := os.ReadFile(specfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if model_input == "" {
		model_input = "meta-llama/Llama-3-8b"
	}
	model := vllm.Chat(model_input)

	result, err := goai.StreamText(context.Background(), model,
		goai.WithPrompt(string(file)),
		goai.WithSystem(prompt),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer(output, result)
}

// writer func (End of requesters!)

func writer(output string, result *goai.TextStream) {
	output_file, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer output_file.Close()
	for text := range result.TextStream() {
		fmt.Fprint(output_file, text)
	}
}
