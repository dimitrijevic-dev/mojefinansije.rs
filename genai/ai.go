package genai

import (
	"context"
	"fon/config"
	"fon/persistence"
	"log"
	"os"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

type PromptRequest struct {
	Prompt string `json:"prompt"`
}

var client openai.Client

func Start() {
	client = openai.NewClient(option.WithAPIKey(config.GetEnv("OPENAI_KEY")))
}

func LessonAsk(prompt string, id int) string {
	p, err := os.ReadFile("genai/lesson.txt")
	if err != nil {
		log.Println(err)
	}

	str := string(p) + persistence.GetLessonByID(id).Body + " USER QUESTION: " + prompt

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(str),
		},
		Model: openai.ChatModelGPT5,
	})
	if err != nil {
		log.Println("Error: Chat completion failed!")
	}

	return chatCompletion.Choices[0].Message.Content
}
