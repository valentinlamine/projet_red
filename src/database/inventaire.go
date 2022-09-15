package database

//Liste des variables des armes
var dague Armes
var claymore Armes
var rapière Armes
var uchigatana Armes
var bâton Armes
var hachequeue Armes

//Liste des variables des boucliers
var bouclier_bois Boucliers
var Bouclier_bois_cerf Boucliers
var Bouclier_fer Boucliers
var Bouclier_acier Boucliers
var Bouclier_mithril Boucliers
var Bouclier_Havel Boucliers

type Inventaire struct {
	Liste_armes        []Armes
	Liste_consommables []Consommable
	Liste_boucliers    []Boucliers
	Liste_sort         []string //TODO Sort
}

func (i *Inventaire) Init() {
	//Initialisation des armes
	dague.Init_Armes("Dague")
	claymore.Init_Armes("Claymore")
	rapière.Init_Armes("Rapière")
	uchigatana.Init_Armes("Uchigatana")
	bâton.Init_Armes("Bâton")
	hachequeue.Init_Armes("Hache queue de gargouille")
	i.Liste_armes = append(i.Liste_armes, dague, claymore, rapière, uchigatana, bâton, hachequeue)
	//Initialisation des boucliers
	bouclier_bois.Init_Bouclier("Bouclier en bois")
	Bouclier_bois_cerf.Init_Bouclier("Bouclier en bois de cerf")
	Bouclier_fer.Init_Bouclier("Bouclier en fer")
	Bouclier_acier.Init_Bouclier("Bouclier en acier")
	Bouclier_mithril.Init_Bouclier("Bouclier en mithril")
	Bouclier_Havel.Init_Bouclier("Bouclier d'Havel")
	i.Liste_boucliers = append(i.Liste_boucliers, bouclier_bois, Bouclier_bois_cerf, Bouclier_fer, Bouclier_acier, Bouclier_mithril, Bouclier_Havel)

}
