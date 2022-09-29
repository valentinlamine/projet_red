package database

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rivo/uniseg"
)

func Affichage(titre string, liste []string, espace bool, attendre bool) {
	longueur := 0
	for _, s := range liste {
		if uniseg.GraphemeClusterCount(s) > longueur {
			longueur = uniseg.GraphemeClusterCount(s)
		}
	}
	if uniseg.GraphemeClusterCount(titre) > longueur {
		longueur = uniseg.GraphemeClusterCount(titre)
	}
	if espace {
		fmt.Print("\n\n\n\n\n\n\n\n\n\n")
		fmt.Print("\n\n\n\n\n\n\n\n\n\n")
		fmt.Print("\n\n\n\n\n\n\n\n\n\n")
		fmt.Print("\n\n\n\n\n\n\n\n\n\n")
	}
	//affichage du haut de la boite
	fmt.Print("╒", strings.Repeat("═", longueur), "╕", "\n")
	//affichage du titre
	fmt.Print("│", titre, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(titre)), "│", "\n")
	//affichage de la ligne
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	//affichage des lignes de texte
	for _, s := range liste {
		fmt.Print("│", s, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(s)), "│", "\n")
	}
	//affichage du bas de la boite
	fmt.Print("╘", strings.Repeat("═", longueur), "╛", "\n")
	if attendre {
		Attendre()
	}
}

func (p *Personnage) Affichage_Personnage() {
	Affichage("Personnage", []string{"Nom : " + p.Nom, "Classe : " + p.Classe, "Niveau : " + strconv.Itoa(p.Niveau), "Ames : " + strconv.Itoa(p.Ames), "PV : " + strconv.Itoa(p.VieAct) + "/" + strconv.Itoa(p.VieMax), "Poids : " + strconv.Itoa(p.PoidsAct) + "/" + strconv.Itoa(p.PoidsMax), "Vitalité: " + strconv.Itoa(p.Vitalite), "Force : " + strconv.Itoa(p.Force), "Dextérité : " + strconv.Itoa(p.Dexterite), "Intelligence : " + strconv.Itoa(p.Intelligence), "Position : " + p.Position.Val["nom"], "ManaAct : " + strconv.Itoa(p.ManaAct) + "/" + strconv.Itoa(p.ManaMax), "Dégâts : " + strconv.Itoa(p.Degats)}, true, false)
	fmt.Println("Appuyez sur 1 pour voir votre inventaire équipé")
	fmt.Println("Appuyez sur 0 pour revenir au menu principal")
	var choix = Choix(0, 1)
	switch choix {
	case 1:
		p.Affichage_Inventaire_Equipe()
		Attendre()
	case 0:
		return
	}
}

