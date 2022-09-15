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
	Inventaire []string
}

func (p *Personnage) Init(nom, classe string) {
	if classe == "Guerrrier" {
		p.Niveau = 1
		p.Pvmax = 10
		p.Vitalite = 11
		p.Force = 13
		p.Dexterite = 13
		p.Intelligence = 9
	} else if classe == "Chevalier" {
		p.Niveau = 1
		p.Pvmax = 10
		p.Vitalite = 14
		p.Force = 11
		p.Dexterite = 11
		p.Intelligence = 9
	} else if classe == "Pyromancien" {
		p.Niveau = 1
		p.Pvmax = 10
		p.Vitalite = 10
		p.Force = 12
		p.Dexterite = 9
		p.Intelligence = 10
	} else if classe == "Mendiant" {
		p.Niveau = 1
		p.Pvmax = 10
		p.Vitalite = 11
		p.Force = 11
		p.Dexterite = 11
		p.Intelligence = 11
	}
	p.Nom = nom
	p.Classe = classe
	p.Pvact = p.Pvmax
}
