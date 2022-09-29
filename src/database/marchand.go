package database

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
		//Armes
		m.Inv.Liste_armes[0].IsUnlocked = true
		m.Inv.Liste_armes[1].IsUnlocked = true
		m.Inv.Liste_armes[2].IsUnlocked = true
		m.Inv.Liste_armes[3].IsUnlocked = true
		m.Inv.Liste_armes[4].IsUnlocked = true
		m.Inv.Liste_armes[5].IsUnlocked = true
		//Boucliers
		m.Inv.Liste_boucliers[0].IsUnlocked = true
		m.Inv.Liste_boucliers[1].IsUnlocked = true
		m.Inv.Liste_boucliers[2].IsUnlocked = true
		m.Inv.Liste_boucliers[3].IsUnlocked = true
		m.Inv.Liste_boucliers[4].IsUnlocked = true
		m.Inv.Liste_boucliers[5].IsUnlocked = true
		//première armure
		m.Inv.Liste_armures_tete[1].IsUnlocked = true
		m.Inv.Liste_armures_torse[1].IsUnlocked = true
		m.Inv.Liste_armures_jambes[1].IsUnlocked = true
		m.Inv.Liste_armures_bras[1].IsUnlocked = true
		//deuxième armure
		m.Inv.Liste_armures_tete[2].IsUnlocked = true
		m.Inv.Liste_armures_torse[2].IsUnlocked = true
		m.Inv.Liste_armures_jambes[2].IsUnlocked = true
		m.Inv.Liste_armures_bras[2].IsUnlocked = true
		//troisième armure
		m.Inv.Liste_armures_tete[3].IsUnlocked = true
		m.Inv.Liste_armures_torse[3].IsUnlocked = true
		m.Inv.Liste_armures_jambes[3].IsUnlocked = true
		m.Inv.Liste_armures_bras[3].IsUnlocked = true
		//quatrième armure
		m.Inv.Liste_armures_tete[4].IsUnlocked = true
		m.Inv.Liste_armures_torse[4].IsUnlocked = true
		m.Inv.Liste_armures_jambes[4].IsUnlocked = true
		m.Inv.Liste_armures_bras[4].IsUnlocked = true
		//cinquième armure
		m.Inv.Liste_armures_tete[5].IsUnlocked = true
		m.Inv.Liste_armures_torse[5].IsUnlocked = true
		m.Inv.Liste_armures_jambes[5].IsUnlocked = true
		m.Inv.Liste_armures_bras[5].IsUnlocked = true
		m.Nb_trade = 22
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
		m.Inv.Liste_sort[5].IsUnlocked = true
		m.Nb_trade = 5
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
	liste = append(liste, "")
	liste = append(liste, "Armes :")
	for i := 0; i < len(m.Inv.Liste_armes); i++ {
		if m.Inv.Liste_armes[i].IsUnlocked {
			liste = append(liste, strconv.Itoa(len(liste)-2)+". "+m.Inv.Liste_armes[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_armes[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "")
	liste = append(liste, "Boucliers :")
	for i := 0; i < len(m.Inv.Liste_boucliers); i++ {
		if m.Inv.Liste_boucliers[i].IsUnlocked {
			liste = append(liste, strconv.Itoa(len(liste)-4)+". "+m.Inv.Liste_boucliers[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_boucliers[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "")
	if m.Nom == "Gardienne du feu" {
		liste = append(liste, "Livre de niveau :")
	} else {
		liste = append(liste, "Consommables :")
	}
	for i := 0; i < len(m.Inv.Liste_consommables); i++ {
		if m.Inv.Liste_consommables[i].Quantite > 0 {
			liste = append(liste, strconv.Itoa(len(liste)-6)+". "+m.Inv.Liste_consommables[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_consommables[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "")
	liste = append(liste, "Casque :")
	for i := 0; i < len(m.Inv.Liste_armures_tete); i++ {
		if m.Inv.Liste_armures_tete[i].IsUnlocked {
			liste = append(liste, strconv.Itoa(len(liste)-8)+". "+m.Inv.Liste_armures_tete[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_armures_tete[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "")
	liste = append(liste, "Plastrons :")
	for i := 0; i < len(m.Inv.Liste_armures_torse); i++ {
		if m.Inv.Liste_armures_torse[i].IsUnlocked {
			liste = append(liste, strconv.Itoa(len(liste)-10)+". "+m.Inv.Liste_armures_torse[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_armures_torse[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "")
	liste = append(liste, "Jambières :")
	for i := 0; i < len(m.Inv.Liste_armures_jambes); i++ {
		if m.Inv.Liste_armures_jambes[i].IsUnlocked {
			liste = append(liste, strconv.Itoa(len(liste)-12)+". "+m.Inv.Liste_armures_jambes[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_armures_jambes[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "")
	liste = append(liste, "Brassards :")
	for i := 0; i < len(m.Inv.Liste_armures_bras); i++ {
		if m.Inv.Liste_armures_bras[i].IsUnlocked {
			liste = append(liste, strconv.Itoa(len(liste)-14)+". "+m.Inv.Liste_armures_bras[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_armures_bras[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "")
	liste = append(liste, "Sorts :")
	for i := 0; i < len(m.Inv.Liste_sort); i++ {
		if m.Inv.Liste_sort[i].IsUnlocked {
			liste = append(liste, strconv.Itoa(len(liste)-16)+". "+m.Inv.Liste_sort[i].Nom+" | "+strconv.Itoa(m.Inv.Liste_sort[i].Prix)+" âmes")
		}
	}
	liste = append(liste, "", strconv.Itoa(len(liste)-16)+". Quitter le menu d'achat")
	Affichage(m.Nom, liste, true, false)
}

func (m *Marchand) Trade(player *Personnage) {
	if m.First_trade {
		m.IsFirst_trade(player)
	}
	m.Affichage_Trade()
	var choix int
	fmt.Scan(&choix)
	for choix < 1 || choix > m.Nb_trade+1 {
		Affichage("Erreur", []string{"Vous devez choisir un choix entre 1 et " + strconv.Itoa(m.Nb_trade+1)}, false, false)
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
						Affichage("Succès", []string{"Vous avez acheté l'arme " + m.Inv.Liste_armes[i].Nom}, true, true)
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
						Affichage("Succès", []string{"Vous avez acheté le bouclier " + m.Inv.Liste_boucliers[i].Nom}, true, true)
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
							m.Inv.Liste_consommables[1].Prix += (player.Niveau * player.Niveau) / 2
							m.Inv.Liste_consommables[2].Prix += (player.Niveau * player.Niveau) / 2
							m.Inv.Liste_consommables[3].Prix += (player.Niveau * player.Niveau) / 2
							m.Inv.Liste_consommables[4].Prix += (player.Niveau * player.Niveau) / 2
							player.PrendrePot(player.Inv.Liste_consommables[i])
							player.Inv.Liste_consommables[i].Quantite--
							m.Trade(player)
							return
						} else {
							Affichage("Succès", []string{"Vous avez acheté le consommable " + m.Inv.Liste_consommables[i].Nom}, true, true)
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
						Affichage("Succès", []string{"Vous avez acheté un objet"}, true, true)
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
						Affichage("Succès", []string{"Vous avez acheté un objet"}, true, true)
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
						Affichage("Succès", []string{"Vous avez acheté un objet"}, true, true)
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
						Affichage("Succès", []string{"Vous avez acheté un objet"}, true, true)
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
						Affichage("Succès", []string{"Vous avez acheté un objet"}, true, true)
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
	var choix int
	fmt.Scan(&choix)
	for choix < 1 || choix > 3 {
		Affichage("Erreur", []string{"Vous devez choisir un choix entre 1 et 3"}, false, false)
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
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"}, true, true)
		}
	} else if item_type == "Boucliers" {
		if IsBuyable(player.Ames, player.Inv.Liste_boucliers[index].Prix) && !player.Inv.Liste_boucliers[index].IsUnlocked {
			tradeable = true
		} else if player.Inv.Liste_boucliers[index].IsUnlocked {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"}, true, true)
		}
	} else if item_type == "Consommables" {
		if IsBuyable(player.Ames, player.Inv.Liste_consommables[index].Prix) && player.Inv.Liste_consommables[index].Quantite <= 100 {
			tradeable = true
		} else if player.Inv.Liste_consommables[index].Quantite > 100 {
			Affichage("Marchand", []string{"Vous possédez déjà la quantité maximale de cet objet"}, true, true)
		}
	} else if item_type == "Armures_tete" {
		if IsBuyable(player.Ames, player.Inv.Liste_armures_tete[index].Prix) && !player.Inv.Liste_armures_tete[index].IsUnlocked {
			tradeable = true
		} else if player.Inv.Liste_armures_tete[index].IsUnlocked {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"}, true, true)
		}
	} else if item_type == "Armures_torse" {
		if IsBuyable(player.Ames, player.Inv.Liste_armures_torse[index].Prix) && !player.Inv.Liste_armures_torse[index].IsUnlocked {
			tradeable = true
		} else if player.Inv.Liste_armures_torse[index].IsUnlocked {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"}, true, true)
		}
	} else if item_type == "Armures_jambes" {
		if IsBuyable(player.Ames, player.Inv.Liste_armures_jambes[index].Prix) && !player.Inv.Liste_armures_jambes[index].IsUnlocked {
			tradeable = true
		} else if player.Inv.Liste_armures_jambes[index].IsUnlocked {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"}, true, true)
		}
	} else if item_type == "Armures_bras" {
		if IsBuyable(player.Ames, player.Inv.Liste_armures_bras[index].Prix) && !player.Inv.Liste_armures_bras[index].IsUnlocked {
			tradeable = true
		} else if player.Inv.Liste_armures_bras[index].IsUnlocked {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"}, true, true)
		}
	} else if item_type == "Sort" {
		if IsBuyable(player.Ames, player.Inv.Liste_sort[index].Prix) && !player.Inv.Liste_sort[index].IsUnlocked {
			tradeable = true
		} else if player.Inv.Liste_sort[index].IsUnlocked {
			Affichage("Marchand", []string{"Vous possédez déjà cet objet"}, true, true)
		}
	}
	return tradeable
}

func IsBuyable(Ames, Prix int) bool {
	if Ames <= Prix {
		Affichage("Erreur", []string{"Vous n'avez pas assez d'âmes pour acheter cet objet"}, true, true)
		return false
	}
	return true
}

func (m *Marchand) IsFirst_trade(player *Personnage) {
	Affichage(m.Nom, []string{"Bienvenue dans mon magasin !", "Je suis le marchand " + m.Nom + ".", "Comme vous avez l'air d'être nouveau, laissez moi vous faire un petit cadeau !"}, true, true)
	player.Inv.Liste_consommables[0].Quantite += 1
	m.First_trade = false
	Affichage("Informations", []string{"Félicitations !!!", "Vous avez obtenu une Fiole d'Estus gratuitement !"}, true, true)
}

func (m *Marchand) Menu_Marchand(p *Personnage) {
	Affichage("Menu du marchand", []string{"Que voulez-vous faire ?", "1. Acheter un objet", "2. Parler", "3. Quitter le menu du marchand"}, true, false)
	var choix = Choix(1, 3)
	if choix == 1 {
		m.Trade(p)
	} else if choix == 2 {
		if m.Nom == "Petrus" {
			Affichage(m.Nom, []string{"Je suis le marchand " + m.Nom + ".", "Je suis un ancien chevalier de l'Ordre de la Lumière.", "Je suis venu ici pour me reposer et me soigner.", "Je ne peux plus combattre, mais je peux vous aider à vous soigner."}, true, true)
			if p.Inv.Liste_consommables[0].Quantite < 3 {
				Affichage(m.Nom, []string{"Je vois que vos fioles sont vides...", "Laissez moi vous remplir !"}, true, true)
				p.Inv.Liste_consommables[0].Quantite = 3
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

func (p *Personnage) Forgeron_amelioration() {
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
		Affichage("Inventaire", []string{"Vous avez choisi l'objet : " + list_objets[choix-1], "Que voulez vous faire avec cet objet ? ", "1. L'améliorer", "2. Afficher ces statistiques", "3. Retourner au menu précédent"}, true, false)
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
		if p.Inv.Liste_items["grand éclat de titanite"] >= 2 && p.Inv.Liste_items["éclat de titanite"] >= 2 && p.Inv.Liste_items["tablette de titanite"] >= 2 && p.Ames >= 2000 {
			return true
		} else {
			return false
		}
	}
	return false
}
