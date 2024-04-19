package main

// question
// find the sum of binary

// algoritham

// in these case add binary
// 1+1 = 0 and reminder =1
// 0+1 // 1+0 = 1
// 0+0 = 0

// binary numbers are gettng to the string
// store a number into a variable and check

func addBinary(a string, b string) string {

	sum := make([]int, 0)

	count := -1
	reminder := 0

	for i := 0; i < len(a+b); i++ {

		if a == "1" && b == "1" || b == "1" && a == "1" {

			if reminder > 0 {

				sum = append(sum, 1)
				reminder = 1
				continue
			}
			count = 0
			sum = append(sum, count)
			reminder = 1
		}

		if a == "0" && b == "0" || b == "0" && a == "0" {

			if reminder > 0 {
				sum = append(sum, 1)
				reminder = 0
				continue
			}

			count = 0

			sum = append(sum, count)
		}

		if a == "1" && b == "0" || b == "1" && a == "0" {

			if reminder > 0 {

				sum = append(sum, 0)
				reminder = 1
				continue
			}

			count = 1

			sum = append(sum, count)
		}
	}

  if a == "" && b == "" {
	break 
  }

  

}
