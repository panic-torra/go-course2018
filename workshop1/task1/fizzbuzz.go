package task1

import (
	"fmt"
	"strconv"
)

func task1() {
	const LoopSize = 50
	var output string

	for i := 1; i <= LoopSize; i++ {
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

		if i != LoopSize {
			output += ", "
		}
	}

	fmt.Println(output)
}
