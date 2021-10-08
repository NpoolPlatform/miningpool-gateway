package main

import (
	"fmt"
	"os"

	"github.com/NpoolPlatform/go-service-framework/pkg/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cli "github.com/urfave/cli/v2"
)

const serviceName = "Service Sample"

func main() {
	var port uint
	var configFile string

	commands := cli.Commands{
		&cli.Command{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "to run the app",
			Flags: []cli.Flag{
				&cli.UintFlag{
					Name:        "port",
					Aliases:     []string{"p"},
					Usage:       "specify this service run port",
					Destination: &port,
				},
			},
			Action: func(c *cli.Context) error {
				// Here to start your service
				return nil
			},
		},
	}

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Usage:       "specify the conofig file",
			Destination: &configFile,
		},
	}
	description := fmt.Sprintf("my %v service cli\nFor help on any individual command run <%v COMMAND -h>\n",
		serviceName, serviceName)
	_app, err := app.NewApp(serviceName, description, "", "", flags, nil, commands)
	if err != nil {
		logger.Sugar().Errorf("fail to create %v: %v", serviceName, err)
	}
	err = _app.Run(os.Args)
	if err != nil {
		logger.Sugar().Errorf("fail to run %v: %v", serviceName, err)
	}
}
