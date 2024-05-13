package main

import "fmt"

var (
	n     int
	lista [1010]int
	p     [1010]int
)

func calcula(a, b, corAtual, pos, resAtual int) int {
	if pos > n {
		return resAtual
	}
	if lista[pos] == corAtual {
		resAtual++
		if corAtual == a {
			corAtual = b
		} else {
			corAtual = a
		}
	}
	if pos == n {
		return resAtual
	}
	return calcula(a, b, corAtual, pos+1, resAtual)
}

func main() {
	var res = 1
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		p[i] = -1
	}

	for i := 1; i <= n; i++ {
		var a int
		fmt.Scan(&a)
		lista[i] = a
		if p[a] == -1 {
			p[a] = i
		}
	}

	for i := 1; i <= n; i++ {
		if p[i] != -1 {
			for j := i + 1; j <= n; j++ {
				if p[j] != -1 {
					res1 := calcula(i, j, i, p[i], 0)
					res2 := calcula(j, i, j, p[j], 0)
					resAtual := max(res1, res2)
					if resAtual > res {
						res = resAtual
					}
				}
			}
		}
	}

	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
