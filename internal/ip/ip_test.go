package ip

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	invalidIP = "a68c:fe4f:867b:79d7:fc17:dc68:f43f:cc10"
	validIP   = "8.8.8.8"
)

type MockDoType func(req *http.Request) (*http.Response, error)

// MockClient is a mock HTTPClient
type MockClient struct {
	MockDo MockDoType
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func TestGetPublicIP(t *testing.T) {
	// Test cases
	tests := []struct {
		name       string
		want       string
		wantErr    bool
		genError   bool
		httpStatus bool
		noIP       bool
		invalidIP  bool
	}{
		{
			name:    "Test correct HTTP GET to remote service.",
			want:    validIP,
			wantErr: false,
		},
		{
			name:     "Test error returned when HTTP error.",
			want:     "",
			wantErr:  true,
			genError: true,
		},
		{
			name:       "Test error returned when HTTP status code not OK.",
			want:       "",
			wantErr:    true,
			httpStatus: true,
		},
		{
			name:    "Test error returned when no IP address",
			want:    "",
			wantErr: true,
			noIP:    true,
		},
		{
			name:      "Test error returned when invalid IP address",
			want:      "",
			wantErr:   true,
			invalidIP: true,
		},
	}

	for _, tt := range tests {
		// Setup Mock HTTP Client
		client = &MockClient{
			MockDo: func(req *http.Request) (*http.Response, error) {
				if tt.wantErr {
					if tt.genError {
						return nil, fmt.Errorf("generic error")
					} else if tt.httpStatus {
						return &http.Response{
							StatusCode: http.StatusBadRequest,
						}, nil
					} else if tt.noIP {
						r := ioutil.NopCloser(bytes.NewReader([]byte("")))
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       r,
						}, nil
					} else if tt.invalidIP {
						r := ioutil.NopCloser(bytes.NewReader([]byte(invalidIP)))
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       r,
						}, nil
					}
				}
				r := ioutil.NopCloser(bytes.NewReader([]byte(validIP)))
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}, nil
			},
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPublicIP()
			if (err != nil) != tt.wantErr {
				t.Errorf("getPublicIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getPublicIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
