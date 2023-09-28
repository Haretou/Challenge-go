package Combat

import (
	"fmt"
	"math/rand"
	"time"
)

// Character représente un personnage dans le jeu.
type Character struct {
	name      string
	class     string
	hp        int
	attack    int
	defense   int
	inventory map[string]int // Inventaire avec des objets et leur quantité
}

// Attaquer effectue une attaque contre un adversaire.
func (c *Character) Attaquer(target *Character) {
	damage := c.attack - target.defense
	if damage < 0 {
		damage = 0
	}
	target.hp -= damage
	fmt.Printf("%s attaque %s et lui inflige %d points de dégâts.\n", c.name, target.name, damage)
}

// UtiliserObjet permet au personnage d'utiliser un objet de son inventaire.
func (c *Character) UtiliserObjet(objet string) {
	if quantite, ok := c.inventory[objet]; ok && quantite > 0 {
		switch objet {
		case "Potion de soin":
			c.hp += 30
			fmt.Printf("%s utilise une potion de soin et récupère 30 points de vie.\n", c.name)
		case "Potion de poison":
			c.hp -= 20
			fmt.Printf("%s utilise une potion de poison et perd 20 points de vie.\n", c.name)
		}
		c.inventory[objet]--
	} else {
		fmt.Printf("%s n'a pas d'%s dans son inventaire.\n", c.name, objet)
	}
}

// EstEnVie vérifie si le personnage est encore en vie.
func (c *Character) EstEnVie() bool {
	return c.hp > 0
}

// Crée un personnage avec une classe donnée.
func CreerPersonnage(nom string, classe string) *Character {
	rand.Seed(time.Now().UnixNano())
	hp := rand.Intn(100) + 100
	attack := rand.Intn(20) + 10
	defense := rand.Intn(10) + 5
	inventory := make(map[string]int)
	inventory["Potion de soin"] = 3
	inventory["Potion de poison"] = 2
	return &Character{name: nom, class: classe, hp: hp, attack: attack, defense: defense, inventory: inventory}
}

func FightRound() {
    guerrier := CreerPersonnage("P1 Guerrier", "Guerrier")
    archer := CreerPersonnage("Archer", "Archer")

    tour := 1
    for guerrier.EstEnVie() && archer.EstEnVie() {
        fmt.Printf("Tour %d\n", tour)
        fmt.Printf("%s a %d points de vie restants, %s a %d points de vie restants.\n", guerrier.name, guerrier.hp, archer.name, archer.hp)

        // Demande au joueur de choisir entre attaquer ou utiliser un objet.
        var choix string
        fmt.Printf("%s, choisissez une action (attaquer/inventaire): ", guerrier.name)
        fmt.Scanln(&choix)

        switch choix {
        case "attaquer":
            guerrier.Attaquer(archer)
        case "inventaire":
            // Affiche l'inventaire du guerrier.
            fmt.Printf("%s, voici votre inventaire:\n", guerrier.name)
            for objet, quantite := range guerrier.inventory {
                fmt.Printf("%s: %d\n", objet, quantite)
            }
            // Demande au joueur de choisir un objet.
            var objetChoisi string
            fmt.Print("Choisissez un objet dans votre inventaire: ")
            fmt.Scanln(&objetChoisi)
            guerrier.UtiliserObjet(objetChoisi)
        default:
            fmt.Println("Action invalide.")
        }

        if !archer.EstEnVie() {
            fmt.Printf("%s a été vaincu ! %s gagne la bataille.\n", archer.name, guerrier.name)
            break
        }

        archer.Attaquer(guerrier)
        if !guerrier.EstEnVie() {
            fmt.Printf("%s a été vaincu ! %s gagne la bataille.\n", guerrier.name, archer.name)
            break
        }

        tour++
    }
}
