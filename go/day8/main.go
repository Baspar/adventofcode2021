package main

import (
	"fmt"
	"reflect"
	"strings"

	utils "github.com/baspar/adventofcode2021/internal"
	set "github.com/baspar/adventofcode2021/internal/set"
)

type Entry struct {
	signals []string
	outputs []string
}
type DayImpl []Entry

//  0000
// 1    2
// 1    2
//  3333
// 4    5
// 4    5
//  6666
var segmentsForNumber = map[int]map[int]struct{}{
	0: {0: {}, 1: {}, 2: {}, 4: {}, 5: {}, 6: {}},
	1: {2: {}, 5: {}},
	2: {0: {}, 2: {}, 3: {}, 4: {}, 6: {}},
	3: {0: {}, 2: {}, 3: {}, 5: {}, 6: {}},
	4: {1: {}, 2: {}, 3: {}, 5: {}},
	5: {0: {}, 1: {}, 3: {}, 5: {}, 6: {}},
	6: {0: {}, 1: {}, 3: {}, 4: {}, 5: {}, 6: {}},
	7: {0: {}, 2: {}, 5: {}},
	8: {0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}},
	9: {0: {}, 1: {}, 2: {}, 3: {}, 5: {}, 6: {}},
}

func (d *DayImpl) Init(lines []string) error {
	*d = make([]Entry, 0)
	for _, line := range lines {
		var entry Entry
		tokens := strings.Split(line, " | ")
		entry.signals = strings.Split(tokens[0], " ")
		entry.outputs = strings.Split(tokens[1], " ")
		*d = append(*d, entry)
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	izi := 0
	for _, entry := range *d {
		for _, output := range entry.outputs {
			if len(output) == 2 || len(output) == 3 || len(output) == 4 || len(output) == 7 {
				izi++
			}
		}
	}
	return fmt.Sprintf("%d", izi), nil
}
func (d *DayImpl) Part2() (string, error) {
	type Permutation []rune

	isCorrectSignal := func(signal string, permutation []rune) bool {
		s := make(map[rune]bool)
		for _, letter := range signal {
			s[letter] = true
		}

		for _, segments := range segmentsForNumber {
			expectedLetters := make(map[rune]bool)
			for segment := range segments {
				expectedLetters[permutation[segment]] = true
			}
			if reflect.DeepEqual(expectedLetters, s) {
				return true
			}
		}
		return false
	}

	convertOutputToDigit := func(letters string, letter2Segment map[rune]int) (int, error) {
		s := make(map[int]struct{})
		for _, letter := range letters {
			s[letter2Segment[letter]] = struct{}{}
		}

		for digit, segments := range segmentsForNumber {
			if reflect.DeepEqual(segments, s) {
				return digit, nil
			}
		}

		return -1, fmt.Errorf("Cannot convert %s to a digit", letters)
	}

	determinePermutations := func(signals []string) []Permutation {
		var (
			possibleValuesForSegment [7][]rune
			valuesForNumber          [10][]rune
		)

		// Find "easy" values
		for _, signal := range signals {
			letters := []rune(signal)
			switch len(letters) {
			case 2:
				valuesForNumber[1] = letters
				possibleValuesForSegment[2] = valuesForNumber[1]
				possibleValuesForSegment[5] = valuesForNumber[1]
			case 3:
				valuesForNumber[7] = letters
			case 4:
				valuesForNumber[4] = letters
			}
		}

		// With 1 and 7, we can and find value for segment 0
		possibleValuesForSegment[0] = set.Difference(valuesForNumber[1], valuesForNumber[7])

		// With 1, and 4, we can narrow down the possible values
		// for segment 1 and 3 (Complimentary)
		possibleValuesForSegment[1] = set.Excluding(valuesForNumber[4], valuesForNumber[1])
		possibleValuesForSegment[3] = possibleValuesForSegment[1]

		// With 1, and 7, we can narrow down the possible values
		// for segment 5 and 2 (Complimentary)
		possibleValuesForSegment[2] = set.Excluding(valuesForNumber[7], possibleValuesForSegment[0])
		possibleValuesForSegment[5] = possibleValuesForSegment[2]

		// Rest of the segments (Complimentary)
		possibleValuesForSegment[4] = set.Excluding([]rune("abcdefg"), possibleValuesForSegment[0], possibleValuesForSegment[1], possibleValuesForSegment[2])
		possibleValuesForSegment[6] = possibleValuesForSegment[4]

		p := possibleValuesForSegment
		var permutations []Permutation
		permutations = append(permutations, Permutation{p[0][0], p[1][0], p[2][0], p[3][1], p[4][0], p[5][1], p[6][1]})
		permutations = append(permutations, Permutation{p[0][0], p[1][1], p[2][0], p[3][0], p[4][0], p[5][1], p[6][1]})
		permutations = append(permutations, Permutation{p[0][0], p[1][0], p[2][1], p[3][1], p[4][0], p[5][0], p[6][1]})
		permutations = append(permutations, Permutation{p[0][0], p[1][1], p[2][1], p[3][0], p[4][0], p[5][0], p[6][1]})
		permutations = append(permutations, Permutation{p[0][0], p[1][0], p[2][0], p[3][1], p[4][1], p[5][1], p[6][0]})
		permutations = append(permutations, Permutation{p[0][0], p[1][1], p[2][0], p[3][0], p[4][1], p[5][1], p[6][0]})
		permutations = append(permutations, Permutation{p[0][0], p[1][0], p[2][1], p[3][1], p[4][1], p[5][0], p[6][0]})
		permutations = append(permutations, Permutation{p[0][0], p[1][1], p[2][1], p[3][0], p[4][1], p[5][0], p[6][0]})
		return permutations
	}

	findCorrectPermutation := func(signals []string) (Permutation, error) {
		for _, permutation := range determinePermutations(signals) {
			allSignalsCorrect := true
			for _, signal := range signals {
				if !isCorrectSignal(signal, permutation) {
					allSignalsCorrect = false
					break
				}
			}
			if allSignalsCorrect {
				return permutation, nil
			}
		}

		return nil, fmt.Errorf("Cannot find permutation for signals %s", signals)
	}

	tot := 0
	for _, entry := range *d {
		var (
			correctPermutation Permutation
			err                error
			digit              int
		)

		if correctPermutation, err = findCorrectPermutation(entry.signals); err != nil {
			return "", err
		}

		letter2Segment := make(map[rune]int)
		for segment, letter := range correctPermutation {
			letter2Segment[letter] = segment
		}

		number := 0
		for _, output := range entry.outputs {
			number *= 10
			if digit, err = convertOutputToDigit(output, letter2Segment); err != nil {
				return "", err
			}
			number += digit
		}
		tot += number
	}

	return fmt.Sprintf("%d", tot), nil
}

func main() {
	utils.Run(&DayImpl{})
}
