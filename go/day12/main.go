package main

import (
	"container/list"
	"fmt"
	"strings"

	utils "github.com/baspar/adventofcode2021/internal"
)

func isSmall(cave string) bool {
	return cave[0] >= 'a' && cave[0] <= 'z'
}

type Walker struct {
	currentCave        string
	visited            map[string]bool
	canVisitSmallTwice bool
}

func NewWalker(cave string, canVisitSmallTwice bool) Walker {
	visited := make(map[string]bool)
	if isSmall(cave) {
		visited[cave] = true
	}
	return Walker{cave, visited, canVisitSmallTwice}
}

func (e *Walker) MoveTo(nextCave string) Walker {
	explNode := NewWalker(nextCave, e.canVisitSmallTwice)

	if isSmall(nextCave) && e.visited[nextCave] {
		explNode.canVisitSmallTwice = false
	}

	for cave := range e.visited {
		explNode.visited[cave] = true
	}

	return explNode
}

type DayImpl struct {
	caves map[string][]string
}

func (d *DayImpl) Init(lines []string) error {
	d.caves = make(map[string][]string)
	for _, line := range lines {
		caves := strings.Split(line, "-")
		if len(caves) < 2 {
			return fmt.Errorf("Cannot parse line %s", line)
		}

		d.caves[caves[0]] = append(d.caves[caves[0]], caves[1])
		d.caves[caves[1]] = append(d.caves[caves[1]], caves[0])
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	var q list.List
	q.PushBack(NewWalker("start", false))

	numberOfPath := 0
	for q.Len() > 0 {
		node := q.Remove(q.Front()).(Walker)

		for _, neighbourCave := range d.caves[node.currentCave] {
			if neighbourCave == "end" {
				numberOfPath++
				continue
			}

			_, visited := node.visited[neighbourCave]
			if !visited {
				q.PushBack(node.MoveTo(neighbourCave))
			}
		}
	}

	return fmt.Sprintf("%d", numberOfPath), nil
}
func (d *DayImpl) Part2() (string, error) {
	var q list.List
	q.PushBack(NewWalker("start", true))

	numberOfPath := 0
	for q.Len() > 0 {
		node := q.Remove(q.Back()).(Walker)

		for _, neighbourCave := range d.caves[node.currentCave] {
			if neighbourCave == "end" {
				numberOfPath++
				continue
			}

			if neighbourCave == "start" {
				continue
			}

			_, visited := node.visited[neighbourCave]
			if !visited || node.canVisitSmallTwice {
				q.PushBack(node.MoveTo(neighbourCave))
			}
		}
	}

	return fmt.Sprintf("%d", numberOfPath), nil
}

func main() {
	utils.Run(&DayImpl{})
}
