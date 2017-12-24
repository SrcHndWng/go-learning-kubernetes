package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	w := httptest.NewRecorder()
	home(w, nil)

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", have, want)
	}

	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	var info versionInfo
	err = json.Unmarshal(result, &info)
	if err != nil {
		t.Fatal(err)
	}

	if info.BuildTime == "" || info.Commit == "" || info.Release == "" {
		t.Errorf("Response Body is wrong. result = %v\n", info)
	}
}
