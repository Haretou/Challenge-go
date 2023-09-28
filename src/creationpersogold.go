package main

import (
	"fmt"
)

// Définition des classes de personnages
type Personnage struct {
	Nom        string
	PointDeVie int
	Attaque    int
}

type Guerrier struct {
	Personnage
}

type Pretre struct {
	Personnage
}

type Chasseur struct {
	Personnage
}

// Fonction pour créer un nouveau personnage
func NouveauPersonnage(nom string, pointDeVie, attaque int) Personnage {
	return Personnage{
		Nom:        nom,
		PointDeVie: pointDeVie,
		Attaque:    attaque,
	}
}

// Fonction pour effectuer une attaque
func (p *Personnage) Attaquer(cible *Personnage) {
	fmt.Printf("%s attaque %s pour %d points de dégâts\n", p.Nom, cible.Nom, p.Attaque)
	cible.PointDeVie -= p.Attaque
	if cible.PointDeVie <= 0 {
		fmt.Printf("%s est vaincu!\n", cible.Nom)
	}
}

// Fonction pour afficher les informations du personnage
func (p *Personnage) AfficherInfos() {
	fmt.Printf("Nom: %s\nPoints de vie: %d\nAttaque: %d\n", p.Nom, p.PointDeVie, p.Attaque)
}

// Structure pour le système économique
type Economie struct {
	Or int
}

func new() {
	// Création de quelques personnages
	guerrier := Guerrier{NouveauPersonnage("Guerrier", 140, 20)}
	pretre := Pretre{NouveauPersonnage("Pretre", 1000, 15)}
	chasseur := Chasseur{NouveauPersonnage("Chasseur", 90, 35)}

	// Affichage des informations des personnages
	guerrier.AfficherInfos()
	pretre.AfficherInfos()
	chasseur.AfficherInfos()

	// Affichage des informations après l'attaque
	guerrier.AfficherInfos()
	pretre.AfficherInfos()

	// Système économique
	economie := Economie{Or: 100}
	fmt.Printf("Or disponible: %d\n", economie.Or)

	// // Simulation d'une transaction
	// montant := 30
	// if economie.Or >= montant {
	// 	economie.Or -= montant
	// 	fmt.Printf("Achat réussi. Or restant: %d\n", economie.Or)
	// } else {
	// 	fmt.Println("Transaction échouée. Pas assez d'or.")
	// }
}

