package evolutionator

import "math"

// Based on:
// https://github.com/nickpoorman/rmse/blob/master/lib/rmse.js

func squaredError(vals []float64) []float64 {
	errors := make([]float64, len(vals))
	for i, val := range vals {
		errors[i] = math.Pow(val, 2)
	}

	return errors
}

func mean(vals []float64) float64 {
	var total float64
	for _, val := range vals {
		total += val
	}

	return total / float64(len(vals))
}

// MSE returns the Mean squared error
func MSE(vals []float64) float64 {
	return mean(squaredError(vals))
}
