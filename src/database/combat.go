package database

import (
	"math/rand"
	"strconv"
	"time"
)

func Combat(player, mob *Personnage) bool {
	Affichage("Combat", []string{"Vous affrontez un " + mob.Nom}, true, false)
	var turnCount int = 1
	if player.Initiative > mob.Initiative {
		for player.Pvact > 0 && mob.Pvact > 0 {
			time.Sleep(1 * time.Second)
			if mob.charTurn(player) {
				return true
			}
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
			if mob.charTurn(player) {
				return true
			}
			player.Mana += 15
			turnCount++
		}
	}
	return false
}

func (mob *Personnage) charTurn(player *Personnage) bool {
	if player.IsDead() {
		return true
	}
	Affichage("Combat", []string{"Que voulez vous faire ?", "1. Attaquer", "2. Attaque lourde", "3. Sorts", "4. Inventaire"}, false, false)
	var Choix = Choix(1, 4)
	switch Choix {
	case 1:
		mob.Pvact -= player.Degats
		if mob.Pvact < 0 {
			mob.Pvact = 0
			player.Ames += mob.Ames
			Affichage("Combat", []string{"Vous avez tué le " + mob.Nom}, true, false)
		} else {
			Affichage(player.Nom, []string{"Vous avez infligé " + strconv.Itoa(player.Degats) + " dégats au " + mob.Nom + ", il lui reste " + strconv.Itoa(mob.Pvact) + " pv"}, false, false)
		}
	case 2:
		mob.Pvact -= player.Degats * 2
		player.Mana -= 80
		if mob.Pvact < 0 {
			mob.Pvact = 0
			player.Ames += mob.Ames
			switch mob.Nom {
			case "Carcasse":
				player.Inv.Liste_items["éclat de titanite"] += rand.Intn(2) * 3
			case "Chevalier mort-vivant":
				player.Inv.Liste_items["grand éclat de titanite"] += rand.Intn(2) * 3
			case "Champion mort-vivant":
				player.Inv.Liste_items["tablette de titanite"] += rand.Intn(2) * 3
			case "Gargouille":
				player.Inv.Liste_items["éclat de titanite"] += rand.Intn(2) * 5
				player.Inv.Liste_items["tablette de titanite"] += rand.Intn(2) * 2
				if !player.Inv.Liste_armes[5].IsUnlocked {
					player.Inv.Liste_armes[5].IsUnlocked = true
					Affichage("Combat", []string{"Vous avez débloqué la hache de guerre"}, true, false)
				}
			case "Démon Capra":
				player.Inv.Liste_items["éclat de titanite"] += rand.Intn(2) * 5
				player.Inv.Liste_items["grand éclat de titanite"] += rand.Intn(2) * 2
			case "Démon taureau":
				player.Inv.Liste_items["éclat de titanite"] += rand.Intn(2) * 7
				player.Inv.Liste_items["grand éclat de titanite"] += rand.Intn(2) * 7
				player.Inv.Liste_items["tablette de titanite"] += rand.Intn(2) * 7
			}
			Affichage("Combat", []string{"Vous avez tué le " + mob.Nom}, true, false)
		} else {
			Affichage(player.Nom, []string{"Vous avez infligé " + strconv.Itoa(player.Degats*2) + " dégats au " + mob.Nom + ", il lui reste " + strconv.Itoa(mob.Pvact) + " pv", "Vous avez maintenant " + strconv.Itoa(player.Mana) + " mana"}, false, false)
		}
	case 3:
		Menu_sort(player, mob)
	case 4:
		player.Affichage_Inventaire()
	}
	return false
}

func (player *Personnage) EnemyTurn(mob *Personnage, turnC int) {
	if player.IsDead() {
		return
	}
	if mob.Pvact > 0 {
		if turnC%3 == 0 {
			player.Pvact -= mob.Degats * 2
			if player.Pvact > 0 {
				Affichage(mob.Nom, []string{"Le " + mob.Nom + " vous a infligé " + strconv.Itoa(mob.Degats*2) + " dégats, vous avez maintenant " + strconv.Itoa(player.Pvact) + " pv"}, false, false)
			}
		} else if turnC%4 == 0 {
			Affichage(mob.Nom, []string{"Le " + mob.Nom + " vous a empoisonné"}, false, false)
			player.Inv.Liste_consommables[5].Quantite++
			player.PrendrePot(mob.Inv.Liste_consommables[5])
		} else {
			player.Pvact -= mob.Degats
			if player.Pvact < 0 {
				player.Pvact = 0
			}
			if player.Pvact > 0 {
				Affichage(mob.Nom, []string{"Le " + mob.Nom + " vous a infligé " + strconv.Itoa(mob.Degats) + " dégats, vous avez maintenant " + strconv.Itoa(player.Pvact) + " pv"}, false, false)
			}
		}
	}
}

func (p *Personnage) Do_combat() bool {
	loop := true
	for loop {
		c, _ := strconv.Atoi(p.Position.Val["mob_nb"])
		if c > 0 {
			Affichage("Déplacement", []string{"Nous nous dirigeons vers " + p.Position.Val["name"], "Cet endroit est infesté de " + p.Position.Val["mob_type"] + " il y en a " + p.Position.Val["mob_nb"], "Que voulez vous faire ?", "1. Combattre", "2. Fuir"}, false, false)
			var Choix = Choix(1, 2)
			if Choix == 1 {
				var mob Personnage
				var inv Inventaire
				inv.Init()
				mob.Inv = inv
				mob.Init("Mob", p.Position.Val["mob_type"])
				if Combat(p, &mob) {
					return false
				}
				c--
				p.Position.Val["mob_nb"] = strconv.Itoa(c)
			} else if Choix == 2 {
				Affichage("Déplacement", []string{"Ce chemin à l'air dangereur, peut être que revenir en arrière était la bonne idée"}, true, true)
				loop = false
			}
		}
		if c == 0 {
			Affichage("Déplacement", []string{"Nous arrivons vers " + p.Position.Val["name"]}, true, true)
			return true
		}
	}
	return false
}
