package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		// Affichage du menu principal
		fmt.Println("Menu principal")
		fmt.Println("1. Play")
		fmt.Println("2. Enemy")
		fmt.Println("3. Quitter")

		var choice int
		fmt.Print("Choisissez une option : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			
			Lors()
			play()
			Market()

		case 2:

			Enemy()
		case 3:
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
	fmt.Println("Le diable Abdela envoilla ses soldats à la conquête du pays.")
	time.Sleep(time.Second * 2)
	fmt.Println("Un jour un jeune homme aux nom Noobie entendit une voie qui lui dit")
	time.Sleep(time.Second * 2)
	fmt.Println("tu es l'héritier qui sauvera le pays et triomphera du diable. ")
	time.Sleep(time.Second * 2)
	fmt.Println("Il prend l'équipement et il va sauver son pays et tuer le diable.")
	fmt.Scanln()
}

func play() {

	fmt.Println("Option")
	fmt.Println("1. combatre")
	fmt.Println("2. Market")
	fmt.Println("3. Quitter")
	fmt.Scanln()
}
func Market(){
	fmt.Println("Bienvenue")
	fmt.Println("Que voulez vous acheter ?")
	fmt.Println("Potion de vie : Free")
	time.Sleep(time.Second * 2)
	fmt.Println("Potion de poison : 5 Po")
	time.Sleep(time.Second * 2)
	fmt.Println("épé en fer : 10 PO")
	fmt.Scanln()
}
func Enemy() {
	fmt.Println("History")
	time.Sleep(time.Second * 2)
	fmt.Println("Noms: Guerrier \n Point de vie: 140 \n damage: 20 \n défense: 10")
	time.Sleep(time.Second * 2)
	fmt.Println("Noms: Pretre \n Point de vie: 500 \n damage: 15 \n défense: 30")
	time.Sleep(time.Second * 2)
	fmt.Println("Noms: Chasseur \n Point de vie: 100 \n damage: 35 \n défense: 14")
	fmt.Scanln()
}