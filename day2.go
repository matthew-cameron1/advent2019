package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
	OPCODES:
	99: program is finished, stop
	1: adds together numbers read from two positions and stores result in a third
		- Three integers following tell you the positions of these numbers
	2: Exactly the same as 1 but multiplies then stores
 */

func toInt(str string) int {

	i, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	objective := 19690720
	for n := 0; n < 99; n++ {
		for v := 0; v < 99; v++ {

			arr := partTwo(n, v)
			if toInt(arr[0]) == objective {
				fmt.Print(100 * n + v)
				break
			}
		}
	}
}

func partTwo(noun, verb int) []string {
	file, error := os.Open("inputs/day2.txt")
	var arr []string
	if error != nil {
		panic(error)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value := scanner.Text()
		arr = strings.Split(value, ",")
		arr[1] = strconv.Itoa(noun)
		arr[2] = strconv.Itoa(verb)

		position := 0
		code, _ := strconv.Atoi(arr[position])

		for {

			if code == 99 {
				break
			}

			pos1, pos2, pos3 := toInt(arr[position+1]), toInt(arr[position+2]), toInt(arr[position+3])

			num1, num2 := toInt(arr[pos1]), toInt(arr[pos2])
			value := 0
			if code == 1 {
				value = num1 + num2
			} else {
				value = num1*num2
			}

			arr[pos3] = strconv.Itoa(value)

			if len(arr) < position + 4 {
				break
			}
			position = position + 4
			code = toInt(arr[position])
		}
	}
	return arr
}
