package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func getDataArray() []int {
	file, _ := os.Open("./day1Input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var lines []int
	for scanner.Scan() {
		i, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		lines = append(lines, int(i))
	}
	return lines
}

func Sum(arr []int) int {
	var value int
	for _, v := range arr {
		value += v
	}
	return value
}

func part2() {

	var windowSize int = 3
	var count int
	data := getDataArray()
	for i, _ := range data[:len(data)-(windowSize)] {
		a := Sum(data[i : i+3])
		b := Sum(data[i+1 : i+4])
		if b > a {
			count++
		}
	}
	fmt.Printf("Part2 count: %d\n", count)

}

func part1() {
	data := getDataArray()
	var count int
	var prev int = data[0]
	for i, _ := range data[0:] {
		cur := data[i]
		if cur > prev {
			count++
		}
		prev = cur
	}
	fmt.Printf("Part1 count: %d\n", count)
}
