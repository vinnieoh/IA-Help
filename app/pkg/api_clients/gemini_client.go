package apiclients

import (
    "net/http"
    "io"
    "bytes"
    "encoding/json"
    "os"
)

type GeminiClient struct {
    apiKey string
}

func NewGeminiClient() *GeminiClient {
    return &GeminiClient{
        apiKey: os.Getenv("API_GEMINI"),
    }
}

func (client *GeminiClient) Ask(question string) (string, error) {
    url := "https://api.gemini.com/v1/engines/davinci-codex/completions"
    payload := map[string]interface{}{"prompt": question, "max_tokens": 100}
    body, _ := json.Marshal(payload)

    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+client.apiKey)

    clientHTTP := &http.Client{}
    resp, err := clientHTTP.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var responseMap map[string]interface{}
    json.Unmarshal(responseBody, &responseMap)

    choices := responseMap["choices"].([]interface{})
    if len(choices) > 0 {
        return choices[0].(map[string]interface{})["text"].(string), nil
    }
    return "", nil
}
