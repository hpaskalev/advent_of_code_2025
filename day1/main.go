package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Diretion int

const (
	Left  Diretion = -1
	Right Diretion = 1
)

const LEFT = "L"
const RIGHT = "R"

type Rotation struct {
	direction Diretion
	positions int
}

const BASE = 100

type Dial struct {
	position int
}

var rotationRegex = regexp.MustCompile(`^(L|R)(\d+)$`)

func (d *Dial) Rotate(rotation Rotation) bool {
	var delta int = rotation.positions
	if rotation.direction == Left {
		delta = delta * -1
	}

	d.position = ((d.position+delta)%BASE + BASE) % BASE
	return d.position == 0
}

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("Usage: %s <input file>\n", args[0])
		os.Exit(1)
	}

	filePath := args[1]

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fScan := bufio.NewScanner(f)

	dial := Dial{50}

	var lineNumber int = 1
	var clicks int = 0

	var matchResult []string
	var line string

	for fScan.Scan() {
		line = fScan.Text()

		matchResult = rotationRegex.FindStringSubmatch(line)
		if matchResult == nil {
			panic(fmt.Sprintf("invalid rotation format at line %d", lineNumber))
		}

		var direction Diretion
		switch matchResult[1] {
		case LEFT:
			direction = Left
		case RIGHT:
			direction = Right
		default:
			panic(fmt.Sprintf("invalid direction %s at line %d", matchResult[1], lineNumber))
		}

		positions, err := strconv.Atoi(matchResult[2])
		if err != nil {
			panic(fmt.Sprintf("invalid rotatation step %s at line %d", matchResult[2], lineNumber))
		}

		if dial.Rotate(Rotation{direction, positions}) {
			clicks++
		}
		lineNumber++
	}

	if err := fScan.Err(); err != nil {
		panic(err)
	}

	if clicks > 0 {
		fmt.Printf("Password is '%d'!\n", clicks)
	} else {
		fmt.Printf("Couldn't guess password!\n")
	}
}
