package connection

import "golang.org/x/crypto/ssh"

type Auth interface {
	// User() should return the user for this call
	User() string
	// AuthMethod should return the ssh.AuthMethod for this specific Auth
	// e.g. ssh.Password() for simple password based Auth
	AuthMethod() ssh.AuthMethod
}

// SSHSimplePasswordAuth represents Simple SSH password authentication
type SSHSimplePasswordAuth struct {
	user     string
	password string
}

func (s *SSHSimplePasswordAuth) User() string {
	return s.user
}

func (s *SSHSimplePasswordAuth) AuthMethod() ssh.AuthMethod {
	return ssh.Password(s.password)
}
