package database

import (
	"fmt"
	"strconv"
)

type Marchand struct {
	Nom      string
	Inv      Inventaire
	Nb_trade int
}

func (m *Marchand) Init(nom string, inv Inventaire) {
	m.Nom = nom
	m.Nb_trade = 0
	m.Inv = inv
}

func (m *Marchand) InitMarchand(number int) {
	//Marchand 1
	var inv Inventaire
	inv.Init()
	if number == 1 {
		m.Init("Marchand mort vivant", inv)
		m.Inv.Liste_armes[4].Set_Armes("isUnlocked", "true")
		m.Inv.Liste_consommables[5].Set_Consommable("isUnlocked", "true")
		m.Inv.Liste_boucliers[2].Set_Boucliers("isUnlocked", "true")
		m.Nombre_Trade()
	}
}

func (m *Marchand) Nombre_Trade() {
	var nombre int
	for i := 0; i < len(m.Inv.Liste_armes); i++ {
		if m.Inv.Liste_armes[i].Get_Armes("isUnlocked") == "true" {
			nombre++
		}
	}
	for i := 0; i < len(m.Inv.Liste_boucliers); i++ {
		if m.Inv.Liste_boucliers[i].Get_Boucliers("isUnlocked") == "true" {
			nombre++
		}
	}
	for i := 0; i < len(m.Inv.Liste_consommables); i++ {
		if m.Inv.Liste_consommables[i].Get_Consommable("isUnlocked") == "true" {
			nombre++
		}
	}
	m.Nb_trade = nombre
}

func (m *Marchand) Affichage_Trade() {
	var liste []string
	liste = append(liste, "Que voulez-vous acheter ?")
	for i := 0; i < len(m.Inv.Liste_armes); i++ {
		if m.Inv.Liste_armes[i].Get_Armes("isUnlocked") == "true" {
			liste = append(liste, strconv.Itoa(len(liste))+". "+m.Inv.Liste_armes[i].nom+" | "+strconv.Itoa(m.Inv.Liste_armes[i].Prix)+" âmes")
		}
	}
	for i := 0; i < len(m.Inv.Liste_boucliers); i++ {
		if m.Inv.Liste_boucliers[i].Get_Boucliers("isUnlocked") == "true" {
			liste = append(liste, strconv.Itoa(len(liste))+". "+m.Inv.Liste_boucliers[i].nom+" | "+strconv.Itoa(m.Inv.Liste_boucliers[i].Prix)+" âmes")
		}
	}
	for i := 0; i < len(m.Inv.Liste_consommables); i++ {
		if m.Inv.Liste_consommables[i].Get_Consommable("isUnlocked") == "true" {
			liste = append(liste, strconv.Itoa(len(liste))+". "+m.Inv.Liste_consommables[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_consommables[i].Prix)+" âmes")
		}
	}
	liste = append(liste, strconv.Itoa(len(liste))+". Quitter le menu d'achat")
	Affichage("Marchand", liste)
}

func (m *Marchand) Trade(player *Personnage) {
	m.Affichage_Trade()
	var choix int
	fmt.Scan(&choix)
	for choix < 1 || choix > m.Nb_trade+1 {
		Affichage("Erreur", []string{"Vous devez choisir un choix entre 1 et " + strconv.Itoa(m.Nb_trade+1)})
		fmt.Scan(&choix)
	}
	if choix == m.Nb_trade+1 {
		m.Menu_Marchand(player)
	}
	for i := 0; i < len(m.Inv.Liste_armes); i++ {
		if m.Inv.Liste_armes[i].Get_Armes("isUnlocked") == "true" {
			choix--
			if choix == 0 {
				if WantTrade(player, player.Inv.Liste_armes[i]) {
					print("test want trade")
					if IsTradeable(player, "Armes", i) {
						print("test is tradeable")
						player.Inv.Liste_armes[i].Set_Armes("isUnlocked", "true")
						player.Ames -= m.Inv.Liste_armes[i].Prix
						Affichage("Succès", []string{"Vous avez acheté un objet"})
						Attendre()
						return
					} else {
						m.Trade(player)
					}
				} else {
					m.Trade(player)
				}
			}
		}
	}
	for i := 0; i < len(m.Inv.Liste_boucliers); i++ {
		if m.Inv.Liste_boucliers[i].Get_Boucliers("isUnlocked") == "true" {
			choix--
			if choix == 0 {
				if WantTrade(player, player.Inv.Liste_boucliers[i]) {
					print("test want trade")
					if IsTradeable(player, "Boucliers", i) {
						print("test is tradeable")
						player.Inv.Liste_boucliers[i].Set_Boucliers("isUnlocked", "true")
						player.Ames -= m.Inv.Liste_boucliers[i].Prix
						Affichage("Succès", []string{"Vous avez acheté un objet"})
						Attendre()
						return
					} else {
						m.Trade(player)
					}
				} else {
					m.Trade(player)
				}
			}
		}
	}
	for i := 0; i < len(m.Inv.Liste_consommables); i++ {
		if m.Inv.Liste_consommables[i].Get_Consommable("isUnlocked") == "true" {
			choix--
			if choix == 0 {
				if WantTrade(player, player.Inv.Liste_consommables[i]) {
					print("test want trade")
					if IsTradeable(player, "Consommables", i) {
						print("test is tradeable")
						player.Inv.Liste_consommables[i].Set_Consommable("isUnlocked", "true")
						player.Ames -= m.Inv.Liste_consommables[i].Prix
						Affichage("Succès", []string{"Vous avez acheté un objet"})
						Attendre()
						return
					} else {
						m.Trade(player)
					}
				} else {
					m.Trade(player)
				}
			}
		}
	}
}

