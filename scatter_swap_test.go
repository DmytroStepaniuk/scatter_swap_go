package scatter_swap

import (
	"testing"
)

func TestAverage(t *testing.T) {
	hashService := Service{Spin: 1}

	digits := []int{
		9,
		90,
		900,
		9000,
		90000,
		900000,
		9000000,
		90000000,
		900000000,
		9999999999,
	}

  for _, digit := range digits {
		result := hashService.HashToString(digit)

		if len(result) != 10 {
			t.Error(
				"For", digit,
				"expected", "string that contains 10 digits",
				"got", result,
			)
		}
  }
}
