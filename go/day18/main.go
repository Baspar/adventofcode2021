package main

import (
	"fmt"
	"strconv"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/baspar/adventofcode2021/internal/math"
)

type Pair struct {
	left    *Pair
	right   *Pair
	literal int
}

func (p1 Pair) add(p2 Pair) (p Pair) {
	p.left = &p1
	p.right = &p2
	p.literal = 0

	for p.attemptToExplode(0, nil, nil) || p.attemptToSplit() {
	}

	return
}
func (p Pair) clone() Pair {
	if p.isLiteral() {
		return Pair{nil, nil, p.literal}
	}

	left := p.left.clone()
	right := p.right.clone()
	return Pair{
		&left,
		&right,
		0,
	}
}
func (p Pair) magnitude() int {
	if p.isLiteral() {
		return p.literal
	}

	return 3*p.left.magnitude() + 2*p.right.magnitude()
}
func (p Pair) isLiteral() bool {
	return p.left == nil
}
func (p *Pair) attemptToSplit() (hasChanged bool) {
	if p.isLiteral() && p.literal >= 10 {
		p.left = &Pair{nil, nil, p.literal / 2}
		p.right = &Pair{nil, nil, p.literal/2 + p.literal%2}
		return true
	}

	if !p.isLiteral() {
		return p.left.attemptToSplit() || p.right.attemptToSplit()
	}

	return false
}
func (p *Pair) attemptToExplode(depth int, closestLeftFork *Pair, closestRightFork *Pair) (hasChanged bool) {
	if p.isLiteral() {
		return false
	}

	if p.left.isLiteral() && p.right.isLiteral() && depth >= 4 {
		if closestLeftFork != nil {
			x := closestLeftFork
			for ; !x.isLiteral(); x = x.right {
			}
			x.literal += p.left.literal
		}
		if closestRightFork != nil {
			x := closestRightFork
			for ; !x.isLiteral(); x = x.left {
			}
			x.literal += p.right.literal
		}
		p.left = nil
		p.right = nil
		p.literal = 0
		return true
	}

	return p.left.attemptToExplode(depth+1, closestLeftFork, p.right) || p.right.attemptToExplode(depth+1, p.left, closestRightFork)
}
func parsePair(i *int, line string) (t Pair, err error) {
	var (
		left    Pair
		right   Pair
		literal int
	)
	if line[*i] == '[' {
		(*i)++ // Skip [
		if left, err = parsePair(i, line); err != nil {
			return
		}
		t.left = &left
		(*i)++ // Skip ,
		if right, err = parsePair(i, line); err != nil {
			return
		}
		t.right = &right
		(*i)++ // Skip ]
	} else {
		var (
			s string
		)
		for {
			c := line[*i]
			if c < '0' || c > '9' {
				break
			}
			s += string(c)
			(*i)++
			if *i >= len(line) {
				break
			}
		}
		if literal, err = strconv.Atoi(s); err != nil {
			err = fmt.Errorf("Cannot parse %s to int: %w", s, err)
			return
		} else {
			t.literal = literal
		}
	}
	return
}

type DayImpl struct {
	tokens []Pair
}

func (d *DayImpl) Init(lines []string) error {
	d.tokens = make([]Pair, 0)
	for _, line := range lines {
		i := 0
		if token, err := parsePair(&i, line); err != nil {
			return err
		} else {
			d.tokens = append(d.tokens, token)
		}
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	t := d.tokens[0]
	for i := 1; i < len(d.tokens); i++ {
		t = t.add(d.tokens[i])
	}
	return fmt.Sprint(t.magnitude()), nil
}
func (d *DayImpl) Part2() (string, error) {
	maxMagnitude := 0
	for i, d1 := range d.tokens {
		for j, d2 := range d.tokens {
			if i != j {
				d := d1.clone().add(d2.clone())
				maxMagnitude = math.Max(maxMagnitude, d.magnitude())
			}
		}
	}
	return fmt.Sprint(maxMagnitude), nil
}

func main() {
	utils.Run(&DayImpl{})
}
