package internal

// MfaSerial ...
type MfaSerial string

// RoleArn ...
type RoleArn string

// RoleSessionName ...
type RoleSessionName string

// SourceProfile ...
type SourceProfile string

// TokenCode ...
type TokenCode string

// Config ...
type Config struct {
	DurationSeconds int
	MfaSerial       MfaSerial
	RoleArn         RoleArn
	RoleSessionName RoleSessionName
	SourceProfile   SourceProfile
	TokenCode       TokenCode
}

// RequireTokenCode ...
func (c *Config) RequireTokenCode() bool {
	return string(c.MfaSerial) != ""
}

// HasEmptyTokenCode ...
func (c *Config) HasEmptyTokenCode() bool {
	return string(c.TokenCode) == ""
}
