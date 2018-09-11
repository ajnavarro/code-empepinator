package sandbox

import (
	"strconv"
	"strings"

	"github.com/robertkrimen/otto"
)

type Javascript struct {
	name, function string

	vm *otto.Otto
}

func NewJavascript(name, function string) *Javascript {
	return &Javascript{
		vm:       otto.New(),
		name:     name,
		function: function,
	}
}

func (e *Javascript) Execute(vals ...float64) (float64, error) {
	_, err := e.vm.Run(e.function)
	if err != nil {
		return 0, err
	}

	var strVals []string
	for _, v := range vals {
		strVals = append(strVals, strconv.FormatFloat(v, 'f', -1, 64))
	}

	params := strings.Join(strVals, ",")

	val, err := e.vm.Run(e.name + "(" + params + ")")
	if err != nil {
		return 0, err
	}

	return val.ToFloat()
}
