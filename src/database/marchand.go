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

// initialise les marchands
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
	fmt.Print(liste)
	Affichage("Marchand", liste)
}

func (m *Marchand) Trade(player *Personnage) {
	m.Affichage_Trade()
	var choix int
	fmt.Scan(&choix)
	for choix < 1 || choix > m.Nb_trade {
		Affichage("Erreur", []string{"Vous devez choisir un choix entre 1 et" + string(m.Nb_trade)})
		fmt.Scan(&choix)
	}
	for i := 0; i < len(m.Inv.Liste_armes); i++ {
		if m.Inv.Liste_armes[i].Get_Armes("isUnlocked") == "true" {
			choix--
			if choix == 0 {
				if !IsBuyable(player.Ames, m.Inv.Liste_armes[i].Prix) {
					m.Trade(player)
				} else if player.Inv.Liste_armes[i].Get_Armes("isUnlocked") == "true" {
					Affichage("Erreur", []string{"Vous possédez déjà cet objet"})
					m.Trade(player)
				} else {
					player.Inv.Liste_armes[i].Set_Armes("isUnlocked", "true")
					return
				}
			}
		}
	}
	for i := 0; i < len(m.Inv.Liste_boucliers); i++ {
		if m.Inv.Liste_boucliers[i].Get_Boucliers("isUnlocked") == "true" {
			choix--
			if choix == 0 {
				if !IsBuyable(player.Ames, m.Inv.Liste_boucliers[i].Prix) {
					m.Trade(player)
				} else if player.Inv.Liste_boucliers[i].Get_Boucliers("isUnlocked") == "true" {
					Affichage("Erreur", []string{"Vous possédez déjà cet objet"})
					m.Trade(player)
				} else {
					player.Inv.Liste_boucliers[i].Set_Boucliers("isUnlocked", "true")
					return
				}
			}
		}
	}
	for i := 0; i < len(m.Inv.Liste_consommables); i++ {
		if m.Inv.Liste_consommables[i].Get_Consommable("isUnlocked") == "true" {
			choix--
			if choix == 0 {
				if !IsBuyable(player.Ames, m.Inv.Liste_consommables[i].Prix) {
					m.Trade(player)
				} else if player.Inv.Liste_consommables[i].Get_Consommable("isUnlocked") == "true" {
					Affichage("Erreur", []string{"Vous possédez déjà cet objet"})
					m.Trade(player)
				} else {
					player.Inv.Liste_consommables[i].Set_Consommable("isUnlocked", "true")
					return
				}
			}
		}
	}
}

func IsBuyable(Ames, Prix int) bool {
	if Ames <= Prix {
		Affichage("Erreur", []string{"Vous n'avez pas assez d'âmes pour acheter cet objet"})
		return false
	}
	return true
}

// Affichage("Menu", []string{"Que voulez-vous faire ?", "1. Accéder aux statistiques du personnage", "2. Accéder à l'inventaire du personnage", "3. Quitter le jeu", "4. Boire une potion de soin", "5. Se déplacer", "6. Aller voir le marchand mort vivant"})
func (m *Marchand) Menu_Marchand(p *Personnage) {
	Affichage("Menu du marchand", []string{"Que voulez-vous faire ?", "1. Acheter un objet", "2. Quitter le menu du marchand"})
	var choix int
	fmt.Scan(&choix)
	for choix < 1 || choix > 2 {
		Affichage("Erreur", []string{"Vous devez choisir un choix entre 1 et 2"})
		fmt.Scan(&choix)
	}
	switch choix {
	case 1:
		m.Trade(p)
	}
}
