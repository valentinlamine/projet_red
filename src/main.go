package main

import (
	"fmt"
	"os"
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
	for !player.IsDead() {
		Menu()
	}
	//à modifier
	if player.IsDead() {
		database.Affichage("Fin du jeu", []string{"Vous êtes mort !"})
	}
}

func setup_personnage() {
	database.Affichage("Création du personnage", []string{"Bienvenue dans le jeu de rôle !", "Pour commencer, vous devez créer votre personnage", "Choisissez un nom"})
	var nom string
	fmt.Scan(&nom)
	//si le nom contient autre chose que des lettres, on demande de recommencer
	for _, c := range nom {
		if !strings.Contains("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZéèàùçâêîôäëïöüÿìò", string(c)) {
			database.Affichage("Erreur", []string{"Le nom ne peut contenir que des lettres", "Veuillez recommencer"})
			fmt.Scan(&nom)
		}
	}
	//on met la première lettre en majuscule et le reste en minuscule
	for i, c := range nom {
		if i == 0 {
			nom = strings.ToUpper(string(c))
		} else {
			nom += strings.ToLower(string(c))
		}
	}
	database.Affichage("Création du personnage", []string{"Bonjour " + nom, "Quelle est la classe de ton personnage ?", "il y a 4 classes : Guerrier, Chevalier, Pyromancien, Mendiant", "Pour choisir guerrier, tapez 1", "Pour choisir chevalier, tapez 2", "Pour choisir pyromancien, tapez 3", "Pour choisir mendiant, tapez 4"})
	player.Inv = inventaire
	var classe int
	fmt.Scan(&classe)
	for classe < 1 || classe > 4 {
		fmt.Println("Vous devez choisir une classe entre 1 et 4")
		fmt.Scan(&classe)
	}
	switch classe {
	case 1:
		player.Init(nom, "Guerrier")
		player.Inv.Liste_armes[3].Set_Armes("isUnlocked", "true")         //débloquer Uchigatana
		player.Inv.Liste_boucliers[0].Set_Boucliers("isUnlocked", "true") //débloquer Bouclier de bois
	case 2:
		player.Init(nom, "Chevalier")
		player.Inv.Liste_armes[1].Set_Armes("isUnlocked", "true")         //débloquer Claymore
		player.Inv.Liste_boucliers[2].Set_Boucliers("isUnlocked", "true") //débloquer Bouclier de fer
	case 3:
		player.Init(nom, "Pyromancien")
		player.Inv.Liste_armes[0].Set_Armes("isUnlocked", "true") //débloquer Dague
	case 4:
		player.Init(nom, "Mendiant")
		player.Inv.Liste_armes[4].Set_Armes("isUnlocked", "true") //débloquer Baton
	}
}

func Menu() {
	database.Affichage("Menu", []string{"Que voulez-vous faire ?", "1. Accéder aux statistiques du personnage", "2. Accéder à l'inventaire du personnage", "3. Quitter le jeu", "6. Interragir avec le marchand mort-vivant"})
	var choix int
	fmt.Scan(&choix)
	for choix < 1 || choix > 6 {
		fmt.Println("Vous devez choisir un choix entre 1 et 3")
		fmt.Scan(&choix)
	}
	switch choix {
	case 1:
		player.Affichage_Personnage()
		attendre()
	case 2:
		player.Affichage_Inventaire()
		attendre()
	case 3:
		database.Affichage("Fin du jeu", []string{"Merci d'avoir joué !"})
		os.Exit(-1)
	case 4:
		player.PrendrePot(player.Inv.Liste_consommables[0])
		player.Affichage_Personnage()
	case 5:
		player.PrendrePot(player.Inv.Liste_consommables[5])
		player.Affichage_Personnage()
	case 6:
		database.Affichage("Marchand Mort-vivant", []string{"1. | 100 | Fiole d'essence de pin pourri", "2. | 100 | Bâton", "3. | 400 | Résine de pin doré"})
		var marchand int
		switch marchand {
		case 1:
			player.Ames -= player.Inv.Liste_consommables[5].Prix
			player.Inv.Liste_consommables[5].Quantite += 1

		}
	}
}

func attendre() {
	fmt.Println("Appuyez sur entrée pour continuer")
	fmt.Scan()
}
