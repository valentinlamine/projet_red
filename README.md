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
## Documentation utilisateur

Le menu des statistiques permet aussi d'afficher les armes et amures équipés.
Le menu de l'inventaire permet de choisir une arme ou une armure à équiper, ou de choisir une potion et la consommer.
Le menu déplacement permet de choisir une direction dans laquelle se déplacer.
Le menu hub permet d'accéder aux marchands et forgeron qui se trouvent au hub, Lige feu.
    André le forgeron permet d'ameliorer une arme ou une armure contre des éclats de Titanite.
    Petrus marchand permet d'acheter des potions et des armes et armures.
    Laurentius est un entraineur sorcier, il permet d'acheter des sorts et de les apprendre.
    La Gardienne du feu est un PNJ qui permet d'améliorer nos statistiques contre des âmes.

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
## Documentation technique

Notre jeu se base sur un système de menu, qui permet de naviguer entre les différentes parties du jeu.
Dans la fonction `main()`, on appelle la fonction `menu()` au travers d'une boucle for qui permet de naviguer entre les différents menus du jeu via un switch case.

Nos objets et personnages sont définit via des structures. Ce qui permet de les manipuler facilement et d'attribuer différentes statistiques selon nos besoins.

Le joueur et les ennemis sont basés sur la même structure, ce qui permet de les manipuler de la même manière. Vu que la vie et les dégâts ne dépendent pas seulement de l'équipement, nous pouvons dans le cas des monstres, les initialisé avec de plus grandes statistiques pour qu'ils représentent un défi au joueur.

L'inventaire est aussi basé sur une structure qui regroupe des listes d'autres structures. Cela permet de manipuler facilement les objets et de les trier selon leur type.
L'inventaire est lui-même un attribut de la structure du joueur, ce qui permet de l'initialiser avec des objets dès le début du jeu.
L'inventaire du joueur est initialiser avec toutes les armes et armures du jeu, mais qui sont pour l'instant bloquées pour le joueur. Il peut les débloquer en achetant les armes où en les ramassant sur des monstres.

Les combats sont basés sur deux fonction qui bouclent. La première fonction, `Faire_Combat()`, s'éxécute tant qu'il y a des monstres, si le joueur choisis d'attaquer, il rentre dans la fonction `Combat()`, sinon il fuit le combat et retourne à la zone précédante. La fonction `Combat()` s'éxécute tant que le joueur ou l'ennemi n'a pas 0 points de vie. Elle appelle la fonction `Tour_Joueur()` puis `Tour_Ennemi()` ou inversement selon la statistique de Dextérité des combatants. La fonction `Tour_Joueur()` laisse le choix entre attaque, attaque lourdes, sorts et inventaire. L'attaque lourde inflige deux fois plus de dégât mais coûte de la mana. La fonction `Tour_Ennemi()` inflige des dégâts au joueur en l'attaquant. Tous les 3 tours le monstre inflige une attaques critique, et tous les 5 tours le monstre empoisonne le joueur. Le nombre de tour est stocké dans la varibale `CompteurTour`. Chaque tour le joueur regagne de la mana.

Notre carte du jeu est une structure Arbre, avec les attributs gauche, droite, centre et parent qui pointent vers d'autres Arbres. Chaque arbre représente une zone du jeu, cette zone peut contenir un marchand, des monstres ou un boss.
(N'allez pas voir l'initialisation de la carte pour votre santé mentalen ça marche juste, posez pas de questions)


## Contribuer

Pour contribuer au projet, il suffit de créer une branche et de faire une pull request.

## Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus d'informations.