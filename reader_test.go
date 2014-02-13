package ottox

import (
	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAddReader(t *testing.T) {

	runtime := otto.New()
	reader := strings.NewReader("Hello ottox!")

	assert.NoError(t, AddReader(runtime, "read", reader))

	// make sure ottox.read method was added
	assert.True(t, Exist(runtime, "read"))

	// read some bytes
	if r, err := runtime.Call(`read`, nil, 5); assert.NoError(t, err) {
		if res, err := r.Object().Get("data"); assert.NoError(t, err) {
			if bytes, err := res.ToString(); assert.NoError(t, err) {
				assert.Equal(t, bytes, "Hello", "data")
			}
		}
		if res, err := r.Object().Get("eof"); assert.NoError(t, err) {
			if eofBool, err := res.ToBoolean(); assert.NoError(t, err) {
				assert.False(t, eofBool, "eof")
			}
		}
	}

	if r, err := runtime.Call(`read`, nil, 7); assert.NoError(t, err) {
		if res, err := r.Object().Get("data"); assert.NoError(t, err) {
			if bytes, err := res.ToString(); assert.NoError(t, err) {
				assert.Equal(t, bytes, " ottox!", "data")
			}
		}
		if res, err := r.Object().Get("eof"); assert.NoError(t, err) {
			if eofBool, err := res.ToBoolean(); assert.NoError(t, err) {
				assert.False(t, eofBool, "eof")
			}
		}
	}

	if r, err := runtime.Call(`read`, nil, 1); assert.NoError(t, err) {
		if res, err := r.Object().Get("data"); assert.NoError(t, err) {
			if bytes, err := res.ToString(); assert.NoError(t, err) {
				assert.Equal(t, bytes, "", "data")
			}
		}
		if res, err := r.Object().Get("eof"); assert.NoError(t, err) {
			if eofBool, err := res.ToBoolean(); assert.NoError(t, err) {
				assert.True(t, eofBool, "eof")
			}
		}
	}

}

func TestAddReader_ReadAll(t *testing.T) {

	runtime := otto.New()
	reader := strings.NewReader("Hello ottox!")

	assert.NoError(t, AddReader(runtime, "read", reader))

	// make sure ottox.read method was added
	assert.True(t, Exist(runtime, "read"))

	// read some bytes
	if r, err := runtime.Call(`read`, nil, -1); assert.NoError(t, err) {
		if res, err := r.Object().Get("data"); assert.NoError(t, err) {
			if bytes, err := res.ToString(); assert.NoError(t, err) {
				assert.Equal(t, bytes, "Hello ottox!", "data")
			}
		}
		if res, err := r.Object().Get("eof"); assert.NoError(t, err) {
			if eofBool, err := res.ToBoolean(); assert.NoError(t, err) {
				assert.True(t, eofBool, "eof")
			}
		}
	}

}
