package main

import "math"

func countSquares(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0

	}

	cnt := 0
	a := len(matrix)
	b := len(matrix[0])

	d := 1
	t := int(math.Min(float64(a), float64(b)))
	for d <= t {
		for i := 0; i < a; i++ {

			for j := 0; j < b; j++ {
				if (i+d <= a) && (j+d <= b) {
					x := 0
					for k := 0; k < d; k++ {
						for kk := 0; kk < d; kk++ {
							x += matrix[i+k][j+kk]
						}
					}
					if x == d*d {
						cnt += 1
					}
				}
			}
		}
		d += 1
	}
	return cnt
}
