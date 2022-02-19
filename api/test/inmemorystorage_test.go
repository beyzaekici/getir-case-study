package test

import (
	"bytes"
	"getir-case/api/data"
	"getir-case/api/data/store/cache"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	wr := httptest.NewRecorder()

	holder := cache.NewCacheProvider()

	err := holder.SetKey("getir", "case")
	if err != nil {
		return
	}
	dataHandler := data.New(holder)

	req := httptest.NewRequest(http.MethodGet, "/getAndSet/get", nil)

	req.Header.Set("key", "getir")

	dataHandler.GetInMemory(wr, req)

	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), `{"key":"getir","value":"case"}`) {
		t.Errorf(
			`response body "%s" does not contain {"key":"getir","value":"case"}`,
			wr.Body.String(),
		)
	}
}

func TestSet(t *testing.T) {
	wr := httptest.NewRecorder()

	holder := cache.NewCacheProvider()

	dataHandler := data.New(holder)

	jsonStr := []byte(`{"key":"getir","value":"case"}`)

	req := httptest.NewRequest(http.MethodPost, "/getAndSet/set", bytes.NewBuffer(jsonStr))

	dataHandler.SetInMemory(wr, req)
	if wr.Code != http.StatusCreated {
		t.Errorf("got HTTP status code %d, expected 201", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), `{"key":"getir","value":"case"}`) {
		t.Errorf(
			`response body "%s" does not contain {"key":"getir","value":"case"}`,
			wr.Body.String(),
		)
	}
}
