package main

import (
	"fmt"
	"math"
)

const s string = "constants"

func main() {
	values()

	variables()

	constants()

	forLoops()
}

// Loops
func forLoops() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// constants
func constants() {
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(math.Sin(n))
}

// variables
func variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

// values
func values() {
	fmt.Println("Go" + "lang")

	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
