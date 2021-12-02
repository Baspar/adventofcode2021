package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	d := &DayImpl{}
	d.Init(input)

	res, err := d.Part1()

	assert.Equal("150", res)
	assert.Nil(err)
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	d := &DayImpl{}
	d.Init(input)

	res, err := d.Part2()

	assert.Equal("", res)
	assert.Nil(err)
}
