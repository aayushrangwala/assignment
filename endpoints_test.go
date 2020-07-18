package main

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestSimplePath(t *testing.T) {

	for i := 0; i < 10; i++ {
		randomPath := generateRandomPath()

		req, err := http.NewRequest("GET", "/"+randomPath, nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Reflect)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.

		handler.ServeHTTP(rr, req)

		t.Logf("response for %s request path is: %s", req.URL.String(), rr.Body.String())
		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK || rr.Body.String() != strings.Trim(randomPath, "/") {
			t.Errorf("endpoint failed: Status codes: [Expected: %v], [Returned: %v], path: [Expected: %v], [Returned: %v]",
				status, http.StatusOK, rr.Body.String(), randomPath)
		}
	}
}

func TestHealth(t *testing.T) {
	for i := 0; i < 10; i++ {
		req, err := http.NewRequest("GET", "/healthz", nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Health)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.

		handler.ServeHTTP(rr, req)

		t.Logf("response for %s request path is: %s", req.URL.String(), rr.Body.String())
		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("endpoint failed: Status codes: [Expected: %v], [Returned: %v]", status, http.StatusOK)
		}
	}
}

func generateRandomPath() string {
	var letterRunes = []rune("/abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
