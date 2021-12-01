package main

import (
	utils "github.com/baspar/adventofcode2021/internal"
)

type DayImpl struct {
	input string
}

func (d *DayImpl) Init(input string) error {
	d.input = input
	return nil
}
func (d *DayImpl) Part1() (response string, err error) {
	response = d.input
	return
}
func (d *DayImpl) Part2() (response string, err error) {
	response = d.input
	return
}

func main() {
	utils.Run(&DayImpl{})
}
