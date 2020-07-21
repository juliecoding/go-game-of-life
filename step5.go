package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	height = 8
	width = 5
)

type Universe [][]bool  // "Use slices rather than arrays so that a universe
						// "can be shared with, and modified by, functions or methods."
						// (without using pointers)

func NewUniverse() Universe {
			  //type	//length & capacity
	u := make(Universe, height)
	for i := range u {				// for i := 0; i < len(u); i++ {
		u[i] = make([]bool, width)	// set each element in u equal to a slice of booleans,
									// with a length of width
	}
	return u
}

// add a Show method to universe
// Go doesn't have classes, but you can add methods to types.
// You can associate the method with the type by using a special argument called a receiver.
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
			if r == 1 {
				u[i][j] = true
			} else {
				u[i][j] = false
			}
		}
	}
}

func (u Universe) Alive(row, col int) bool {
	// "A complication arises when the cell is outside of the universe.
	// "Is (-1, -1) dead or alive?"
	// Solution: Wrap around!

	//	If y is less than 0, add the height to it.
	// "If y exceeds the height of the grid, you can turn to the modulus operator
	// "Use % to divide y by height and keep the remainder. 
	// "The same goes for x and width."
	bufferedRow := row
	bufferedCol := col

	if row > height {
		bufferedRow = row % width
	} else if row < 0 {
		bufferedRow = row + width
	}
	if col > width {
		bufferedCol = col % width
	} else if col < 0 {
		bufferedCol = col + width
	}


	return u[bufferedRow][bufferedCol]
}

func (u Universe) CountAlives() int {
	count := 0
	for i := range(u) {
		for j := range(u[i]) {
			if u.Alive(i, j) {
				count++	
			}
		}
	}
	return count
}

func (u Universe) Neighbors(row, col int) int {
	count := 0
	for i := row - 1; i <= row + 1; i++ {
		for j := col - 1; j <= col + 1; j++ {
			if i == row && j == col {
				continue;
			}
			if u.Alive(bufferedRow, bufferedCol) {
				count += 1
			}
		}
	}
	return count
}

func (u Universe) Next(row, col int) bool {
	// A live cell with fewer than 2 live neighbors dies (isolation)
	// A live cell with 2 or 3 live neighbors lives on (community!)
	// A live cell with more than 3 live neighbors dies (overcrowding) 
	// A dead cell with exactly three live neighbors becomes a live cell
	isAlive := u.Alive(row, col)
	livingNeighbors := u.Neighbors(row, col)
	var livesOn bool

	if !isAlive {
		if livingNeighbors == 3 {
			livesOn = true
		}
	} else {
		switch {
			case livingNeighbors < 2:
				livesOn = false
			case livingNeighbors == 2 || livingNeighbors == 3:
				livesOn = true
			case livingNeighbors > 3:
				livesOn = false
		}
	}

	return livesOn
}

func Step(a, b Universe) {
	for i := range a {
		for j := range a[i] {
			b[i][j] = a.Next(i, j)
		}
	}
	a, b = b, a
}

func main() {
	rand.Seed(time.Now().UnixNano())
	u := NewUniverse()
	u.Seed()
	u.Show()
}
