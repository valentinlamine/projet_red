package database

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rivo/uniseg"
)

type Marchand struct {
	Nom          string
	Inv          Inventaire
	NombreTrade  int
	PremierTrade bool
}

func (m *Marchand) Initialisation(Nom string, inv Inventaire) {
	m.Nom = Nom
	m.NombreTrade = 0
	m.Inv = inv
	m.PremierTrade = true
}

func (m *Marchand) Initialisation_Marchand(nombre int) {
	//Marchand 1
	var inv Inventaire
	inv.Initialisation()
	if nombre == 1 {
		m.Initialisation("Marchand mort vivant", inv)
		//Armes
		m.Inv.Liste_Armes[0].EstDebloque = true
		m.Inv.Liste_Armes[1].EstDebloque = true
		m.Inv.Liste_Armes[2].EstDebloque = true
		m.Inv.Liste_Armes[3].EstDebloque = true
		m.Inv.Liste_Armes[5].EstDebloque = true
		//Boucliers
		m.Inv.Liste_Boucliers[0].EstDebloque = true
		m.Inv.Liste_Boucliers[1].EstDebloque = true
		m.Inv.Liste_Boucliers[2].EstDebloque = true
		m.Inv.Liste_Boucliers[3].EstDebloque = true
		m.Inv.Liste_Boucliers[4].EstDebloque = true
		m.Inv.Liste_Boucliers[5].EstDebloque = true
		//première armure
		m.Inv.Liste_Armures_Tete[1].EstDebloque = true
		m.Inv.Liste_Armures_Torse[1].EstDebloque = true
		m.Inv.Liste_Armures_Jambes[1].EstDebloque = true
		m.Inv.Liste_Armures_Bras[1].EstDebloque = true
		//deuxième armure
		m.Inv.Liste_Armures_Tete[2].EstDebloque = true
		m.Inv.Liste_Armures_Torse[2].EstDebloque = true
		m.Inv.Liste_Armures_Jambes[2].EstDebloque = true
		m.Inv.Liste_Armures_Bras[2].EstDebloque = true
		//troisième armure
		m.Inv.Liste_Armures_Tete[3].EstDebloque = true
		m.Inv.Liste_Armures_Torse[3].EstDebloque = true
		m.Inv.Liste_Armures_Jambes[3].EstDebloque = true
		m.Inv.Liste_Armures_Bras[3].EstDebloque = true
		//quatrième armure
		m.Inv.Liste_Armures_Tete[4].EstDebloque = true
		m.Inv.Liste_Armures_Torse[4].EstDebloque = true
		m.Inv.Liste_Armures_Jambes[4].EstDebloque = true
		m.Inv.Liste_Armures_Bras[4].EstDebloque = true
		//cinquième armure
		m.Inv.Liste_Armures_Tete[5].EstDebloque = true
		m.Inv.Liste_Armures_Torse[5].EstDebloque = true
		m.Inv.Liste_Armures_Jambes[5].EstDebloque = true
		m.Inv.Liste_Armures_Bras[5].EstDebloque = true
		m.PremierTrade = false
		m.NombreTrade = 33
	} else if nombre == 2 {
		m.Initialisation("Gardienne du feu", inv)
		m.Inv.Liste_Consommables[1].Quantite = 100
		m.Inv.Liste_Consommables[2].Quantite = 100
		m.Inv.Liste_Consommables[3].Quantite = 100
		m.Inv.Liste_Consommables[4].Quantite = 100
		m.PremierTrade = false
		m.NombreTrade = 4
	} else if nombre == 3 {
		m.Initialisation("Laurentius", inv)
		m.Inv.Liste_Sorts[1].EstDebloque = true
		m.Inv.Liste_Sorts[2].EstDebloque = true
		m.Inv.Liste_Sorts[3].EstDebloque = true
		m.Inv.Liste_Sorts[4].EstDebloque = true
		m.Inv.Liste_Sorts[5].EstDebloque = true
		m.PremierTrade = false
		m.NombreTrade = 5
	} else if nombre == 4 {
		m.Initialisation("Petrus", inv)
		m.Inv.Liste_Consommables[0].Quantite = 100
		m.Inv.Liste_Consommables[5].Quantite = 100
		m.NombreTrade = 2
	}
}

