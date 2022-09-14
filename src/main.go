package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Personnage struct {
	nom    string // Nom du personnage
	classe string // Classe du personnage (Guerrier, Chevalier, pyromancien, mendiant)
	niveau int    // Niveau du personnage
	//stats améliorables
	pvmax        int // pvmax du personnage (points de vie)
	vitalite     int // Vitalité du personnage (modificateur de pvmax)
	force        int // Force du personnage (dégâts infligés, ports d'armes, etc.)
	dexterite    int // Dextérité du personnage (dégâts infligés, esquive, etc.)
	intelligence int // Intelligence du personnage (dégâts magiques, etc.)
	//état du personnage
	pvact      int // Points de vie actuels
	inventaire []string
}

func main() {
	var p1 Personnage
	fmt.Println("Création de votre personnage")
	fmt.Println("Quel est le nom de votre personnage ?")
	nom := input()
	affichage([]string{"Quelle est la classe de votre personnage ?", "il y a 4 classes : Guerrier, Chevalier, Pyromancien, Mendiant", "Pour choisir guerrier, tapez 1", "Pour choisir chevalier, tapez 2", "Pour choisir pyromancien, tapez 3", "Pour choisir mendiant, tapez 4"})
	classe := input()
	classe_int, _ := strconv.Atoi(classe)
	if classe_int == 1 {
		p1.Init(nom, "Guerrier")
	} else if classe_int == 2 {
		p1.Init(nom, "Chevalier")
	} else if classe_int == 3 {
		p1.Init(nom, "Pyromancien")
	} else if classe_int == 4 {
		p1.Init(nom, "Mendiant")
	} else {
		fmt.Println("Classe inconnue, Veuiilez choisir entre Guerrier, Chevalier, Pyromancien ou Mendiant")
	}
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func affichage(list []string) {
	longest := 0
	for _, s := range list {
		if len(s) > longest {
			longest = len(s)
		}
	}
	fmt.Println("╒", strings.Repeat("═", longest+1), "╕")
	for _, s := range list {
		fmt.Println("│", s, strings.Repeat(" ", longest-len(s)), "│")
	}
	fmt.Println("╘", strings.Repeat("═", longest+1), "╛")

}

func (p *Personnage) Init(nom, classe string) {
	if classe == "Guerrrier" {
		p.nom = nom
		p.classe = classe
		p.niveau = 1
		p.pvmax = 10
		p.vitalite = 11
		p.force = 13
		p.dexterite = 13
		p.intelligence = 9
		p.pvact = p.pvmax
	} else if classe == "Chevalier" {
		p.nom = nom
		p.classe = classe
		p.niveau = 1
		p.pvmax = 10
		p.vitalite = 14
		p.force = 11
		p.dexterite = 11
		p.intelligence = 9
		p.pvact = p.pvmax
	} else if classe == "Pyromancien" {
		p.nom = nom
		p.classe = classe
		p.niveau = 1
		p.pvmax = 10
		p.vitalite = 10
		p.force = 12
		p.dexterite = 9
		p.intelligence = 10
		p.pvact = p.pvmax
	} else if classe == "Mendiant" {
		p.nom = nom
		p.classe = classe
		p.niveau = 1
		p.pvmax = 10
		p.vitalite = 11
		p.force = 11
		p.dexterite = 11
		p.intelligence = 11
		p.pvact = p.pvmax
	}
}
