package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing Game")
	fmt.Println("A number will be drawn. Try to guess it. The number is an integer between 1 and 100.")

	drawnNumber := rand.Intn(101)
	scanner := bufio.NewScanner(os.Stdin)
	attempts := make([]int, 0, 10)

	for attempt := range make([]struct{}, 10) {
		fmt.Println("What is your guess?")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		guessInt, err := strconv.Atoi(guess)

		if err != nil {
			fmt.Println("Your guess must be an integer")
			attempt--
			continue
		}

		attempts = append(attempts, guessInt)

		if guessInt < drawnNumber {
			fmt.Println("You missed. The drawn number is greater than", guessInt)
		} else if guessInt > drawnNumber {
			fmt.Println("You missed. The drawn number is less than", guessInt)
		} else {
			fmt.Printf(
				"Congratulations! You guessed it! The number was: %d\n"+
					"You guessed it in %d attempts\n"+
					"These were your attempts: %v\n",
				drawnNumber, attempt+1, attempts,
			)
			return
		}
	}

	fmt.Printf(
		"Unfortunately, you did not guess the number, which was: %d\n"+
			"You had 10 attempts.\n"+
			"These were your attempts: %v\n",
		drawnNumber, attempts,
	)
}
