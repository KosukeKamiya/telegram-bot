package main

import (
	"io"
	"net/http"
	"testing"

	"google.golang.org/appengine/aetest"
)



func createTestRequest(method string, path string, r io.Reader, t *testing.T) *http.Request {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest(method, path, r)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}
	return req
}
