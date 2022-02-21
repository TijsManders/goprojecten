package Hoogbouw

import (
	"bufio"
	"fmt"
	"os"
)

func Hoogteverschil() {
	// Geef 3 gebouwen met hoogtes op
	fmt.Println("Wat is de naam van het eerste gebouw?")
	// bufio.newscanner is maar 1 keer nodig?
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	NaamGebouw1 := input.Text()
	fmt.Println("Wat is de hoogte van", NaamGebouw1, "in meters?")
	input.Scan()
	HoogteGebouw1 := input.Text()

	fmt.Println("Wat is de naam van het tweede gebouw?")
	input.Scan()
	NaamGebouw2 := input.Text()
	fmt.Println("Wat is de hoogte van", NaamGebouw2, "in meters?")
	input.Scan()
	HoogteGebouw2 := input.Text()

	fmt.Println("Wat is de naam van het derde gebouw?")
	input.Scan()
	NaamGebouw3 := input.Text()
	fmt.Println("Wat is de hoogte van", NaamGebouw3, "in meters?")
	input.Scan()
	HoogteGebouw3 := input.Text()

	fmt.Println("De namen van de eerste 3 gebouwen zijn:", NaamGebouw1, NaamGebouw2, NaamGebouw3, "En de hoogte hiervan zijn:", HoogteGebouw1, HoogteGebouw2, HoogteGebouw3)

	//fmt.Println("Wil je nog meer gebouwen toevoegen?")

	// Rekensom

	// Uitslag wegschrijven naar tekstbestand
}
