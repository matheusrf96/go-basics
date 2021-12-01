package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the command line application ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search for IPs and server names on internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search IP addresses on internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "en.wikipedia.org",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
