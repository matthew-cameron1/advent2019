package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main () {
	file, error := os.Open("inputs/day1.txt")

	if error != nil {
		panic(error)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalFuel := 0

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		totalFuel += calculateFuelRecursive(mass)
	}

	fmt.Println("Required fuel is:", totalFuel)
}

func calculateFuelRecursive(mass int) int {
	fuel := calculateFuel(mass)

	if fuel <= 0 {
		return 0
	} else {
		return fuel + calculateFuelRecursive(fuel)
	}
}

func calculateFuel(mass int) int {
	rounded := int(mass / 3)
	return rounded - 2
}