package main

import "fmt"

func duplicateZeros(arr []int) {

	for i := 0; i < len(arr); i++ {

		if arr[i] == 0 {
			fmt.Println("index of ", i)

			for j := len(arr) - 1; j > i; j-- {

				arr[j] = arr[j-1]

			}
			i++
		}

	}

}