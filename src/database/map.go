package database

type Arbre struct {
	Val    map[string]string
	Gauche *Arbre
	Droite *Arbre
	Centre *Arbre
}

// number of key in the val : ("name", "mob_nb", "mob_type", "boss_type", "feu", "forge", "secret_destination", "secret_unlock")
func (a *Arbre) Insert(val map[string]string) {
	if a.Val == nil {
		a.Val = val
		a.Gauche = &Arbre{}
		a.Droite = &Arbre{}
		a.Centre = &Arbre{}
	}
}

func (hub *Arbre) Init() {
	hub.Insert(map[string]string{"name": "hub", "mob_nb": "0", "mob_type": "none", "boss_type": "none", "feu": "vrai", "forge": "false", "secret_destination": "none", "secret_unlock": "false"})
	hub.Centre.Insert(map[string]string{"name": "Camps de mort vivant", "mob_nb": "4", "mob_type": "none", "boss_type": "none", "feu": "false", "forge": "false", "secret_destination": "none", "secret_unlock": "false"})
	hub.Centre.Droite.Insert(map[string]string{"name": "Forteresse", "mob_nb": "2", "mob_type": "none", "boss_type": "none", "feu": "false", "forge": "false", "secret_destination": "none", "secret_unlock": "false"})
	hub.Centre.Gauche.Insert(map[string]string{"name": "Batiment détruit", "mob_nb": "0", "mob_type": "none", "boss_type": "none", "feu": "true", "forge": "false", "secret_destination": "none", "secret_unlock": "false"})
	hub.Centre.Gauche.Centre.Insert(map[string]string{"name": "Tour ouest", "mob_nb": "6", "mob_type": "none", "boss_type": "none", "feu": "false", "forge": "false", "secret_destination": "none", "secret_unlock": "false"})
	hub.Centre.Gauche.Centre.Centre.Insert(map[string]string{"name": "Muraille", "mob_nb": "0", "mob_type": "none", "boss_type": "Gargouille", "feu": "false", "forge": "false", "secret_destination": "none", "secret_unlock": "false"})
	hub.Centre.Gauche.Centre.Centre.Centre.Insert(map[string]string{"name": "Tour nord", "mob_nb": "0", "mob_type": "none", "boss_type": "none", "feu": "false", "forge": "false", "secret_destination": "none", "secret_unlock": "false"})
}

func (p *Personnage) Deplacement(base Arbre, direction string) bool {
	restart := false
	if direction == "centre" {
		if p.Position.Centre.Val != nil {
			p.Position = *p.Position.Centre
			Affichage("Déplacement", []string{"félicitation nous sommes désormais arrivé à " + p.Position.Val["name"]})
			Attendre()
		} else {
			Affichage("Déplacement", []string{"On dirait que nous nous sommes perdu...", "Il n'y a rien ici.", "Il faut trouver un autre chemin."})
			Attendre()
			restart = true
		}
	} else if direction == "gauche" {
		if p.Position.Gauche.Val != nil {
			p.Position = *p.Position.Gauche
			Affichage("Déplacement", []string{"félicitation nous sommes désormais arrivé à " + p.Position.Val["name"]})
			Attendre()
		} else {
			Affichage("Déplacement", []string{"On dirait que nous nous sommes perdu...", "Il n'y a rien ici.", "Il faut trouver un autre chemin."})
			Attendre()
			restart = true
		}
	} else if direction == "droite" {
		if p.Position.Droite.Val != nil {
			p.Position = *p.Position.Droite
			Affichage("Déplacement", []string{"félicitation nous sommes désormais arrivé à " + p.Position.Val["name"]})
			Attendre()
		} else {
			Affichage("Déplacement", []string{"On dirait que nous nous sommes perdu...", "Il n'y a rien ici.", "Il faut trouver un autre chemin."})
			Attendre()
			restart = true
		}
	} else if direction == "retour" {
		if p.Position.Val["name"] == "hub" {
			Affichage("Déplacement", []string{"Nous sommes déjà au hub."})
			Attendre()
			restart = true
		} else {
			p.Position = base
			Affichage("Déplacement", []string{"félicitation nous sommes désormais arrivé au hub."})
			Attendre()
		}
	}
	return restart
}
