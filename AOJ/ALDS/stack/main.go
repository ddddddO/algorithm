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
	splited := strings.Split(sc.Text(), " ")

	stack := []int{}
	for _, sp := range splited {
		if !isOperator(sp) {
			num, err := strconv.Atoi(sp)
			if err != nil {
				panic(err)
			}
			stack = append(stack, num)
			continue
		}

		sum := calc(sp, stack[len(stack)-2], stack[len(stack)-1])
		stack = stack[:len(stack)-2]
		stack = append(stack, sum)
	}

	fmt.Println(stack[len(stack)-1])
}

func isOperator(s string) bool {
	flg := false

	switch s {
	case "+":
		flg = true
	case "-":
		flg = true
	case "*":
		flg = true
	default:
		flg = false
	}
	return flg
}

func calc(op string, val0, val1 int) int {
	rslt := 0
	switch op {
	case "+":
		rslt = val0 + val1
	case "-":
		rslt = val0 - val1
	case "*":
		rslt = val0 * val1
	default:
		panic("XXX")
	}

	return rslt
}
