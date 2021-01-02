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
	"log"
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
					cf, err := cloudflare.NewCloudflare(token, zone, record)
					if err != nil {
						return err
					}

					err = cf.UpdateDNSRecord(c.Bool("proxy"))
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
