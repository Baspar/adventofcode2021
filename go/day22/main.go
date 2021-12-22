package main

import (
	"fmt"

	utils "github.com/baspar/adventofcode2021/internal"
	"github.com/baspar/adventofcode2021/internal/math"
)

type Position struct {
	x int
	y int
	z int
}
type Grid map[Position]bool
type Cube struct {
	x1, x2 int
	y1, y2 int
	z1, z2 int
	on     bool
}

func (c Cube) volume() int {
	return (c.x2 - c.x1) * (c.y2 - c.y1) * (c.z2 - c.z1)
}
func (c Cube) contains(p Position) bool {
	if p.x < c.x1 || p.x > c.x2 {
		return false
	}
	if p.y < c.y1 || p.y > c.y2 {
		return false
	}
	if p.z < c.z1 || p.z > c.z2 {
		return false
	}

	return true
}
func (c1 Cube) overlaps(c2 Cube) bool {
	corners := []Position{
		{c2.x1, c2.y1, c2.z1},
		{c2.x1, c2.y1, c2.z2},
		{c2.x1, c2.y2, c2.z1},
		{c2.x1, c2.y2, c2.z2},
		{c2.x2, c2.y1, c2.z1},
		{c2.x2, c2.y1, c2.z2},
		{c2.x2, c2.y2, c2.z1},
		{c2.x2, c2.y2, c2.z2},
	}

	for _, corner := range corners {
		if c1.contains(corner) {
			return true
		}
	}
	return false
}
func (c1 Cube) exclude(c2 Cube) (c1Sub []Cube, intersection Cube, c2Sub []Cube) { // TODO
	intersection = Cube{
		math.Max(c1.x1, c2.x1), math.Min(c1.x2, c2.x2),
		math.Max(c1.y1, c2.y1), math.Min(c1.y2, c2.y2),
		math.Max(c1.z1, c2.z1), math.Min(c1.z2, c2.z2),
		true,
	}

	type Limits struct {
		x [4]int
		y [4]int
		z [4]int
	}
	getSubCubesOf := func(c Cube) (subCubes []Cube) {
		limits := Limits{
			x: [4]int{c.x1, intersection.x1, intersection.x2, c.x2},
			y: [4]int{c.y1, intersection.y1, intersection.y2, c.y2},
			z: [4]int{c.z1, intersection.z1, intersection.z2, c.z2},
		}

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				for k := 0; k < 3; k++ {
					x1, x2 := limits.x[i], limits.x[i+1]
					y1, y2 := limits.y[j], limits.y[j+1]
					z1, z2 := limits.z[k], limits.z[k+1]
					cube := Cube{
						x1, x2,
						y1, y2,
						z1, z2,
						true,
					}

					if cube == intersection || x1 == x2 || y1 == y2 || z1 == z2 {
						continue
					}

					subCubes = append(subCubes, cube)
				}
			}
		}
		return
	}

	return getSubCubesOf(c1), intersection, getSubCubesOf(c2)
}

type DayImpl struct {
	cubes []Cube
}

func substractCube(existingCubes []Cube, offCube Cube) (cubes []Cube) {
	for _, existingCube := range existingCubes {
		if existingCube.overlaps(offCube) || offCube.overlaps(existingCube) {
			leftoversExistingCube, _, _ := existingCube.exclude(offCube)
			cubes = append(cubes, leftoversExistingCube...)
		} else {
			cubes = append(cubes, existingCube)
		}
	}
	return
}
func addCube(existingCubes []Cube, onCube Cube) (cubes []Cube) {
	for _, existingCube := range existingCubes {
		if existingCube.overlaps(onCube) || onCube.overlaps(existingCube) {
			cs, _, _ := existingCube.exclude(onCube)
			cubes = append(cubes, cs...)
		} else {
			cubes = append(cubes, existingCube)
		}
	}
	cubes = append(cubes, onCube)
	return
}

func (d *DayImpl) Init(lines []string) error {
	d.cubes = make([]Cube, len(lines))
	for i, line := range lines {
		cube := Cube{}
		var w string
		if _, err := fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d",
			&w,
			&cube.x1, &cube.x2,
			&cube.y1, &cube.y2,
			&cube.z1, &cube.z2,
		); err != nil {
			return fmt.Errorf("Cannot parse Cube %s: %w", line, err)
		} else {
			cube.on = w == "on"
			cube.x2++
			cube.y2++
			cube.z2++
		}

		d.cubes[i] = cube
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	existingCubes := []Cube{}

	for _, newCube := range d.cubes {
		if newCube.x1 < -50 || newCube.x2 > 51 {
			continue
		}
		if newCube.y1 < -50 || newCube.y2 > 51 {
			continue
		}
		if newCube.z1 < -50 || newCube.z2 > 51 {
			continue
		}
		if newCube.on {
			existingCubes = addCube(existingCubes, newCube)
		} else {
			existingCubes = substractCube(existingCubes, newCube)
		}
	}

	volume := 0
	for _, existingCube := range existingCubes {
		volume += existingCube.volume()
	}
	return fmt.Sprint(volume), nil
}
func (d *DayImpl) Part2() (string, error) {
	existingCubes := []Cube{}

	for _, newCube := range d.cubes {
		if newCube.on {
			existingCubes = addCube(existingCubes, newCube)
		} else {
			existingCubes = substractCube(existingCubes, newCube)
		}
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
