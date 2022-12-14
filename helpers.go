package exterror

import (
	"net"
)

// -----------------------------------------------------------------------------

// IsNetworkError returns true if the provided error object is related to a network error.
func IsNetworkError(err error) bool {
	if err != nil {
		switch err.(type) {
		case net.Error:
			return true
		case *net.OpError:
			return true
		case *net.DNSError:
			return true
		}
	}
	return false
}
