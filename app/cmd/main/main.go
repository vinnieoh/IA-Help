package main

import (
    "log"

    "github.com/vinnieoh/IA-Help/app/cmd/root" 
    
    "github.com/joho/godotenv"
    
)

func main() {
    // Carrega variáveis de ambiente do arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
    }

    // Executa o comando raiz
    root.Execute()
}