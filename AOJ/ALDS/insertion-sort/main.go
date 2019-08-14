package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	inputN := sc.Text()
	sc.Scan()
	inputLine := sc.Text()

	n, err := strconv.Atoi(inputN)
	if err != nil {
		panic(err)
	}

	list := []int{}
	newInput := strings.Split(inputLine, " ")
	for _, v := range newInput {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		list = append(list, num)
	}

	insertionSort(list, n)
}

func insertionSort(list []int, n int) {
	listDisplay(list)

	for i := 1; i < n; i++ {
		v := list[i]
		j := i - 1
		for ; j >= 0; j-- {
			if list[j] <= v {
				break
			}
			list[j+1] = list[j]
		}
		list[j+1] = v
		listDisplay(list)
	}
}

func listDisplay(list []int) {
	row := ""
	for _, v := range list {
		row += fmt.Sprintf("%d ", v)
	}
	row = strings.TrimRight(row, " ")
	fmt.Println(row)
}
