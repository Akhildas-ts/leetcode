package main

import (
	"fmt"
	"strconv"
)

// question is we need to find the gretest value in these case ..

func main() {
	i := []int{1, 5, 2, 6, 8}
	fmt.Println(findRelativeRanks(i))
}




func findRelativeRanks(score []int) []string {
	var frist int
	var second int
	var third int

	outPut := make([]string, 0)
	for i := 0; i < len(score); i++ {

		if frist < score[i] {

			third = second
			second = frist
			frist = score[i]
		}

	}

	for j := 0; j < len(score); j++ {

		str := strconv.Itoa(score[j])

		if score[j] == frist {

			outPut = append(outPut, "frist")
		} else if score[j] == second {

			outPut = append(outPut, "second")
		} else if score[j] == third {
			outPut = append(outPut, "third")
		} else {
			outPut = append(outPut, str)

		}

	}

	return outPut

}
