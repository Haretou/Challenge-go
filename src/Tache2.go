package main

import "fmt"

// Définir la structure Personnage
type Personnage struct {
    Name            string
    Classe          string
    Level           int
    MaxHealth       int
    Health          int
    Inventory      []string
}

func main2() {
    // Création du personnage
    Player := Personnage{
        Name:             "TonPere",
        Classe:          "Guerrier",
        Level:          10,
        MaxHealth:  100,
        Health: 100,
        Inventory:      []string{"Épée", "Armure", "Potion de vie "},
    }
    // Afficher les informations de Player
    fmt.Printf("Name: %s\n", Player.Name)
    fmt.Printf("Classe: %s\n", Player.Classe)
    fmt.Printf("Niveau: %d\n", Player.Level)
    fmt.Printf("Points de Vie Max: %d\n", Player.MaxHealth)
    fmt.Printf("Points de Vie Actuels: %d\n", Player.Health)
    fmt.Printf("Inventaire: %v\n", Player.Inventory)
}