package main

import (
	"fmt"

	"github.com/baspar/adventofcode2021/internal"
	"github.com/baspar/adventofcode2021/internal/math"
)

type Coord struct {
	x1 int
	y1 int
	x2 int
	y2 int
}
type DayImpl []Coord

func markCell(x int, y int, hist *(map[int]map[int]int)) {
	if (*hist)[x] == nil {
		(*hist)[x] = make(map[int]int)
	}

	(*hist)[x][y] += 1
}
func direction(from int, to int) int {
	if from == to {
		return 0
	}

	d := (to - from)
	return d / math.Abs(d)
}

func (d *DayImpl) Init(lines []string) error {
	*d = nil
	for _, line := range lines {
		var c Coord
		if _, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &c.x1, &c.y1, &c.x2, &c.y2); err != nil {
			return fmt.Errorf("Cannot parse line '%s'", line)
		}
		*d = append(*d, c)
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	hist := make(map[int]map[int]int)

	for _, coord := range *d {
		isRow := coord.x1 == coord.x2
		isColumn := coord.y1 == coord.y2
		if !(isRow || isColumn) {
			continue
		}

		dx := direction(coord.x1, coord.x2)
		dy := direction(coord.y1, coord.y2)
		for x, y := coord.x1, coord.y1; ; x, y = x+dx, y+dy {
			markCell(x, y, &hist)
			if x == coord.x2 && y == coord.y2 {
				break
			}
		}
	}

	tot := 0
	for _, m := range hist {
		for _, count := range m {
			if count >= 2 {
				tot++
			}
		}
	}

	return fmt.Sprint(tot), nil
}
func (d *DayImpl) Part2() (string, error) {
	hist := make(map[int]map[int]int)

	for _, coord := range *d {
		isRow := coord.x1 == coord.x2
		isColumn := coord.y1 == coord.y2
		isDiagonal := math.Abs(coord.x1 - coord.x2) == math.Abs(coord.y1 - coord.y2)
		if !(isRow || isColumn || isDiagonal) {
			continue
		}

		dx := direction(coord.x1, coord.x2)
		dy := direction(coord.y1, coord.y2)
		for x, y := coord.x1, coord.y1; ; x, y = x+dx, y+dy {
			markCell(x, y, &hist)
			if x == coord.x2 && y == coord.y2 {
				break
			}
		}
	}

	tot := 0
	for _, m := range hist {
		for _, count := range m {
			if count >= 2 {
				tot++
			}
		}
	}

	return fmt.Sprint(tot), nil
}

func main() {
	utils.Run(&DayImpl{})
}
