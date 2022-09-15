package database

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
	Inventaire [][]string
}

func (p *Personnage) InitIntern(nom, classe string, Vit, For, Dex, Int int) {
	p.Nom = nom
	p.Classe = classe
	p.Niveau = 1
	p.Pvmax = 20 * p.Vitalite
	p.Vitalite = Vit
	p.Force = For
	p.Dexterite = Dex
	p.Intelligence = Int
	p.Pvact = p.Pvmax

}

func (p *Personnage) Init(nom, classe string) {
	if classe == "Guerrrier" {
		p.Init(nom, "Guerrier", 11, 12, 9, 8)
	} else if classe == "Chevalier" {
		p.Init(nom, "Chevalier", 10, 11, 8, 10)
	} else if classe == "Pyromancien" {
		p.Init(nom, "Pyromancien", 9, 9, 11, 12)
	} else if classe == "Mendiant" {
		p.Init(nom, "Mendiant", 10, 10, 10, 10)
	}
}
