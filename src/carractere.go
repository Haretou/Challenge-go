package main

import (
    "fmt"
)

// Fonction pour ajouter un objet à l'inventaire
func (inv *Inventory) AddItem(item item) {
	inv.Items = append(inv.Items, item)
}

// Fonction pour afficher le contenu de l'inventaire
func (inv Inventory) DisplayInventory() {
	fmt.Println("Inventory:")
	for i, item := range inv.Items {
		fmt.Printf("%d. %s (Value: %d)\n", i+1, item.Name, item.Value)
	}
}

func AffInventory() {
	// Créer un inventaire vide
	myInventory := Inventory{}
	// Ajouter des objets à l'inventaire
	item1 := item{Name: "Épée", Value: 100}
	item2 := item{Name: "Armure", Value: 150}
	item3 := item{Name: "Potion de guérison", Value: 50}
    // On "place" les objets dans l'inventaire
	myInventory.AddItem(item1)
	myInventory.AddItem(item2)
	myInventory.AddItem(item3)
	// Afficher l'inventaire
	myInventory.DisplayInventory()
}
