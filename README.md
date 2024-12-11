package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// generazione di un numero casuale compreso nell'intervallo che va da 0 a 99 (estremi inclusi)
	var numeroGenerato int = rand.Intn(100)

	var n int
	var count int
	for {

		count++
		fmt.Printf("Tentativo numero %d: ", count)
		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Println("Inserisci un valore valido.\nExit status...")
			break
		}
		if n < numeroGenerato {
			fmt.Println("Troppo basso! Riprova!")
			continue
		} else if n > numeroGenerato {
			fmt.Println("Troppo alto! Riprova!")
			continue
		}
		if n == numeroGenerato {
			fmt.Println("Hai indovinato in ", count, "tentativi!")
			break
		}
	}
}
