package main

import (
	"bufio"
	"errors"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/madeadi/aoc2024/util"
)

func main() {
	array1 := readFile("input.txt")
	var start time.Time
	var totalSafe int
	var elapsed time.Duration

	// start = time.Now()
	// totalSafe = WithGoRoutine(array1)
	// elapsed = time.Since(start)
	// slog.Info("With Go Routine", "elapsed", elapsed, "totalSafe", totalSafe)

	start = time.Now()
	totalSafe = WithoutGoRoutine(array1)
	elapsed = time.Since(start)
	slog.Info("Without Go Routine", "elapsed", elapsed, "totalSafe", totalSafe)

}

func readFile(filename string) (array1 [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var array []int
		line := scanner.Text()
		fields := strings.Fields(line)
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				log.Fatal(err)
			}
			array = append(array, num)
		}
		array1 = append(array1, array)
	}

	return
}

func IsSafe(array []int) (int, error) {
	if len(array) < 2 {
		return 0, errors.New("array is less than 2")
	}

	asc := array[0] < array[1]
	// slog.Info("asc", "asc", asc, "line", array)

	for i := 0; i < len(array)-1; i++ {
		// slog.Info("inside", "item", array[i], "next", array[i+1], "asc", asc, "analysis", array[i] >= array[i+1])
		if asc && array[i] >= array[i+1] {
			return i, errors.New("array is not ascending")
		}

		if !asc && array[i] <= array[i+1] {
			return i, errors.New("array is not descending")
		}

		delta := util.Abs(array[i] - array[i+1])
		if delta == 0 || delta > 3 {
			return i, errors.New("distance is 0 or greater than 3")
		}

	}

	return 0, nil
}

func WithGoRoutine(array1 [][]int) int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	totalSafe := 0
	for _, line := range array1 {
		wg.Add(1)
		go func(line []int) {
			defer wg.Done()
			_, err := IsSafe(line)
			if err == nil {
				mu.Lock()
				totalSafe++
				mu.Unlock()
			}
		}(line)
	}

	wg.Wait()
	return totalSafe
}

func WithoutGoRoutine(array1 [][]int) int {
	totalSafe := 0
	for _, line := range array1 {

		errIndex, err := IsSafe(line)
		if err == nil {
			totalSafe++
		} else {
			err := Dampen(line, errIndex)
			if err == nil {
				totalSafe++
			}
		}

	}

	return totalSafe
}

func remove(slice []int, s int) []int {
	newSlice := make([]int, len(slice)-1)
	copy(newSlice, slice[:s])
	copy(newSlice[s:], slice[s+1:])
	return newSlice
}

func Dampen(slice []int, errIndex int) error {
	// index to remove
	itr := []int{0, errIndex}
	if errIndex+1 < len(slice) {
		itr = append(itr, errIndex+1)
	}

	for _, i := range itr {
		line := remove(slice, i)
		_, err := IsSafe(line)
		if err == nil {
			slog.Info("Can dampen", "line", slice, "dampened", line)
			return nil
		}
	}

	slog.Info("cannot dampen", "line", slice)
	return errors.New("cannot dampen")
}
