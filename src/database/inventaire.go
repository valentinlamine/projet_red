package database

type Inventaire struct {
	Liste_armes          []Armes
	Liste_armures_tete   []Armures
	Liste_armures_torse  []Armures
	Liste_armures_bras   []Armures
	Liste_armures_jambes []Armures
	Liste_boucliers      []Boucliers
	Liste_consommables   []Consommable
	Liste_sort           []Sort
	Liste_items          map[string]int
}

func (i *Inventaire) Init() {
	//Initialisation des armes
	for j := 1; j < 7; j++ {
		var a Armes
		a.Init_Armes(j)
		i.Liste_armes = append(i.Liste_armes, a)
	}
	//Initialisation des boucliers
	for j := 1; j < 7; j++ {
		var b Boucliers
		b.Init_Bouclier(j)
		i.Liste_boucliers = append(i.Liste_boucliers, b)
	}
	//Initialisation des consommables
	for j := 1; j < 7; j++ {
		var c Consommable
		c.Init_Consommable(j)
		i.Liste_consommables = append(i.Liste_consommables, c)
	}
	//Initialisation des armures
	for j := 1; j < 25; j++ {
		var a Armures
		a.Init_Armures(j)
		if j < 7 {
			i.Liste_armures_tete = append(i.Liste_armures_tete, a)
		} else if j < 13 {
			i.Liste_armures_torse = append(i.Liste_armures_torse, a)
		} else if j < 19 {
			i.Liste_armures_bras = append(i.Liste_armures_bras, a)
		} else {
			i.Liste_armures_jambes = append(i.Liste_armures_jambes, a)
		}
	}
	//initialisation des sorts
	for j := 1; j < 7; j++ {
		var s Sort
		s.InitSort(j)
		i.Liste_sort = append(i.Liste_sort, s)
	}
	//initialisation des item
	i.Liste_items = make(map[string]int)
	i.Liste_items["éclat de titanite"] = 0
	i.Liste_items["grand éclat de titanite"] = 0
	i.Liste_items["tablette de titanite"] = 0
}

