package main

import (
	"encoding/csv"
	"flag"
	"fmt"
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

func csvWritable(solutionInt []int) [][]string {
	var solution [][]string
	solution = append(solution, []string{"index"})
	for _, v := range solutionInt {
		solution = append(solution, []string{strconv.Itoa(v)})
	}
	return solution
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

func main() {
	var (
		i      = flag.String("input", "0", "input number 0-6")
		solver = flag.String("solver", "", "solver name")
	)
	flag.Parse()
	cities := read_cities(*i)

	var solution [][]string
	distances := getDistance(cities)
	n := len(cities)
	switch *solver {
	case "nn":
		solutionInt, length := solveNN(n, distances)
		fmt.Println(length)
		solution = csvWritable(solutionInt)
	case "2opt":
		solutionInt, length := solveNN(n, distances)
		solutionInt, length = twoOpt(n, distances, solutionInt, length)
		solution = csvWritable(solutionInt)
		fmt.Println(length)
	case "dist": // TODO
		edges := makeEdges(getDistance(cities))
		sortEdges(edges)
		return
	default:
		solution, _ = solve(cities)
		return
	}
	write_solution(*i, solution)
}
