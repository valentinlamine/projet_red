package database

import "strconv"

type Armures struct {
	// Nom
	Nom   string
	Class string
	Prix  int
	// Stat Min
	NivMinFor int
	NivMinDex int
	NivMinInt int
	Niv       int
	// Stat VieBonus
	VieBonus    int
	Poids       int
	EstDebloque bool
	IsEquiped   bool
}

func (a *Armures) Initialisation_Interne_Armures(nom, class string, prix, NivMinFor, NivMinDex, NivMinInt, VieBonus, poids int) {
	a.Nom = nom
	a.Class = class
	a.Prix = prix
	a.NivMinFor = NivMinFor
	a.NivMinDex = NivMinDex
	a.NivMinInt = NivMinInt
	a.VieBonus = VieBonus
	a.Poids = poids
	a.EstDebloque = false
	a.EstDebloque = false
	a.Niv = 1
}

// Initialisation des casques
func (a *Armures) Initialisation_Armures(nombre int) {

	//Initialisation des casques
	switch nombre {
	case 1:
		a.Initialisation_Interne_Armures("Casque de Carcasse", "casque", 300, 8, 8, 6, 20, 3)
	case 2:
		a.Initialisation_Interne_Armures("Casque d'Havel", "casque", 5000, 16, 12, 8, 70, 13)
	case 3:
		a.Initialisation_Interne_Armures("Bandeau de cuir noir ", "casque", 500, 9, 12, 10, 35, 3)
	case 4:
		a.Initialisation_Interne_Armures("Heaume de Chevalier", "casque", 700, 13, 12, 10, 40, 6)
	case 5:
		a.Initialisation_Interne_Armures("Capuche de feu", "casque", 750, 7, 7, 15, 45, 3)
	case 6:
		a.Initialisation_Interne_Armures("Tête de Dragon", "casque", 10000, 20, 10, 15, 100, 10)
	//Initialisation des plastrons
	case 7:
		a.Initialisation_Interne_Armures("Plastron de Carcasse", "plastron", 300, 8, 8, 6, 30, 7)
	case 8:
		a.Initialisation_Interne_Armures("Plastron d'Havel", "plastron", 7000, 17, 13, 8, 120, 35)
	case 9:
		a.Initialisation_Interne_Armures("Veste de cuir noir", "plastron", 700, 9, 12, 10, 45, 5)
	case 10:
		a.Initialisation_Interne_Armures("Plastron de Chevalier", "plastron", 1000, 12, 12, 10, 40, 8)
	case 11:
		a.Initialisation_Interne_Armures("Manteau de feu", "plastron", 800, 7, 11, 15, 45, 4)
	case 12:
		a.Initialisation_Interne_Armures("Ecailles de Dragon", "plastron", 1000, 20, 11, 16, 220, 27)
	// Initialisation des gantelets
	case 13:
		a.Initialisation_Interne_Armures("Plastron de Carcasse", "gantelet", 200, 8, 8, 6, 20, 3)
	case 14:
		a.Initialisation_Interne_Armures("Gantelet d'Havel", "gantelet", 3000, 14, 11, 8, 50, 10)
	case 15:
		a.Initialisation_Interne_Armures("Gant de cuir noir", "gantelet", 300, 9, 12, 9, 25, 5)
	case 16:
		a.Initialisation_Interne_Armures("Gantelet de Chevalier", "gantelet", 500, 12, 12, 10, 35, 6)
	case 17:
		a.Initialisation_Interne_Armures("Manchette de feu", "gantelet", 400, 10, 12, 15, 30, 3)
	case 18:
		a.Initialisation_Interne_Armures("Griffes de Dragon", "gantelet", 10000, 20, 10, 15, 70, 8)
	// Initialisation des jambieres
	case 19:
		a.Initialisation_Interne_Armures("Jambières de Carcasse", "jambieres", 500, 8, 8, 6, 25, 8)
	case 20:
		a.Initialisation_Interne_Armures("Jambières d'Havel", "jambieres", 5000, 11, 9, 8, 80, 25)
	case 21:
		a.Initialisation_Interne_Armures("Bottes de cuir  noir", "jambieres", 500, 9, 12, 9, 35, 5)
	case 22:
		a.Initialisation_Interne_Armures("Jambières de Chevalier", "jambieres", 600, 12, 14, 10, 40, 10)
	case 23:
		a.Initialisation_Interne_Armures("Bottes de feu", "jambieres", 1000, 7, 11, 15, 45, 4)
	case 24:
		a.Initialisation_Interne_Armures("Pattes de Dragon", "jambieres", 1000, 20, 10, 16, 120, 15)
	}
}

func (a *Armures) Ameliorer_Armures(p *Personnage) {
	a.Niv++
	a.VieBonus = a.VieBonus + a.VieBonus/5
	switch a.Niv {
	case 2:
		p.Ames -= 100
		p.Inv.Liste_Items["éclat de titanite"] -= 6
	case 3:
		p.Ames -= 500
		p.Inv.Liste_Items["éclat de titanite"] -= 3
		p.Inv.Liste_Items["grand éclat de titanite"] -= 3
	case 4:
		p.Ames -= 2000
		p.Inv.Liste_Items["éclat de titanite"] -= 2
		p.Inv.Liste_Items["grand éclat de titanite"] -= 2
		p.Inv.Liste_Items["tablette de titanite"] -= 2
	}
	Affichage("Amélioration", []string{"Félécitation, vous venez d'améliorer votre armure !", "Désormais votre " + a.Nom + " est au niveau " + strconv.Itoa(a.Niv), "Désormais votre " + a.Nom + " donne " + strconv.Itoa(a.VieBonus) + " points de vie bonus !"}, true, true)
}
