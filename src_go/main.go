package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	min, max := 1, 100

	rand.Seed(time.Now().UnixNano()) // otherwise we will get the same secret
	secretNumber := rand.Intn(max-min) + min

	tries := 0

	fmt.Println("-- Please guess a number between 1 and 100.")
	fmt.Println("-- Input your guess.")

	for {
		tries++
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("-- Error when reading input!", err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("-- Invalid input - that is not a number! Please enter a number.")
			continue
		}

		fmt.Println("-- Your guess is: ", guess)

		if guess > secretNumber {
			fmt.Println("-- That is higher than the secret number.")
			fmt.Println("-- Please try again...")
		} else if guess < secretNumber {
			fmt.Println("-- That is lower than the secret number.")
			fmt.Println("-- Please try again...")
		} else {
			fmt.Println("-- That is the correct guess after", tries, "tries!")
			break
		}
	}
}
