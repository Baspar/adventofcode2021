package main

import (
	"fmt"

	utils "github.com/baspar/adventofcode2021/internal"
)

const COS int = 0
const SIN int = 1

type Matrix [3][3]int

var rotX = Matrix{{1, 0, 0}, {0, COS, -SIN}, {0, SIN, COS}}
var rotY = Matrix{{COS, 0, SIN}, {0, 1, 0}, {-SIN, 0, COS}}
var rotZ = Matrix{{COS, -SIN, 0}, {SIN, COS, 0}, {0, 0, 1}}

type Point [3]int

func (p1 Point) minus(p2 Point) (p Point) {
	for i := range p {
		p[i] = p1[i] - p2[i]
	}
	return
}

type Scanner []Point

func (scanner Scanner) shiftBy(p Point) (s Scanner) {
	for _, point := range scanner {
		s = append(s, point.minus(p))
	}
	return
}
func (s Scanner) matMul(mat Matrix) (ns Scanner) {
	ns = make(Scanner, len(s))
	for index, point := range s {
		for i := range mat {
			ns[index][0] += mat[i][0] * point[i]
			ns[index][1] += mat[i][1] * point[i]
			ns[index][2] += mat[i][2] * point[i]
		}
	}
	return
}
func (scanner Scanner) allPermutations() (scanners []Scanner) {
	s := scanner

	// Front
	for i := 0; i < 4; i++ {
		s = s.matMul(rotX)
		scanners = append(scanners, s)
	}

	// Left
	s = s.matMul(rotZ)
	for i := 0; i < 4; i++ {
		s = s.matMul(rotY)
		scanners = append(scanners, s)
	}

	// Back
	s = s.matMul(rotZ)
	for i := 0; i < 4; i++ {
		s = s.matMul(rotX)
		scanners = append(scanners, s)
	}

	// Right
	s = s.matMul(rotZ)
	for i := 0; i < 4; i++ {
		s = s.matMul(rotY)
		scanners = append(scanners, s)
	}

	// Up
	s = s.matMul(rotX)
	for i := 0; i < 4; i++ {
		s = s.matMul(rotZ)
		scanners = append(scanners, s)
	}

	// Down
	s = s.matMul(rotX).matMul(rotX)
	for i := 0; i < 4; i++ {
		s = s.matMul(rotZ)
		scanners = append(scanners, s)
	}
	return
}
func (s1 Scanner) nbAlignedPoints(s2 Scanner) (nbAlignedPoints int) {
	pointsInS1 := make(map[Point]bool)
	for _, p := range s1 {
		pointsInS1[p] = true
	}
	for _, p := range s2 {
		if pointsInS1[p] {
			nbAlignedPoints++
		}
	}
	return
}
func (s1 Scanner) isAlignedWith(s2 Scanner) (isAligned bool, anchorPoint1 Point, anchorPoint2 Point) {
	for _, anchorPoint1 = range s1 {
		for _, anchorPoint2 = range s2 {
			nbAlignedPoints := s1.nbAlignedPoints(s2.shiftBy(anchorPoint2.minus(anchorPoint1)))
			if nbAlignedPoints >= 12 {
				return true, anchorPoint1, anchorPoint2
			}
		}
	}
	return false, Point{}, Point{}
}

type DayImpl struct {
	scanners []Scanner
}

func (d DayImpl) alignScanners() (alignedScanners map[int]Scanner) {
	alignedScanners = make(map[int]Scanner)
	otherScannersOrientations := make(map[int][]Scanner)
	for i, scanner := range d.scanners {
		if i == 0 {
			alignedScanners[i] = scanner
		} else {
			otherScannersOrientations[i] = scanner.allPermutations()
		}
	}

	tryToAlign2Scanners := func() {
		for i, s1 := range alignedScanners {
			for scannerId, otherScannerOrientations := range otherScannersOrientations {
				for _, s2 := range otherScannerOrientations {
					if isAligned, p1, p2 := s1.isAlignedWith(s2); isAligned {
						alignedScanners[scannerId] = s2.shiftBy(p2.minus(p1))
						delete(otherScannersOrientations, scannerId)
						fmt.Printf("Aligning scanner %d with %d (%d/%d) <= [%d/%d]\n", scannerId, i, len(alignedScanners), len(d.scanners), len(otherScannersOrientations), len(d.scanners))
						return
					}
				}
			}
		}
		fmt.Println("Found nothing ಥ_ಥ")
	}

	for len(otherScannersOrientations) > 0 {
		tryToAlign2Scanners()
	}

	return
}
func (d *DayImpl) Init(lines []string) error {
	scanner := make(Scanner, 0)
	for _, line := range lines {
		if line == "" {
			d.scanners = append(d.scanners, scanner)
			scanner = make(Scanner, 0)
		}

		var x, y, z int
		if _, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z); err == nil {
			scanner = append(scanner, Point{x, y, z})
		}
	}
	d.scanners = append(d.scanners, scanner)
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	alignedScanners := d.alignScanners()

	points := make(map[Point]bool)
	for _, scanner := range alignedScanners {
		for _, point := range scanner {
			points[point] = true
		}
	}
	return fmt.Sprint(len(points)), nil
}
func (d *DayImpl) Part2() (string, error) {
	return "", nil
}

func main() {
	utils.Run(&DayImpl{})
}
