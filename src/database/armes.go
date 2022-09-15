package database

type Armes struct {
	// Nom
	nom string
	// Stat Min
	lvlMinFor int
	lvlMinDex int
	lvlMinInt int
	// Stat Damage
	deg   int
	poids int
}

func (a *Armes) InitIntern(nom string, lvlMinFor, lvlMinDex, lvlMinInt, deg, poids int) {
	a.nom = nom
	a.lvlMinFor = lvlMinFor
	a.lvlMinDex = lvlMinDex
	a.lvlMinInt = lvlMinInt
	a.deg = deg
	a.poids = poids

}

func (a *Armes) Init(nom string) {
	if nom == "Dague" {
		a.Init("Dague", 9, 11, 8, 20, 3)
	} else if nom == "Claymore" {
		a.Init("Claymore", 11, 9, 8, 50, 8)
	} else if nom == "Rapière" {
		a.Init("Rapière", 9, 12, 9, 35, 5)
	} else if nom == "Uchigatana" {
		a.Init("Uchigatana", 10, 14, 10, 40, 6)
	} else if nom == "Bâton" {
		a.Init("Bâton", 7, 7, 7, 15, 3)
	} else if nom == "Hache queue de gargouille" {
		a.Init("Hache queue de gargouille", 14, 10, 8, 60, 12)
	}
}
