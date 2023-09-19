package main

// Structure pour représenter un personnage de MMORPG
type Player struct {
	// attributs
    Name string
    Gender string
    Classe string
    Level int // permet d'augmenter toutes les autre stats ( sauf Money ) en montant 
    Health int
    Damage int
    Def int // baisse les dégats subit en fonction de valeur de cette variable
    Money  int
}

type Enemy struct {
    Name string
    Level int
    Health int
    Damage int
    Def int
    Money int // Pour qu'il puisse "drop" de l'argent au joueurs en mourrant
}

// Structure représentant un objet dans l'inventaire
type item struct {
	Name  string
	Value int // Prix / valeur de l'objet
}

// Structure représentant l'inventaire
type Inventory struct {
	Items []item
}
