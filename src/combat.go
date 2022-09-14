package main

var p1 Personnage
var mob Personnage
p1.Init(p1, "Guerrier")
mob.Init(mob, "Mendiant") 

func combat(p1, p2 Personnage) {

}

		//Fichier Personnage

type Personnage struct {
	nom    string // Nom du personnage
	classe string // Classe du personnage (Guerrier, Chevalier, pyromancien, mendiant)
	niveau int    // Niveau du personnage
	//stats améliorables
	pvmax        int // pvmax du personnage (points de vie)
	vitalite     int // Vitalité du personnage (modificateur de pvmax)
	force        int // Force du personnage (dégâts infligés, ports d'armes, etc.)
	dexterite    int // Dextérité du personnage (dégâts infligés, esquive, etc.)
	intelligence int // Intelligence du personnage (dégâts magiques, etc.)
	//état du personnage
	pvact      int // Points de vie actuels
	inventaire []string
}
func (p *Personnage) Init(nom, classe string) {
	if classe == "Guerrrier" {
		p.nom = nom
		p.classe = classe
		p.niveau = 1
		p.pvmax = 10
		p.vitalite = 11
		p.force = 13
		p.dexterite = 13
		p.intelligence = 9
		p.pvact = p.pvmax
	} else if classe == "Chevalier" {
		p.nom = nom
		p.classe = classe
		p.niveau = 1
		p.pvmax = 10
		p.vitalite = 14
		p.force = 11
		p.dexterite = 11
		p.intelligence = 9
		p.pvact = p.pvmax
	} else if classe == "Pyromancien" {
		p.nom = nom
		p.classe = classe
		p.niveau = 1
		p.pvmax = 10
		p.vitalite = 10
		p.force = 12
		p.dexterite = 9
		p.intelligence = 10
		p.pvact = p.pvmax
	} else if classe == "Mendiant" {
		p.nom = nom
		p.classe = classe
		p.niveau = 1
		p.pvmax = 10
		p.vitalite = 11
		p.force = 11
		p.dexterite = 11
		p.intelligence = 11
		p.pvact = p.pvmax
	}
}