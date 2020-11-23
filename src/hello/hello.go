package main

import (
	"fmt"
)

func main() {
	name := "Cristiano"
	version := 1.1
	fmt.Println("Olá Senhor", name)
	fmt.Println("Este programa esta na versão:", version)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var command int
	fmt.Scan(&command)

	fmt.Println("Opção escolhida:", command)
}
