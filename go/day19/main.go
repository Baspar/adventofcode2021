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
func (s Scanner) isAlignedWithCloud(cloud map[Point][]int) (isAligned bool, shift Point) {
		for p1 := range cloud {
			for _, p2 := range s {
				shift = p2.minus(p1)

				alignedPoints := make(map[int]int)
				for _, p := range s.shiftBy(shift) {
					if scanners, exists := cloud[p]; exists {
						for _, scannerId := range scanners {
							alignedPoints[scannerId]++
							if alignedPoints[scannerId] >= 12 {
								return true, shift
							}
						}
					}
				}
			}
		}
		return false, shift
}

type DayImpl struct {
	scanners []Scanner
}
func (d DayImpl) alignScanners() (alignedScanners map[int]Scanner, pointCloud map[Point][]int, err error) {
	alignedScanners = make(map[int]Scanner)
	otherScannersOrientations := make(map[int][]Scanner)
	for i, scanner := range d.scanners {
		if i == 0 {
			alignedScanners[i] = scanner
		} else {
			otherScannersOrientations[i] = scanner.allPermutations()
		}
	}

	pointCloud = make(map[Point][]int)
	for _, point := range d.scanners[0] {
		pointCloud[point] = []int{0}
	}

	alignScannerWithCloud := func() (err error) {
		for scannerId, otherScannerOrientations := range otherScannersOrientations {
			for _, s := range otherScannerOrientations {
				if isAligned, shift := s.isAlignedWithCloud(pointCloud); isAligned {
					alignedScanners[scannerId] = s
					delete(otherScannersOrientations, scannerId)
					fmt.Printf("\nAligned scanner %d (%d/%d)   \r\033[1F", scannerId, len(alignedScanners), len(d.scanners))
					for _, p := range s {
						p = p.minus(shift)
						pointCloud[p] = append(pointCloud[p], scannerId)
					}
					return
				}
			}
		}
		return fmt.Errorf("Cannot align any scanner")
	}

	for len(otherScannersOrientations) > 0 {
		if err := alignScannerWithCloud(); err != nil {
			return nil, nil, err
		}
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
	_, pointCloud, err := d.alignScanners()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(len(pointCloud)), nil
}
func (d *DayImpl) Part2() (string, error) {
	return "", nil
}

func main() {
	utils.Run(&DayImpl{})
}
