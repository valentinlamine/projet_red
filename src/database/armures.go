package database

type Armures struct {
	// Nom
	nom   string
	class string
	// Stat Min
	lvlMinFor int
	lvlMinDex int
	lvlMinInt int
	// Stat PvBonus
	pvbonus    int
	poids      int
	isUnlocked bool
}

func (a *Armures) InitIntern_Armures(nom, class string, lvlMinFor, lvlMinDex, lvlMinInt, pvbonus, poids int) {
	a.nom = nom
	a.class = class
	a.lvlMinFor = lvlMinFor
	a.lvlMinDex = lvlMinDex
	a.lvlMinInt = lvlMinInt
	a.pvbonus = pvbonus
	a.poids = poids
}

//TODO Changer les stats des armures

// Init des casques
func (a *Armures) Init_Armures(nom string) {

	//Init des casques
	if nom == "Casque de Carcasse" {
		a.InitIntern_Armures("Casque de Carcasse", "casque", 9, 11, 8, 20, 3)
	} else if nom == "Casque d'Havel" {
		a.InitIntern_Armures("Casque d'Havel", "casque", 11, 9, 8, 50, 8)
	} else if nom == "Bandeau de cuir noir" {
		a.InitIntern_Armures("Bandeau de cuir noir ", "casque", 9, 12, 9, 35, 5)
	} else if nom == "Heaume de Chevalier" {
		a.InitIntern_Armures("Heaume de Chevalier", "casque", 10, 14, 10, 40, 6)
	} else if nom == "Capuche de feu" {
		a.InitIntern_Armures("Capuche de feu", "casque", 7, 7, 7, 15, 3)
	} else if nom == "Tête de Dragon" {
		a.InitIntern_Armures("Tête de Dragon", "casque", 14, 10, 8, 60, 12)

		// Init des armures
	} else if nom == "Plastron de Carcasse" {
		a.InitIntern_Armures("Plastron de Carcasse", "plastron", 9, 11, 8, 20, 3)
	} else if nom == "Plastron d'Havel" {
		a.InitIntern_Armures("Plastron d'Havel", "plastron", 11, 9, 8, 50, 8)
	} else if nom == "Veste de cuir  noir" {
		a.InitIntern_Armures("Veste de cuir  noir", "plastron", 9, 12, 9, 35, 5)
	} else if nom == "Plastron de Chevalier" {
		a.InitIntern_Armures("Plastron de Chevalier", "plastron", 10, 14, 10, 40, 6)
	} else if nom == "Manteau de feu" {
		a.InitIntern_Armures("Manteau de feu", "plastron", 7, 7, 7, 15, 3)
	} else if nom == "Ecailles de Dragon" {
		a.InitIntern_Armures("Ecailles de Dragon", "plastron", 14, 10, 8, 60, 12)

		// Init des gauntlets
	} else if nom == "Gantelet de Carcasse" {
		a.InitIntern_Armures("Plastron de Carcasse", "gantelet", 9, 11, 8, 20, 3)
	} else if nom == "Gantelet d'Havel" {
		a.InitIntern_Armures("Gantelet d'Havel", "gantelet", 11, 9, 8, 50, 8)
	} else if nom == "Gant de cuir  noir" {
		a.InitIntern_Armures("Bandeau de cuir  noir", "gantelet", 9, 12, 9, 35, 5)
	} else if nom == "Gantelet de Chevalier" {
		a.InitIntern_Armures("Plastron de Chevalier", "gantelet", 10, 14, 10, 40, 6)
	} else if nom == "Manchette de feu" {
		a.InitIntern_Armures("Manchette de feu", "gantelet", 7, 7, 7, 15, 3)
	} else if nom == "Griffes de Dragon" {
		a.InitIntern_Armures("Griffes de Dragon", "gantelet", 14, 10, 8, 60, 12)

		// Init des jambieres
	} else if nom == "Jambières de Carcasse" {
		a.InitIntern_Armures("Jambières de Carcasse", "jambieres", 9, 11, 8, 20, 3)
	} else if nom == "Jambières d'Havel" {
		a.InitIntern_Armures("Jambières d'Havel", "jambieres", 11, 9, 8, 50, 8)
	} else if nom == "Bottes de cuir  noir" {
		a.InitIntern_Armures("Bottes de cuir  noir", "jambieres", 9, 12, 9, 35, 5)
	} else if nom == "Jambières de Chevalier" {
		a.InitIntern_Armures("Jambières de Chevalier", "jambieres", 10, 14, 10, 40, 6)
	} else if nom == "Bottes de feu" {
		a.InitIntern_Armures("Bottes de feu", "jambieres", 7, 7, 7, 15, 3)
	} else if nom == "Pattes de Dragon" {
		a.InitIntern_Armures("Pattes de Dragon", "jambieres", 14, 10, 8, 60, 12)
	}
}
