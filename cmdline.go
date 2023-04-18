package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func ConfigCmd(defaultAction func(*cli.Context) error) *cli.App {
    return &cli.App{
        Action: defaultAction, // only run if no sub command

        // run if sub command not correct
		CommandNotFound: func(context *cli.Context, s string) {
			fmt.Println("command not find, use -h or --help show help")
		},

        Commands: []*cli.Command{
            {
				Name:  "config",
				Usage: "config command",
				Subcommands: []*cli.Command{
					{
						Name:  "get_token",
						Usage: "get token",
						Action: func(clictx *cli.Context) error {
                            appConfig := loadConfig(BINARY_DIR)
                            fmt.Println(appConfig.cfg.Token)
                            return nil
						},
					},
					{
						Name:  "set",
						Usage: "set token",
						Flags: []cli.Flag{
                            &cli.StringFlag{Name: "token", Required: false},
                        },
						Action: func(clictx *cli.Context) error {
                            // Change value in map and marshal back into yaml
                            tokenSet := clictx.String("token")

                            appConfig := loadConfig(BINARY_DIR)
                            appConfig.cfg.Token = tokenSet
                            appConfig.updateConfig()
                            return nil
						},
					},
				},
			},
        },

    }
}
