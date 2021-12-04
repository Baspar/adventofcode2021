package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	utils "github.com/baspar/adventofcode2021/internal"
)

type Cell struct {
	number int
	marked bool
}

type Grid [][]Cell

func (g *Grid) print() {
	for _, line := range *g {
		for _, cell := range line {
			if cell.marked {
				fmt.Printf("*% 2d*", cell.number)
			} else {
				fmt.Printf(" % 2d ", cell.number)
			}
		}
		fmt.Println()
	}
}
func (g *Grid) SumUnmarked() int {
	total := 0
	for _, line := range *g {
		for _, cell := range line {
			if !cell.marked {
				total += cell.number
			}
		}
	}
	return total
}
func (g *Grid) CheckGrid(i int, j int) bool {
	lineComplete := true
	for d := 0; lineComplete && d < 5; d++ {
		lineComplete = (*g)[i][d].marked
	}

	rowComplete := true
	for d := 0; rowComplete && d < 5; d++ {
		rowComplete = (*g)[d][j].marked
	}

	return lineComplete || rowComplete
}
func (g *Grid) MarkNumberAndCheckForWinning(number int) bool {
	gridSolved := false
	for i, line := range *g {
		for j, cell := range line {
			if cell.number == number {
				(*g)[i][j].marked = true
				gridSolved = gridSolved || g.CheckGrid(i, j)
			}
		}
	}
	return gridSolved
}

type DayImpl struct {
	numbers []int
	grids   []Grid
}

func (d *DayImpl) playUntilOneWinningGrid() (int, error) {
	for _, number := range d.numbers {
		for _, grid := range d.grids {
			if grid.MarkNumberAndCheckForWinning(number) {
				return grid.SumUnmarked() * number, nil
			}
		}
	}

	return -1, errors.New("No winning grid has been found")
}
func (d *DayImpl) playUntilLastWinningGrid() (int, error) {
	hasWon := make(map[int]bool)

	for _, number := range d.numbers {
		for gridIndex, grid := range d.grids {
			if !hasWon[gridIndex] && grid.MarkNumberAndCheckForWinning(number) {
				hasWon[gridIndex] = true
				if len(hasWon) == len(d.grids) {
					return grid.SumUnmarked() * number, nil
				}
			}
		}
	}

	return -1, errors.New("No winning grid has been found")
}

func (d *DayImpl) Init(lines []string) error {
	var (
		number int
		err    error
	)

	numberOfGrids := (len(lines) - 1) / 6

	// Parse bingo numbers
	for _, numberStr := range strings.Split(lines[0], ",") {
		if number, err = strconv.Atoi(numberStr); err != nil {
			return fmt.Errorf("Cannot convert '%s' to int", numberStr)
		}
		d.numbers = append(d.numbers, number)
	}

	// Parse grids
	for gridIndex := 0; gridIndex < numberOfGrids; gridIndex++ {
		from := 2 + gridIndex*6
		gridLines := lines[from : from+5]
		grid := make(Grid, 5)
		for lineIndex, gridLine := range gridLines {
			var n int
			line := make([]Cell, 5)
			r := strings.NewReader(gridLine)
			for cellIndex := range line {
				if _, err = fmt.Fscanf(r, "%d", &n); err != nil {
					return fmt.Errorf("Cannot read line")
				}
				line[cellIndex].number = n
			}
			grid[lineIndex] = line
		}
		d.grids = append(d.grids, grid)
	}

	return nil
}
func (d *DayImpl) Part1() (string, error) {
	finalScore, err := d.playUntilOneWinningGrid()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(finalScore), nil
}
func (d *DayImpl) Part2() (string, error) {
	finalScore, err := d.playUntilLastWinningGrid()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(finalScore), nil
}

func main() {
	utils.Run(&DayImpl{})
}
