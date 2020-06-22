package main

import (
	"strconv"
)

func copyMapCities(cities map[int][]float64) map[int][]float64 {
	citiesCopied := make(map[int][]float64)
	for k, v := range cities {
		citiesCopied[k] = v
	}
	return citiesCopied
}

func solveGreedy(cities map[int][]float64) ([][]string, float64) {
	var solution [][]string

	minLen := 1000000.0
	for id := 0; id < len(cities); id++ {
		var solutionI [][]string
		solutionI = append(solutionI, []string{"index"})
		lengthI := 0.0
		citiesCopied := copyMapCities(cities)
		x0 := citiesCopied[id][0]
		y0 := citiesCopied[id][1]
		for {
			solutionI = append(solutionI, []string{strconv.Itoa(id)})
			lengths := distancesFrom(id, citiesCopied)
			newID, minLen := min(lengths)
			if len(citiesCopied) == 2 {
				solutionI = append(solutionI, []string{strconv.Itoa(newID)})
				lengthI += minLen
				xq := citiesCopied[newID][0]
				yq := citiesCopied[newID][1]
				lengthI += distance(x0, y0, xq, yq)
				break
			}
			delete(citiesCopied, id)
			id = newID
			lengthI += minLen
		}

		if lengthI < minLen {
			minLen = lengthI
			solution = solutionI
		}
	}

	return solution, minLen
}

func distancesFrom(i int, cities map[int][]float64) map[int]float64 {
	lengths := make(map[int]float64)
	x := cities[i][0]
	y := cities[i][1]
	for k, v := range cities {
		if k == i {
			continue
		}
		lengths[k] = distance(x, y, v[0], v[1])
	}
	return lengths
}
