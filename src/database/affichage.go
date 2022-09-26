package database

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rivo/uniseg"
)

func NouvelAffichage(titre string, contenu []string) {
	//laisser 8 ligne de vide au dessus
	fmt.Print("\n\n\n\n\n\n\n\n")
	//affichage
	Affichage(titre, contenu)
}

func Affichage(titre string, list []string) {
	longest := 0
	for _, s := range list {
		if uniseg.GraphemeClusterCount(s) > longest {
			longest = uniseg.GraphemeClusterCount(s)
		}
	}
	if uniseg.GraphemeClusterCount(titre) > longest {
		longest = uniseg.GraphemeClusterCount(titre)
	}
	//affichage du haut de la boite
	fmt.Print("╒", strings.Repeat("═", longest), "╕", "\n")
	//affichage du titre
	fmt.Print("│", titre, strings.Repeat(" ", longest-uniseg.GraphemeClusterCount(titre)), "│", "\n")
	//affichage de la ligne
	fmt.Print("├", strings.Repeat("─", longest), "┤", "\n")
	//affichage des lignes de texte
	for _, s := range list {
		//fmt.Println(longest, len(s), longest-len(s))
		fmt.Print("│", s, strings.Repeat(" ", longest-uniseg.GraphemeClusterCount(s)), "│", "\n")
	}
	//affichage du bas de la boite
	fmt.Print("╘", strings.Repeat("═", longest), "╛", "\n")
}

func (p *Personnage) Affichage_Personnage() {
	Affichage("Personnage", []string{"Nom : " + p.Nom, "Classe : " + p.Classe, "Niveau : " + strconv.Itoa(p.Niveau), "Ames : " + strconv.Itoa(p.Ames), "PV : " + strconv.Itoa(p.Pvact) + "/" + strconv.Itoa(p.Pvmax), "Poids : " + strconv.Itoa(p.PoidsEquip) + "/" + strconv.Itoa(p.PoidsMax), "Vitalité: " + strconv.Itoa(p.Vitalite), "Force : " + strconv.Itoa(p.Force), "Dextérité : " + strconv.Itoa(p.Dexterite), "Intelligence : " + strconv.Itoa(p.Intelligence), "Position : " + p.Position.Val["name"], "Mana : " + strconv.Itoa(p.Mana) + "/" + strconv.Itoa(p.ManaMax), "Dégâts : " + strconv.Itoa(p.Degats)})
	fmt.Println("Appuyez sur 1 pour voir votre inventaire équipé")
	fmt.Println("Appuyez sur 0 pour revenir au menu principal")
	var choix = Choix(0, 1)
	switch choix {
	case 1:
		p.Affichage_inventaire_equipe()
		Attendre()
	case 0:
		return
	}
}

