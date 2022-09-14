package main

type arbre struct {
	Val    map[string]string
	Gauche *arbre
	Droite *arbre
	Centre *arbre
}

// number of key in the val : ("name", "mob_nb", "mob_type", "boss_type", "feu", "forge", "secret_destination", "secret_unlock")
func (a *arbre) Insert(val map[string]string) {
	if a.Val == nil {
		a.Val = val
		a.Gauche = &arbre{}
		a.Droite = &arbre{}
		a.Centre = &arbre{}
	}
}

func main() {
	var hub arbre
	hub.Insert(map[string]string{"name": "hub", "mob_nb": "0", "mob_type": "none", "boss_type": "none", "feu": "vrai", "forge": "false", "secret_destination": "none", "secret_unlock": "false"})
	hub.Centre.Insert(map[string]string{"name": "Camps de mort vivant", "mob_nb": "4", "mob_type": "none", "boss_type": "none", "feu": "false", "forge": "false", "secret_destination": "none", "secret_unlock": "false"})
	hub.Centre.Droite.Insert(map[string]string{"name": "Forteresse", "mob_nb": "2", "mob_type": "none", "boss_type": "none", "feu": "false", "forge": "false", "secret_destination": "none", "secret_unlock": "false"})
	hub.Centre.Gauche.Insert(map[string]string{"name": "Batiment détruit", "mob_nb": "0", "mob_type": "none", "boss_type": "none", "feu": "true", "forge": "false", "secret_destination": "none", "secret_unlock": "false"})
	hub.Centre.Gauche.Centre.Insert(map[string]string{"name": "Tour ouest", "mob_nb": "6", "mob_type": "none", "boss_type": "none", "feu": "false", "forge": "false", "secret_destination": "none", "secret_unlock": "false"})
	hub.Centre.Gauche.Centre.Centre.Insert(map[string]string{"name": "Muraille", "mob_nb": "0", "mob_type": "none", "boss_type": "Gargouille", "feu": "false", "forge": "false", "secret_destination": "none", "secret_unlock": "false"})
	hub.Centre.Gauche.Centre.Centre.Centre.Insert(map[string]string{"name": "Tour nord", "mob_nb": "0", "mob_type": "none", "boss_type": "none", "feu": "false", "forge": "false", "secret_destination": "none", "secret_unlock": "false"})
}
