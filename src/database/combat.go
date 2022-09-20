package database

import "fmt"

func (p *Personnage) Combat() {
	Affichage("COMBAT", []string{"Que voulez-vous faire ?", "1. Attaquer", "2. Esquiver", "3. Ouvrir inventaire", "4. Fuir", "5. Se déplacer", "6. Aller voir le marchand mort vivant"})
	var choix int
	fmt.Scan(&choix)
	for choix < 1 || choix > 6 {
		fmt.Println("Vous devez choisir un choix entre 1 et 6")
		fmt.Scan(&choix)
	}
}
