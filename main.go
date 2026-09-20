package main

import (
	"fmt"
)
type Soldat struct {
    nom     string
    vie     int
    attaque int
}

func afficherEquipe(equipe [6]Soldat){

	println("=== ÉQUIPE ===")  
	println() 

	for i := 0; i < len(equipe);  i++ {

		fmt.Println(equipe[i].nom)
		fmt.Println("Vie :", equipe[i].vie)
		fmt.Println("Attaque :", equipe[i].attaque) 
		fmt.Println()
	}
}

func main() {

	equipe := [6]Soldat{
		{"Firas", 1200, 250},
		{"Julien", 850, 320},
		{"Brice", 1500, 180},
		{"Kantin", 700, 400},
		{"Chritophe", 1000, 280},
		{"Issa", 500, 450},
	}
	afficherEquipe(equipe)

	// exo 2 affichage
	println("=== ANALYSE ===")
	println()

	plusDeVie := trouverPlusDeVie(equipe)
	println("Soldat avec le plus de vie :", plusDeVie.nom)
	println("Vie :", plusDeVie.vie)
	println()

	plusDAttaque := trouverPlusDAttaque(equipe)
	println("Soldat avec la plus grande attaque :", plusDAttaque.nom)
	println("Attaque :", plusDAttaque.attaque)
	println()

	vieMoyenne := calculerVieMoyenne(equipe)
	fmt.Printf("Vie moyenne : %.2f\n", vieMoyenne)  // %.2f pour 2 chiffres après la virgule (et fmt. parce que ca existe pas Printf tout seul)
	println() 

	faibles := compterFaibles(equipe)
	println("Soldats avec moins de 800 PV :", faibles)

	// exo 3 affichage
	println()
	var degats int

	print("Dégâts de l'ennemi : ")
	fmt.Scan(&degats)

	attaquerEquipe(&equipe, degats)

	println()
	println("=== APRÈS L'ATTAQUE ===")
	afficherEquipe(equipe)

	// exo 4 affichage
	println()
	afficherEtat(equipe)

	// exo 5 affichage avec boucle 
	println()
	println("=== BATAILLE ===")
	println()

	var nombreAttaques int

	print("Nombre d'attaques ennemies : ")
	fmt.Scan(&nombreAttaques)

	for i := 1; i <= nombreAttaques; i++ {

	var degats int

	print("Attaque ", i, " : ")
	fmt.Scan(&degats)

	attaquerEquipe(&equipe, degats)

	println()
	print("=== APRÈS L'ATTAQUE ", i, " ===\n")
	println()

	afficherEtat(equipe)
	println()
}

}


// exo 2!!!!

func trouverPlusDeVie(equipe [6]Soldat) Soldat {
	plusDeVie := equipe[0]

	for i := 1; i < len(equipe); i++ {
		if equipe[i].vie > plusDeVie.vie {
			plusDeVie = equipe[i]
		}
	}
	return plusDeVie
}

func trouverPlusDAttaque(equipe [6]Soldat) Soldat {
	plusDAttaque := equipe[0]

	for i := 1; i < len(equipe); i++ {
		if equipe[i].attaque > plusDAttaque.attaque {
			plusDAttaque = equipe[i]
		}
	}
	return plusDAttaque
}

func calculerVieMoyenne(equipe [6]Soldat) float64 {
	total := 0

	for i := 0; i < len(equipe); i++ {
		total = total + equipe[i].vie
	}
	return float64(total)/float64(len(equipe))
}

func compterFaibles(equipe [6]Soldat) int {
	compteur := 0

	for i := 0; i < len(equipe); i++ {
		if equipe[i].vie < 800 {
			compteur++  // on ajoute au compteur si oui
		}
	}
	return compteur
}

// exo 3 

func attaquerEquipe(equipe *[6]Soldat, degats int) {
	for i := 0; i < len(equipe); i++ {	

		if equipe[i].vie > 0 {
			equipe[i].vie = equipe[i].vie - degats

			if equipe[i].vie < 0 {
				equipe[i].vie = 0
			}

			if equipe[i].vie == 0 {
				println(equipe[i].nom, "est KO comme Jordan Zebo!")
			}
		}
	}
}

// exo 4

func afficherEtat(equipe [6]Soldat) {
	println("=== ÉTAT DE L'ÉQUIPE ===")
	println()

	for i := 0; i < len(equipe); i++ {
		if equipe[i].vie > 0 {
			println(equipe[i].nom, ":", equipe[i].vie, "PV")
		} else {
			println(equipe[i].nom, ": KO")
		}
	}
}
