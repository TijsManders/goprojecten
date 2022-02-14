package Planeten

import (
	"fmt"
	"math"
)

func Snelheid() {
	var straal float64
	var uren float64

	fmt.Println("Wat is de afstand tussen de planeet en de ster in kilometers")
	fmt.Scanln(&straal)

	fmt.Println("Vul in hoeveel uren dit duurt")
	fmt.Scanln(&uren)

	km := straal * 2 * math.Pi
	fmt.Println("De planeet gaat", km/uren, "km/h")
}
