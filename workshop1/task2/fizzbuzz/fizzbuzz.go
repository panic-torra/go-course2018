package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(loopSize int) (output string) {
	for i := 1; i <= loopSize; i++ {
		fizz := (i % 3) == 0
		buzz := (i % 5) == 0

		if fizz && buzz {
			output += "Fizz Buzz"
		} else if fizz {
			output += "Fizz"
		} else if buzz {
			output += "Buzz"
		} else {
			output += strconv.Itoa(i)
		}

		if i != loopSize {
			output += ", "
		}
	}

	return output
}
