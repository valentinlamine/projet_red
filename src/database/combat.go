package database

import "fmt"

func Combat(player, mob *Personnage) {
	Affichage("Combat", []string{"Vous avez affronté un " + mob.Nom})
	for player.Pvact > 0 && mob.Pvact > 0 {
		var turnCount int
		charTurn(player, mob, turnCount)
		EnemyTurn(player, mob, turnCount)
	}
}

func (p *Personnage) Attaquer(target *Personnage) {
	target.Pvact -= p.Degats
}

func charTurn(player, mob *Personnage, turnC int) {
	var Choix int
	Affichage("Combat", []string{"Que voulez vous faire ?", "1. Attaquer", "2. Inventaire", "3. Fuir"})
	fmt.Scan(&Choix)
	switch Choix {
	case 1:
		player.Attaquer(mob)
		turnC++
	case 2:
		player.Affichage_Inventaire()
		turnC++
	}
}

func EnemyTurn(player, mob *Personnage, turnC int) {
	if turnC%3 == 0 {
		player.Pvact -= mob.Degats * 2
	} else {
		player.Pvact -= mob.Degats
	}
	turnC++
}
