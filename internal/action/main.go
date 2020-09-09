package action

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/shirakiya/alug/internal"
	"github.com/shirakiya/alug/internal/aws"
	"github.com/shirakiya/alug/internal/config"
	"github.com/urfave/cli/v2"
)

func createConfig(profile string, tokenCode string) (*internal.Config, error) {
	c := new(internal.Config)

	awsFileConfig, err := config.GetAwsFileConfig(config.AwsConfigPath, profile)
	if err != nil {
		return c, err
	}

	// duration-seconds (have default value)
	durationSeconds := awsFileConfig.DurationSeconds
	if durationSeconds == 0 {
		durationSeconds = aws.DefaultDurationSeconds
	}

	// role-arn (disallow empty)
	roleArn := awsFileConfig.RoleArn
	if roleArn == "" {
		return c, errors.New("role-arn must be defined")
	}

	// role-session-name (have default value)
	roleSessionName := awsFileConfig.RoleSessionName
	if roleSessionName == "" {
		roleSessionName = config.CreateDefaultRoleSessionName()
	}

	// source-profile (exists default value)
	sourceProfile := awsFileConfig.SourceProfile
	if sourceProfile == "" {
		sourceProfile = aws.DefaultSourceProfile
	}

	c = &internal.Config{
		DurationSeconds: durationSeconds,
		MfaSerial:       internal.MfaSerial(awsFileConfig.MfaSerial),
		RoleArn:         internal.RoleArn(roleArn),
		RoleSessionName: internal.RoleSessionName(roleSessionName),
		SourceProfile:   internal.SourceProfile(sourceProfile),
		TokenCode:       internal.TokenCode(tokenCode),
	}

	return c, nil
}

func questionAndSetTokenCode(c *internal.Config) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Input MFA Code: ")
		input, _ := reader.ReadString('\n')
		// convert CRLF to LF
		input = strings.ReplaceAll(input, "\n", "")

		if input == "" || len(input) != 6 {
			fmt.Println("Invalid input, continue")
			continue
		}

		c.TokenCode = internal.TokenCode(input)
		break
	}
}

// CommandAction ...
func CommandAction(c *cli.Context) error {
	profile := c.String("profile")
	tokenCode := c.String("token-code")

	config, err := createConfig(profile, tokenCode)
	if err != nil {
		return err
	}

	if config.RequireTokenCode() && config.HasEmptyTokenCode() {
		questionAndSetTokenCode(config)
	}

	url, err := aws.CreateLoginURL(config)
	if err != nil {
		return err
	}

	// Final output.
	fmt.Printf("\n%s\n", url)

	return nil
}
