package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Day interface {
	Init(input string) error
	Part1() (string, error)
	Part2() (string, error)
}

func Run(day Day) {
	var (
		err       error
		stdinInfo os.FileInfo
		content   []byte
		part1     string
		part2     string
		start     time.Time
	)

	stdin := os.Stdin
	if stdinInfo, err = stdin.Stat(); err != nil {
		fmt.Printf("Cannot analyze stdin %s\n", err)
		return
	}

	if (stdinInfo.Mode() & os.ModeCharDevice) == 0 {
		if content, err = ioutil.ReadAll(stdin); err != nil {
			fmt.Printf("Cannot read stdin %s\n", err)
		} else {
			fmt.Print("Using stdin\n\n")
		}
	} else if content, err = ioutil.ReadFile("./input.txt"); err != nil {
		fmt.Printf("Input reading failed: %s\n", err)
		return
	}

	if err = day.Init(string(content)); err != nil {
		fmt.Printf("Init failed: %s\n", err)
		return
	}

	fmt.Print("Part1:\n======\n")
	start = time.Now()
	if part1, err = day.Part1(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("(%s)\n", time.Since(start))
	fmt.Println(part1)

	fmt.Print("\nPart1:\n======\n")
	start = time.Now()
	if part2, err = day.Part2(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("(%s)\n", time.Since(start))
	fmt.Println(part2)
}
