package main

import (
    "fmt"
    "Player"
)
func takePot() {
    if Player.Health < Player.MaxHealth {
        Player.Health += 50 
        fmt.Println("Vous vous etes soigné !", Player.Health)
    } else {
        fmt.Println("Votre vie est déja au maximum !")
    }
}
