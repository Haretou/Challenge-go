package main

import (
	"fmt"
)

// Fonction de combat
func combat(Player, Enemy) {
    // Boucle de combat
    for Player.Health > 0 && Enemy.Health > 0 { // tant que aucuns des deux est mort 
        // Player attaque enemy
		if DamagePlayer > 0 {
			HealthEnemy = DamagePlayer - DefEnemy // vie -= (degats subit - def)
			break
		}
		if DamageEnemy > 0 {
			HealthPlayer = DamageEnemy - DefPlayer // vie -= (degats subit - def)
			break
		}
	}
		// Vérifier si Enemy est KO
		if Enemy.Health <= 0 {
			fmt.Printf("%s win %s!\n", Player.Name, Enemy.Name)
			break
		}
		// Vérifier si Player est KO
		if Player.Health <= 0 {
			fmt.Print("%s win %s!\n", Enemy.Name, Player.Name)
			break
		}
}
