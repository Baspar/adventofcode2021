package main

import (
	"testing"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/stretchr/testify/assert"
)

var input = utils.SanitizeInput(`target area: x=20..30, y=-10..-5`)

var d = &DayImpl{}

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	d.Init(input)

	res, err := d.Part1()

	assert.Equal("45", res)
	assert.Nil(err)
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	d.Init(input)

	res, err := d.Part2()

	assert.Equal("112", res)
	assert.Nil(err)
}
