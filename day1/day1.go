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

func part2() {

	var windowSize int = 3
	var windowSums []int

	file, _ := os.Open("./day1Input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var lines []int
	for scanner.Scan() {
		i, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		lines = append(lines, int(i))
	}
	for i := 0; i < len(lines)-(windowSize-1); i++ {
		windowSums = append(windowSums, lines[i]+lines[i+1]+lines[i+2])
	}
	var count int
	var prev int = windowSums[0]
	for i := 1; i < len(windowSums); i++ {
		cur := windowSums[i]
		if cur > prev {
			count++
		}
		prev = cur
	}
	fmt.Printf("Part2 count: %d\n", count)

}

func part1() {
	file, _ := os.Open("./day1Input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var count int64
	scanner.Scan()
	prev, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	//fmt.Println(prev)
	for scanner.Scan() {
		cur, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		//fmt.Println(cur)
		if cur > prev {
			count++
		}
		prev = cur
	}
	fmt.Printf("Part1 count: %d\n", count)
}
