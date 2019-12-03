package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("day01/input.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	fuel := 0
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		check(err)
		fuel += fuelRequired(i)
	}

	fmt.Print(fuel)
}

func fuelRequired(module int) int {
	return int(math.Floor(float64(module) / 3) - 2)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