func (p *Personnage) Affichage_Inventaire() {
	var nom string
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
	//affichage des armes débloquées
	fmt.Print("│", "Armes débloquées :", strings.Repeat(" ", longueur-18), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_armes); index++ {
		if p.Inv.Liste_armes[index].IsUnlocked {
			nom = p.Inv.Liste_armes[index].nom
			nb_objets++
			list_objets = append(list_objets, nom)
			fmt.Print("│  ", nb_objets, ". ", nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(nom)-5), "│", "\n")
		}
	}
	//affichage des boucliers débloqués
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Boucliers débloqués :", strings.Repeat(" ", longueur-21), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_boucliers); index++ {
		if p.Inv.Liste_boucliers[index].IsUnlocked {
			nom = p.Inv.Liste_boucliers[index].nom
			nb_objets++
			list_objets = append(list_objets, nom)
			fmt.Print("│  ", nb_objets, ". ", nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(nom)-5), "│", "\n")
		}
	}
	//affichage des consommables débloqués
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Consommables débloqués :", strings.Repeat(" ", longueur-24), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_consommables); index++ {
		if p.Inv.Liste_consommables[index].Quantite > 0 {
			nom = p.Inv.Liste_consommables[index].Nom
			nb_objets++
			list_objets = append(list_objets, nom)
			fmt.Print("│  ", nb_objets, ". ", nom, " ("+strconv.Itoa(p.Inv.Liste_consommables[index].Quantite)+")", strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(nom)-(8+len(strconv.Itoa(p.Inv.Liste_consommables[index].Quantite)))), "│", "\n")
		}
	}
	//affichage des armures de tête débloquées
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Casques débloquées :", strings.Repeat(" ", longueur-20), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_armures_tete); index++ {
		if p.Inv.Liste_armures_tete[index].isUnlocked {
			nom = p.Inv.Liste_armures_tete[index].nom
			nb_objets++
			list_objets = append(list_objets, nom)
			fmt.Print("│  ", nb_objets, ". ", nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(nom)-5), "│", "\n")
		}
	}
	//affichage des armures de torse débloquées
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Plastrons débloquées :", strings.Repeat(" ", longueur-22), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_armures_torse); index++ {
		if p.Inv.Liste_armures_torse[index].isUnlocked {
			nom = p.Inv.Liste_armures_torse[index].nom
			nb_objets++
			list_objets = append(list_objets, nom)
			fmt.Print("│  ", nb_objets, ". ", nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(nom)-5), "│", "\n")
		}
	}
	//affichage des armures de jambes débloquées
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Jambières débloquées :", strings.Repeat(" ", longueur-22), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_armures_jambes); index++ {
		if p.Inv.Liste_armures_jambes[index].isUnlocked {
			nom = p.Inv.Liste_armures_jambes[index].nom
			nb_objets++
			list_objets = append(list_objets, nom)
			fmt.Print("│  ", nb_objets, ". ", nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(nom)-6), "│", "\n")
		}
	}
	//affichage des armures de bras débloquées
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Brassards débloquées :", strings.Repeat(" ", longueur-22), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_armures_bras); index++ {
		if p.Inv.Liste_armures_bras[index].isUnlocked {
			nom = p.Inv.Liste_armures_bras[index].nom
			nb_objets++
			list_objets = append(list_objets, nom)
			fmt.Print("│  ", nb_objets, ". ", nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(nom)-6), "│", "\n")
		}
	}
	//affichage du bas de la boite
	fmt.Print("╘", strings.Repeat("═", longueur), "╛", "\n")

	//choix de l'objet à afficher
	fmt.Print("Choisissez un objet à afficher (0 pour quitter) : ")
	var choix = Choix(0, len(list_objets))
	if choix != 0 {
		//affichage de l'objet choisi en utilisant la fonction Affichage
		item := p.Inv.Get_Item(list_objets[choix-1])
		switch item := item.(type) {
		case Consommable:
			Affichage("Inventaire", []string{"Vous avez choisi l'objet : " + item.Nom, "Que voulez vous faire avec cet objet ? ", "1. Le consommer", "2. Afficher ces statistiques", "3. Retourner au menu précédent"})
		default:
			Affichage("Inventaire", []string{"Vous avez choisi l'objet : " + list_objets[choix-1], "Que voulez vous faire avec cet objet ? ", "1. L'équiper", "2. Afficher ces statistiques", "3. Retourner au menu précédent"})
		}
		var choix2 int
		fmt.Scan(&choix2)
		for choix2 < 1 || choix2 > 3 {
			fmt.Print("Choix incorrect, veuillez réessayer : ")
			fmt.Scan(&choix2)
		}
		if choix2 == 1 {
			p.Equiper(item) //si consommable la fonction Equiper() va le consommer
		} else if choix2 == 2 {
			switch item := item.(type) {
			case Armes:
				item.Affichage()
				Attendre()
				p.Affichage_Inventaire()
			case Boucliers:
				item.Affichage()
				Attendre()
				p.Affichage_Inventaire()
			case Consommable:
				item.Affichage()
				Attendre()
				p.Affichage_Inventaire()
			case Armures:
				item.Affichage()
				Attendre()
				p.Affichage_Inventaire()
			default:
				print("Erreur")
			}
		} else if choix2 == 3 {
			p.Affichage_Inventaire()
		}
	}
}

