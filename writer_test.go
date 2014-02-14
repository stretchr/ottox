package ottox

import (
	"bytes"
	"errors"
	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddWriter(t *testing.T) {

	runtime := otto.New()
	var buf bytes.Buffer

	assert.NoError(t, AddWriter(runtime, "write", &buf))

	// make sure ottox.read method was added
	assert.True(t, Exist(runtime, "write"))

	if r, err := runtime.Call(`write`, nil, "abcde"); assert.NoError(t, err) {
		if res, err := r.Object().Get("len"); assert.NoError(t, err) {
			if lenInt, err := res.ToInteger(); assert.NoError(t, err) {
				assert.Equal(t, 5, lenInt, "len")
				assert.Equal(t, "abcde", buf.String())
			}
		}
	}

	if r, err := runtime.Call(`write`, nil, "fghij"); assert.NoError(t, err) {
		if res, err := r.Object().Get("len"); assert.NoError(t, err) {
			if lenInt, err := res.ToInteger(); assert.NoError(t, err) {
				assert.Equal(t, 5, lenInt, "len")
				assert.Equal(t, "abcdefghij", buf.String())
			}
		}
	}

}

type writerThatThrowsError struct{}

func (w writerThatThrowsError) Write(data []byte) (int, error) {
	return 0, errors.New("Something went wrong")
}

func TestAddWriter_Error(t *testing.T) {

	runtime := otto.New()

	assert.NoError(t, AddWriter(runtime, "write", &writerThatThrowsError{}))

	// make sure ottox.read method was added
	assert.True(t, Exist(runtime, "write"))

	if r, err := runtime.Call(`write`, nil, "abcde"); assert.NoError(t, err) {
		if res, err := r.Object().Get("len"); assert.NoError(t, err) {
			if lenInt, err := res.ToInteger(); assert.NoError(t, err) {
				assert.Equal(t, 0, lenInt, "len")
				if res, err := r.Object().Get("error"); assert.NoError(t, err) {
					if errString, err := res.ToString(); assert.NoError(t, err) {
						assert.Equal(t, "Something went wrong", errString)
					}
				}
			}
		}
	}

}
