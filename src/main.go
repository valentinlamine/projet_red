package main

import (
	"fmt"
	"src/database"
	"strings"
)

// initialisation des variables
var p1 database.Personnage
var carte database.Arbre

func main() {
	setup_personnage()
	Affichage_Personnage(p1)
	carte.Init()

}

func setup_personnage() {
	Affichage("Création du personnage", []string{"Bienvenue dans le jeu de rôle !", "Pour commencer, vous devez créer votre personnage", "Choisissez un nom"})
	var nom string
	fmt.Scan(&nom)
	Affichage("Création du personnage", []string{"Quelle est la classe de votre personnage ?", "il y a 4 classes : Guerrier, Chevalier, Pyromancien, Mendiant", "Pour choisir guerrier, tapez 1", "Pour choisir chevalier, tapez 2", "Pour choisir pyromancien, tapez 3", "Pour choisir mendiant, tapez 4"})
	var classe int
	fmt.Scan(&classe)
	for classe < 1 || classe > 4 {
		fmt.Println("Vous devez choisir une classe entre 1 et 4")
		fmt.Scan(&classe)
	}
	switch classe {
	case 1:
		p1.Init2(nom, "Guerrier")
	case 2:
		p1.Init2(nom, "Chevalier")
	case 3:
		p1.Init2(nom, "Pyromancien")
	case 4:
		p1.Init2(nom, "Mendiant")
	}
}

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

func Affichage_Personnage(p database.Personnage) {
	fmt.Println("disabled")
	//Affichage("Personnage", []string{"Nom : " + p.Nom, "Classe : " + p.Classe, "Niveau : " + string(p.Niveau), "Pvmax : " + string(rune(p.Pvmax)), "Vitalite : " + string(p.Vitalite), "Force : " + string(p.Force), "Dexterite : " + string(p.Dexterite), "Intelligence : " + string(p.Intelligence), "Pvact : " + string(p.Pvact)})
}
