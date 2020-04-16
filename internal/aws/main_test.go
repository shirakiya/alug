package aws

import (
	"testing"
)

func TestBuildLoginURL(t *testing.T) {
	r := &SigninTokenResponse{SigninToken: "SigninToken"}
	got, err := buildLoginURL(r)

	expected := "https://signin.aws.amazon.com/federation?Action=login&Destination=https%3A%2F%2Fconsole.aws.amazon.com%2F&Issuer=IssuedByAlug&SigninToken=SigninToken"

	if got != expected {
		t.Errorf("expect %v, got %v", expected, got)
	}
	if err != nil {
		t.Errorf("expect %v, got %v", nil, err)
	}
}
