package exercises

import (
	"math/rand"
)

func Pi(N int) float64 {
	inside := 0
	result := 0.0

	for i := 0; i < N; i++ {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1

		if x*x+y*y <= 1 {
			inside++
		}

		result = 4 * float64(inside) / float64(N)
	}

	return result
}
