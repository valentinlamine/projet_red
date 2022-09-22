package database

//Liste des variables des armes
var dague Armes
var claymore Armes
var rapière Armes
var uchigatana Armes
var bâton Armes
var hachequeue Armes

//Liste des variables des boucliers
var bouclier_bois Boucliers
var Bouclier_bois_cerf Boucliers
var Bouclier_fer Boucliers
var Bouclier_acier Boucliers
var Bouclier_mithril Boucliers
var Bouclier_Havel Boucliers

//Liste des variables des consommables
var Fiole_Estus Consommable
var Résine_pin_doré Consommable
var Résine_pin_brulé Consommable
var Résine_pin_pourri Consommable
var Potion_poids_max Consommable
var Fiole_Pin_Pourri Consommable

//Liste des variables des armures de tête
var Casque_carcasse Armures
var Casque_Havel Armures
var Casque_cuir Armures
var Casque_chevalier Armures
var Casque_feu Armures
var Casque_dragon Armures
var Casque_vide Armures

//Liste des variables des armures de torse
var Armure_carcasse Armures
var Armure_Havel Armures
var Armure_cuir Armures
var Armure_chevalier Armures
var Armure_feu Armures
var Armure_dragon Armures
var Armure_vide Armures

//Liste des variables des armures de bras
var Gantelets_carcasse Armures
var Gantelets_Havel Armures
var Gantelets_cuir Armures
var Gantelets_chevalier Armures
var Gantelets_feu Armures
var Gantelets_dragon Armures
var Gantelets_vide Armures

//Liste des variables des armures de jambes
var Jambières_carcasse Armures
var Jambières_Havel Armures
var Jambières_cuir Armures
var Jambières_chevalier Armures
var Jambières_feu Armures
var Jambières_dragon Armures
var Jambières_vide Armures

type Inventaire struct {
	Liste_armes          []Armes
	Liste_armures_tete   []Armures
	Liste_armures_torse  []Armures
	Liste_armures_bras   []Armures
	Liste_armures_jambes []Armures
	Liste_boucliers      []Boucliers
	Liste_consommables   []Consommable
	Liste_sort           []string //TODO Sort
}

func (i *Inventaire) Init() {
	//Initialisation des armes
	dague.Init_Armes("Dague")
	claymore.Init_Armes("Claymore")
	rapière.Init_Armes("Rapière")
	uchigatana.Init_Armes("Uchigatana")
	bâton.Init_Armes("Bâton")
	hachequeue.Init_Armes("Hache queue de gargouille")
	i.Liste_armes = append(i.Liste_armes,
		dague,
		claymore,
		rapière,
		uchigatana,
		bâton,
		hachequeue)

	//Initialisation des boucliers
	bouclier_bois.Init_Bouclier("Bouclier en bois")
	Bouclier_bois_cerf.Init_Bouclier("Bouclier en bois de cerf")
	Bouclier_fer.Init_Bouclier("Bouclier en fer")
	Bouclier_acier.Init_Bouclier("Bouclier en acier")
	Bouclier_mithril.Init_Bouclier("Bouclier en mithril")
	Bouclier_Havel.Init_Bouclier("Bouclier d'Havel")
	i.Liste_boucliers = append(i.Liste_boucliers,
		bouclier_bois,
		Bouclier_bois_cerf,
		Bouclier_fer,
		Bouclier_acier,
		Bouclier_mithril,
		Bouclier_Havel)

	//Initialisation des consommables
	Fiole_Estus.Init_Consommable("Fiole d'Estus")
	Résine_pin_doré.Init_Consommable("Résine de pin doré")
	Résine_pin_brulé.Init_Consommable("Résine de pin brulé")
	Résine_pin_pourri.Init_Consommable("Résine de pin pourri")
	Potion_poids_max.Init_Consommable("Potion de poids max")
	Fiole_Pin_Pourri.Init_Consommable("Fiole d'essence de pin pourri")
	i.Liste_consommables = append(i.Liste_consommables,
		Fiole_Estus,
		Résine_pin_doré,
		Résine_pin_brulé,
		Résine_pin_pourri,
		Potion_poids_max,
		Fiole_Pin_Pourri)

	//Initialisation des armures de tête
	Casque_carcasse.Init_Armures("Casque de Carcasse")
	Casque_Havel.Init_Armures("Casque d'Havel")
	Casque_cuir.Init_Armures("Bandeau de cuir noir")
	Casque_chevalier.Init_Armures("Heaume de Chevalier")
	Casque_feu.Init_Armures("Capuche de feu")
	Casque_dragon.Init_Armures("Tête de Dragon")
	Casque_vide.Init_Armures("vide")
	i.Liste_armures_tete = append(i.Liste_armures_tete,
		Casque_carcasse,
		Casque_Havel,
		Casque_cuir,
		Casque_chevalier,
		Casque_feu,
		Casque_dragon,
		Casque_vide)

	//Initialisation des armures de torse
	Armure_carcasse.Init_Armures("Plastron de Carcasse")
	Armure_Havel.Init_Armures("Plastron d'Havel")
	Armure_cuir.Init_Armures("Veste de cuir  noir")
	Armure_chevalier.Init_Armures("Plastron de Chevalier")
	Armure_feu.Init_Armures("Manteau de feu")
	Armure_dragon.Init_Armures("Ecailles de Dragon")
	Armure_vide.Init_Armures("vide")
	i.Liste_armures_torse = append(i.Liste_armures_torse,
		Armure_carcasse,
		Armure_Havel,
		Armure_cuir,
		Armure_chevalier,
		Armure_feu,
		Armure_dragon,
		Armure_vide)

	//Initialisation des armures de bras
	Gantelets_carcasse.Init_Armures("Gantelet de Carcasse")
	Gantelets_Havel.Init_Armures("Gantelet d'Havel")
	Gantelets_cuir.Init_Armures("Bandeau de cuir  noir")
	Gantelets_chevalier.Init_Armures("Plastron de Chevalier")
	Gantelets_feu.Init_Armures("Manchette de feu")
	Gantelets_dragon.Init_Armures("Griffes de Dragon")
	Gantelets_vide.Init_Armures("vide")
	i.Liste_armures_bras = append(i.Liste_armures_bras,
		Gantelets_carcasse,
		Gantelets_Havel,
		Gantelets_cuir,
		Gantelets_chevalier,
		Gantelets_feu,
		Gantelets_dragon,
		Gantelets_vide)

	//Initialisation des armures de jambes
	Jambières_carcasse.Init_Armures("Jambières de Carcasse")
	Jambières_Havel.Init_Armures("Jambières d'Havel")
	Jambières_cuir.Init_Armures("Bottes de cuir  noir")
	Jambières_chevalier.Init_Armures("Jambières de Chevalier")
	Jambières_feu.Init_Armures("Bottes de feu")
	Jambières_dragon.Init_Armures("Pattes de Dragon")
	Jambières_vide.Init_Armures("vide")
	i.Liste_armures_jambes = append(i.Liste_armures_jambes,
		Jambières_carcasse,
		Jambières_Havel,
		Jambières_cuir,
		Jambières_chevalier,
		Jambières_feu,
		Jambières_dragon,
		Jambières_vide)

}
