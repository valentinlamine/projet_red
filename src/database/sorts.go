package database

import "strconv"

type Sort struct {
	Nom        string
	Prix       int
	CoutMana   int
	Degats     int
	BoostPv    int
	Type       int
	IsUnlocked bool
}

const (
	DPS int = iota
	HEAL
)

func (s *Sort) Init(nom string, prix, coutMana, degats, boostPv, Type int) {
	s.Nom = nom
	s.Prix = prix
	s.CoutMana = coutMana
	s.Degats = degats
	s.BoostPv = boostPv
	s.Type = Type
	s.IsUnlocked = false
}

func (s *Sort) InitSort(number int) {
	if number == 1 {
		s.Init("Coup de poing", 500, 10, 15, 0, DPS)
	} else if number == 2 {
		s.Init("Boule de feu", 500, 15, 20, 0, DPS)
	} else if number == 3 {
		s.Init("Grande boule de feu", 500, 25, 30, 0, DPS)
	} else if number == 4 {
		s.Init("Soin", 500, 20, 0, 30, HEAL)
	} else if number == 5 {
		s.Init("Grand soin", 500, 40, 0, 70, HEAL)
	}
}

func (p *Personnage) SubirSort(s Sort, target *Personnage) {
	if p.Mana >= s.CoutMana {
		switch s.Type {
		case DPS:
			target.Pvact -= s.Degats
			Affichage("Combat", []string{"Vous avez infligé " + strconv.Itoa(s.Degats) + " points de dégats à " + target.Nom, "Il vous reste " + strconv.Itoa(p.Mana-s.CoutMana) + " points de mana"})
		case HEAL:
			p.Pvact += s.BoostPv
			Affichage("Combat", []string{"Vous avez gagné " + strconv.Itoa(s.BoostPv) + " points de vie"})
		}
	} else {
		Affichage("Combat", []string{"Vous n'avez pas assez de mana pour utiliser ce sort"})
	}
}

func Menu_sort(p *Personnage, target *Personnage) {
	var liste_affichage []string
	var liste_nom []string
	liste_affichage = append(liste_affichage, "Quel sort voulez-vous utiliser ?")
	for i := 0; i < len(p.Inv.Liste_sort); i++ {
		if p.Inv.Liste_sort[i].IsUnlocked {
			liste_nom = append(liste_nom, p.Inv.Liste_sort[i].Nom)
			liste_affichage = append(liste_affichage, strconv.Itoa(len(liste_nom))+". "+p.Inv.Liste_sort[i].Nom)
		}
	}
	liste_affichage = append(liste_affichage, "0. Retour")
	NouvelAffichage("Sort", liste_affichage)
	var choix = Choix(0, len(liste_nom))
	if choix != 0 {
		item := p.Inv.Get_Item(liste_nom[choix-1])
		switch item := item.(type) {
		case Sort:
			p.SubirSort(item, target)
		default:
			println("Erreur should never happend")
		}
	}
}
