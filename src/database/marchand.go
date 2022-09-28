package database

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rivo/uniseg"
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
		m.Inv.Liste_boucliers[2].IsUnlocked = true
		m.Inv.Liste_armures_tete[1].IsUnlocked = true
		m.Inv.Liste_armures_torse[1].IsUnlocked = true
		m.Inv.Liste_armures_jambes[1].IsUnlocked = true
		m.Inv.Liste_armures_bras[1].IsUnlocked = true
		m.Nb_trade = 7
	} else if number == 2 {
		m.Init("Gardienne du feu", inv)
		m.Inv.Liste_consommables[1].Quantite = 100
		m.Inv.Liste_consommables[2].Quantite = 100
		m.Inv.Liste_consommables[3].Quantite = 100
		m.Inv.Liste_consommables[4].Quantite = 100
		m.Nb_trade = 4
	} else if number == 3 {
		m.Init("Laurentius", inv)
		m.Inv.Liste_sort[1].IsUnlocked = true
		m.Inv.Liste_sort[2].IsUnlocked = true
		m.Inv.Liste_sort[3].IsUnlocked = true
		m.Inv.Liste_sort[4].IsUnlocked = true
		m.Nb_trade = 4
	} else if number == 4 {
		m.Init("Petrus", inv)
		m.Inv.Liste_consommables[0].Quantite = 100
		m.Inv.Liste_consommables[5].Quantite = 100
		m.Nb_trade = 2
	}
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
						Affichage("Succès", []string{"Vous avez acheté l'arme " + m.Inv.Liste_armes[i].Nom})
						Attendre()
						m.Trade(player)
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
						Affichage("Succès", []string{"Vous avez acheté le bouclier " + m.Inv.Liste_boucliers[i].Nom})
						Attendre()
						m.Trade(player)
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
							m.Inv.Liste_consommables[i].Prix = int(float64(m.Inv.Liste_consommables[i].Prix) * 1.5)
							player.PrendrePot(player.Inv.Liste_consommables[i])
							player.Inv.Liste_consommables[i].Quantite--
							m.Trade(player)
							return
						} else {
							Affichage("Succès", []string{"Vous avez acheté le consommable " + m.Inv.Liste_consommables[i].Nom})
							Attendre()
							m.Trade(player)
							return
						}
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
		if m.Nom == "Petrus" {
			Affichage(m.Nom, []string{"Je suis le marchand " + m.Nom + ".", "Je suis un ancien chevalier de l'Ordre de la Lumière.", "Je suis venu ici pour me reposer et me soigner.", "Je ne peux plus combattre, mais je peux vous aider à vous soigner."})
			Attendre()
			if p.Inv.Liste_consommables[0].Quantite < 3 {
				Affichage(m.Nom, []string{"Je vois que vos fioles sont vides...", "Laissez moi vous remplir !"})
				Attendre()
				p.Inv.Liste_consommables[0].Quantite = 3
				Affichage("Informations", []string{"Vous possédez maintenant 3 fioles d'Estus !"})
			}
		} else {
			Affichage("Marchand", []string{"Heh Heh Heh...",
				"Si vous êtes en manque d'âmes, tuez des monstres !",
				"Plus ils sont gros, plus ils ont d'âmes",
				"Donnez les moi je donnerai des objets de valeur... Heh Heh Heh..."})
			Attendre()
		}
		m.Menu_Marchand(p)
	}
}

