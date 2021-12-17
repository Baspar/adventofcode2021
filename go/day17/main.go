package main

import (
	"errors"
	"fmt"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/baspar/adventofcode2021/internal/math"
)

type Position struct {
	x  int
	y  int
	dx int
	dy int
}

func (p *Position) step() {
	p.x += p.dx
	p.y += p.dy
	p.dy--
	if p.dx > 0 {
		p.dx--
	} else if p.dx < 0 {
		p.dx++
	}
}
func (p *Position) run(t Target) (maxHeight int, err error) {
	maxHeight = p.y
	for {
		maxHeight = math.Max(maxHeight, p.y)
		if t.contains(*p) {
			return maxHeight, nil
		} else if t.overshot(*p) {
			return 0, errors.New("Overshot")
		} else if t.undershot(*p) {
			return 0, errors.New("Undershot")
		}

		p.step()
	}
}

type Target struct {
	x1 int
	x2 int
	y1 int
	y2 int
}

func (t Target) contains(pos Position) bool {
	return pos.x >= t.x1 &&
		pos.y >= t.y1 &&
		pos.x <= t.x2 &&
		pos.y <= t.y2
}
func (t Target) overshot(pos Position) bool {
	return pos.x > t.x2
}
func (t Target) undershot(pos Position) bool {
	return pos.y < t.y1
}

type DayImpl struct {
	target Target
}

func (d *DayImpl) Init(lines []string) error {
	_, err := fmt.Sscanf(lines[0], "target area: x=%d..%d, y=%d..%d", &d.target.x1, &d.target.x2, &d.target.y1, &d.target.y2)
	return err
}
func (d *DayImpl) Part1() (string, error) {
	maxHeight := 0
	for dx := 0; dx < 1000; dx++ {
		for dy := 0; dy < 1000; dy++ {
			p := Position{0, 0, dx, dy}
			if max, err := p.run(d.target); err == nil {
				maxHeight = math.Max(maxHeight, max)
			}
		}
	}
	return fmt.Sprint(maxHeight), nil
}
func (d *DayImpl) Part2() (string, error) {
	tot := 0
	for dx := -1000; dx < 1000; dx++ {
		for dy := -1000; dy < 1000; dy++ {
			p := Position{0, 0, dx, dy}
			if _, err := p.run(d.target); err == nil {
				tot++
			}
		}
	}
	return fmt.Sprint(tot), nil
}

func main() {
	utils.Run(&DayImpl{})
}
