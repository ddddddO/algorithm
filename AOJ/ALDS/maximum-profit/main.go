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

	max := list[n-1] - list[n-2]
	for i := n - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			tmp := list[i] - list[j]
			if max < tmp {
				max = tmp
			}
		}
	}

	fmt.Println(max)
}
