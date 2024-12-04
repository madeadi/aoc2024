package main

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
	"sort"
	"time"

	"github.com/madeadi/aoc2024/util"
)

func main() {
	part1()
	part2()
}

func readFile() (array1, array2 []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
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

	return
}

func part1() {
	start := time.Now()
	array1, array2 := readFile()

	sort.Ints(array1)
	sort.Ints(array2)

	totalDistance := 0
	// loop through array1 and array2
	for i := 0; i < len(array1); i++ {
		totalDistance += util.Abs(array1[i] - array2[i])
	}

	elapsed := time.Since(start)
	slog.Info("part1", "totalDistance", totalDistance, "elapsed", elapsed)
}

func part2() {
	start := time.Now()
	array1, array2 := readFile()

	sort.Ints(array1)
	sort.Ints(array2)

	totalSimilarity := 0
	for i := 0; i < len(array1); i++ {
		count := 0
		for j := 0; j < len(array2); j++ {
			if array2[j] == array1[i] {
				count++
			} else if array2[j] > array1[i] {
				totalSimilarity += (count * array1[i])
				break
			} else {
				continue
			}

		}
	}

	elapsed := time.Since(start)
	slog.Info("part2", "elapsed", elapsed, "similarity", totalSimilarity)

}
