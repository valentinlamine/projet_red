package database

import "strings"

type Personnage struct {
	Nom    string // Nom du personnage
	Classe string // Classe du personnage (Guerrier, Chevalier, pyromancien, mendiant)
	Niveau int    // Niveau du personnage
	//stats améliorables
	Pvmax        int // Pvmax du personnage (points de vie)
	Vitalite     int // Vitalité du personnage (modificateur de Pvmax)
	Force        int // Force du personnage (dégâts infligés, ports d'armes, etc.)
	Dexterite    int // Dextérité du personnage (dégâts infligés, esquive, etc.)
	Intelligence int // Intelligence du personnage (dégâts magiques, etc.)
	//état du personnage
	Pvact int // Points de vie actuels
	Inv   Inventaire
	Ames  int
}

func (p *Personnage) InitIntern(nom, classe string, Vit, For, Dex, Int int) {
	p.Nom = nom
	p.Classe = classe
	p.Niveau = 1
	p.Vitalite = Vit
	p.Pvmax = 20 * p.Vitalite
	p.Force = For
	p.Dexterite = Dex
	p.Intelligence = Int
	p.Pvact = p.Pvmax / 2
	p.Ames = 0

}

func (p *Personnage) Init(nom, classe string) {
	//modifie le nom pour mettre la première lettre en majuscule et le reste en minuscule
	for i, c := range nom {
		if i == 0 {
			nom = strings.ToUpper(string(c))
		} else {
			nom += strings.ToLower(string(c))
		}
	}

	print(nom)
	if classe == "Guerrrier" {
		p.InitIntern(nom, "Guerrier", 11, 12, 9, 8)
	} else if classe == "Chevalier" {
		p.InitIntern(nom, "Chevalier", 10, 11, 8, 10)
	} else if classe == "Pyromancien" {
		p.InitIntern(nom, "Pyromancien", 9, 9, 11, 12)
	} else if classe == "Mendiant" {
		p.InitIntern(nom, "Mendiant", 9, 9, 9, 9)
	}
}
