package main

import (
	"container/list"
	"fmt"
	"sort"
	"strconv"

	utils "github.com/baspar/adventofcode2021/internal"
)

type Cell struct {
	x   int
	y   int
	val int
}
type DayImpl struct {
	grid [][]Cell
}

func (d *DayImpl) getNeighbours(cell Cell) (values []Cell) {
	i, j := cell.x, cell.y
	if i-1 >= 0 {
		values = append(values, d.grid[i-1][j])
	}
	if j-1 >= 0 {
		values = append(values, d.grid[i][j-1])
	}
	if i+1 < len(d.grid) {
		values = append(values, d.grid[i+1][j])
	}
	if j+1 < len(d.grid[i]) {
		values = append(values, d.grid[i][j+1])
	}
	return
}
func (d *DayImpl) findAllMins() (out []Cell) {
	for i, line := range d.grid {
		for j, cell := range line {
			isLocalMin := true
			for _, neighbour := range d.getNeighbours(cell) {
				if cell.val >= neighbour.val {
					isLocalMin = false
					break
				}
			}

			if isLocalMin {
				out = append(out, Cell{i, j, cell.val})
			}
		}
	}
	return
}

func (d *DayImpl) Init(lines []string) error {
	var (
		n   int
		err error
	)

	d.grid = make([][]Cell, len(lines))

	for i, line := range lines {
		row := make([]Cell, len(line))
		for j, char := range line {
			if n, err = strconv.Atoi(string(char)); err != nil {
				return fmt.Errorf("Cannot convert '%c' to int", char)
			}
			row[j] = Cell{i, j, n}
		}
		d.grid[i] = row
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	tot := 0
	for _, cell := range d.findAllMins() {
		tot += cell.val + 1
	}
	return fmt.Sprintf("%d", tot), nil
}
func (d *DayImpl) Part2() (string, error) {
	seen := make(map[Cell]bool)
	var bassinSizes []int

	for _, minCell := range d.findAllMins() {
		var q list.List
		q.PushBack(minCell)

		bassinSize := 0
		for q.Len() != 0 {
			cell := q.Remove(q.Front()).(Cell)

			if seen[cell] || cell.val == 9 {
				continue
			}

			seen[cell] = true
			bassinSize++

			for _, neighbour := range d.getNeighbours(cell) {
				q.PushBack(neighbour)
			}
		}
		bassinSizes = append(bassinSizes, bassinSize)
	}

	sort.Slice(bassinSizes, func(i, j int) bool {
		return bassinSizes[i] > bassinSizes[j]
	})

	if len(bassinSizes) < 3 {
		return "", fmt.Errorf("Found less than 3 bassins")
	}

	return fmt.Sprintf("%d", bassinSizes[0]*bassinSizes[1]*bassinSizes[2]), nil
}

func main() {
	utils.Run(&DayImpl{})
}
