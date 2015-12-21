package main

import "fmt"

const PRESENTS = 29000000

func partOne() int {
	H := make([]int, PRESENTS/10)
	for e := 1; e < PRESENTS/10; e++ {
		for h := e; h < PRESENTS/10; h += e {
			H[h] += e * 10
		}
	}
	var house int
	for i, v := range H {
		house = i
		if v >= PRESENTS {
			break
		}
	}
	return house
}

func partTwo() int {
	H := make([]int, PRESENTS/10)
	for e := 1; e < PRESENTS/10; e++ {
		numHouse := 0
		for h := e; h < PRESENTS/10; h += e {
			H[h] += e * 11
			numHouse++
			if numHouse > 50 {
				break
			}
		}
	}
	var house int
	for i, v := range H {
		house = i
		if v >= PRESENTS {
			break
		}
	}
	return house
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
