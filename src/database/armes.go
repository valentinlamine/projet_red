package database

type Armes struct {
	// Nom et Prix
	Nom  string
	Prix int
	// Stat Min
	LvlMinFor int
	LvlMinDex int
	LvlMinInt int
	// Stat Damage
	Deg        int
	Poids      int
	IsUnlocked bool
	IsEquiped  bool
}

type Boucliers struct {
	// Nom et Prix
	Nom  string
	Prix int
	// Stat Min
	LvlMinFor int
	LvlMinDex int
	LvlMinInt int
	// Stat Damage
	Pvbonus    int
	Poids      int
	IsUnlocked bool
	IsEquiped  bool
}

func (a *Armes) InitIntern_Armes(nom string, Prix, lvlMinFor, lvlMinDex, lvlMinInt, deg, poids int) {
	a.Nom = nom
	a.Prix = Prix
	a.LvlMinFor = lvlMinFor
	a.LvlMinDex = lvlMinDex
	a.LvlMinInt = lvlMinInt
	a.Deg = deg
	a.Poids = poids
	a.IsUnlocked = false
	a.IsEquiped = false

}

func (a *Armes) Init_Armes(nom string) {
	if nom == "Dague" {
		a.InitIntern_Armes("Dague", 100, 9, 11, 8, 20, 3)
	} else if nom == "Claymore" {
		a.InitIntern_Armes("Claymore", 2000, 11, 9, 8, 50, 8)
	} else if nom == "Rapière" {
		a.InitIntern_Armes("Rapière", 500, 9, 12, 9, 35, 5)
	} else if nom == "Uchigatana" {
		a.InitIntern_Armes("Uchigatana", 1000, 10, 14, 10, 40, 6)
	} else if nom == "Bâton" {
		a.InitIntern_Armes("Bâton", 100, 7, 7, 7, 15, 3)
	} else if nom == "Hache queue de gargouille" {
		a.InitIntern_Armes("Hache queue de gargouille", 5000, 14, 10, 8, 60, 12)
	} else if nom == "vide" {
		a.InitIntern_Armes("vide", 0, 0, 0, 0, 0, 0)
	}
}

func (b *Boucliers) InitIntern_Bouclier(nom string, Prix, lvlMinFor, lvlMinDex, lvlMinInt, pvbonus, poids int) {
	b.Nom = nom
	b.Prix = Prix
	b.LvlMinFor = lvlMinFor
	b.LvlMinDex = lvlMinDex
	b.LvlMinInt = lvlMinInt
	b.Pvbonus = pvbonus
	b.Poids = poids
	b.IsUnlocked = false
	b.IsEquiped = false

}

func (b *Boucliers) Init_Bouclier(nom string) {
	if nom == "Bouclier en bois" {
		b.InitIntern_Bouclier("Bouclier en bois", 50, 9, 8, 8, 20, 3)
	} else if nom == "Bouclier en bois de cerf" {
		b.InitIntern_Bouclier("Bouclier en bois de cerf", 100, 10, 8, 8, 40, 5)
	} else if nom == "Bouclier en fer" {
		b.InitIntern_Bouclier("Bouclier en fer", 700, 12, 8, 8, 75, 9)
	} else if nom == "Bouclier en acier" {
		b.InitIntern_Bouclier("Bouclier en acier", 500, 10, 12, 9, 65, 5)
	} else if nom == "Bouclier en mithril" {
		b.InitIntern_Bouclier("Bouclier en mithril", 1200, 12, 16, 10, 120, 6)
	} else if nom == "Bouclier d'Havel" {
		b.InitIntern_Bouclier("Bouclier d'Havel", 5000, 20, 10, 10, 175, 20)
	} else if nom == "vide" {
		b.InitIntern_Bouclier("vide", 0, 0, 0, 0, 0, 0)
	}
}

func (a Armes) ReturnIsEquiped() bool {
	return a.IsEquiped
}

func (b Boucliers) ReturnIsEquiped() bool {
	return b.IsEquiped
}
