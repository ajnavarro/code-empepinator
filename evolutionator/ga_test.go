package evolutionator

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ajnavarro/code-empepinator/sandbox"

	"github.com/stretchr/testify/assert"
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

func TestEvolutionator(t *testing.T) {
	fmt.Println("STARTING")
	var pairs []*Pair
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
		p := &Pair{
			Input:  params,
			Output: result,
		}

		pairs = append(pairs, p)
	}

	fmt.Println("PAIRS GENERATED CORRECTLY")

	err := executor.Parse(function)
	assert.NoError(t, err)

	_, err = Optimize(executor.AST, pairs, name)

	assert.NoError(t, err)
}
