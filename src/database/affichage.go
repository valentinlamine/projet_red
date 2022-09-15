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

func Affichage_Personnage(p *Personnage) {
	Affichage("Personnage", []string{"Nom : " + p.Nom, "Classe : " + p.Classe, "Niveau : " + string(p.Niveau), "Vie" + string(p.Pvact), "Pv max : " + string(p.Pvmax), "Vitalité : " + string(p.Vitalite), "Force : " + string(p.Force), "Dexterite : " + string(p.Dexterite), "Intelligence : " + string(p.Intelligence)})
}
