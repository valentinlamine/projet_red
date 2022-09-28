package database

import "strconv"

type Armes struct {
	// Nom et Prix
	Nom  string
	Prix int
	// Stat Min
	LvlMinFor int
	LvlMinDex int
	LvlMinInt int
	Lvl       int
	// Stat Damage
	Deg        int
	Poids      int
	IsUnlocked bool
	IsEquiped  bool
}

type Boucliers struct {
	// Nom et Prix
	Nom  string
	Prix int
	// Stat Min
	LvlMinFor int
	LvlMinDex int
	LvlMinInt int
	Lvl       int
	// Stat Damage
	Pvbonus    int
	Poids      int
	IsUnlocked bool
	IsEquiped  bool
}

func (a *Armes) InitIntern_Armes(nom string, Prix, lvlMinFor, lvlMinDex, lvlMinInt, deg, poids int) {
	a.Nom = nom
	a.Prix = Prix
	a.LvlMinFor = lvlMinFor
	a.LvlMinDex = lvlMinDex
	a.LvlMinInt = lvlMinInt
	a.Deg = deg
	a.Poids = poids
	a.IsUnlocked = false
	a.IsEquiped = false
	a.Lvl = 1

}

func (a *Armes) Init_Armes(number int) {
	switch number {
	case 1:
		a.InitIntern_Armes("Dague", 100, 9, 11, 8, 20, 3)
	case 2:
		a.InitIntern_Armes("Claymore", 2000, 11, 9, 8, 50, 8)
	case 3:
		a.InitIntern_Armes("Rapière", 500, 9, 12, 9, 35, 5)
	case 4:
		a.InitIntern_Armes("Uchigatana", 1000, 10, 14, 10, 40, 6)
	case 5:
		a.InitIntern_Armes("Bâton", 100, 7, 7, 7, 15, 3)
	case 6:
		a.InitIntern_Armes("Hache queue de gargouille", 5000, 14, 10, 8, 60, 12)
	}
}

func (b *Boucliers) InitIntern_Bouclier(nom string, Prix, lvlMinFor, lvlMinDex, lvlMinInt, pvbonus, poids int) {
	b.Nom = nom
	b.Prix = Prix
	b.LvlMinFor = lvlMinFor
	b.LvlMinDex = lvlMinDex
	b.LvlMinInt = lvlMinInt
	b.Pvbonus = pvbonus
	b.Poids = poids
	b.IsUnlocked = false
	b.IsEquiped = false
	b.Lvl = 1

}

func (b *Boucliers) Init_Bouclier(number int) {
	switch number {
	case 1:
		b.InitIntern_Bouclier("Bouclier en bois", 50, 9, 8, 8, 20, 3)
	case 2:
		b.InitIntern_Bouclier("Bouclier en bois de cerf", 100, 10, 8, 8, 40, 5)
	case 3:
		b.InitIntern_Bouclier("Bouclier en fer", 700, 12, 8, 8, 75, 9)
	case 4:
		b.InitIntern_Bouclier("Bouclier en acier", 500, 10, 12, 9, 65, 5)
	case 5:
		b.InitIntern_Bouclier("Bouclier en mithril", 1200, 12, 16, 10, 120, 6)
	case 6:
		b.InitIntern_Bouclier("Bouclier d'Havel", 5000, 20, 10, 10, 175, 20)
	}
}

func (a *Armes) ReturnIsEquiped() bool {
	return a.IsEquiped
}

func (b *Boucliers) ReturnIsEquiped() bool {
	return b.IsEquiped
}

func (a *Armes) Ameliorer_arme(p *Personnage) {
	if a.Nom == "Bâton" {
		a.Lvl++
		a.Deg *= 3
	} else {
		a.Lvl++
		a.Deg = a.Deg + a.Deg/5
	}
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
	Affichage("Amélioration", []string{"Félécitation, vous venez d'améliorer votre arme !", "Désormais votre " + a.Nom + " est au niveau " + strconv.Itoa(a.Lvl), "Désormais votre " + a.Nom + " fait " + strconv.Itoa(a.Deg) + " dégats !"})
}

func (b *Boucliers) Ameliorer_bouclier(p *Personnage) {
	b.Lvl++
	b.Pvbonus = b.Pvbonus + b.Pvbonus/5
	switch b.Lvl {
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
	Affichage("Amélioration", []string{"Félécitation, vous venez d'améliorer votre bouclier !", "Désormais votre " + b.Nom + " est au niveau " + strconv.Itoa(b.Lvl), "Désormais votre " + b.Nom + " donne " + strconv.Itoa(b.Pvbonus) + " points de vie bonus !"})
}
