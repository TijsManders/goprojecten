package Hoogbouw

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Hoogteverschil() {
	Gebouwen := []Gebouw{}

	fmt.Println("Hoeveel gebouwen Wil je toevoegen?")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	HoeveelGebouwen, err := strconv.Atoi(input.Text())

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < HoeveelGebouwen; i++ {
		fmt.Println("Wat is de naam van gebouw", i+1, "?")
		input.Scan()
		NaamGebouw := input.Text()
		fmt.Println("Wat is de hoogte van", NaamGebouw, "in meters?")
		input.Scan()
		HoogteGebouw, err := strconv.Atoi(input.Text())
		if err != nil {
			log.Fatal(err)
		}
		NieuwGebouw := Gebouw{
			Naam:   NaamGebouw,
			Hoogte: HoogteGebouw,
		}
		Gebouwen = append(Gebouwen, NieuwGebouw)
	}
	for _, Info := range Gebouwen {
		fmt.Println("Gebouw", Info.Naam, "is", Info.Hoogte, "meter hoog.")
	}

}

// Rekensom
// sort.splice onderzoeken

//fmt.Println(Hoogste1, "is het hoogst, het is ", Verschil1, "meter hoger dan", Hoogste2, "en ", Verschil2, "meter hoger dan ", Hoogste3, ".")

// Uitslag wegschrijven naar tekstbestand

type Gebouw struct {
	Naam   string
	Hoogte int
}
