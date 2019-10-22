package main

import (
	"fmt"
	"math"
	"os"
)

func main()  {
	var a, v0, s0, t float64
	fmt.Println("Please, enter a value for acceleration:")
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Please, enter a value for initial velocity:")
	if _, err := fmt.Scan(&v0); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Please, enter a value for initial position:")
	if _, err := fmt.Scan(&s0); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fn := GenDisplaceFn(a, v0, s0)

	fmt.Println("Great! now please enter a value for time:")
	fmt.Scan(&t)

	fmt.Printf("The displacement after %.1f seconds is %4.4f", t, fn(t))
}

// GenDisplaceFn recives arguments acceleration, initial velocity and initial displacement
func GenDisplaceFn(a, v0, s0 float64) func (float64) float64 {
	fn := func (time float64) float64 {
		return a/2 * math.Pow(time, 2) + v0 * time + s0
	}
	return fn
}