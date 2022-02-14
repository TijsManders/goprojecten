package main

import (
	"fmt"
)

func main() {

	var optie string

	const menu = "Kies 1 voor Hello World \n" +
		"Kies 2 voor Alarm\n" +
		"Kies 3 voor Planeten\n"
	fmt.Println(menu)

	fmt.Scanln(&optie)
	switch optie {

	default:
		fmt.Println("Je hebt niet voor een valide optie gekozen")

	case "1":

	}
}
