package main

import (
	"container/heap"
	"fmt"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/baspar/adventofcode2021/internal/math"
)

type Heap []State

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	return h[i].score <= h[j].score
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

type Room struct {
	isHallway bool
	amphipods string
}
type Rooms [11]Room
type State struct {
	rooms     Rooms
	prevState *State
	score     int
}
type DayImpl struct {
	rooms Rooms
}

var AmphipodCost = map[rune]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}
var RoomAmphipod = map[int]rune{
	2: 'A',
	4: 'B',
	6: 'C',
	8: 'D',
}

func (s State) print() {
	for state := s; state.prevState != nil; state = *state.prevState {
		state._print()
	}
}
func (s State) _print() {
	fmt.Println()
	fmt.Println(s)
	fmt.Println("#############")

	fmt.Print("#")
	for _, room := range s.rooms {
		if room.isHallway && len(room.amphipods) > 0 {
			fmt.Print(room.amphipods)
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println("#")

	fmt.Print("###")
	for _, room := range s.rooms {
		if !room.isHallway {
			if len(room.amphipods) == 2 {
				fmt.Printf("%c#", room.amphipods[0])
			} else {
				fmt.Print(" #")
			}
		}
	}
	fmt.Println("##")

	fmt.Print("  #")
	for _, room := range s.rooms {
		if !room.isHallway {
			if len(room.amphipods) > 0 {
				fmt.Printf("%c#", room.amphipods[len(room.amphipods)-1])
			} else {
				fmt.Print(" #")
			}
		}
	}
	fmt.Println("#")

	// ###B#C#B#D###
	// 	#A#D#C#A#
	fmt.Println("  #########")
}
func (s State) costToMove(fromRoomID, toRoomID int) (distance int) {
	amphipod := rune(s.rooms[fromRoomID].amphipods[0])

	distance = math.Abs(fromRoomID - toRoomID)

	fromRoom, toRoom := s.rooms[fromRoomID], s.rooms[toRoomID]
	if !fromRoom.isHallway {
		distance += (3 - len(fromRoom.amphipods))
	}
	if !toRoom.isHallway {
		distance += (2 - len(toRoom.amphipods))
	}

	return distance * AmphipodCost[amphipod]
}
func (s State) canMove(fromRoomID, toRoomID int) bool {
	direction := 1
	if fromRoomID < toRoomID {
		direction = -1
	}

	for i := toRoomID; i != fromRoomID; i += direction {
		room := s.rooms[i]
		if room.isHallway && len(room.amphipods) > 0 {
			return false
		}
	}

	return true
}
func (s State) roomIsPartiallyComplete(roomID int) bool {
	for _, amphipod := range s.rooms[roomID].amphipods {
		if amphipod != RoomAmphipod[roomID] {
			return false
		}
	}

	return true
}
func (s State) roomIsComplete(roomID int) bool {
	room := s.rooms[roomID]
	return len(room.amphipods) == 2 && s.roomIsPartiallyComplete(roomID)
}
func (s State) roomsAreComplete() bool {
	for roomID := range RoomAmphipod {
		if !s.roomIsComplete(roomID) {
			return false
		}
	}
	return true
}
func (s State) canGoInRoom(amphipod rune, roomID int) bool {
	if amphipod != RoomAmphipod[roomID] {
		return false
	}

	room := s.rooms[roomID]

	for _, a := range room.amphipods {
		if a != amphipod {
			return false
		}
	}

	if len(room.amphipods) == 2 {
		return false
	}

	return true
}
func (s State) nextStates() (states []State) {
	for fromRoomID, fromRoom := range s.rooms {
		for toRoomID, toRoom := range s.rooms {
			// Nothing to move
			if len(fromRoom.amphipods) == 0 {
				continue
			}

			// Obstructed path
			if !s.canMove(fromRoomID, toRoomID) {
				continue
			}

			// Need to change type of room
			if fromRoom.isHallway == toRoom.isHallway {
				continue
			}

			amphipod := rune(fromRoom.amphipods[0])

			// Hallway => Room
			if fromRoom.isHallway && !s.canGoInRoom(amphipod, toRoomID) {
				continue
			}

			// Room => Hallway
			if !fromRoom.isHallway {
				if s.roomIsPartiallyComplete(fromRoomID) {
					continue
				}
			}

			newRooms := s.rooms
			newRooms[fromRoomID].amphipods = fromRoom.amphipods[1:]
			newRooms[toRoomID].amphipods = string(amphipod) + toRoom.amphipods
			states = append(states, State{
				newRooms,
				&s,
				s.score + s.costToMove(fromRoomID, toRoomID),
			})
		}
	}

	return
}

func (d *DayImpl) Init(lines []string) error {
	d.rooms = Rooms{
		{true, ""},
		{true, ""},
		{false, fmt.Sprintf("%c%c", lines[2][3] , lines[3][3] )},
		{true, ""},
		{false, fmt.Sprintf("%c%c", lines[2][5] , lines[3][5] )},
		{true, ""},
		{false, fmt.Sprintf("%c%c", lines[2][7] , lines[3][7] )},
		{true, ""},
		{false, fmt.Sprintf("%c%c", lines[2][9] , lines[3][9] )},
		{true, ""},
		{true, ""},
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	q := &Heap{}
	initialState := State{
		d.rooms,
		nil,
		0,
	}
	heap.Push(q, initialState)

	initialState.print()
	seen := map[Rooms]bool{}
	for q.Len() > 0 {
		state := heap.Pop(q).(State)

		if state.roomsAreComplete() {
			return fmt.Sprint(state.score), nil
		}

		if seen[state.rooms] {
			continue
		}
		seen[state.rooms] = true

		for _, newState := range state.nextStates() {
			heap.Push(q, newState)
		}
	}

	return "", nil
}
func (d *DayImpl) Part2() (string, error) {
	return "", nil
}

func main() {
	utils.Run(&DayImpl{})
}
