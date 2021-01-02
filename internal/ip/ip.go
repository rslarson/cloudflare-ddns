// Copyright (C) 2021 Randall Larson
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

package ip

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var client HTTPClient

func init() {
	client = &http.Client{
		Timeout: 5 * time.Second,
	}
}

// GetPublicIP returns the public (WAN) IP address
func GetPublicIP() (string, error) {
	req, err := http.NewRequest("GET", "https://api.ipify.org?format=text", nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	} else if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP Status Code not OK: %v", resp.Status)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	ip := string(b)

	if i := net.ParseIP(ip); i == nil {
		return "", fmt.Errorf("Invalid IP recieved: %v", i)
	} else if j := i.To4(); j == nil {
		return "", fmt.Errorf("Expected IPv4 Address, recieved IPv6: %v", i)
	}

	return ip, nil
}
