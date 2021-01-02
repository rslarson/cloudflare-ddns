// Copyright (C) 2020 Randall Larson
//
// This file is part of CloudFlare DynamicDNS Updater.
//
// CloudFlare DynamicDNS Updater is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// CloudFlare DynamicDNS Updater is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with CloudFlare DynamicDNS Updater.  If not, see <http://www.gnu.org/licenses/>.

package cloudflare

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	// "reflect"
	"testing"
	// cf "github.com/cloudflare/cloudflare-go"
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

/* func TestNewCloudflare(t *testing.T) {
	type args struct {
		token  string
		zone   string
		record string
	}
	tests := []struct {
		name    string
		args    args
		want    *Cloudflare
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCloudflare(tt.args.token, tt.args.zone, tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCloudflare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCloudflare() = %v, want %v", got, tt.want)
			}
		})
	}
} */

/* func TestCloudflare_GetRecordDetails(t *testing.T) {
	type fields struct {
		API      *cf.API
		ZoneID   string
		RecordID string
		PublicIP string
	}
	tests := []struct {
		name    string
		fields  fields
		want    cf.DNSRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cloudflare{
				API:      tt.fields.API,
				ZoneID:   tt.fields.ZoneID,
				RecordID: tt.fields.RecordID,
				PublicIP: tt.fields.PublicIP,
			}
			got, err := c.GetRecordDetails()
			if (err != nil) != tt.wantErr {
				t.Errorf("Cloudflare.GetRecordDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cloudflare.GetRecordDetails() = %v, want %v", got, tt.want)
			}
		})
	}
} */

/* func TestCloudflare_UpdateDNSRecord(t *testing.T) {
	type fields struct {
		API      *cf.API
		ZoneID   string
		RecordID string
		PublicIP string
	}
	type args struct {
		proxy bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cloudflare{
				API:      tt.fields.API,
				ZoneID:   tt.fields.ZoneID,
				RecordID: tt.fields.RecordID,
				PublicIP: tt.fields.PublicIP,
			}
			if err := c.UpdateDNSRecord(tt.args.proxy); (err != nil) != tt.wantErr {
				t.Errorf("Cloudflare.UpdateDNSRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
} */

func Test_getPublicIP(t *testing.T) {
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
			got, err := getPublicIP()
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
