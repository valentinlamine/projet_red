package database

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rivo/uniseg"
)

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
	//laisser 8 ligne de vide au dessus
	fmt.Print("\n\n\n\n\n\n\n\n")
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
	//Bizarre que ça marche grâce au 1er Get_Personnage sans avoir besoin de changer les autres appels
	Affichage("Personnage", []string{"Nom : " + p.Nom, "Niveau : " + strconv.Itoa(p.Niveau), "PV : " + strconv.Itoa(p.Pvact), "PV Max : " + strconv.Itoa(p.Pvmax), "Vitalité: " + strconv.Itoa(p.Vitalite), "Force : " + strconv.Itoa(p.Force), "Dextérité : " + strconv.Itoa(p.Dexterite), "Intelligence : " + strconv.Itoa(p.Intelligence), "Ames : " + strconv.Itoa(p.Ames)})
}

func (p *Personnage) Affichage_Inventaire() {
	nom := "None"
	longueur := 50
	var list_objets []string
	nb_objets := 0
	//laisser 8 ligne de vide au dessus
	fmt.Print("\n\n\n\n\n\n\n\n")
	//affichage du haut de la boite
	fmt.Print("╒", strings.Repeat("═", longueur), "╕", "\n")
	//affichage du titre
	fmt.Print("│", "Inventaire", strings.Repeat(" ", longueur-len("Inventaire")), "│", "\n")
	//affichage de la ligne
	fmt.Print("╞", strings.Repeat("═", longueur), "╡", "\n")
	//affichage des armes débloquées
	fmt.Print("│", "Armes débloquées :", strings.Repeat(" ", longueur-len("Armes débloquées")), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_armes); index++ {
		if p.Inv.Liste_armes[index].Get_Armes("isUnlocked") == "true" {
			nom = p.Inv.Liste_armes[index].Get_Armes("nom")
			nb_objets++
			list_objets = append(list_objets, nom)
			fmt.Print("│  ", nb_objets, ". ", nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(nom)-5), "│", "\n")
		}
	}
	//affichage des boucliers débloqués
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Boucliers débloqués :", strings.Repeat(" ", longueur-len("Boucliers débloqués")), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_boucliers); index++ {
		if p.Inv.Liste_boucliers[index].Get_Boucliers("isUnlocked") == "true" {
			nom = p.Inv.Liste_boucliers[index].Get_Boucliers("nom")
			nb_objets++
			list_objets = append(list_objets, nom)
			fmt.Print("│  ", nb_objets, ". ", nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(nom)-5), "│", "\n")
		}
	}
	//affichage des consommables débloqués
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Consommables débloqués :", strings.Repeat(" ", longueur-len("Consommables débloqués")), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_consommables); index++ {
		if p.Inv.Liste_consommables[index].Get_Consommable("isUnlocked") == "true" {
			nom = p.Inv.Liste_consommables[index].Get_Consommable("nom")
			nb_objets++
			list_objets = append(list_objets, nom)
			fmt.Print("│  ", nb_objets, ". ", nom, strings.Repeat(" ", longueur-uniseg.GraphemeClusterCount(nom)-5), "│", "\n")
		}
	}
	//affichage du bas de la boite
	fmt.Print("╘", strings.Repeat("═", longueur), "╛", "\n")

	//choix de l'objet à afficher
	fmt.Print("Choisissez un objet à afficher (0 pour quitter) : ")
	var choix int
	fmt.Scan(&choix)
	for choix < 0 || choix > len(list_objets) {
		fmt.Print("Choix incorrect, veuillez réessayer : ")
		fmt.Scan(&choix)
	}
	if choix != 0 {
		//affichage de l'objet choisi en utilisant la fonction Affichage
		item := p.Inv.Get_Item(list_objets[choix-1])
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
		default:
			print("Erreur")
		}
	}
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

func Attendre() {
	fmt.Println("Appuyez sur 0 pour continuer")
	var choix int
	fmt.Scan(&choix)
	for choix != 0 {
		fmt.Println("Appuyez sur 0 pour continuer")
		fmt.Scan(&choix)
	}
}
