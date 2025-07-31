package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleDeeplink(t *testing.T) {
	tests := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedHeader string
	}{
		{
			name:           "Valid deeplink",
			queryParam:     "deeplink=myapp://home",
			expectedStatus: http.StatusFound,
			expectedHeader: "myapp://home",
		},
		{
			name:           "Missing deeplink parameter",
			queryParam:     "",
			expectedStatus: http.StatusBadRequest,
			expectedHeader: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/redirect?"+tt.queryParam, nil)
			w := httptest.NewRecorder()

			handleDeeplink(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d",
					tt.expectedStatus, resp.StatusCode)
			}

			if tt.expectedHeader != "" {
				location := resp.Header.Get("Location")
				if location != tt.expectedHeader {
					t.Errorf("expected Location header %s, got %s",
						tt.expectedHeader, location)
				}
			}
		})
	}
}
