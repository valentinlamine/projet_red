package database

import (
	"strconv"
	"time"
)

func Combat(player, mob *Personnage) {
	NouvelAffichage("Combat", []string{"Vous affrontez un " + mob.Nom})
	var turnCount int = 1
	if player.Initiative > mob.Initiative {
		for player.Pvact > 0 && mob.Pvact > 0 {
			time.Sleep(1 * time.Second)
			mob.charTurn(player)
			time.Sleep(1 * time.Second)
			player.EnemyTurn(mob, turnCount)
			player.Mana += 15
			turnCount++
		}
	} else {
		for player.Pvact > 0 && mob.Pvact > 0 {
			time.Sleep(1 * time.Second)
			player.EnemyTurn(mob, turnCount)
			time.Sleep(1 * time.Second)
			mob.charTurn(player)
			player.Mana += 15
			turnCount++
		}
	}
}

func (mob *Personnage) charTurn(player *Personnage) {
	Affichage("Combat", []string{"Que voulez vous faire ?", "1. Attaquer", "2. Attaque lourde", "3. Sorts", "4. Inventaire"})
	var Choix = Choix(1, 4)
	switch Choix {
	case 1:
		mob.Pvact -= player.Degats
		if mob.Pvact < 0 {
			mob.Pvact = 0
			player.Ames += mob.Ames
			NouvelAffichage("Combat", []string{"Vous avez tué le " + mob.Nom})
		} else {
			NouvelAffichage(player.Nom, []string{"Vous avez infligé " + strconv.Itoa(player.Degats) + " dégats au " + mob.Nom + ", il lui reste " + strconv.Itoa(mob.Pvact) + " pv"})
		}
	case 2:
		mob.Pvact -= player.Degats * 2
		player.Mana -= 35
		if mob.Pvact < 0 {
			mob.Pvact = 0
			player.Ames += mob.Ames
			NouvelAffichage("Combat", []string{"Vous avez tué le " + mob.Nom})
		} else {
			NouvelAffichage(player.Nom, []string{"Vous avez infligé " + strconv.Itoa(player.Degats*2) + " dégats au " + mob.Nom + ", il lui reste " + strconv.Itoa(mob.Pvact) + " pv"})
		}
	case 3:
		Menu_sort(player, mob)
	case 4:
		player.Affichage_Inventaire()
	}
}

func (player *Personnage) EnemyTurn(mob *Personnage, turnC int) {
	if mob.Pvact > 0 {
		if turnC%3 == 0 {
			player.Pvact -= mob.Degats * 2
			Affichage(mob.Nom, []string{"Le " + mob.Nom + " vous a infligé " + strconv.Itoa(mob.Degats*2) + " dégats, il vous reste " + strconv.Itoa(player.Pvact) + " pv"})
		} else if turnC%4 == 0 {
			Affichage(mob.Nom, []string{"Le " + mob.Nom + " vous a empoisonné"})
			player.PrendrePot(mob.Inv.Liste_consommables[5])
		} else {
			player.Pvact -= mob.Degats
			if player.Pvact < 0 {
				player.Pvact = 0
			}
			Affichage(mob.Nom, []string{"Le " + mob.Nom + " vous a infligé " + strconv.Itoa(mob.Degats) + " dégats, vous avez maintenant " + strconv.Itoa(player.Pvact) + " pv"})
		}
	}
}

func (p *Personnage) Do_combat() bool {
	loop := true
	for loop {
		c, _ := strconv.Atoi(p.Position.Val["mob_nb"])
		if c > 0 {
			Affichage("Déplacement", []string{"Nous nous dirigeons vers " + p.Position.Val["name"], "Cet endroit est infesté de " + p.Position.Val["mob_type"] + " il y en a " + p.Position.Val["mob_nb"], "Que voulez vous faire ?", "1. Combattre", "2. Fuir"})
			var Choix = Choix(1, 2)
			if Choix == 1 {
				var mob Personnage
				var inv Inventaire
				inv.Init()
				mob.Inv = inv
				mob.Init("Mob", p.Position.Val["mob_type"])
				Combat(p, &mob)
				if p.IsDead() {
					Affichage("Combat", []string{"Vous êtes mort"})
					return false
				}
				c--
				p.Position.Val["mob_nb"] = strconv.Itoa(c)
			} else if Choix == 2 {
				Affichage("Déplacement", []string{"Ce chemin à l'air dangereur, peut être que revenir en arrière était la bonne idée"})
				Attendre()
				loop = false
			}
		}
		if c == 0 {
			Affichage("Déplacement", []string{"Nous arrivons vers " + p.Position.Val["name"]})
			Attendre()
			return true
		}
	}
	return false
}
