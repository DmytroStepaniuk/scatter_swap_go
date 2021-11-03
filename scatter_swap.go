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

func New(spin int) Service {
	return Service{
		Spin: spin,
	}
}

func (s Service) Hash(digit int) []int {
	zeroPad := fmt.Sprintf("%010d", digit)

	arrayOfStrings := strings.Split(zeroPad, "")
	var arrayOfDigits []int

	for _, el := range arrayOfStrings {
		newEl, _ := strconv.Atoi(el)
		arrayOfDigits = append(arrayOfDigits, newEl)
	}

	return s.swap(arrayOfDigits)
}

func (s Service) HashToString(digit int) string {
	result := s.Hash(digit)

	var tmpArray []string
	for _, el := range result {
		tmpArray = append(tmpArray, strconv.Itoa(el))
	}

	return strings.Join(tmpArray, "")
}

func (s Service) Unhash(digit string) []int {
	workingArray := strings.Split(digit, "")
	var array []int

	for _, el := range workingArray {
		newEl, _ := strconv.Atoi(el)
		array = append(array, newEl)
	}

	result := s.unscatter(array)
	result = s.unswap(result)

	return result
}

func (s Service) swapperMap(index int) []int {
	var newArray []int
	workingArray := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < 10; i++ {
		maxJ := index + i ^ s.Spin
		lastElIndex := len(workingArray) - 1

		maxJ = maxJ % len(workingArray)
		for j := 0; j < maxJ; j++ {
			workingArray = append(workingArray[1:], workingArray[0])
		}

		newArray = append(newArray, workingArray[lastElIndex])
		workingArray = workingArray[:lastElIndex]
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
			workingArray = append(workingArray[1:], workingArray[0])
		}

		newArray = append(newArray, workingArray[lastElIndex])
		workingArray = workingArray[:lastElIndex]
	}

	return newArray
}

func (s Service) unscatter(workingArray []int) []int {
	var sumOfDigits int = 0

	for _, el := range workingArray { sumOfDigits += el }

	var newArray []int

	for i := 0; i < 10; i++ {
		lastElIndex := len(workingArray) - 1
		newArray = append(newArray, workingArray[lastElIndex])

		workingArray = workingArray[:lastElIndex]

		maxJ := (sumOfDigits ^ s.Spin) % len(newArray)
		for j := 0; j < maxJ; j++ {
			lastElIndex := len(newArray) - 1
			newArray = append([]int{newArray[lastElIndex]}, newArray[:lastElIndex]...)
		}
	}

	return newArray
}

func (s Service) unswap(workingArray []int) []int {
	var newArray []int

	for idx, digit := range workingArray {
		array := s.swapperMap(idx)

		for index := len(array)-1; index >= 0 ; index-- {
			if digit == array[index] {
				newArray = append(newArray, index)
				break
			}
		}
	}

	return newArray
}
