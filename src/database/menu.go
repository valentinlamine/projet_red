package database

import (
	"fmt"
	"os"
)

func Menu() {
	joueur.Est_Mort()
	Affichage("Menu - "+joueur.Position.Val["nom"], []string{"Que voulez-vous faire ?", "1. Accéder aux statistiques du personnage", "2. Accéder à l'inventaire du personnage", "3. Se déplacer", "4. Accèder au menu de lige feu", "5. Quitter le jeu"}, true, false)
	var choix = Choix(1, 6)
	switch choix {
	case 1:
		joueur.Affichage_Personnage()
	case 2:
		joueur.Affichage_Inventaire()
	case 3:
		Menu_Deplacement()
	case 4:
		Menu_Hub()
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
		Menu_Triche()

	}
}

func Menu_Triche() {
	Affichage("Menu cheat", []string{"Quel cheat souhaitez vous utilisez ?", "0. Quittez le menu", "1. Ajouter 1000 ames", "2. Ajouter 100 éclat de titanite", "3. Ajouter 100 grands éclats de titanite", "4. Ajouter 100 tablettes éclats de titanite", "5. Mettre une vie infini"}, true, false)
	var choix = Choix(0, 5)
	switch choix {
	case 0:
		return
	case 1:
		joueur.Ames += 1000
		Affichage("Menu cheat", []string{"Vous avez maintenant " + fmt.Sprint(joueur.Ames) + " ames"}, false, false)
		Menu_Triche()
	case 2:
		joueur.Inv.Liste_Items["éclat de titanite"] += 100
		Affichage("Menu cheat", []string{"Vous avez maintenant " + fmt.Sprint(joueur.Inv.Liste_Items["éclat de titanite"]) + " éclats de titanite"}, false, false)
		Menu_Triche()
	case 3:
		joueur.Inv.Liste_Items["grand éclat de titanite"] += 100
		Affichage("Menu cheat", []string{"Vous avez maintenant " + fmt.Sprint(joueur.Inv.Liste_Items["grand éclat de titanite"]) + " grands éclats de titanite"}, false, false)
		Menu_Triche()
	case 4:
		joueur.Inv.Liste_Items["tablette de titanite"] += 100
		Affichage("Menu cheat", []string{"Vous avez maintenant " + fmt.Sprint(joueur.Inv.Liste_Items["tablette de titanite"]) + " tablettes éclats de titanite"}, false, false)
		Menu_Triche()
	case 5:
		joueur.VieMax = 999999999
		joueur.VieAct = 999999999
		Affichage("Menu cheat", []string{"Vous avez maintenant une vie infini"}, false, false)
		Menu_Triche()
	}

}

func Menu_Deplacement() {
	boucle := false
	Affichage("Menu déplacement", []string{"Vous êtes à " + joueur.Position.Val["nom"], "Où voulez-vous aller ?", "1. Devant", "2. Gauche", "3. Droite", "4. Revenir en arrière", "5. Retour à Lige feu", "6. Retour au menu principal"}, true, false)
	var choix = Choix(1, 6)
	switch choix {
	case 1:
		boucle = joueur.Deplacement("centre")
	case 2:
		boucle = joueur.Deplacement("gauche")
	case 3:
		boucle = joueur.Deplacement("droite")
	case 4:
		boucle = joueur.Deplacement("retour")
	case 5:
		Affichage("Avertissement", []string{"Vous allez retourner à Lige feu", "Si vous retournez à Lige feu voici ce qu'il va se produire :", "● Vous allez perdre 20 ames", "● La carte sera regénéré", "", "Voulez-vous continuer ?", "1. Oui", "2. Non"}, false, false)
		choix = Choix(1, 2)
		switch choix {
		case 1:
			boucle = joueur.Deplacement("hub")
		case 2:
			boucle = true
		}
	case 6:
		return
	}
	if boucle {
		Menu_Deplacement()
	}
}

func Menu_Hub() {
	if joueur.Position.Val["nom"] != "Lige feu" {
		Affichage("Avertissement", []string{"Vous n'êtes pas à Lige feu", "Vous ne pouvez donc pas accéder au menu de Lige Feu", "Voulez-vous retourner à Lige feu ?", "1. Oui", "2. Non"}, false, false)
		var choix = Choix(1, 2)
		switch choix {
		case 1:
			Affichage("Avertissement", []string{"Vous allez retourner à Lige feu", "Si vous retournez à Lige feu voici ce qu'il va se produire :", "● Vous allez perdre 20 ames", "● La carte sera regénéré", "", "Voulez-vous continuer ?", "1. Oui", "2. Non"}, false, false)
			choix = Choix(1, 2)
			switch choix {
			case 1:
				joueur.Deplacement("hub")
			case 2:
				return
			}
		case 2:
			return
		}
	}
	Affichage("Lige feu", []string{"Vous êtes à " + joueur.Position.Val["nom"], "Que voulez-vous faire ?", "1. Aller voir André forgeron", "2. Aller voir André marchand", "3. Aller voir la gardienne du feu", "4. Aller vois Laurentius", "5. Aller voir Pétrus", "6. Revenir au menu principal"}, true, false)
	var choix = Choix(1, 6)
	switch choix {
	case 1:
		Menu_Forgeron()
	case 2:
		m1.Menu_Marchand(&joueur)
	case 3:
		m2.Menu_Marchand(&joueur)
	case 4:
		m3.Menu_Marchand(&joueur)
	case 5:
		m4.Menu_Marchand(&joueur)
	case 6:
		return
	}
}

func Menu_Forgeron() {
	Affichage("Forgeron", []string{"Bonjour " + joueur.Nom, "que veux tu faire aujourd'hui ?", "1. Améliorer un item", "2. Discuter", "3. Revenir au menu principal"}, true, false)
	var choix = Choix(1, 3)
	switch choix {
	case 1:
		joueur.Forgeron_Amelioration()
	case 2:
		Affichage("André", []string{"Bonjour aventurier", "Je me présente, Je suis le forgeron de ce village"}, true, true)
		Affichage("André", []string{"Je peux améliorer tes armes, tes boucliers et tes armures", "Mais pour cela il me faut des ressources que tu peux récupérer sur les monstres", "Et évidemment je reste un buisness man", "Chaque amélioration te coutera des Ames"}, true, true)
		Affichage("André", []string{"Voici le prix de mes améliorations :", "", "Pour une amélioration de tier 1, il te faudra :  ", "  ● 6 éclats de titanite", "  ● 100 âmes", "Pour une amélioration de tier 2, il te faudra :  ", "  ● 3 éclats de titanite", "  ● 3 grands éclats de titanite", "  ● 500 âmes", "Pour une amélioration de tier 3, il te faudra :  ", "  ● 2 éclats de titanite", "  ● 2 grands éclats de titanite", "  ● 2 tablettes éclats de titanite", "  ● 2000 âmes"}, true, true)
		Menu_Forgeron()
	}
}
