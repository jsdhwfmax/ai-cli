package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Config represents the universal configuration for any OpenAI-compatible API
type Config struct {
	ApiBaseURL string `json:"api_base_url"` // The API endpoint
	APIKey     string `json:"api_key"`      // The secret key
	Model      string `json:"model"`        // The model name
}

// OpenAI-compatible payload structures
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestPayload struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

type ResponsePayload struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

// Helper to get the config file path (~/.ai-cli.json)
func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".ai-cli.json"), nil
}

// Helper to load or initialize the config file
func loadConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory")
	}

	// If config file doesn't exist, create a template (Default to DeepSeek as an example) and exit
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfig := Config{
			ApiBaseURL: "https://api.deepseek.com/chat/completions",
			APIKey:     "YOUR_API_KEY_HERE",
			Model:      "deepseek-chat",
		}
		data, _ := json.MarshalIndent(defaultConfig, "", "  ")
		os.WriteFile(configPath, data, 0600)

		fmt.Println("🚀 Welcome to AI-CLI!")
		fmt.Printf("A configuration file has been created at: \033[93m%s\033[0m\n", configPath)
		fmt.Println("Please open this file, enter your API Key, and run the command again.")
		fmt.Println("You can also change the 'api_base_url' and 'model' to use any other AI provider (OpenAI, Qwen, Kimi, Ollama, etc.)")
		os.Exit(0)
	}

	// Read existing config
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file")
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("config file is corrupted. Please check the JSON format")
	}

	if config.APIKey == "" || config.APIKey == "YOUR_API_KEY_HERE" {
		return nil, fmt.Errorf("API Key is missing. Please update %s", configPath)
	}
	if config.ApiBaseURL == "" || config.Model == "" {
		return nil, fmt.Errorf("API Base URL or Model is missing in the config file")
	}

	return &config, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ai '<your request>'")
		fmt.Println("Example: ai 'undo my last git commit'")
		os.Exit(1)
	}
	userInput := strings.Join(os.Args[1:], " ")

	// Load configuration
	config, err := loadConfig()
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("🤖 AI Engine [%s] is thinking...\n\n", config.Model)

	// The system prompt
	systemPrompt := `You are an ultimate geek master of Linux, macOS, and developer tools.
Your task is to convert the user's natural language request into a single-line terminal command.
CRITICAL RULES:
1. ONLY output the exact command. Do NOT output any explanations.
2. Do NOT use Markdown formatting or wrap the text in backticks.
3. Return plain text only.`

	payload := RequestPayload{
		Model:       config.Model,
		Temperature: 0.1,
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userInput},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("❌ Internal error: Data packing failed.")
		os.Exit(1)
	}

	// Send HTTP Request using the custom URL from config
	req, err := http.NewRequest("POST", config.ApiBaseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("❌ Internal error: Request creation failed.")
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ Network error: Please check your internet connection or the API Base URL.")
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("❌ Error: Failed to read API response.")
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("❌ API Request Failed! (HTTP Status %d)\n", resp.StatusCode)
		fmt.Printf("Provider response: %s\n", string(body))
		os.Exit(1)
	}

	var responseData ResponsePayload
	if err := json.Unmarshal(body, &responseData); err != nil || len(responseData.Choices) == 0 {
		fmt.Println("❌ Error: Failed to parse API JSON response. Ensure the provider uses OpenAI-compatible format.")
		os.Exit(1)
	}

	command := strings.TrimSpace(responseData.Choices[0].Message.Content)

	fmt.Println("💡 Done! Here is your command:")
	fmt.Println("----------------------------------------")
	fmt.Printf("\033[92m%s\033[0m\n", command)
	fmt.Println("----------------------------------------")
}