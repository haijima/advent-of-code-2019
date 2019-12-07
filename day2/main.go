package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	ints := readFile("./day2/input.txt")
	//fmt.Println(run1202Program(ints))
	fmt.Println(findNounAndVerb(ints, 19690720))
}

func run1202Program(ints []int) int {
	ints[1] = 12
	ints[2] = 2
	return runIntCodeProgram(ints)
}

func findNounAndVerb(ints []int, output int) int {
	for noun := 0; noun < 100; noun++ {
		tmp := make([]int, len(ints))
		for verb := 0; verb < 100; verb++ {
			copy(tmp, ints)
			tmp[1] = noun
			tmp[2] = verb
			if runIntCodeProgram(tmp) == output {
				return noun*100 + verb
			}
		}
	}
	panic("Couldn't find")
}

func readFile(filepath string) []int {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	line, err := csv.NewReader(file).Read()
	if err != nil {
		panic(err)
	}
	ints := make([]int, len(line))
	for i, elem := range line {
		num, err := strconv.Atoi(elem)
		if err != nil {
			panic(err)
		}
		ints[i] = num
	}
	return ints
}

func runIntCodeProgram(ints []int) int {
	l := len(ints)
	for i := 0; i < len(ints); i += 4 {
		op := ints[i]
		if op == 99 {
			break
		}

		if ints[i+1] >= l || ints[i+2] >= l || ints[i+3] >= l {
			return -1
		}

		if op == 1 {
			ints[ints[i+3]] = ints[ints[i+1]] + ints[ints[i+2]]
		} else if op == 2 {
			ints[ints[i+3]] = ints[ints[i+1]] * ints[ints[i+2]]
		}
	}
	return ints[0]
}
