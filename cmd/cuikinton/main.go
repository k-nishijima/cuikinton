package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/k-nishijima/cuikinton"
)

func main() {
	app := cli.NewApp()
	app.Name = "cuikinton"
	app.Usage = "A terminal interface for kintone"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "domain, d",
			Usage:    "your kintone domain",
			Required: true,
			EnvVar:   "KINTONE_DOMAIN",
		},
		cli.StringFlag{
			Name:   "user, u",
			Usage:  "kintone user name",
			EnvVar: "KINTONE_USERNAME",
		},
		cli.StringFlag{
			Name:   "password, p",
			Usage:  "kintone password",
			EnvVar: "KINTONE_PASSWORD",
		},
		cli.StringFlag{
			Name:   "apitoken, t",
			Usage:  "kintone api token",
			EnvVar: "KINTONE_APITOKEN",
		},
		cli.Uint64Flag{
			Name:     "appid, aid",
			Usage:    "appId to connection",
			Required: true,
		},
		cli.Uint64Flag{
			Name:  "guestspaceid, gid",
			Usage: "GuestSpaceID to connection",
			Value: 0,
		},
	}

	app.Action = func(c *cli.Context) error {
		config := cuikinton.KintoneConfig{
			Domain:       c.String("domain"),
			AppId:        c.Uint64("appid"),
			GuestSpaceId: c.Uint64("guestspaceid"),
		}
		apitoken := c.String("apitoken")
		if apitoken != "" {
			config.ApiToken = apitoken
		} else {
			config.User = c.String("user")
			config.Password = c.String("password")
		}

		cuikinton.Run(config)
		return nil
	}

	app.Run(os.Args)
}
