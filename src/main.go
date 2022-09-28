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
var m2 database.Marchand
var m3 database.Marchand
var m4 database.Marchand

func main() {
	carte.Init()
	inventaire.Init()
	setup_personnage()
	//Initialisation des marchands
	m1.InitMarchand(1)
	m2.InitMarchand(2)
	m3.InitMarchand(3)
	m4.InitMarchand(4)
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
	var classe = database.Choix(1, 4)
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
	database.Affichage("Menu", []string{"Que voulez-vous faire ?", "1. Accéder aux statistiques du personnage", "2. Accéder à l'inventaire du personnage", "3. Se déplacer", "4. Accèder au menu de la carte", "5. Quitter le jeu", "6. Menu cheat"})
	var choix = database.Choix(1, 8)
	switch choix {
	case 1:
		player.Affichage_Personnage()
	case 2:
		player.Affichage_Inventaire()
	case 3:
		Menu_deplacement()
	case 4:
		Menu_hub()
	case 5:
		database.Affichage("Quitter le jeu", []string{"Voulez-vous vraiment quitter le jeu ?", "1. Oui", "2. Non"})
		var choix = database.Choix(1, 2)
		switch choix {
		case 1:
			database.Affichage("Fin du jeu", []string{"Merci d'avoir joué !"})
			os.Exit(-1)
		case 2:
			return
		}
	case 6:
		menu_cheat()

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
	database.Affichage("Lige feu", []string{"Vous êtes à " + player.Position.Val["name"], "Que voulez-vous faire ?", "1. Aller voir le forgeron", "2. Aller voir le marchand forgeron", "3. Aller voir le marchand de niveau", "4. Aller vois le marchand de sort", "5. Aller voir le marchand de consommable"})
	var choix = database.Choix(1, 5)
	switch choix {
	case 1:
		Menu_forgeron()
	case 2:
		m1.Menu_Marchand(&player)
	case 3:
		m2.Menu_Marchand(&player)
	case 4:
		m3.Menu_Marchand(&player)
	case 5:
		m4.Menu_Marchand(&player)
	}
}

func Menu_forgeron() {
	database.Affichage("Forgeron", []string{"Bonjour " + player.Nom, "que veux tu faire aujourd'hui ?", "1. Améliorer un item", "2. Discuter", "3. Revenir au menu principal"})
	var choix = database.Choix(1, 3)
	switch choix {
	case 1:
		player.Forgeron_amelioration()
	case 2:
		database.Affichage("André", []string{"Bonjour aventurier", "Je me présente, Je suis le forgeron de ce village"})
		database.Attendre()
		database.Affichage("André", []string{"Je peux améliorer tes armes, tes boucliers et tes armures", "Mais pour cela il me faut des ressources que tu peux récupérer sur les monstres", "Et évidemment je reste un buisness man", "Chaque amélioration te coutera des Ames"})
		database.Attendre()
		database.Affichage("André", []string{"Voici le prix de mes améliorations :", "", "Pour une amélioration de tier 1, il te faudra :  ", "  ● 6 éclats de titanite", "  ● 100 âmes", "Pour une amélioration de tier 2, il te faudra :  ", "  ● 3 éclats de titanite", "  ● 3 grands éclats de titanite", "  ● 500 âmes", "Pour une amélioration de tier 3, il te faudra :  ", "  ● 2 éclats de titanite", "  ● 2 grands éclats de titanite", "  ● 2 tablettes éclats de titanite", "  ● 2000 âmes"})
		database.Attendre()
		Menu_forgeron()
	}
}
