package aws

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/shirakiya/alug/internal"
)

// DefaultDurationSeconds ...
const DefaultDurationSeconds = 3600

// DefaultSourceProfile ...
const DefaultSourceProfile = "default"

// FederationURL ...
const FederationURL = "https://signin.aws.amazon.com/federation"

// ConsoleURL ...
const ConsoleURL = "https://console.aws.amazon.com/"

// Issuer ...
const Issuer = "IssuedByAlug"

// SigninTokenCredentials ...
type SigninTokenCredentials struct {
	SessionID    string `json:"sessionId"`
	SessionKey   string `json:"sessionKey"`
	SessionToken string `json:"sessionToken"`
}

// SigninTokenResponse ...
type SigninTokenResponse struct {
	SigninToken string `json:"SigninToken"`
}

func createAwsSession(c *internal.Config) *session.Session {
	akid := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	// Use credentials in environments if set.
	if akid != "" && secretKey != "" {
		return session.Must(session.NewSession())
	}

	return session.Must(session.NewSessionWithOptions(session.Options{
		Profile: string(c.SourceProfile),
	}))
}

func assumeRole(sess *session.Session, c *internal.Config) (*sts.AssumeRoleOutput, error) {
	svc := sts.New(sess)

	serialNumber := string(c.MfaSerial)
	roleArn := string(c.RoleArn)
	roleSessionName := string(c.RoleSessionName)
	tokenCode := string(c.TokenCode)

	input := &sts.AssumeRoleInput{
		RoleArn:         &roleArn,
		RoleSessionName: &roleSessionName,
		SerialNumber:    &serialNumber,
		TokenCode:       &tokenCode,
	}

	return svc.AssumeRole(input)
}

func getSigninToken(output *sts.AssumeRoleOutput, c *internal.Config) (*SigninTokenResponse, error) {
	cred := SigninTokenCredentials{
		SessionID:    *output.Credentials.AccessKeyId,
		SessionKey:   *output.Credentials.SecretAccessKey,
		SessionToken: *output.Credentials.SessionToken,
	}
	credJSON, err := json.Marshal(cred)
	if err != nil {
		return new(SigninTokenResponse), err
	}

	u, _ := url.Parse(FederationURL)
	q := u.Query()
	q.Set("Action", "getSigninToken")
	q.Set("SessionDuration", strconv.Itoa(c.DurationSeconds))
	q.Set("Session", string(credJSON))
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return new(SigninTokenResponse), err
	}
	defer resp.Body.Close()
	byteBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return new(SigninTokenResponse), errors.New(string(byteBody))
	}

	signinTokenResponse := new(SigninTokenResponse)
	err = json.Unmarshal(byteBody, signinTokenResponse)
	if err != nil {
		return new(SigninTokenResponse), err
	}

	return signinTokenResponse, nil
}

func buildLoginURL(r *SigninTokenResponse) (string, error) {
	u, _ := url.Parse(FederationURL)
	q := u.Query()
	q.Set("Action", "login")
	q.Set("Issuer", Issuer)
	q.Set("Destination", ConsoleURL)
	q.Set("SigninToken", r.SigninToken)
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// CreateLoginURL ...
func CreateLoginURL(c *internal.Config) (string, error) {
	sess := createAwsSession(c)
	assumeRoleOutput, err := assumeRole(sess, c)
	if err != nil {
		return "", err
	}

	signinTokenResponse, err := getSigninToken(assumeRoleOutput, c)
	if err != nil {
		return "", err
	}

	return buildLoginURL(signinTokenResponse)
}
