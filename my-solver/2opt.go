package main

func twoOpt(n int, distances map[int]map[int]float64, solution []int, length float64) ([]int, float64) {
	for i := 0; i < n-2; i++ {
		u1, v1 := solution[i], solution[i+1]
		for j := i + 1; j < n; j++ {
			u2, v2 := solution[j], solution[(j+1)%n]
			d1, d2 := distances[u1][v1], distances[u2][v2] // current path
			d3, d4 := distances[u1][u2], distances[v1][v2] // candidate to swap
			if d3+d4 < d1+d2 {
				l := i + 1
				r := j
				for l < r { // update and reverse.
					solution[l], solution[r] = solution[r], solution[l]
					l += 1
					r -= 1
				}
				length = length - (d1 + d2) + (d3 + d4)
			}
		}
	}
	return solution, length
}
