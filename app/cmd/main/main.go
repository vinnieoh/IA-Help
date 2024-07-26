package main

import (
    "log"
    
    "github.com/joho/godotenv"
    "github.com/vinnieoh/IA-Help/app/cmd/root" 
)

func main() {
    
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
    }

    // Executa o comando raiz
    cmd.Execute()
}
