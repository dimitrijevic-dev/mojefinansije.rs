package genai

import (
	"context"
	"fon/config"
	"fon/persistence"
	"log"
	"os"
	"strings"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

type PromptRequest struct {
	Prompt string `json:"prompt"`
}

type Flashcard struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
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
	return ask(str)
}

func ask(prompt string) string {
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: openai.ChatModelGPT5,
	})
	if err != nil {
		log.Println("Error: Chat completion failed!")
	}

	return chatCompletion.Choices[0].Message.Content
}

func GenerateFlashcards(id int) []Flashcard {
	p, err := os.ReadFile("genai/prompt.txt")
	if err != nil {
		log.Println(err)
	}

	str := string(p) + persistence.GetLessonByID(id).Body
	lines := strings.Split(ask(str), "\n")
	var arr []Flashcard
	for _, line := range lines {
		splitLine := strings.Split(line, "|")
		arr = append(arr, Flashcard{Question: splitLine[0], Answer: splitLine[1]})
	}
	return arr
}
