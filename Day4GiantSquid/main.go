package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Case struct {
	value int
	check bool
}

type Board struct {
	cases [][]Case
}

func (b *Board) checkVictory() bool {
	ret := true
	for i := range b.cases {
		ret = true
		for j := range b.cases[i] {
			if !b.cases[i][j].check {
				ret = false
				break
			}
		}
		if ret {
			return ret
		}
		ret = true
		for j := range b.cases[i] {
			if !b.cases[j][i].check {
				ret = false
				break
			}
		}
		if ret {
			return ret
		}
	}
	return ret
}

func (b *Board) check(value int) {
	for i := range b.cases {
		for j := range b.cases[i] {
			if b.cases[i][j].value == value {
				b.cases[i][j].check = true
			}
		}
	}
}

func initBoard() ([]Board, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	values := strings.Split(scanner.Text(), ",")
	ret := make([]int, len(values))
	for i := range values {
		tmp, _ := strconv.Atoi(values[i])
		ret[i] = tmp
	}
	boards := make([]Board, 0)
	scanner.Scan()
	tmp := Board{}
	tmp.cases = make([][]Case, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			boards = append(boards, tmp)
			tmp = Board{}
			tmp.cases = make([][]Case, 0)
		} else {
			val := strings.Split(line, " ")
			col := make([]Case, 0)
			for i := range val {
				if val[i] != "" {
					tmp, _ := strconv.Atoi(val[i])
					col = append(col, Case{tmp, false})
				}
			}
			tmp.cases = append(tmp.cases, col)
		}
	}
	boards = append(boards, tmp)

	return boards, ret
}

func (b *Board) getSum() int {
	ret := 0
	for i := range b.cases {
		for j := range b.cases[i] {
			if !b.cases[i][j].check {
				ret += b.cases[i][j].value
			}
		}
	}
	return ret
}

func remove(slice []Board, s int) []Board {
	return append(slice[:s], slice[s+1:]...)
}

func (b *Board) print() {
	for i := range b.cases {
		fmt.Println(b.cases[i])
	}
	fmt.Println()
}

func main() {
	boards, values := initBoard()
	for i := range values {
		for j := 0; j < len(boards); j++ {
			boards[j].check(values[i])
			if boards[j].checkVictory() {
				if len(boards) == 1 {
					fmt.Println(boards[0].getSum() * values[i])
					return
				}
				boards = remove(boards, j)
				j--
			}
		}
	}
}
