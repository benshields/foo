package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	rands := []int{
		12, 28, 0, 63, 26, 38, 64, 17, 74, 67, 51, 44, 77, 32, 6, 10, 52, 47, 61, 46, 50, 29, 15, 1, 39, 37, 13, 66, 45, 8, 68, 96, 53, 40, 76, 72, 21, 93, 16, 83, 62, 48, 11, 9, 20, 36, 91, 19, 5, 42, 99, 84, 4, 95, 92, 89, 7, 71, 34, 35, 55, 22, 59, 18, 49, 14, 54, 85, 82, 58, 24, 73, 31, 97, 69, 43, 65, 27, 81, 56, 87, 70, 33, 88, 60, 2, 75, 90, 57, 94, 23, 30, 78, 80, 41, 3, 98, 25, 79, 86,
	}
	fmt.Println(len(rands))

	input, err := ioutil.ReadFile("./04/board.txt")
	if err != nil {
		panic(err)
	}

	boards := strings.Split(string(input), "\n\n")
	for _, board := range boards {
		fmt.Println(board)
		rows := strings.Split(board, "\n")
		cols := make([][]string, 5)
		for ri, row := range rows {
			nums := strings.Split(row, " ")
			for ci, num := range nums {
				cols[ci][ri] = num
			}
			fmt.Println("row:", row)
		}
		for _, col := range cols {
			fmt.Println("col:", col)
		}
	}

	/*
		fmt.Println(len(inp))

		a := string(inp)
		b := strings.Split(a, "\n\n")
		fmt.Println("len(b):", len(b))
		fmt.Println("b[0]:", b[0])
		fmt.Println()

		c := strings.Split(b[0], "\n")
		fmt.Println("c[0]:", c[0])
		fmt.Println()

		d := strings.Split(c[0], " ")
		fmt.Println("d:", d)
		fmt.Println()
	*/
}
