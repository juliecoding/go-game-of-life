package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	height = 10
	width = 5
)

type Universe [][]bool  // Use slices rather than arrays so that a universe
						// can be shared with, and modified by, functions or methods
						// (without using pointers)

func NewUniverse() Universe {
			  //type	//length & capacity
	u := make(Universe, height)
	for i := range u {				// for i := 0; i < len(u); i++ {
		u[i] = make([]bool, width)	// set each element in u equal to a one-dimensional array of booleans,
									// with a length of width
	}
	return u
}

// add a show method to universe
// Go doesn't have classes, but you can add methods to types.
// you can associate the method with the type by using a special argument called a receiver
func (u Universe) Show() {
	str := ""
	for i := 0; i < len(u); i++ {
		for j := 0; j < len(u[i]); j++ {
			if (u[i][j]) {
				str += "*"
			} else {
				str += "-"
			}
		}
		str += "\n"
	}

	fmt.Println(str)
}

func (u Universe) Seed() {
	for i := range u {
		for j := range u[i] {
			r := 1 + rand.Intn(4) // create a random integer, minimum 1 and maximum 4
			if (r == 1) {
				u[i][j] = true
			} else {
				u[i][j] = false
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	u := NewUniverse()
	u.Show()
	u.Seed()
	u.Show()
}