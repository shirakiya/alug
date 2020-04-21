package alug

import (
	"os"

	"github.com/shirakiya/alug/internal/action"
	"github.com/shirakiya/alug/internal/config"
	"github.com/urfave/cli/v2"
)

const usage string = "AWS Login URL Generator \"alug\". Create URL attached federation token to login AWS."

// GetConsoleLoginURL ...
func GetConsoleLoginURL(appName string, version string) error {
	app := &cli.App{
		Name:    appName,
		Version: version,
		Usage:   usage,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "profile",
				Usage: "AWS profile assumed role",
			},
			&cli.IntFlag{
				Name:        "duration-seconds",
				Usage:       "Set duration time of session (unit: second)",
				DefaultText: "3600",
			},
			&cli.StringFlag{
				Name:  "mfa-serial",
				Usage: "The identification number of MFA device or ARN",
			},
			&cli.StringFlag{
				Name:  "role-arn",
				Usage: "IAM role ARN to login",
			},
			&cli.StringFlag{
				Name:        "role-session-name",
				Usage:       "Name for the assumed role session",
				Value:       config.CreateDefaultRoleSessionName(),
				DefaultText: "assume-by-alug-{unixtimestamp}",
			},
			&cli.StringFlag{
				Name:        "source-profile",
				Usage:       "Source AWS profile to perform switching role",
				DefaultText: "default",
			},
			&cli.StringFlag{
				Name:  "token-code",
				Usage: "Token provided by MFA device if MFA is required",
			},
		},
		Action: action.CommandAction,
	}

	return app.Run(os.Args)
}
