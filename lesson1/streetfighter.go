package main

import (
	"flag"
	"fmt"
)

const GameName = "Street Fighter"

type Character struct {
	Name        string
	SuperSpeech string
	Health      int
	Power       int
}

func doSuperMove(char *Character) {
	fmt.Println(char.SuperSpeech)
	char.Power -= 10
}

func main() {

	playerOne := &Character{
		"Ryu",
		"Hadouken!",
		100,
		50,
	}

	playerTwo := &Character{
		"Sagat",
		"Tiger Uppercut!",
		100,
		50,
	}

	fmt.Println("Welcome to", GameName)

	var fightDuration int
	flag.IntVar(&fightDuration, "fight_duration", 99, "The duration of the fight in seconds")

	fmt.Println("Fight!")
	defer fmt.Println("Fight Over!")

	for i := playerOne.Power; i > 0; i -= 10 {
		doSuperMove(playerOne)
	}

	fmt.Println(playerOne.Name, "has no power left...")

	for i := playerTwo.Power; i > 0; i -= 10 {
		doSuperMove(playerTwo)
	}

	fmt.Println(playerTwo.Name, "has no power left...")

}
