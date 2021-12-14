package main

import (
	"container/list"
	"fmt"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/baspar/adventofcode2021/internal/math"
)

type Rules map[rune]map[rune]rune
type Level struct {
	l1, l2         rune
	remainingDepth int
}
type Mem map[Level]Hist
type Hist map[rune]int
func (h *Hist) Extremum() (int, int) {
	var l []int
	for _, count := range *h {
		l = append(l, count)
	}
	return math.Extremum(l...)
}
func (h *Hist) Merge(h2 Hist) {
		for letter, count := range h2 {
			(*h)[letter] += count
		}
}

type DayImpl struct {
	polymer list.List
	rules   Rules
}
func (d *DayImpl) Init(lines []string) error {
	d.polymer.Init()
	for _, r := range lines[0] {
		d.polymer.PushBack(r)
	}

	d.rules = make(Rules)
	for i := 2; i < len(lines); i++ {
		var l1, l2, next rune
		if _, err := fmt.Sscanf(lines[i], "%c%c -> %c", &l1, &l2, &next); err != nil {
			return fmt.Errorf("Cannot parse %s: %w", lines[i], err)
		}
		if _, exists := d.rules[l1]; !exists {
			d.rules[l1] = make(map[rune]rune)
		}
		d.rules[l1][l2] = next
	}

	return nil
}
func (d *DayImpl) mem_recur(l Level, mem *Mem) Hist {
	if _, exists := (*mem)[l]; !exists {
		(*mem)[l] = d.recur(l, mem)
	}
	return (*mem)[l]
}
func (d *DayImpl) recur(l Level, mem *Mem) Hist {
	l1, l2, remainingDepth := l.l1, l.l2, l.remainingDepth

	hist := make(Hist)

	if remainingDepth == 0 {
		hist[l1]++
		return hist
	}

	intermediate := d.rules[l1][l2]
	hist.Merge(d.mem_recur(Level{l1, intermediate, remainingDepth - 1}, mem))
	hist.Merge(d.mem_recur(Level{intermediate, l2, remainingDepth - 1}, mem))

	return hist
}
func (d *DayImpl) Part1() (string, error) {
	mem := make(Mem)
	hist := make(Hist)

	for l := d.polymer.Front(); l.Next() != nil; l = l.Next() {
		l1 := l.Value.(rune)
		l2 := l.Next().Value.(rune)
		hist.Merge(d.mem_recur(Level{l1, l2, 10}, &mem))
	}

	hist[d.polymer.Back().Value.(rune)]++

	a, b := hist.Extremum()
	
	return fmt.Sprintf("%d", b-a), nil
}
func (d *DayImpl) Part2() (string, error) {
	mem := make(Mem)
	hist := make(Hist)

	for l := d.polymer.Front(); l.Next() != nil; l = l.Next() {
		l1 := l.Value.(rune)
		l2 := l.Next().Value.(rune)
		hist.Merge(d.mem_recur(Level{l1, l2, 40}, &mem))
	}

	hist[d.polymer.Back().Value.(rune)]++

	a, b := hist.Extremum()
	
	return fmt.Sprintf("%d", b-a), nil
}

func main() {
	utils.Run(&DayImpl{})
}
