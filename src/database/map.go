package database

type Arbre struct {
	Val    map[string]string
	Gauche *Arbre
	Droite *Arbre
	Centre *Arbre
	Parent *Arbre
}

// nombre of key in the val : ("nom", "monstre_nombre", "monstre_type", "boss_type", "feu", "forge", "secret_destination", "secret_unlock")
func (a *Arbre) Inserer(val map[string]string, parent *Arbre) {
	if a.Val == nil {
		a.Val = val
		a.Gauche = &Arbre{}
		a.Droite = &Arbre{}
		a.Centre = &Arbre{}
		a.Parent = &Arbre{}
	}
}

func (hub *Arbre) Initialisation() {
	hub.Inserer(map[string]string{"nom": "Lige feu", "monstre_nombre": "0", "monstre_type": "none", "feu": "true", "forge": "true", "loot": "0"}, nil)
	//camps des morts vivants
	hub.Centre.Inserer(map[string]string{"nom": "Camps des morts vivants", "monstre_nombre": "4", "monstre_type": "Carcasse", "feu": "false", "forge": "false", "loot": "0"}, hub)
	//partie gauche
	hub.Centre.Gauche.Inserer(map[string]string{"nom": "Batiment détruit", "monstre_nombre": "0", "monstre_type": "none", "feu": "true", "forge": "false", "loot": "0"}, hub.Centre)
	hub.Centre.Gauche.Centre.Inserer(map[string]string{"nom": "Tour ouest", "monstre_nombre": "2", "monstre_type": "Chevalier mort-vivant", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Gauche)
	hub.Centre.Gauche.Centre.Gauche.Inserer(map[string]string{"nom": "Muraille", "monstre_nombre": "1", "monstre_type": "Démon taureau", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Gauche.Centre)
	hub.Centre.Gauche.Centre.Gauche.Droite.Inserer(map[string]string{"nom": "Tour nord", "monstre_nombre": "0", "monstre_type": "none", "feu": "true", "forge": "false", "loot": "0"}, hub.Centre.Gauche.Centre.Gauche)
	hub.Centre.Gauche.Centre.Gauche.Droite.Centre.Inserer(map[string]string{"nom": "Anor Londo", "monstre_nombre": "0", "monstre_type": "none", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Gauche.Centre.Gauche.Droite)
	//partie droite
	hub.Centre.Droite.Inserer(map[string]string{"nom": "Forteresse", "monstre_nombre": "0", "monstre_type": "none", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre)
	hub.Centre.Droite.Centre.Inserer(map[string]string{"nom": "Annexe", "monstre_nombre": "3", "monstre_type": "Carcasse", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite)
	hub.Centre.Droite.Centre.Gauche.Inserer(map[string]string{"nom": "Forge", "monstre_nombre": "0", "monstre_type": "none", "feu": "true", "forge": "true", "loot": "0"}, hub.Centre.Droite.Centre)
	hub.Centre.Droite.Centre.Droite.Inserer(map[string]string{"nom": "Couloir", "monstre_nombre": "0", "monstre_type": "none", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre)
	hub.Centre.Droite.Centre.Droite.Gauche.Inserer(map[string]string{"nom": "Cuisine", "monstre_nombre": "2", "monstre_type": "Champion mort-vivant", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Inserer(map[string]string{"nom": "Garde manger", "monstre_nombre": "0", "monstre_type": "none", "feu": "true", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche)
	//cours intérieur
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Inserer(map[string]string{"nom": "Cours intérieure", "monstre_nombre": "0", "monstre_type": "none", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite)
	//partie gauche
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Inserer(map[string]string{"nom": "Chambre du prince", "monstre_nombre": "0", "monstre_type": "none", "feu": "false", "forge": "false", "loot": "1"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Inserer(map[string]string{"nom": "Balcon de la chambre", "monstre_nombre": "0", "monstre_type": "none", "feu": "true", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Inserer(map[string]string{"nom": "Échelle vers le toit", "monstre_nombre": "0", "monstre_type": "none", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Droite.Inserer(map[string]string{"nom": "Toit de la chambre", "monstre_nombre": "2", "monstre_type": "Champion mort-vivant", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche)
	//partie droite
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Droite.Inserer(map[string]string{"nom": "Réserve d'armes", "monstre_nombre": "3", "monstre_type": "Champion mort-vivant", "feu": "false", "forge": "false", "loot": "1"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Droite.Gauche.Inserer(map[string]string{"nom": "Tour de guet intérieur", "monstre_nombre": "2", "monstre_type": "Carcasse", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Droite)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Droite.Gauche.Centre.Inserer(map[string]string{"nom": "Tour de guet extérieur", "monstre_nombre": "0", "monstre_type": "none", "feu": "false", "forge": "false", "loot": "2"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Droite.Gauche)
	//Toit de la chambre
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Droite.Gauche.Inserer(map[string]string{"nom": "Toit du garde manger", "monstre_nombre": "0", "monstre_type": "none", "feu": "false", "forge": "false", "loot": "3"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Droite.Droite.Inserer(map[string]string{"nom": "Toit de la réserve d'armes", "monstre_nombre": "7", "monstre_type": "Champion mort-vivant", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Droite.Centre.Inserer(map[string]string{"nom": "Toit du palais", "monstre_nombre": "0", "monstre_type": "none", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre)
	//Toit du palais
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Droite.Centre.Centre.Inserer(map[string]string{"nom": "Toit du clocher", "monstre_nombre": "1", "monstre_type": "Gargouille", "feu": "false", "forge": "false", "loot": "0"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Droite.Centre)
	hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Droite.Centre.Centre.Gauche.Inserer(map[string]string{"nom": "Tour du clocher", "monstre_nombre": "0", "monstre_type": "none", "feu": "false", "forge": "false", "loot": "4"}, hub.Centre.Droite.Centre.Droite.Gauche.Droite.Centre.Gauche.Gauche.Centre.Droite.Centre.Centre)

}

func (p *Personnage) Deplacement(direction string) bool {
	redemarrage := false
	if direction == "centre" {
		if p.Position.Centre.Val != nil {
			p.ListeRetour = append(p.ListeRetour, p.Position)
			p.Position = *p.Position.Centre
			if p.Faire_Combat() {
				p.ManaAct = p.ManaMax
				Affichage("Déplacement", []string{"félicitation nous sommes désormais arrivé à " + p.Position.Val["nom"]}, true, true)
			} else {
				if p.Ames != 0 && p.ManaAct != p.ManaMax && p.VieAct != p.VieMax && p.Position.Val["nom"] != "Lige feu" { //condition de post mortem
					p.Position = p.ListeRetour[len(p.ListeRetour)-1]
					p.ListeRetour = p.ListeRetour[:len(p.ListeRetour)-1]
				}
			}
		} else {
			Affichage("Déplacement", []string{"On dirait que nous nous sommes perdu...", "Il n'y a rien ici.", "Il faut trouver un autre chemin."}, true, true)
			redemarrage = true
		}
	} else if direction == "gauche" {
		if p.Position.Gauche.Val != nil {
			p.ListeRetour = append(p.ListeRetour, p.Position)
			p.Position = *p.Position.Gauche
			if p.Faire_Combat() {
				Affichage("Déplacement", []string{"félicitation nous sommes désormais arrivé à " + p.Position.Val["nom"]}, true, true)
			} else {
				if p.Ames != 0 && p.ManaAct != p.ManaMax && p.VieAct != p.VieMax && p.Position.Val["nom"] != "Lige feu" { //condition de post mortem
					p.Position = p.ListeRetour[len(p.ListeRetour)-1]
					p.ListeRetour = p.ListeRetour[:len(p.ListeRetour)-1]
				}
			}
		} else {
			Affichage("Déplacement", []string{"On dirait que nous nous sommes perdu...", "Il n'y a rien ici.", "Il faut trouver un autre chemin."}, true, true)
			redemarrage = true
		}
	} else if direction == "droite" {
		if p.Position.Droite.Val != nil {
			p.ListeRetour = append(p.ListeRetour, p.Position)
			p.Position = *p.Position.Droite
			if p.Faire_Combat() {
				Affichage("Déplacement", []string{"félicitation nous sommes désormais arrivé à " + p.Position.Val["nom"]}, true, true)
			} else {
				if p.Ames != 0 && p.ManaAct != p.ManaMax && p.VieAct != p.VieMax && p.Position.Val["nom"] != "Lige feu" { //condition de post mortem
					p.Position = p.ListeRetour[len(p.ListeRetour)-1]
					p.ListeRetour = p.ListeRetour[:len(p.ListeRetour)-1]
				}
			}
		} else {
			Affichage("Déplacement", []string{"On dirait que nous nous sommes perdu...", "Il n'y a rien ici.", "Il faut trouver un autre chemin."}, true, true)
			redemarrage = true
		}
	} else if direction == "retour" {
		if len(p.ListeRetour) > 0 {
			p.Position = p.ListeRetour[len(p.ListeRetour)-1]
			p.ListeRetour = p.ListeRetour[:len(p.ListeRetour)-1]
			Affichage("Déplacement", []string{"félicitation nous sommes désormais arrivé à " + p.Position.Val["nom"]}, true, true)
		} else {
			Affichage("Déplacement", []string{"On dirait que nous nous sommes perdu...", "Il n'y a rien ici.", "Il faut trouver un autre chemin."}, true, true)
			redemarrage = true
		}
	} else if direction == "hub" {
		var base Arbre
		base.Initialisation()
		p.Position = base
		p.VieAct = p.VieMax
		if p.Inv.Liste_Consommables[0].Quantite < 3 {
			p.Inv.Liste_Consommables[0].Quantite = 3
		}
		p.ListeRetour = []Arbre{}
		if p.Ames > 20 {
			p.Ames -= 20
			Affichage("Déplacement", []string{"félicitation nous sommes désormais arrivé au hub."}, true, true)
		} else {
			Affichage("Déplacement", []string{"On dirait que nous n'avons pas assez d'âmes pour revenir au hub."}, true, true)
			redemarrage = true
		}
	}
	return redemarrage
}
