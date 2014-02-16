package ottox

import (
	"fmt"
	"sync"
)

var (
	methodGenCounter     = 0
	methodGenCounterLock sync.Mutex
)

const defaultMethodPrefix string = "ottoxMethod"

// generateMethodName generates a unique usable method name for the client
// side, with the specified optional prefix.
//
// Calling this method with no arguments will cause it to use the defaultMethodPrefix
// or one argument will specify the prefix.  Other arguments will be ignored.
func generateMethodName(optionalPrefix ...string) string {
	var prefix string
	if len(optionalPrefix) > 0 {
		prefix = optionalPrefix[0]
	}
	if len(prefix) == 0 {
		prefix = defaultMethodPrefix
	}
	methodGenCounterLock.Lock()
	methodGenCounter++
	methodGenCounterLock.Unlock()
	return fmt.Sprintf("%s_%d", prefix, methodGenCounter)
}
