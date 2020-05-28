package main

import (
	"fmt"
	"time"
)

var (
	aika, saa, vesi, kala int
	aamu                  = string("Kirkkaat värit. Tummemmat sävyt valon vähetessä.")
	ilta                  = string("Tummemmat sävyt valon vähetessä. Aggressiiviset ja arsyttävät värisävyt.")
	kirkas                = []string{"Hopea", "Kulta", "Keltainen", "Vihreä"}
	samea                 = []string{"Punainen", "Ruskea"}
	aurinkoinen           = []string{"Sininen", "Keltainen", "Ruskea"}
	pilvinen              = []string{"Mattavärit", "Punainen", "Vihreä", "Violetti"}
	hauki                 = string("Sopiva vieheen koko on 10-14cm")
	ahven                 = string("Sopiva vieheen koko on 2.5-7cm")
	kuha                  = string("Sopiva vieheen koko on 7-13cm")
	lohi                  = string("Sopiva vieheen koko on 9-13cm")
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
	fmt.Println("Menetkö kalaan nyt:")
	fmt.Println("[1] Kyllä")
	fmt.Println("[2] Ei")
	fmt.Scan(&aika)
	if aika > 2 {
		fmt.Println("Valinta väärin. Valitse uudelleen. [0] ohittaa ")
		selvitaAika()
	}
	if aika == 1 {
		t := time.Now()
		switch {
		case t.Hour() < 12:
			aika = 1
			return
		case t.Hour() > 12:
			aika = 2
			return
		}
	}

	fmt.Println("Mihin aikaan kalastat:")
	fmt.Println("[1] Ennen puoltapäivää:")
	fmt.Println("[2] Puolenpäivän jälkeen:")
	fmt.Scan(&aika)
	if aika > 2 {
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
	//SELVITETÄÄN OLOSUHTEET//
	selvitaAika()
	selvitaVesi()
	selvitaSaa()
	selvitaKala()

	//PRINTATAAN YLLÄ SAADUT TULOKSET//
	switch vesi {
	case 1:
		fmt.Println("Kirkkaaseen veteen sopivat värit:", kirkas)
		if saa == 1 {
			fmt.Println("Aurinkoiseen säähän sopivat värit:", aurinkoinen)
			if aika == 1 {
				fmt.Println("Ennen puoltapäivää:", aamu)
			}
			if aika == 2 {
				fmt.Println("Puolenpäivän jälkeen:", ilta)
			}
		}
		if saa == 2 {
			fmt.Println("Pilviseen säähän sopivat värit:", pilvinen)
			if aika == 1 {
				fmt.Println("Ennen puoltapäivää:", aamu)
			}
			if aika == 2 {
				fmt.Println("Puolenpäivän jälkeen:", ilta)
			}
		}

	case 2:
		fmt.Println("Sameaan veteen sopivat värit:", samea)
		if saa == 1 {
			fmt.Println("Aurinkoiseen säähän sopivat värit:", aurinkoinen)
			if aika == 1 {
				fmt.Println("Ennen puoltapäivää:", aamu)
			}
			if aika == 2 {
				fmt.Println("Puolenpäivän jälkeen:", ilta)
			}
		}
		if saa == 2 {
			fmt.Println("Pilviseen säähän sopivat värit:", pilvinen)
			if aika == 1 {
				fmt.Println("Ennen puoltapäivää:", aamu)
			}
			if aika == 2 {
				fmt.Println("Puolenpäivän jälkeen", ilta)
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
