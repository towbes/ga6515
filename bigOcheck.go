
package main

import (
	"log"
	"time"
	"math"
)

// trackTime Taken from https://medium.com/better-programming/dynamic-programming-in-go-a95d32ee9953
// Track execution time
func trackTime(start time.Time, result float64, name string) {
	elapsed := time.Since(start)
	log.Printf("---> %s solution | result: %v | took %s", name, result, elapsed)
}


//f(n)
func fofn(n float64) float64 {

	result := (math.Pow(n, 2)) / math.Log(n)
	defer trackTime (time.Now(), result, "F of n")

	return result
}

//g(n)
func gofn(n float64) float64 {

	result := n * math.Pow(math.Log(n), 2)
	defer trackTime (time.Now(), result, "G of n")

	return result
}

func main() {
	input := float64(99999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999)

	fofn(input)
	gofn(input)
}