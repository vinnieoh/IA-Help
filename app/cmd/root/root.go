package root

import (
    "fmt"
    _interface "github.com/vinnieoh/IA-Help/app/interfaces"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "ask",
    Short: "Ask questions to ChatGPT and Gemini",
    Run: func(cmd *cobra.Command, args []string) {
        terminalInterface := _interface.NewTerminalInterface()
        terminalInterface.Run()
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
