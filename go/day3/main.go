package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	utils "github.com/baspar/adventofcode2021/internal"
)

type DayImpl struct {
	input [][]int
}

func (d *DayImpl) Init(input string) error {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	for _, line := range lines {
		var inputLine []int
		for _, ch := range line {
			if val, err := strconv.Atoi(string(ch)); err != nil {
				return errors.New(fmt.Sprintf("Unknown value: %c", ch))
			} else {
				inputLine = append(inputLine, val)
			}
		}
		d.input = append(d.input, inputLine)
	}

	if len(d.input) == 0 {
		return errors.New("Didn't find any line in input")
	}

	return nil
}
func (d *DayImpl) Part1() (string, error) {
	nbBits := len(d.input[0])

	gammaRate := 0
	epsilonRate := 0

	for i := 0; i < nbBits; i++ {
		hist := make(map[int]int)

		for _, line := range d.input {
			hist[line[i]] += 1
		}

		gammaRate *= 2
		epsilonRate *= 2
		if hist[1] > hist[0] {
			gammaRate += 1
		} else {
			epsilonRate += 1
		}
	}

	return fmt.Sprintf("%d", gammaRate*epsilonRate), nil
}
func (d *DayImpl) Part2() (response string, err error) {
	return
}

func main() {
	utils.Run(&DayImpl{})
}
