package main

import (
	"fmt"
	"math/rand"
)

func random_insults() {
	answers := []string{
		"But mom, I dont wanna!",
		"Why should I? You're not the boss of me",
		".....................................................\n\n................................\n\n.........................\n\nSorry did you say something?\n I wasn't paying attention",
		"BAD COMMAND OR FILENAME",
		"Didnt your parents teach you to use please?",
		"Whaaattt? Sorry didnt hear you. Learn to type properly kid, come back when you've grown up. LOL",
	}
	randomIndex := rand.Intn(len(answers))
	pick := answers[randomIndex]
	fmt.Println(pick)

}

func windows_commands() {
	fmt.Println("We dont support Weenie-Doze, loser. Learn a real operating system")
}
