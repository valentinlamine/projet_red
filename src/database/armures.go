package database

type Armures struct {
	// Nom
	nom string
	// Stat Min
	lvlMinFor int
	lvlMinDex int
	lvlMinInt int
	// Stat PvBonus
	pvbonus    int
	poids      int
	isUnlocked bool
	isEquiped  bool
}

func (a *Armures) InitIntern_Armures(nom string, lvlMinFor, lvlMinDex, lvlMinInt, pvbonus, poids int) {
	a.nom = nom
	a.lvlMinFor = lvlMinFor
	a.lvlMinDex = lvlMinDex
	a.lvlMinInt = lvlMinInt
	a.pvbonus = pvbonus
	a.poids = poids
	a.isUnlocked = false
	a.isEquiped = false

}

//TODO Changer les stats des armures

func (a *Armures) Init_Armures(nom string) {
	if nom == "Armure de Carcasse" {
		a.InitIntern_Armures("Armure de Carcasse", 9, 11, 8, 20, 3)
	} else if nom == "Armure de fer" {
		a.InitIntern_Armures("Armure de fer", 11, 9, 8, 50, 8)
	} else if nom == "Armure de cuir" {
		a.InitIntern_Armures("Armure de cuir", 9, 12, 9, 35, 5)
	} else if nom == "Armure de mithril" {
		a.InitIntern_Armures("Armure de mithril", 10, 14, 10, 40, 6)
	} else if nom == "Armure de bronze" {
		a.InitIntern_Armures("Armure de bronze", 7, 7, 7, 15, 3)
	} else if nom == "Armure de platine" {
		a.InitIntern_Armures("Armure de platine", 14, 10, 8, 60, 12)
	}
}
