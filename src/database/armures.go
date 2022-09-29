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
func (a *Armures) Init_Armures(number int) {

	//Init des casques
	switch number {
	case 1:
		a.InitIntern_Armures("Casque de Carcasse", "casque", 300, 8, 8, 6, 20, 3)
	case 2:
		a.InitIntern_Armures("Casque d'Havel", "casque", 5000, 16, 12, 8, 70, 13)
	case 3:
		a.InitIntern_Armures("Bandeau de cuir noir ", "casque", 500, 9, 12, 10, 35, 3)
	case 4:
		a.InitIntern_Armures("Heaume de Chevalier", "casque", 700, 13, 12, 10, 40, 6)
	case 5:
		a.InitIntern_Armures("Capuche de feu", "casque", 750, 7, 7, 15, 45, 3)
	case 6:
		a.InitIntern_Armures("Tête de Dragon", "casque", 10000, 20, 10, 15, 100, 10)
	//Init des plastrons
	case 7:
		a.InitIntern_Armures("Plastron de Carcasse", "plastron", 300, 8, 8, 6, 30, 7)
	case 8:
		a.InitIntern_Armures("Plastron d'Havel", "plastron", 7000, 17, 13, 8, 120, 35)
	case 9:
		a.InitIntern_Armures("Veste de cuir noir", "plastron", 700, 9, 12, 10, 45, 5)
	case 10:
		a.InitIntern_Armures("Plastron de Chevalier", "plastron", 1000, 12, 12, 10, 40, 8)
	case 11:
		a.InitIntern_Armures("Manteau de feu", "plastron", 800, 7, 11, 15, 45, 4)
	case 12:
		a.InitIntern_Armures("Ecailles de Dragon", "plastron", 1000, 20, 11, 16, 220, 27)
	// Init des gantelets
	case 13:
		a.InitIntern_Armures("Plastron de Carcasse", "gantelet", 200, 8, 8, 6, 20, 3)
	case 14:
		a.InitIntern_Armures("Gantelet d'Havel", "gantelet", 3000, 14, 11, 8, 50, 10)
	case 15:
		a.InitIntern_Armures("Gant de cuir noir", "gantelet", 300, 9, 12, 9, 25, 5)
	case 16:
		a.InitIntern_Armures("Gantelet de Chevalier", "gantelet", 500, 12, 12, 10, 35, 6)
	case 17:
		a.InitIntern_Armures("Manchette de feu", "gantelet", 400, 10, 12, 15, 30, 3)
	case 18:
		a.InitIntern_Armures("Griffes de Dragon", "gantelet", 10000, 20, 10, 15, 70, 8)
	// Init des jambieres
	case 19:
		a.InitIntern_Armures("Jambières de Carcasse", "jambieres", 500, 8, 8, 6, 25, 8)
	case 20:
		a.InitIntern_Armures("Jambières d'Havel", "jambieres", 5000, 11, 9, 8, 80, 25)
	case 21:
		a.InitIntern_Armures("Bottes de cuir  noir", "jambieres", 500, 9, 12, 9, 35, 5)
	case 22:
		a.InitIntern_Armures("Jambières de Chevalier", "jambieres", 600, 12, 14, 10, 40, 10)
	case 23:
		a.InitIntern_Armures("Bottes de feu", "jambieres", 1000, 7, 11, 15, 45, 4)
	case 24:
		a.InitIntern_Armures("Pattes de Dragon", "jambieres", 1000, 20, 10, 16, 120, 15)
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
		p.Inv.Liste_items["tablette de titanite"] -= 2
	}
	Affichage("Amélioration", []string{"Félécitation, vous venez d'améliorer votre armure !", "Désormais votre " + a.Nom + " est au niveau " + strconv.Itoa(a.Lvl), "Désormais votre " + a.Nom + " donne " + strconv.Itoa(a.Pvbonus) + " points de vie bonus !"}, true, true)
}
