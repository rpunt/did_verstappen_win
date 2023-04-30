package main

import (
	"fmt"

	"github.com/rpunt/f1apireader"
)

func main() {
	race, err := f1apireader.RaceResults()
	if err != nil {
		fmt.Println("err:", err)
	}

	raceStatus, err := race.Status()
	if err != nil {
		fmt.Println("err:", err)
	}
	raceWinner, err := race.Winner()
	if err != nil {
		fmt.Println("err:", err)
	}

	if raceStatus == "completed" {
		if raceWinner.DriverTLA == "VER" {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else {
		fmt.Printf("race is %s\n", raceStatus)
	}
}
