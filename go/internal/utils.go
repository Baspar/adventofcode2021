package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Day interface {
	Init(input []string) error
	Part1() (string, error)
	Part2() (string, error)
}

var Frames = []string{"⡆", "⠇", "⠋", "⠙", "⠸", "⢰", "⣠", "⣄"}
// ⠁⠂⠃⠄⠅⠆⠇⡀⡁⡂⡃⡄⡅⡆⡇⠈⠉⠊⠋⠌⠍⠎⠏⡈⡉⡊⡋⡌⡍⡎⡏⠐⠑⠒⠓⠔⠕⠖⠗⡐⡑⡒⡓⡔⡕⡖⡗⠘⠙⠚⠛⠜⠝⠞⠟⡘⡙⡚⡛⡜⡝⡞⡟⠠⠡⠢⠣⠤⠥⠦⠧⡠⡡⡢⡣⡤⡥⡦⡧⠨⠩⠪⠫⠬⠭⠮⠯⡨⡩⡪⡫⡬⡭⡮⡯⠰⠱⠲⠳⠴⠵⠶⠷⡰⡱⡲⡳⡴⡵⡶⡷⠸⠹⠺⠻⠼⠽⠾⠿⡸⡹⡺⡻⡼⡽⡾⡿⢀⢁⢂⢃⢄⢅⢆⢇⣀⣁⣂⣃⣄⣅⣆⣇⢈⢉⢊⢋⢌⢍⢎⢏⣈⣉⣊⣋⣌⣍⣎⣏⢐⢑⢒⢓⢔⢕⢖⢗⣐⣑⣒⣓⣔⣕⣖⣗⢘⢙⢚⢛⢜⢝⢞⢟⣘⣙⣚⣛⣜⣝⣞⣟⢠⢡⢢⢣⢤⢥⢦⢧⣠⣡⣢⣣⣤⣥⣦⣧⢨⢩⢪⢫⢬⢭⢮⢯⣨⣩⣪⣫⣬⣭⣮⣯⢰⢱⢲⢳⢴⢵⢶⢷⣰⣱⣲⣳⣴⣵⣶⣷⢸⢹⢺⢻⢼⢽⢾⢿⣸⣹⣺⣻⣼⣽⣾⣿

func SanitizeInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func Run(day Day) {
	var (
		err       error
		stdinInfo os.FileInfo
		content   []byte
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

	lines := SanitizeInput(string(content))
	if err = day.Init(lines); err != nil {
		fmt.Printf("Init failed: %s\n", err)
		return
	}

	processPart := func(part int, f func() (string, error)) {
		var start time.Time

		currentFrame := 0
		ansChan := make(chan string)
		errChan := make(chan error)
		ticker := time.NewTicker(100 * time.Millisecond)

		defer ticker.Stop()
		defer close(ansChan)
		defer close(errChan)

		go func() {
			start = time.Now()
			if ans, err := f(); err != nil {
				errChan <- err
			} else {
				ansChan <- ans
			}
		}()

		fmt.Printf("\033[1;33mPart%d: %s\033[0m", part, Frames[currentFrame])
	Loop:
		for {
			select {
			case <-ticker.C:
				currentFrame = (currentFrame + 1) % 8
				fmt.Printf("\r\033[1;33mPart%d: %s\033[0m", part, Frames[currentFrame])
			case ans := <-ansChan:
				fmt.Printf("\r\033[1;32mPart%d: ✓\033[0m (%s)\n\n", part, time.Since(start))
				fmt.Println(ans)
				break Loop
			case err := <-errChan:
				fmt.Printf("\r\033[1;31mPart%d: 𐄂\033[0m (%s)\n\n", part, time.Since(start))
				fmt.Println(err)
				break Loop
			}
		}
		fmt.Println("")
	}

	processPart(1, day.Part1)
	processPart(2, day.Part2)
}
