package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	// totalRows = 5
	// totalColumns = 10
	totalRows    = 20 // SWITCH COMMENTS TO SHOW SMALLER GRID
	totalColumns = 50
)

type Universe [][]bool

func NewUniverse() Universe {
	u := make(Universe, totalRows)
	for row := range u {
		u[row] = make([]bool, totalColumns)
	}
	return u
}

func (u Universe) Show() {
	str := ""
	for row := range u {
		for col := range u[row] {
			if u[row][col] {
				str += "*"
			} else {
				// str += "-"                   // SWITCH COMMENTS TO MAKE BLANKS VISIBLE
				str += " "
			}
		}
		str += "\n"
	}

	// fmt.Print("\033c")           // UNCOMMENT TO CLEAR SCREEN EACH GENERATION
	fmt.Print("\n", str)
}

func (u Universe) Seed() {
	for i := 0; i < (totalColumns * totalRows / 4); i++ {
		u.Set(rand.Intn(totalColumns), rand.Intn(totalRows), true)
	}
}

func (u Universe) Set(col, row int, b bool) {
	u[row][col] = b
}

func (u Universe) Alive(col, row int) bool {
	col = (col + totalColumns) % totalColumns
	row = (row + totalRows) % totalRows
	return u[row][col]
}

func (u Universe) CountAlives() int {
	count := 0
	for row := range u {
		for col := range u[row] {
			if u.Alive(col, row) {
				count++
			}
		}
	}
	return count
}

func (u Universe) Neighbors(col, row int) int {
	count := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i == row && j == col {
				continue
			}
			if u.Alive(j, i) {
				count += 1
			}
		}
	}
	return count
}

func (u Universe) Next(col, row int) bool {
	n := u.Neighbors(col, row)
	return n == 3 || n == 2 && u.Alive(col, row)
}

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
	// for i := 0; i < 5; i++ {
	for i := 0; i < 1000; i++ { // SWITCH COMMENTS TO RUN SHORTER PROGRAM
		Step(a, b)
		a.Show()
		// time.Sleep(time.Second)
		time.Sleep(time.Second / 100) // SWITCH COMMENTS TO SLOW DOWN
		a, b = b, a
	}
}
