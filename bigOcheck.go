
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

	result := math.Pow(n, 0.1)
	defer trackTime (time.Now(), result, "F of n")

	return result
}

//g(n)
func gofn(n float64) float64 {

	result := math.Pow(math.Log(n), 10)
	defer trackTime (time.Now(), result, "G of n")

	return result
}

func main() {

	inputs := make([]float64, 10)
	timesten := float64(1)

	for i := 0; i < 10; i++ {
		inputs[i] = float64(1) * math.Pow(10, timesten)
		timesten += 3 
	}

	for i := 0; i < 10; i++ {
		x := fofn(inputs[i])
		y := gofn(inputs[i])
		difference := math.Abs(x - y)
		log.Printf("Difference: %v", difference)
	}

}