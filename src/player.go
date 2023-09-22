package main
 
type Player struct {
	name string
	classe string
	Level int
	Health int
	Damage int
	def int
	Money int
}

func player() {
	var personne Player
		personne.name = "Nooby"
		personne.classe= "adventurer"
		personne.Level= 1
		personne.Health= 100
		personne.Damage= 10
		personne.def= 5
		personne.Money= 0
		
}
