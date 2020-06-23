package main

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

func solveNN(n int, distances map[int]map[int]float64) ([]int, float64) {
	var solution []int
	minLen := 1000000.0
	for id := 0; id < n; id++ {
		var solutionI []int

		// Current city. Start with id.
		current := id
		solutionI = append(solutionI, current)
		lengthI := 0.0

		// Visited or not.
		visited := make(map[int]bool)
		for i := 0; i < n; i++ {
			visited[i] = false
		}
		visited[current] = true
		// Number of unvisited cities.
		unvisited := n - 1

		distancesCopied := copyDistances(distances)
		for unvisited > 0 {
			newID, newLen := min(distancesCopied[current])
			if visited[newID] {
				delete(distancesCopied[current], newID)
				continue
			}
			current = newID
			solutionI = append(solutionI, current)
			visited[current] = true
			unvisited -= 1
			lengthI += newLen
		}
		// Add path to go back to the start city.
		lengthI += distances[id][current]

		// Update if it is optimal.
		if lengthI < minLen {
			minLen = lengthI
			solution = solutionI
		}
	}

	return solution, minLen
}
