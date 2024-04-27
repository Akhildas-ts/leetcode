package main

import "fmt"


func main() {
    s := []byte{'h','e','l','l','o'}
	reverseString(s)
}

func reverseString(s []byte) {
	

	for i := 0; i < len(s)/2; i++ {
        s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
    }
	fmt.Println(string(s))

}