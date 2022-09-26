package database

import "fmt"

type Personnage struct {
	Nom    string // Nom du personnage
	Classe string // Classe du personnage (Guerrier, Chevalier, pyromancien, mendiant)
	Niveau int    // Niveau du personnage
	//stats améliorables
	Vitalite     int // Vitalité du personnage (modificateur de Pvmax)
	Force        int // Force du personnage (dégâts infligés, ports d'armes, etc.)
	Dexterite    int // Dextérité du personnage (dégâts infligés, esquive, etc.)
	Intelligence int // Intelligence du personnage (dégâts magiques, etc.)
	//état du personnage
	Position            Arbre // Position du personnage sur la carte
	Pvmax               int   // Pvmax du personnage (points de vie)
	Pvact               int   // Points de vie actuels
	ManaMax             int   // Mana maximum du personnage
	Mana                int   // Mana du personnage
	Inv                 Inventaire
	PoidsEquip          int // Poids total des objets équipés
	PoidsMax            int // Poids maximum que peut porter le personnage
	Degats              int
	EquipementArmes     Armes
	EquipementBoucliers Boucliers
	EquipementArmures   map[string]Armures
	Initiative          int
	Ames                int // Nombre d'âmes du personnage(Argent/EXP)
}

func (p *Personnage) InitIntern(nom, classe string, Vit, For, Dex, Int, Ames int) {
	//Attribut
	p.Nom = nom
	p.Classe = classe
	p.Niveau = 1
	p.Vitalite = Vit
	p.Force = For
	p.Dexterite = Dex
	p.Intelligence = Int
	//Etat Max
	p.Pvmax = 20 * p.Vitalite
	p.ManaMax = 20 * p.Intelligence
	p.PoidsMax = 5 * p.Force
	//Etat actuel
	p.Initiative = 10 + p.Dexterite
	p.Pvact = p.Pvmax / 2
	p.Mana = p.ManaMax
	p.PoidsEquip = 0
	p.Degats = p.EquipementArmes.Deg + p.Force
	p.EquipementArmures = make(map[string]Armures)
	p.Ames = Ames
}

func (p *Personnage) Init(nom, classe string) {
	if classe == "Guerrier" {
		p.InitIntern(nom, "Guerrier", 11, 12, 9, 8, 0)
		p.Inv.Liste_armes[3].IsUnlocked = true     //débloquer Uchigatana
		p.Inv.Liste_boucliers[0].IsUnlocked = true //débloquer Bouclier de bois
		p.Equiper(p.Inv.Liste_armes[3])            //équiper Uchigatana
		p.Equiper(p.Inv.Liste_boucliers[0])        //équiper Bouclier de bois

	} else if classe == "Chevalier" {
		p.InitIntern(nom, "Chevalier", 10, 11, 8, 10, 0)
		p.Inv.Liste_armes[1].IsUnlocked = true           //débloquer Claymore
		p.Inv.Liste_boucliers[2].IsUnlocked = true       //débloquer Bouclier de fer
		p.Equiper(p.Inv.Liste_armes[1])                  //équiper Claymore
		p.EquipementBoucliers = p.Inv.Liste_boucliers[2] //équiper Bouclier de fer

	} else if classe == "Pyromancien" {
		p.InitIntern(nom, "Pyromancien", 9, 9, 11, 12, 0)
		p.Inv.Liste_armes[0].IsUnlocked = true //débloquer Dague
		p.Equiper(p.Inv.Liste_armes[0])        //équiper Dague

	} else if classe == "Mendiant" {
		p.InitIntern(nom, "Mendiant", 9, 9, 9, 9, 0)
		p.Inv.Liste_armes[4].IsUnlocked = true          //débloquer Baton
		p.Inv.Liste_armures_tete[0].IsUnlocked = true   //débloquer première armure de tête
		p.Inv.Liste_armures_torse[0].IsUnlocked = true  //débloquer première armure de torse
		p.Inv.Liste_armures_bras[0].IsUnlocked = true   //débloquer première armure de bras
		p.Inv.Liste_armures_jambes[0].IsUnlocked = true //débloquer première armure de jambes

		p.Equiper(p.Inv.Liste_armes[4])          //équiper Baton
		p.Equiper(p.Inv.Liste_armures_tete[0])   //équiper première armure de tête
		p.Equiper(p.Inv.Liste_armures_torse[0])  //équiper première armure de torse
		p.Equiper(p.Inv.Liste_armures_bras[0])   //équiper première armure de bras
		p.Equiper(p.Inv.Liste_armures_jambes[0]) //équiper première armure de jambes
		p.Pvmax = p.Pvmax +
			p.EquipementArmures["Tete"].Pvbonus +
			p.EquipementArmures["Torse"].Pvbonus +
			p.EquipementArmures["Jambes"].Pvbonus +
			p.EquipementArmures["Pieds"].Pvbonus

		//Init des mobs
	} else if classe == "Carcasse" {
		p.InitIntern("Carcasse", "Carcasse", 7, 9, 8, 5, 200)
		p.Inv.Liste_armes[4].IsUnlocked = true //débloquer Baton
		p.Equiper(p.Inv.Liste_armes[4])        //équiper Baton

	} else if classe == "Chevalier mort-vivant" {
		p.InitIntern("Chevalier mort-vivant", "Chevalier mort-vivant", 10, 9, 7, 5, 400)
		p.Inv.Liste_armes[4].IsUnlocked = true //débloquer Baton
		p.Equiper(p.Inv.Liste_armes[4])        //équiper Baton

	} else if classe == "Chambion mort-vivant" {
		p.InitIntern("Chambion mort-vivant", "Chambion mort-vivant", 13, 15, 13, 5, 700)
		p.Inv.Liste_armes[4].IsUnlocked = true //débloquer Baton
		p.Equiper(p.Inv.Liste_armes[4])        //équiper Baton

		//Init des boss
	} else if classe == "Gargouille" {
		p.InitIntern("Gargouille", "Gargouille", 40, 50, 20, 5, 4000)
		p.Inv.Liste_armes[4].IsUnlocked = true //débloquer Baton
		p.Equiper(p.Inv.Liste_armes[4])        //équiper Baton

	} else if classe == "Démon Capra" {
		p.InitIntern("Démon Capra", "Démon Capra", 20, 20, 15, 5, 2000)
		p.Inv.Liste_armes[4].IsUnlocked = true //débloquer Baton
		p.Equiper(p.Inv.Liste_armes[4])        //équiper Baton

	} else if classe == "Démon taureau" {
		p.InitIntern("Démon taureau", "Démon taureau", 30, 30, 11, 5, 3000)
		p.Inv.Liste_armes[4].IsUnlocked = true //débloquer Baton
		p.Equiper(p.Inv.Liste_armes[4])        //équiper Baton
	}
}

func (p *Personnage) IsDead() bool {
	if p.Pvact <= 0 {
		p.Ames = 0
		p.Pvact = p.Pvmax
		var hub Arbre
		hub.Init()
		p.Position = hub
		p.Pvact = p.Pvmax / 2
		fmt.Println("\033[31mVous êtes mort !\033[0m")
		Attendre()
		return true
	} else {
		return false
	}
}
