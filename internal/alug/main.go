package alug

import (
	"fmt"
	"os"

	"github.com/shirakiya/alug/internal/action"
	"github.com/urfave/cli/v2"
)

const usage string = "AWS console Login URL Generator \"alug\". Create URL attached federation token to login AWS console."

// GetConsoleLoginURL ...
func GetConsoleLoginURL(appName string, version string) error {
	app := &cli.App{
		Name:      appName,
		Version:   version,
		Usage:     usage,
		UsageText: fmt.Sprintf("%s [global options]", appName),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "profile",
				Aliases:  []string{"p"},
				Usage:    "AWS profile name assumed role",
				EnvVars:  []string{"ALUG_PROFILE"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "token-code",
				Aliases: []string{"t"},
				Usage:   "Token provided by MFA device if MFA is required",
			},
		},
		Action: action.CommandAction,
	}

	return app.Run(os.Args)
}
