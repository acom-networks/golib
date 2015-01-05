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
		if r.URL.Query().Get("mesg") != "fookey" {
			t.Errorf("Expected 'mesg' == %q; got %q", "fookey", r.URL.Query().Get("mesg"))
		}
	}))
	defer ts.Close()

	msg := fmt.Sprintf("%s/?mesg=fookey", ts.URL)
	sendSms(msg)
}
