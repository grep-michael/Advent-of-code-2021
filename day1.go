package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
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
	fmt.Println(count)

}
