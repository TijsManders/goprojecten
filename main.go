package main

import (
	"fmt"

	"github.com/TijsManders/goprojecten/Alarm"
	"github.com/TijsManders/goprojecten/HelloWorld"
	"github.com/TijsManders/goprojecten/Planeten"
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
		HelloWorld.HelloWorld()

	case "2":
		Alarm.Alarm()

	case "3":
		Planeten.Snelheid()

	case "4":
		initialenmaker.initialen()
	}
}
