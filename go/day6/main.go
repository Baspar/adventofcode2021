package main

import (
	"fmt"
	"strconv"
	"strings"

	utils "github.com/baspar/adventofcode2021/internal"
)

type DayImpl struct {
	fishes [9]int64
}

func (d *DayImpl) Init(lines []string) error {
	var (
		t   int
		err error
	)

	for i := range d.fishes {
		d.fishes[i] = 0
	}

	for _, timer := range strings.Split(lines[0], ",") {
		t, err = strconv.Atoi(timer)
		if err != nil {
			return fmt.Errorf("Cannot parse '%s' to int", timer)
		}
		if t < 0 || t > 8 {
			return fmt.Errorf("'%d' is out of the limit", t)
		}

		d.fishes[t] += 1
	}
	return nil
}
func (d *DayImpl) runOneDay() {
	newFishes := d.fishes[0]
	for i := 1; i<len(d.fishes); i++ {
		d.fishes[i-1] = d.fishes[i]
	}
	d.fishes[6] += newFishes
	d.fishes[8] = newFishes
}
func (d *DayImpl) Part1() (string, error) {
	for i := 0; i < 80; i++ {
		d.runOneDay()
	}
	var total int64 = 0
	for _, nb := range d.fishes {
		total += nb
	}
	return fmt.Sprint(total), nil
}
func (d *DayImpl) Part2() (string, error) {
	for i := 0; i < 256; i++ {
		d.runOneDay()
	}
	var total int64 = 0
	for _, nb := range d.fishes {
		total += nb
	}
	return fmt.Sprint(total), nil
}

func main() {
	utils.Run(&DayImpl{})
}
