package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fileToArr() ([]string, []string) {
	arr := make([]string, 0)
	arr1 := make([]string, 0)
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			arr = append(arr, line)
			arr1 = append(arr1, line)
		}
	}

	return arr, arr1
}

func filterArr(arr []string, index int, cond bool) []string {
	if len(arr) == 1 {
		return arr
	}
	if len(arr) == 2 {
		for i := range arr[0] {
			if cond {
				if arr[0][i] != '1' {
					return []string{arr[1]}
				} else if arr[1][i] != '1' {
					return []string{arr[0]}
				}
			} else {
				if arr[0][i] != '0' {
					return []string{arr[1]}
				} else if arr[1][i] != '0' {
					return []string{arr[0]}
				}
			}
		}
	}
	res := make([]string, 0)
	res1 := make([]string, 0)
	count := 0
	for i := range arr {
		if arr[i][index] == '1' {
			count++
		}
	}
	for i := range arr {
		if arr[i][index] == '1' && count > len(arr)/2 {
			res = append(res, arr[i])
		} else if arr[i][index] == '0' && count <= len(arr)/2 {
			res = append(res, arr[i])
		} else {
			res1 = append(res1, arr[i])
		}
	}
	if cond {
		fmt.Println(count, arr, res)
		return res
	}
	return res1
}

func main() {
	arr, arr1 := fileToArr()
	count := 0

	for len(arr) != 1 || len(arr1) != 1 {
		arr = filterArr(arr, count, true)
		arr1 = filterArr(arr1, count, false)
		count++
	}

	i, _ := strconv.ParseInt(arr[0], 2, 64)
	j, _ := strconv.ParseInt(arr1[0], 2, 64)
	fmt.Println(i*j, arr[0], arr1[0])
}
