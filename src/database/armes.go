package database

import "strconv"

type Armes struct {
	// Nom et Prix
	Nom  string
	Prix int
	// Stat Min
	NivMinFor int
	NivMinDex int
	NivMinInt int
	Niv       int
	// Stat Damage
	Degats      int
	Poids       int
	EstDebloque bool
	IsEquiped   bool
}

type Boucliers struct {
	// Nom et Prix
	Nom  string
	Prix int
	// Stat Min
	NivMinFor int
	NivMinDex int
	NivMinInt int
	Niv       int
	// Stat Damage
	VieBonus    int
	Poids       int
	EstDebloque bool
	IsEquiped   bool
}

func (a *Armes) Initialisation_Interne_Armes(nom string, Prix, NivMinFor, NivMinDex, NivMinInt, Degats, poids int) {
	a.Nom = nom
	a.Prix = Prix
	a.NivMinFor = NivMinFor
	a.NivMinDex = NivMinDex
	a.NivMinInt = NivMinInt
	a.Degats = Degats
	a.Poids = poids
	a.EstDebloque = false
	a.IsEquiped = false
	a.Niv = 1

}

func (a *Armes) Initialisation_Armes(nombre int) {
	switch nombre {
	case 1:
		a.Initialisation_Interne_Armes("Dague", 100, 9, 11, 8, 20, 3)
	case 2:
		a.Initialisation_Interne_Armes("Claymore", 2000, 11, 9, 8, 50, 8)
	case 3:
		a.Initialisation_Interne_Armes("Rapière", 500, 9, 12, 9, 35, 5)
	case 4:
		a.Initialisation_Interne_Armes("Uchigatana", 1000, 10, 14, 10, 40, 6)
	case 5:
		a.Initialisation_Interne_Armes("Bâton", 100, 7, 7, 7, 15, 3)
	case 6:
		a.Initialisation_Interne_Armes("Hache queue de gargouille", 5000, 14, 10, 8, 60, 12)
	}
}

func (b *Boucliers) Initialisation_Interne_Bouclier(nom string, Prix, NivMinFor, NivMinDex, NivMinInt, VieBonus, poids int) {
	b.Nom = nom
	b.Prix = Prix
	b.NivMinFor = NivMinFor
	b.NivMinDex = NivMinDex
	b.NivMinInt = NivMinInt
	b.VieBonus = VieBonus
	b.Poids = poids
	b.EstDebloque = false
	b.IsEquiped = false
	b.Niv = 1

}

func (b *Boucliers) Initialisation_Bouclier(nombre int) {
	switch nombre {
	case 1:
		b.Initialisation_Interne_Bouclier("Bouclier en bois", 50, 9, 8, 8, 20, 3)
	case 2:
		b.Initialisation_Interne_Bouclier("Bouclier en bois de cerf", 100, 10, 8, 8, 40, 5)
	case 3:
		b.Initialisation_Interne_Bouclier("Bouclier en fer", 700, 12, 8, 8, 75, 9)
	case 4:
		b.Initialisation_Interne_Bouclier("Bouclier en acier", 500, 10, 12, 9, 65, 5)
	case 5:
		b.Initialisation_Interne_Bouclier("Bouclier en mithril", 1200, 12, 16, 10, 120, 6)
	case 6:
		b.Initialisation_Interne_Bouclier("Bouclier d'Havel", 5000, 20, 10, 10, 175, 20)
	}
}

func (a *Armes) Ameliorer_Arme(p *Personnage) {
	if a.Nom == "Bâton" {
		a.Niv++
		a.Degats *= 3
	} else {
		a.Niv++
		a.Degats = a.Degats + a.Degats/5
	}
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
	Affichage("Amélioration", []string{"Félécitation, vous venez d'améliorer votre arme !", "Désormais votre " + a.Nom + " est au niveau " + strconv.Itoa(a.Niv), "Désormais votre " + a.Nom + " fait " + strconv.Itoa(a.Degats) + " dégats !"}, true, true)
}

func (b *Boucliers) Ameliorer_Bouclier(p *Personnage) {
	b.Niv++
	b.VieBonus = b.VieBonus + b.VieBonus/5
	switch b.Niv {
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
	Affichage("Amélioration", []string{"Félécitation, vous venez d'améliorer votre bouclier !", "Désormais votre " + b.Nom + " est au niveau " + strconv.Itoa(b.Niv), "Désormais votre " + b.Nom + " donne " + strconv.Itoa(b.VieBonus) + " points de vie bonus !"}, true, true)
}
