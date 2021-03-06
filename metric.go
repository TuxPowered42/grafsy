package main

import (
	"bufio"
	"os"
)

type MetricData struct {
	value  float64
	amount int64
}

// Reading metrics from file and remove file afterwords
func readMetricsFromFile(file string) []string {
	var results_list []string
	f, err := os.Open(file)

	if err != nil {
		f.Close()
		return results_list
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		results_list = append(results_list, scanner.Text())
	}
	f.Close()
	// Go will close file automatically?
	os.Remove(file)

	return results_list
}

// Reading metrics from file and remove file afterwords
func getSizeInLinesFromFile(file string) int {
	f, err := os.Open(file)
	defer f.Close()

	res := 0
	if err != nil {
		return res
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		res++
	}
	return res
}
