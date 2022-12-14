# go-exterror

Extended error routines.

## AugmentedError

An error object which wraps another one and adds supports for extended fields.


```golang
package main

import (
	"errors"

	"github.com/randlabs/go-exterror"
)

func main() {
	err := exterror.NewAugmentedError(
		errors.New("wrapped error"),
		"some example message", map[string]interface{}{
			"value2": 1,
			"value1": "hello",
		},
	)
	
	//...
}
```
## AtomicError

An error object that can be set only once and implements context behavior.

```golang
package main

import (
	"errors"

	"github.com/randlabs/go-exterror"
)

func main() {
	err := exterror.NewAtomicError()

	// Set an error in a separate go routine
	go func() {
		err.Set(errors.New("error"))
	}()

	// Wait for the error to be set
	<-err.Done()
}
```

## License
See `LICENSE` file for details.

