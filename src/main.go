package main

import "fmt"

type Personnage struct {
	nom    string // Nom du personnage
	classe string // Classe du personnage (Guerrier, Chevalier, pyromancien, mendiant)
	niveau int    // Niveau du personnage
	//stats améliorables
	pvmax        int // Vitalité du personnage (points de vie)
	force        int // Force du personnage (dégâts infligés, ports d'armes, etc.)
	dexterite    int // Dextérité du personnage (dégâts infligés, esquive, etc.)
	intelligence int // Intelligence du personnage (dégâts magiques, etc.)
	//état du personnage
	pvact int // Points de vie actuels

	inventaire []string
}

func main() {
	var p1 Personnage
	p1.Init("Théo", "Elfe")
	p2.Init("Valentin", "Humain")
	p3.Init("Alexandre", "Nain")
	p4.Init("Romain", "Gobelin")
}

func (p *Personnage) Init(nom, classe string) {
	if classe == "Guerrrier" {
		continue
	} else if classe == "Chevalier" {
		continue
	} else if classe == "Pyromancien" {
		continue
	} else if classe == "Mendiant" {
		continue
	} else {
		fmt.Println("Classe inconnue")
	}

}
