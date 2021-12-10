package main

import (
	"container/list"
	"fmt"
	"sort"

	utils "github.com/baspar/adventofcode2021/internal"
)

type DayImpl struct {
	lines []string
}

func (d *DayImpl) processLine(line string) (isCompleteLine bool, char rune, stack list.List) {
	pair := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	for _, char = range line {
		if _, isOpeningToken := pair[char]; isOpeningToken {
			stack.PushBack(char)
		} else if pair[stack.Back().Value.(rune)] == char {
			stack.Remove(stack.Back())
		} else {
			return false, char, stack
		}
	}

	return true, char, stack
}

func (d *DayImpl) Init(lines []string) error {
	d.lines = lines
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	char2Score := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	score := 0
	for _, line := range d.lines {
		if isCompleteLine, illegalChar, _ := d.processLine(line); !isCompleteLine {
			score += char2Score[illegalChar]
		}
	}

	return fmt.Sprintf("%d", score), nil
}
func (d *DayImpl) Part2() (string, error) {
	char2Score := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	var scores []int
	for _, line := range d.lines {
		if isCompleteLine, _, remainingStack := d.processLine(line); isCompleteLine {
			score := 0
			for e := remainingStack.Back(); e != nil; e = e.Prev() {
				char := e.Value.(rune)
				score *= 5
				score += char2Score[char]
			}
			scores = append(scores, score)
		}
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})

	score := scores[len(scores)/2]

	return fmt.Sprintf("%d", score), nil
}

func main() {
	utils.Run(&DayImpl{})
}
