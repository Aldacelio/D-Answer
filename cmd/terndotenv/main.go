package main

import (
    // "fmt"
    // "os"
    "os/exec"

    "github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        panic(err)
    }

    // Imprime o diretório de trabalho atual
    // dir, err := os.Getwd()
    // if err != nil {
    //     panic(err)
    // }
    // fmt.Println("Diretório de trabalho atual:", dir)

    cmd := exec.Command(
        "tern", 
        "migrate", 
        "--migrations", 
        "./internal/store/pgstore/migrations", 
        "--config", 
        "./internal/store/pgstore/migrations/tern.conf",
    )
    if err := cmd.Run(); err != nil {
        panic(err)
    }
    
    // Imprime o comando para depuração
    // fmt.Println("Executando comando:", cmd.String())

    // Captura a saída e o erro do comando
    // output, err := cmd.CombinedOutput()
    // if err != nil {
    //     // Imprime a saída e o erro para depuração
    //     fmt.Printf("Erro ao executar comando: %s\n", err)
    //     fmt.Printf("Saída do comando: %s\n", output)
    //     panic(err)
    // }
    
    // Se o comando for bem-sucedido, imprime a saída
    // fmt.Println("Saída do comando:", string(output))
}
