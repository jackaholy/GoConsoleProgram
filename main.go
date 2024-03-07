package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

var scores []int

func newScore() string {
	for {
		fmt.Println("Enter a new score: ")
		var score int
		_, err := fmt.Scan(&score)
		if err != nil {
			fmt.Println("\nInvalid input. Please enter an integer.")
		} else {
			scores = append(scores, score)
			// sort the list of scores
			sort.Ints(scores)
			return "Added " + strconv.Itoa(score) + " to list of scores."
		}
	}
}

func pastScore() {
	if len(scores) == 0 {
		fmt.Println("\nThere are currently no scores entered.")
	} else {
		fmt.Println("Past Scores: ")
		for i, score := range scores {
			fmt.Printf("%d: %d\n", i+1, score)
		}
	}
}

func lowestToHighest() {

}

func getMean() {
	if len(scores) == 0 {
		fmt.Println("\nThere are currently no scores entered.")
		return
	}
	sum := 0
	for _, score := range scores {
		sum += score
	}
	mean := float64(sum) / float64(len(scores))
	fmt.Printf("Mean: %.2f\n", mean)
}

func main() {
	for {
		var input1 string
		fmt.Println("\nScore Tracker:")
		fmt.Println("Press [n] to enter a NEW score\n" +
			"Press [p] to view PAST scores\n" +
			"Press [m] to see the MEAN of the scores\n" +
			"Press [q] to QUIT the program")
		fmt.Scan(&input1)
		switch input1 {
		case "n":
			fmt.Println(newScore())
		case "p":
			pastScore()
		case "m":
			getMean()
		case "q":
			fmt.Println("Quitting program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid input. Please try again")
		}
	}
}
