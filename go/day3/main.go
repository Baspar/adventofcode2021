package main

import (
	"errors"
	"fmt"

	utils "github.com/baspar/adventofcode2021/internal"
)

type DayImpl struct {
	input [][]bool
}

func (d *DayImpl) Init(lines []string) error {
	d.input = nil
	for _, line := range lines {
		var inputLine []bool
		for _, ch := range line {
			inputLine = append(inputLine, ch == '1')
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
		hist := make(map[bool]int)

		for _, line := range d.input {
			hist[line[i]] += 1
		}

		gammaRate *= 2
		epsilonRate *= 2
		if hist[true] > hist[false] {
			gammaRate += 1
		} else {
			epsilonRate += 1
		}
	}

	return fmt.Sprintf("%d", gammaRate*epsilonRate), nil
}

func (d *DayImpl) Part2() (response string, err error) {
	process := func(getCO2 bool) []bool {
		validLineIndexes := make(map[int]struct{})
		for index := range d.input {
			validLineIndexes[index] = struct{}{}
		}

		nbBits := len(d.input[0])
		for currentBit := 0; currentBit < nbBits && len(validLineIndexes) > 1; currentBit++ {
			hist := make(map[bool]int)

			for lineIndex := range validLineIndexes {
				hist[d.input[lineIndex][currentBit]] += 1
			}

			valueToRemove := getCO2 == (hist[true] >= hist[false])

			for lineIndex := range validLineIndexes {
				if d.input[lineIndex][currentBit] == valueToRemove {
					delete(validLineIndexes, lineIndex)
				}
			}
		}

		var out []bool
		for lineIndex := range validLineIndexes {
			out = d.input[lineIndex]
		}
		return out
	}
	binary2Int := func(binary []bool) int {
		out := 0
		for _, b := range binary {
			out *= 2
			if b {
				out += 1
			}
		}
		return out
	}

	oxyGenRating := binary2Int(process(false))
	co2ScrubRating := binary2Int(process(true))

	return fmt.Sprintf("%d", oxyGenRating*co2ScrubRating), nil
}

func main() {
	utils.Run(&DayImpl{})
}
