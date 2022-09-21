package database

import "fmt"

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
	Pvact      int // Points de vie actuels
	Inv        Inventaire
	PoidsEquip int // Poids total des objets équipés
	PoidsMax   int // Poids maximum que peut porter le personnage
	Degats     int
	Ames       int // Nombre d'âmes du personnage(Argent/EXP)
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
	p.PoidsEquip = 0
	p.PoidsMax = 5 * p.Force
	p.Ames = 0

}

func (p *Personnage) Init(nom, classe string) {
	//modifie le nom pour mettre la première lettre en majuscule et le reste en minuscule
	if classe == "Guerrier" {
		p.InitIntern(nom, "Guerrier", 11, 12, 9, 8)
	} else if classe == "Chevalier" {
		p.InitIntern(nom, "Chevalier", 10, 11, 8, 10)
	} else if classe == "Pyromancien" {
		p.InitIntern(nom, "Pyromancien", 9, 9, 11, 12)
	} else if classe == "Mendiant" {
		p.InitIntern(nom, "Mendiant", 9, 9, 9, 9)
	}
}

func (p *Personnage) IsDead() bool {
	if p.Pvact <= 0 {
		p.Ames = 0
		//print in red "Vous êtes mort !"
		fmt.Println("\033[31mVous êtes mort !\033[0m")
		return true
	} else {
		return false
	}
}
