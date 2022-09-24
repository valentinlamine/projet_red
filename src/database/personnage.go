package database

import "fmt"

type Personnage struct {
	Nom    string // Nom du personnage
	Classe string // Classe du personnage (Guerrier, Chevalier, pyromancien, mendiant)
	Niveau int    // Niveau du personnage
	//stats améliorables
	Pvmax        int // Pvmax du personnage (points de vie)
	Vitalite     int // Vitalité du personnage (modificateur de Pvmax)
	Force        int // Force du personnage (dégâts infligés, ports d'armes, etc.)
	Dexterite    int // Dextérité du personnage (dégâts infligés, esquive, etc.)
	Intelligence int // Intelligence du personnage (dégâts magiques, etc.)
	//état du personnage
	Position            Arbre // Position du personnage sur la carte
	Pvact               int   // Points de vie actuels
	Inv                 Inventaire
	PoidsEquip          int // Poids total des objets équipés
	PoidsMax            int // Poids maximum que peut porter le personnage
	Degats              int
	EquipementArmes     [1]Armes
	EquipementBoucliers [1]Boucliers
	EquipementArmures   map[string]Armures
	Ames                int // Nombre d'âmes du personnage(Argent/EXP)
}

func (p *Personnage) InitIntern(nom, classe string, Vit, For, Dex, Int int) {
	p.Nom = nom
	p.Classe = classe
	p.Niveau = 1
	p.Vitalite = Vit
	p.Pvmax = 20 * p.Vitalite
	p.Force = For
	p.Dexterite = Dex
	p.Intelligence = Int
	p.Pvact = p.Pvmax / 2
	p.PoidsEquip = 0
	p.PoidsMax = 5 * p.Force
	p.Degats = p.EquipementArmes[0].deg
	p.Ames = 0
}

func (p *Personnage) Init(nom, classe string) {
	if classe == "Guerrier" {
		p.InitIntern(nom, "Guerrier", 11, 12, 9, 8)
		p.Inv.Liste_armes[3].Set_Armes("isUnlocked", "true")         //débloquer Uchigatana
		p.Inv.Liste_boucliers[0].Set_Boucliers("isUnlocked", "true") //débloquer Bouclier de bois
		p.EquipementArmes[0] = p.Inv.Liste_armes[3]                  //équiper Uchigatana
		p.EquipementBoucliers[0] = p.Inv.Liste_boucliers[0]          //équiper Bouclier de bois

	} else if classe == "Chevalier" {
		p.InitIntern(nom, "Chevalier", 10, 11, 8, 10)
		p.Inv.Liste_armes[1].Set_Armes("isUnlocked", "true")         //débloquer Claymore
		p.Inv.Liste_boucliers[2].Set_Boucliers("isUnlocked", "true") //débloquer Bouclier de fer
		p.EquipementArmes[0] = p.Inv.Liste_armes[1]                  //équiper Claymore
		p.EquipementBoucliers[0] = p.Inv.Liste_boucliers[2]          //équiper Bouclier de fer

	} else if classe == "Pyromancien" {
		p.InitIntern(nom, "Pyromancien", 9, 9, 11, 12)
		p.Inv.Liste_armes[0].Set_Armes("isUnlocked", "true") //débloquer Dague
		p.EquipementArmes[0] = p.Inv.Liste_armes[0]          //équiper Dague

	} else if classe == "Mendiant" {
		p.InitIntern(nom, "Mendiant", 9, 9, 9, 9)
		p.Inv.Liste_armes[4].IsUnlocked = true          //débloquer Baton
		p.Inv.Liste_armures_tete[6].isUnlocked = true   //débloquer Vide
		p.Inv.Liste_armures_torse[6].isUnlocked = true  //débloquer Vide
		p.Inv.Liste_armures_bras[6].isUnlocked = true   //débloquer Vide
		p.Inv.Liste_armures_jambes[6].isUnlocked = true //débloquer Vide

		p.EquipementArmes[0] = p.Inv.Liste_armes[4]                   //équiper Baton
		p.EquipementArmures["Tete"] = p.Inv.Liste_armures_tete[6]     //équiper Vide
		p.EquipementArmures["Torse"] = p.Inv.Liste_armures_torse[6]   //équiper Vide
		p.EquipementArmures["Bras"] = p.Inv.Liste_armures_bras[6]     //équiper Vide
		p.EquipementArmures["Jambes"] = p.Inv.Liste_armures_jambes[6] //équiper Vide
	}
}

func (p *Personnage) IsDead() bool {
	if p.Pvact <= 0 {
		p.Ames = 0
		//print in red "Vous êtes mort !"
		fmt.Println("\033[31mVous êtes mort !\033[0m")
		return true
	} else {
		return false
	}
}
