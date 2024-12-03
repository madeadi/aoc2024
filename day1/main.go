package main

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
	"sort"
)

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var array1, array2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var num1, num2 int
		_, err := fmt.Sscanf(line, "%d %d", &num1, &num2)
		if err != nil {
			log.Fatal(err)
		}

		array1 = append(array1, num1)
		array2 = append(array2, num2)
	}

	sort.Ints(array1)
	sort.Ints(array2)

	totalDistance := 0
	// loop through array1 and array2
	for i := 0; i < len(array1); i++ {
		totalDistance += abs(array1[i] - array2[i])
	}

	slog.Info("totalDistance", "totalDistance", totalDistance)
}
