package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

func read_cities(i string) map[int][]float64 {
	path := "../input_" + i + ".csv"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var line []string
	reader := csv.NewReader(file)
	n := -1
	cities := make(map[int][]float64)
	for {
		line, err = reader.Read()
		if n == -1 {
			n++
			continue
		}
		if err != nil {
			break
		}
		if x, err := strconv.ParseFloat(line[0], 32); err == nil {
			cities[n] = append(cities[n], x)
		}
		if y, err := strconv.ParseFloat(line[1], 32); err == nil {
			cities[n] = append(cities[n], y)
		}
		n++
	}
	return cities
}

// solve do nothing.
func solve(cities map[int][]float64) ([][]string, float64) {
	var solution [][]string
	solution = append(solution, []string{"index"})
	for k := range cities {
		solution = append(solution, []string{strconv.Itoa(k)})
	}
	return solution, 0
}

func write_solution(i string, solution [][]string) {
	path := "../solution_yours_" + i + ".csv"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err := writer.WriteAll(solution); err != nil {
		panic(err)
	}
}

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

func dist(x0, y0, x1, y1 float64) float64 {
	return math.Sqrt(math.Pow((x1-x0), 2) + math.Pow((y1-y0), 2))
}

func main() {
	var i string
	fmt.Scan(&i)
	cities := read_cities(i)

	var option string
	fmt.Scan(&option)
	var solution [][]string
	var length float64
	switch option {
	case "greedy":
		solution, length = solveGreedy(cities)
	default:
		solution, length = solve(cities)
	}
	write_solution(i, solution)
	fmt.Println(length)
}
