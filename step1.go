package main

import "fmt"

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
		u[i] = make([]bool, width)	// set each element in u equal to a one-dimensional slice of booleans,
									// each with a length of width
	}								// zero state = false
	return u
}

// add a show method to universe
// Go doesn't have classes, but you can add methods to types.
// you can associate the method with the type by using a special argument called a receiver
// For any universe we create, we can call these methods on it.
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

func main() {
	u := NewUniverse()
	u.Show()
}