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

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/rslarson/cloudflare-ddns/internal/cloudflare"

	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli/v2"
)

var (
	app = &cli.App{
		Name:  "cloudflare-ddns",
		Usage: "cloudflare-ddns is a simple utility to update Cloudflare Dynamic DNS records",
		Commands: []*cli.Command{
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "Update Dynamic DNS Record",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "proxy",
						Aliases: []string{"p"},
						Usage:   "Enable CloudFlare Proxy",
						Value:   false,
					},
				},
				Action: func(c *cli.Context) error {
					ip, err := getPublicIP()
					if err != nil {
						return err
					}

					cf, err := cloudflare.NewCloudflare(token, zone, record)
					if err != nil {
						return err
					}

					err = cf.UpdateDNSRecord(ip, c.Bool("proxy"))
					if err != nil {
						return nil
					}

					rd, err := cf.GetRecordDetails()
					if err != nil {
						return err
					}
					spew.Dump(rd)
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "token",
				Aliases:     []string{"t"},
				Usage:       "Cloudflare API token",
				Required:    true,
				Destination: &token,
			},
			&cli.StringFlag{
				Name:        "zone",
				Aliases:     []string{"z"},
				Usage:       "DNS Zone",
				Required:    true,
				DefaultText: "example.com",
				Destination: &zone,
			},
			&cli.StringFlag{
				Name:        "record",
				Aliases:     []string{"r"},
				Usage:       "DNS record",
				Required:    true,
				DefaultText: "service.example.com",
				Destination: &record,
			},
		},
		Compiled: time.Time{},
		Authors: []*cli.Author{
			{
				Name:  "Randall Larson",
				Email: "rslarson147@pm.me",
			},
		},
		Action: func(c *cli.Context) error {
			cf, err := cloudflare.NewCloudflare(token, zone, record)
			if err != nil {
				return err
			}

			rd, err := cf.GetRecordDetails()
			if err != nil {
				return err
			}
			spew.Dump(rd)
			return nil
		},
	}

	record string
	token  string
	zone   string
)

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getPublicIP() (string, error) {
	c := http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://api.ipify.org?format=text", nil)
	if err != nil {
		return "", err
	}

	resp, err := c.Do(req)
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
