package main

import (
	"fmt"
	"src/database"
	"strings"
)

func main() {
	database.Test()
	//var p1 Personnage
	fmt.Println("Création de votre personnage")
	fmt.Println("Quel est le nom de votre personnage ?")
	var nom string
	fmt.Scan(&nom)
	affichage([]string{"Quelle est la classe de votre personnage ?", "il y a 4 classes : Guerrier, Chevalier, Pyromancien, Mendiant", "Pour choisir guerrier, tapez 1", "Pour choisir chevalier, tapez 2", "Pour choisir pyromancien, tapez 3", "Pour choisir mendiant, tapez 4"})
	var classe int
	fmt.Scan(&classe)
	/*	if classe == 1 {
			p1.Init(nom, "Guerrier")
		} else if classe == 2 {
			p1.Init(nom, "Chevalier")
		} else if classe == 3 {
			p1.Init(nom, "Pyromancien")
		} else if classe == 4 {
			p1.Init(nom, "Mendiant")
		} else {
			fmt.Println("Classe inconnue, Veuiilez choisir entre Guerrier, Chevalier, Pyromancien ou Mendiant")
		}*/
}

func affichage(list []string) {
	longest := 0
	for _, s := range list {
		if len(s) > longest {
			longest = len(s)
		}
	}
	fmt.Println("╒", strings.Repeat("═", longest+1), "╕")
	for _, s := range list {
		fmt.Println("│", s, strings.Repeat(" ", longest-len(s)), "│")
	}
	fmt.Println("╘", strings.Repeat("═", longest+1), "╛")

}
