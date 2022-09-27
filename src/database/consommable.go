package database

import (
	"strconv"
	"time"
)

type Consommable struct {
	//Nom
	Nom string
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
		c.InitIntern_Consommable(nom, 20, 3, 50, 0, 0, 0, 0)
	} else if nom == "Vitalité" { //Augmente la vitalité, se recupère aux Feux
		c.InitIntern_Consommable(nom, 300, 0, 1, 0, 0, 0, 0)
	} else if nom == "Force" { //Potion de force
		c.InitIntern_Consommable(nom, 300, 0, 0, 1, 0, 0, 0)
	} else if nom == "Dextérité" { //Potion de dextérité
		c.InitIntern_Consommable(nom, 300, 0, 0, 0, 1, 0, 0)
	} else if nom == "Intelligence" { //Potion d'intelligence
		c.InitIntern_Consommable(nom, 300, 0, 0, 0, 0, 1, 0)
	} else if nom == "Potion de poids max" {
		c.InitIntern_Consommable(nom, 300, 0, 0, 0, 0, 0, 1)
	} else if nom == "Fiole d'essence de pin pourri" { //Potion de poison
		c.InitIntern_Consommable(nom, 50, 0, 30, 0, 0, 0, 0)
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
			Affichage("Inventaire", []string{"Vous avez bu une potion de vie, vous avez maintenant " + strconv.Itoa(p.Pvact) + " pv"})
			Attendre()
		}
	} else if c.Nom == "Vitalité" {
		p.Vitalite++
		p.Niveau++
	} else if c.Nom == "Force" {
		p.Force++
		p.Niveau++
	} else if c.Nom == "Dextérité" {
		p.Dexterite++
		p.Niveau++
	} else if c.Nom == "Intelligence" {
		p.Intelligence++
		p.Niveau++
	} else if c.Nom == "Fiole d'essence de pin pourri" {
		if p.IsInInv(5) {
			p.Inv.Liste_consommables[5].Quantite -= 1
			for i := 0; i < 3; i++ {
				time.Sleep(1 * time.Second)
				p.Pvact -= c.PvBonus / 3
				Affichage("Poison", []string{"ouch, ça fait mal !", "Vous subissez " + strconv.Itoa(c.PvBonus/3) + " de dégats", "Il vous reste " + strconv.Itoa(p.Pvact) + " pv"})
				if p.Pvact <= 0 {
					p.Pvact = 0
				}
			}
			Attendre()
		}
	}
}

func (p *Personnage) IsInInv(n int) bool {
	if p.Inv.Liste_consommables[n].Quantite > 0 {
		return true
	} else {
		Affichage("Inventaire", []string{"Vous n'avez plus de " + p.Inv.Liste_consommables[n].Nom})
		return false
	}
}