func (p *Personnage) Equiper(item interface{}, annonce bool) {
	//TODO : vérifier si l'item est déjà équipé
	a, b := "", []string{""}
	isconso := false
	switch item := item.(type) {
	case Armes:
		if p.PoidsMax < p.PoidsEquip+item.Poids {
			a, b[0] = "Erreur", "Vous ne pouvez pas porter autant de poids"
		} else if p.Force < item.LvlMinFor && p.Dexterite < item.LvlMinDex && p.Intelligence < item.LvlMinInt {
			a, b[0] = "Erreur", "Vous n'avez pas le niveau requis pour équiper cet objet"
		} else if p.EquipementArmes == item {
			a, b[0] = "Erreur", "Vous avez déjà équipé cet objet"
		} else {
			p.PoidsEquip -= p.EquipementArmes.Poids
			p.Degats -= p.EquipementArmes.Deg
			p.EquipementArmes = item
			p.PoidsEquip += item.Poids
			p.Degats += item.Deg
			a, b[0] = "Succès", "Vous avez équipé "+item.Nom
		}
	case Boucliers:
		if p.PoidsMax < p.PoidsEquip+item.Poids {
			a, b[0] = "Erreur", "Vous ne pouvez pas porter autant de poids"
		} else if p.Force < item.LvlMinFor && p.Dexterite < item.LvlMinDex && p.Intelligence < item.LvlMinInt {
			a, b[0] = "Erreur", "Vous n'avez pas le niveau requis pour équiper cet objet"
		} else if p.EquipementBoucliers == item {
			a, b[0] = "Erreur", "Vous avez déjà équipé cet objet"
		} else {
			p.PoidsEquip -= p.EquipementBoucliers.Poids
			p.Pvmax -= p.EquipementBoucliers.Pvbonus
			p.EquipementBoucliers = item
			p.PoidsEquip += item.Poids
			p.Pvmax += item.Pvbonus
			a, b[0] = "Succès", "Vous avez équipé "+item.Nom
		}
	case Armures:
		for _, v := range p.EquipementArmures {
			if v == item {
				a, b[0] = "Erreur", "Vous avez déjà équipé cet objet"
				return
			}
		}
		if p.PoidsMax < p.PoidsEquip+item.Poids {
			a, b[0] = "Erreur", "Vous ne pouvez pas porter autant de poids"
		} else if p.Force < item.LvlMinFor && p.Dexterite < item.LvlMinDex && p.Intelligence < item.LvlMinInt {
			a, b[0] = "Erreur", "Vous n'avez pas le niveau requis pour équiper cet objet"
		} else {
			switch item.Class {
			case "casque":
				p.PoidsEquip -= p.EquipementArmures["tete"].Poids
				p.Pvmax -= p.EquipementArmures["tete"].Pvbonus
				p.EquipementArmures["tete"] = item
			case "plastron":
				p.PoidsEquip -= p.EquipementArmures["torse"].Poids
				p.Pvmax -= p.EquipementArmures["torse"].Pvbonus
				p.EquipementArmures["torse"] = item
			case "gantelet":
				p.PoidsEquip -= p.EquipementArmures["bras"].Poids
				p.Pvmax -= p.EquipementArmures["bras"].Pvbonus
				p.EquipementArmures["bras"] = item
			case "jambieres":
				p.PoidsEquip -= p.EquipementArmures["jambes"].Poids
				p.Pvmax -= p.EquipementArmures["jambes"].Pvbonus
				p.EquipementArmures["jambes"] = item
			default:
				a, b[0] = "Erreur", "Erreur lors de l'équipement"
			}
			p.PoidsEquip += item.Poids
			p.Pvmax += item.Pvbonus
			a, b[0] = "Succès", "Vous avez équipé "+item.Nom
		}
	case Consommable:
		p.PrendrePot(item)
		isconso = true
	}
	if annonce && !isconso {
		Affichage(a, b)
	}
}

func (p *Personnage) Remplacer_item(item interface{}) {
	switch item := item.(type) {
	case Armes:
		for i, v := range p.Inv.Liste_armes {
			if v.Nom == item.Nom {
				p.Inv.Liste_armes[i] = item
				if p.EquipementArmes.Nom == item.Nom {
					p.Equiper(item, false)
				}
			}
		}
	case Boucliers:
		for i, v := range p.Inv.Liste_boucliers {
			if v.Nom == item.Nom {
				p.Inv.Liste_boucliers[i] = item
				if p.EquipementBoucliers.Nom == item.Nom {
					p.Equiper(item, false)
				}
			}
		}
	case Armures:
		switch item.Class {
		case "casque":
			for i, v := range p.Inv.Liste_armures_tete {
				if v.Nom == item.Nom {
					p.Inv.Liste_armures_tete[i] = item
					if p.EquipementArmures["tete"].Nom == item.Nom {
						p.Equiper(item, false)
					}
				}
			}
		case "plastron":
			for i, v := range p.Inv.Liste_armures_torse {
				if v.Nom == item.Nom {
					p.Inv.Liste_armures_torse[i] = item
					if p.EquipementArmures["torse"].Nom == item.Nom {
						p.Equiper(item, false)
					}
				}
			}
		case "gantelet":
			for i, v := range p.Inv.Liste_armures_bras {
				if v.Nom == item.Nom {
					p.Inv.Liste_armures_bras[i] = item
					if p.EquipementArmures["bras"].Nom == item.Nom {
						p.Equiper(item, false)
					}
				}
			}
		case "jambieres":
			for i, v := range p.Inv.Liste_armures_jambes {
				if v.Nom == item.Nom {
					p.Inv.Liste_armures_jambes[i] = item
					if p.EquipementArmures["jambes"].Nom == item.Nom {
						p.Equiper(item, false)
					}
				}
			}
		}
	}
}
