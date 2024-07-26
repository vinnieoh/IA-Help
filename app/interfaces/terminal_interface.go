package _interface

import (
    "bufio"
    "fmt"
    "os"
    "github.com/vinnieoh/IA-Help/app/logic"
)

type TerminalInterface struct {
    processor *logic.QuestionProcessor
}

func NewTerminalInterface() *TerminalInterface {
    return &TerminalInterface{
        processor: logic.NewQuestionProcessor(),
    }
}

func (ti *TerminalInterface) Run() {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Enter your question: ")
        question, _ := reader.ReadString('\n')
        response, err := ti.processor.ProcessQuestion(question)
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }
        fmt.Println("Response:", response)
    }
}