// Package postaje definira postajo in vsebuje funkcije za upravljane s postajami
package vreme

import (
	"fmt"
)

// Postaja predstavlja vremensko postajo z imenom, krajem in meritvami
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo doda novo postajo v map
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{ime: ime, kraj: kraj}
}

// DodajMeritev doda meritev izbrani postaji
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, minTemp, maxTemp float64) {
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Println("neveljavna meritev!")
		return
	}
	e, ok := postaje[oznaka]
	if ok {
		e.meritve = append(e.meritve, temperatura)
		postaje[oznaka] = e
	} else {
		fmt.Println("postaja ne obstaja!")
	}

}

func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	e, obstaja := postaje[oznaka]

	if !obstaja || stMeritev < 1 || len(e.meritve) < stMeritev {
		return 0, 0, 0, false
	}

	najnizja = e.meritve[0]
	najvisja = e.meritve[0]

	var sum float64

	for _, val := range e.meritve {
		sum += val

		if val < najnizja {
			najnizja = val
		}

		if val > najvisja {
			najvisja = val
		}
	}

	povprecna = sum / float64(len(e.meritve))

	return najnizja, najvisja, povprecna, true
}

// IzpisMeritev izpiše vse postaje in njihove meritve
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for key, value := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", key, value.ime, value.kraj, value.meritve)
	}
}

// IzpisPorocila izpiše statistiko za vse postaje
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for key, value := range postaje {
		najnizja, najvisja, povprecna, ok := statistika(postaje, key, stMeritev)

		if ok {
			fmt.Printf("%s: min %f °C, max %f °C, povprečje %f °C -> ", value.ime, najnizja, najvisja, povprecna)

			switch {
			case povprecna >= 25:
				fmt.Printf("Vroče\n")
			case povprecna < 25 && povprecna >= 15:
				fmt.Printf("Toplo\n")
			case povprecna < 15 && povprecna >= 5:
				fmt.Printf("Hladno\n")
			case povprecna < 5:
				fmt.Printf("Mrzlo\n")
			}
		} else {
			fmt.Printf("%s: premalo meritev\n", value.ime)

		}
	}
}
