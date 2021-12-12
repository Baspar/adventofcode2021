package main

import (
	"testing"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/stretchr/testify/assert"
)

var input1 = utils.SanitizeInput(`start-A
start-b
A-c
A-b
b-d
A-end
b-end`)
var input2 = utils.SanitizeInput(`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`)
var input3 = utils.SanitizeInput(`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`)

type TestSet struct {
	input []string
	expected string
}

var d = &DayImpl{}

func TestPart1(t *testing.T) {
	inputs := []TestSet{
		{input1, "10"},
		{input2, "19"},
		{input3, "226"},
	}

	assert := assert.New(t)

	for _, testSet := range inputs {
		d.Init(testSet.input)

		res, err := d.Part1()

		assert.Equal(testSet.expected, res)
		assert.Nil(err)
	}
}

func TestPart2(t *testing.T) {
	inputs := []TestSet{
		{input1, "36"},
		{input2, "103"},
		{input3, "3509"},
	}

	assert := assert.New(t)

	for _, testSet := range inputs {
		d.Init(testSet.input)

		res, err := d.Part2()

		assert.Equal(testSet.expected, res)
		assert.Nil(err)
	}
}
