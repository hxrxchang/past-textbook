package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, _ := strconv.Atoi(f)
		n = append(n, i)
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	scanner.Scan()
	cards := numbers(scanner.Text())
	fmt.Println(cards)

	scanner.Scan()
	Q, _ := strconv.Atoi(scanner.Text())
}
