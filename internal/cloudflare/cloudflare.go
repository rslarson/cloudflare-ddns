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
	"fmt"

	"github.com/rslarson/cloudflare-ddns/internal/ip"

	cf "github.com/cloudflare/cloudflare-go"
)

// Cloudflare represents the Cloudflare API and zoneID
type Cloudflare struct {
	API      *cf.API
	ZoneID   string
	RecordID string
	PublicIP string
}

// NewCloudflare retuns a Cloudflare struct
func NewCloudflare(token, zone, record string) (*Cloudflare, error) {
	api, err := cf.NewWithAPIToken(token)
	if err != nil {
		return nil, err
	}

	zid, err := api.ZoneIDByName(zone)
	if err != nil {
		return nil, err
	}

	req := cf.DNSRecord{
		Type:   "A",
		Name:   record,
		ZoneID: zid,
	}

	resp, err := api.DNSRecords(zid, req)
	if err != nil {
		return nil, err
	} else if len(resp) > 1 {
		return nil, fmt.Errorf("Multiple records for provided DNS record: \n%v", resp)
	} else if resp == nil {
		return nil, fmt.Errorf("No records returned for provided DNS record")
	}

	ip, err := ip.GetPublicIP()
	if err != nil {
		return nil, err
	}

	return &Cloudflare{
		API:      api,
		ZoneID:   zid,
		RecordID: resp[0].ID,
		PublicIP: ip,
	}, nil
}

// GetRecordDetails returns the Cloudflare DNS zone details
func (c *Cloudflare) GetRecordDetails() (cf.DNSRecord, error) {
	details, err := c.API.DNSRecord(c.ZoneID, c.RecordID)
	if err != nil {
		return cf.DNSRecord{}, err
	}
	return details, nil
}

// UpdateDNSRecord updates the provided DNS record with the provided IP address
func (c *Cloudflare) UpdateDNSRecord(proxy bool) error {
	newRecord := cf.DNSRecord{
		ID:        c.RecordID,
		Type:      "A",
		Content:   c.PublicIP,
		Proxiable: true,
		Proxied:   proxy,
		ZoneID:    c.ZoneID,
	}

	err := c.API.UpdateDNSRecord(c.ZoneID, c.RecordID, newRecord)
	if err != nil {
		return err
	}
	return nil
}
