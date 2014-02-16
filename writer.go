package ottox

import (
	"github.com/robertkrimen/otto"
	"io"
)

// NewWriter adds a write method to the specified runtime that allows
// client code to write to the specified io.Writer.
//
// The client function created has the following syntax:
//
//     var response = writeMethodName(contentToWrite)
//
// Response object:
//
//     {
//       len: bytes_written,
//        error: error|undefined
//     }
func NewWriter(runtime *otto.Otto, methodName string, writer io.Writer) error {

	runtime.Set(methodName, func(call otto.FunctionCall) otto.Value {

		var data string
		var count int
		var err error
		var val otto.Value

		if data, err = call.Argument(0).ToString(); err == nil {
			if count, err = writer.Write([]byte(data)); err == nil {
				if val, err = makeMap(runtime, map[string]interface{}{"len": count}); err != nil {
					raiseError(runtime, "Failed to create output object: %s", err)
				} else {
					return val
				}
			}
		}

		if err != nil {
			if val, err := makeMap(runtime, map[string]interface{}{"len": 0, "error": err.Error()}); err != nil {
				raiseError(runtime, "Failed to create output object: %s", err)
				return otto.UndefinedValue()
			} else {
				return val
			}
		}

		return otto.UndefinedValue()

	})

	return nil
}
