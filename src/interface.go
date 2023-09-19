package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	clearScreen()
	fmt.Println("...!")

	for {
		fmt.Println("\nQue voulez-vous faire?")
		fmt.Println("1. Star")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Quitter")

		var choix int
		fmt.Print("Choix: ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			fmt.Println("Vous explorez le monde...")
		case 2:
			fmt.Println("Vous ouvrez votre inventaire...")
		case 3:
			fmt.Println("Au revoir!")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}

//Reset de la fenetre 
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
