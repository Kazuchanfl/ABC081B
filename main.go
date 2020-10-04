package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	var a []int
	var times int = 0
	var result bool = true

	fmt.Fscan(r, &n)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	for result  {
		for i := 0; i < n; i++ {
			if a[i] % 2 == 1 {
				result = false
			}
			a[i] /= 2
		}
		if result {
			times ++
		}
	}

	fmt.Fprintln(w, times)
}
