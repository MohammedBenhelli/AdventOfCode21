package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	x := 0
	y := 0
	aim := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			tab := strings.Split(line, " ")
			val, _ := strconv.Atoi(string(tab[1]))
			switch string(tab[0]) {
			case "forward":
				y += val
				x += val * aim
			case "down":
				aim += val
			case "up":
				aim -= val
			}
		}
	}
	fmt.Println(x*y, x, y)
}
