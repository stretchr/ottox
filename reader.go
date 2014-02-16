package ottox

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"io"
	"io/ioutil"
)

// NewReader adds a read method to the specified runtime that allows
// client code to read from the specified reader.
//
// The client function created has the following syntax:
//
//     var response = readMethodName(bytesToRead);
//
// Response object:
//
//     {
//       data: "This was read from the reader.",
//       eof: false|true,
//       error: error|undefined
//     }
//
// If eof is false, client code should keep calling the read method
// until all data has been read.
//
// For example, to stream from the Stdin pipe:
//
//     // in go...
//     NewReader(js, "readStdIn", os.Stdin)
//
//     // in a script...
//     var data = "";
//     var response = {"eof":false};
//     while (!response.eof) {
//       response = readStdIn(255);
//       data += response.data;
//     }
//
// Passing -1 as the bytesToRead will read the entire contents from
// the reader immediately.
func NewReader(runtime *otto.Otto, methodName string, reader io.Reader) error {

	runtime.Set(methodName, func(call otto.FunctionCall) otto.Value {

		var err error
		var l int64
		var all []byte
		var val otto.Value
		var bytesRead int

		if l, err = call.Argument(0).ToInteger(); err != nil {
			raiseError(runtime, "First argument to read methods must be an integer: %s", err)
		} else {

			if l == -1 {

				all, err = ioutil.ReadAll(reader)
				if err != nil {
					raiseError(runtime, "Failed to read from io.Reader: %s", err)
				} else {

					if val, err = makeMap(runtime, map[string]interface{}{"data": string(all), "eof": true}); err != nil {
						raiseError(runtime, "Failed to create output object: %s", err)
					} else {
						return val
					}

				}

			} else {

				// read x bytes from the reader
				buf := make([]byte, l)
				bytesRead, err = reader.Read(buf)
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

				if val, err = makeMap(runtime, map[string]interface{}{"data": dataStr, "eof": isEof}); err != nil {
					raiseError(runtime, "Failed to create output object: %s", err)
				} else {
					return val
				}

			}

		}

		if err != nil {
			if val, err := makeMap(runtime, map[string]interface{}{"eof": true, "error": err.Error()}); err != nil {
				raiseError(runtime, "Failed to create output object: %s", err)
				return otto.UndefinedValue()
			} else {
				return val
			}
		}

		// nothing to return
		return otto.UndefinedValue()
	})

	return nil
}

// raiseError raises an error inside the runtime.
//
// TODO: improve this so it behaves like an error instead of just logging.
func raiseError(runtime *otto.Otto, format string, args ...interface{}) {
	var msg = fmt.Sprintf("ERROR ottox: %s", fmt.Sprintf(format, args...))
	runtime.Call("console.info", nil, msg)
}
