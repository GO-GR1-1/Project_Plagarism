package main

import (
	"fmt"

	//"github.com/agnivade/levenshtein"
)

func main() {
	s1 := "Il fait beau"
	s2 := "Il a fait beau"
	distance := levenshtein.ComputeDistance(s1, s2)
	fmt.Printf("The distance between %s and %s is %d.\n", s1, s2, distance)
	// Output:
	// The distance between kitten and sitting is 3.
}