func (p *Personnage) Affichage_Inventaire() {
	var Nom string
	var list_objets []string
	var nb_objets int
	longueur := 50
	//laisser 8 ligne de vide au dessus
	fmt.Print("\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n")
	//affichage du haut de la boite
	fmt.Print("╒", strings.Repeat("═", longueur), "╕", "\n")
	//affichage du titre
	fmt.Print("│", "Inventaire", strings.Repeat(" ", longueur-10), "│", "\n")
	//affichage de la ligne
	fmt.Print("╞", strings.Repeat("═", longueur), "╡", "\n")
	//affichage des armes débloquées
	fmt.Print("│", "Armes débloquées :", strings.Repeat(" ", longueur-18), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_Armes); index++ {
		if p.Inv.Liste_Armes[index].EstDebloque {
			Nom = p.Inv.Liste_Armes[index].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage des boucliers débloqués
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Boucliers débloqués :", strings.Repeat(" ", longueur-21), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_Boucliers); index++ {
		if p.Inv.Liste_Boucliers[index].EstDebloque {
			Nom = p.Inv.Liste_Boucliers[index].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage des consommables débloqués
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Consommables débloqués :", strings.Repeat(" ", longueur-24), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_Consommables); index++ {
		if p.Inv.Liste_Consommables[index].Quantite > 0 {
			Nom = p.Inv.Liste_Consommables[index].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, " ("+strconv.Itoa(p.Inv.Liste_Consommables[index].Quantite)+")", strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-(8+len(strconv.Itoa(p.Inv.Liste_Consommables[index].Quantite)))), "│", "\n")
		}
	}
	//affichage des armures de tête débloquées
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Casques débloquées :", strings.Repeat(" ", longueur-20), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_Armures_Tete); index++ {
		if p.Inv.Liste_Armures_Tete[index].EstDebloque {
			Nom = p.Inv.Liste_Armures_Tete[index].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage des armures de torse débloquées
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Plastrons débloquées :", strings.Repeat(" ", longueur-22), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_Armures_Torse); index++ {
		if p.Inv.Liste_Armures_Torse[index].EstDebloque {
			Nom = p.Inv.Liste_Armures_Torse[index].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage des armures de jambes débloquées
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Jambières débloquées :", strings.Repeat(" ", longueur-22), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_Armures_Jambes); index++ {
		if p.Inv.Liste_Armures_Jambes[index].EstDebloque {
			Nom = p.Inv.Liste_Armures_Jambes[index].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage des armures de bras débloquées
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Brassards débloquées :", strings.Repeat(" ", longueur-22), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_Armures_Bras); index++ {
		if p.Inv.Liste_Armures_Bras[index].EstDebloque {
			Nom = p.Inv.Liste_Armures_Bras[index].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//sorts débloqués
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Sorts débloqués :", strings.Repeat(" ", longueur-17), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_Sorts); index++ {
		if p.Inv.Liste_Sorts[index].EstDebloque {
			Nom = p.Inv.Liste_Sorts[index].Nom
			nb_objets++
			list_objets = append(list_objets, Nom)
			fmt.Print("│  ", nb_objets, ". ", Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(Nom)-5), "│", "\n")
		}
	}
	//affichage des items débloqués
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Items débloqués :", strings.Repeat(" ", longueur-17), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for i := range p.Inv.Liste_Items {
		if p.Inv.Liste_Items[i] > 0 {
			nb_objets++
			fmt.Print("│  ", p.Inv.Liste_Items[i], "x ", i, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(i)-(4+len(strconv.Itoa(p.Inv.Liste_Items[i])))), "│", "\n")
		}
	}
	//affichage du bas de la boite
	fmt.Print("╘", strings.Repeat("═", longueur), "╛", "\n")

	//choix de l'objet à afficher
	fmt.Print("Choisissez un objet à afficher (0 pour quitter) : ")
	var choix = Choix(0, len(list_objets))
	if choix != 0 {
		//affichage de l'objet choisi en utilisant la fonction Affichage
		item := p.Inv.Recuperer_Item(list_objets[choix-1])
		var choix2 int
		switch item := item.(type) {
		case Consommable:
			Affichage("Inventaire", []string{"Vous avez choisi l'objet : " + item.Nom, "Que voulez vous faire avec cet objet ? ", "1. Le consommer", "2. Afficher ces statistiques", "3. Retourner au menu précédent"}, false, false)
			choix2 = Choix(1, 3)
		case Sort:
			Affichage("Inventaire", []string{"Vous avez choisi le sort : " + item.Nom, "Que voulez vous faire avec ce sort ? ", "2. Afficher ces statistiques", "3. Retourner au menu précédent"}, false, false)
			choix2 = Choix(2, 3)
		default:
			Affichage("Inventaire", []string{"Vous avez choisi l'objet : " + list_objets[choix-1], "Que voulez vous faire avec cet objet ? ", "1. L'équiper", "2. Afficher ces statistiques", "3. Retourner au menu précédent"}, false, false)
			choix2 = Choix(1, 3)
		}

		if choix2 == 1 {
			p.Equiper(item, true) //si consommable la fonction Equiper() va le consommer
		} else if choix2 == 2 {
			switch item := item.(type) {
			case Armes:
				item.Affichage()
				p.Affichage_Inventaire()
			case Boucliers:
				item.Affichage()
				p.Affichage_Inventaire()
			case Consommable:
				item.Affichage()
				p.Affichage_Inventaire()
			case Armures:
				item.Affichage()
				p.Affichage_Inventaire()
			case Sort:
				item.Affichage()
				p.Affichage_Inventaire()
			default:
				print("Erreur")
			}
		} else if choix2 == 3 {
			p.Affichage_Inventaire()
		}
	}
}

