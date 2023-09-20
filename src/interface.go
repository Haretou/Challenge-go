package main

import (
    "fmt"
    "os"
    "os/exec"
)

func AfficherMenu() {
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
            fmt.Println("Bienvenue dans ...")
        case 2:
            fmt.Println("Vous ouvrez vos sacs ...")
        case 3:
            fmt.Println("Au revoir!")
            os.Exit(0)
        default:
            fmt.Println("Choix invalide. Veuillez réessayer.")
        }
    }
}

//Reset de la fenetre 
func ClearAffichage() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}