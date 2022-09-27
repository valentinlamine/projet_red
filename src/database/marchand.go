package database

import (
	"fmt"
	"strconv"
)

type Marchand struct {
	Nom         string
	Inv         Inventaire
	Nb_trade    int
	First_trade bool
}

func (m *Marchand) Init(Nom string, inv Inventaire) {
	m.Nom = Nom
	m.Nb_trade = 0
	m.Inv = inv
	m.First_trade = true
}

func (m *Marchand) InitMarchand(number int) {
	//Marchand 1
	var inv Inventaire
	inv.Init()
	if number == 1 {
		m.Init("Marchand mort vivant", inv)
		m.Inv.Liste_armes[4].IsUnlocked = true
		m.Inv.Liste_consommables[5].Quantite = 100
		m.Inv.Liste_boucliers[2].IsUnlocked = true
		m.Inv.Liste_armures_tete[1].IsUnlocked = true
		m.Inv.Liste_armures_torse[1].IsUnlocked = true
		m.Inv.Liste_armures_jambes[1].IsUnlocked = true
		m.Inv.Liste_armures_bras[1].IsUnlocked = true
		m.Inv.Liste_sort[1].IsUnlocked = true
		m.Nombre_Trade()
	} else if number == 2 {
		m.Init("Laurentius", inv)
		//TODO
	} else if number == 3 {
		m.Init("Gardienne du feu", inv)
		m.Inv.Liste_consommables[1].Quantite = 100
		m.Inv.Liste_consommables[2].Quantite = 100
		m.Inv.Liste_consommables[3].Quantite = 100
		m.Inv.Liste_consommables[4].Quantite = 100
	}
}

func (m *Marchand) Nombre_Trade() {
	var nombre int
	for i := 0; i < len(m.Inv.Liste_armes); i++ {
		if m.Inv.Liste_armes[i].IsUnlocked {
			nombre++
		}
	}
	for i := 0; i < len(m.Inv.Liste_boucliers); i++ {
		if m.Inv.Liste_boucliers[i].IsUnlocked {
			nombre++
		}
	}
	for i := 0; i < len(m.Inv.Liste_consommables); i++ {
		if m.Inv.Liste_consommables[i].Quantite > 0 {
			nombre++
		}
	}
	for i := 0; i < len(m.Inv.Liste_armures_tete); i++ {
		if m.Inv.Liste_armures_tete[i].IsUnlocked {
			nombre++
		}
	}
	for i := 0; i < len(m.Inv.Liste_armures_torse); i++ {
		if m.Inv.Liste_armures_torse[i].IsUnlocked {
			nombre++
		}
	}
	for i := 0; i < len(m.Inv.Liste_armures_jambes); i++ {
		if m.Inv.Liste_armures_jambes[i].IsUnlocked {
			nombre++
		}
	}
	for i := 0; i < len(m.Inv.Liste_armures_bras); i++ {
		if m.Inv.Liste_armures_bras[i].IsUnlocked {
			nombre++
		}
	}
	for i := 0; i < len(m.Inv.Liste_sort); i++ {
		if m.Inv.Liste_sort[i].IsUnlocked {
			nombre++
		}
	}

	m.Nb_trade = nombre
}

