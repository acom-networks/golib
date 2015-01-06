package acom

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendSms(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method %q; got %q", "GET", r.Method)
		}
		if r.Header == nil {
			t.Errorf("Expected non-nil request Header")
		}
		if r.URL.Query().Get("mesg") != "fookey 123" {
			t.Errorf("Expected 'mesg' == %q; got %q", "fookey 123", r.URL.Query().Get("mesg"))
		}
	}))
	defer ts.Close()

	uri := fmt.Sprintf("%s/?mesg=", ts.URL)
	SendSms(uri, "fookey 123")
}
