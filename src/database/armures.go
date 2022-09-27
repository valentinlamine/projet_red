package database

type Armures struct {
	// Nom
	Nom   string
	Class string
	Prix  int
	// Stat Min
	LvlMinFor int
	LvlMinDex int
	LvlMinInt int
	Lvl       int
	// Stat PvBonus
	Pvbonus    int
	Poids      int
	IsUnlocked bool
	IsEquiped  bool
}

func (a *Armures) InitIntern_Armures(nom, class string, prix, lvlMinFor, lvlMinDex, lvlMinInt, pvbonus, poids int) {
	a.Nom = nom
	a.Class = class
	a.Prix = prix
	a.LvlMinFor = lvlMinFor
	a.LvlMinDex = lvlMinDex
	a.LvlMinInt = lvlMinInt
	a.Pvbonus = pvbonus
	a.Poids = poids
	a.IsUnlocked = false
	a.IsUnlocked = false
	a.Lvl = 1
}

//TODO Changer les stats des armures

// Init des casques
func (a *Armures) Init_Armures(nom string) {

	//Init des casques
	if nom == "Casque de Carcasse" {
		a.InitIntern_Armures("Casque de Carcasse", "casque", 300, 8, 8, 6, 20, 3)
	} else if nom == "Casque d'Havel" {
		a.InitIntern_Armures("Casque d'Havel", "casque", 5000, 16, 12, 8, 70, 13)
	} else if nom == "Bandeau de cuir noir" {
		a.InitIntern_Armures("Bandeau de cuir noir ", "casque", 500, 9, 12, 10, 35, 3)
	} else if nom == "Heaume de Chevalier" {
		a.InitIntern_Armures("Heaume de Chevalier", "casque", 700, 13, 12, 10, 40, 6)
	} else if nom == "Capuche de feu" {
		a.InitIntern_Armures("Capuche de feu", "casque", 750, 7, 7, 15, 45, 3)
	} else if nom == "Tête de Dragon" {
		a.InitIntern_Armures("Tête de Dragon", "casque", 10000, 20, 10, 15, 100, 10)

		// Init des armures
	} else if nom == "Plastron de Carcasse" {
		a.InitIntern_Armures("Plastron de Carcasse", "plastron", 300, 8, 8, 6, 30, 7)
	} else if nom == "Plastron d'Havel" {
		a.InitIntern_Armures("Plastron d'Havel", "plastron", 7000, 17, 13, 8, 120, 35)
	} else if nom == "Veste de cuir  noir" {
		a.InitIntern_Armures("Veste de cuir  noir", "plastron", 700, 9, 12, 10, 45, 5)
	} else if nom == "Plastron de Chevalier" {
		a.InitIntern_Armures("Plastron de Chevalier", "plastron", 1000, 12, 12, 10, 40, 8)
	} else if nom == "Manteau de feu" {
		a.InitIntern_Armures("Manteau de feu", "plastron", 800, 7, 11, 15, 45, 4)
	} else if nom == "Ecailles de Dragon" {
		a.InitIntern_Armures("Ecailles de Dragon", "plastron", 1000, 20, 11, 16, 220, 27)

		// Init des gauntlets
	} else if nom == "Gantelet de Carcasse" {
		a.InitIntern_Armures("Plastron de Carcasse", "gantelet", 200, 8, 8, 6, 20, 3)
	} else if nom == "Gantelet d'Havel" {
		a.InitIntern_Armures("Gantelet d'Havel", "gantelet", 3000, 14, 11, 8, 50, 10)
	} else if nom == "Gant de cuir  noir" {
		a.InitIntern_Armures("Gant de cuir noir", "gantelet", 300, 9, 12, 9, 25, 5)
	} else if nom == "Gantelet de Chevalier" {
		a.InitIntern_Armures("Gantelet de Chevalier", "gantelet", 500, 12, 12, 10, 35, 6)
	} else if nom == "Manchette de feu" {
		a.InitIntern_Armures("Manchette de feu", "gantelet", 400, 10, 12, 15, 30, 3)
	} else if nom == "Griffes de Dragon" {
		a.InitIntern_Armures("Griffes de Dragon", "gantelet", 10000, 20, 10, 15, 70, 8)

		// Init des jambieres
	} else if nom == "Jambières de Carcasse" {
		a.InitIntern_Armures("Jambières de Carcasse", "jambieres", 500, 8, 8, 6, 25, 8)
	} else if nom == "Jambières d'Havel" {
		a.InitIntern_Armures("Jambières d'Havel", "jambieres", 500, 11, 9, 8, 80, 25)
	} else if nom == "Bottes de cuir noir" {
		a.InitIntern_Armures("Bottes de cuir  noir", "jambieres", 500, 9, 12, 9, 35, 5)
	} else if nom == "Jambières de Chevalier" {
		a.InitIntern_Armures("Jambières de Chevalier", "jambieres", 600, 12, 14, 10, 40, 10)
	} else if nom == "Bottes de feu" {
		a.InitIntern_Armures("Bottes de feu", "jambieres", 1000, 7, 11, 15, 45, 4)
	} else if nom == "Pattes de Dragon" {
		a.InitIntern_Armures("Pattes de Dragon", "jambieres", 1000, 20, 10, 16, 120, 15)
	}
}

func (a Armures) ReturnIsEquiped() bool {
	return a.IsEquiped
}
