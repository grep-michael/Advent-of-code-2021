package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	part1()
	part2()
}

const length_of_strings = 12

var data [length_of_strings][]string = getData()

func getData() [length_of_strings][]string {
	file, _ := os.Open("./day3Input.txt")
	scanner := bufio.NewScanner(file)
	var bits [length_of_strings][]string
	for scanner.Scan() {
		s := scanner.Text()
		for i := 0; i < length_of_strings; i++ {
			bits[i] = append(bits[i], string(s[i]))
		}
	}
	return bits
}
func getBitsFromPosition(i int) (string, string) {
	// i is the position we want to sum up
	//returns (mostCommon, leastCommon)
	var count [2]int //index 0 = 0's and 1 = 1's
	for _, v := range data[i] {
		intForm := int(rune(v[0])) - 48 // convert "1" -> 1 and "0" -> 0
		count[intForm] += 1
	}
	//TODO: theres a better way to do it i bet
	max := func() string {
		if count[0] > count[1] {
			return "0"
		} else {
			return "1"
		}
	}()
	min := func() string {
		if count[0] < count[1] {
			return "0"
		} else {
			return "1"
		}
	}()
	return max, min
}
func stringBinaryToInt(s string) int {
	var out int
	for i, _ := range s {
		pow := (len(s) - 1) - i
		mul := int(rune(s[i])) - 48
		out += mul * int(math.Pow(2, float64(pow)))
	}
	return out
}

func part2() {

}

func part1() {
	var gamma string
	var epsilon string
	for i, _ := range data {
		max, min := getBitsFromPosition(i)
		gamma += max
		epsilon += min
	}
	gamNum := stringBinaryToInt(gamma)
	epNum := stringBinaryToInt(epsilon)
	fmt.Printf("Part1:=%d\n", gamNum*epNum)
}
