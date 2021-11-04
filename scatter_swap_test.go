package scatter_swap

import (
	"testing"
)

func TestAverage(t *testing.T) {
	hashService := Service{Spin: 1}

  for i := 0; i <= 9_999_999_999; i++ {
		result := hashService.HashToString(i)

		if len(result) != 10 {
			t.Error(
				"For", i,
				"expected", "string that contains 10 digits",
				"got", result,
			)
		}
  }
}
