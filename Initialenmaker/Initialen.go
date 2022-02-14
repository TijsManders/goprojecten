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

	VolleNamen := strings.ReplaceAll(Naam, " der ", " ")
	VolleNamen1 := strings.ReplaceAll(VolleNamen, " de ", " ")
	VolleNamen2 := strings.ReplaceAll(VolleNamen1, " van ", " ")
	VolleNamen3 := strings.ReplaceAll(VolleNamen2, " den ", " ")

	Namen := strings.Split(VolleNamen3, " ")

	for index := range Namen {
		Letter := Namen[index][0:1]
		fmt.Print(strings.ToUpper(Letter) + ".")

	}
}

/* func split(s string, sep rune) []string {
    var a []string
    var j int
    for i, r := range s {
        if r == sep {
            a = append(a, s[j:i+1])
            j = i + 1
        }
    }
    a = append(a, s[j:])
    return a
}
*/
