package database

type Inventaire struct {
	Liste_Armes          []Armes
	Liste_Armures_Tete   []Armures
	Liste_Armures_Torse  []Armures
	Liste_Armures_Bras   []Armures
	Liste_Armures_Jambes []Armures
	Liste_Boucliers      []Boucliers
	Liste_Consommables   []Consommable
	Liste_Sorts          []Sort
	Liste_Items          map[string]int
}

func (i *Inventaire) Initialisation() {
	//Initialisation des armes
	for j := 1; j < 7; j++ {
		var a Armes
		a.Initialisation_Armes(j)
		i.Liste_Armes = append(i.Liste_Armes, a)
	}
	//Initialisation des boucliers
	for j := 1; j < 7; j++ {
		var b Boucliers
		b.Initialisation_Bouclier(j)
		i.Liste_Boucliers = append(i.Liste_Boucliers, b)
	}
	//Initialisation des consommables
	for j := 1; j < 7; j++ {
		var c Consommable
		c.Initialisation_Consommable(j)
		i.Liste_Consommables = append(i.Liste_Consommables, c)
	}
	//Initialisation des armures
	for j := 1; j < 25; j++ {
		var a Armures
		a.Initialisation_Armures(j)
		if j < 7 {
			i.Liste_Armures_Tete = append(i.Liste_Armures_Tete, a)
		} else if j < 13 {
			i.Liste_Armures_Torse = append(i.Liste_Armures_Torse, a)
		} else if j < 19 {
			i.Liste_Armures_Bras = append(i.Liste_Armures_Bras, a)
		} else {
			i.Liste_Armures_Jambes = append(i.Liste_Armures_Jambes, a)
		}
	}
	//initialisation des sorts
	for j := 1; j < 7; j++ {
		var s Sort
		s.Initialisation_Sort(j)
		i.Liste_Sorts = append(i.Liste_Sorts, s)
	}
	//initialisation des item
	i.Liste_Items = make(map[string]int)
	i.Liste_Items["éclat de titanite"] = 0
	i.Liste_Items["grand éclat de titanite"] = 0
	i.Liste_Items["tablette de titanite"] = 0
}

