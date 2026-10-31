package main

import (
	"io"
	"net/http"
	"strings"
)

func (b *Backend) healthCheck() bool {
	// format url & send health request
	url := strings.TrimRight(b.Address, "/") + "/health"
	resp, err := http.Get(url)

	// check for error
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	// consume the remaining response byte and discard
	// the body to make the connection reusable
	_, err = io.Copy(io.Discard, resp.Body)

	return resp.StatusCode == http.StatusOK
}
