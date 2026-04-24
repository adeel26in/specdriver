package main

import (
	"flag"
)

func main() {
	provider_flag := flag.String("provider", "Anthropic", "The provider to choose")
	model_flag := flag.String("model", "", "The model to choose (Auto fallback)")
	specfile_flag := flag.String("specfile", "specdriver_file.md", "The file with the spec")
	output_file := flag.String("output", "specdriver_output.txt", "Name of the output file")
	prompt_caching_flag := flag.Bool("prompt-caching", true, "Enable/Disable prompt caching for Anthropic or OpenAI (Purely here for user control, shouldn't ideally be turned off!)")
	flag.Parse()

	switch *provider_flag {
	case "Anthropic":
		anthropic_func(*specfile_flag, *output_file, *model_flag, *prompt_caching_flag)
	case "OpenAI":
		openai_func(*specfile_flag, *output_file, *model_flag, *prompt_caching_flag)
	case "Gemini":
		google_func(*specfile_flag, *output_file, *model_flag)
	case "Vertex":
		vertex_func(*specfile_flag, *model_flag, *output_file)
	case "Cohere":
		cohere_func(*specfile_flag, *output_file, *model_flag)
	case "Mistral":
		mistral_func(*specfile_flag, *output_file, *model_flag)
	case "xAI", "Grok":
		xAI_func(*specfile_flag, *output_file, *model_flag)
	case "Groq":
		groq_func(*specfile_flag, *output_file, *model_flag)
	case "Deepseek":
		deepseek_func(*specfile_flag, *output_file, *model_flag)
	case "Minimax":
		minimax_func(*specfile_flag, *output_file, *model_flag)
	case "Fireworks":
		fireworks_func(*specfile_flag, *output_file, *model_flag)
	case "Together":
		together_func(*specfile_flag, *output_file, *model_flag)
	case "Deepinfra":
		deepinfra_func(*specfile_flag, *output_file, *model_flag)
	case "Openrouter":
		openrouter_func(*specfile_flag, *output_file, *model_flag)
	case "Perplexity":
		perplexity_func(*specfile_flag, *output_file, *model_flag)
	case "Cerebras":
		cerebras_func(*specfile_flag, *output_file, *model_flag)
	case "NVIDIA NIM", "NIM":
		nvidianim_func(*specfile_flag, *output_file, *model_flag)
	case "Cloudflare", "Cloudflare Workers":
		cloudflare_func(*specfile_flag, *output_file, *model_flag)
	case "Ollama":
		ollama_func(*specfile_flag, *output_file, *model_flag)
	case "vLLM":
		vLLM_func(*specfile_flag, *output_file, *model_flag)
	}
}
