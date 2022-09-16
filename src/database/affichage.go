package database

import (
	"fmt"
	"strconv"
	"strings"
)

func Affichage(titre string, list []string) {
	longest := 0
	for _, s := range list {
		if len(s) > longest {
			longest = len(s)
		}
	}
	if len(titre) > longest {
		longest = len(titre)
	}
	//affichage du haut de la boite
	fmt.Print("╒", strings.Repeat("═", longest), "╕", "\n")
	//affichage du titre
	fmt.Print("│", titre, strings.Repeat(" ", longest-len(titre)+1), "│", "\n")
	//affichage de la ligne
	fmt.Print("├", strings.Repeat("─", longest), "┤", "\n")
	//affichage des lignes de texte
	for _, s := range list {
		//fmt.Println(longest, len(s), longest-len(s))
		fmt.Print("│", s, strings.Repeat(" ", longest-len(s)), "│", "\n")
	}
	//affichage du bas de la boite
	fmt.Print("╘", strings.Repeat("═", longest), "╛", "\n")
}

func (p *Personnage) Affichage_Personnage() {
	//Bizarre que ça marche grâce au 1er Get_Personnage sans avoir besoin de changer les autres appels
	Affichage("Personnage", []string{"Nom : " + p.Get_Personnage("Nom"), "Niveau : " + p.Get_Personnage("Niveau"), "PV : " + strconv.Itoa(p.Pvact), "PV Max : " + strconv.Itoa(p.Pvmax), "Vitalité: " + strconv.Itoa(p.Vitalite), "Force : " + strconv.Itoa(p.Force), "Dextérité : " + strconv.Itoa(p.Dexterite), "Intelligence : " + strconv.Itoa(p.Intelligence), "Ames : " + strconv.Itoa(p.Ames)})
}

func (p *Personnage) Affichage_Inventaire() {
	nom := "None"
	longueur := 50
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
			fmt.Print("│  ● ", nom, strings.Repeat(" ", longueur-len(nom)-4), "│", "\n")
		}
	}
	//affichage des boucliers débloqués
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Boucliers débloqués :", strings.Repeat(" ", longueur-len("Boucliers débloqués")), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_boucliers); index++ {
		if p.Inv.Liste_boucliers[index].Get_Boucliers("isUnlocked") == "true" {
			nom = p.Inv.Liste_boucliers[index].Get_Boucliers("nom")
			fmt.Print("│  ● ", nom, strings.Repeat(" ", longueur-len(nom)-4), "│", "\n")
		}
	}
	//affichage des consommables débloqués
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Consommables débloqués :", strings.Repeat(" ", longueur-len("Consommables débloqués")), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index := 0; index < len(p.Inv.Liste_consommables); index++ {
		if p.Inv.Liste_consommables[index].Get_Consommable("isUnlocked") == "true" {
			nom = p.Inv.Liste_consommables[index].Get_Consommable("nom")
			fmt.Print("│  ● ", nom, strings.Repeat(" ", longueur-len(nom)-4), "│", "\n")
		}
	}
	//affichage du bas de la boite
	fmt.Print("╘", strings.Repeat("═", longueur), "╛", "\n")
}

func (a *Armes) Affichage_Arme() {
	Affichage("Arme", []string{"Nom : " + a.nom, "Stat Min Force : " + strconv.Itoa(a.lvlMinFor), "Stat Min Dextérité : " + strconv.Itoa(a.lvlMinDex), "Stat Min Intelligence : " + strconv.Itoa(a.lvlMinInt), "Dégâts : " + strconv.Itoa(a.deg), "Poids : " + strconv.Itoa(a.poids), "Débloqué : " + strconv.FormatBool(a.isUnlocked), "Équipé : " + strconv.FormatBool(a.isEquiped)})
}

func (b *Boucliers) Affichage_Bouclier() {
	Affichage("Bouclier", []string{"Nom : " + b.nom, "Stat Min Force : " + strconv.Itoa(b.lvlMinFor), "Stat Min Dextérité : " + strconv.Itoa(b.lvlMinDex), "Stat Min Intelligence : " + strconv.Itoa(b.lvlMinInt), "PV Bonus : " + strconv.Itoa(b.pvbonus), "Poids : " + strconv.Itoa(b.poids), "Débloqué : " + strconv.FormatBool(b.isUnlocked), "Équipé : " + strconv.FormatBool(b.isEquiped)})
}

//player.Inv.Liste_armes[0].Affichage_Arme()
