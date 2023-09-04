// Set up the module:
// go mod init github.com/pietergreyling/web3_programming_with_gno/m/v2

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

// Example session:
// go run .
// -- Please guess a number between 1 and 100.
// 50
// -- Your guess is:  50
// -- That is lower than the secret number.
// -- Please try again...
// 75
// -- Your guess is:  75
// -- That is higher than the secret number.
// -- Please try again...
// 62
// -- Your guess is:  62
// -- That is lower than the secret number.
// -- Please try again...
// 68
// -- Your guess is:  68
// -- That is higher than the secret number.
// -- Please try again...
// 65
// -- Your guess is:  65
// -- That is lower than the secret number.
// -- Please try again...
// 67
// -- Your guess is:  67
// -- That is the correct number, congratulations!
// -- You tried 6 times!

// Main holds the global state in this program - - - -
func main() {

	min, max := 1, 100
	tries := 0
	correct := false

	rand.Seed(time.Now().UnixNano()) // otherwise we will get the same secret
	secret := rand.Intn(max-min) + min

	fmt.Println("-- Please guess a number between 1 and 100.")

	reader := bufio.NewReader(os.Stdin)

	for {
		tries++

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("-- Error when reading input!", err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("-- Invalid input - that is not a number!")
			fmt.Println("-- Please enter a number.")
			continue
		}

		correct = TestGuess(guess, secret, min, max)
		if correct {
			break
		}
	}

	fmt.Println("-- You tried", tries, "times!")
}

// No global state is held here - - - -
func TestGuess(guess int, secret int, min int, max int) bool {

	correct := false

	fmt.Println("-- Your guess is: ", guess)

	if guess > secret {
		correct = false
		fmt.Println("-- That is higher than the secret number.")
		fmt.Println("-- Please try again...")
	} else if guess < secret {
		correct = false
		fmt.Println("-- That is lower than the secret number.")
		fmt.Println("-- Please try again...")
	} else {
		correct = true
		fmt.Println("-- That is the correct number, congratulations!")
	}

	return correct
}
