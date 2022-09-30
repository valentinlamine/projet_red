package database

import (
	"fmt"
)

type Personnage struct {
	Nom    string // Nom du personnage
	Classe string // Classe du personnage (Guerrier, Chevalier, pyromancien, mendiant)
	Niveau int    // Niveau du personnage
	//stats améliorables
	Vitalite     int // Vitalité du personnage (modificateur de VieMax)
	Force        int // Force du personnage (dégâts infligés, ports d'armes, etc.)
	Dexterite    int // Dextérité du personnage (dégâts infligés, esquive, etc.)
	Intelligence int // Intelligence du personnage (dégâts magiques, etc.)
	//état du personnage
	Position            Arbre // Position du personnage sur la carte
	ListeRetour         []Arbre
	VieMax              int // VieMax du personnage (points de vie)
	VieAct              int // Points de vie actuels
	ManaMax             int // ManaAct maximum du personnage
	ManaAct             int // ManaAct du personnage
	Inv                 Inventaire
	PoidsMax            int // Poids maximum que peut porter le personnage
	PoidsAct            int // Poids total des objets équipés
	Degats              int
	EquipementArmes     Armes
	EquipementBoucliers Boucliers
	EquipementArmures   map[string]Armures
	Initiative          int
	Ames                int // Nombre d'âmes du personnage(Argent/EXP)
}

func (p *Personnage) Initialisation_Interne(nom, classe string, Vit, For, Dex, Int, Ames int) {
	//Attribut
	p.Nom = nom
	p.Classe = classe
	p.Niveau = 10
	p.Vitalite = Vit
	p.Force = For
	p.Dexterite = Dex
	p.Intelligence = Int
	//Etat Max
	p.VieMax = 20 * p.Vitalite
	p.ManaMax = 20 * p.Intelligence
	p.PoidsMax = 5 * p.Force
	//Etat actuel
	p.Initiative = 10 + p.Dexterite
	p.VieAct = p.VieMax
	p.ManaAct = p.ManaMax
	p.PoidsAct = 0
	p.Degats = p.EquipementArmes.Degats + p.Force
	p.EquipementArmures = make(map[string]Armures)
	p.Ames = Ames
	p.ListeRetour = make([]Arbre, 0)
}

