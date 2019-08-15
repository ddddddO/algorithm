package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var (
		err error
	)
	if sc.Scan() {
		_, err = strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
	}

	list := []string{}
	if sc.Scan() {
		list = strings.Split(sc.Text(), " ")
	}

	checks := [][]string{}
	for i := range list {
		iv := strings.TrimLeftFunc(list[i], func(r rune) bool {
			return unicode.IsUpper(r)
		})
		if err != nil {
			panic(err)
		}

		exist := false
		for _, check := range checks {
			for _, c := range check {
				if strings.Contains(c, iv) {
					exist = true
					break
				}
			}
			if exist {
				break
			}
		}
		if exist {
			continue
		}

		check := []string{list[i]}
		for j := i + 1; j < len(list); j++ {
			if strings.Contains(list[j], iv) {
				check = append(check, list[j])
			}
		}
		if len(check) > 1 {
			checks = append(checks, check)
		}
	}

	listForBubble := make([]string, len(list))
	copy(listForBubble, list)

	listForSelect := make([]string, len(list))
	copy(listForSelect, list)

	sortedBubble := bubbleSort(listForBubble)
	sb := ""
	for _, ss := range sortedBubble {
		sb = sb + ss + " "
	}
	sb = strings.TrimRight(sb, " ")
	fmt.Println(sb)
	isStable(checks, sortedBubble)

	sortedSelect := selectionSort(listForSelect)
	ss := ""
	for _, sss := range sortedSelect {
		ss = ss + sss + " "
	}
	ss = strings.TrimRight(ss, " ")
	fmt.Println(ss)
	isStable(checks, sortedSelect)
}

func bubbleSort(list []string) []string {
	for i := 0; i < len(list); i++ {
		for j := len(list) - 1; j > i; j-- {
			pre, err := strconv.Atoi(strings.TrimLeftFunc(list[j-1], func(r rune) bool {
				return unicode.IsUpper(r)
			}))
			if err != nil {
				panic(err)
			}

			crr, err := strconv.Atoi(strings.TrimLeftFunc(list[j], func(r rune) bool {
				return unicode.IsUpper(r)
			}))
			if err != nil {
				panic(err)
			}

			if pre > crr {
				tmp := list[j-1]
				list[j-1] = list[j]
				list[j] = tmp
			}
		}
	}

	return list
}

func selectionSort(list []string) []string {
	for i := 0; i < len(list); i++ {
		min := i
		for j := i; j < len(list); j++ {
			minv, err := strconv.Atoi(strings.TrimLeftFunc(list[min], func(r rune) bool {
				return unicode.IsUpper(r)
			}))
			if err != nil {
				panic(err)
			}

			jv, err := strconv.Atoi(strings.TrimLeftFunc(list[j], func(r rune) bool {
				return unicode.IsUpper(r)
			}))
			if err != nil {
				panic(err)
			}

			if jv < minv {
				min = j
			}
		}
		//if i != min {
		tmp := list[i]
		list[i] = list[min]
		list[min] = tmp
		//}
	}
	return list
}

func isStable(checks [][]string, target []string) {
	for _, check := range checks {
		pre := -1
		for _, c := range check {
			for i, t := range target {
				if strings.Contains(t, c) {
					if pre > i {
						fmt.Println("Not stable")
						return
					}
					pre = i
					break
				}
			}
		}
	}
	fmt.Println("Stable")
}
