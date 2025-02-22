package maths

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func HandleArgs(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("wrong number of arguments")
	}
	if args[1] != "data.txt" {
		return fmt.Errorf("wrong filename")
	}
	file, err := ReadFile(args[1])
	if err != nil {
		return err
	}
	average := MyAverage(file)
	fmt.Printf("Average: %d\n", int(average))
	median := MyMedian(file)
	fmt.Printf("Median: %d\n", int(median))
	variance := MyVariance(file, average)
	fmt.Printf("Variance: %d\n", int(variance))
	deviation := MyStandardDeviation(file, variance)
	fmt.Printf("Standard Deviation: %d\n", int(deviation))
	return nil
}

func ReadFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result []int
	var vlakeia error
	var temp int
	for scanner.Scan() {
		temp, vlakeia = strconv.Atoi(scanner.Text())
		if vlakeia != nil {
			return nil, vlakeia
		}
		result = append(result, temp)
	}
	return result, nil
}

func MyAverage(result []int) float64 {
	var average float64
	var sum int
	for _, v := range result {
		sum += v
	}
	if len(result) > 0 {
		average = float64(sum) / float64(len(result))
	}
	average = math.Round(average)
	return average
}

func MyMedian(result []int) float64 {
	var median float64
	if len(result) > 0 {
		sort.Ints(result)
		if len(result)%2 != 0 {
			median = float64(result[len(result)/2])
		} else {
			median = float64((result[len(result)/2] + result[len(result)/2-1])) / 2
		}
	}
	median = math.Round(median)
	return median
}

func MyVariance(result []int, average float64) float64 {
	var variance float64
	if len(result) > 0 {
		for i := 0; i < len(result); i++ {
			variance += math.Pow(float64(result[i])-average, 2)
		}
		variance /= float64(len(result))
	}
	variance = math.Round(variance)
	return variance
}

func MyStandardDeviation(result []int, variance float64) float64 {
	var deviation float64 = math.Sqrt(variance)
	deviation = math.Round(deviation)
	return deviation
}
