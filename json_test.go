package ottox

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToJSON(t *testing.T) {

	assert.Equal(t, `"Hello"`, toJson("Hello"))
	assert.Equal(t, `123`, toJson(123))
	assert.Equal(t, `{"age":31,"name":"Mat"}`, toJson(map[string]interface{}{"name": "Mat", "age": 31}))

}

func TestMakeMap(t *testing.T) {
	// TODO: this
}
