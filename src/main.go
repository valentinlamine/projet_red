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
var m1 database.Marchand

func main() {
	inventaire.Init()
	carte.Init()
	m1.InitMarchand(1)
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
	case 2:
		player.Init(nom, "Chevalier")
	case 3:
		player.Init(nom, "Pyromancien")
	case 4:
		player.Init(nom, "Mendiant")
	}
	player.Position = carte
}

func Menu() {
	database.Affichage("Menu", []string{"Que voulez-vous faire ?", "1. Accéder aux statistiques du personnage", "2. Accéder à l'inventaire du personnage", "3. Se déplacer", "4. Boire une potion de soin", "5. Boire une potion de poison", "6. Aller voir le marchand mort vivant", "7. Menu triche", "8. Quitter le jeu"})
	var choix = Choix(1, 8)
	switch choix {
	case 1:
		player.Affichage_Personnage()
		database.Attendre()
	case 2:
		player.Affichage_Inventaire()
	case 3:
		Menu_deplacement()
	case 4:
		player.PrendrePot(player.Inv.Liste_consommables[0])
		player.Affichage_Personnage()
	case 5:
		player.PrendrePot(player.Inv.Liste_consommables[5])
		player.Affichage_Personnage()
	case 6:
		m1.Menu_Marchand(&player)
	case 7:
		menu_cheat()
	case 8:
		database.Affichage("Fin du jeu", []string{"Merci d'avoir joué !"})
		os.Exit(-1)
	}
}

func menu_cheat() {
	database.Affichage("Menu cheat", []string{"Quel cheat souhaitez vous utilisez ?", "0. Quittez le menu", "1. Ajouter 1000 ames"})
	var choix int
	fmt.Scan(&choix)
	for choix < 0 || choix > 1 {
		fmt.Println("Vous devez choisir un choix entre 0 et 1")
		fmt.Scan(&choix)
	}
	switch choix {
	case 0:
		return
	case 1:
		player.Ames += 1000
		database.Affichage("Menu cheat", []string{"Vous avez maintenant " + fmt.Sprint(player.Ames) + " ames"})
		database.Attendre()
	}
}

func Menu_deplacement() {
	boucle := false
	database.Affichage("Menu déplacement", []string{"Vous êtes à " + player.Position.Val["name"], "Où voulez-vous aller ?", "1. Devant", "2. Gauche", "3. Droite", "4. Retour au hub", "5. Retour au menu principal"})
	var choix = Choix(1, 5)
	switch choix {
	case 1:
		boucle = player.Deplacement(carte, "centre")
	case 2:
		boucle = player.Deplacement(carte, "gauche")
	case 3:
		boucle = player.Deplacement(carte, "droite")
	case 4:
		database.Affichage("Avertissement", []string{"Vous allez retourner au hub", "Si vous retournez au hub voici ce qu'il va se produire :", "● Vous allez perdre 10 ames", "● La carte sera regénéré", "", "Voulez-vous continuer ?", "1. Oui", "2. Non"})
		choix = Choix(1, 2)
		switch choix {
		case 1:
			player.Ames -= 10
			boucle = player.Deplacement(carte, "retour")
		case 2:
			boucle = true
		}
	case 5:
		return
	}
	if boucle {
		Menu_deplacement()
	}
}

func Choix(nb1, nb2 int) int {
	var a int
	fmt.Scan(&a)
	for a < nb1 || a > nb2 {
		fmt.Println("Vous devez choisir un nombre entre", nb1, "et", nb2)
		fmt.Scan(&a)
	}
	return a
}
