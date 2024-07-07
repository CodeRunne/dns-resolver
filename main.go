package main

import (
	"os"
	"log"
	"net"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command {
			{
				Name: "nameserver",
				Aliases: []string{"ns"},
				Usage: "Check nameserver of any host",
				Action: func(c *cli.Context) error {
					if c.Args().Len() >= 2 {
						color.Red.Println("Can only lookup one domain at a time")
						return nil
					}

					// get host
					host := c.Args().First()

					// lookup host
					ns, err := net.LookupNS(host)
					if err != nil {
						color.Red.Println(err.Error())
						return nil
					}

					color.Style{color.FgCyan, color.OpBold}.Println("Nameservers for", host, ":")
					for i := 0; i < len(ns); i++ {
						color.Green.Println(ns[i].Host)
					}

					return nil
				},
			},
			{
				Name: "mailserver",
				Aliases: []string{"mx"},
				Usage: "Get mail servers for a domain",
				Action: func(c *cli.Context) error {
					if c.Args().Len() >= 2 {
						color.Red.Println("Can only lookup one domain mx at a time")
						return nil
					}

					// host
					host := c.Args().First()

					// lookup mail server
					mx, err := net.LookupMX(host)
					if err != nil {
						color.Red.Println(err.Error())
						return nil
					}

					color.Style{color.FgCyan, color.OpBold}.Println("Mailservers for", host, ":")
					for i := 0; i < len(mx); i++ {
						color.Green.Println("Mail Server Host:", mx[i].Host)
					}

					return nil
				},
			},
			{
				Name: "ip",
				Usage: "Lookup IP address of a host",
				Action: func(c *cli.Context) error {
					if c.Args().Len() >= 2 {
						color.Red.Println("Can only lookup one domain ip at a time")
						return nil
					}

					// get host
					host := c.Args().First()

					ip, err := net.LookupIP(host)
					if err != nil {
						color.Red.Println(err.Error())
						return nil
					}

					color.Style{color.FgCyan, color.OpBold}.Println("IP Addresses for", host, ":")
					for i := 0; i < len(ip); i++ {
						color.Green.Println("IP Address for:", ip[i].To16())
					}

					return nil
				},
			},
			{
				Name: "cname",
				Usage: "Lookup cname of a host",
				Action: func(c *cli.Context) error {
					if c.Args().Len() >= 2 {
						color.Red.Println("Can only lookup one domain ip at a time")
						return nil
					}

					// get host
					host := c.Args().First()

					cname, err := net.LookupCNAME(host)
					if err != nil {
						color.Red.Println(err.Error())
						return nil
					}

					color.Style{color.FgCyan, color.OpBold}.Println("CNAME Record for", host, ":")
					color.Green.Println("CNAME Record for:", cname)
					return nil
				},
			},
			{
				Name: "txt",
				Usage: "Lookup txt of a host",
				Action: func(c *cli.Context) error {
					if c.Args().Len() >= 2 {
						color.Red.Println("Can only lookup one domain ip at a time")
						return nil
					}

					// get host
					host := c.Args().First()

					txt, err := net.LookupTXT(host)
					if err != nil {
						color.Red.Println(err.Error())
						return nil
					}

					color.Style{color.FgCyan, color.OpBold}.Println("TXT Record for", host, ":")
					for i := 0; i < len(txt); i++ {
						color.Green.Println(txt[i])
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}