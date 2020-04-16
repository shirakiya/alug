package config

import (
	"testing"
	"time"

	"bou.ke/monkey"
)

func TestCreateDefaultRoleSessionName(t *testing.T) {
	mockDay := time.Date(2017, time.July, 14, 2, 40, 0, 0, time.UTC)
	patch := monkey.Patch(time.Now, func() time.Time { return mockDay })
	defer patch.Unpatch()

	expected := "assume-by-alug-1500000000"
	got := CreateDefaultRoleSessionName()

	if got != expected {
		t.Errorf("expect %v, got %v", expected, got)
	}
}

func TestGetAwsFileConfigFileNotExist(t *testing.T) {
	expected := AwsFileConfig{}

	awsFileConfig, err := GetAwsFileConfig("not_exist_path", "")

	if awsFileConfig != expected {
		t.Errorf("expect %+v, got %+v", expected, awsFileConfig)
	}
	if err != nil {
		t.Errorf("expect %v, got %v", nil, err)
	}
}

func TestGetAwsFileConfig(t *testing.T) {
	path := "./aws_config.ini"

	cases := []struct {
		profile  string
		expected AwsFileConfig
	}{
		{
			profile: "",
			expected: AwsFileConfig{
				MfaSerial:       "DEFAULT_MFA_SERIAL",
				RoleArn:         "DEFAULT_MFA_ARN",
				RoleSessionName: "DEFAULT_ROLE_SESSION_NAME",
				DurationSeconds: 1800,
				SourceProfile:   "DEFAULT_SOURCE_PROFILE",
			},
		},
		{
			profile: "test1",
			expected: AwsFileConfig{
				RoleArn:         "TEST1_MFA_ARN",
				DurationSeconds: 7200,
			},
		},
		{
			profile: "test2",
			expected: AwsFileConfig{
				RoleArn:         "TEST2_MFA_ARN",
				DurationSeconds: 0,
			},
		},
		{
			profile: "test3",
			expected: AwsFileConfig{
				RoleArn:         "",
				DurationSeconds: 600,
			},
		},
	}

	for _, c := range cases {
		awsFileConfig, err := GetAwsFileConfig(path, c.profile)

		if awsFileConfig != c.expected {
			t.Errorf("expect %+v, got %+v", c.expected, awsFileConfig)
		}
		if err != nil {
			t.Errorf("expect %v, got %v", nil, err)
		}
	}
}
