package internal

import (
	"testing"
)

func TestConfig_RequireTokenCode(t *testing.T) {
	cases := []struct {
		mfaSerial string
		expected  bool
	}{
		{
			mfaSerial: "",
			expected:  false,
		},
		{
			mfaSerial: "MFA_SERIAL",
			expected:  true,
		},
	}

	for _, c := range cases {
		config := Config{MfaSerial: MfaSerial(c.mfaSerial)}

		if config.RequireTokenCode() != c.expected {
			t.Errorf("expect %v, got %v", c.expected, config.RequireTokenCode())
		}
	}
}

func TestConfig_HasEmptyTokenCode(t *testing.T) {
	cases := []struct {
		tokenCode string
		expected  bool
	}{
		{
			tokenCode: "",
			expected:  true,
		},
		{
			tokenCode: "123456",
			expected:  false,
		},
	}

	for _, c := range cases {
		config := Config{TokenCode: TokenCode(c.tokenCode)}

		if config.HasEmptyTokenCode() != c.expected {
			t.Errorf("expect %v, got %v", c.expected, config.HasEmptyTokenCode())
		}
	}
}