func (p *Personnage) Forgeron_amelioration() {
	var Nom string
	longueur := 50
	var list_objets []string
	nb_objets := 0
	//laisser 8 ligne de vide au dessus
	fmt.Print("\n\n\n\n\n\n\n\n")
	//affichage du haut de la boite
	fmt.Print("╒", strings.Repeat("═", longueur), "╕", "\n")
	//affichage du titre
	fmt.Print("│", "Inventaire", strings.Repeat(" ", longueur-10), "│", "\n")
	//affichage de la ligne
	fmt.Print("╞", strings.Repeat("═", longueur), "╡", "\n")
	//affichage des armes
	fmt.Print("│", "Armes Améliorables :", strings.Repeat(" ", longueur-20), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for i := 0; i < len(p.Inv.Liste_armes); i++ {
		if p.Have_item(p.Inv.Liste_armes[i].Lvl) && p.Inv.Liste_armes[i].IsUnlocked {
			Nom = p.Inv.Liste_armes[i].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage des boucliers
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Boucliers Améliorables :", strings.Repeat(" ", longueur-24), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for i := 0; i < len(p.Inv.Liste_boucliers); i++ {
		if p.Have_item(p.Inv.Liste_boucliers[i].Lvl) && p.Inv.Liste_boucliers[i].IsUnlocked {
			Nom = p.Inv.Liste_boucliers[i].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage des casques
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Casques Améliorables :", strings.Repeat(" ", longueur-22), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for i := 0; i < len(p.Inv.Liste_armures_tete); i++ {
		if p.Have_item(p.Inv.Liste_armures_tete[i].Lvl) && p.Inv.Liste_armures_tete[i].IsUnlocked {
			Nom = p.Inv.Liste_armures_tete[i].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage des plastrons
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Plastrons Améliorables :", strings.Repeat(" ", longueur-24), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for i := 0; i < len(p.Inv.Liste_armures_torse); i++ {
		if p.Have_item(p.Inv.Liste_armures_torse[i].Lvl) && p.Inv.Liste_armures_torse[i].IsUnlocked {
			Nom = p.Inv.Liste_armures_torse[i].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage des jambières
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Jambières Améliorables :", strings.Repeat(" ", longueur-24), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for i := 0; i < len(p.Inv.Liste_armures_jambes); i++ {
		if p.Have_item(p.Inv.Liste_armures_jambes[i].Lvl) && p.Inv.Liste_armures_jambes[i].IsUnlocked {
			Nom = p.Inv.Liste_armures_jambes[i].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage des brassards
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Brassards Améliorables :", strings.Repeat(" ", longueur-24), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for i := 0; i < len(p.Inv.Liste_armures_bras); i++ {
		if p.Have_item(p.Inv.Liste_armures_bras[i].Lvl) && p.Inv.Liste_armures_bras[i].IsUnlocked {
			Nom = p.Inv.Liste_armures_bras[i].Nom
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
		item := p.Inv.Get_Item(list_objets[choix-1])
		Affichage("Inventaire", []string{"Vous avez choisi l'objet : " + list_objets[choix-1], "Que voulez vous faire avec cet objet ? ", "1. L'améliorer", "2. Afficher ces statistiques", "3. Retourner au menu précédent"})
		var choix2 = Choix(1, 3)
		switch choix2 {
		case 1:
			switch item := item.(type) {
			case Armes:
				item.Ameliorer_arme(p)
				p.Remplacer_item(item)
				Attendre()
			case Boucliers:
				item.Ameliorer_bouclier(p)
				p.Remplacer_item(item)
				Attendre()
			case Armures:
				item.Ameliorer_Armures(p)
				p.Remplacer_item(item)
				Attendre()
			}
		case 2:
			switch item := item.(type) {
			case Armes:
				item.Affichage()
				Attendre()
				p.Forgeron_amelioration()
			case Boucliers:
				item.Affichage()
				Attendre()
				p.Forgeron_amelioration()
			case Armures:
				item.Affichage()
				Attendre()
				p.Forgeron_amelioration()
			}
		case 3:
			p.Forgeron_amelioration()
		}
	}
}

func (p *Personnage) Have_item(number int) bool {
	if number > 3 {
		return false
	}
	switch number {
	case 1:
		if p.Inv.Liste_items["éclat de titanite"] >= 6 && p.Ames >= 100 {
			return true
		} else {
			return false
		}
	case 2:
		if p.Inv.Liste_items["grand éclat de titanite"] >= 3 && p.Inv.Liste_items["éclat de titanite"] >= 3 && p.Ames >= 500 {
			return true
		} else {
			return false
		}
	case 3:
		if p.Inv.Liste_items["grand éclat de titanite"] >= 2 && p.Inv.Liste_items["éclat de titanite"] >= 2 && p.Inv.Liste_items["tablette éclat de titanite"] >= 2 && p.Ames >= 2000 {
			return true
		} else {
			return false
		}
	}
	return false
}
