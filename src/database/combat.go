package database

import "fmt"

func (p *Personnage) Combat(mob *Personnage) {
	Affichage("COMBAT", []string{"Que voulez-vous faire ?", "1. Attaquer", "2. Esquiver", "3. Bloquer", "4. Ouvrir inventaire", "5. Fuir"})
	var choixCombat int
	fmt.Scan(&choixCombat)
	for choixCombat < 1 || choixCombat > 6 {
		fmt.Println("Vous devez choisir un choix entre 1 et 6")
		fmt.Scan(&choixCombat)
	}
	switch choixCombat {
	case 1:
		p.Inv

	}
}
