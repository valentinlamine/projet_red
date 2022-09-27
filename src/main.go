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

var mob database.Personnage
var inventaireMob database.Inventaire

func main() {
	inventaireMob.Init()
	mob.Inv = inventaireMob
	mob.Init("Mob", "Carcasse")
	inventaire.Init()
	carte.Init()
	m1.InitMarchand(1)
	setup_personnage()
	for {
		player.IsDead()
		Menu()
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
	database.Affichage("Menu", []string{"Que voulez-vous faire ?", "1. Accéder aux statistiques du personnage", "2. Accéder à l'inventaire du personnage", "3. Se déplacer", "4. Aller voir le marchand mort vivant", "5. Accéder au menu du hub", "6. Quitter le jeu", "7. Menu cheat"})
	var choix = database.Choix(1, 8)
	switch choix {
	case 1:
		player.Affichage_Personnage()
	case 2:
		fmt.Println("Vous avez maintenant ", player.Degats, " DPS")
		player.Affichage_Inventaire()
	case 3:
		Menu_deplacement()
	case 4:
		m1.Menu_Marchand(&player)
	case 5:
		Menu_hub()
	case 6:
		database.Affichage("Fin du jeu", []string{"Merci d'avoir joué !"})
		os.Exit(-1)
	case 7:
		menu_cheat()
	case 8:
		database.Combat(&player, &mob)
	}
}

func menu_cheat() {
	database.Affichage("Menu cheat", []string{"Quel cheat souhaitez vous utilisez ?", "0. Quittez le menu", "1. Ajouter 1000 ames", "2. Ajouter 100 éclat de titanite", "3. Ajouter 100 grands éclats de titanite", "4. Ajouter 100 tablettes éclats de titanite", "5. Mettre une vie infini"})
	var choix = database.Choix(0, 5)
	switch choix {
	case 0:
		return
	case 1:
		player.Ames += 1000
		database.Affichage("Menu cheat", []string{"Vous avez maintenant " + fmt.Sprint(player.Ames) + " ames"})
		menu_cheat()
	case 2:
		player.Inv.Liste_items["éclat de titanite"] += 100
		database.Affichage("Menu cheat", []string{"Vous avez maintenant " + fmt.Sprint(player.Inv.Liste_items["éclat de titanite"]) + " éclats de titanite"})
		menu_cheat()
	case 3:
		player.Inv.Liste_items["grand éclat de titanite"] += 100
		database.Affichage("Menu cheat", []string{"Vous avez maintenant " + fmt.Sprint(player.Inv.Liste_items["grand éclat de titanite"]) + " grands éclats de titanite"})
		menu_cheat()
	case 4:
		player.Inv.Liste_items["tablette éclat de titanite"] += 100
		database.Affichage("Menu cheat", []string{"Vous avez maintenant " + fmt.Sprint(player.Inv.Liste_items["tablette éclat de titanite"]) + " tablettes éclats de titanite"})
		menu_cheat()
	case 5:
		player.Pvmax = 99999
		player.Pvact = 99999
		database.Affichage("Menu cheat", []string{"Vous avez maintenant une vie infini"})
		menu_cheat()
	}

}

func Menu_deplacement() {
	boucle := false
	database.Affichage("Menu déplacement", []string{"Vous êtes à " + player.Position.Val["name"], "Où voulez-vous aller ?", "1. Devant", "2. Gauche", "3. Droite", "4. Revenir en arrière", "5. Retour au hub", "6. Retour au menu principal"})
	var choix = database.Choix(1, 5)
	switch choix {
	case 1:
		boucle = player.Deplacement("centre")
	case 2:
		boucle = player.Deplacement("gauche")
	case 3:
		boucle = player.Deplacement("droite")
	case 4:
		boucle = player.Deplacement("retour")
	case 5:
		database.Affichage("Avertissement", []string{"Vous allez retourner au hub", "Si vous retournez au hub voici ce qu'il va se produire :", "● Vous allez perdre 20 ames", "● La carte sera regénéré", "", "Voulez-vous continuer ?", "1. Oui", "2. Non"})
		choix = database.Choix(1, 2)
		switch choix {
		case 1:
			boucle = player.Deplacement("hub")
		case 2:
			boucle = true
		}
	case 6:
		return
	}
	if boucle {
		Menu_deplacement()
	}
}

func Menu_hub() {
	database.Affichage("Lige feu", []string{"Vous êtes à Lige feu", "Que voulez-vous faire ?", "1. Aller voir nos meilleurs marchand de la région", "2. Aller voir notre forgeron", ""})
}
