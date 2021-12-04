package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {

	part1()
}

type Boardoperator interface {
}
type node struct {
	value   int
	checked bool
	left    *node
	right   *node
	up      *node
	down    *node
}

type Board struct {
	startNode node // top left node
}

func parseBoards() {
	file, _ := os.Open("./day4Input.txt")
	scanner := bufio.NewScanner(file)
	var boards []Board
	//get directions
	scanner.Scan()
	//order := strings.Split(scanner.Text(), ",")

	//loop over boards
	for scanner.Scan() {
		b := Board()
		for i := 5; i > 0; i-- {
			scanner.Scan()
			row := strings.Split(scanner.Text(), ",")

		}

		break
	}

}

func part1() {
	parseBoards()
}
