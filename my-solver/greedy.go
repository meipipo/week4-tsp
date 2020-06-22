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

func copyDistances(distances map[int]map[int]float64) map[int]map[int]float64 {
	distancesCopied := make(map[int]map[int]float64)
	for k, v := range distances {
		for k2, v2 := range v {
			if _, ok := distancesCopied[k]; !ok {
				distancesCopied[k] = make(map[int]float64)
			}
			distancesCopied[k][k2] = v2
		}
	}
	return distancesCopied
}

func solveNN(cities map[int][]float64) ([][]string, float64) {
	distances := getDistance(cities)

	var solution [][]string
	minLen := 1000000.0
	for id := 0; id < len(cities); id++ {
		var solutionI [][]string
		solutionI = append(solutionI, []string{"index"})

		// Current city. Start with id.
		current := id
		solutionI = append(solutionI, []string{strconv.Itoa(current)})
		lengthI := 0.0

		// Visited or not.
		visited := make(map[int]bool)
		for i := 0; i < len(cities); i++ {
			visited[i] = false
		}
		visited[current] = true
		// Number of unvisited cities.
		unvisited := len(cities) - 1

		distancesCopied := copyDistances(distances)
		for unvisited > 0 {
			newID, newLen := min(distancesCopied[current])
			if visited[newID] {
				delete(distancesCopied[current], newID)
				continue
			}
			current = newID
			solutionI = append(solutionI, []string{strconv.Itoa(current)})
			visited[current] = true
			unvisited -= 1
			lengthI += newLen
		}
		// Return to the start city.
		solutionI = append(solutionI, []string{strconv.Itoa(id)})
		lengthI += distances[id][current]

		// Update if it is optimal.
		if lengthI < minLen {
			minLen = lengthI
			solution = solutionI
		}
	}

	return solution, minLen
}
