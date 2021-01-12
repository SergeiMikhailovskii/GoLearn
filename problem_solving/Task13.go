package main

import "fmt"

func main() {
	var m, n, a, b, k, max, x int

	_, _ = fmt.Scanf("%d", &n)
	_, _ = fmt.Scanf("%d", &m)

	arr := make([]int, n+1)

	for i := 0; i < m; i++ {
		_, _ = fmt.Scanf("%d", &a)
		_, _ = fmt.Scanf("%d", &b)
		_, _ = fmt.Scanf("%d", &k)

		arr[a] += k

		if b+1 <= n {
			arr[b+1] -= k
		}
	}

	for i := 1; i <= n; i++ {
		x += arr[i]
		if x > max {
			max = x
		}
	}
	fmt.Println(max)
}
