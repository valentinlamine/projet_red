package database

type Armes struct {
	// Nom
	nom string
	// Stat Min
	lvlMinFor int
	lvlMinDex int
	lvlMinInt int
	// Stat Damage
	deg        int
	poids      int
	isUnlocked bool
	isEquiped  bool
}

type Boucliers struct {
	// Nom
	nom string
	// Stat Min
	lvlMinFor int
	lvlMinDex int
	lvlMinInt int
	// Stat Damage
	pvbonus    int
	poids      int
	isUnlocked bool
	isEquiped  bool
}

func (a *Armes) InitIntern_Armes(nom string, lvlMinFor, lvlMinDex, lvlMinInt, deg, poids int) {
	a.nom = nom
	a.lvlMinFor = lvlMinFor
	a.lvlMinDex = lvlMinDex
	a.lvlMinInt = lvlMinInt
	a.deg = deg
	a.poids = poids
	a.isUnlocked = false

}

func (a *Armes) Init_Armes(nom string) {
	if nom == "Dague" {
		a.InitIntern_Armes("Dague", 9, 11, 8, 20, 3)
	} else if nom == "Claymore" {
		a.InitIntern_Armes("Claymore", 11, 9, 8, 50, 8)
	} else if nom == "Rapière" {
		a.InitIntern_Armes("Rapière", 9, 12, 9, 35, 5)
	} else if nom == "Uchigatana" {
		a.InitIntern_Armes("Uchigatana", 10, 14, 10, 40, 6)
	} else if nom == "Bâton" {
		a.InitIntern_Armes("Bâton", 7, 7, 7, 15, 3)
	} else if nom == "Hache queue de gargouille" {
		a.InitIntern_Armes("Hache queue de gargouille", 14, 10, 8, 60, 12)
	}
}

func (b *Boucliers) InitIntern_Bouclier(nom string, lvlMinFor, lvlMinDex, lvlMinInt, pvbonus, poids int) {
	b.nom = nom
	b.lvlMinFor = lvlMinFor
	b.lvlMinDex = lvlMinDex
	b.lvlMinInt = lvlMinInt
	b.pvbonus = pvbonus
	b.poids = poids
	b.isUnlocked = false

}

func (b *Boucliers) Init_Bouclier(nom string) {
	if nom == "Bouclier en bois" {
		b.InitIntern_Bouclier("Bouclier en bois", 9, 11, 8, 20, 3)
	} else if nom == "Bouclier en fer" {
		b.InitIntern_Bouclier("Bouclier en fer", 11, 9, 8, 50, 8)
	} else if nom == "Bouclier en acier" {
		b.InitIntern_Bouclier("Bouclier en acier", 9, 12, 9, 35, 5)
	} else if nom == "Bouclier en mithril" {
		b.InitIntern_Bouclier("Bouclier en mithril", 10, 14, 10, 40, 6)
	} else if nom == "Bouclier en bois de cerf" {
		b.InitIntern_Bouclier("Bouclier en bois de cerf", 14, 10, 8, 60, 12)
	}
}
