package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

type Instruction struct {
	direction string
	ammount   int
}

var instructions []Instruction = buildInstructionList()

func buildInstructionList() []Instruction {
	var ins []Instruction
	file, _ := os.Open("./day2Input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		i, _ := strconv.ParseInt(s[1], 10, 32)
		ins = append(ins, Instruction{direction: s[0], ammount: int(i)})

	}
	return ins
}

func part1() {
	var horizontal int
	var depth int
	for _, v := range instructions {
		switch v.direction {
		case "forward":
			horizontal += v.ammount
		case "down":
			depth += v.ammount
		case "up":
			depth -= v.ammount
		default:
			fmt.Printf("Default %s %d\n", v.direction, v.ammount)
		}
	}
	fmt.Printf("part1: %d\n", horizontal*depth)

}

func part2() {
	var aim int
	var horizontal int
	var depth int
	for _, v := range instructions {
		switch v.direction {
		case "forward":
			horizontal += v.ammount
			depth += aim * v.ammount
		case "down":
			aim += v.ammount
		case "up":
			aim -= v.ammount
		}
	}
	fmt.Printf("Part2: %d\n", horizontal*depth)
}
