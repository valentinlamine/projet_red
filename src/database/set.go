package database

import "strconv"

func (a *Armes) Set_Armes(data string, value string) {
	if data == "nom" {
		a.nom = value
	} else if data == "lvlMinFor" {
		a.lvlMinFor, _ = strconv.Atoi(value)
	} else if data == "lvlMinDex" {
		a.lvlMinDex, _ = strconv.Atoi(value)
	} else if data == "lvlMinInt" {
		a.lvlMinInt, _ = strconv.Atoi(value)
	} else if data == "deg" {
		a.deg, _ = strconv.Atoi(value)
	} else if data == "poids" {
		a.poids, _ = strconv.Atoi(value)
	} else if data == "isUnlocked" {
		a.isUnlocked, _ = strconv.ParseBool(value)
	} else if data == "isEquiped" {
		a.isEquiped, _ = strconv.ParseBool(value)
	}
}

func (b *Boucliers) Set_Boucliers(data string, value string) {
	if data == "nom" {
		b.nom = value
	} else if data == "lvlMinFor" {
		b.lvlMinFor, _ = strconv.Atoi(value)
	} else if data == "lvlMinDex" {
		b.lvlMinDex, _ = strconv.Atoi(value)
	} else if data == "lvlMinInt" {
		b.lvlMinInt, _ = strconv.Atoi(value)
	} else if data == "pvbonus" {
		b.pvbonus, _ = strconv.Atoi(value)
	} else if data == "poids" {
		b.poids, _ = strconv.Atoi(value)
	} else if data == "isUnlocked" {
		b.isUnlocked, _ = strconv.ParseBool(value)
	} else if data == "isEquiped" {
		b.isEquiped, _ = strconv.ParseBool(value)
	}
}

func (p *Personnage) Set_Personnage(data string, value string) {
	if data == "nom" {
		p.Nom = value
	} else if data == "classe" {
		p.Classe = value
	} else if data == "niveau" {
		p.Niveau, _ = strconv.Atoi(value)
	} else if data == "pvmax" {
		p.Pvmax, _ = strconv.Atoi(value)
	} else if data == "vitalite" {
		p.Vitalite, _ = strconv.Atoi(value)
	} else if data == "force" {
		p.Force, _ = strconv.Atoi(value)
	} else if data == "dexterite" {
		p.Dexterite, _ = strconv.Atoi(value)
	} else if data == "intelligence" {
		p.Intelligence, _ = strconv.Atoi(value)
	} else if data == "pvact" {
		p.Pvact, _ = strconv.Atoi(value)
	}
}

func (c *Consommable) Set_Consommable(data string, value string) {
	if data == "nom" {
		c.Nom = value
	} else if data == "classe" {
		c.Classe = value
	} else if data == "prix" {
		c.Prix, _ = strconv.Atoi(value)
	} else if data == "quantite" {
		c.Quantite, _ = strconv.Atoi(value)
	} else if data == "pvbonus" {
		c.PvBonus, _ = strconv.Atoi(value)
	} else if data == "multilvlfor" {
		c.MultiLvlFor, _ = strconv.Atoi(value)
	} else if data == "multilvldex" {
		c.MultiLvlDex, _ = strconv.Atoi(value)
	} else if data == "multilvlint" {
		c.MultiLvlInt, _ = strconv.Atoi(value)
	} else if data == "multilvlpoidsmax" {
		c.MultiLvlPoidsMax, _ = strconv.Atoi(value)
	}
}
