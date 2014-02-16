package ottox

import (
	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterMethodFileOpen(t *testing.T) {

	runtime := otto.New()

	// reset the counter to predict the auto generated method name
	methodGenCounterLock.Lock()
	methodGenCounter = 0
	methodGenCounterLock.Unlock()

	if assert.NoError(t, RegisterMethodFileOpen(runtime, "openFile")) {

		// call the method and make assertions about the response
		if val, err := runtime.Run(`openFile("test/test-files/text.txt")`); assert.NoError(t, err) {
			if assert.NotEqual(t, otto.UndefinedValue(), val, "val == undefined") {

				if assert.True(t, val.IsObject()) {
					if okObj, err := val.Object().Get("ok"); assert.NoError(t, err) {
						if okBool, err := okObj.ToBoolean(); assert.NoError(t, err) {
							assert.True(t, okBool)
						}
					}
					if readerMethodNameObj, err := val.Object().Get("reader"); assert.NoError(t, err) {
						if readerMethodName, err := readerMethodNameObj.ToString(); assert.NoError(t, err) {
							assert.Equal(t, "fileRead_1", readerMethodName)
						}
					}
					if closerMethodNameObj, err := val.Object().Get("closer"); assert.NoError(t, err) {
						if closerMethodName, err := closerMethodNameObj.ToString(); assert.NoError(t, err) {
							assert.Equal(t, "fileClose_2", closerMethodName)
						}
					}
				}

			}
		}

	}

}
