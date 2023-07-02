package main

import (
	"log"
	"os"

	"github.com/cxfksword/httpcap/capture"
	"github.com/cxfksword/httpcap/flags"
	"github.com/cxfksword/httpcap/output"
	"github.com/cxfksword/httpcap/utils"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

var (
	AppName = "httpcap"
	AppDesc = "A simple network analyzer that capture http network traffic."

	/*********Will auto update by ci build *********/
	Version = "unknown"
	/*********Will auto update by ci build *********/
)

func main() {
	app := &cli.App{
		Name:     AppName,
		HelpName: AppDesc,
		Version:  Version,
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "show all interfaces",
				Action: func(c *cli.Context) error {
					utils.ShowAllInterfaces()
					return nil
				},
			},
			{
				Name:  "run",
				Usage: "run capture",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "interface",
						Aliases:     []string{"i"},
						Usage:       "interface to listen on (e.g. eth0, en1, or 192.168.1.1, 127.0.0.1 etc.)",
						Value:       "any",
						Destination: &flags.Options.InterfaceName,
					},
					&cli.IntFlag{
						Name:        "port",
						Aliases:     []string{"p"},
						Usage:       "port to listen on (default listen on all port)",
						Destination: &flags.Options.Port,
					},
					&cli.StringFlag{
						Name:        "ip",
						Usage:       "capture traffic to and from ip",
						Destination: &flags.Options.Ip,
					},
					&cli.StringFlag{
						Name:        "keyword",
						Aliases:     []string{"k"},
						Usage:       "filte output match the keyword",
						Destination: &flags.Options.Keyword,
					},
					&cli.BoolFlag{
						Name:        "body",
						Usage:       "print body content (only support text content body)",
						Destination: &flags.Options.Body,
					},
					&cli.BoolFlag{
						Name:        "raw",
						Usage:       "print raw request / response",
						Destination: &flags.Options.Raw,
					},
					&cli.BoolFlag{
						Name:        "long",
						Aliases:     []string{"l"},
						Usage:       "print more detail (ex. source ip)",
						Destination: &flags.Options.MoreDetail,
					},
					&cli.BoolFlag{
						Name:        "debug",
						Aliases:     []string{"vv"},
						Usage:       "print debug message",
						Destination: &flags.Options.Debug,
					},
					&cli.BoolFlag{
						Name:        "verbose",
						Aliases:     []string{"vvv"},
						Usage:       "print more debug message",
						Destination: &flags.Options.Verbose,
					},
				},
				Action: func(c *cli.Context) error {
					zerolog.SetGlobalLevel(zerolog.InfoLevel)
					zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr})
					if flags.Options.Debug {
						zerolog.SetGlobalLevel(zerolog.DebugLevel)
					}
					if flags.Options.Verbose {
						zerolog.SetGlobalLevel(zerolog.TraceLevel)
					}

					output.Init()
					capture.Start()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
