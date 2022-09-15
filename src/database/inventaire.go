package databse

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
	Liste_sort         []Sort
}

func (i *Inventaire) Init() {
	i.Liste_armes = append(i.Liste_armes, Armes.Init("Dague"))
	i.Liste_armes = append(i.Liste_armes, Armes.Init("Claymore"))
	i.Liste_armes = append(i.Liste_armes, Armes.Init("Rapière"))
	i.Liste_armes = append(i.Liste_armes, Armes.Init("Uchigatana"))
	i.Liste_armes = append(i.Liste_armes, Armes.Init("Bâton"))
	i.Liste_armes = append(i.Liste_armes, Armes.Init("Hache queue de gargouille"))

}
