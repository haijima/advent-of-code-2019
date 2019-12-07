package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(calculateTotalFuel("./day1/input.txt"))
}

func calculateTotalFuel(filepath string) int {
	var fp *os.File
	var err error

	fp, err = os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var sumFuel int
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		sumFuel += calculateFuel2(mass)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return sumFuel
}

func calculateFuel1(mass int) int {
	return mass/3 - 2
}

func calculateFuel2(mass int) int {
	fuel := mass/3 - 2
	if fuel <= 0 {
		return 0
	}
	return fuel + calculateFuel2(fuel)
}
