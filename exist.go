package ottox

import (
	"github.com/robertkrimen/otto"
)

// Exist gets whether a variable exists in the runtime or not.
func Exist(runtime *otto.Otto, name string) bool {

	if val, err := runtime.Get(name); err == nil {
		return val != otto.UndefinedValue()
	}

	return false
}
