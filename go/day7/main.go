package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"math"

	utils "github.com/baspar/adventofcode2021/internal"
	m "github.com/baspar/adventofcode2021/internal/math"
)

type DayImpl struct {
	positions []int
}

func (d *DayImpl) Init(lines []string) error {
	d.positions = nil
	values := strings.Split(lines[0], ",")
	for _, value := range values {
		n, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("Cannot parse '%s'", value)
		}
		d.positions = append(d.positions, n)
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	sort.Ints(d.positions)

	medianIndex := len(d.positions) / 2
	median := d.positions[medianIndex]

	distance := 0
	for _, position := range d.positions {
		distance += m.Abs(position - median)
	}
	return fmt.Sprintf("%d", distance), nil
}
func (d *DayImpl) Part2() (string, error) {
	min, max := m.Extremum(d.positions...)

	distance := func(pos int, level int) int {
		nbSteps := m.Abs(pos - level)
		return nbSteps * (1 + nbSteps) / 2 
	}

	minFuel := math.MaxInt
	for level := min; level <= max; level++ {
		fuel := 0
		for _, pos := range d.positions {
			fuel += distance(pos, level)
		}
		minFuel = m.Min(minFuel, fuel)
	}
	return fmt.Sprintf("%d", minFuel), nil
}

func main() {
	utils.Run(&DayImpl{})
}
