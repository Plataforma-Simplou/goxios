package goxios

import (
	"bytes"
	"context"
	"net/http"
	"testing"
)

func TestClientMethods(t *testing.T) {
	ts := getTestServer(t)
	defer ts.Close()

	client := NewClient(context.Background())

	requestJSON := JSON{
		"username": "gabriel",
	}
	b, err := requestJSON.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	testMethods := ts.Methods(b)

	for _, tm := range testMethods {
		var res *http.Response
		var err error
		contentType := Header{Key: "Content-Type", Value: "application/json"}
		switch tm.method {
		case http.MethodGet:
			res, err = client.Get(tm.url, []Header{})
		case http.MethodPost:
			res, err = client.Post(tm.url, []Header{contentType}, bytes.NewBuffer(tm.body))
		case http.MethodPut:
			res, err = client.Put(tm.url, []Header{contentType}, bytes.NewBuffer(tm.body))
		case http.MethodPatch:
			res, err = client.Patch(tm.url, []Header{contentType}, bytes.NewBuffer(tm.body))
		case http.MethodDelete:
			res, err = client.Delete(tm.url, []Header{contentType}, bytes.NewBuffer(tm.body))
		}
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != tm.expectedStatus {
			t.Errorf("Expected status code %d, got %d", tm.expectedStatus, res.StatusCode)
		}
	}
}
