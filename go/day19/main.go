package main

import (
	"fmt"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/baspar/adventofcode2021/internal/math"
)

const COS int = 0
const SIN int = 1

type Matrix [3][3]int
var rotX = Matrix{{1, 0, 0}, {0, COS, -SIN}, {0, SIN, COS}}
var rotY = Matrix{{COS, 0, SIN}, {0, 1, 0}, {-SIN, 0, COS}}
var rotZ = Matrix{{COS, -SIN, 0}, {SIN, COS, 0}, {0, 0, 1}}

type Point [3]int
func (p1 Point) manhattanDistance(p2 Point) (dist int) {
	for i := range p1 {
		dist += math.Abs(p1[i] - p2[i])
	}
	return
}
func (p1 Point) minus(p2 Point) (p Point) {
	for i := range p {
		p[i] = p1[i] - p2[i]
	}
	return
}

type Scanner []Point
func (s Scanner) shiftBy(p Point) (scanner Scanner) {
	for _, point := range s {
		scanner = append(scanner, point.minus(p))
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
func (s Scanner) allRotations() (scanners []Scanner) {
	instructions := []struct {
		n   int
		mat Matrix
		add bool
	}{
		{4, rotX, true}, // Front face
		{1, rotZ, false},
		{4, rotY, true}, // Left face
		{1, rotZ, false},
		{4, rotX, true}, // Back face
		{1, rotZ, false},
		{4, rotY, true}, // Right face
		{1, rotX, false},
		{4, rotZ, true}, // Top face
		{2, rotX, false},
		{4, rotZ, true}, // Bottom face
	}

	for _, instr := range instructions {
		for i := 0; i < instr.n; i++ {
			s = s.matMul(instr.mat)
			if instr.add {
				scanners = append(scanners, s)
			}
		}
	}

	return
}
func (s Scanner) isAlignedWithPointCloud(pointCloud map[Point][]int) (isAligned bool, shift Point) {
	freqShift := make(map[Point]map[int]int)
	for alignPoint1, scanners := range pointCloud {
		for _, alignPoint2 := range s {
			shift = alignPoint2.minus(alignPoint1)
			if _, exists := freqShift[shift]; !exists {
				freqShift[shift] = make(map[int]int)
			}

			for _, scannerId := range scanners {
				freqShift[shift][scannerId]++
				if freqShift[shift][scannerId] >= 12 {
					return true, shift
				}
			}
		}
	}
	return false, shift
}

type DayImpl struct {
	scanners []Scanner
}
func (d DayImpl) alignNextScanner(unalignedScannersRotations map[int][]Scanner, pointCloud map[Point][]int) (err error, scannerId int, scanner Scanner, shift Point) {
	for scannerId, otherScannerOrientations := range unalignedScannersRotations {
		for _, s := range otherScannerOrientations {
			if isAligned, shift := s.isAlignedWithPointCloud(pointCloud); isAligned {
				return nil, scannerId, s, shift
			}
		}
	}
	return fmt.Errorf("Cannot align any scanner"), 0, nil, Point{}
}
func (d DayImpl) alignAllScanners() (scannerLocations []Point, pointCloud map[Point][]int, err error) {
	// Contains all unaligned scanners and their rotation (scanner[0] is considered valid)
	unalignedScannersRotations := make(map[int][]Scanner)
	for i := 1; i < len(d.scanners); i++ {
		unalignedScannersRotations[i] = d.scanners[i].allRotations()
	}

	// Contains the scanner coord relative to scanners[0]
	scannerLocations = make([]Point, len(d.scanners))
	scannerLocations[0] = Point{0, 0, 0}

	// Contains the aligned points and their corresponding scanner IDs
	pointCloud = make(map[Point][]int)
	for _, point := range d.scanners[0] {
		pointCloud[point] = []int{0}
	}

	for len(unalignedScannersRotations) > 0 {
		err, scannerId, scanner, shift := d.alignNextScanner(unalignedScannersRotations, pointCloud)
		if err != nil {
			return nil, nil, err
		}

		delete(unalignedScannersRotations, scannerId)
		scannerLocations[scannerId] = shift
		fmt.Printf("\nAligned scanner %d (%d/%d)   \r\033[1F", scannerId, len(d.scanners)-len(unalignedScannersRotations), len(d.scanners))
		for _, p := range scanner.shiftBy(shift) {
			pointCloud[p] = append(pointCloud[p], scannerId)
		}
	}

	return
}
func (d *DayImpl) Init(lines []string) error {
	d.scanners = make([]Scanner, 0)
	scanner := make(Scanner, 0)
	for i, line := range lines {
		var x, y, z int
		if _, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z); err == nil {
			scanner = append(scanner, Point{x, y, z})
		}

		if line == "" || i == len(lines)-1 {
			d.scanners = append(d.scanners, scanner)
			scanner = make(Scanner, 0)
		}
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	_, pointCloud, err := d.alignAllScanners()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(len(pointCloud)), nil
}
func (d *DayImpl) Part2() (string, error) {
	scannerLocations, _, err := d.alignAllScanners()
	if err != nil {
		return "", err
	}

	maxDistance := 0
	for i := 0; i < len(scannerLocations); i++ {
		for j := i + 1; j < len(scannerLocations); j++ {
			maxDistance = math.Max(maxDistance, scannerLocations[i].manhattanDistance(scannerLocations[j]))
		}
	}
	return fmt.Sprint(maxDistance), nil
}

func main() {
	utils.Run(&DayImpl{})
}
