package exterror_test

import (
	"net"
	"testing"

	"github.com/randlabs/go-exterror"
)

// -----------------------------------------------------------------------------

func TestIsNetworkError(t *testing.T) {
	_, err := net.LookupIP("non-existent-domain.123")
	if !exterror.IsNetworkError(err) {
		t.FailNow()
	}
}
