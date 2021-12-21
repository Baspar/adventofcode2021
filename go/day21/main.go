package main

import (
	"fmt"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/baspar/adventofcode2021/internal/math"
)

type DayImpl struct {
	pos [2]int
}

func incMod(n int, inc int, mod int) int {
	return (n+inc-1)%mod + 1
}

func (d *DayImpl) Init(lines []string) error {
	fmt.Sscanf(lines[0], "Player 1 starting position: %d", &d.pos[0])
	fmt.Sscanf(lines[1], "Player 2 starting position: %d", &d.pos[1])
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	die := 1
	dieThrown := 0
	currentPlayer := 0
	var scores [2]int
	for scores[0] < 1000 && scores[1] < 1000 {
		for i := 0; i < 3; i++ {
			d.pos[currentPlayer] = incMod(d.pos[currentPlayer], die, 10)
			die = incMod(die, 1, 100)
			dieThrown++
		}
		scores[currentPlayer] += d.pos[currentPlayer]
		currentPlayer = (currentPlayer + 1) % 2
	}
	return fmt.Sprint(dieThrown * math.Min(scores[0], scores[1])), nil
}
func (d *DayImpl) Part2() (string, error) {
	return "", nil
}

func main() {
	utils.Run(&DayImpl{})
}
