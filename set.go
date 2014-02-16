package ottox

import (
	"github.com/robertkrimen/otto"
)

// SetOnce sets a name or value only if it doesn't exist.
//
// Returns whether the value was set or not, or an error if it failed to
// set it.
//
// For acceptable values, see http://godoc.org/github.com/robertkrimen/otto#Otto.Set
func SetOnce(runtime *otto.Otto, name string, value interface{}) (bool, error) {

	if !Exist(runtime, name) {

		if err := runtime.Set(name, value); err != nil {
			return false, err
		}

		return true, nil

	}

	// nothing to do
	return false, nil

}
