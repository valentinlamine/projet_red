package database

type Armes struct {
	// Nom
	Nom string
	// Stat Min
	LvlMinFor int
	LvlMinDex int
	LvlMinInt int
	// Stat Damage
	Deg   int
	Poids int
}

func (a *Armes) Init(nom string, lvlMinFor, lvlMinDex, lvlMinInt, deg, poids int) {
	a.Nom = nom
	a.LvlMinFor = lvlMinFor
	a.LvlMinDex = lvlMinDex
	a.LvlMinInt = lvlMinInt
	a.Deg = deg
	a.Poids = poids

}
