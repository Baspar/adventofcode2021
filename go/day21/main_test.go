package main

import (
	"testing"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/stretchr/testify/assert"
)

var input = utils.SanitizeInput(`Player 1 starting position: 4
Player 2 starting position: 8`)

var d = &DayImpl{}

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	d.Init(input)

	res, err := d.Part1()

	assert.Equal("739785", res)
	assert.Nil(err)
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	d.Init(input)

	res, err := d.Part2()

	assert.Equal("", res)
	assert.Nil(err)
}
