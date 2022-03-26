package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Row []bool
type Graph []Row

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, _ := strconv.Atoi(f)
		n = append(n, i)
	}
	return n
}

func main() {
	var n, m, q int
	fmt.Scan(&n, &m, &q)

	var graph Graph
	for i := 0; i < n; i++ {
		var row Row
		for i := 0; i < n; i++ {
			row = append(row, false)
		}
		graph = append(graph, row)
	}

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		u -= 1
		v -= 1
		graph[u][v] = true
		graph[v][u] = true
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	colors := numbers(scanner.Text())

	for i := 0; i < q; i++ {
		scanner.Scan()
		query := numbers(scanner.Text())

		if query[0] == 1 {
			x := query[1] - 1
			fmt.Println(colors[x])

			for i := 0; i < n; i++ {
				if graph[x][i] {
					colors[i] = colors[x]
				}
			}
		}

		if query[0] == 2 {
			x := query[1] - 1
			y := query[2]
			fmt.Println(colors[x])
			colors[x] = y
		}
	}
}
