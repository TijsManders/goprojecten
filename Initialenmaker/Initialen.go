package Initialenmaker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Naam string

func Initialen() {

	fmt.Println("Vul hier je volledige naam in")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	Naam = scanner.Text()
	Namen := strings.Split(Naam, " ")
	fmt.Println(Namen[0])
	fmt.Println(Namen[1])
}
