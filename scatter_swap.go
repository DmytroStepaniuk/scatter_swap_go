package scatter_swap

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	Service struct {
		Spin int
	}
)

// New makes new service instance
func New(spin int) Service {
	return Service{
		Spin: spin,
	}
}

// Hash returns array of digits
func (s Service) Hash(digit int) []int {
	zeroPad := fmt.Sprintf("%010d", digit)

	arrayOfStrings := strings.Split(zeroPad, "")
	var arrayOfDigits []int

	// convert strings to integers
	for _, el := range arrayOfStrings {
		newEl, _ := strconv.Atoi(el)
		arrayOfDigits = append(arrayOfDigits, newEl)
	}

	return s.swap(arrayOfDigits)
}

// HashToString returns hashed string
func (s Service) HashToString(digit int) string {
	result := s.Hash(digit)

	// convert all integers to strings
	var tmpArray []string
	for _, el := range result {
		tmpArray = append(tmpArray, strconv.Itoa(el))
	}

	return strings.Join(tmpArray, "")
}

// Unhash de-obfuscates string into array of integers
func (s Service) Unhash(digit string) []int {
	workingArray := strings.Split(digit, "")
	var array []int

	// convert all integers to strings
	for _, el := range workingArray {
		newEl, _ := strconv.Atoi(el)
		array = append(array, newEl)
	}

	result := s.unscatter(array)
	result = s.unswap(result)

	return result
}

// UnhashToInt de-obfuscates string into integer
func (s Service) UnhashToInt(digit string) int {
	result := s.Unhash(digit)

	var tmpArray []string
	for _, el := range result {
		tmpArray = append(tmpArray, strconv.Itoa(el))
	}

	digitAsAString := strings.Join(tmpArray, "")

	value, _ := strconv.Atoi(digitAsAString)

	return value
}

func (s Service) swapperMap(index int) []int {
	var newArray []int
	workingArray := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < 10; i++ {
		maxJ := index + i ^ s.Spin
		lastElIndex := len(workingArray) - 1

		maxJ = maxJ % len(workingArray)
		for j := 0; j < maxJ; j++ {
			workingArray = append(workingArray[1:], workingArray[0]) // rotate array from the right to left
		}

		newArray = append(newArray, workingArray[lastElIndex]) // push last element
		workingArray = workingArray[:lastElIndex]              // drop last element
	}

	return newArray
}

func (s Service) swap(workingArray []int) []int {
	var newArray []int

	for index, digit := range workingArray {
		resultArray := s.swapperMap(index)
		result := resultArray[digit]
		newArray = append(newArray, result)
	}

	newArray = s.scatter(newArray)

	return newArray
}

func (s Service) scatter(workingArray []int) []int {
	var sumOfDigits int = 0

	for _, el := range workingArray {
		sumOfDigits += el
	}

	var newArray []int

	for i := 0; i < 10; i++ {
		maxJ := s.Spin ^ sumOfDigits
		maxJ = maxJ % len(workingArray)
		lastElIndex := len(workingArray) - 1

		for j := 0; j < maxJ; j++ {
			workingArray = append(workingArray[1:], workingArray[0]) // rotate array from the right to left
		}

		newArray = append(newArray, workingArray[lastElIndex]) // push last element
		workingArray = workingArray[:lastElIndex]              // drop last element
	}

	return newArray
}

func (s Service) unscatter(workingArray []int) []int {
	var sumOfDigits int = 0

	for _, el := range workingArray {
		sumOfDigits += el
	}

	var newArray []int

	for i := 0; i < 10; i++ {
		lastElIndex := len(workingArray) - 1
		newArray = append(newArray, workingArray[lastElIndex])

		workingArray = workingArray[:lastElIndex]

		maxJ := (sumOfDigits ^ s.Spin) % len(newArray)
		for j := 0; j < maxJ; j++ {
			lastElIndex := len(newArray) - 1
			newArray = append([]int{newArray[lastElIndex]}, newArray[:lastElIndex]...) // rotate array from the left to right
		}
	}

	return newArray
}

func (s Service) unswap(workingArray []int) []int {
	var newArray []int

	for idx, digit := range workingArray {
		array := s.swapperMap(idx)

		// find in array first match, consider direction from the right to left
		for index := len(array) - 1; index >= 0; index-- {
			if digit == array[index] {
				newArray = append(newArray, index)
				break
			}
		}
	}

	return newArray
}
