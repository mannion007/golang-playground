package main

import (
    "fmt"
)

type Character struct {
    Name        string
    SuperSpeech string
    Health      int
    Power       int
}

func DoSuperMove(char *Character) {
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

    fmt.Println("Fight!")
    defer fmt.Println("Fight Over!")

    for i := playerOne.Power; i > 0; i -= 10 {
        DoSuperMove(playerOne)
    }

    fmt.Println(playerOne.Name, "has no power left...")

    for i := playerTwo.Power; i > 0; i -= 10 {
        DoSuperMove(playerTwo)
    }
    
    fmt.Println(playerTwo.Name, "has no power left...")
}
