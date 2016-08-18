package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
)


func CountInversions(input []int) (int, []int) {

	if len(input) < 2 {
		return 0, input
	}

	cutpoint := len(input) / 2

	inversions := 0

	leftInversions, left := CountInversions(input[:cutpoint])
	rightInversions, right := CountInversions(input[cutpoint:])

	inversions += leftInversions + rightInversions

	sorted := []int{}

	i, j := 0, 0
	for i < len(left) || j < len(right) {
		if i >= len(left) {
			sorted = append(sorted, right[j:]...)
			break
		}
		if j >=len(right) {
			sorted = append(sorted, left[i:]...)
			break
		}

		if left[i] <= right[j] {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
			inversions += len(left) - i
		}
	}

	return inversions, sorted

}

func main() {

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	invs,_ := CountInversions(numbers)

	//fmt.Println(numbers)
	//fmt.Println(sorted)
	fmt.Printf("Number of Inversions: %d\n", invs)

}


