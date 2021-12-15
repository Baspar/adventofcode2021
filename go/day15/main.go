package main

import (
	"container/heap"
	"errors"
	"fmt"
	"strconv"

	utils "github.com/baspar/adventofcode2021/internal"
)

type Position struct {
	x int
	y int
}
type Expl struct {
	position     Position
	prevPosition Position
	score        int
}

type Heap []Expl
func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].score <= h[j].score }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(Expl)) }
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Grid [][]int
func (g Grid) getNeighbourOf(pos Position, scale int) (out []Position) {
	if pos.x-1 > 0 {
		out = append(out, Position{pos.x - 1, pos.y})
	}
	if pos.y-1 > 0 {
		out = append(out, Position{pos.x, pos.y - 1})
	}
	if pos.x+1 < scale*len(g) {
		out = append(out, Position{pos.x + 1, pos.y})
	}
	if pos.y+1 < scale*len(g[0]) {
		out = append(out, Position{pos.x, pos.y + 1})
	}
	return
}
func (g Grid) getScoreOf(pos Position, scale int) int {
	width, height := len(g), len(g[0])
	val := g[pos.x%width][pos.y%height]
	val = (val+pos.x/width+pos.y/height-1)%9 + 1

	return val
}
func (g Grid) dijkstra(initialPosition Position, scale int) (string, error) {
	previousPosition := make(map[Position]Position)

	q := &Heap{}
	heap.Init(q)
	heap.Push(q, Expl{initialPosition, initialPosition, 0})

	endingX, endingY := scale*len(g)-1, scale*len(g[0])-1

	for q.Len() > 0 {
		expl := heap.Pop(q).(Expl)

		if _, exists := previousPosition[expl.position]; exists {
			continue
		}

		previousPosition[expl.position] = expl.prevPosition

		if expl.position.x == endingX && expl.position.y == endingY {
			return fmt.Sprintf("%d", expl.score), nil
		}

		for _, pos := range g.getNeighbourOf(expl.position, scale) {
			heap.Push(q, Expl{pos, expl.position, expl.score + g.getScoreOf(pos, scale)})
		}
	}
	return "", errors.New("Cannot find path to last cell")
}
func (g Grid) str(prevPosition map[Position]Position, scale int) string {
	isInPath := make(map[Position]bool)
	pos := Position{scale*len(g)-1, scale*len(g[0])-1}
	for {
		isInPath[pos] = true
		if prevPosition[pos] == pos {
			break
		}
		pos=prevPosition[pos]
	}

	s := ""
	width, height := len(g)*scale, len(g[0])*scale
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := fmt.Sprint(g.getScoreOf(Position{x, y}, scale))
			if isInPath[Position{x, y}] {
				c = fmt.Sprintf("\033[1;32m%s\033[0m", c)
			} else {
				c = fmt.Sprintf("\033[1;33m%s\033[0m", c)
			}
			s += c
		}
		s += "\n"
	}

	return s
}

type DayImpl struct {
	grid Grid
}

func (d *DayImpl) Init(lines []string) error {
	d.grid = make(Grid, len(lines))
	for x, line := range lines {
		d.grid[x] = make([]int, len(line))
		for y, val := range line {
			if n, err := strconv.Atoi(string(val)); err != nil {
				return fmt.Errorf("Cannot parse '%c': %w", val, err)
			} else {
				d.grid[x][y] = n
			}
		}
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	initialPosition := Position{0, 0}
	return d.grid.dijkstra(initialPosition, 1)
}
func (d *DayImpl) Part2() (string, error) {
	initialPosition := Position{0, 0}
	return d.grid.dijkstra(initialPosition, 5)
}

func main() {
	utils.Run(&DayImpl{})
}
