package database

import "strconv"

func (a *Armes) Get_Armes(data string) string {
	if data == "nom" {
		return a.nom
	} else if data == "lvlMinFor" {
		return strconv.Itoa(a.lvlMinFor)
	} else if data == "lvlMinDex" {
		return strconv.Itoa(a.lvlMinDex)
	} else if data == "lvlMinInt" {
		return strconv.Itoa(a.lvlMinInt)
	} else if data == "deg" {
		return strconv.Itoa(a.deg)
	} else if data == "poids" {
		return strconv.Itoa(a.poids)
	} else if data == "isUnlocked" {
		return strconv.FormatBool(a.isUnlocked)
	} else if data == "isEquiped" {
		return strconv.FormatBool(a.isEquiped)
	}
	return ""
}

func (b *Boucliers) Get_Boucliers(data string) string {
	if data == "nom" {
		return b.nom
	} else if data == "lvlMinFor" {
		return strconv.Itoa(b.lvlMinFor)
	} else if data == "lvlMinDex" {
		return strconv.Itoa(b.lvlMinDex)
	} else if data == "lvlMinInt" {
		return strconv.Itoa(b.lvlMinInt)
	} else if data == "pvbonus" {
		return strconv.Itoa(b.pvbonus)
	} else if data == "poids" {
		return strconv.Itoa(b.poids)
	} else if data == "isUnlocked" {
		return strconv.FormatBool(b.isUnlocked)
	} else if data == "isEquiped" {
		return strconv.FormatBool(b.isEquiped)
	}
	return ""
}

/*
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
*/
func (p *Personnage) Get_Personnage(data string) string {
	if data == "Nom" {
		return p.Nom
	} else if data == "Classe" {
		return p.Classe
	} else if data == "Niveau" {
		return strconv.Itoa(p.Niveau)
	} else if data == "Pvmax" {
		return strconv.Itoa(p.Pvmax)
	} else if data == "Vitalite" {
		return strconv.Itoa(p.Vitalite)
	} else if data == "Force" {
		return strconv.Itoa(p.Force)
	} else if data == "Dexterite" {
		return strconv.Itoa(p.Dexterite)
	} else if data == "Intelligence" {
		return strconv.Itoa(p.Intelligence)
	} else if data == "Pvact" {
		return strconv.Itoa(p.Pvact)
	}
	return ""
}
