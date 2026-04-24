specdriver
Spec Driven Development (SDD) made easy!

Supported providers (20):

1. Anthropic
2. OpenAI
3. Gemini
4. Vertex
5. Cohere
6. Mistral
7. xAI
8. Groq
9. Deepseek
10. Minimax
11. Fireworks
12. Together
13. Deepinfra
14. OpenRouter
15. Perplexity
16. Cerebras
17. NVIDIA NIM
18. Cloudflare Workers
19. Ollama
20. vLLM

Config?: Make a config file in . or /.config or for Windows users /AppData/Local and name the file specdriver.yaml

How to get started?: Install the binary for your system or build from source

Make sure you have Go installed: go.dev

git clone https://github.com/adeel26in/specdriver.git

cd specdriver/src

CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -buildid="

sudo mv specdriver /usr/bin

specdriver --help
