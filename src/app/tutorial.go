package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/benjgrad/go-tutorial/src/app/dice"
)

func main() {
	//import and random seeding
	rand.Seed(time.Now().UnixNano())

	//Initializing and setting multiple values
	//:= is a variable declaration shorthand
	a, b := dice.Roll()

	//Print with formatting
	fmt.Printf("Dice 1: %v Dice 2: %v Roll value: %v\n", a, b, Add(a, b))

}
