package main

import (
	"fmt"
	"os"
	"strconv"
)

type ScoreSheet struct {
	scores       []int
	numScores    int
	highestScore int
	lowestScore  int
	mean         float32
}

func newScore(sheet *ScoreSheet) string {
	for {
		fmt.Println("\nEnter a new score: ")
		var score int
		_, err := fmt.Scanln(&score)
		if err != nil {
			return "Invalid input. Please enter an integer."
		} else {
			sheet.scores = append(sheet.scores, score)
			pastScore(sheet)
			return "Added " + strconv.Itoa(score) + " to list."
		}
	}
}

func pastScore(sheet *ScoreSheet) {
	fmt.Println("Past Scores: ")
	for i, score := range sheet.scores {
		fmt.Printf("%d: %d\n", i+1, score)
	}
}

func getMean(scores []int) float64 {
	sum := 0.0
	for score := 0; score < len(scores); score++ {
		sum += float64(score)
	}
	mean := sum / float64(len(scores))
	return mean
}

func main() {
	for {
		// A blank sheet to store the input1 scores.
		sheet1 := &ScoreSheet{
			scores:       []int{},
			numScores:    0,
			highestScore: 0,
			lowestScore:  0,
			mean:         0}
		var input1 string
		fmt.Println("\nWelcome to your score keeper")
		fmt.Println("Press [n] to enter a New score\n" +
			"Press [p] to view past scores\n" +
			"Press [q] to quit the program")
		fmt.Scanln(&input1)
		switch input1 {
		case "n":
			fmt.Println(newScore(sheet1))
		case "p":
			pastScore(sheet1)
		case "q":
			fmt.Println("Quitting program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid input. Please try again")
		}
	}
}
