package database

import (
	"fmt"
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
	fmt.Println(longest)
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
	stat := p.GetStatList()
	Affichage("Personnage", stat)
}

func (p *Personnage) Affichage_Inventaire() {
	nom := "None"
	longueur := 50
	//affichage du haut de la boite
	fmt.Print("╒", strings.Repeat("═", longueur), "╕", "\n")
	//affichage du titre
	fmt.Print("│", "Inventaire", strings.Repeat(" ", longueur-len("Inventaire")+1), "│", "\n")
	//affichage de la ligne
	fmt.Print("╞", strings.Repeat("═", longueur), "╡", "\n")
	//affichage des armes débloquées
	fmt.Print("│", "Armes débloquées", strings.Repeat(" ", longueur-len("Armes débloquées")+1), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index, _ := range p.Inv.Liste_armes {
		if p.Inv.Liste_armes[index].Get_Armes("isUnlocked") == "true" {
			nom = p.Inv.Liste_armes[index].Get_Armes("nom")
			fmt.Print("│", nom, strings.Repeat(" ", longueur-len(nom)), "│", "\n")
		}
	}
	//affichage des boucliers débloqués
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	fmt.Print("│", "Boucliers débloqués", strings.Repeat(" ", longueur-len("Boucliers débloqués")+1), "│", "\n")
	fmt.Print("├", strings.Repeat("─", longueur), "┤", "\n")
	for index, _ := range p.Inv.Liste_boucliers {
		if p.Inv.Liste_boucliers[index].Get_Boucliers("isUnlocked") == "true" {
			nom = p.Inv.Liste_boucliers[index].Get_Boucliers("nom")
			fmt.Print("│", nom, strings.Repeat(" ", longueur-len(nom)), "│", "\n")
		}
	}
}

//player.Inv.Liste_armes[0].Affichage_Arme()
