package database

import (
	"fmt"
	"strings"
)

var inventaire Inventaire
var carte Arbre
var player Personnage
var m1 Marchand
var m2 Marchand
var m3 Marchand
var m4 Marchand

func Init() {
	carte.Init()
	inventaire.Init()
	Setup_Personnage()
	//Initialisation des marchands
	m1.InitMarchand(1)
	m2.InitMarchand(2)
	m3.InitMarchand(3)
	m4.InitMarchand(4)
}

func Setup_Personnage() {
	Affichage("Création du personnage", []string{"Bienvenue dans le jeu de rôle !", "Pour commencer, vous devez créer votre personnage", "Choisissez un nom"}, true, false)
	var nom string
	fmt.Scan(&nom)
	//si le nom contient autre chose que des lettres, on demande de recommencer
	for _, c := range nom {
		if !strings.Contains("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZéèàùçâêîôäëïöüÿìò", string(c)) {
			Affichage("Erreur", []string{"Le nom ne peut contenir que des lettres", "Veuillez recommencer"}, false, false)
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
	Affichage("Création du personnage", []string{"Bonjour " + nom, "Quelle est la classe de ton personnage ?", "il y a 4 classes : Guerrier, Chevalier, Pyromancien, Mendiant", "Pour choisir guerrier, tapez 1", "Pour choisir chevalier, tapez 2", "Pour choisir pyromancien, tapez 3", "Pour choisir mendiant, tapez 4"}, true, false)
	player.Inv = inventaire
	var classe = Choix(1, 4)
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
	player.Position = carte
}
