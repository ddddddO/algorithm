package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	for i := 0; i < n; i++ {
		sc.Scan()
		in, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		list[i] = in
	}

	// O(n)
	min := list[0]
	max := list[n-1] - list[n-2]
	for i := 1; i < n; i++ {
		if max < list[i]-min {
			max = list[i] - min
		}

		if min > list[i] {
			min = list[i]
		}
	}

	/*
		// O(n^2)
		max := list[n-1] - list[n-2]
		for i := n - 1; i > 0; i-- {
			for j := i - 1; j >= 0; j-- {
				tmp := list[i] - list[j]
				if max < tmp {
					max = tmp
				}
			}
		}
	*/
	fmt.Println(max)
}
