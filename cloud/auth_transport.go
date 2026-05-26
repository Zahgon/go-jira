package cloud

import "net/http"

// cloneRequest returns a clone of the provided *http.Request.
// The clone is a shallow copy of the struct and its Header map.
func cloneRequest(r *http.Request) *http.Request {
	_ = "STUB: not implemented"
	// shallow copy of the struct
	return nil
}

// deep copy of the Header
