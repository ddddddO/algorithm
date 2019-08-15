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

	var (
		n   int
		err error
	)
	if sc.Scan() {
		n, err = strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
	}

	list := make([]int, n)
	if sc.Scan() {
		splited := strings.Split(sc.Text(), " ")
		for i, v := range splited {
			vv, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			list[i] = vv
		}
	}

	cnt, sorted := selectionSort(list)

	s := ""
	for _, ss := range sorted {
		s = s + strconv.Itoa(ss) + " "
	}
	s = strings.TrimRight(s, " ")

	fmt.Println(s)
	fmt.Println(cnt)
}

func selectionSort(list []int) (int, []int) {
	cnt := 0
	for i := 0; i < len(list); i++ {
		min := i
		for j := i + 1; j < len(list); j++ {
			if list[min] > list[j] {
				min = j
			}
		}
		if i != min {
			tmp := list[i]
			list[i] = list[min]
			list[min] = tmp
			cnt++
		}
	}

	return cnt, list
}