func WantTrade(player *Personnage, item interface{}) bool {
	//vérifie si item est une structure Armes, Boucliers ou Consommable
	switch item := item.(type) {
	case Armes:
		Affichage("Marchand", []string{"Que voulez-vous faire ?", "1. Acheter l'item " + item.nom, "2. Afficher les caractéristiques de l'objet", "3. Revenir en arrière"})
	case Boucliers:
		Affichage("Marchand", []string{"Que voulez-vous faire ?", "1. Acheter l'item " + item.nom, "2. Afficher les caractéristiques de l'objet", "3. Revenir en arrière"})
	case Consommable:
		Affichage("Marchand", []string{"Que voulez-vous faire ?", "1. Acheter l'item " + item.Nom, "2. Afficher les caractéristiques de l'objet", "3. Revenir en arrière"})
	}
	var choix int
	fmt.Scan(&choix)
	for choix < 1 || choix > 3 {
		Affichage("Erreur", []string{"Vous devez choisir un choix entre 1 et 3"})
		fmt.Scan(&choix)
	}
	switch choix {
	case 1:
		return true
	case 2:
		switch item := item.(type) {
		case Armes:
			item.Affichage()
			Attendre()
		case Boucliers:
			item.Affichage()
			Attendre()
		case Consommable:
			item.Affichage()
			Attendre()
		default:
			print("Erreur")
		}
	}
	return false
}

func IsTradeable(player *Personnage, item_type string, index int) bool {
	tradeable := false
	if item_type == "Armes" {
		if IsBuyable(player.Ames, player.Inv.Liste_armes[index].Prix) && !player.Inv.Liste_armes[index].IsUnlocked {
			tradeable = true
		} else if player.Inv.Liste_armes[index].IsUnlocked {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"})
			Attendre()
		}
	} else if item_type == "Boucliers" {
		if IsBuyable(player.Ames, player.Inv.Liste_boucliers[index].Prix) && !player.Inv.Liste_boucliers[index].IsUnlocked {
			tradeable = true
		} else if player.Inv.Liste_boucliers[index].IsUnlocked {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"})
			Attendre()
		}
	} else if item_type == "Consommables" {
		if IsBuyable(player.Ames, player.Inv.Liste_consommables[index].Prix) && player.Inv.Liste_consommables[index].Quantite <= 100 {
			tradeable = true
		} else if player.Inv.Liste_consommables[index].Quantite > 100 {
			Affichage("Marchand", []string{"Vous possédez déjà la quantité maximale de cet objet"})
			Attendre()
		}
	}
	return tradeable
}

func IsBuyable(Ames, Prix int) bool {
	if Ames <= Prix {
		Affichage("Erreur", []string{"Vous n'avez pas assez d'âmes pour acheter cet objet"})
		Attendre()
		return false
	}
	return true
}

func (m *Marchand) Menu_Marchand(p *Personnage) {
	Affichage("Menu du marchand", []string{"Que voulez-vous faire ?", "1. Acheter un objet", "2. Quitter le menu du marchand"})
	var choix int
	fmt.Scan(&choix)
	for choix < 1 || choix > 2 {
		Affichage("Erreur", []string{"Vous devez choisir un choix entre 1 et 2"})
		fmt.Scan(&choix)
	}
	if choix == 1 {
		m.Trade(p)
	}
}
