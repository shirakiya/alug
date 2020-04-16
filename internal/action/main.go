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

func createConfig(
	profile string,
	durationSeconds int,
	mfaSerial string,
	roleArn string,
	roleSessionName string,
	sourceProfile string,
	tokenCode string,
) (*internal.Config, error) {
	c := new(internal.Config)

	awsFileConfig, err := config.GetAwsFileConfig(config.AwsConfigPath, profile)
	if err != nil {
		return c, err
	}

	// duration-seconds (exists default value)
	if durationSeconds == 0 {
		durationSeconds = awsFileConfig.DurationSeconds
		if durationSeconds == 0 {
			durationSeconds = aws.DefaultDurationSeconds
		}
	}

	// mfa-serial
	if mfaSerial == "" {
		mfaSerial = awsFileConfig.MfaSerial
	}

	// role-arn (disallow empty)
	if roleArn == "" {
		roleArn = awsFileConfig.RoleArn
		if roleArn == "" {
			return c, errors.New("role-arn must be defined")
		}
	}

	// role-session-name (disallow empty)
	if awsFileConfig.RoleSessionName != "" {
		roleSessionName = awsFileConfig.RoleSessionName
	}
	if roleSessionName == "" {
		return c, errors.New("Empty role-session-name is not allowed")
	}

	// source-profile (exists default value)
	if sourceProfile == "" {
		sourceProfile = awsFileConfig.SourceProfile
		if sourceProfile == "" {
			sourceProfile = aws.DefaultSourceProfile
		}
	}

	c = &internal.Config{
		DurationSeconds: durationSeconds,
		MfaSerial:       internal.MfaSerial(mfaSerial),
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
		input = strings.Replace(input, "\n", "", -1)

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
	durationSeconds := c.Int("duration-seconds")
	mfaSerial := c.String("mfa-serial")
	roleArn := c.String("role-arn")
	roleSessionName := c.String("role-session-name")
	sourceProfile := c.String("source-profile")
	tokenCode := c.String("token-code")

	config, err := createConfig(
		profile,
		durationSeconds,
		mfaSerial,
		roleArn,
		roleSessionName,
		sourceProfile,
		tokenCode,
	)
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
	fmt.Println(url)

	return nil
}
