package main

import (
	"fmt"
	"time"
)

// Structure pour représenter un élément de l'inventaire
type InventoryItem struct {
    Name  string
    Usage string
}

// Fonction pour afficher l'inventaire
func accessInventory(inventory []InventoryItem) {
    fmt.Println("Contenu de l'inventaire")
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
        fmt.Println("Menu principal") 
        fmt.Println("1. Play")
        fmt.Println("2. Inventaire")
        fmt.Println("3. Marchand")
        fmt.Println("4. Quitter")

        var choice int
        fmt.Print("Choisissez une option : ")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            Lors()
        case 2:
            // Afficher l'inventaire en appelant la fonction accessInventory
            accessInventory(inventory)

            // Attendre une entrée du joueur pour revenir au menu principal
            fmt.Println("Appuyez sur Entrée pour continuer...")
            fmt.Scanln()
        case 3:
            fmt.Println("Que voulez vous acheter ? \n 1.Potion de vie : Gratuit \n 2.Potion de poison : 5 P.o \n 3.épée en fer : 20 P.o \n 4.Quitter ")
        case 4:
            // Quitter le programme
            fmt.Println("Au revoir !")
            return
        default:
            fmt.Println("Option invalide. Veuillez réessayer.")
        }
    }
}
func Lors() { 
    fmt.Println("Bienvenue dans le monde Foresta")
    time.Sleep(time.Second * 2)
    fmt.Println("Le diable ????? envoilla ses soldats à la conquête du pays." )
    time.Sleep(time.Second * 2)
    fmt.Println("Un jour un jeune homme aux nom ???? entendit une voie qui lui dit")
    time.Sleep(time.Second * 2)
    fmt.Println("tu es l'héritier qui sauvera le pays et triomphera du diable. ")
    time.Sleep(time.Second * 2)
    fmt.Println("Il prend l'équipement et il va sauver son pays et tuer le diable.")
    fmt.Scanln()
}