package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	totalRows = 3
	totalColumns = 5
)

type Universe [][]bool  // "Use slices rather than arrays so that a universe
						// "can be shared with, and modified by, functions or methods."
						// (without using pointers)

func NewUniverse() Universe {
			  //type	//length & capacity
	u := make(Universe, totalRows)
	for row := range u {					// for i := 0; i < len(u); i++ {
		u[row] = make([]bool, totalColumns)	// set each element in u equal to a slice of booleans,
											// with a length of totalColumns
	}
	return u
}

// add a Show method to universe
// Go doesn't have classes, but you can add methods to types.
// You can associate the method with the type by using a receiver.
func (u Universe) Show() {
	str := ""
	for row := range u {
		for col := range u[row] {
			if (u[row][col]) {
				str += "*"
			} else {
				str += "-"
			}
		}
		str += "\n"
	}

	// fmt.Print("\x0c")
	fmt.Print("\033[2J", str)
}

func (u Universe) Seed() {
	for row := range u {
		for col := range u[row] {
			r := 1 + rand.Intn(4) // create a random integer, minimum 1 and maximum 4
			if r == 1 {
				u[row][col] = true
			} else {
				u[row][col] = false
			}
		}
	}
}

// func (u Universe) Seed2() {
// 	for i := 0; i < (width * height / 4); i++ {
// 		u.Set(rand.Intn(width), rand.Intn(height), true)
// 	}
// }

func (u Universe) Set(col, row int, b bool) {
	u[row][col] = b
}

func (u Universe) Alive(col, row int) bool {
	// "A complication arises when the cell is outside of the universe.
	// "Is (-1, -1) dead or alive?"
	// Solution: Wrap around!

	//	If y is less than 0, add the height to it.
	// "If y exceeds the height of the grid, you can turn to the modulus operator
	// "Use % to divide y by height and keep the remainder. 
	// "The same goes for x and width."
	col = (col + totalColumns) % totalColumns  	// Why does this work?
	row = (row + totalRows) % totalRows			// Think if you had 100 rows, and you were passed 150
												// 150 + 100 = 250 
												// 250 % 100 = 50 

	return u[row][col]
}

func (u Universe) CountAlives() int {
	count := 0
	for row := range(u) {
		for col := range(u[row]) {
			if u.Alive(col, row) {
				count++	
			}
		}
	}
	return count
}

func (u Universe) Neighbors(col, row int) int {
	count := 0
	for i := row - 1; i <= row + 1; i++ {
		for j := col - 1; j <= col + 1; j++ {
			if i == row && j == col {
				continue;
			}
			if u.Alive(j, i) {
				count += 1
			}
		}
	}
	return count
}

// func (u Universe) Next1(col, row int) bool {
// 	// A live cell with fewer than 2 live neighbors dies (isolation)
// 	// A live cell with 2 or 3 live neighbors lives on (community!)
// 	// A live cell with more than 3 live neighbors dies (overcrowding) 
// 	// A dead cell with exactly three live neighbors becomes a live cell (birth)
// 	isAlive := u.Alive(col, row)
// 	livingNeighbors := u.Neighbors(col, row)
// 	var livesOn bool

// 	if !isAlive {
// 		if livingNeighbors == 3 {
// 			livesOn = true
// 		}
// 	} else {
// 		switch {
// 			case livingNeighbors < 2:
// 				livesOn = false
// 			case livingNeighbors == 2 || livingNeighbors == 3:
// 				livesOn = true
// 			case livingNeighbors > 3:
// 				livesOn = false
// 		}
// 	}

// 	return livesOn
// }

func (u Universe) Next(col, row int) bool {
	n := u.Neighbors(col, row)
	return n == 3 || n == 2 && u.Alive(col, row)
}

// "Update the state of the next universe (b)
// "From the current universe (a)."
func Step(a, b Universe) {	
	for row := 0; row < totalRows; row++ {
		for col := 0; col < totalColumns; col++ {
			b.Set(col, row, a.Next(col, row))
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a, b := NewUniverse(), NewUniverse()
	a.Seed()
	for i := 0; i < 5; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second) // We use the Sleep function from the `time` package to slow down the animation
		a, b = b, a
	}
}
