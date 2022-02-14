package Initialenmaker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Initialen() {

	fmt.Println("Vul hier je volledige naam in")

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	Naam := input.Text()
	Namen := strings.Split(Afkorten(Naam), " ")

	for index := range Namen {
		Letter := Namen[index][0:1]
		fmt.Print(strings.ToUpper(Letter) + ".")
	}
}

func Afkorten(MetTussenvoegsels string) string {

	Tussenvoegsels := []string{" de ", " der ", " den ", " van "}
	for Tussenvoegsel := range Tussenvoegsels {
		MetTussenvoegsels = strings.ReplaceAll(MetTussenvoegsels, Tussenvoegsels[Tussenvoegsel], " ")
	}
	return MetTussenvoegsels
}
