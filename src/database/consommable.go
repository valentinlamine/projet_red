package database

type Consomable struct {
	//Nom
	nom string
	//Prix
	prix int
	//Quantité
	quantite int
	//Bonus
	pvBonus          int
	multiLvlFor      int
	multiLvlDex      int
	multiLvlInt      int
	multiLvlPoidsMax int
}

func (c *Consomable) InitIntern(nom, classe string, prix, quantite, pvBonus, multiLvlFor, multiLvlDex, multiLvlInt, multiLvlPoidsMax int) {
	c.Nom = nom
	c.Classe = classe
	c.Prix = prix
	c.Quantite = quantite
	c.PvBonus = pvBonus
	c.MultiLvlFor = multiLvlFor
	c.MultiLvlDex = multiLvlDex
	c.MultiLvlInt = multiLvlInt
	c.MultiLvlPoidsMax = multiLvlPoidsMax
}

func (c *Consomable) Init(nom string) {
	if nom == "Fiole d'Estus" {
		
}
