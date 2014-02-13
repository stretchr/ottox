package ottox

import (
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	"io"
	"io/ioutil"
)

func toJson(o interface{}) string {
	if bytes, err := json.Marshal(o); err == nil {
		return string(bytes)
	}
	return ""
}

// AddReader adds a read method to the specified runtime that allows
// client code to read from the specified reader.
//
// The client function created has the following syntax:
//
//     var response = methodName(bytesToRead);
//
// Response object:
//
//     {
//       data: "This was read from the reader.",
//       eof: false|true
//     }
//
// If eof is false, client code should keep calling the read method
// until all data has been read.
//
// Passing -1 as the bytesToRead will real the entire contents from
// the reader.
func AddReader(runtime *otto.Otto, methodName string, reader io.Reader) error {

	runtime.Set(methodName, func(call otto.FunctionCall) otto.Value {

		if l, err := call.Argument(0).ToInteger(); err != nil {
			raiseError(runtime, "First argument to read methods must be an integer: %s", err)
		} else {

			if l == -1 {

				all, err := ioutil.ReadAll(reader)
				if err != nil {
					raiseError(runtime, "Failed to read from io.Reader: %s", err)
				}

				if val, err := makeMap(runtime, map[string]interface{}{"data": string(all), "eof": true}); err != nil {
					raiseError(runtime, "Failed to create output object: %s", err)
				} else {
					return val
				}

			} else {

				// read x bytes from the reader
				buf := make([]byte, l)
				bytesRead, err := reader.Read(buf)
				var isEof bool = false
				if err == io.EOF {
					isEof = true
				} else if err != nil {
					raiseError(runtime, "Failed to read from io.Reader: %s", err)
				}

				// get the data
				var dataStr string
				if bytesRead > 0 {
					dataStr = string(buf)
				}

				if val, err := makeMap(runtime, map[string]interface{}{"data": dataStr, "eof": isEof}); err != nil {
					raiseError(runtime, "Failed to create output object: %s", err)
				} else {
					return val
				}

			}

		}

		// nothing to return
		return otto.UndefinedValue()
	})

	return nil
}

func raiseError(runtime *otto.Otto, format string, args ...interface{}) {
	var msg = fmt.Sprintf(format, args...)
	runtime.Call("console.info", nil, msg)
}

func makeMap(runtime *otto.Otto, m map[string]interface{}) (otto.Value, error) {
	return runtime.Run(fmt.Sprintf("eval(%s);", toJson(m)))
}
