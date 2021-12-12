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

type DayImpl struct {
	caves map[string][]string
}

type ExplNode struct {
	currentCave string
	visited     map[string]bool
}

func NewExplNode(cave string) ExplNode {
	visited := make(map[string]bool)
	if isSmall(cave) {
		visited[cave] = true
	}

	return ExplNode{cave, visited}
}

func (e *ExplNode) MoveTo(nextCave string) ExplNode {
	explNode := NewExplNode(nextCave)

	for cave := range e.visited {
		explNode.visited[cave] = true
	}

	return explNode
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
	q.PushBack(NewExplNode("start"))

	numberOfPath := 0
	for q.Len() > 0 {
		node := q.Remove(q.Front()).(ExplNode)

		if node.currentCave == "end" {
			numberOfPath++
		}

		for _, neighbourCave := range d.caves[node.currentCave] {
			_, visited := node.visited[neighbourCave]
			if !visited {
				q.PushBack(node.MoveTo(neighbourCave))
			}
		}
	}

	return fmt.Sprintf("%d", numberOfPath), nil
}
func (d *DayImpl) Part2() (string, error) {
	return "", nil
}

func main() {
	utils.Run(&DayImpl{})
}
