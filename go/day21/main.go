package main

import (
	"container/heap"
	"fmt"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/baspar/adventofcode2021/internal/math"
)

type DayImpl struct {
	pos [2]int
}

type State struct {
	pos           [2]int
	scores        [2]int
	currentPlayer int
}

type Heap []State

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	return h[i].scores[0]+h[i].scores[1] < h[j].scores[0]+h[j].scores[1]
}
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(State)) }
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func genDiracDiceDistribution() (out map[int]int64) {
	out = map[int]int64{0: 0}
	for d1 := 1; d1 <= 3; d1++ {
		for d2 := 1; d2 <= 3; d2++ {
			for d3 := 1; d3 <= 3; d3++ {
				out[d1+d2+d3]++
			}
		}
	}
	return
}

var diracDiceDistribution = genDiracDiceDistribution()

func (s State) won(winningScore int) (isOver bool, winner int) {
	for i, score := range s.scores {
		if score >= winningScore {
			return true, i
		}
	}

	return false, winner
}
func (oldState State) generateNextState(die *int) (s State) {
	s = oldState
	for i := 0; i < 3; i++ {
		incMod(&s.pos[s.currentPlayer], *die, 10)
		incMod(die, 1, 100)
	}
	s.scores[s.currentPlayer] += s.pos[s.currentPlayer]
	s.currentPlayer = (s.currentPlayer + 1) % 2
	return
}
func (oldState State) generateNextDiracState() (states map[State]int64) {
	states = make(map[State]int64)
	for dice, occurences := range diracDiceDistribution {
		s := oldState
		incMod(&s.pos[s.currentPlayer], dice, 10)
		s.scores[s.currentPlayer] += s.pos[s.currentPlayer]
		s.currentPlayer = (s.currentPlayer + 1) % 2
		states[s] = occurences
	}
	return
}

func incMod(n *int, inc int, mod int) {
	*n = (*n+inc-1)%mod + 1
}

func (d *DayImpl) Init(lines []string) error {
	fmt.Sscanf(lines[0], "Player 1 starting position: %d", &d.pos[0])
	fmt.Sscanf(lines[1], "Player 2 starting position: %d", &d.pos[1])
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	die := 1
	dieThrown := 0

	state := State{d.pos, [2]int{0, 0}, 0}

	for {
		isOver, _ := state.won(1000)
		if isOver {
			break
		}
		state = state.generateNextState(&die)
		dieThrown += 3
	}

	return fmt.Sprint(dieThrown * math.Min(state.scores[0], state.scores[1])), nil
}
func (d *DayImpl) Part2() (string, error) {
	initialState := State{d.pos, [2]int{0, 0}, 0}

	winStates := make(map[State]int)

	universeOccurrences := make(map[State]int64)
	universeOccurrences[initialState] = 1

	q := &Heap{}
	heap.Push(q, initialState)

	for q.Len() > 0 {
		state := heap.Pop(q).(State)

		if isOver, winner := state.won(21); isOver {
			winStates[state] = winner
			continue
		}

		for nextState, occurrences := range state.generateNextDiracState() {
			if _, exists := universeOccurrences[nextState]; !exists {
				heap.Push(q, nextState)
			}
			universeOccurrences[nextState] += occurrences * universeOccurrences[state]
		}
	}

	totalWins := [2]int64{0, 0}
	for winState, winner := range winStates {
		totalWins[winner] += universeOccurrences[winState]
	}

	if totalWins[0] > totalWins[1] {
		return fmt.Sprint(totalWins[0]), nil
	} else {
		return fmt.Sprint(totalWins[1]), nil
	}
}

func main() {
	utils.Run(&DayImpl{})
}
