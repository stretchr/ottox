package ottox

import (
	"github.com/robertkrimen/otto"
	"io"
)

func AddWriter(runtime *otto.Otto, methodName string, writer io.Writer) error {

	runtime.Set(methodName, func(call otto.FunctionCall) otto.Value {

		if data, err := call.Argument(0).ToString(); err == nil {
			if count, err := writer.Write([]byte(data)); err == nil {
				if val, err := makeMap(runtime, map[string]interface{}{"len": count}); err != nil {
					raiseError(runtime, "Failed to create output object: %s", err)
				} else {
					return val
				}
			}
		}

		return otto.UndefinedValue()
	})

	return nil
}
