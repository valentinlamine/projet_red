package databse



type Inventaire struct {
	Armes        []Armes
	Consommables []Consommable
	Boucliers    []Boucliers
	Sort 	     []Sort
}

func (i *Inventaire) Init() {
	i.Armes 
