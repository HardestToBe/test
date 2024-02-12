package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	universe := make(Universe, height)
	for i := range universe {
		universe[i] = make([]bool, width)
	}
	return universe
}
func (universe Universe) Show() {
	for i := range universe {
		for c := range universe[i] {
			if universe[i][c] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func (u Universe) Seed() {
	for column := range u {
		for row := range u[column] {
			num := rand.Intn(4) + 1
			if num == 1 {
				u[column][row] = true
			} else {
				u[column][row] = false
			}
		}
	}
}
func (u Universe) Alive(x, y int) bool {
	/*if x >= width {
		x %= width
	} else if x < 0 {
		x += width
	}
	if y >= height {
		y %= height
	} else if y < 0 {
		y += height
	}*/
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}
func (u Universe) Neighbors(x, y int) int {
	neighbours := 0
	for dy := y - 1; dy <= y+1; dy++ {
		for dx := x - 1; dx <= x+1; dx++ {
			if dy == y && dx == x {
				continue
			}
			if u.Alive(dx, dy) {
				neighbours++
			}
		}
	}
	return neighbours
}
func (u Universe) Next(x, y int) bool {
	neighbours := u.Neighbors(x, y)
	if u[y][x] {
		if neighbours < 2 {
			return false
		} else if neighbours == 2 || neighbours == 3 {
			return true
		} else if neighbours > 3 {
			return false
		}
	} else {
		if neighbours == 3 {
			return true
		}
	}
	return u[y][x]
}
func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			alive := a.Next(x, y)
			b[y][x] = alive
		}
	}

}

func main() {
	a := NewUniverse()
	b := NewUniverse()
	a.Seed()
	for i := 0; i < 2; i++ {
		Step(a, b)
		a, b = b, a
		fmt.Print("\x0c")
		time.Sleep(1 * time.Second)
		a.Show()
		time.Sleep(3 * time.Second)

	}

}
