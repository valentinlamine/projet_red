package main

import (
	"fmt"
	"src/database"
	"strings"
)

// initialisation des variables
var inventaire database.Inventaire
var carte database.Arbre
var player database.Personnage

func main() {
	inventaire.Init()
	carte.Init()
	setup_personnage()
	player.Affichage_Personnage()
	Menu()
}

func setup_personnage() {
	database.Affichage("Création du personnage", []string{"Bienvenue dans le jeu de rôle !", "Pour commencer, vous devez créer votre personnage", "Choisissez un nom"})
	var nom string
	fmt.Scan(&nom)
	//si le nom contient autre chose que des lettres, on demande de recommencer
	for _, c := range nom {
		if !strings.Contains("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", string(c)) {
			database.Affichage("Erreur", []string{"Le nom ne peut contenir que des lettres", "Veuillez recommencer"})
			fmt.Scan(&nom)
		}
	}
	database.Affichage("Création du personnage", []string{"Bonjour jeune aventurier !", "Quelle est la classe de ton personnage ?", "il y a 4 classes : Guerrier, Chevalier, Pyromancien, Mendiant", "Pour choisir guerrier, tapez 1", "Pour choisir chevalier, tapez 2", "Pour choisir pyromancien, tapez 3", "Pour choisir mendiant, tapez 4"})
	var classe int
	fmt.Scan(&classe)
	for classe < 1 || classe > 4 {
		fmt.Println("Vous devez choisir une classe entre 1 et 4")
		fmt.Scan(&classe)
	}
	switch classe {
	case 1:
		player.Init(nom, "Guerrier")
	case 2:
		player.Init(nom, "Chevalier")
	case 3:
		player.Init(nom, "Pyromancien")
	case 4:
		player.Init(nom, "Mendiant")
	}
	player.Inv = inventaire
}

func Menu() {
	database.Affichage("Menu", []string{"Que voulez-vous faire ?", "1. Accéder aux statistiques du personnage", "2. Accéder à l'inventaire du personnage", "3. Quitter le jeu"})
	var choix int
	fmt.Scan(&choix)
	for choix < 1 || choix > 3 {
		fmt.Println("Vous devez choisir un choix entre 1 et 3")
		fmt.Scan(&choix)
	}
	switch choix {
	case 1:
		player.Affichage_Personnage()
	case 2:
		player.Affichage_Inventaire() //TODO
	case 3:
		fmt.Println("quitter")
	}
}
