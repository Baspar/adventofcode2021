package main

import (
	"fmt"

	utils "github.com/baspar/adventofcode2021/internal"
)

type Point struct {
	x int
	y int
}
type Grid map[Point]bool
func (g Grid) getNbLit(infiniteValue bool) (nbLit int, err error) {
	if infiniteValue {
		return nbLit, fmt.Errorf("There is an infinite number of lit cell")
	}

	for _, lit := range g {
		if lit {
			nbLit++
		}
	}
	return nbLit, nil
}
func (g Grid) getWindowBinValue(p Point, infiniteValue bool) int {
	val := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			val *= 2
			b, exists := g[Point{p.x + dx, p.y + dy}]
			if !exists {
				b = infiniteValue
			}

			if b {
				val++
			}
		}
	}
	return val
}

type DayImpl struct {
	rules                  [512]bool
	grid                   Grid
	minX, minY, maxX, maxY int
	infiniteValue          bool
}
func (d *DayImpl) step() {
	d.minX--
	d.minY--
	d.maxX++
	d.maxY++

	grid := make(Grid)
	for x := d.minX; x < d.maxX; x++ {
		for y := d.minY; y < d.maxY; y++ {
			val := d.grid.getWindowBinValue(Point{x, y}, d.infiniteValue)
			grid[Point{x, y}] = d.rules[val]
		}
	}

	d.grid = grid

	if d.infiniteValue {
		d.infiniteValue = d.rules[511]
	} else {
		d.infiniteValue = d.rules[0]
	}
}
func (d *DayImpl) Init(lines []string) error {
	for i, char := range lines[0] {
		d.rules[i] = char == '#'
	}

	d.minX, d.minY = 0, 0
	d.maxX, d.maxY = len(lines)-2, len(lines)-2

	d.infiniteValue = false

	d.grid = make(Grid)
	for x, line := range lines[2:] {
		for y, cell := range line {
			d.grid[Point{x, y}] = cell == '#'
		}
	}

	return nil
}
func (d *DayImpl) Part1() (string, error) {
	for i := 0; i < 2; i++ {
		d.step()
	}

	if nbLit, err := d.grid.getNbLit(d.infiniteValue); err != nil {
		return "", err
	} else {
		return fmt.Sprint(nbLit), nil
	}
}
func (d *DayImpl) Part2() (string, error) {
	for i := 0; i < 50; i++ {
		d.step()
	}

	if nbLit, err := d.grid.getNbLit(d.infiniteValue); err != nil {
		return "", err
	} else {
		return fmt.Sprint(nbLit), nil
	}
}

func main() {
	utils.Run(&DayImpl{})
}
