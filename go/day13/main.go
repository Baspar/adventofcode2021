package main

import (
	"fmt"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/baspar/adventofcode2021/internal/math"
)

type Fold struct {
	isVertical bool
	coord      int
}
type Cell struct {
	x int
	y int
}
type Grid map[Cell]bool
func (g *Grid) toString() (s string) {
	var xs, ys []int
	for cell := range *g {
		xs = append(xs, cell.x)
		ys = append(ys, cell.y)
	}

	minX, maxX := math.Extremum(xs...)
	minY, maxY := math.Extremum(ys...)

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if (*g)[Cell{x, y}] {
				s += "██"
			} else {
				s += "  "
			}
		}
		s += "\n"
	}
	return
}
func (g *Grid) fold(fold Fold) {
	newG := make(Grid)
	for cell := range *g {
		x, y := cell.x, cell.y
		if fold.isVertical && x > fold.coord {
			x = 2*fold.coord - x
		}

		if !fold.isVertical && y > fold.coord {
			y = 2*fold.coord - y
		}

		newG[Cell{x, y}] = true
	}
	*g = newG
}

type DayImpl struct {
	g     Grid
	folds []Fold
}

func (d *DayImpl) Init(lines []string) error {
	readingGrid := true
	d.g = make(Grid)
	for _, line := range lines {
		if line == "" {
			readingGrid = false
		} else if readingGrid {
			var x, y int
			if _, err := fmt.Sscanf(line, "%d,%d", &y, &x); err != nil {
				return fmt.Errorf("Cannot read Cell for '%s': %w", line, err)
			}
			d.g[Cell{x, y}] = true
		} else {
			var (
				dir   rune
				coord int
			)
			if _, err := fmt.Sscanf(line, "fold along %c=%d", &dir, &coord); err != nil {
				return fmt.Errorf("Cannot read fold for '%s': %w", line, err)
			}
			d.folds = append(d.folds, Fold{isVertical: dir == 'y', coord: coord})
		}
	}

	return nil
}
func (d *DayImpl) Part1() (string, error) {
	d.g.fold(d.folds[0])

	return fmt.Sprintf("%d", len(d.g)), nil
}
func (d *DayImpl) Part2() (string, error) {
	for _, fold := range d.folds {
		d.g.fold(fold)
	}

	return d.g.toString(), nil
}

func main() {
	utils.Run(&DayImpl{})
}
