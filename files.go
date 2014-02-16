package ottox

import (
	"github.com/robertkrimen/otto"
	"io"
	"os"
)

// RegisterMethodFileOpen registers a method that is capable of opening files.
//
// Response object:
//
//     {
//       ok: true|false,
//       reader: "readFile" // the method to use to read from the file
//     }
//
// To read the entire file into a variable:
//
//     // in Go...
//     RegisterMethodFileOpen(js, "openFile")
//
//     // in the script...
//     var file = openFile("filename"); // open the file
//     var read = eval(file.reader); // get the reader method
//     var close = eval(file.closer); // get the closer method
//
//     var fileContents = read(-1); // read everything
//     close(); // close the file
//
func RegisterMethodFileOpen(runtime *otto.Otto, methodName string) error {

	runtime.Set(methodName, func(call otto.FunctionCall) otto.Value {

		// check the arguments
		if len(call.ArgumentList) != 1 {
			raiseError(runtime, "%s takes 1 arguments: %s(filename)", methodName, methodName)
			return otto.UndefinedValue()
		}

		// get the arguments
		var filename string
		var err error
		if filename, err = call.Argument(0).ToString(); err != nil {
			raiseError(runtime, "%s first argument must be a string containing the filename", methodName)
			return otto.UndefinedValue()
		}

		// open the file
		var file io.ReadCloser
		if file, err = os.Open(filename); err != nil {
			raiseError(runtime, "Failed to open file '%s': %s", filename, err)
			return otto.UndefinedValue()
		}

		// add the reader
		var readerMethodName string = generateMethodName("fileRead")
		NewReader(runtime, readerMethodName, file)

		// add the closer
		var closerMethodName string = generateMethodName("fileClose")
		runtime.Set(closerMethodName, func(call otto.FunctionCall) otto.Value {
			if err := file.Close(); err != nil {
				raiseError(runtime, "Failed to close file '%s': %s", filename, err)
				return otto.FalseValue()
			}
			return otto.TrueValue()
		})

		var response otto.Value
		if response, err = makeMap(runtime, map[string]interface{}{"reader": readerMethodName, "closer": closerMethodName, "ok": true}); err != nil {
			raiseError(runtime, "%s failed to make response map.", methodName)
			return otto.UndefinedValue()
		}

		// done
		return response
	})

	return nil
}
