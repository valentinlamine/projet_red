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
	} else if nom == "Résine de pin doré" { //Potion de force
		c.InitIntern_Consommable(nom, 50, 0, 0, 2, 0, 0, 0)
	} else if nom == "Résine de pin brulé" { //Potion de dextérité
		c.InitIntern_Consommable(nom, 50, 0, 0, 0, 2, 0, 0)
	} else if nom == "Résine de pin pourri" { //Potion d'intelligence
		c.InitIntern_Consommable(nom, 50, 0, 0, 0, 0, 1, 0)
	} else if nom == "Potion de poids max" {
		c.InitIntern_Consommable(nom, 50, 0, 0, 0, 0, 0, 1)
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
	} else if c.Nom == "Résine de pin doré" {
		if p.IsInInv(1) {
			p.Inv.Liste_consommables[1].Quantite -= 1
			p.Force += c.MultiLvlFor
			Affichage("Inventaire", []string{"Vous avez bu une potion de force, vous avez maintenant " + strconv.Itoa(p.Force) + " de force"})
			Attendre()
		}
	} else if c.Nom == "Résine de pin brulé" {
		if p.IsInInv(2) {
			p.Inv.Liste_consommables[2].Quantite -= 1
			p.Dexterite += c.MultiLvlDex
			Affichage("Inventaire", []string{"Vous avez bu une potion de dextérité, vous avez maintenant " + strconv.Itoa(p.Dexterite) + " de dextérité"})
			Attendre()
		}
	} else if c.Nom == "Résine de pin pourri" {
		if p.IsInInv(3) {
			p.Inv.Liste_consommables[3].Quantite -= 1
			p.Intelligence += c.MultiLvlInt
			Affichage("Inventaire", []string{"Vous avez bu une potion d'intelligence, vous avez maintenant " + strconv.Itoa(p.Intelligence) + " d'intelligence"})
			Attendre()
		}
	} else if c.Nom == "Potion de poids max" {
		if p.IsInInv(4) {
			p.Inv.Liste_consommables[4].Quantite -= 1
			p.PoidsMax += c.MultiLvlPoidsMax
			Affichage("Inventaire", []string{"Vous avez bu une potion de poids max, vous avez maintenant " + strconv.Itoa(p.PoidsMax) + " de poids max"})
			Attendre()
		}

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
