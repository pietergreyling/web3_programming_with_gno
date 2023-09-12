// guess.gno (package)

// see: https://docs.onbloc.xyz/docs/building-a-realm/realm-examples/hello.gno

package guess

func IsGuessCorrect(guess int, secret int, min int, max int) int {
	var correct int = 0

	// fmt.Println("-- Your guess is: ", guess)

	if guess > secret {
		correct = 1
		// fmt.Println("-- That is higher than the secret number.")
		// fmt.Println("-- Please try again...")
	} else if guess < secret {
		correct = -1
		// fmt.Println("-- That is lower than the secret number.")
		// fmt.Println("-- Please try again...")
	} else {
		correct = 0
		// fmt.Println("-- That is the correct number, congratulations!")
	}

	return correct
}
