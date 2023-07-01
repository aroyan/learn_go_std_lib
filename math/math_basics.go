package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	x := 10.0

	fmt.Println(math.Pi)
	fmt.Println(math.Phi)

	fmt.Println(math.Ceil(math.Pi))
	fmt.Println(math.Floor(math.Pi))

	// Modulo operator for float type
	fmt.Println(math.Mod(18.5, 4.5))
	fmt.Println(math.Mod(18.0, 4.0))

	fmt.Println(math.Round(3.14))
	fmt.Println(math.Round(-3.14))

	fmt.Println(math.RoundToEven(3.14))
	fmt.Println(math.RoundToEven(-3.14))

	// Trigonometry
	fmt.Println(math.Pow(x, 2.0))

	// Exponent
	fmt.Println(math.E, math.Exp(x))

	// Trigonometry
	fmt.Println(math.Cos(math.Pi))
	fmt.Println(math.Sin(math.Pi / 2))
	fmt.Println(math.Tan(math.Pi / 2))

	fmt.Println(math.Log(10))

	fmt.Println(math.Sqrt(49))
	fmt.Println(math.Sqrt(64))
	fmt.Println(math.Sqrt(81))

	fmt.Println(math.Cbrt(125))

	fmt.Println(math.Hypot(30, 40))

	// Random Numbers
	rand.NewSource(time.Now().UnixNano())

	min := 10
	max := 30

	fmt.Println(rand.Intn(max-min) + min)

}
