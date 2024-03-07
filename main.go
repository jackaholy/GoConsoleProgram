/*
Author: Jack Holy
Description: A Go program which allows the user to keep track of a scores.
The user has the option to add new scores, look at the past scores, check
the mean (average) of all the scores, and quit the program. Any whole number
can be entered as a score, and strings are not valid.
*/
package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Create an empty array of scores.
var scores []int

/*
Allows the user to enter a new score. Allows type int as valid input,
and adds it to the array.
@return string
*/
func newScore() string {
	for {
		fmt.Println("Enter a new score: ")
		var score int
		// Get input from the user
		_, err := fmt.Scan(&score)
		// Catch errors that are passed into the array of scores.
		if err != nil {
			fmt.Println("\nInvalid input. Please enter an integer.")
		} else {
			scores = append(scores, score)
			// Sort the array of scores
			sort.Ints(scores)
			return "Added " + strconv.Itoa(score) + " to list of scores."
		}
	}
}

/*
Allow the user to look at the scores stored in the array. Iterate through
the list printing each entry.
*/
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

/*
Get the average from the array of scores.
*/
func getMean() {
	// Check to see if any scores are present.
	if len(scores) == 0 {
		fmt.Println("\nThere are currently no scores entered.")
		return
	}
	// Get the sum of the scores
	sum := 0
	for _, score := range scores {
		sum += score
	}
	// Calculate the mean
	mean := float64(sum) / float64(len(scores))
	fmt.Printf("Mean: %.2f\n", mean)
}

/*
Runs the code. Gives the user an option to choose between what they want to
do with the program.
*/
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
