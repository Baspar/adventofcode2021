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
	rooms      Rooms
	roomHeight int
	score      int
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

func (s State) costToMove(fromRoomID, toRoomID int) (distance int) {
	amphipod := rune(s.rooms[fromRoomID].amphipods[0])

	distance = math.Abs(fromRoomID - toRoomID)

	fromRoom, toRoom := s.rooms[fromRoomID], s.rooms[toRoomID]
	if !fromRoom.isHallway {
		distance += (s.roomHeight - len(fromRoom.amphipods) + 1)
	}
	if !toRoom.isHallway {
		distance += (s.roomHeight - len(toRoom.amphipods))
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
	return len(room.amphipods) == s.roomHeight && s.roomIsPartiallyComplete(roomID)
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
	// Is not the correct room for the amphipod
	if amphipod != RoomAmphipod[roomID] {
		return false
	}

	// The room still contains wrong amphipod
	if !s.roomIsPartiallyComplete(roomID) {
		return false
	}

	// The room is full
	if len(s.rooms[roomID].amphipods) == s.roomHeight {
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
			if !fromRoom.isHallway && s.roomIsPartiallyComplete(fromRoomID) {
				continue
			}

			newRooms := s.rooms
			newRooms[fromRoomID].amphipods = fromRoom.amphipods[1:]
			newRooms[toRoomID].amphipods = string(amphipod) + toRoom.amphipods
			states = append(states, State{
				rooms:      newRooms,
				roomHeight: s.roomHeight,
				score:      s.score + s.costToMove(fromRoomID, toRoomID),
			})
		}
	}

	return
}

func Dijkstra(initialState State) (string, error) {
	q := &Heap{}
	heap.Push(q, initialState)

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

type DayImpl struct {
	rooms Rooms
}

func (d *DayImpl) Init(lines []string) error {
	d.rooms = Rooms{
		{true, ""},
		{true, ""},
		{false, fmt.Sprintf("%c%c", lines[2][3], lines[3][3])},
		{true, ""},
		{false, fmt.Sprintf("%c%c", lines[2][5], lines[3][5])},
		{true, ""},
		{false, fmt.Sprintf("%c%c", lines[2][7], lines[3][7])},
		{true, ""},
		{false, fmt.Sprintf("%c%c", lines[2][9], lines[3][9])},
		{true, ""},
		{true, ""},
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	initialState := State{
		rooms:      d.rooms,
		roomHeight: 2,
		score:      0,
	}
	return Dijkstra(initialState)
}
func (d *DayImpl) Part2() (string, error) {
	rooms := d.rooms
	rooms[2].amphipods = fmt.Sprintf("%cDD%c", rooms[2].amphipods[0], rooms[2].amphipods[1])
	rooms[4].amphipods = fmt.Sprintf("%cCB%c", rooms[4].amphipods[0], rooms[4].amphipods[1])
	rooms[6].amphipods = fmt.Sprintf("%cBA%c", rooms[6].amphipods[0], rooms[6].amphipods[1])
	rooms[8].amphipods = fmt.Sprintf("%cAC%c", rooms[8].amphipods[0], rooms[8].amphipods[1])
	initialState := State{
		rooms:      rooms,
		roomHeight: 4,
		score:      0,
	}
	return Dijkstra(initialState)
}

func main() {
	utils.Run(&DayImpl{})
}
