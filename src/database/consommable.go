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
	if nom == "Fiole d'Estus" { //Potion de vie
		c.InitIntern_Consommable(nom, 0, 3, 50, 0, 0, 0, 0)
	} else if nom == "Résine de pin doré" { //Potion de force
		c.InitIntern_Consommable(nom, 0, 3, 0, 2, 0, 0, 0)
	} else if nom == "Résine de pin brulé" { //Potion de dextérité
		c.InitIntern_Consommable(nom, 0, 3, 0, 0, 2, 0, 0)
	} else if nom == "Résine de pin pourri" { //Potion d'intelligence
		c.InitIntern_Consommable(nom, 0, 3, 0, 0, 0, 1, 0)
	} else if nom == "Potion de poids max" {
		c.InitIntern_Consommable(nom, 0, 3, 0, 0, 0, 0, 1)
	}
}
