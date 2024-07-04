package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Pergunta ao usuário qual navegador deseja abrir
	fmt.Println("Qual navegador você deseja abrir?")
	fmt.Println("1. Mozilla Firefox")
	fmt.Println("2. Brave")

	// Ler a escolha do usuário
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Escolha: ")
	choiceStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Erro ao ler a escolha: %v", err)
	}

	// Remove o caractere de nova linha do final da escolha
	choiceStr = strings.TrimSpace(choiceStr)

	// Executa o comando conforme a escolha do usuário
	switch choiceStr {
	case "1":
		openBrowser("firefox")
	case "2":
		openBrowser("brave")
	default:
		fmt.Println("Escolha inválida.")
	}
}

func openBrowser(browser string) {
	var cmd *exec.Cmd

	switch browser {
	case "firefox":
		cmd = exec.Command("cmd", "/C", "start", "firefox")
	case "brave":
		cmd = exec.Command("cmd", "/C", "start", "brave")
	default:
		fmt.Printf("Navegador '%s' não suportado.\n", browser)
		return
	}

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Erro ao abrir o navegador %s: %v", browser, err)
	}
}
