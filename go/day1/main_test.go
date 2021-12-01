package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `199
200
208
210
200
207
240
269
260
263`

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	d := &DayImpl{}
	d.Init(input)

	res, err := d.Part1()

	assert.Equal("7", res)
	assert.Nil(err)
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	d := &DayImpl{}
	d.Init(input)

	res, err := d.Part2()

	assert.Equal("5", res)
	assert.Nil(err)
}
