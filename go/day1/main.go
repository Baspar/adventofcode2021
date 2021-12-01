package main

import (
	"fmt"
	"strconv"
	"strings"

	utils "github.com/baspar/adventofcode2021/internal"
)

type DayImpl struct {
	numbers []int
}

func (d *DayImpl) Init(input string) error {
	var (
		val int
		err error
	)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		val, err = strconv.Atoi(line)
		if err != nil {
			return err
		}
		d.numbers = append(d.numbers, val)
	}
	return nil
}

func (d *DayImpl) sumWithDelta(delta int) (total int) {
	for i := 0; i < len(d.numbers)-delta; i++ {
		if d.numbers[i] < d.numbers[i+delta] {
			total += 1
		}
	}
	return
}

func (d *DayImpl) Part1() (string, error) {
	return fmt.Sprint(d.sumWithDelta(1)), nil
}
func (d *DayImpl) Part2() (string, error) {
	return fmt.Sprint(d.sumWithDelta(3)), nil
}

func main() {
	utils.Run(&DayImpl{})
}
