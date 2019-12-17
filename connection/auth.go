package connection

import (
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
)

// Auth represents the authentication method used for the connection
// and is used by the Connection to gather metadata about how to authenticate
// the connection.
type Auth interface {
	// User() should return the user for this call
	User() string
	// AuthMethod should return the ssh.AuthMethod for this specific Auth
	// e.g. ssh.Password() for simple password based Auth
	AuthMethod() ssh.AuthMethod
}

type AuthMetadata struct {
	SimplePasswordAuth *SSHSimplePasswordAuth `yaml:"simplePasswordAuth"`
}

func AuthMetadataFromYaml(data []byte) (*AuthMetadata, error) {
	authMetadata := AuthMetadata{}
	if err := yaml.Unmarshal(data, &authMetadata); err != nil {
		return nil, err
	}

	return &authMetadata, nil
}

func (a *AuthMetadata) GetAuth() Auth {
	if a.SimplePasswordAuth != nil {
		return a.SimplePasswordAuth
	}

	return nil
}

// SSHSimplePasswordAuth represents Simple SSH password authentication
type SSHSimplePasswordAuth struct {
	UserName string `yaml:"user"`
	Password string `yaml:"password"`
}

func (s *SSHSimplePasswordAuth) User() string {
	return s.UserName
}

func (s *SSHSimplePasswordAuth) AuthMethod() ssh.AuthMethod {
	return ssh.Password(s.Password)
}