func (m *Marchand) Affichage_Trade() {
	var liste []string
	liste = append(liste, "Que voulez-vous acheter ?")
	for i := 0; i < len(m.Inv.Liste_armes); i++ {
		if m.Inv.Liste_armes[i].IsUnlocked {
			liste = append(liste, strconv.Itoa(len(liste))+". "+m.Inv.Liste_armes[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_armes[i].Prix)+" âmes")
		}
	}
	for i := 0; i < len(m.Inv.Liste_boucliers); i++ {
		if m.Inv.Liste_boucliers[i].IsUnlocked {
			liste = append(liste, strconv.Itoa(len(liste))+". "+m.Inv.Liste_boucliers[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_boucliers[i].Prix)+" âmes")
		}
	}
	for i := 0; i < len(m.Inv.Liste_consommables); i++ {
		if m.Inv.Liste_consommables[i].Quantite > 0 {
			liste = append(liste, strconv.Itoa(len(liste))+". "+m.Inv.Liste_consommables[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_consommables[i].Prix)+" âmes")
		}
	}
	for i := 0; i < len(m.Inv.Liste_armures_tete); i++ {
		if m.Inv.Liste_armures_tete[i].IsUnlocked {
			liste = append(liste, strconv.Itoa(len(liste))+". "+m.Inv.Liste_armures_tete[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_armures_tete[i].Prix)+" âmes")
		}
	}
	for i := 0; i < len(m.Inv.Liste_armures_torse); i++ {
		if m.Inv.Liste_armures_torse[i].IsUnlocked {
			liste = append(liste, strconv.Itoa(len(liste))+". "+m.Inv.Liste_armures_torse[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_armures_torse[i].Prix)+" âmes")
		}
	}
	for i := 0; i < len(m.Inv.Liste_armures_jambes); i++ {
		if m.Inv.Liste_armures_jambes[i].IsUnlocked {
			liste = append(liste, strconv.Itoa(len(liste))+". "+m.Inv.Liste_armures_jambes[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_armures_jambes[i].Prix)+" âmes")
		}
	}
	for i := 0; i < len(m.Inv.Liste_armures_bras); i++ {
		if m.Inv.Liste_armures_bras[i].IsUnlocked {
			liste = append(liste, strconv.Itoa(len(liste))+". "+m.Inv.Liste_armures_bras[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_armures_bras[i].Prix)+" âmes")
		}
	}
	for i := 0; i < len(m.Inv.Liste_sort); i++ {
		if m.Inv.Liste_sort[i].IsUnlocked {
			liste = append(liste, strconv.Itoa(len(liste))+". "+m.Inv.Liste_sort[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_sort[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "", strconv.Itoa(len(liste))+". Quitter le menu d'achat")
	Affichage(m.Nom, liste)
}

func (m *Marchand) Trade(player *Personnage) {
	if m.First_trade {
		m.IsFirst_trade(player)
	}
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
		if m.Inv.Liste_armes[i].IsUnlocked {
			choix--
			if choix == 0 {
				if WantTrade(player, player.Inv.Liste_armes[i]) {
					if IsTradeable(player, "Armes", i) {
						player.Inv.Liste_armes[i].IsUnlocked = true
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
		if m.Inv.Liste_boucliers[i].IsUnlocked {
			choix--
			if choix == 0 {
				if WantTrade(player, player.Inv.Liste_boucliers[i]) {
					if IsTradeable(player, "Boucliers", i) {
						player.Inv.Liste_boucliers[i].IsUnlocked = true
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
		if m.Inv.Liste_consommables[i].Quantite > 0 {
			choix--
			if choix == 0 {
				if WantTrade(player, player.Inv.Liste_consommables[i]) {
					if IsTradeable(player, "Consommables", i) {
						player.Inv.Liste_consommables[i].Quantite++
						player.Ames -= m.Inv.Liste_consommables[i].Prix
						if i > 0 && i < 5 {
							player.Inv.Liste_consommables[i].Prix = int(float64(player.Inv.Liste_consommables[i].Prix) * 1.5)
						}
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
	for i := 0; i < len(m.Inv.Liste_armures_tete); i++ {
		if m.Inv.Liste_armures_tete[i].IsUnlocked {
			choix--
			if choix == 0 {
				if WantTrade(player, player.Inv.Liste_armures_tete[i]) {
					if IsTradeable(player, "Armures_tete", i) {
						player.Inv.Liste_armures_tete[i].IsUnlocked = true
						player.Ames -= m.Inv.Liste_armures_tete[i].Prix
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
	for i := 0; i < len(m.Inv.Liste_armures_torse); i++ {
		if m.Inv.Liste_armures_torse[i].IsUnlocked {
			choix--
			if choix == 0 {
				if WantTrade(player, player.Inv.Liste_armures_torse[i]) {
					if IsTradeable(player, "Armures_torse", i) {
						player.Inv.Liste_armures_torse[i].IsUnlocked = true
						player.Ames -= m.Inv.Liste_armures_torse[i].Prix
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
	for i := 0; i < len(m.Inv.Liste_armures_jambes); i++ {
		if m.Inv.Liste_armures_jambes[i].IsUnlocked {
			choix--
			if choix == 0 {
				if WantTrade(player, player.Inv.Liste_armures_jambes[i]) {
					if IsTradeable(player, "Armures_jambes", i) {
						player.Inv.Liste_armures_jambes[i].IsUnlocked = true
						player.Ames -= m.Inv.Liste_armures_jambes[i].Prix
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
	for i := 0; i < len(m.Inv.Liste_armures_bras); i++ {
		if m.Inv.Liste_armures_bras[i].IsUnlocked {
			choix--
			if choix == 0 {
				if WantTrade(player, player.Inv.Liste_armures_bras[i]) {
					if IsTradeable(player, "Armures_bras", i) {
						player.Inv.Liste_armures_bras[i].IsUnlocked = true
						player.Ames -= m.Inv.Liste_armures_bras[i].Prix
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
	for i := 0; i < len(m.Inv.Liste_sort); i++ {
		if m.Inv.Liste_sort[i].IsUnlocked {
			choix--
			if choix == 0 {
				if WantTrade(player, player.Inv.Liste_sort[i]) {
					if IsTradeable(player, "Sort", i) {
						player.Inv.Liste_sort[i].IsUnlocked = true
						player.Ames -= m.Inv.Liste_sort[i].Prix
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
		Affichage("Marchand", []string{"Que voulez-vous faire ?", "1. Acheter l'item " + item.Nom, "2. Afficher les caractéristiques de l'objet", "3. Revenir en arrière"})
	case Boucliers:
		Affichage("Marchand", []string{"Que voulez-vous faire ?", "1. Acheter l'item " + item.Nom, "2. Afficher les caractéristiques de l'objet", "3. Revenir en arrière"})
	case Consommable:
		Affichage("Marchand", []string{"Que voulez-vous faire ?", "1. Acheter l'item " + item.Nom, "2. Afficher les caractéristiques de l'objet", "3. Revenir en arrière"})
	case Armures:
		Affichage("Marchand", []string{"Que voulez-vous faire ?", "1. Acheter l'item " + item.Nom, "2. Afficher les caractéristiques de l'objet", "3. Revenir en arrière"})
	case Sort:
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
		case Armures:
			item.Affichage()
			Attendre()
		case Sort:
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
	} else if item_type == "Armures_tete" {
		if IsBuyable(player.Ames, player.Inv.Liste_armures_tete[index].Prix) && !player.Inv.Liste_armures_tete[index].IsUnlocked {
			tradeable = true
		} else if player.Inv.Liste_armures_tete[index].IsUnlocked {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"})
			Attendre()
		}
	} else if item_type == "Armures_torse" {
		if IsBuyable(player.Ames, player.Inv.Liste_armures_torse[index].Prix) && !player.Inv.Liste_armures_torse[index].IsUnlocked {
			tradeable = true
		} else if player.Inv.Liste_armures_torse[index].IsUnlocked {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"})
			Attendre()
		}
	} else if item_type == "Armures_jambes" {
		if IsBuyable(player.Ames, player.Inv.Liste_armures_jambes[index].Prix) && !player.Inv.Liste_armures_jambes[index].IsUnlocked {
			tradeable = true
		} else if player.Inv.Liste_armures_jambes[index].IsUnlocked {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"})
			Attendre()
		}
	} else if item_type == "Armures_bras" {
		if IsBuyable(player.Ames, player.Inv.Liste_armures_bras[index].Prix) && !player.Inv.Liste_armures_bras[index].IsUnlocked {
			tradeable = true
		} else if player.Inv.Liste_armures_bras[index].IsUnlocked {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"})
			Attendre()
		}
	} else if item_type == "Sort" {
		if IsBuyable(player.Ames, player.Inv.Liste_sort[index].Prix) && !player.Inv.Liste_sort[index].IsUnlocked {
			tradeable = true
		} else if player.Inv.Liste_sort[index].IsUnlocked {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"})
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

func (m *Marchand) IsFirst_trade(player *Personnage) {
	Affichage(m.Nom, []string{"Bienvenue dans mon magasin !", "Je suis le marchand " + m.Nom + ".", "Comme vous avez l'air d'être nouveau, laissez moi vous faire un petit cadeau !"})
	Attendre()
	player.Inv.Liste_consommables[0].Quantite += 1
	m.First_trade = false
	Affichage("Informations", []string{"Félicitations !!!", "Vous avez obtenu une Fiole d'Estus gratuitement !"})
	Attendre()
}

func (m *Marchand) Menu_Marchand(p *Personnage) {
	NouvelAffichage("Menu du marchand", []string{"Que voulez-vous faire ?", "1. Acheter un objet", "2. Parler", "3. Quitter le menu du marchand"})
	var choix = Choix(1, 3)
	if choix == 1 {
		m.Trade(p)
	} else if choix == 2 {
		Affichage("Marchand", []string{"Heh Heh Heh...",
			"Si vous êtes en manque d'âmes, tuez des monstres !",
			"Plus ils sont gros, plus ils ont d'âmes",
			"Donnez les moi je donnerai des objets de valeur... Heh Heh Heh..."})
		Attendre()
		m.Menu_Marchand(p)
	}
}
