package action

import (
	"errors"
	"testing"

	"bou.ke/monkey"
	"github.com/shirakiya/alug/internal"
	"github.com/shirakiya/alug/internal/aws"
	"github.com/shirakiya/alug/internal/config"
)

func TestCreateConfigAsGetAwsFileConfigThrowsError(t *testing.T) {
	expected := internal.Config{}
	expectedErr := errors.New("")

	mockAwsFileConfig := config.AwsFileConfig{}
	patch := monkey.Patch(config.GetAwsFileConfig, func(path string, profile string) (config.AwsFileConfig, error) {
		return mockAwsFileConfig, expectedErr
	})
	defer patch.Unpatch()

	got, err := createConfig("", "")

	if *got != expected {
		t.Errorf("expect %+v, got %+v", expected, *got)
	}
	if err != expectedErr {
		t.Errorf("expect %v, got %v", expectedErr, err)
	}
}

func TestCreateConfigSetFromAwsFileConfig(t *testing.T) {
	mockAwsFileConfig := config.AwsFileConfig{
		MfaSerial:       "FromAwsFileConfig",
		RoleArn:         "FromAwsFileConfig",
		RoleSessionName: "FromAwsFileConfig",
		DurationSeconds: 100,
		SourceProfile:   "FromAwsFileConfig",
	}

	patch := monkey.Patch(config.GetAwsFileConfig, func(path string, profile string) (config.AwsFileConfig, error) {
		return mockAwsFileConfig, nil
	})
	defer patch.Unpatch()

	got, err := createConfig("", "")

	expected := internal.Config{
		MfaSerial:       internal.MfaSerial("FromAwsFileConfig"),
		RoleArn:         internal.RoleArn("FromAwsFileConfig"),
		RoleSessionName: internal.RoleSessionName("FromAwsFileConfig"),
		DurationSeconds: 100,
		SourceProfile:   internal.SourceProfile("FromAwsFileConfig"),
	}

	if *got != expected {
		t.Errorf("expect %+v, got %+v", expected, *got)
	}
	if err != nil {
		t.Errorf("expect %v, got %v", nil, err)
	}
}

func TestCreateConfigSetDefaultValue(t *testing.T) {
	mockAwsFileConfig := config.AwsFileConfig{
		MfaSerial:       "FromAwsFileConfig",
		RoleArn:         "FromAwsFileConfig",
		RoleSessionName: "",
		DurationSeconds: 0,
		SourceProfile:   "",
	}

	patch := monkey.Patch(config.GetAwsFileConfig, func(path string, profile string) (config.AwsFileConfig, error) {
		return mockAwsFileConfig, nil
	})
	defer patch.Unpatch()

	got, err := createConfig("", "")

	if got.DurationSeconds != aws.DefaultDurationSeconds {
		t.Errorf("expect %v, got %v", aws.DefaultDurationSeconds, got.DurationSeconds)
	}
	expectedRoleSessionName := config.CreateDefaultRoleSessionName()
	if string(got.RoleSessionName) != expectedRoleSessionName {
		t.Errorf("expect %v, got %v", got.RoleSessionName, expectedRoleSessionName)
	}
	if got.SourceProfile != aws.DefaultSourceProfile {
		t.Errorf("expect %v, got %v", aws.DefaultSourceProfile, got.SourceProfile)
	}
	if err != nil {
		t.Errorf("expect %v, got %v", nil, err)
	}
}

func TestCreateConfigAllowsEmptyValue(t *testing.T) {
	mockAwsFileConfig := config.AwsFileConfig{
		MfaSerial:       "",
		RoleArn:         "FromAwsFileConfig",
		RoleSessionName: "FromAwsFileConfig",
		DurationSeconds: 0,
		SourceProfile:   "FromAwsFileConfig",
	}

	patch := monkey.Patch(config.GetAwsFileConfig, func(path string, profile string) (config.AwsFileConfig, error) {
		return mockAwsFileConfig, nil
	})
	defer patch.Unpatch()

	got, err := createConfig("", "")

	if got.MfaSerial != "" {
		t.Errorf("expect %v, got %v", "", got.MfaSerial)
	}
	if err != nil {
		t.Errorf("expect %v, got %v", nil, err)
	}
}

func TestCreateConfigDisallowsEmptyRoleName(t *testing.T) {
	mockAwsFileConfig := config.AwsFileConfig{
		MfaSerial:       "FromAwsFileConfig",
		RoleArn:         "",
		RoleSessionName: "FromAwsFileConfig",
		DurationSeconds: 0,
		SourceProfile:   "FromAwsFileConfig",
	}

	patch := monkey.Patch(config.GetAwsFileConfig, func(path string, profile string) (config.AwsFileConfig, error) {
		return mockAwsFileConfig, nil
	})
	defer patch.Unpatch()

	_, err := createConfig("", "")

	if err == nil {
		t.Errorf("expect %v, got %v", "role-arn must be defined", err)
	}
}
