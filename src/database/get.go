package database

func (i *Inventaire) Get_Item(nom string) interface{} {
	for _, v := range i.Liste_armes {
		if v.Nom == nom {
			return v
		}
	}
	for _, v := range i.Liste_boucliers {
		if v.Nom == nom {
			return v
		}
	}
	for _, v := range i.Liste_consommables {
		if v.Nom == nom {
			return v
		}
	}
	for _, v := range i.Liste_armures_tete {
		if v.Nom == nom {
			return v
		}
	}
	for _, v := range i.Liste_armures_torse {
		if v.Nom == nom {
			return v
		}
	}
	for _, v := range i.Liste_armures_jambes {
		if v.Nom == nom {
			return v
		}
	}
	for _, v := range i.Liste_armures_bras {
		if v.Nom == nom {
			return v
		}
	}
	for _, v := range i.Liste_sort {
		if v.Nom == nom {
			return v
		}
	}
	return nil
}
