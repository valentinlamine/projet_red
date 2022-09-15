package database

//Liste des variables des armes
var dague Armes
var claymore Armes
var rapière Armes
var uchigatana Armes
var bâton Armes
var hachequeue Armes

//Liste des variables des boucliers

type Inventaire struct {
	Liste_armes        []Armes
	Liste_consommables []Consommable
	Liste_boucliers    []Boucliers
	Liste_sort         []string //TODO Sort
}

func (i *Inventaire) Init() {
	dague.Init_Armes("Dague")
	claymore.Init_Armes("Claymore")
	rapière.Init_Armes("Rapière")
	uchigatana.Init_Armes("Uchigatana")
	bâton.Init_Armes("Bâton")
	hachequeue.Init_Armes("Hache queue de gargouille")
	i.Liste_armes = append(i.Liste_armes, dague)
	i.Liste_armes = append(i.Liste_armes, claymore)
	i.Liste_armes = append(i.Liste_armes, rapière)
	i.Liste_armes = append(i.Liste_armes, uchigatana)
	i.Liste_armes = append(i.Liste_armes, bâton)
	i.Liste_armes = append(i.Liste_armes, hachequeue)

}
