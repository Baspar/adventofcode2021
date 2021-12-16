package main

import (
	"fmt"

	utils "github.com/baspar/adventofcode2021/internal"
)

var translation = map[rune][]int{
	'0': {0, 0, 0, 0},
	'1': {0, 0, 0, 1},
	'2': {0, 0, 1, 0},
	'3': {0, 0, 1, 1},
	'4': {0, 1, 0, 0},
	'5': {0, 1, 0, 1},
	'6': {0, 1, 1, 0},
	'7': {0, 1, 1, 1},
	'8': {1, 0, 0, 0},
	'9': {1, 0, 0, 1},
	'A': {1, 0, 1, 0},
	'B': {1, 0, 1, 1},
	'C': {1, 1, 0, 0},
	'D': {1, 1, 0, 1},
	'E': {1, 1, 1, 0},
	'F': {1, 1, 1, 1},
}

type Packet struct {
	version    int
	typeID     int
	literal    int
	subpackets []Packet
}
func (p Packet) getSumVersions() (sum int) {
	sum += p.version
	for _, subpacket := range p.subpackets {
		sum += subpacket.getSumVersions()
	}
	return
}

type Bits []int
func (b Bits) readInt(ptr *int, size int) (value int) {
	for i := 0; i < size; i++ {
		value *= 2
		value += b[*ptr]
		*ptr++
	}
	return
}
func (b Bits) readSubpackets(ptr *int) (subpackets []Packet) {
	lengthTypeID := b[*ptr]
	*ptr++
	if lengthTypeID == 0 {
		length := b.readInt(ptr, 15)
		target := *ptr + length
		for *ptr != target {
			subpackets = append(subpackets, b.parsePacket(ptr))
		}
	} else {
		nbPackets := b.readInt(ptr, 11)
		for i := 0; i < nbPackets; i++ {
			subpackets = append(subpackets, b.parsePacket(ptr))
		}
	}
	return
}
func (b Bits) readLiteral(ptr *int) (literal int) {
	for {
		isLastGroup, block := b.readLiteralBlock(ptr)
		literal *= 16
		literal += block
		if isLastGroup {
			break
		}
	}
	return
}
func (b Bits) readLiteralBlock(ptr *int) (isLastGroup bool, value int) {
	isLastGroup = b[*ptr] == 0
	*ptr++
	value = b.readInt(ptr, 4)
	return
}
func (b Bits) parsePacket(ptr *int) (packet Packet) {
	packet.version = b.readInt(ptr, 3)
	packet.typeID = b.readInt(ptr, 3)

	if packet.typeID == 4 {
		packet.literal = b.readLiteral(ptr)
	} else {
		packet.subpackets = b.readSubpackets(ptr)
	}

	return
}

type DayImpl struct {
	bits Bits
}
func (d *DayImpl) Init(lines []string) error {
	for _, r := range lines[0] {
		d.bits = append(d.bits, translation[r]...)
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	ptr := 0
	ans := d.bits.parsePacket(&ptr).getSumVersions()
	return fmt.Sprint(ans), nil
}
func (d *DayImpl) Part2() (string, error) {
	return "", nil
}

func main() {
	utils.Run(&DayImpl{})
}
