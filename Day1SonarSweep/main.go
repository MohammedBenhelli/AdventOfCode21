package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseFile() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var arr []int
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			val, _ := strconv.Atoi(line)
			arr = append(arr, val)
		}
	}
	return arr
}

func toSumSlice(arr []int) []int {
	ret := make([]int, len(arr)-2)
	for i := 0; i < len(arr)-2; i++ {
		sum := arr[i] + arr[i+1] + arr[i+2]
		ret[i] = sum
	}
	return ret
}

func main() {
	arr := toSumSlice(parseFile())
	tmp := arr[0]
	count := 0

	for i := 1; i < len(arr); i++ {
		if tmp < arr[i] {
			count += 1
		}
		tmp = arr[i]
	}

	fmt.Println(count)
}
