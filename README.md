# How to get started with Web3 Programming using Gno

This is a short tutorial on how to get started with Web3 programming using Gnolang and Gnoland.

You only need some basic Go programming language knowledge. 

Gnolang is a version of Golang but it is built for Web3 programming. 

This means that you can directly transfer your Go skills to Web3 development.

## Source Code Editing Tools

For editing the code of this project I will be using [Visual Studio Code (VSCode)](https://code.visualstudio.com/).

The following two VSCode extensions are highly recommended:

- [Go in Visual Studio Code](https://code.visualstudio.com/docs/languages/go)

- [Gno for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=harry-hov.gno)

However, please feel free to use any other editing environment that you prefer.

## Introductory Gno

### Overview and Concepts

[What is Gnoland?](https://docs.onbloc.xyz/introduction-to-gnoland/what-is-gnoland)

- [Overview](https://docs.onbloc.xyz/introduction-to-gnoland/what-is-gnoland/overview)

- [Concepts](https://docs.onbloc.xyz/introduction-to-gnoland/what-is-gnoland/concepts/)

### Packages

[Packages](https://docs.onbloc.xyz/docs/packages) encompass functionalities that are more closely aligned with the characteristics and capabilities of realms, as opposed to standard libraries.

- A unit that contains functionalities and utilities that can be used in realms.
- Gnolang [Packages](https://docs.onbloc.xyz/docs/building-a-realm/overview#packages) are stateless.
- The default import path is ```gno.land/p/~~~```.
- Can be imported to other realms or packages.

### Realms

A Gnolang [Realm](https://docs.onbloc.xyz/docs/building-a-realm/overview#realms) is where the state of an application lives.

- Realms are [Smart Contracts](https://docs.onbloc.xyz/docs/building-a-realm) in Gnolang.
- Realms are stateful.
- The default import path is ```gno.land/r/~~~```.
- Each realm has the capability to publicly export the function Render(path string) string, which performs rendering when passed a valid markdown as a parameter for the specified path.

## The Example Go Project that we will port to Gno

The example project for this tutorial is a simple number guessing game.

In the Go programming language, the essential logic for a command line guessing game program looks as follows:

```golang

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

```

Please notice that we have made sure to separate the stateful main function from the stateless number guessing code. This is in order to be in line with the Gnolang concepts of Realms and Packages:

- The stateful main function will become a Realm
- The stateless number guessing code will become a Package

As preparation for porting this to a Gnolang realm and package architecture we run a quick test by executing the following:

```shell
cd src_go
go run .
-- Please guess a number between 1 and 100.
50
-- Your guess is:  50
-- That is lower than the secret number.
-- Please try again...
75
-- Your guess is:  75
-- That is higher than the secret number.
-- Please try again...
62
-- Your guess is:  62
-- That is lower than the secret number.
-- Please try again...
68
-- Your guess is:  68
-- That is higher than the secret number.
-- Please try again...
65
-- Your guess is:  65
-- That is lower than the secret number.
-- Please try again...
67
-- Your guess is:  67
-- That is the correct number, congratulations!
-- You tried 6 times!
```

We have now verified that the program is working as expected.

## References

- [Gnoland on YouTube](https://www.youtube.com/@_gnoland)

  - [Jae Kwon: Gnoland the Inevitable Next Generation Smart Contract Platform](https://youtu.be/IJ0xel8lr4c?list=PLc_cC_y0BSHymUiUnIoZPeg_BpYwitw8V)

  - [GNO Example Demo](https://youtu.be/-BlnEXCs0eI)

  - [Quick Intro to Gno](https://youtu.be/bA-lMd60Lts?list=PLc_cC_y0BSHymUiUnIoZPeg_BpYwitw8V)

  - [Go to Gno: How to build a microblog](https://youtu.be/F-_dadxcRJM?list=PLc_cC_y0BSHymUiUnIoZPeg_BpYwitw8V)

- [Gnolang on GitHub](https://github.com/gnolang/gno)

  - [Gno(Lang) & Gno.land](https://github.com/gnolang)

  - [Quickstart Guide](https://github.com/gnolang/gno/blob/master/examples/gno.land/r/demo/boards/README.md)

  - [Gnoland Developer Portal](https://docs.onbloc.xyz)

  - [Awesome Gno](https://github.com/gnolang/awesome-gno)

  - [Awesome Gno Tutorials](https://github.com/gnolang/awesome-gno#tutorials)

  - [Gnolang Examples](https://github.com/gnolang/gno/tree/master/examples)

- [Gno By Example](https://gno-by-example.com)