func (p *Personnage) Affichage_Inventaire_Equipe() {
	longueur := 50
	fmt.Print("\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n")
	//affichage du haut de la boite
	fmt.Print("╒", strings.Repeat("═", longueur), "╕", "\n")
	fmt.Print("│", "Inventaire équipé :", strings.Repeat(" ", longueur-19), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	//affichage de l'arme équipée
	fmt.Print("│", "Arme équipée :", strings.Repeat(" ", longueur-14), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│ ● ", p.EquipementArmes.Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(p.EquipementArmes.Nom)-3), "│", "\n")
	//affichage du bouclier équipé
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Bouclier équipé :", strings.Repeat(" ", longueur-17), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│ ● ", p.EquipementBoucliers.Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(p.EquipementBoucliers.Nom)-3), "│", "\n")
	//affichage de l'armure de tête équipée
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Casque équipée :", strings.Repeat(" ", longueur-16), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│ ● ", p.EquipementArmures["tete"].Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(p.EquipementArmures["tete"].Nom)-3), "│", "\n")
	//affichage de l'armure de torse équipée
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Plastron équipée :", strings.Repeat(" ", longueur-18), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│ ● ", p.EquipementArmures["torse"].Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(p.EquipementArmures["torse"].Nom)-3), "│", "\n")
	//affichage de l'armure de jambes équipée
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Jambières équipée :", strings.Repeat(" ", longueur-19), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│ ● ", p.EquipementArmures["jambes"].Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(p.EquipementArmures["jambes"].Nom)-3), "│", "\n")
	//affichage de l'armure de bras équipée
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Brassards équipée :", strings.Repeat(" ", longueur-19), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│ ● ", p.EquipementArmures["bras"].Nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(p.EquipementArmures["bras"].Nom)-3), "│", "\n")
	//affichage du bas de la boite
	fmt.Print("└", strings.Repeat("─", longueur), "┘", "\n")
}

func (a *Armes) Affichage() {
	Affichage("Arme", []string{"Nom : " + a.Nom, "Niveau : " + strconv.Itoa(a.Niv), "Stat Min Force : " + strconv.Itoa(a.NivMinFor), "Stat Min Dextérité : " + strconv.Itoa(a.NivMinDex), "Stat Min Intelligence : " + strconv.Itoa(a.NivMinInt), "Dégâts : " + strconv.Itoa(a.Degats), "Poids : " + strconv.Itoa(a.Poids), "Équipé : " + strconv.FormatBool(a.IsEquiped)}, true, true)
}

func (b *Boucliers) Affichage() {
	Affichage("Bouclier", []string{"Nom : " + b.Nom, "Niveau : " + strconv.Itoa(b.Niv), "Stat Min Force : " + strconv.Itoa(b.NivMinFor), "Stat Min Dextérité : " + strconv.Itoa(b.NivMinDex), "Stat Min Intelligence : " + strconv.Itoa(b.NivMinInt), "PV Bonus : " + strconv.Itoa(b.VieBonus), "Poids : " + strconv.Itoa(b.Poids), "Équipé : " + strconv.FormatBool(b.IsEquiped)}, true, true)
}

func (c *Consommable) Affichage() {
	Affichage("Consommable", []string{"Nom : " + c.Nom, "Prix : " + strconv.Itoa(c.Prix), "Quantité : " + strconv.Itoa(c.Quantite), "PV Bonus : " + strconv.Itoa(c.VieBonus), "Poids Bonus : " + strconv.Itoa(c.PoidsBonus), "Dégats bonus : " + strconv.Itoa(c.DegatsBonus), "Mana bonus : " + strconv.Itoa(c.ManaBonus)}, true, true)
}

func (a *Armures) Affichage() {
	Affichage("Armure", []string{"Nom : " + a.Nom, "Niveau : " + strconv.Itoa(a.Niv), "Stat Min Force : " + strconv.Itoa(a.NivMinFor), "Stat Min Dextérité : " + strconv.Itoa(a.NivMinDex), "Stat Min Intelligence : " + strconv.Itoa(a.NivMinInt), "PV Bonus : " + strconv.Itoa(a.VieBonus), "Poids : " + strconv.Itoa(a.Poids), "Débloqué : " + strconv.FormatBool(a.EstDebloque)}, true, true)
}

func (s *Sort) Affichage() {
	Affichage("Sort", []string{"Nom : " + s.Nom, "Prix : " + strconv.Itoa(s.Prix), "Cout Mana : " + strconv.Itoa(s.ManaCout), "Dégâts : " + strconv.Itoa(s.Degats), "Boost PV : " + strconv.Itoa(s.BoostVie), "Type : " + strconv.Itoa(s.Type), "Débloqué : " + strconv.FormatBool(s.EstDebloque)}, true, true)
}

func Attendre() {
	fmt.Println("Appuyez sur 0 pour continuer")
	var choix string
	fmt.Scan(&choix)
	for choix != "0" {
		fmt.Println("Appuyez sur 0 pour continuer")
		fmt.Scan(&choix)
	}
}

func Choix(nb1, nb2 int) int {
	var choix string
	fmt.Scan(&choix)
	chiffre, err := strconv.Atoi(choix)
	for err != nil || chiffre < nb1 || chiffre > nb2 {
		fmt.Println("Veuillez entrer un nombre entre", nb1, "et", nb2)
		fmt.Scan(&choix)
		chiffre, err = strconv.Atoi(choix)
	}
	return chiffre
}
