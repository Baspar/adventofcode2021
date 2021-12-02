package main

import (
	"errors"
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

func (d *DayImpl) Init(input string) error {
	var (
		err  error
		move Move
	)

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for index, line := range lines {
		tokens := strings.Split(line, " ")
		if len(tokens) < 2 {
			return errors.New(fmt.Sprintf("(#%d) Line '%s' is misshappen", index, line))
		}

		move.instruction = tokens[0]
		move.value, err = strconv.Atoi(tokens[1])
		if err != nil {
			return errors.New(fmt.Sprintf("(#%d) Cannot convert %s to int", index, tokens[1]))
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
func (d *DayImpl) Part2() (response string, err error) {
	return
}

func main() {
	utils.Run(&DayImpl{})
}
