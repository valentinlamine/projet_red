package database

import (
	"bufio"
	"os"
	"strings"
)

var inventaire Inventaire
var carte Arbre
var joueur Personnage
var m1 Marchand
var m2 Marchand
var m3 Marchand
var m4 Marchand

var nom string

func Initialisation() {
	carte.Initialisation()
	inventaire.Initialisation()
	Initialisation_Personnage()
	//Initialisation des marchands
	m1.Initialisation_Marchand(1)
	m2.Initialisation_Marchand(2)
	m3.Initialisation_Marchand(3)
	m4.Initialisation_Marchand(4)
}

func Initialisation_Personnage() {
	Affichage("Création du personnage", []string{"Bienvenue dans le jeu de rôle !", "Pour commencer, vous devez créer votre personnage", "Choisissez un nom"}, true, false)
	//on met la première lettre en majuscule et le reste en minuscule
	Initialisation_Nom()
	for i, c := range nom {
		if i == 0 {
			nom = strings.ToUpper(string(c))
		} else {
			nom += strings.ToLower(string(c))
		}
	}
	Affichage("Création du personnage", []string{"Bonjour " + nom, "Quelle est la classe de ton personnage ?", "il y a 4 classes : Guerrier, Chevalier, Pyromancien, Mendiant", "Pour choisir guerrier, tapez 1", "Pour choisir chevalier, tapez 2", "Pour choisir pyromancien, tapez 3", "Pour choisir mendiant, tapez 4"}, true, false)
	joueur.Inv = inventaire
	var classe = Choix(1, 4)
	switch classe {
	case 1:
		joueur.Initialisation(nom, "Guerrier")
	case 2:
		joueur.Initialisation(nom, "Chevalier")
	case 3:
		joueur.Initialisation(nom, "Pyromancien")
	case 4:
		joueur.Initialisation(nom, "Mendiant")
	}
	joueur.Position = carte
}

func Initialisation_Nom() {
	boucle := true
	scanner := bufio.NewScanner(os.Stdin)

	for boucle {
		if scanner.Scan() {
			line := scanner.Text()
			nom = line
		}
		if len(nom) < 1 {
			Affichage("Création du personnage", []string{"Le nom ne peut être vide", "Choisissez un nom"}, false, false)
		} else {
			for _, letter := range nom {
				if (letter < 'A' || letter > 'Z') && (letter < 'a' || letter > 'z') {
					Affichage("Création du personnage", []string{"Le nom ne peut contenir que des lettres majuscule ou minuscule", "Choisissez un nom"}, false, false)
					Initialisation_Nom()
					return
				}
			}
			boucle = false
		}

	}
}
