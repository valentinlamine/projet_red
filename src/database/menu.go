package database

import (
	"fmt"
	"os"
)

func Menu() {
	player.IsDead()
	Affichage("Menu - "+player.Position.Val["name"], []string{"Que voulez-vous faire ?", "1. Accéder aux statistiques du personnage", "2. Accéder à l'inventaire du personnage", "3. Se déplacer", "4. Accèder au menu de la carte", "5. Quitter le jeu", "6. Menu cheat"}, true, false)
	var choix = Choix(1, 6)
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
		Affichage("Quitter le jeu", []string{"Voulez-vous vraiment quitter le jeu ?", "1. Oui", "2. Non"}, true, false)
		var choix = Choix(1, 2)
		switch choix {
		case 1:
			Affichage("Fin du jeu", []string{"Merci d'avoir joué !"}, false, false)
			os.Exit(-1)
		case 2:
			return
		}
	case 6:
		menu_cheat()

	}
}

func menu_cheat() {
	Affichage("Menu cheat", []string{"Quel cheat souhaitez vous utilisez ?", "0. Quittez le menu", "1. Ajouter 1000 ames", "2. Ajouter 100 éclat de titanite", "3. Ajouter 100 grands éclats de titanite", "4. Ajouter 100 tablettes éclats de titanite", "5. Mettre une vie infini"}, true, false)
	var choix = Choix(0, 5)
	switch choix {
	case 0:
		return
	case 1:
		player.Ames += 1000
		Affichage("Menu cheat", []string{"Vous avez maintenant " + fmt.Sprint(player.Ames) + " ames"}, false, false)
		menu_cheat()
	case 2:
		player.Inv.Liste_items["éclat de titanite"] += 100
		Affichage("Menu cheat", []string{"Vous avez maintenant " + fmt.Sprint(player.Inv.Liste_items["éclat de titanite"]) + " éclats de titanite"}, false, false)
		menu_cheat()
	case 3:
		player.Inv.Liste_items["grand éclat de titanite"] += 100
		Affichage("Menu cheat", []string{"Vous avez maintenant " + fmt.Sprint(player.Inv.Liste_items["grand éclat de titanite"]) + " grands éclats de titanite"}, false, false)
		menu_cheat()
	case 4:
		player.Inv.Liste_items["tablette de titanite"] += 100
		Affichage("Menu cheat", []string{"Vous avez maintenant " + fmt.Sprint(player.Inv.Liste_items["tablette de titanite"]) + " tablettes éclats de titanite"}, false, false)
		menu_cheat()
	case 5:
		player.Pvmax = 999999999
		player.Pvact = 999999999
		Affichage("Menu cheat", []string{"Vous avez maintenant une vie infini"}, false, false)
		menu_cheat()
	}

}

func Menu_deplacement() {
	boucle := false
	Affichage("Menu déplacement", []string{"Vous êtes à " + player.Position.Val["name"], "Où voulez-vous aller ?", "1. Devant", "2. Gauche", "3. Droite", "4. Revenir en arrière", "5. Retour au hub", "6. Retour au menu principal"}, true, false)
	var choix = Choix(1, 6)
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
		Affichage("Avertissement", []string{"Vous allez retourner au hub", "Si vous retournez au hub voici ce qu'il va se produire :", "● Vous allez perdre 20 ames", "● La carte sera regénéré", "", "Voulez-vous continuer ?", "1. Oui", "2. Non"}, false, false)
		choix = Choix(1, 2)
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
	Affichage("Lige feu", []string{"Vous êtes à " + player.Position.Val["name"], "Que voulez-vous faire ?", "1. Aller voir le forgeron", "2. Aller voir le marchand forgeron", "3. Aller voir le marchand de niveau", "4. Aller vois le marchand de sort", "5. Aller voir le marchand de consommable"}, true, false)
	var choix = Choix(1, 5)
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
	Affichage("Forgeron", []string{"Bonjour " + player.Nom, "que veux tu faire aujourd'hui ?", "1. Améliorer un item", "2. Discuter", "3. Revenir au menu principal"}, true, false)
	var choix = Choix(1, 3)
	switch choix {
	case 1:
		player.Forgeron_amelioration()
	case 2:
		Affichage("André", []string{"Bonjour aventurier", "Je me présente, Je suis le forgeron de ce village"}, true, true)
		Affichage("André", []string{"Je peux améliorer tes armes, tes boucliers et tes armures", "Mais pour cela il me faut des ressources que tu peux récupérer sur les monstres", "Et évidemment je reste un buisness man", "Chaque amélioration te coutera des Ames"}, true, true)
		Affichage("André", []string{"Voici le prix de mes améliorations :", "", "Pour une amélioration de tier 1, il te faudra :  ", "  ● 6 éclats de titanite", "  ● 100 âmes", "Pour une amélioration de tier 2, il te faudra :  ", "  ● 3 éclats de titanite", "  ● 3 grands éclats de titanite", "  ● 500 âmes", "Pour une amélioration de tier 3, il te faudra :  ", "  ● 2 éclats de titanite", "  ● 2 grands éclats de titanite", "  ● 2 tablettes éclats de titanite", "  ● 2000 âmes"}, true, true)
		Menu_forgeron()
	}
}
