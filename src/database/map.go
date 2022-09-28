package database

type Arbre struct {
	Val    map[string]string
	Gauche *Arbre
	Droite *Arbre
	Centre *Arbre
	Parent *Arbre
}

// number of key in the val : ("name", "mob_nb", "mob_type", "boss_type", "feu", "forge", "secret_destination", "secret_unlock")
func (a *Arbre) Insert(val map[string]string, parent *Arbre) {
	if a.Val == nil {
		a.Val = val
		a.Gauche = &Arbre{}
		a.Droite = &Arbre{}
		a.Centre = &Arbre{}
		a.Parent = &Arbre{}
	}
}

func (hub *Arbre) Init() {
	hub.Insert(map[string]string{"name": "Lige feu", "mob_nb": "0", "mob_type": "none", "feu": "true", "forge": "true", "loot": "0"}, nil)
	//camps des morts vivants
	hub.Centre.Insert(map[string]string{"name": "Camps des morts vivants", "mob_nb": "4", "mob_type": "Carcasse", "feu": "false", "forge": "false", "loot": "0"}, hub)
	//partie gauche
	hub.Centre.Gauche.Insert(map[string]string{"name": "Batiment détruit", "mob_nb": "0", "mob_type": "none", "feu": "true", "forge": "false", "loot": "0"}, hub.Centre)
	hub.Centre.Gauche.Centre.Insert(map[string]string{"name": "Tour ouest", "mob_nb": "2", "mob_type": "Chevalier mort-vivant", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Gauche)
	hub.Centre.Gauche.Centre.Gauche.Insert(map[string]string{"name": "Muraille", "mob_nb": "1", "mob_type": "Démon taureau", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Gauche.Centre)
	hub.Centre.Gauche.Centre.Gauche.Droite.Insert(map[string]string{"name": "Tour nord", "mob_nb": "0", "mob_type": "none", "feu": "true", "forge": "false", "loot": "0"}, hub.Centre.Gauche.Centre.Gauche)
	hub.Centre.Gauche.Centre.Gauche.Droite.Centre.Insert(map[string]string{"name": "Anor Londo", "mob_nb": "0", "mob_type": "none", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Gauche.Centre.Gauche.Droite)
	//partie droite
	hub.Centre.Droite.Insert(map[string]string{"name": "Forteresse", "mob_nb": "0", "mob_type": "none", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre)
	hub.Centre.Droite.Centre.Insert(map[string]string{"name": "Annexe", "mob_nb": "3", "mob_type": "Carcasse", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite)
	hub.Centre.Droite.Centre.Gauche.Insert(map[string]string{"name": "Forge", "mob_nb": "0", "mob_type": "none", "feu": "true", "forge": "true", "loot": "0"}, hub.Centre.Droite.Centre)
	hub.Centre.Droite.Centre.Droite.Insert(map[string]string{"name": "Couloir", "mob_nb": "0", "mob_type": "none", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre)
	hub.Centre.Droite.Centre.Droite.Gauche.Insert(map[string]string{"name": "Cuisine", "mob_nb": "2", "mob_type": "Champion mort-vivant", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Insert(map[string]string{"name": "Garde manger", "mob_nb": "0", "mob_type": "none", "feu": "true", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche)
	//cours inétieure
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Insert(map[string]string{"name": "Cours intérieure", "mob_nb": "0", "mob_type": "none", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite)
	//partie gauche
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Insert(map[string]string{"name": "Chambre du prince", "mob_nb": "0", "mob_type": "none", "feu": "false", "forge": "false", "loot": "1"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Insert(map[string]string{"name": "Balcon de la chambre", "mob_nb": "0", "mob_type": "none", "feu": "true", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Insert(map[string]string{"name": "Échelle vers le toit", "mob_nb": "0", "mob_type": "none", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Droite.Insert(map[string]string{"name": "Toit de la chambre", "mob_nb": "2", "mob_type": "Champion mort-vivant", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche)
	//partie droite
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Droite.Insert(map[string]string{"name": "Réserve d'arme", "mob_nb": "3", "mob_type": "Champion mort-vivant", "feu": "false", "forge": "false", "loot": "1"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Droite.Gauche.Insert(map[string]string{"name": "Tour de guet intérieur", "mob_nb": "2", "mob_type": "Carcasse", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Droite)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Droite.Gauche.Centre.Insert(map[string]string{"name": "Tour de guet extérieur", "mob_nb": "0", "mob_type": "none", "feu": "false", "forge": "false", "loot": "2"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Droite.Gauche)
	//Toit de la chambre
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Droite.Gauche.Insert(map[string]string{"name": "Toit du garde manger", "mob_nb": "0", "mob_type": "none", "feu": "false", "forge": "false", "loot": "3"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Droite.Droite.Insert(map[string]string{"name": "Toit de la réserve d'arme", "mob_nb": "7", "mob_type": "Champion mort-vivant", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Droite.Centre.Insert(map[string]string{"name": "Toit du palais", "mob_nb": "0", "mob_type": "none", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre)
	//Toit du palais
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Droite.Centre.Centre.Insert(map[string]string{"name": "Toit du clocher", "mob_nb": "1", "mob_type": "Gargouille", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Droite.Centre)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Droite.Centre.Centre.Gauche.Insert(map[string]string{"name": "Tour du clocher", "mob_nb": "0", "mob_type": "none", "feu": "false", "forge": "false", "loot": "4"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Droite.Centre.Centre)

}

func (p *Personnage) Deplacement(direction string) bool {
	restart := false
	if direction == "centre" {
		if p.Position.Centre.Val != nil {
			p.Back_list = append(p.Back_list, p.Position)
			p.Position = *p.Position.Centre
			if p.Do_combat() {
				p.Mana = p.ManaMax
				Affichage("Déplacement", []string{"félicitation nous sommes désormais arrivé à " + p.Position.Val["name"]})
				Attendre()
			} else {
				p.Position = p.Back_list[len(p.Back_list)-1]
				p.Back_list = p.Back_list[:len(p.Back_list)-1]
			}
		} else {
			Affichage("Déplacement", []string{"On dirait que nous nous sommes perdu...", "Il n'y a rien ici.", "Il faut trouver un autre chemin."})
			Attendre()
			restart = true
		}
	} else if direction == "gauche" {
		if p.Position.Gauche.Val != nil {
			p.Back_list = append(p.Back_list, p.Position)
			p.Position = *p.Position.Gauche
			if p.Do_combat() {
				Affichage("Déplacement", []string{"félicitation nous sommes désormais arrivé à " + p.Position.Val["name"]})
				Attendre()
			} else {
				p.Position = p.Back_list[len(p.Back_list)-1]
				p.Back_list = p.Back_list[:len(p.Back_list)-1]
			}
		} else {
			Affichage("Déplacement", []string{"On dirait que nous nous sommes perdu...", "Il n'y a rien ici.", "Il faut trouver un autre chemin."})
			Attendre()
			restart = true
		}
	} else if direction == "droite" {
		if p.Position.Droite.Val != nil {
			p.Back_list = append(p.Back_list, p.Position)
			p.Position = *p.Position.Droite
			if p.Do_combat() {
				Affichage("Déplacement", []string{"félicitation nous sommes désormais arrivé à " + p.Position.Val["name"]})
				Attendre()
			} else {
				p.Position = p.Back_list[len(p.Back_list)-1]
				p.Back_list = p.Back_list[:len(p.Back_list)-1]
			}
		} else {
			Affichage("Déplacement", []string{"On dirait que nous nous sommes perdu...", "Il n'y a rien ici.", "Il faut trouver un autre chemin."})
			Attendre()
			restart = true
		}
	} else if direction == "retour" {
		if len(p.Back_list) > 0 {
			p.Position = p.Back_list[len(p.Back_list)-1]
			p.Back_list = p.Back_list[:len(p.Back_list)-1]
			Affichage("Déplacement", []string{"félicitation nous sommes désormais arrivé à " + p.Position.Val["name"]})
			Attendre()
		} else {
			Affichage("Déplacement", []string{"On dirait que nous nous sommes perdu...", "Il n'y a rien ici.", "Il faut trouver un autre chemin."})
			Attendre()
			restart = true
		}
	} else if direction == "hub" {
		var base Arbre
		base.Init()
		p.Position = base
		p.Back_list = []Arbre{}
		if p.Ames > 20 {
			p.Ames -= 20
			Affichage("Déplacement", []string{"félicitation nous sommes désormais arrivé au hub."})
			Attendre()
		} else {
			Affichage("Déplacement", []string{"On dirait que nous n'avons pas assez d'âmes pour revenir au hub."})
			Attendre()
			restart = true
		}
	}
	return restart
}
