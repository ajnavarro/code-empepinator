package main

import (
	"fmt"
	"math/rand"

	"github.com/ajnavarro/code-empepinator/sandbox"
)

const samplesCount = 500

const name = "multiply"
const function = `function multiply(p1,p2) {
	return p1 * p2;
}
`
const paramsCount = 2

type Pair struct {
	Input  []float64
	Output float64
}

func main() {
	var pairs []*Pair
	executor := sandbox.NewJavascript(name, function)
	for i := 0; i < samplesCount; i++ {
		var params []float64
		for j := 0; j < paramsCount; j++ {
			params = append(params, rand.Float64()*float64(rand.Int()))
		}

		result, err := executor.Execute(params...)
		if err != nil {
			panic(err)
		}

		p := &Pair{
			Input:  params,
			Output: result,
		}

		pairs = append(pairs, p)
	}

	// TODO remove
	pairs = append(pairs, &Pair{[]float64{2, 2}, 4})
	for _, p := range pairs {
		fmt.Println("PAIR", p.Input, p.Output)
	}
}
