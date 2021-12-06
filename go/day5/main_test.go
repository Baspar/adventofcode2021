package main

import (
	"testing"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/stretchr/testify/assert"
)

var input = utils.SanitizeInput(`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`)

var d = &DayImpl{}

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	d.Init(input)

	res, err := d.Part1()

	assert.Equal("5", res)
	assert.Nil(err)
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	d.Init(input)

	res, err := d.Part2()

	assert.Equal("12", res)
	assert.Nil(err)
}
