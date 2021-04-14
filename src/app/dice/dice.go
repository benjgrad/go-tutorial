package dice

import (
	"math/rand"
)

//Multiple return values
func Roll() (int, int) {
	//var declares variables without initializing them
	var x, y int
	x, y = rand.Intn(6), rand.Intn(6)
	return x, y
}
