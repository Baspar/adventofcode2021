package main

import (
	"fmt"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/baspar/adventofcode2021/internal/math"
)

type Cube struct {
	x1, x2 int
	y1, y2 int
	z1, z2 int
}
func (c Cube) volume() int {
	return (c.x2 - c.x1) * (c.y2 - c.y1) * (c.z2 - c.z1)
}
func (c Cube) isEmpty() bool {
	return c.x1 >= c.x2 || c.y1 >= c.y2 || c.z1 >= c.z2
}
func (c1 Cube) intersection(c2 Cube) Cube {
	return Cube{
		math.Max(c1.x1, c2.x1), math.Min(c1.x2, c2.x2),
		math.Max(c1.y1, c2.y1), math.Min(c1.y2, c2.y2),
		math.Max(c1.z1, c2.z1), math.Min(c1.z2, c2.z2),
	}
}
func (c1 Cube) overlaps(c2 Cube) bool {
	return !c1.intersection(c2).isEmpty()
}
func (c Cube) splitAround(cutter Cube) (cubes []Cube) {
	type Limits struct {
		x [4]int
		y [4]int
		z [4]int
	}
	limits := Limits{
		x: [4]int{c.x1, cutter.x1, cutter.x2, c.x2},
		y: [4]int{c.y1, cutter.y1, cutter.y2, c.y2},
		z: [4]int{c.z1, cutter.z1, cutter.z2, c.z2},
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				cube := Cube{
					limits.x[i], limits.x[i+1],
					limits.y[j], limits.y[j+1],
					limits.z[k], limits.z[k+1],
				}

				if cube.isEmpty() || cube == cutter {
					continue
				}

				cubes = append(cubes, cube)
			}
		}
	}
	return
}
func (c1 Cube) minus(c2 Cube) (cubes []Cube) {
	return c1.splitAround(c1.intersection(c2))
}

type Instr struct {
	Cube
	on bool
}
func (instr Instr) applyTo(existingCubes []Cube) (cubes []Cube) {
	if instr.on {
		cubes = append(cubes, instr.Cube)
	}

	for _, existingCube := range existingCubes {
		if existingCube.overlaps(instr.Cube) {
			cubes = append(cubes, existingCube.minus(instr.Cube)...)
		} else {
			cubes = append(cubes, existingCube)
		}
	}
	return
}

type DayImpl struct {
	instructions []Instr
}
func (d *DayImpl) Init(lines []string) error {
	d.instructions = make([]Instr, len(lines))
	for i, line := range lines {
		instr := Instr{}
		var w string
		if _, err := fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d",
			&w,
			&instr.x1, &instr.x2,
			&instr.y1, &instr.y2,
			&instr.z1, &instr.z2,
		); err != nil {
			return fmt.Errorf("Cannot parse Cube %s: %w", line, err)
		} else {
			instr.on = w == "on"
			instr.x2++
			instr.y2++
			instr.z2++
		}

		d.instructions[i] = instr
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	existingCubes := []Cube{}

	for _, instr := range d.instructions {
		if instr.x1 < -50 || instr.x2 > 51 {
			continue
		}
		if instr.y1 < -50 || instr.y2 > 51 {
			continue
		}
		if instr.z1 < -50 || instr.z2 > 51 {
			continue
		}

		existingCubes = instr.applyTo(existingCubes)
	}

	volume := 0
	for _, existingCube := range existingCubes {
		volume += existingCube.volume()
	}
	return fmt.Sprint(volume), nil
}
func (d *DayImpl) Part2() (string, error) {
	existingCubes := []Cube{}

	for _, instr := range d.instructions {
		existingCubes = instr.applyTo(existingCubes)
	}

	volume := 0
	for _, existingCube := range existingCubes {
		volume += existingCube.volume()
	}
	return fmt.Sprint(volume), nil
}

func main() {
	utils.Run(&DayImpl{})
}
