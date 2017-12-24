package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHome(t *testing.T) {
	w := httptest.NewRecorder()
	buildTime := time.Now().Format("20060102_03:04:05")
	commit := "some test hash"
	release := "0.0.8"
	h := home(buildTime, commit, release)
	h(w, nil)

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

	if info.BuildTime != buildTime || info.Commit != commit || info.Release != release {
		t.Errorf("Response Body is wrong. result = %v\n", info)
	}
}
