package database

type Consommable struct {
	//Nom
	Nom string
	//Classe
	Classe string
	//Prix
	Prix int
	//Quantité
	Quantite int
	//Bonus
	PvBonus          int
	MultiLvlFor      int
	MultiLvlDex      int
	MultiLvlInt      int
	MultiLvlPoidsMax int
}

func (c *Consommable) InitIntern_Consommable(nom string, prix, quantite, pvBonus, multiLvlFor, multiLvlDex, multiLvlInt, multiLvlPoidsMax int) {
	c.Nom = nom
	c.Prix = prix
	c.Quantite = quantite
	c.PvBonus = pvBonus
	c.MultiLvlFor = multiLvlFor
	c.MultiLvlDex = multiLvlDex
	c.MultiLvlInt = multiLvlInt
	c.MultiLvlPoidsMax = multiLvlPoidsMax
}

func (c *Consommable) Init_Consommable(nom string) {
	if nom == "Fiole d'Estus" { //Potion de vie, se recupère aux Feux
		c.InitIntern_Consommable(nom, 0, 3, 50, 0, 0, 0, 0)
	} else if nom == "Résine de pin doré" { //Potion de force
		c.InitIntern_Consommable(nom, 0, 3, 0, 2, 0, 0, 0)
	} else if nom == "Résine de pin brulé" { //Potion de dextérité
		c.InitIntern_Consommable(nom, 0, 3, 0, 0, 2, 0, 0)
	} else if nom == "Résine de pin pourri" { //Potion d'intelligence
		c.InitIntern_Consommable(nom, 0, 3, 0, 0, 0, 1, 0)
	} else if nom == "Potion de poids max" {
		c.InitIntern_Consommable(nom, 0, 3, 0, 0, 0, 0, 1)
	} else if nom == "Fiole d'essence de pin pourri" { //Potion de poison
		c.InitIntern_Consommable(nom, 0, 100, 50, 0, 0, 0, 0)
	}
}

func (p *Personnage) PrendrePot(c Consommable) {
	if c.Nom == "Fiole d'Estus" {
		if p.IsInInv(0) {
			p.Pvact += c.PvBonus
			p.Inv.Liste_consommables[0].Quantite -= 1
			if p.Pvact > p.Pvmax {
				p.Pvact = p.Pvmax
			}
		}
	} else if c.Nom == "Résine de pin doré" {
		if p.IsInInv(1) {
			p.Inv.Liste_consommables[1].Quantite -= 1
			p.Force += c.MultiLvlFor
		}
	} else if c.Nom == "Résine de pin brulé" {
		if p.IsInInv(2) {
			p.Inv.Liste_consommables[2].Quantite -= 1
			p.Dexterite += c.MultiLvlDex
		}
	} else if c.Nom == "Résine de pin pourri" {
		if p.IsInInv(3) {
			p.Inv.Liste_consommables[3].Quantite -= 1
			p.Intelligence += c.MultiLvlInt
		}
	} else if c.Nom == "Potion de poids max" {
		if p.IsInInv(4) {
			p.Inv.Liste_consommables[4].Quantite -= 1
			p.PoidsMax += c.MultiLvlPoidsMax
		}

	} else if c.Nom == "Fiole d'essence de pin pourri" {
		if p.IsInInv(5) {
			p.Inv.Liste_consommables[5].Quantite -= 1
			p.Pvact -= c.PvBonus
			if p.Pvact <= 0 {
				p.Pvact = 0
			}
		}
	}
}

func (p *Personnage) IsInInv(n int) bool {
	if p.Inv.Liste_consommables[n].Quantite > 0 {
		return true
	} else {
		return false
	}
}
