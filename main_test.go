package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

const fixturesDir = "test-fixtures"

// readFixture is a helper function that loads a file's contents to a string.
func readFixture(t *testing.T, f string) string {
	ff, err := ioutil.ReadFile(f)
	if err != nil {
		t.Fatalf("Failed to load fixture file: %v", f)
	}
	return string(ff)
}

// TestIndexHandlerPut tests against valid and invalid PUT request cases by loding from fixed set
// of test fixtures. Each test case loads the mock request payload and expected response from
// "golden" files on disk.
func TestIndexHandlerPostRequest(t *testing.T) {
	cases := []struct {
		desc, fixture string
		isValid       bool
	}{
		{"Ideal case", "typical", true},
		{"Invalid JSON (missing comma after value)", "missing-comma", false},
		{"Invalid JSON (extra comma after last list item)", "extra-comma", false},
		{"Malformed JSON", "malformed", false},
	}

	for _, tc := range cases {
		// If we mark the test case as valid, we load the expected result from disk
		// othewise we simply load the canned "error" file.
		var inputPath, goldenPath string
		var expectedStatus int
		if tc.isValid {
			inputPath = "valid-" + tc.fixture
			goldenPath = "expected-" + tc.fixture
			expectedStatus = http.StatusOK
		} else {
			inputPath = "invalid-" + tc.fixture
			goldenPath = "expected-err"
			expectedStatus = http.StatusBadRequest
		}

		payload := readFixture(t, filepath.Join(fixturesDir, inputPath))
		req, err := http.NewRequest("POST", "/", strings.NewReader(payload))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(indexHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != expectedStatus {
			format := "Handler returned wrong status code. Test case: \"%v\"\nActual: %v, Expected: %v"
			t.Errorf(format, tc.desc, status, expectedStatus)
		}

		expectedBody := readFixture(t, filepath.Join(fixturesDir, goldenPath))
		if rr.Body.String() != string(expectedBody) {
			format := "Handler returned wrong status response body. Test case: \"%v\"\nActual: %v\nExpected: %v"
			t.Errorf(format, tc.desc, rr.Body.String(), expectedBody)
		}
	}
}

// TestIndexHandlerOtherRequests tests permutations of HTTP methods and endpoints that are not allowed
func TestIndexHandlerOtherRequests(t *testing.T) {
	cases := []struct {
		method, endpoint string
		expectedStatus   int
	}{
		{"GET", "/", http.StatusMethodNotAllowed},
		{"GET", "/foo", http.StatusMethodNotAllowed},
		{"PUT", "/", http.StatusMethodNotAllowed},
		{"PUT", "/foo", http.StatusMethodNotAllowed},
		{"DELETE", "/", http.StatusMethodNotAllowed},
		{"POST", "/foo", http.StatusForbidden},
	}
	for _, tc := range cases {
		req, err := http.NewRequest(tc.method, tc.endpoint, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(indexHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != tc.expectedStatus {
			format := "Handler returned wrong status code.\nActual: %v, Expected: %v"
			t.Errorf(format, status, tc.expectedStatus)
		}
	}
}
