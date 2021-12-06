package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day2(commands []string) int {
	var horizontal, depth int
	for _, cmd := range commands {
		split := strings.Split(cmd, " ")
		val, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "forward":
			horizontal = horizontal + val
		case "down":
			depth = depth + val
		case "up":
			depth = depth - val
		}
	}
	return horizontal * depth
}

func day2Part2(commands []string) int {
	var horizontal, depth, aim int
	for _, cmd := range commands {
		split := strings.Split(cmd, " ")
		val, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Printf("error parsing value: %s\n", err)
		}
		switch split[0] {
		case "forward":
			horizontal += val
			depth += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}
	return horizontal * depth
}

func loadInputDay2(filename string) ([]string, error) {
	var input []string
	file, err := os.Open(filename)
	if err != nil {
		return input, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input, nil

}

func main() {
	input, _ := loadInputDay2("day2_input.txt")
	fmt.Println(day2(input))
	fmt.Println(day2Part2(input))
}
