package svc

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"discovergy/internal"

	"github.com/gorilla/mux"
)

func TestSimplePath(t *testing.T) {

	for i := 0; i < 10; i++ {
		randomPath := generateRandomPath()

		req, err := http.NewRequest(http.MethodGet, "/"+randomPath, nil)
		if err != nil {
			t.Error(err)
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
		req, err := http.NewRequest(http.MethodGet, "/healthz", nil)
		if err != nil {
			t.Error(err)
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

func TestCeaserCipherEncode(t *testing.T) {
	var testCases = []struct {
		in, out string
		shift   int
	}{
		{"if-he-had-anything-confidential", "ol-nk-ngj-gteznotm-iutlojktzogr", 32},
		{"to-say", "zu-yge", 32},
		{"by-so-changing-the", "if-zv-johunpun-aol", 7},
		{"order-of-the", "xamna-xo-cqn", 9},
		{"letters of the", "exmmxkl hy max", 19},
		{"alphabet, that", "ximexybq, qexq", 23},
		{"not-a-word-could", "yze-l-hzco-nzfwo", 37},
		{"be made out.", "mp xlop zfe.", 11},
	}

	for _, tc := range testCases {
		out := internal.CeaserCipher(tc.in, tc.shift, internal.CeaserCipherEncode)
		t.Logf("Ceaser cipher, Input: [%s], Shift: [%d], Expected: [%s], out: [%s]", tc.in, tc.shift, tc.out, out)

		if out != tc.out {
			t.Errorf("[FAILED]. expected: %v, returned: %v ", tc.out, out)
		}
	}
}

func TestCeaserCipherDecode(t *testing.T) {
	var testCases = []struct {
		in, out string
		shift   int
	}{
		{"ol-nk-ngj-gteznotm-iutlojktzogr", "if-he-had-anything-confidential", 32},
		{"zu-yge", "to-say", 32},
		{"if-zv-johunpun-aol", "by-so-changing-the", 7},
		{"xamna-xo-cqn", "order-of-the", 9},
		{"exmmxkl hy max", "letters of the", 19},
		{"ximexybq, qexq", "alphabet, that", 23},
		{"yze-l-hzco-nzfwo", "not-a-word-could", 37},
		{"mp xlop zfe.", "be made out.", 11},
	}

	for _, tc := range testCases {
		out := internal.CeaserCipher(tc.in, tc.shift, internal.CeaserCipherDecode)
		t.Logf("Ceaser cipher, Input: [%s], Shift: [%d], Expected: [%s], out: [%s]", tc.in, tc.shift, tc.out, out)

		if out != tc.out {
			t.Errorf("[FAILED]. expected: %v, returned: %v ", tc.out, out)
		}
	}
}

func TestEncode(t *testing.T) {
	var testCases = []struct {
		in, out string
	}{
		{"If-he-had-anything-confidential", "Ol-nk-ngj-gteznotm-iutlojktzogr"},
		{"to-say", "zu-yge"},
		{"by-so-changing-the", "he-yu-ingtmotm-znk"},
		{"order-of-the", "uxjkx-ul-znk"},
		{"letters-of-the", "rkzzkxy-ul-znk"},
		{"alphabet,-that", "grvnghkz,-zngz"},
		{"not-a-word-could", "tuz-g-cuxj-iuarj"},
		{"be-made-out.", "hk-sgjk-uaz."},
	}

	for _, tc := range testCases {
		req, err := http.NewRequest(http.MethodGet, "/encode", nil)
		if err != nil {
			t.Error(err)
		}

		req = mux.SetURLVars(req, map[string]string{"input": tc.in})

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Encode)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		t.Logf("request, Input: [%s], Expected: [%s], out: [%s]", tc.in, tc.out, strings.Trim(rr.Body.String(), "\n"))
		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK && rr.Body.String() != tc.out {
			t.Errorf("endpoint failed: Status codes: [Expected: %v], [Returned: %v], Cipher: [Expected %v], [Returned %v]",
				status, http.StatusOK, tc.out, rr.Body.String())
		}
	}
}

func TestDecode(t *testing.T) {
	var testCases = []struct {
		in, out string
	}{
		{"Ol-nk-ngj-gteznotm-iutlojktzogr", "If-he-had-anything-confidential"},
		{"zu-yge", "to-say"},
		{"he-yu-ingtmotm-znk", "by-so-changing-the"},
		{"uxjkx-ul-znk", "order-of-the"},
		{"rkzzkxy-ul-znk", "letters-of-the"},
		{"grvnghkz,-zngz", "alphabet,-that"},
		{"tuz-g-cuxj-iuarj", "not-a-word-could"},
		{"hk-sgjk-uaz.", "be-made-out."},
	}

	for _, tc := range testCases {
		req, err := http.NewRequest(http.MethodGet, "/decode", nil)
		if err != nil {
			t.Error(err)
		}

		req = mux.SetURLVars(req, map[string]string{"input": tc.in})

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Decode)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		t.Logf("request, Input: [%s], Expected: [%s], out: [%s]", tc.in, tc.out, strings.Trim(rr.Body.String(), "\n"))
		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK && rr.Body.String() != tc.out {
			t.Errorf("endpoint failed: Status codes: [Expected: %v], [Returned: %v], Cipher: [Expected %v], [Returned %v]",
				status, http.StatusOK, tc.out, rr.Body.String())
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
