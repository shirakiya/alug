package config

import (
	"os"
	"strconv"
	"time"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/ini.v1"
)

// AwsConfigPath ...
const AwsConfigPath = "~/.aws/config"

// AwsFileConfig ...
type AwsFileConfig struct {
	MfaSerial       string
	RoleArn         string
	RoleSessionName string
	DurationSeconds int
	SourceProfile   string
}

// CreateDefaultRoleSessionName ...
func CreateDefaultRoleSessionName() string {
	timestamp := time.Now().Unix()
	return "assume-by-alug-" + strconv.FormatInt(timestamp, 10)
}

// GetAwsFileConfig ...
func GetAwsFileConfig(path string, profile string) (AwsFileConfig, error) {
	var c AwsFileConfig

	path, err := homedir.Expand(path)
	if err != nil {
		return c, err
	}

	// It's OK that config file is not exist.
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return c, nil
	}

	cfg, err := ini.Load(path)
	if err != nil {
		return c, err
	}

	if profile == "" {
		profile = "default"
	} else {
		profile = "profile " + profile
	}

	sec, err := cfg.GetSection(profile)
	if err != nil {
		return c, nil
	}

	c = AwsFileConfig{
		MfaSerial:       sec.Key("mfa_serial").String(),
		RoleArn:         sec.Key("role_arn").String(),
		RoleSessionName: sec.Key("role_session_name").String(),
		DurationSeconds: sec.Key("duration_seconds").MustInt(),
		SourceProfile:   sec.Key("source_profile").String(),
	}

	return c, nil
}