func (m *Marchand) Affichage_Echange(joueur *Personnage) {
	var liste []string
	liste = append(liste, "Que voulez-vous acheter ?")
	liste = append(liste, "Vous possédez : "+strconv.Itoa(joueur.Ames)+" Ames")
	liste = append(liste, "")
	liste = append(liste, "Armes :")
	for i := 0; i < len(m.Inv.Liste_Armes); i++ {
		if m.Inv.Liste_Armes[i].EstDebloque {
			liste = append(liste, strconv.Itoa(len(liste)-3)+". "+m.Inv.Liste_Armes[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_Armes[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "")
	liste = append(liste, "Boucliers :")
	for i := 0; i < len(m.Inv.Liste_Boucliers); i++ {
		if m.Inv.Liste_Boucliers[i].EstDebloque {
			liste = append(liste, strconv.Itoa(len(liste)-5)+". "+m.Inv.Liste_Boucliers[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_Boucliers[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "")
	if m.Nom == "Gardienne du feu" {
		liste = append(liste, "Livre de niveau :")
	} else {
		liste = append(liste, "Consommables :")
	}
	for i := 0; i < len(m.Inv.Liste_Consommables); i++ {
		if m.Inv.Liste_Consommables[i].Quantite > 0 {
			liste = append(liste, strconv.Itoa(len(liste)-7)+". "+m.Inv.Liste_Consommables[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_Consommables[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "")
	liste = append(liste, "Casque :")
	for i := 0; i < len(m.Inv.Liste_Armures_Tete); i++ {
		if m.Inv.Liste_Armures_Tete[i].EstDebloque {
			liste = append(liste, strconv.Itoa(len(liste)-9)+". "+m.Inv.Liste_Armures_Tete[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_Armures_Tete[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "")
	liste = append(liste, "Plastrons :")
	for i := 0; i < len(m.Inv.Liste_Armures_Torse); i++ {
		if m.Inv.Liste_Armures_Torse[i].EstDebloque {
			liste = append(liste, strconv.Itoa(len(liste)-11)+". "+m.Inv.Liste_Armures_Torse[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_Armures_Torse[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "")
	liste = append(liste, "Jambières :")
	for i := 0; i < len(m.Inv.Liste_Armures_Jambes); i++ {
		if m.Inv.Liste_Armures_Jambes[i].EstDebloque {
			liste = append(liste, strconv.Itoa(len(liste)-13)+". "+m.Inv.Liste_Armures_Jambes[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_Armures_Jambes[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "")
	liste = append(liste, "Brassards :")
	for i := 0; i < len(m.Inv.Liste_Armures_Bras); i++ {
		if m.Inv.Liste_Armures_Bras[i].EstDebloque {
			liste = append(liste, strconv.Itoa(len(liste)-15)+". "+m.Inv.Liste_Armures_Bras[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_Armures_Bras[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "")
	liste = append(liste, "Sorts :")
	for i := 0; i < len(m.Inv.Liste_Sorts); i++ {
		if m.Inv.Liste_Sorts[i].EstDebloque {
			liste = append(liste, strconv.Itoa(len(liste)-17)+". "+m.Inv.Liste_Sorts[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_Sorts[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "", strconv.Itoa(len(liste)-17)+". Quitter le menu d'achat")
	Affichage(m.Nom, liste, true, false)
}

func (m *Marchand) Echange(joueur *Personnage) {
	if m.PremierTrade {
		m.Est_Premier_Echange(joueur)
	}
	m.Affichage_Echange(joueur)
	var choix = Choix(1, m.NombreTrade+1)
	if choix == m.NombreTrade+1 {
		m.Menu_Marchand(joueur)
	}
	for i := 0; i < len(m.Inv.Liste_Armes); i++ {
		if m.Inv.Liste_Armes[i].EstDebloque {
			choix--
			if choix == 0 {
				if Souhaite_Echanger(joueur, joueur.Inv.Liste_Armes[i]) {
					if Est_Echangeable(joueur, "Armes", i) {
						joueur.Inv.Liste_Armes[i].EstDebloque = true
						joueur.Ames -= m.Inv.Liste_Armes[i].Prix
						Affichage("Succès", []string{"Vous avez acheté l'arme " + m.Inv.Liste_Armes[i].Nom}, true, true)
						m.Echange(joueur)
						return
					} else {
						m.Echange(joueur)
					}
				} else {
					m.Echange(joueur)
				}
			}
		}
	}
	for i := 0; i < len(m.Inv.Liste_Boucliers); i++ {
		if m.Inv.Liste_Boucliers[i].EstDebloque {
			choix--
			if choix == 0 {
				if Souhaite_Echanger(joueur, joueur.Inv.Liste_Boucliers[i]) {
					if Est_Echangeable(joueur, "Boucliers", i) {
						joueur.Inv.Liste_Boucliers[i].EstDebloque = true
						joueur.Ames -= m.Inv.Liste_Boucliers[i].Prix
						Affichage("Succès", []string{"Vous avez acheté le bouclier " + m.Inv.Liste_Boucliers[i].Nom}, true, true)
						m.Echange(joueur)
						return
					} else {
						m.Echange(joueur)
					}
				} else {
					m.Echange(joueur)
				}
			}
		}
	}
	for i := 0; i < len(m.Inv.Liste_Consommables); i++ {
		if m.Inv.Liste_Consommables[i].Quantite > 0 {
			choix--
			if choix == 0 {
				if Souhaite_Echanger(joueur, joueur.Inv.Liste_Consommables[i]) {
					if Est_Echangeable(joueur, "Consommables", i) {
						joueur.Inv.Liste_Consommables[i].Quantite++
						joueur.Ames -= m.Inv.Liste_Consommables[i].Prix
						if i > 0 && i < 5 {
							m.Inv.Liste_Consommables[1].Prix += (joueur.Niveau * joueur.Niveau) / 2
							m.Inv.Liste_Consommables[2].Prix += (joueur.Niveau * joueur.Niveau) / 2
							m.Inv.Liste_Consommables[3].Prix += (joueur.Niveau * joueur.Niveau) / 2
							m.Inv.Liste_Consommables[4].Prix += (joueur.Niveau * joueur.Niveau) / 2
							joueur.Prendre_Potion(joueur.Inv.Liste_Consommables[i])
							joueur.Inv.Liste_Consommables[i].Quantite--
							m.Echange(joueur)
							return
						} else {
							Affichage("Succès", []string{"Vous avez acheté le consommable " + m.Inv.Liste_Consommables[i].Nom}, true, true)
							m.Echange(joueur)
							return
						}
					} else {
						m.Echange(joueur)
					}
				} else {
					m.Echange(joueur)
				}
			}
		}
	}
	for i := 0; i < len(m.Inv.Liste_Armures_Tete); i++ {
		if m.Inv.Liste_Armures_Tete[i].EstDebloque {
			choix--
			if choix == 0 {
				if Souhaite_Echanger(joueur, joueur.Inv.Liste_Armures_Tete[i]) {
					if Est_Echangeable(joueur, "Armures_tete", i) {
						joueur.Inv.Liste_Armures_Tete[i].EstDebloque = true
						joueur.Ames -= m.Inv.Liste_Armures_Tete[i].Prix
						Affichage("Succès", []string{"Vous avez acheté un objet"}, true, true)
						return
					} else {
						m.Echange(joueur)
					}
				} else {
					m.Echange(joueur)
				}
			}
		}
	}
	for i := 0; i < len(m.Inv.Liste_Armures_Torse); i++ {
		if m.Inv.Liste_Armures_Torse[i].EstDebloque {
			choix--
			if choix == 0 {
				if Souhaite_Echanger(joueur, joueur.Inv.Liste_Armures_Torse[i]) {
					if Est_Echangeable(joueur, "Armures_torse", i) {
						joueur.Inv.Liste_Armures_Torse[i].EstDebloque = true
						joueur.Ames -= m.Inv.Liste_Armures_Torse[i].Prix
						Affichage("Succès", []string{"Vous avez acheté un objet"}, true, true)
						return
					} else {
						m.Echange(joueur)
					}
				} else {
					m.Echange(joueur)
				}
			}
		}
	}
	for i := 0; i < len(m.Inv.Liste_Armures_Jambes); i++ {
		if m.Inv.Liste_Armures_Jambes[i].EstDebloque {
			choix--
			if choix == 0 {
				if Souhaite_Echanger(joueur, joueur.Inv.Liste_Armures_Jambes[i]) {
					if Est_Echangeable(joueur, "Armures_jambes", i) {
						joueur.Inv.Liste_Armures_Jambes[i].EstDebloque = true
						joueur.Ames -= m.Inv.Liste_Armures_Jambes[i].Prix
						Affichage("Succès", []string{"Vous avez acheté un objet"}, true, true)
						return
					} else {
						m.Echange(joueur)
					}
				} else {
					m.Echange(joueur)
				}
			}
		}
	}
	for i := 0; i < len(m.Inv.Liste_Armures_Bras); i++ {
		if m.Inv.Liste_Armures_Bras[i].EstDebloque {
			choix--
			if choix == 0 {
				if Souhaite_Echanger(joueur, joueur.Inv.Liste_Armures_Bras[i]) {
					if Est_Echangeable(joueur, "Armures_bras", i) {
						joueur.Inv.Liste_Armures_Bras[i].EstDebloque = true
						joueur.Ames -= m.Inv.Liste_Armures_Bras[i].Prix
						Affichage("Succès", []string{"Vous avez acheté un objet"}, true, true)
						return
					} else {
						m.Echange(joueur)
					}
				} else {
					m.Echange(joueur)
				}
			}
		}
	}
	for i := 0; i < len(m.Inv.Liste_Sorts); i++ {
		if m.Inv.Liste_Sorts[i].EstDebloque {
			choix--
			if choix == 0 {
				if Souhaite_Echanger(joueur, joueur.Inv.Liste_Sorts[i]) {
					if Est_Echangeable(joueur, "Sort", i) {
						joueur.Inv.Liste_Sorts[i].EstDebloque = true
						joueur.Ames -= m.Inv.Liste_Sorts[i].Prix
						Affichage("Succès", []string{"Vous avez acheté un objet"}, true, true)
						return
					} else {
						m.Echange(joueur)
					}
				} else {
					m.Echange(joueur)
				}
			}
		}
	}
}

func Souhaite_Echanger(joueur *Personnage, item interface{}) bool {
	//vérifie si item est une structure Armes, Boucliers ou Consommable
	switch item := item.(type) {
	case Armes:
		Affichage("Marchand", []string{"Que voulez-vous faire ?", "1. Acheter l'item " + item.Nom, "2. Afficher les caractéristiques de l'objet", "3. Revenir en arrière"}, false, false)
	case Boucliers:
		Affichage("Marchand", []string{"Que voulez-vous faire ?", "1. Acheter l'item " + item.Nom, "2. Afficher les caractéristiques de l'objet", "3. Revenir en arrière"}, false, false)
	case Consommable:
		Affichage("Marchand", []string{"Que voulez-vous faire ?", "1. Acheter l'item " + item.Nom, "2. Afficher les caractéristiques de l'objet", "3. Revenir en arrière"}, false, false)
	case Armures:
		Affichage("Marchand", []string{"Que voulez-vous faire ?", "1. Acheter l'item " + item.Nom, "2. Afficher les caractéristiques de l'objet", "3. Revenir en arrière"}, false, false)
	case Sort:
		Affichage("Marchand", []string{"Que voulez-vous faire ?", "1. Acheter l'item " + item.Nom, "2. Afficher les caractéristiques de l'objet", "3. Revenir en arrière"}, false, false)
	}
	var choix = Choix(1, 3)
	switch choix {
	case 1:
		return true
	case 2:
		switch item := item.(type) {
		case Armes:
			item.Affichage()
		case Boucliers:
			item.Affichage()
		case Consommable:
			item.Affichage()
		case Armures:
			item.Affichage()
		case Sort:
			item.Affichage()
		default:
			print("Erreur")
		}
	}
	return false
}

func Est_Echangeable(joueur *Personnage, item_type string, index int) bool {
	tradeable := false
	if item_type == "Armes" {
		if Est_Achetable(joueur.Ames, joueur.Inv.Liste_Armes[index].Prix) && !joueur.Inv.Liste_Armes[index].EstDebloque {
			tradeable = true
		} else if joueur.Inv.Liste_Armes[index].EstDebloque {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"}, true, true)
		}
	} else if item_type == "Boucliers" {
		if Est_Achetable(joueur.Ames, joueur.Inv.Liste_Boucliers[index].Prix) && !joueur.Inv.Liste_Boucliers[index].EstDebloque {
			tradeable = true
		} else if joueur.Inv.Liste_Boucliers[index].EstDebloque {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"}, true, true)
		}
	} else if item_type == "Consommables" {
		if Est_Achetable(joueur.Ames, joueur.Inv.Liste_Consommables[index].Prix) && joueur.Inv.Liste_Consommables[index].Quantite <= 100 {
			tradeable = true
		} else if joueur.Inv.Liste_Consommables[index].Quantite > 100 {
			Affichage("Marchand", []string{"Vous possédez déjà la quantité maximale de cet objet"}, true, true)
		}
	} else if item_type == "Armures_tete" {
		if Est_Achetable(joueur.Ames, joueur.Inv.Liste_Armures_Tete[index].Prix) && !joueur.Inv.Liste_Armures_Tete[index].EstDebloque {
			tradeable = true
		} else if joueur.Inv.Liste_Armures_Tete[index].EstDebloque {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"}, true, true)
		}
	} else if item_type == "Armures_torse" {
		if Est_Achetable(joueur.Ames, joueur.Inv.Liste_Armures_Torse[index].Prix) && !joueur.Inv.Liste_Armures_Torse[index].EstDebloque {
			tradeable = true
		} else if joueur.Inv.Liste_Armures_Torse[index].EstDebloque {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"}, true, true)
		}
	} else if item_type == "Armures_jambes" {
		if Est_Achetable(joueur.Ames, joueur.Inv.Liste_Armures_Jambes[index].Prix) && !joueur.Inv.Liste_Armures_Jambes[index].EstDebloque {
			tradeable = true
		} else if joueur.Inv.Liste_Armures_Jambes[index].EstDebloque {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"}, true, true)
		}
	} else if item_type == "Armures_bras" {
		if Est_Achetable(joueur.Ames, joueur.Inv.Liste_Armures_Bras[index].Prix) && !joueur.Inv.Liste_Armures_Bras[index].EstDebloque {
			tradeable = true
		} else if joueur.Inv.Liste_Armures_Bras[index].EstDebloque {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"}, true, true)
		}
	} else if item_type == "Sort" {
		if Est_Achetable(joueur.Ames, joueur.Inv.Liste_Sorts[index].Prix) && !joueur.Inv.Liste_Sorts[index].EstDebloque {
			tradeable = true
		} else if joueur.Inv.Liste_Sorts[index].EstDebloque {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"}, true, true)
		}
	}
	return tradeable
}

func Est_Achetable(Ames, Prix int) bool {
	if Ames <= Prix {
		Affichage("Erreur", []string{"Vous n'avez pas assez d'âmes pour acheter cet objet"}, true, true)
		return false
	}
	return true
}

func (m *Marchand) Est_Premier_Echange(joueur *Personnage) {
	Affichage(m.Nom, []string{"Bienvenue dans mon magasin !", "Je suis le marchand " + m.Nom + ".", "Comme vous avez l'air d'être nouveau, laissez moi vous faire un petit cadeau !"}, true, true)
	joueur.Inv.Liste_Consommables[0].Quantite += 1
	m.PremierTrade = false
	Affichage("Informations", []string{"Félicitations !!!", "Vous avez obtenu une Fiole d'Estus gratuitement !"}, true, true)
}

func (m *Marchand) Menu_Marchand(p *Personnage) {
	Affichage("Menu du marchand", []string{"Que voulez-vous faire ?", "1. Acheter un objet", "2. Parler", "3. Quitter le menu du marchand"}, true, false)
	var choix = Choix(1, 3)
	if choix == 1 {
		m.Echange(p)
	} else if choix == 2 {
		if m.Nom == "Petrus" {
			Affichage(m.Nom, []string{"Je suis le marchand " + m.Nom + ".", "Je suis un ancien chevalier de l'Ordre de la Lumière.", "Je suis venu ici pour me reposer et me soigner.", "Je ne peux plus combattre, mais je peux vous aider à vous soigner."}, true, true)
			if p.Inv.Liste_Consommables[0].Quantite < 3 {
				Affichage(m.Nom, []string{"Je vois que vos fioles sont vides...", "Laissez moi vous remplir !"}, true, true)
				p.Inv.Liste_Consommables[0].Quantite = 3
				Affichage("Informations", []string{"Vous possédez maintenant 3 fioles d'Estus !"}, true, true)
			}
		} else if m.Nom == "Laurentius" {
			Affichage(m.Nom, []string{"Hmm hmmm hmmmm..."}, true, false)
			time.Sleep(3 * time.Second)
			Attendre()
			Affichage(m.Nom, []string{"Oh excusez moi... J'étais perdu dans mes pensées."}, true, true)
			Affichage(m.Nom, []string{"Je suis Laurentius. J'ai des amis formidables vous savez."}, true, true)
			Affichage(m.Nom, []string{"Albertus Brutus Baldus Alphus... Un grand musicien.", "Mais vous connaissez sûrement beaucoup plus Steven Spielberg."}, true, true)
			Affichage(m.Nom, []string{"Grand réalisateur de film, Ready Player One, Jurassic Park...", "Je suis un fan de Steven Spielberg."}, true, false)
			time.Sleep(2 * time.Second)
			Attendre()
			Affichage(m.Nom, []string{"Je suis un grand fan de Steven Spielberg..."}, true, false)
			time.Sleep(2 * time.Second)
			Attendre()
			Affichage(m.Nom, []string{"Bref, je suis un grand fan de Steven Spielberg."}, true, true)
			Affichage(m.Nom, []string{"Où en étais-je ?"}, true, true)
		} else if m.Nom == "Gardienne du feu" {
			Affichage(m.Nom, []string{"Je suis la gardienne du feu.", "Je peux améliorer tes compétences contre de l'argent."}, true, true)
			Affichage(m.Nom, []string{"Si tu veux apprendre de nouveaux sorts, tu peux aller voir mon ami Laurientus.", "de ce qu'il m'a dit, il est un grand fan de Steven Spielberg."}, true, true)
		} else {
			Affichage("Marchand", []string{"Heh Heh Heh...",
				"Si vous êtes en manque d'âmes, tuez des monstres !",
				"Plus ils sont gros, plus ils ont d'âmes",
				"Donnez les moi je donnerai des objets de valeur... Heh Heh Heh..."}, true, true)
		}
		m.Menu_Marchand(p)
	}
}

func (p *Personnage) Forgeron_Amelioration() {
	var Nom string
	longueur := 50
	var list_objets []string
	nb_objets := 0
	//laisser 8 ligne de vide au dessus
	fmt.Print("\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n")
	//affichage du haut de la boite
	fmt.Print("╒", strings.Repeat("═", longueur), "╕", "\n")
	//affichage du titre
	fmt.Print("│", "Inventaire", strings.Repeat(" ", longueur-10), "│", "\n")
	//affichage de la ligne
	fmt.Print("╞", strings.Repeat("═", longueur), "╡", "\n")
	//affichage des armes
	fmt.Print("│", "Armes Améliorables :", strings.Repeat(" ", longueur-20), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for i := 0; i < len(p.Inv.Liste_Armes); i++ {
		if p.Possede_Item(p.Inv.Liste_Armes[i].Niv) && p.Inv.Liste_Armes[i].EstDebloque {
			Nom = p.Inv.Liste_Armes[i].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage des boucliers
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Boucliers Améliorables :", strings.Repeat(" ", longueur-24), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for i := 0; i < len(p.Inv.Liste_Boucliers); i++ {
		if p.Possede_Item(p.Inv.Liste_Boucliers[i].Niv) && p.Inv.Liste_Boucliers[i].EstDebloque {
			Nom = p.Inv.Liste_Boucliers[i].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage des casques
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Casques Améliorables :", strings.Repeat(" ", longueur-22), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for i := 0; i < len(p.Inv.Liste_Armures_Tete); i++ {
		if p.Possede_Item(p.Inv.Liste_Armures_Tete[i].Niv) && p.Inv.Liste_Armures_Tete[i].EstDebloque {
			Nom = p.Inv.Liste_Armures_Tete[i].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage des plastrons
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Plastrons Améliorables :", strings.Repeat(" ", longueur-24), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for i := 0; i < len(p.Inv.Liste_Armures_Torse); i++ {
		if p.Possede_Item(p.Inv.Liste_Armures_Torse[i].Niv) && p.Inv.Liste_Armures_Torse[i].EstDebloque {
			Nom = p.Inv.Liste_Armures_Torse[i].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage des jambières
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Jambières Améliorables :", strings.Repeat(" ", longueur-24), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for i := 0; i < len(p.Inv.Liste_Armures_Jambes); i++ {
		if p.Possede_Item(p.Inv.Liste_Armures_Jambes[i].Niv) && p.Inv.Liste_Armures_Jambes[i].EstDebloque {
			Nom = p.Inv.Liste_Armures_Jambes[i].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage des brassards
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Brassards Améliorables :", strings.Repeat(" ", longueur-24), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for i := 0; i < len(p.Inv.Liste_Armures_Bras); i++ {
		if p.Possede_Item(p.Inv.Liste_Armures_Bras[i].Niv) && p.Inv.Liste_Armures_Bras[i].EstDebloque {
			Nom = p.Inv.Liste_Armures_Bras[i].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage du bas du tableau
	fmt.Print("└", strings.Repeat("─", longueur), "┘", "\n")

	//choix de l'objet à améliorer
	fmt.Print("Choisissez l'objet à améliorer (0 pour quitter) : ")
	var choix = Choix(0, len(list_objets))
	if choix != 0 {
		item := p.Inv.Recuperer_Item(list_objets[choix-1])
		Affichage("Inventaire", []string{"Vous avez choisi l'objet : " + list_objets[choix-1], "Que voulez vous faire avec cet objet ? ", "1. L'améliorer", "2. Afficher ces statistiques", "3. Retourner au menu précédent"}, true, false)
		var choix2 = Choix(1, 3)
		switch choix2 {
		case 1:
			switch item := item.(type) {
			case Armes:
				item.Ameliorer_Arme(p)
				p.Remplacer_Item(item)
			case Boucliers:
				item.Ameliorer_Bouclier(p)
				p.Remplacer_Item(item)
			case Armures:
				item.Ameliorer_Armures(p)
				p.Remplacer_Item(item)

			}
		case 2:
			switch item := item.(type) {
			case Armes:
				item.Affichage()

				p.Forgeron_Amelioration()
			case Boucliers:
				item.Affichage()
				p.Forgeron_Amelioration()
			case Armures:
				item.Affichage()
				p.Forgeron_Amelioration()
			}
		case 3:
			p.Forgeron_Amelioration()
		}
	}
}

func (p *Personnage) Possede_Item(nombre int) bool {
	if nombre > 3 {
		return false
	}
	switch nombre {
	case 1:
		if p.Inv.Liste_Items["éclat de titanite"] >= 6 && p.Ames >= 100 {
			return true
		} else {
			return false
		}
	case 2:
		if p.Inv.Liste_Items["grand éclat de titanite"] >= 3 && p.Inv.Liste_Items["éclat de titanite"] >= 3 && p.Ames >= 500 {
			return true
		} else {
			return false
		}
	case 3:
		if p.Inv.Liste_Items["grand éclat de titanite"] >= 2 && p.Inv.Liste_Items["éclat de titanite"] >= 2 && p.Inv.Liste_Items["tablette de titanite"] >= 2 && p.Ames >= 2000 {
			return true
		} else {
			return false
		}
	}
	return false
}