func (p *Personnage) Initialisation(nom, classe string) {
	if classe == "Guerrier" {
		p.Initialisation_Interne(nom, "Guerrier", 11, 12, 9, 6, 0)
		p.Inv.Liste_Consommables[0].Quantite = 3
		p.Inv.Liste_Sorts[0].EstDebloque = true
		p.Inv.Liste_Armes[3].EstDebloque = true     //débloquer Uchigatana
		p.Inv.Liste_Boucliers[0].EstDebloque = true //débloquer Bouclier de bois
		p.Equiper(p.Inv.Liste_Armes[3], false)      //équiper Uchigatana
		p.Equiper(p.Inv.Liste_Boucliers[0], false)  //équiper Bouclier de bois

		//Débloquer armures
		p.Inv.Liste_Armures_Tete[0].EstDebloque = true
		p.Inv.Liste_Armures_Torse[0].EstDebloque = true
		p.Inv.Liste_Armures_Bras[0].EstDebloque = true
		p.Inv.Liste_Armures_Jambes[0].EstDebloque = true
		//Equiper armures
		p.Equiper(p.Inv.Liste_Armures_Tete[0], false)
		p.Equiper(p.Inv.Liste_Armures_Torse[0], false)
		p.Equiper(p.Inv.Liste_Armures_Bras[0], false)
		p.Equiper(p.Inv.Liste_Armures_Jambes[0], false)

	} else if classe == "Chevalier" {
		p.Initialisation_Interne(nom, "Chevalier", 10, 11, 8, 10, 0)
		p.Inv.Liste_Consommables[0].Quantite = 3
		p.Inv.Liste_Sorts[0].EstDebloque = true
		p.Inv.Liste_Armes[1].EstDebloque = true          //débloquer Claymore
		p.Inv.Liste_Boucliers[2].EstDebloque = true      //débloquer Bouclier de fer
		p.Equiper(p.Inv.Liste_Armes[1], false)           //équiper Claymore
		p.EquipementBoucliers = p.Inv.Liste_Boucliers[2] //équiper Bouclier de fer

		//Débloquer armures
		p.Inv.Liste_Armures_Tete[0].EstDebloque = true
		p.Inv.Liste_Armures_Torse[0].EstDebloque = true
		p.Inv.Liste_Armures_Bras[0].EstDebloque = true
		p.Inv.Liste_Armures_Jambes[0].EstDebloque = true
		//Equiper armures
		p.Equiper(p.Inv.Liste_Armures_Tete[0], false)
		p.Equiper(p.Inv.Liste_Armures_Torse[0], false)
		p.Equiper(p.Inv.Liste_Armures_Bras[0], false)
		p.Equiper(p.Inv.Liste_Armures_Jambes[0], false)

	} else if classe == "Pyromancien" {
		p.Initialisation_Interne(nom, "Pyromancien", 9, 7, 12, 12, 0)
		p.Inv.Liste_Consommables[0].Quantite = 3
		p.Inv.Liste_Sorts[0].EstDebloque = true
		p.Inv.Liste_Armes[0].EstDebloque = true //débloquer Dague
		p.Equiper(p.Inv.Liste_Armes[0], false)  //équiper Dague
		p.Inv.Liste_Sorts[1].EstDebloque = true //débloquer boule de Feu

		//Débloquer armures
		p.Inv.Liste_Armures_Tete[2].EstDebloque = true
		p.Inv.Liste_Armures_Torse[2].EstDebloque = true
		p.Inv.Liste_Armures_Bras[2].EstDebloque = true
		p.Inv.Liste_Armures_Jambes[2].EstDebloque = true
		//Equiper armures
		p.Equiper(p.Inv.Liste_Armures_Tete[2], false)
		p.Equiper(p.Inv.Liste_Armures_Torse[2], false)
		p.Equiper(p.Inv.Liste_Armures_Bras[2], false)
		p.Equiper(p.Inv.Liste_Armures_Jambes[2], false)

	} else if classe == "Mendiant" {
		p.Initialisation_Interne(nom, "Mendiant", 9, 9, 9, 9, 0)
		p.Inv.Liste_Consommables[0].Quantite = 3
		p.Inv.Liste_Sorts[0].EstDebloque = true
		p.Inv.Liste_Armes[4].EstDebloque = true //débloquer Baton
		p.Equiper(p.Inv.Liste_Armes[4], false)  //équiper Baton

		//Le pauvre haha le loser, il est pauvre et il a pas d'armure haha

		//Initialisation des mobs
	} else if classe == "Carcasse" {
		p.Initialisation_Interne("Carcasse", "Carcasse", 7, 9, 10, 5, 200)
		p.Inv.Liste_Armes[4].EstDebloque = true //débloquer Baton
		p.Equiper(p.Inv.Liste_Armes[4], false)  //équiper Baton

	} else if classe == "Chevalier mort-vivant" {
		p.Initialisation_Interne("Chevalier mort-vivant", "Chevalier mort-vivant", 10, 9, 13, 5, 400)
		p.Inv.Liste_Armes[4].EstDebloque = true //débloquer Baton
		p.Equiper(p.Inv.Liste_Armes[4], false)  //équiper Baton

	} else if classe == "Champion mort-vivant" {
		p.Initialisation_Interne("Champion mort-vivant", "Champion mort-vivant", 13, 15, 15, 5, 700)
		p.Inv.Liste_Armes[4].EstDebloque = true //débloquer Baton
		p.Equiper(p.Inv.Liste_Armes[4], false)  //équiper Baton

		//Initialisation des boss
	} else if classe == "Gargouille" {
		p.Initialisation_Interne("Gargouille", "Gargouille", 40, 50, 20, 5, 4000)
		p.Inv.Liste_Armes[4].EstDebloque = true //débloquer Baton
		p.Equiper(p.Inv.Liste_Armes[4], false)  //équiper Baton

	} else if classe == "Démon Capra" {
		p.Initialisation_Interne("Démon Capra", "Démon Capra", 20, 20, 15, 5, 2000)
		p.Inv.Liste_Armes[4].EstDebloque = true //débloquer Baton
		p.Equiper(p.Inv.Liste_Armes[4], false)  //équiper Baton

	} else if classe == "Démon taureau" {
		p.Initialisation_Interne("Démon taureau", "Démon taureau", 30, 30, 11, 5, 3000)
		p.Inv.Liste_Armes[4].EstDebloque = true //débloquer Baton
		p.Equiper(p.Inv.Liste_Armes[4], false)  //équiper Baton
	}
}

func (p *Personnage) Est_Mort() bool {
	if p.VieAct <= 0 {
		p.Ames = 0
		p.VieAct = p.VieMax
		var hub Arbre
		hub.Initialisation()
		p.ListeRetour = make([]Arbre, 0)
		p.Position = hub
		p.VieAct = p.VieMax
		p.ManaAct = p.ManaMax
		fmt.Println("\033[31mVous êtes mort !\033[0m")
		Attendre()
		return true
	} else {
		return false
	}
}
