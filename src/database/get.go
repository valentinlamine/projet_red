package database

import "strconv"

func (a *Armes) Affichage_Arme() {
	Affichage("Arme", []string{"Nom : " + a.nom, "Stat Min Force : " + strconv.Itoa(a.lvlMinFor), "Stat Min Dextérité : " + strconv.Itoa(a.lvlMinDex), "Stat Min Intelligence : " + strconv.Itoa(a.lvlMinInt), "Dégâts : " + strconv.Itoa(a.deg), "Poids : " + strconv.Itoa(a.poids), "Débloqué : " + strconv.FormatBool(a.isUnlocked), "Équipé : " + strconv.FormatBool(a.isEquiped)})
}

func (b *Boucliers) Affichage_Bouclier() {
	Affichage("Bouclier", []string{"Nom : " + b.nom, "Stat Min Force : " + strconv.Itoa(b.lvlMinFor), "Stat Min Dextérité : " + strconv.Itoa(b.lvlMinDex), "Stat Min Intelligence : " + strconv.Itoa(b.lvlMinInt), "PV Bonus : " + strconv.Itoa(b.pvbonus), "Poids : " + strconv.Itoa(b.poids), "Débloqué : " + strconv.FormatBool(b.isUnlocked), "Équipé : " + strconv.FormatBool(b.isEquiped)})
}

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

func (b *Boucliers) Get_Bouclier(data string) string {
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

func (b *Boucliers) Set_Bouclier(data string, value string) {
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
