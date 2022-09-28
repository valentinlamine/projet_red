package database

import "strconv"

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
		a.InitIntern_Armures("Casque de Carcasse", "casque", 500, 9, 11, 8, 20, 3)
	} else if nom == "Casque d'Havel" {
		a.InitIntern_Armures("Casque d'Havel", "casque", 500, 11, 9, 8, 50, 8)
	} else if nom == "Bandeau de cuir noir" {
		a.InitIntern_Armures("Bandeau de cuir noir ", "casque", 500, 9, 12, 9, 35, 5)
	} else if nom == "Heaume de Chevalier" {
		a.InitIntern_Armures("Heaume de Chevalier", "casque", 500, 10, 14, 10, 40, 6)
	} else if nom == "Capuche de feu" {
		a.InitIntern_Armures("Capuche de feu", "casque", 500, 7, 7, 7, 15, 3)
	} else if nom == "Tête de Dragon" {
		a.InitIntern_Armures("Tête de Dragon", "casque", 500, 14, 10, 8, 60, 12)

		// Init des armures
	} else if nom == "Plastron de Carcasse" {
		a.InitIntern_Armures("Plastron de Carcasse", "plastron", 500, 9, 11, 8, 20, 3)
	} else if nom == "Plastron d'Havel" {
		a.InitIntern_Armures("Plastron d'Havel", "plastron", 500, 11, 9, 8, 50, 8)
	} else if nom == "Veste de cuir  noir" {
		a.InitIntern_Armures("Veste de cuir  noir", "plastron", 500, 9, 12, 9, 35, 5)
	} else if nom == "Plastron de Chevalier" {
		a.InitIntern_Armures("Plastron de Chevalier", "plastron", 500, 10, 14, 10, 40, 6)
	} else if nom == "Manteau de feu" {
		a.InitIntern_Armures("Manteau de feu", "plastron", 500, 7, 7, 7, 15, 3)
	} else if nom == "Ecailles de Dragon" {
		a.InitIntern_Armures("Ecailles de Dragon", "plastron", 500, 14, 10, 8, 60, 12)

		// Init des gauntlets
	} else if nom == "Gantelet de Carcasse" {
		a.InitIntern_Armures("Plastron de Carcasse", "gantelet", 500, 9, 11, 8, 20, 3)
	} else if nom == "Gantelet d'Havel" {
		a.InitIntern_Armures("Gantelet d'Havel", "gantelet", 500, 11, 9, 8, 50, 8)
	} else if nom == "Gant de cuir  noir" {
		a.InitIntern_Armures("Bandeau de cuir  noir", "gantelet", 500, 9, 12, 9, 35, 5)
	} else if nom == "Gantelet de Chevalier" {
		a.InitIntern_Armures("Plastron de Chevalier", "gantelet", 500, 10, 14, 10, 40, 6)
	} else if nom == "Manchette de feu" {
		a.InitIntern_Armures("Manchette de feu", "gantelet", 500, 7, 7, 7, 15, 3)
	} else if nom == "Griffes de Dragon" {
		a.InitIntern_Armures("Griffes de Dragon", "gantelet", 500, 14, 10, 8, 60, 12)

		// Init des jambieres
	} else if nom == "Jambières de Carcasse" {
		a.InitIntern_Armures("Jambières de Carcasse", "jambieres", 500, 9, 11, 8, 20, 3)
	} else if nom == "Jambières d'Havel" {
		a.InitIntern_Armures("Jambières d'Havel", "jambieres", 500, 11, 9, 8, 50, 8)
	} else if nom == "Bottes de cuir  noir" {
		a.InitIntern_Armures("Bottes de cuir  noir", "jambieres", 500, 9, 12, 9, 35, 5)
	} else if nom == "Jambières de Chevalier" {
		a.InitIntern_Armures("Jambières de Chevalier", "jambieres", 500, 10, 14, 10, 40, 6)
	} else if nom == "Bottes de feu" {
		a.InitIntern_Armures("Bottes de feu", "jambieres", 500, 7, 7, 7, 15, 3)
	} else if nom == "Pattes de Dragon" {
		a.InitIntern_Armures("Pattes de Dragon", "jambieres", 500, 14, 10, 8, 60, 12)
	}
}

func (a *Armures) ReturnIsEquiped() bool {
	return a.IsEquiped
}

func (a *Armures) Ameliorer_Armures(p *Personnage) {
	a.Lvl++
	a.Pvbonus = a.Pvbonus + a.Pvbonus/5
	switch a.Lvl {
	case 2:
		p.Ames -= 100
		p.Inv.Liste_items["éclat de titanite"] -= 6
	case 3:
		p.Ames -= 500
		p.Inv.Liste_items["éclat de titanite"] -= 3
		p.Inv.Liste_items["grand éclat de titanite"] -= 3
	case 4:
		p.Ames -= 2000
		p.Inv.Liste_items["éclat de titanite"] -= 2
		p.Inv.Liste_items["grand éclat de titanite"] -= 2
		p.Inv.Liste_items["tablette éclat de titanite"] -= 2
	}
	Affichage("Amélioration", []string{"Félécitation, vous venez d'améliorer votre armure !", "Désormais votre " + a.Nom + " est au niveau " + strconv.Itoa(a.Lvl), "Désormais votre " + a.Nom + " donne " + strconv.Itoa(a.Pvbonus) + " points de vie bonus !"})
}
