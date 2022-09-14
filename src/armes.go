package main

import "fmt"

type Armes struct {
	// Nom
	nom string
	// Stat Min
	lvlMinFor int
	lvlMinDex int
	lvlMinInt int

	// Stat Damage
	deg int
	poids int
}

func (a *Armes) Init(nom string, lvlMinFor, lvlMinDex, lvlMinInt, deg, poids int) {
	a.nom = nom

	a.lvlMinFor = lvlMinFor
	a.lvlMinDex = lvlMinDex
	a.lvlMinInt = lvlMinInt

	a.deg = deg
	a.poids = poids
	
}

//Test
var Dague Armes
Dague.Init(Dague, 9, 11, 8, 20, 3)
