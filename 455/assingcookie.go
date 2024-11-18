package main

import "sort"

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	findContentChildren(g, s)
}
func findContentChildren(g []int, s []int) int {
	count := 0
	sort.Ints(g) //reason for sorting is some test is not wokring

	sort.Ints(s)

	for i := count; i < len(s); i++ {

		for j := count; j < len(g); j++ {

			if s[i] >= g[j] {
				count++

				break

			}
		}
	}

	return count
}
