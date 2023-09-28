package main

import (
    "fmt"
)

func main2() {
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
        fmt.Println("4. Marchand")

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
        case 4:
            fmt.Println("Que voulez vous acheter ?\n potion de vie : Free \n potion de poison : 5 Po \n épée en fer \n Quitter")
        default:
            fmt.Println("Option invalide. Veuillez réessayer.")
        }
    }
}
