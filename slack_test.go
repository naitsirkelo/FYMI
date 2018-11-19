package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func Test_sendPayload(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(IdHandler))
	defer ts.Close()
}
