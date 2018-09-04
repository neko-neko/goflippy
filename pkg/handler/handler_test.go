package handler_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/neko-neko/goflippy/pkg/handler"
	"github.com/stretchr/testify/assert"
)

func TestServeHTTP(t *testing.T) {
	cases := []struct {
		input struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		expected int
	}{
		{
			input: struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			}{
				ID:   1000,
				Name: "Test-A",
			},
			expected: 200,
		},
		{
			input: struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			}{
				ID:   1001,
				Name: "Test-B",
			},
			expected: 201,
		},
	}

	for _, c := range cases {
		h := HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
			return c.expected, c.input, nil
		})

		eh := HTTPErrorHandlerFunc(func(code int, err error, w http.ResponseWriter) {
		})

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "http://example.com/foo", nil)

		target := Handler(h, eh)
		target.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, c.expected, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-type"))

		expectedBody, _ := json.Marshal(c.input)
		actualBody, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, string(expectedBody), strings.Trim(string(actualBody), "\n"))
	}
}

func TestServeHTTPWhenError(t *testing.T) {
	cases := []struct {
		input struct {
			Message string `json:"message"`
		}
		expected int
	}{
		{
			input: struct {
				Message string `json:"message"`
			}{
				Message: "Error-A",
			},
			expected: http.StatusInternalServerError,
		},
		{
			input: struct {
				Message string `json:"message"`
			}{
				Message: "Error-A",
			},
			expected: http.StatusInternalServerError,
		},
	}

	for _, c := range cases {
		h := HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
			return c.expected, struct{}{}, errors.New("TEST ERROR")
		})

		eh := HTTPErrorHandlerFunc(func(code int, err error, w http.ResponseWriter) {
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(c.input)
		})

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "http://example.com/foo", nil)

		target := Handler(h, eh)
		target.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, c.expected, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-type"))

		expectedBody, _ := json.Marshal(c.input)
		actualBody, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, string(expectedBody), strings.Trim(string(actualBody), "\n"))
	}
}
