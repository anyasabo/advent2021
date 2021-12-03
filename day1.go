package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://adventofcode.com/2021/day/1

func day1(measurements []int) int {
	var prev int
	increases := 0
	for i, val := range measurements {
		if i == 0 {
			prev = val
			continue
		}
		if prev < val {
			increases++
		}
		prev = val

	}
	return increases
}

// since each set to compare contains 2 of the same values
// we can get away with comparing the current value and the
// 3rd most recent value
func day1part2(vals []int) int {
	// yolo
	i := vals[0]
	j := vals[1]
	k := vals[2]
	var increases int
	for _, current := range vals[3:] {
		if current > i {
			increases++
		}
		i = j
		j = k
		k = current
	}
	return increases
}

func loadInput(filename string) ([]int, error) {
	var input []int
	file, err := os.Open(filename)
	if err != nil {
		return input, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return input, err
		}
		input = append(input, val)

	}
	return input, nil

}

func main() {
	input, err := loadInput("day1_input.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("part1: %d\n", day1(input))
		fmt.Printf("part2: %d\n", day1part2(input))
	}
}
