package main

import (
	"fmt"
	"strconv"
	"strings"

	utils "github.com/baspar/adventofcode2021/internal"
)

type Move struct {
	instruction string
	value       int
}
type DayImpl struct {
	input []Move
}

func (d *DayImpl) Init(lines []string) error {
	var (
		err  error
		move Move
	)

	for index, line := range lines {
		tokens := strings.Split(line, " ")
		if len(tokens) < 2 {
			return fmt.Errorf("(#%d) Line '%s' is misshappen", index, line)
		}

		move.instruction = tokens[0]
		move.value, err = strconv.Atoi(tokens[1])
		if err != nil {
			return fmt.Errorf("(#%d) Cannot convert %s to int", index, tokens[1])
		}

		d.input = append(d.input, move)
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	x := 0
	y := 0
	for _, move := range d.input {
		switch move.instruction {
		case "forward":
			x += move.value
		case "up":
			y -= move.value
		case "down":
			y += move.value
		}
	}
	return fmt.Sprintf("%d", x * y), nil
}
func (d *DayImpl) Part2() (string, error) {
	aim := 0
	x := 0
	y := 0
	for _, move := range d.input {
		switch move.instruction {
		case "forward":
			x += move.value
			y += move.value * aim
		case "up":
			aim -= move.value
		case "down":
			aim += move.value
		}
	}
	return fmt.Sprintf("%d", x * y), nil
}

func main() {
	utils.Run(&DayImpl{})
}
