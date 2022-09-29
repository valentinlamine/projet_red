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
	DegatsBonus int
	VieBonus    int
	PoidsBonus  int
	ManaBonus   int
}

func (c *Consommable) Initialisation_Interne_Consommable(nom string, prix, quantite, VieBonus, DegatsBonus, PoixBonus, Manabonus int) {
	c.Nom = nom
	c.Prix = prix
	c.Quantite = quantite
	c.VieBonus = VieBonus
	c.DegatsBonus = DegatsBonus
	c.PoidsBonus = PoixBonus
	c.ManaBonus = Manabonus
}

func (c *Consommable) Initialisation_Consommable(nombre int) {
	switch nombre {
	case 1:
		c.Initialisation_Interne_Consommable("Fiole d'Estus", 100, 0, 70, 0, 0, 0)
	case 2:
		c.Initialisation_Interne_Consommable("1 niveau de Vitalité", 300, 0, 20, 0, 0, 0)
	case 3:
		c.Initialisation_Interne_Consommable("1 niveau de Force", 300, 0, 0, 5, 5, 0)
	case 4:
		c.Initialisation_Interne_Consommable("1 niveau de Dextérité", 300, 0, 0, 0, 0, 0)
	case 5:
		c.Initialisation_Interne_Consommable("1 niveau de Intelligence", 300, 0, 0, 0, 0, 20)
	case 6:
		c.Initialisation_Interne_Consommable("Fiole d'essence de pin pourri", 100, 0, 30, 0, 0, 0)
	}
}

func (p *Personnage) Prendre_Potion(c Consommable) {
	if c.Nom == "Fiole d'Estus" {
		if p.Est_Dans_Inventaire(0) {
			p.VieAct += c.VieBonus
			p.Inv.Liste_Consommables[0].Quantite -= 1
			if p.VieAct > p.VieMax {
				p.VieAct = p.VieMax
			}
			Affichage("Information", []string{"Vous avez bu une potion de vie, vous avez maintenant " + strconv.Itoa(p.VieAct) + " de vie"}, true, true)
		}
	} else if c.Nom == "1 niveau de Vitalité" {
		p.Vitalite++
		p.Niveau++
		p.VieMax += c.VieBonus
		Affichage("Information", []string{"Vous avez pris un niveau de vitalité, vous avez maintenant " + strconv.Itoa(p.Vitalite) + " de vitalité"}, true, true)
	} else if c.Nom == "1 niveau de Force" {
		p.Force++
		p.Niveau++
		p.Degats += c.DegatsBonus
		p.PoidsMax += c.PoidsBonus
		Affichage("Information", []string{"Vous avez pris un niveau de force, vous avez maintenant " + strconv.Itoa(p.Force) + " de force"}, true, true)
	} else if c.Nom == "1 niveau de Dextérité" {
		p.Dexterite++
		p.Niveau++
		p.Initiative++
		Affichage("Information", []string{"Vous avez pris un niveau de dextérité, vous avez maintenant " + strconv.Itoa(p.Dexterite) + " de dextérité"}, true, true)
	} else if c.Nom == "1 niveau de Intelligence" {
		p.Intelligence++
		p.Niveau++
		p.ManaMax += c.ManaBonus
		for _, v := range p.Inv.Liste_Sorts {
			v.Degats = int(float64(v.Degats) * 1.3)
			v.BoostVie = int(float64(v.BoostVie) * 1.3)
		}
		Affichage("Information", []string{"Vous avez pris un niveau d'intelligence, vous avez maintenant " + strconv.Itoa(p.Intelligence) + " d'intelligence"}, true, true)
	} else if c.Nom == "Fiole d'essence de pin pourri" {
		if p.Est_Dans_Inventaire(5) {
			p.Inv.Liste_Consommables[5].Quantite -= 1
			for i := 0; i < 3; i++ {
				time.Sleep(1 * time.Second)
				p.VieAct -= c.VieBonus / 3
				Affichage("Poison", []string{"ouch, ça fait mal !", "Vous subissez " + strconv.Itoa(c.VieBonus/3) + " de dégats", "Il vous reste " + strconv.Itoa(p.VieAct) + " de vie"}, false, false)
				if p.VieAct <= 0 {
					p.VieAct = 0
				}
			}
			Attendre()
		}
	}
}

func (p *Personnage) Est_Dans_Inventaire(n int) bool {
	if p.Inv.Liste_Consommables[n].Quantite > 0 {
		return true
	} else {
		Affichage("Inventaire", []string{"Vous n'avez plus de " + p.Inv.Liste_Consommables[n].Nom}, true, true)
		return false
	}
}
