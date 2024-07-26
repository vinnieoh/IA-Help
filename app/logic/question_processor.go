package logic

import (
    "github.com/vinnieoh/IA-Help/app/pkg/api_clients"
)

type QuestionProcessor struct {
    chatgptClient *apiclients.ChatGPTClient
    geminiClient  *apiclients.GeminiClient
}

func NewQuestionProcessor() *QuestionProcessor {
    return &QuestionProcessor{
        chatgptClient: apiclients.NewChatGPTClient(),
        geminiClient:  apiclients.NewGeminiClient(),
    }
}

func (qp *QuestionProcessor) ProcessQuestion(question string) (string, error) {
    // Call ChatGPT API
    chatGPTResponse, err := qp.chatgptClient.Ask(question)
    if err != nil {
        return "", err
    }
    // Call Gemini API
    geminiResponse, err := qp.geminiClient.Ask(question)
    if err != nil {
        return "", err
    }
    // Combine and return responses
    return chatGPTResponse + "\n" + geminiResponse, nil
}