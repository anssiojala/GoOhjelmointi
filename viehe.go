package main

import (
	"fmt"
)

var (
	aika, saa, vesi, kala int
)

func selvitaVesi() {
	fmt.Println("Millainen vesistö")
	fmt.Println("[1] Kirkas vesi")
	fmt.Println("[2] Samea vesi")
	fmt.Scan(&vesi)
	if vesi > 2 {
		fmt.Println("Valinta väärin. Valitse uudelleen. [0] ohittaa ")
		selvitaVesi()
	}
}

func selvitaSaa() {
	fmt.Println("Onko sää nyt:")
	fmt.Println("[1] Aurinkoinen")
	fmt.Println("[2] Pilvinen")
	fmt.Scan(&saa)
	if saa > 2 {
		fmt.Println("Valinta väärin. Valitse uudelleen. [0] ohittaa ")
		selvitaSaa()
	}
}

func selvitaAika() {
	fmt.Println("Onko aika nyt:")
	fmt.Println("[1] Aamu")
	fmt.Println("[2] Päivä")
	fmt.Println("[3] Ilta")
	fmt.Scan(&aika)
	if aika > 3 {
		fmt.Println("Valinta väärin. Valitse uudelleen. [0] ohittaa ")
		selvitaAika()
	}
}

func selvitaKala() {
	fmt.Println("Haetko tiettyä kalaa:")
	fmt.Println("[1] Hauki")
	fmt.Println("[2] Ahven")
	fmt.Println("[3] Kuha")
	fmt.Println("[4] Lohi")
	fmt.Println("[5] Ei tiettyä")
	fmt.Scan(&kala)
	if kala > 5 {
		fmt.Println("Valinta väärin. Valitse uudelleen. [0] ohittaa ")
		selvitaKala()
	}
}

func main() {

	//VASTAUKSET
	kirkas := []string{"Hopea", "Kulta", "Keltainen", "Vihreä"}
	samea := []string{"Punainen", "Ruskea"}
	aurinkoinen := []string{"Sininen", "Keltainen", "Ruskea"}
	pilvinen := []string{"Mattavärit", "Punainen", "Vihreä", "Violetti"}
	aamu := string("Kirkkaat värit.")
	paiva := string("Kirkkaat värit. Tummemmat sävyt valon vähetessä")
	ilta := string("Tummemmat sävyt valon vähetessä. Aggressiiviset ja arsyttävät värisävyt.")
	hauki := string("Sopiva vieheen koko on 10-14cm")
	ahven := string("Sopiva vieheen koko on 2.5-7cm")
	kuha := string("Sopiva vieheen koko on 7-13cm")
	lohi := string("Sopiva vieheen koko on 9-13cm")

	//SELVITETÄÄN OLOSUHTEET//
	selvitaVesi()
	selvitaSaa()
	selvitaAika()
	selvitaKala()

	//PRINTATAAN YLLÄ SAADUT TULOKSET//
	switch vesi {
	case 1:
		fmt.Println("Vesi on kirkasta.", kirkas)
		if saa == 1 {
			fmt.Println("Sää on aurinkoinen.", aurinkoinen)
			if aika == 1 {
				fmt.Println("On aamu.", aamu)
			}
			if aika == 2 {
				fmt.Println("On päivä.", paiva)
			}
			if aika == 3 {
				fmt.Println("On ilta.", ilta)
			}
		}
		if saa == 2 {
			fmt.Println("Sää on pilvinen.", pilvinen)
			if aika == 1 {
				fmt.Println("On aamu.", aamu)
			}
			if aika == 2 {
				fmt.Println("On päivä.", paiva)
			}
			if aika == 3 {
				fmt.Println("On ilta.", ilta)
			}
		}

	case 2:
		fmt.Println("Vesi on sameaa", samea)
		if saa == 1 {
			fmt.Println("Sää on aurinkoinen.", aurinkoinen)
			if aika == 1 {
				fmt.Println("On aamu.", aamu)
			}
			if aika == 2 {
				fmt.Println("On päivä.", paiva)
			}
			if aika == 3 {
				fmt.Println("On ilta.", ilta)
			}
		}
		if saa == 2 {
			fmt.Println("Sää on pilvinen.", pilvinen)
			if aika == 1 {
				fmt.Println("On aamu.", aamu)
			}
			if aika == 2 {
				fmt.Println("On päivä.", paiva)
			}
			if aika == 3 {
				fmt.Println("On ilta.", ilta)
			}
		}

	}

	switch kala {
	case 1:
		fmt.Println(hauki)
	case 2:
		fmt.Println(ahven)
	case 3:
		fmt.Println(kuha)
	case 4:
		fmt.Println(lohi)
	case 5:
		break
	}
}
