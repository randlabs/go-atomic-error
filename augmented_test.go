package exterror_test

import (
	"errors"
	"testing"

	"github.com/randlabs/go-exterror"
)

// -----------------------------------------------------------------------------

func TestAugmentedError(t *testing.T) {
	err := exterror.NewAugmentedError(
		errors.New("dummy wrapped error"),
		"dummy message error", map[string]interface{}{
			"value2": 1000,
			"value1": "hello",
		},
	)

	if err.Error() != "dummy message error [value1=hello] [value2=1000] [err=dummy wrapped error]" {
		t.FailNow()
	}
}
