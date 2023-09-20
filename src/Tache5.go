package main

import (
	"fmt"
)

// Structure pour représenter un élément de l'inventaire
type InventoryItem struct {
	Name  string
	Usage string
}

// Fonction pour afficher l'inventaire
func accessInventory(inventory []InventoryItem) {
	fmt.Println("Contenu de l'inventaire :")
	for i, item := range inventory {
		fmt.Printf("%d. %s - %s\n", i+1, item.Name, item.Usage)
	}

	// Option pour retourner au menu précédent
	fmt.Println("0. Retour au menu précédent")
}

func main() {
	// Exemple d'inventaire
	inventory := []InventoryItem{
		{"Potion de vie", "Rend 50 points de vie"},
		{"Épée en fer", "Arme de base"},
		{"Livre de sort", "Lance un sort de Givre"},
	}

	for {
		// Affichage du menu principal
		fmt.Println("Menu principal :")
		fmt.Println("1. Play")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Quitter")

		var choice int
		fmt.Print("Choisissez une option : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Start.")
		case 2:
			// Afficher l'inventaire en appelant la fonction accessInventory
			accessInventory(inventory)

			// Attendre une entrée du joueur pour revenir au menu principal
			fmt.Println("Appuyez sur Entrée pour continuer...")
			fmt.Scanln()
		case 3:
			// Quitter le programme
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Option invalide. Veuillez réessayer.")
		}
	}
}
