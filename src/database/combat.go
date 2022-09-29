package database

import (
	"math/rand"
	"strconv"
	"time"
)

func Combat(joueur, monstre *Personnage) bool {
	Affichage("Combat", []string{"Vous affrontez un " + monstre.Nom}, true, false)
	var CompteurTour int = 1
	if joueur.Initiative > monstre.Initiative {
		for joueur.VieAct > 0 && monstre.VieAct > 0 {
			time.Sleep(1 * time.Second)
			if monstre.Tour_Joueur(joueur) {
				return true
			}
			time.Sleep(1 * time.Second)
			joueur.Tour_Ennemi(monstre, CompteurTour)
			joueur.ManaAct += 15
			CompteurTour++
		}
	} else {
		for joueur.VieAct > 0 && monstre.VieAct > 0 {
			time.Sleep(1 * time.Second)
			joueur.Tour_Ennemi(monstre, CompteurTour)
			time.Sleep(1 * time.Second)
			if monstre.Tour_Joueur(joueur) {
				return true
			}
			joueur.ManaAct += 15
			CompteurTour++
		}
	}
	return false
}

func (monstre *Personnage) Tour_Joueur(joueur *Personnage) bool {
	var degats int
	if joueur.Est_Mort() {
		return true
	}
	Affichage("Combat", []string{"Que voulez vous faire ?", "1. Attaquer", "2. Attaque lourde", "3. Sorts", "4. Inventaire"}, false, false)
	var Choix = Choix(1, 4)
	switch Choix {
	case 1:
		monstre.VieAct -= joueur.Degats
		degats = joueur.Degats
	case 2:
		monstre.VieAct -= joueur.Degats * 2
		joueur.ManaAct -= 80
		degats = joueur.Degats * 2

	case 3:
		degats = Menu_Sort(joueur, monstre)
	case 4:
		joueur.Affichage_Inventaire()
	}
	if monstre.VieAct < 0 {
		monstre.VieAct = 0
		joueur.Ames += monstre.Ames
		switch monstre.Nom {
		case "Carcasse":
			joueur.Inv.Liste_Items["éclat de titanite"] += rand.Intn(2) * 3
		case "Chevalier mort-vivant":
			joueur.Inv.Liste_Items["grand éclat de titanite"] += rand.Intn(2) * 3
		case "Champion mort-vivant":
			joueur.Inv.Liste_Items["tablette de titanite"] += rand.Intn(2) * 3
		case "Gargouille":
			joueur.Inv.Liste_Items["éclat de titanite"] += rand.Intn(2) * 5
			joueur.Inv.Liste_Items["tablette de titanite"] += rand.Intn(2) * 2
			if !joueur.Inv.Liste_Armes[5].EstDebloque {
				joueur.Inv.Liste_Armes[5].EstDebloque = true
				Affichage("Combat", []string{"Vous avez débloqué la hache de guerre"}, true, false)
			}
		case "Démon Capra":
			joueur.Inv.Liste_Items["éclat de titanite"] += rand.Intn(2) * 5
			joueur.Inv.Liste_Items["grand éclat de titanite"] += rand.Intn(2) * 2
		case "Démon taureau":
			joueur.Inv.Liste_Items["éclat de titanite"] += rand.Intn(2) * 7
			joueur.Inv.Liste_Items["grand éclat de titanite"] += rand.Intn(2) * 7
			joueur.Inv.Liste_Items["tablette de titanite"] += rand.Intn(2) * 7
		}
		Affichage("Combat", []string{"Vous avez tué le " + monstre.Nom}, true, false)
	} else {
		Affichage(joueur.Nom, []string{"Vous avez infligé " + strconv.Itoa(degats) + " dégats au " + monstre.Nom + ", il lui reste " + strconv.Itoa(monstre.VieAct) + " de vie", "Vous avez maintenant " + strconv.Itoa(joueur.ManaAct) + " mana"}, false, false)
	}
	return false
}

func (joueur *Personnage) Tour_Ennemi(monstre *Personnage, CompteurTour int) {
	if joueur.Est_Mort() {
		return
	}
	if monstre.VieAct > 0 {
		if CompteurTour%3 == 0 {
			joueur.VieAct -= monstre.Degats * 2
			if joueur.VieAct > 0 {
				Affichage(monstre.Nom, []string{"Le " + monstre.Nom + " vous a infligé " + strconv.Itoa(monstre.Degats*2) + " dégats, vous avez maintenant " + strconv.Itoa(joueur.VieAct) + " de vie"}, false, false)
			}
		} else if CompteurTour%4 == 0 {
			Affichage(monstre.Nom, []string{"Le " + monstre.Nom + " vous a empoisonné"}, false, false)
			joueur.Inv.Liste_Consommables[5].Quantite++
			joueur.Prendre_Potion(monstre.Inv.Liste_Consommables[5])
		} else {
			joueur.VieAct -= monstre.Degats
			if joueur.VieAct < 0 {
				joueur.VieAct = 0
			}
			if joueur.VieAct > 0 {
				Affichage(monstre.Nom, []string{"Le " + monstre.Nom + " vous a infligé " + strconv.Itoa(monstre.Degats) + " dégats, vous avez maintenant " + strconv.Itoa(joueur.VieAct) + " de vie"}, false, false)
			}
		}
	}
}

func (p *Personnage) Faire_Combat() bool {
	boucle := true
	for boucle {
		c, _ := strconv.Atoi(p.Position.Val["monstre_nombre"])
		if c > 0 {
			Affichage("Déplacement", []string{"Nous nous dirigeons vers " + p.Position.Val["nom"], "Cet endroit est infesté de " + p.Position.Val["monstre_type"] + " il y en a " + p.Position.Val["monstre_nombre"], "Que voulez vous faire ?", "1. Combattre", "2. Fuir"}, false, false)
			var Choix = Choix(1, 2)
			if Choix == 1 {
				var monstre Personnage
				var inv Inventaire
				inv.Initialisation()
				monstre.Inv = inv
				monstre.Initialisation("Mob", p.Position.Val["monstre_type"])
				if Combat(p, &monstre) {
					return false
				}
				c--
				p.Position.Val["monstre_nombre"] = strconv.Itoa(c)
			} else if Choix == 2 {
				Affichage("Déplacement", []string{"Ce chemin à l'air dangereur, peut être que revenir en arrière était la bonne idée"}, true, true)
				boucle = false
			}
		}
		if c == 0 {
			Affichage("Déplacement", []string{"Nous arrivons vers " + p.Position.Val["nom"]}, true, true)
			return true
		}
	}
	return false
}
