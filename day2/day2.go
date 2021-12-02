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
}

func buildMapOfInput() map[string]int {
	m := make(map[string]int)
	file, _ := os.Open("./day2Input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		//fmt.Println(s[0])
		i, _ := strconv.ParseInt(s[1], 10, 32)
		m[s[0]] += int(i)
	}
	return m
}

func part1() {
	m := buildMapOfInput()
	var x int
	var y int
	x += m["forward"]
	y += m["down"]
	y -= m["up"]
	fmt.Printf("x=%d,y=%d , product=%d\n", x, y, x*y)
}
