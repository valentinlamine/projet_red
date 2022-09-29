package database

func (i *Inventaire) Recuperer_Item(nom string) interface{} {
	for _, v := range i.Liste_Armes {
		if v.Nom == nom {
			return v
		}
	}
	for _, v := range i.Liste_Boucliers {
		if v.Nom == nom {
			return v
		}
	}
	for _, v := range i.Liste_Consommables {
		if v.Nom == nom {
			return v
		}
	}
	for _, v := range i.Liste_Armures_Tete {
		if v.Nom == nom {
			return v
		}
	}
	for _, v := range i.Liste_Armures_Torse {
		if v.Nom == nom {
			return v
		}
	}
	for _, v := range i.Liste_Armures_Jambes {
		if v.Nom == nom {
			return v
		}
	}
	for _, v := range i.Liste_Armures_Bras {
		if v.Nom == nom {
			return v
		}
	}
	for _, v := range i.Liste_Sorts {
		if v.Nom == nom {
			return v
		}
	}
	return nil
}
