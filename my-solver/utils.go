package main

import (
	"math"
)

func min(lengths map[int]float64) (int, float64) {
	minLen := 1000000.0 // from previous score
	minID := -1
	for k, v := range lengths {
		if v < minLen {
			minLen = v
			minID = k
		}
	}
	return minID, minLen
}

func distance(x0, y0, x1, y1 float64) float64 {
	return math.Sqrt(math.Pow((x1-x0), 2) + math.Pow((y1-y0), 2))
}

func getDistance(cities map[int][]float64) map[int]map[int]float64 {
	dist := make(map[int]map[int]float64)
	for i := 0; i < len(cities); i++ {
		for j := i + 1; j < len(cities); j++ {
			dist[i][j] = distance(cities[i][0], cities[i][1], cities[j][0], cities[j][1])
			dist[j][i] = distance(cities[i][0], cities[i][1], cities[j][0], cities[j][1])
		}
	}
	return dist
}