func (p *Personnage) Equiper(item interface{}, annonce bool) {
	a, b := "", []string{""}
	EstConsommable := false
	switch item := item.(type) {
	case Armes:
		if p.PoidsMax < p.PoidsAct+item.Poids {
			a, b[0] = "Erreur", "Vous ne pouvez pas porter autant de poids"
		} else if p.Force < item.NivMinFor && p.Dexterite < item.NivMinDex && p.Intelligence < item.NivMinInt {
			a, b[0] = "Erreur", "Vous n'avez pas le niveau requis pour équiper cet objet"
		} else if p.EquipementArmes == item {
			a, b[0] = "Erreur", "Vous avez déjà équipé cet objet"
		} else {
			p.PoidsAct -= p.EquipementArmes.Poids
			p.Degats -= p.EquipementArmes.Degats
			p.EquipementArmes = item
			p.PoidsAct += item.Poids
			p.Degats += item.Degats
			a, b[0] = "Succès", "Vous avez équipé "+item.Nom
		}
	case Boucliers:
		if p.PoidsMax < p.PoidsAct+item.Poids {
			a, b[0] = "Erreur", "Vous ne pouvez pas porter autant de poids"
		} else if p.Force < item.NivMinFor && p.Dexterite < item.NivMinDex && p.Intelligence < item.NivMinInt {
			a, b[0] = "Erreur", "Vous n'avez pas le niveau requis pour équiper cet objet"
		} else if p.EquipementBoucliers == item {
			a, b[0] = "Erreur", "Vous avez déjà équipé cet objet"
		} else {
			p.PoidsAct -= p.EquipementBoucliers.Poids
			p.VieMax -= p.EquipementBoucliers.VieBonus
			p.EquipementBoucliers = item
			p.PoidsAct += item.Poids
			p.VieMax += item.VieBonus
			a, b[0] = "Succès", "Vous avez équipé "+item.Nom
		}
	case Armures:
		for _, v := range p.EquipementArmures {
			if v == item {
				a, b[0] = "Erreur", "Vous avez déjà équipé cet objet"
				return
			}
		}
		if p.PoidsMax < p.PoidsAct+item.Poids {
			a, b[0] = "Erreur", "Vous ne pouvez pas porter autant de poids"
		} else if p.Force < item.NivMinFor && p.Dexterite < item.NivMinDex && p.Intelligence < item.NivMinInt {
			a, b[0] = "Erreur", "Vous n'avez pas le niveau requis pour équiper cet objet"
		} else {
			switch item.Class {
			case "casque":
				p.PoidsAct -= p.EquipementArmures["tete"].Poids
				p.VieMax -= p.EquipementArmures["tete"].VieBonus
				p.EquipementArmures["tete"] = item
			case "plastron":
				p.PoidsAct -= p.EquipementArmures["torse"].Poids
				p.VieMax -= p.EquipementArmures["torse"].VieBonus
				p.EquipementArmures["torse"] = item
			case "gantelet":
				p.PoidsAct -= p.EquipementArmures["bras"].Poids
				p.VieMax -= p.EquipementArmures["bras"].VieBonus
				p.EquipementArmures["bras"] = item
			case "jambieres":
				p.PoidsAct -= p.EquipementArmures["jambes"].Poids
				p.VieMax -= p.EquipementArmures["jambes"].VieBonus
				p.EquipementArmures["jambes"] = item
			default:
				a, b[0] = "Erreur", "Erreur lors de l'équipement"
			}
			p.PoidsAct += item.Poids
			p.VieMax += item.VieBonus
			a, b[0] = "Succès", "Vous avez équipé "+item.Nom
		}
	case Consommable:
		p.Prendre_Potion(item)
		EstConsommable = true
	}
	if annonce && !EstConsommable {
		Affichage(a, b, true, true)
	}
}

func (p *Personnage) Remplacer_Item(item interface{}) {
	switch item := item.(type) {
	case Armes:
		for i, v := range p.Inv.Liste_Armes {
			if v.Nom == item.Nom {
				p.Inv.Liste_Armes[i] = item
				if p.EquipementArmes.Nom == item.Nom {
					p.Equiper(item, false)
				}
			}
		}
	case Boucliers:
		for i, v := range p.Inv.Liste_Boucliers {
			if v.Nom == item.Nom {
				p.Inv.Liste_Boucliers[i] = item
				if p.EquipementBoucliers.Nom == item.Nom {
					p.Equiper(item, false)
				}
			}
		}
	case Armures:
		switch item.Class {
		case "casque":
			for i, v := range p.Inv.Liste_Armures_Tete {
				if v.Nom == item.Nom {
					p.Inv.Liste_Armures_Tete[i] = item
					if p.EquipementArmures["tete"].Nom == item.Nom {
						p.Equiper(item, false)
					}
				}
			}
		case "plastron":
			for i, v := range p.Inv.Liste_Armures_Torse {
				if v.Nom == item.Nom {
					p.Inv.Liste_Armures_Torse[i] = item
					if p.EquipementArmures["torse"].Nom == item.Nom {
						p.Equiper(item, false)
					}
				}
			}
		case "gantelet":
			for i, v := range p.Inv.Liste_Armures_Bras {
				if v.Nom == item.Nom {
					p.Inv.Liste_Armures_Bras[i] = item
					if p.EquipementArmures["bras"].Nom == item.Nom {
						p.Equiper(item, false)
					}
				}
			}
		case "jambieres":
			for i, v := range p.Inv.Liste_Armures_Jambes {
				if v.Nom == item.Nom {
					p.Inv.Liste_Armures_Jambes[i] = item
					if p.EquipementArmures["jambes"].Nom == item.Nom {
						p.Equiper(item, false)
					}
				}
			}
		}
	}
}
