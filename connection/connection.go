package connection

import "golang.org/x/crypto/ssh"

// CloseFunc represents a close method of a connection 
type CloseFunc func() error

type SSHConnection struct {
	auth Auth
}

// GetConnection returns a ssh.Client and a CloseFunc. The CloseFunc should
// be called and defered to prevent leakage of the connection resources  
func (c *SSHConnection) GetConnection() (*ssh.Client, CloseFunc, error) {
	config := &ssh.ClientConfig{
		User: c.auth.User(),
		Auth: []ssh.AuthMethod{
			c.auth.AuthMethod(),
		},
		//TODO pass option to 
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", "host", config)
	if err != nil {
		return nil, nil, err
	}

	return conn, conn.Close, nil
}
