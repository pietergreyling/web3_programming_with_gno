# How to get started with Web3 Programming using Gno

This is a short tutorial on how to get started with Web3 programming using Gnolang and Gnoland. 

You only need some basic Go programming language knowledge. 

Gnolang is a version of Golang but it is built for Web3 programming. 

This means that you can directly transfer your Go skills to Web3 development.

## Introductory Gnolang Concepts

### Realms

A Gnolang Realm is where the state of an application lives.

Realms are the equivalent of Smart Contracts.

- https://docs.onbloc.xyz/docs/building-a-realm


### Packages

Gnolang Packages hold no state, only functionality.

## The Example Gno Project

The example project for this tutorial is a simple number guessing game.

In the Go programming language, the essential logic for a command line guessing game program looks as follows.

```golang

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
    tries := 0
	rand.Seed(time.Now().UnixNano()) // otherwise we will get the same secret
	secretNumber := rand.Intn(max-min) + min

	fmt.Println("Guess a number between 1 and 100")
	fmt.Println("Input your guess")

	for {
		tries++
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error when reading input", err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input - not a number. Please enter a number")
			continue
		}

		fmt.Println("Your guess is", guess)

		if guess > secretNumber {
			fmt.Println("Higher than the secret number. Try again")
		} else if guess < secretNumber {
			fmt.Println("Lower than the secret number. Try again")
		} else {
			fmt.Println("Correct guess after", tries, "tries!")
			break
		}
	}
}

```

## References

Jae Kwon: Gnoland the Inevitable Next Generation Smart Contract Platform

- https://youtu.be/IJ0xel8lr4c?list=PLc_cC_y0BSHymUiUnIoZPeg_BpYwitw8V

GNO Example Demo

- https://youtu.be/-BlnEXCs0eI

Go to Gno: How to build a microblog

- https://youtu.be/F-_dadxcRJM?list=PLc_cC_y0BSHymUiUnIoZPeg_BpYwitw8V

Quick Intro to Gno - Manfred Touron

- https://youtu.be/bA-lMd60Lts?list=PLc_cC_y0BSHymUiUnIoZPeg_BpYwitw8V

Gnoland on YouTube

- https://www.youtube.com/@_gnoland

Gno on GitHub

- https://github.com/gnolang/gno

Quickstart Guide

- https://github.com/gnolang/gno/blob/master/examples/gno.land/r/demo/boards/README.md

Gnoland Developer Portal

- https://docs.onbloc.xyz

Awesome Gno

- https://github.com/gnolang/awesome-gno

Awesome Gno Tutorials

- https://github.com/gnolang/awesome-gno#tutorials

Gno By Example

- https://gno-by-example.com

Gnolang Examples

- https://github.com/gnolang/gno/tree/master/examples

