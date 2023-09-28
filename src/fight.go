package main

import (
	"fmt"
	"Player"
	"Enemy"
)

func fight() {
	fmt.Println(Player.Health, Enemy.Health)
	fmt.Print(("Premier tours ..."))
	if Player.Health > 0 && Enemy.Health > 0 {
		// tours de Player 
		Enemy.Health -= Player.attaque
		// tour de Enemy
		Player.Health -= Enemy.attaque
		fmt.Print("Fin du tours")
	} else {
		fmt.Print("Combat terminé")
	}
}
