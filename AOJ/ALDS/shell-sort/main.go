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

	list := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		a, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		list = append(list, a)
	}

	// -- output --
	G, cnt := shellSort(list)
	fmt.Println(len(G))

	gs := ""
	for _, g := range G {
		gg := strconv.Itoa(g)
		gs = gs + gg + " "
	}
	gs = strings.TrimRight(gs, " ")
	fmt.Println(gs)

	fmt.Println(cnt)

	for _, l := range list {
		fmt.Println(l)
	}
}

// 間隔の式：https://web.archive.org/web/20170212072405/http://www.programming-magic.com/20100507074241/
func hClosure() func() int {
	h := 0
	return func() int {
		h = 3*h + 1
		return h
	}
}

func insertionSort(list []int, g int, cnt int) int {
	for i := g; i < len(list); i++ {
		v := list[i]
		j := i - g
		for ; j >= 0 && list[j] > v; j -= g {
			list[j+g] = list[j]
			cnt++
		}
		list[j+g] = v
	}
	return cnt
}

func shellSort(list []int) ([]int, int) {
	tmpG, G := []int{}, []int{}
	h := hClosure()
	for i := 0; i < len(list); i++ {
		hRange := h()
		if hRange > len(list) {
			break
		}
		tmpG = append(tmpG, hRange)
	}
	for i := len(tmpG) - 1; i >= 0; i-- {
		G = append(G, tmpG[i])
	}

	cnt := 0
	m := len(G)
	for i := 0; i < m; i++ {
		cnt = insertionSort(list, G[i], cnt)
	}

	return G, cnt
}
