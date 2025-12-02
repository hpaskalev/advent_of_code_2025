package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func countDigits(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Floor(math.Log10(math.Abs(float64(n))))) + 1
}

func digit(n int, pos int) int {
	num := int(math.Abs(float64(n)))
	digits := countDigits(n)
	return (num / int(math.Pow10(digits-pos))) % 10
}

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("Usage: %s <input file>\n", args[0])
		os.Exit(1)
	}

	filePath := args[1]

	f, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var ranges []int = make([]int, 0, 10)
	var currentRange strings.Builder
	var inputScanner = io.ByteScanner(bytes.NewBuffer(f))
	var number int

	for {
		b, err := inputScanner.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		if isDigit(b) {
			currentRange.WriteByte(b)
			continue
		}

		if number, err = strconv.Atoi(currentRange.String()); err != nil {
			panic(err)
		}

		ranges = append(ranges, number)
		currentRange.Reset()
	}

	var start int = 0
	var end int = 0

	var digits int

	var invalidSequence bool = false
	var invalidSequencesSum int = 0

	var i int = 0
	for {
		start = ranges[i]
		end = ranges[i+1]

		for j := start; j <= end; j++ {
			digits = countDigits(j)
			if digits%2 != 0 {
				continue
			}

			mid := digits / 2
			s, e := 1, mid+1
			if digit(j, s) != digit(j, e) {
				continue
			}

			if e == digits {
				invalidSequencesSum += j
				continue
			}

			invalidSequence = true
			for {
				s++
				e++

				if e > digits || s >= mid+1 {
					break
				}

				if digit(j, s) != digit(j, e) {
					invalidSequence = false
					break
				}

			}

			if invalidSequence {
				invalidSequencesSum += j
			}
			invalidSequence = false
		}

		i = i + 2
		if i > len(ranges)-1 {
			break
		}
	}

	fmt.Printf("Sum of all invalid ids: %d\n", invalidSequencesSum)
}
