# Projet Red: Text Souls

Text Souls est un jeu purement et sincèrement sorti de l'imagination de l'équipe de DEPUISLOGICIEL, notre studio.
Toute ressemblance avec un autre jeu est fortuite 

## Langage de programmation

- Go

## Auteurs

COLLIOT--MARTINEZ Théo
LAMINE Valentin

## Installation

Pour installer le jeu, il suffit de cloner le dépôt git et de lancer le fichier `main.go` avec la commande `go run main.go`

## Utilisation

Le jeu est jouable en ligne de commande, il suffit de suivre les instructions.

## Documentation technique

Notre jeu se base sur un système de menu, qui permet de naviguer entre les différentes parties du jeu.
Dans la fonction `main()`, on appelle la fonction `menu()` au travers d'une boucle for qui permet de naviguer entre les différents menus du jeu via un switch case.

Nos objets et personnages sont définit via des structures. Ce qui permet de les manipuler facilement et d'attribuer différentes statistiques selon nos besoins.

Le joueur et les ennemis sont basés sur la même structure, ce qui permet de les manipuler de la même manière. Vu que la vie et les dégâts ne dépendent pas seulement de l'équipement, nous pouvons dans le cas des monstres, les initialisé avec de plus grandes statistiques pour qu'ils représentent un défi au joueur.

L'inventaire est aussi basé sur une structure qui regroupe des listes d'autres structures. Cela permet de manipuler facilement les objets et de les trier selon leur type.
L'inventaire est lui-même un attribut de la structure du joueur, ce qui permet de l'initialiser avec des objets dès le début du jeu.

## Documentation utilisateur

Le menu des statistiques permet aussi d'afficher les armes et amures équipés.
Le menu de l'inventaire permet de choisir une arme ou une armure à équiper, ou de choisir une potion et la consommer.
Le menu déplacement permet de choisir une direction dans laquelle se déplacer.
Le menu hub permet d'accéder aux marchands et forgeron qui se trouvent au hub, Lige feu.
    Le menu forgeron permet d'ameliorer une arme ou une armure contre des éclats de Titanite.
    Le menu marchand permet d'acheter des potions et des armes et armures.
    Un des marchands est un entraineur sorcier, il permet d'acheter des sorts et de les apprendre.

Lors des déplacements, le joueur peut rencontrer des ennemis, qui sont des PNJ de la structure Personnage, la même que le joueur.
Lors d'une rencontre d'ennemi, le joueur peut choisir de combattre ou de fuir.
S'il choisit de combattre, il peut choisir d'attaquer, de lancer un sort ou d'utiliser une potion.

C'est en tuant des ennemis que le joueur obtiens des âmes, qui lui permettent d'augmenter ses statistiques et qui servent aussi de monnaie.

Le joueur peut mourir, dans ce cas il perd ses âmes, les ennimis réapparaissent et il est ramené au hub.

### Statistiques

Le joueur a 4 statistiques: la Vitalité, la Force, la Dextérité et l'Intelligence. Ces statistique sont améliorable auprès de la Gardienne du feu, qui se trouve au hub.

La Vitalité augmente la vie du joueur.
La Force augmente les dégâts des attaques physiques et permet d'augemanter la limite de poids maximum équiper.
La Dextérité augmente les chances de commancer le combat en premier.
L'Intelligence augmente les dégâts des attaques magiques.

### Equipement

Tout l'équipement du joueur possède des statistiques minimum pour être équipé.
Les armures et boucliers donnent des point de vie bonus.

## Contribuer

Pour contribuer au projet, il suffit de créer une branche et de faire une pull request.

## Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus d'informations.