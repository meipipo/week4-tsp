package main

import (
	"encoding/csv"
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

func solve(cities map[int][]float64) [][]string {
	var solution [][]string
	solution = append(solution, []string{"index"})
	for k := range cities {
		solution = append(solution, []string{strconv.Itoa(k)})
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
	var i string
	fmt.Scan(&i)
	cities := read_cities(i)

	solution := solve(cities)
	write_solution(i, solution)
}
