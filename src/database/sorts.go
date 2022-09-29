package database

import (
	"fmt"
	"strconv"
)

type Sort struct {
	Nom         string
	Prix        int
	ManaCout    int
	Degats      int
	BoostVie    int
	Type        int
	EstDebloque bool
}

const (
	DPS int = iota
	VIE
)

func (s *Sort) Initialisation(nom string, prix, ManaCout, degats, BoostVie, Type int) {
	s.Nom = nom
	s.Prix = prix
	s.ManaCout = ManaCout
	s.Degats = degats
	s.BoostVie = BoostVie
	s.Type = Type
	s.EstDebloque = false
}

func (s *Sort) Initialisation_Sort(nombre int) {
	if nombre == 1 {
		s.Initialisation("Coup de poing", 100, 10, 15, 0, DPS)
	} else if nombre == 2 {
		s.Initialisation("Boule de feu", 500, 25, 25, 0, DPS)
	} else if nombre == 3 {
		s.Initialisation("Grande boule de feu", 2000, 60, 50, 0, DPS)
	} else if nombre == 4 {
		s.Initialisation("Boule du Chaos", 5000, 120, 110, 0, DPS)
	} else if nombre == 5 {
		s.Initialisation("Soin", 500, 30, 0, 50, VIE)
	} else if nombre == 6 {
		s.Initialisation("Grand soin", 2000, 60, 0, 120, VIE)
	}
}

func (p *Personnage) Subir_Sort(s Sort, cible *Personnage) {
	fmt.Println(p.ManaAct, s.ManaCout)
	if p.ManaAct >= s.ManaCout {
		switch s.Type {
		case DPS:
			cible.VieAct -= s.Degats
			Affichage("Combat", []string{"Vous avez infligé " + strconv.Itoa(s.Degats) + " points de dégats à " + cible.Nom, "Il vous reste " + strconv.Itoa(p.ManaAct-s.ManaCout) + " points de mana"}, false, false)
		case VIE:
			p.VieAct += s.BoostVie
			Affichage("Combat", []string{"Vous avez gagné " + strconv.Itoa(s.BoostVie) + " points de vie"}, false, false)
		}
	} else {
		Affichage("Combat", []string{"Vous n'avez pas assez de mana pour utiliser ce sort"}, false, false)
	}
}

func Menu_Sort(p *Personnage, cible *Personnage) int {
	var degats int
	var liste_affichage []string
	var liste_nom []string
	liste_affichage = append(liste_affichage, "Quel sort voulez-vous utiliser ?")
	for i := 0; i < len(p.Inv.Liste_Sorts); i++ {
		if p.Inv.Liste_Sorts[i].EstDebloque {
			liste_nom = append(liste_nom, p.Inv.Liste_Sorts[i].Nom)
			liste_affichage = append(liste_affichage, strconv.Itoa(len(liste_nom))+". "+p.Inv.Liste_Sorts[i].Nom)
		}
	}
	liste_affichage = append(liste_affichage, "0. Retour")
	Affichage("Sort", liste_affichage, true, false)
	var choix = Choix(0, len(liste_nom))
	if choix != 0 {
		item := p.Inv.Recuperer_Item(liste_nom[choix-1])
		switch item := item.(type) {
		case Sort:
			p.Subir_Sort(item, cible)
			degats = item.Degats
		default:
			println("Erreur should never happend")
		}
	}
	return degats
}
