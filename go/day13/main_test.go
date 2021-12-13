package main

import (
	"strings"
	"testing"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/stretchr/testify/assert"
)

var input = utils.SanitizeInput(`6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`)

var d = &DayImpl{}

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	d.Init(input)

	res, err := d.Part1()

	assert.Equal("17", res)
	assert.Nil(err)
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	d.Init(input)

	res, err := d.Part2()

	assert.Equal(strings.TrimLeft(`
██████████
██      ██
██      ██
██      ██
██████████
`, "\n"), res)
	assert.Nil(err)
}
