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

func mark(x int, y int, hist *(map[int]map[int]int)) {
	if (*hist)[x] == nil {
		(*hist)[x] = make(map[int]int)
	}

	if _, exist := (*hist)[x][y]; !exist {
		(*hist)[x][y] = 0
	}
	(*hist)[x][y] += 1
}

func (d *DayImpl) Init(lines []string) error {
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
		dirX := math.Direction(coord.x1, coord.x2)
		dirY := math.Direction(coord.y1, coord.y2)

		if coord.x1 == coord.x2 {
			for y := coord.y1; ; y += dirY {
				mark(coord.x1, y, &hist)
				if y == coord.y2 {
					break
				}
			}
		} else if coord.y1 == coord.y2 {
			for x := coord.x1; ; x += dirX {
				mark(x, coord.y1, &hist)
				if x == coord.x2 {
					break
				}
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
		dirX := math.Direction(coord.x1, coord.x2)
		dx := math.Abs(coord.x1 - coord.x2)
		dirY := math.Direction(coord.y1, coord.y2)
		dy := math.Abs(coord.y1 - coord.y2)

		if coord.x1 == coord.x2 {
			for y := coord.y1; ; y += dirY {
				mark(coord.x1, y, &hist)
				if y == coord.y2 {
					break
				}
			}
		} else if coord.y1 == coord.y2 {
			for x := coord.x1; ; x += dirX {
				mark(x, coord.y1, &hist)
				if x == coord.x2 {
					break
				}
			}
		} else if dx == dy {
			for d := 0; d <= dx; d++ {
				mark(coord.x1+d*dirX, coord.y1+d*dirY, &hist)
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