func (p *Personnage) Affichage_inventaire_equipe() {
	longueur := 50
	//affichage du haut de la boite
	fmt.Print("╒", strings.Repeat("═", longueur), "╕", "\n")
	fmt.Print("│", "Inventaire équipé :", strings.Repeat(" ", longueur-19), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	//affichage de l'arme équipée
	fmt.Print("│", "Arme équipée :", strings.Repeat(" ", longueur-14), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│ ● ", p.EquipementArmes.nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(p.EquipementArmes.nom)-3), "│", "\n")
	//affichage du bouclier équipé
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Bouclier équipé :", strings.Repeat(" ", longueur-17), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│ ● ", p.EquipementBoucliers.nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(p.EquipementBoucliers.nom)-3), "│", "\n")
	//affichage de l'armure de tête équipée
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Casque équipée :", strings.Repeat(" ", longueur-16), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│ ● ", p.EquipementArmures["Tete"].nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(p.EquipementArmures["Tete"].nom)-3), "│", "\n")
	//affichage de l'armure de torse équipée
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Plastron équipée :", strings.Repeat(" ", longueur-18), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│ ● ", p.EquipementArmures["Torse"].nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(p.EquipementArmures["Torse"].nom)-3), "│", "\n")
	//affichage de l'armure de jambes équipée
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Jambières équipée :", strings.Repeat(" ", longueur-19), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│ ● ", p.EquipementArmures["Jambes"].nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(p.EquipementArmures["Jambes"].nom)-3), "│", "\n")
	//affichage de l'armure de bras équipée
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Brassards équipée :", strings.Repeat(" ", longueur-19), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│ ● ", p.EquipementArmures["Bras"].nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(p.EquipementArmures["Bras"].nom)-3), "│", "\n")
	//affichage du bas de la boite
	fmt.Print("└", strings.Repeat("─", longueur), "┘", "\n")
}

func (a *Armes) Affichage() {
	Affichage("Arme", []string{"Nom : " + a.nom, "Stat Min Force : " + strconv.Itoa(a.lvlMinFor), "Stat Min Dextérité : " + strconv.Itoa(a.lvlMinDex), "Stat Min Intelligence : " + strconv.Itoa(a.lvlMinInt), "Dégâts : " + strconv.Itoa(a.deg), "Poids : " + strconv.Itoa(a.poids), "Débloqué : " + strconv.FormatBool(a.IsUnlocked), "Équipé : " + strconv.FormatBool(a.IsEquiped)})
}

func (b *Boucliers) Affichage() {
	Affichage("Bouclier", []string{"Nom : " + b.nom, "Stat Min Force : " + strconv.Itoa(b.lvlMinFor), "Stat Min Dextérité : " + strconv.Itoa(b.lvlMinDex), "Stat Min Intelligence : " + strconv.Itoa(b.lvlMinInt), "PV Bonus : " + strconv.Itoa(b.pvbonus), "Poids : " + strconv.Itoa(b.poids), "Débloqué : " + strconv.FormatBool(b.IsUnlocked), "Équipé : " + strconv.FormatBool(b.IsEquiped)})
}

func (c *Consommable) Affichage() {
	Affichage("Consommable", []string{"Nom : " + c.Nom, "Prix : " + strconv.Itoa(c.Prix), "Quantité : " + strconv.Itoa(c.Quantite), "PV Bonus : " + strconv.Itoa(c.PvBonus), "Bonus Force : " + strconv.Itoa(c.MultiLvlFor), "Bonus Dextérité : " + strconv.Itoa(c.MultiLvlDex), "Bonus Intelligence : " + strconv.Itoa(c.MultiLvlInt), "Bonus Poids Max : " + strconv.Itoa(c.MultiLvlPoidsMax)})
}

func (a *Armures) Affichage() {
	Affichage("Armure", []string{"Nom : " + a.nom, "Stat Min Force : " + strconv.Itoa(a.lvlMinFor), "Stat Min Dextérité : " + strconv.Itoa(a.lvlMinDex), "Stat Min Intelligence : " + strconv.Itoa(a.lvlMinInt), "PV Bonus : " + strconv.Itoa(a.pvbonus), "Poids : " + strconv.Itoa(a.poids), "Débloqué : " + strconv.FormatBool(a.isUnlocked)})
}

func Attendre() {
	fmt.Println("Appuyez sur 0 pour continuer")
	var choix int
	fmt.Scan(&choix)
	for choix != 0 {
		fmt.Println("Appuyez sur 0 pour continuer")
		fmt.Scan(&choix)
	}
}

func Choix(nb1, nb2 int) int {
	var a int
	fmt.Scan(&a)
	for a < nb1 || a > nb2 {
		fmt.Println("Vous devez choisir un nombre entre", nb1, "et", nb2)
		fmt.Scan(&a)
	}
	return a
}
