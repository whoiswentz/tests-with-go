package signal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Fatalf("http.NewRequest() err = %s", err)
	}
	Handler(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Handler() status = %d; want = %d", http.StatusOK, resp.StatusCode)
	}

	ct := resp.Header.Get("Content-Type")
	if ct != "application/json" {
		t.Errorf("Handler() Content-Type = %q; want = %q", ct, "application/json")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll(resp.Body) err = %s", err)
	}

	var p Person
	if err := json.Unmarshal(data, &p); err != nil {
		t.Fatalf("json.Unmarshal(data, &p) err = %s", err)
	}
}
