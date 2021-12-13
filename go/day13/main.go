package main

import (
	"fmt"
	"strings"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/baspar/adventofcode2021/internal/math"
)

type Grid map[int]map[int]bool

func (g *Grid) add(x, y int) {
	if g == nil {
		*g = make(Grid)
	}
	if _, exists := (*g)[x]; !exists {
		(*g)[x] = make(map[int]bool)
	}
	(*g)[x][y] = true
}

type Fold struct {
	isVertical bool
	coord      int
}
type DayImpl struct {
	g     Grid
	folds []Fold
}

func (d *DayImpl) fold(i int) {
	newG := make(Grid)
	fold := d.folds[i]

	for x, line := range d.g {
		for y := range line {
			if fold.isVertical && x > fold.coord {
				x = 2*fold.coord - x
			}

			if !fold.isVertical && y > fold.coord {
				y = 2*fold.coord - y
			}

			newG.add(x, y)
		}
	}
	d.g = newG
}
func (d *DayImpl) Init(lines []string) error {
	readingGrid := true
	d.g = make(Grid)
	for _, line := range lines {
		if line == "" {
			readingGrid = false
			continue
		}

		if readingGrid {
			var x, y int
			if _, err := fmt.Sscanf(line, "%d,%d", &y, &x); err != nil {
				return fmt.Errorf("Cannot read Cell for '%s': %w", line, err)
			}
			d.g.add(x, y)
		} else {
			tokens := strings.Split(line, " ")
			var (
				dir   rune
				coord int
			)
			token := tokens[len(tokens)-1]
			if _, err := fmt.Sscanf(token, "%c=%d", &dir, &coord); err != nil {
				return fmt.Errorf("Cannot read fold for '%s': %w", token, err)
			}
			d.folds = append(d.folds, Fold{isVertical: dir == 'y', coord: coord})
		}
	}

	return nil
}
func (d *DayImpl) Part1() (string, error) {
	d.fold(0)

	tot := 0
	for _, line := range d.g {
		tot += len(line)
	}

	return fmt.Sprintf("%d", tot), nil
}
func (d *DayImpl) Part2() (string, error) {
	for i := range d.folds {
		d.fold(i)
	}

	var xs, ys []int
	for x := range d.g {
		xs = append(xs, x)
		for y := range d.g[x] {
			ys = append(ys, y)
		}
	}

	minX, maxX := math.Extremum(xs...)
	minY, maxY := math.Extremum(ys...)

	s := ""
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if d.g[x][y] {
				s += "██"
			} else {
				s += "  "
			}
		}
		s += "\n"
	}

	return s, nil
}

func main() {
	utils.Run(&DayImpl{})
}
