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
	DegBonus   int
	PvBonus    int
	PoidsBonus int
	ManaBonus  int
}

func (c *Consommable) InitIntern_Consommable(nom string, prix, quantite, pvBonus, degBonus, PoixBonus, Manabonus int) {
	c.Nom = nom
	c.Prix = prix
	c.Quantite = quantite
	c.PvBonus = pvBonus
	c.DegBonus = degBonus
	c.PoidsBonus = PoixBonus
	c.ManaBonus = Manabonus
}

func (c *Consommable) Init_Consommable(number int) {
	switch number {
	case 1:
		c.InitIntern_Consommable("Fiole d'Estus", 100, 0, 0, 70, 0, 0)
	case 2:
		c.InitIntern_Consommable("1 niveau de Vitalité", 300, 0, 20, 0, 0, 0)
	case 3:
		c.InitIntern_Consommable("1 niveau de Force", 300, 0, 0, 5, 5, 0)
	case 4:
		c.InitIntern_Consommable("1 niveau de Dextérité", 300, 0, 0, 0, 0, 0)
	case 5:
		c.InitIntern_Consommable("1 niveau de Intelligence", 300, 0, 0, 0, 0, 20)
	case 6:
		c.InitIntern_Consommable("Fiole d'essence de pin pourri", 100, 0, 30, 0, 0, 0)
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
			Affichage("Information", []string{"Vous avez bu une potion de vie, vous avez maintenant " + strconv.Itoa(p.Pvact) + " pv"})
			Attendre()
		}
	} else if c.Nom == "1 niveau de Vitalité" {
		p.Vitalite++
		p.Niveau++
		p.Pvmax += c.PvBonus
		Affichage("Information", []string{"Vous avez pris un niveau de vitalité, vous avez maintenant " + strconv.Itoa(p.Vitalite) + " de vitalité"})
		Attendre()
	} else if c.Nom == "1 niveau de Force" {
		p.Force++
		p.Niveau++
		p.Degats += c.DegBonus
		p.PoidsMax += c.PoidsBonus
		Affichage("Information", []string{"Vous avez pris un niveau de force, vous avez maintenant " + strconv.Itoa(p.Force) + " de force"})
		Attendre()
	} else if c.Nom == "1 niveau de Dextérité" {
		p.Dexterite++
		p.Niveau++
		p.Initiative++
		Affichage("Information", []string{"Vous avez pris un niveau de dextérité, vous avez maintenant " + strconv.Itoa(p.Dexterite) + " de dextérité"})
		Attendre()
	} else if c.Nom == "1 niveau de Intelligence" {
		p.Intelligence++
		p.Niveau++
		p.ManaMax += c.ManaBonus
		for _, v := range p.Inv.Liste_sort {
			v.Degats = int(float64(v.Degats) * 1.3)
			v.BoostPv = int(float64(v.BoostPv) * 1.3)
		}
		Affichage("Information", []string{"Vous avez pris un niveau d'intelligence, vous avez maintenant " + strconv.Itoa(p.Intelligence) + " d'intelligence"})
		Attendre()
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
