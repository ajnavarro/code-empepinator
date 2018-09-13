package main

import (
	"fmt"
	"math/rand"

	"github.com/ajnavarro/code-empepinator/evolutionator"
	"github.com/ajnavarro/code-empepinator/sandbox"
)

const samplesCount = 5

const name = "multiply"

var function = `
function multiply(p1,p2) {
    var a = p1;
	var b = p2;
    var result = a * b;

    return result;
}`

const paramsCount = 2

func main() {
	fmt.Println("STARTING")
	var pairs []*evolutionator.Pair
	executor := sandbox.NewJavascript(name)
	for i := 0; i < samplesCount; i++ {
		var params []float64
		for j := 0; j < paramsCount; j++ {
			params = append(params, rand.Float64()*float64(rand.Int()))
		}

		result, err := executor.Execute(function, params...)
		if err != nil {
			panic(err)
		}

		fmt.Println("RESULT", result)
		p := &evolutionator.Pair{
			Input:  params,
			Output: result,
		}

		pairs = append(pairs, p)
	}

	fmt.Println("PAIRS GENERATED CORRECTLY")

	err := executor.Parse(function)

	if err != nil {
		panic(err)
	}
	_, err = evolutionator.Optimize(executor.AST, pairs, name)

	if err != nil {
		panic(err)
	}
}
