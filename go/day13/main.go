package main

import (
	"fmt"
	"strings"

	utils "github.com/baspar/adventofcode2021/internal"
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
	readingFold := false
	d.g = make(Grid)
	for _, line := range lines {
		if line == "" {
			readingFold = true
			continue
		}

		if readingFold {
			tokens := strings.Split(line, " ")
			var (
				dir   rune
				coord int
			)
			token := tokens[len(tokens)-1]
			if _, err := fmt.Sscanf(token, "%c=%d", &dir, &coord); err != nil {
				return fmt.Errorf("Cannot read fold for '%s': %w", token, err)
			}
			d.folds = append(d.folds, Fold{isVertical: dir == 'x', coord: coord})
		} else {
			var x, y int
			if _, err := fmt.Sscanf(line, "%d,%d", &x, &y); err != nil {
				return fmt.Errorf("Cannot read Cell for '%s': %w", line, err)
			}
			d.g.add(x, y)
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
	return "", nil
}

func main() {
	utils.Run(&DayImpl{})
}
