package main

import (
	"testing"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/stretchr/testify/assert"
)

var input = utils.SanitizeInput(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`)

func TestPart1(t *testing.T) {
	var (
		err error
		res string
	)
	assert := assert.New(t)

	d := &DayImpl{}
	
	err = d.Init(input)
	assert.Nil(err)

	res, err = d.Part1()

	assert.Equal("198", res)
	assert.Nil(err)
}

func TestPart2(t *testing.T) {
	var (
		err error
		res string
	)
	assert := assert.New(t)

	d := &DayImpl{}
	
	err = d.Init(input)
	assert.Nil(err)

	res, err = d.Part2()

	assert.Equal("230", res)
	assert.Nil(err)
}
