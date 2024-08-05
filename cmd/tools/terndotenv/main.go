package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations/",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Erro ao executar o comando: %v\nSaída de erro: %s", err, stderr.String())
		panic(err)
	}
	fmt.Printf("Saída do comando:\n%s", out.String())
}
