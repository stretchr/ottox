package ottox

import (
	"fmt"
	"github.com/robertkrimen/otto"
)

// raiseError raises an error inside the runtime.
//
// TODO: improve this so it behaves like an error instead of just logging.
func raiseError(runtime *otto.Otto, format string, args ...interface{}) {
	var msg = fmt.Sprintf("ERROR ottox: %s", fmt.Sprintf(format, args...))
	runtime.Call("console.info", nil, msg)
}
