package main

import (
	"container/list"
	"fmt"

	utils "github.com/baspar/adventofcode2021/internal"
)

type cell struct {
	x     int
	y     int
	level int
}
type Cell *cell
type DayImpl struct {
	levels [][]Cell
}

func (d *DayImpl) getNeighboursOf(cell Cell) (neighbours []Cell) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			i, j := cell.x+dx, cell.y+dy

			if i >= 0 && j >= 0 && i < len(d.levels) && j < len(d.levels[i]) {
				neighbours = append(neighbours, d.levels[i][j])
			}
		}
	}
	return
}
func (d *DayImpl) runeOneStep() (flashes int) {
	var (
		octoToFlash list.List
		octotoBeFlashed = make(map[Cell]bool)
	)

	attemptToFlash := func(cell Cell) {
		if cell.level > 9 {
			octotoBeFlashed[cell] = true
			octoToFlash.PushBack(cell)
		}
	}

	flash := func(cell Cell) {
		flashes++
		cell.level = 0

		for _, neighbour := range d.getNeighboursOf(cell) {
			if !octotoBeFlashed[neighbour] {
				neighbour.level++
				attemptToFlash(neighbour)
			}
		}
	}

	// Initial discovery
	for _, line := range d.levels {
		for _, cell := range line {
			cell.level++
			attemptToFlash(cell)
		}
	}

	for octoToFlash.Len() > 0 {
		cell := octoToFlash.Remove(octoToFlash.Front()).(Cell)
		flash(cell)
	}

	return
}

func (d *DayImpl) Init(lines []string) error {
	d.levels = make([][]Cell, 0)
	for x, line := range lines {
		d.levels = append(d.levels, make([]Cell, 0))
		for y, level := range line {
			d.levels[x] = append(d.levels[x], &cell{x, y, int(level - '0')})
		}
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	flashes := 0
	for step := 0; step < 100; step++ {
		flashes += d.runeOneStep()
	}
	return fmt.Sprintf("%d", flashes), nil
}
func (d *DayImpl) Part2() (string, error) {
	expectedFlashes := 0
	for _, line := range d.levels {
		expectedFlashes += len(line)
	}

	currentStep := 1
	for d.runeOneStep() != expectedFlashes {
		currentStep++
	}

	return fmt.Sprintf("%d", currentStep), nil
}

func main() {
	utils.Run(&DayImpl{})
}
