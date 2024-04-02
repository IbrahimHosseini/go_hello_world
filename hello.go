package main

import (
	"fmt"
	"math"
	"time"
)

const s string = "constants"

func main() {
	values()

	variables()

	constants()

	forLoops()

	ifElse()

	switchSample()

	array()
}

// Array
func array() {
	var a [5]int
	fmt.Println(a)

	a[3] = 14
	a[1] = 20
	fmt.Println(a)

	fmt.Println("Lenght: ", len(a))

	b := [3]int{1, 2, 3}
	fmt.Println(len(b))

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2D => ", twoD)
}

// switch
func switchSample() {
	switch time.Now().Weekday() {
	case time.Friday:
		fmt.Println("Friday")
	case time.Monday:
		fmt.Println(" Monday")
	case time.Thursday:
		fmt.Println(" Thursday")
	case time.Wednesday:
		fmt.Println(" Wednesday")
	default:
		fmt.Println("Other day")
	}
}

// If/else
func ifElse() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
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
