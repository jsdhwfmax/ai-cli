# 🤖 AI-CLI: Your Universal Terminal Telepath

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-green.svg)

> Stop Googling terminal commands. Just ask your terminal directly.

**AI-CLI** is an ultra-lightweight, lightning-fast command-line assistant. You type natural language, and it instantly generates the exact Linux/macOS/Git/Docker command you need. 

Unlike other tools that lock you into one ecosystem, **AI-CLI is 100% model-agnostic**. You can bring your own API key from ANY provider that supports the OpenAI-compatible format (DeepSeek, Gemini, OpenAI, Claude, or even local Ollama models).

## ✨ Features

- ⚡ **Lightning Fast**: Built with pure Go. Returns commands in milliseconds.
- 🎯 **Zero Bullshit**: Outputs ONLY the exact command. No Markdown, no chatty explanations. Ready to copy/execute.
- 🌍 **Universal AI Support**: Connects to ANY OpenAI-compatible API endpoint.
- 📦 **Zero Dependencies**: Compiles into a single binary. No Python, no Node.js needed.

---

## 🚀 Installation

Make sure you have [Go](https://go.dev/) installed on your machine.

```bash
# 1. Clone the repo
git clone [https://github.com/jsdhwfmaX/ai-cli.git](https://github.com/jsdhwfmaX/ai-cli.git)
cd ai-cli

# 2. Build the binary
go build -o ai main.go

# 3. Move to your system path (macOS/Linux)
sudo mv ai /usr/local/bin/
` ``

---

## ⚙️ Configuration

Run the command for the first time to generate the config file:
```bash
ai "hello"
` ``
It will create a configuration file at `~/.ai-cli.json` in your home directory. Open it and replace the content with your preferred AI provider:

### 🔹 Option 1: DeepSeek (Default)
```json
{
  "api_base_url": "[https://api.deepseek.com/chat/completions](https://api.deepseek.com/chat/completions)",
  "api_key": "sk-your-deepseek-key-here",
  "model": "deepseek-chat"
}
` ``

### 🔹 Option 2: Gemini
```json
{
  "api_base_url": "https://generativelanguage.googleapis.com/v1beta/openai/chat/completions",
  "api_key": "AIzaSy-your-gemini-key-here",
  "model": "gemini-2.5-flash"
}
` ``

### 🔹 Option 3: OpenAI
```json
{
  "api_base_url": "[https://api.openai.com/v1/chat/completions](https://api.openai.com/v1/chat/completions)",
  "api_key": "sk-your-openai-key-here",
  "model": "gpt-4o-mini"
}
` ``

> **🔥 For Advanced Users:** You can point `ai-cli` to your local Ollama instance (e.g., `http://localhost:11434/v1/chat/completions`) or any other custom endpoints!

---

## 💡 Usage

Just type `ai` followed by your request in **any language**:

```bash
$ ai "kill the process running on port 8080"
🤖 AI Engine [deepseek-chat] is thinking...

💡 Done! Here is your command:
----------------------------------------
lsof -ti:8080 | xargs kill -9
----------------------------------------
` ``

```bash
$ ai "convert all mp4 files in this directory to mp3"
🤖 AI Engine [gemini-2.5-flash] is thinking...

💡 Done! Here is your command:
----------------------------------------
for i in *.mp4; do ffmpeg -i "$i" "${i%.*}.mp3"; done
----------------------------------------
` ``

## 🤝 Contributing
Issues and Pull Requests are deeply appreciated! Let's make terminal life easier together.
