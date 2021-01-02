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

// "reflect"

// cf "github.com/cloudflare/cloudflare-go"

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
