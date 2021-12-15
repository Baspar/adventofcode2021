package main

import (
	"testing"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/stretchr/testify/assert"
)

var input = utils.SanitizeInput(`1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`)

var d = &DayImpl{}

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	d.Init(input)

	res, err := d.Part1()

	assert.Equal("40", res)
	assert.Nil(err)
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	d.Init(input)

	res, err := d.Part2()

	assert.Equal("315", res)
	assert.Nil(err)
}
