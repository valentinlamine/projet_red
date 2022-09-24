package database

import "fmt"

func Combat(player, mob *Personnage) {
	var fuite bool
	Affichage("Combat", []string{"Vous avez affronté un " + mob.Nom})
	for player.Pvact > 0 && mob.Pvact > 0 || fuite == true {
		var turnCount int
		mob.charTurn(player, turnCount)
		player.EnemyTurn(mob, turnCount)
	}
}

func (mob *Personnage) charTurn(player *Personnage, turnC int) bool {
	var Choix int
	Affichage("Combat", []string{"Que voulez vous faire ?", "1. Attaquer", "2. Inventaire", "3. Fuir"})
	fmt.Scan(&Choix)
	switch Choix {
	case 1:
		mob.Pvact -= player.Degats
		fmt.Print("Vous avez infligé ", player.Degats, " dégats au ", mob.Nom, " il a maintenant ", mob.Pvact, " pv\n")
		turnC++
	case 2:
		player.Affichage_Inventaire()
		turnC++
	case 3:
		return true

	}
	return false
}

func (player *Personnage) EnemyTurn(mob *Personnage, turnC int) {
	if turnC%3 == 0 {
		player.Pvact -= mob.Degats * 2
		fmt.Print("Le ", mob.Nom, " vous a infligé ", mob.Degats*2, " dégats, vous avez maintenant ", player.Pvact, " pv")
	} else {
		player.Pvact -= mob.Degats
		fmt.Print("Le ", mob.Nom, " vous a infligé ", mob.Degats, " dégats, vous avez maintenant ", player.Pvact, " pv")
	}
	turnC++
}
