package Alarm

import (
	"fmt"
	"time"
)

func Alarm() {
	var AantalAlarm int

	fmt.Println("Hoe vaak moet het alarm af gaan?")

	fmt.Scanln(&AantalAlarm)

	fmt.Println("Je alarm gaat", AantalAlarm, "keer af")

	for i := 1; i <= AantalAlarm; i++ {
		fmt.Println("Alarm", i, "!")
		time.Sleep(1000 * time.Millisecond)
	}

}
