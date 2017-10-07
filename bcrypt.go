package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "passwort123"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil { // Error-Handling
		fmt.Println(err)
	}
	fmt.Println("Klartext:", s)     // gibt noch einmal den zu hash-enden String aus
	fmt.Println("Bcrypt-Hash:", bs) // gibt den String in ge-hash-ter Form aus

	// Vergleich zwischen ge-hash-ten und unge-hash-ten Strings:

	testpasswd := "passwort123"
	testhash := []byte{36, 50, 97, 36, 49, 48, 36, 71, 50, 111, 105, 120, 84, 79, 52, 71, 50, 54, 102, 104, 87, 117, 105, 49, 114, 75, 85, 65, 117, 98, 118, 71, 109, 114, 48, 53, 87, 82, 78, 66, 84, 57, 80, 82, 67, 74, 82, 54, 67, 113, 68, 82, 102, 71, 111, 104, 106, 75, 82, 87}
	err = bcrypt.CompareHashAndPassword(testhash, []byte(testpasswd))
	if err != nil { // Wenn es nicht das richtige Passwort war
		fmt.Println("Falsches Passwort!")
		return
	}
	fmt.Println("Richtiges Passwort!")

	// Cost-Parameter ermitteln
	/* 	Anmerkung:
		*	Neuere Passwort-Hashing-Funktionen haben einen Cost-Paramter.append
		*	Dieser gibt an, wieviel "Arbeit" (z.B. CPU-Zeit) aufgewendet wurde um diesen Hash zu erstellen
		*	Definition:
		*	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	    *	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	    *	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
	*/
	c, err := bcrypt.Cost(testhash)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Cost:", c)

}
