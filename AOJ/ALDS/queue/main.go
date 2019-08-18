package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type process struct {
	name string
	time int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	var (
		n, quantum int
		err        error
	)
	splited := strings.Split(sc.Text(), " ")
	n, err = strconv.Atoi(splited[0])
	if err != nil {
		panic(err)
	}
	quantum, err = strconv.Atoi(splited[1])
	if err != nil {
		panic(err)
	}

	q := make(chan process, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		splited = strings.Split(sc.Text(), " ")
		t, err := strconv.Atoi(splited[1])
		if err != nil {
			panic(err)
		}

		p := process{
			name: splited[0],
			time: t,
		}

		q <- p
	}

	progress := 0
	for p := range q {
		if p.time > quantum {
			p.time -= quantum
			progress += quantum
			q <- p
		} else {
			progress += p.time
			fmt.Printf("%s %d\n", p.name, progress)

			if len(q) == 0 {
				close(q)
			}
		}
	}
}
