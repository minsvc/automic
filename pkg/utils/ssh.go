package utils

import (
	"automic/global"
	"bytes"
	"context"
	"fmt"
	"golang.org/x/crypto/ssh"
)

func SshCommand(data interface{}, ip string, port string, user string, password string) (string, error) {
	ctx := context.Background()
	client, err := ssh.Dial("tcp", ip+":"+port, &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		global.Logger.Errorf(ctx, "Ssh connection errs: %v", err.Error())
		return "ssh connection fail", err
	}
	session, err := client.NewSession()
	defer session.Close()
	if err != nil {
		global.Logger.Errorf(ctx, "SSH session errs: %v", err.Error())
	}

	cmd := fmt.Sprintf("%v", data)
	output, err := session.Output(cmd)

	if err != nil {
		global.Logger.Errorf(ctx, "SSH command errs: %v", err.Error())
		return "ssh command fail", err
	}
	result := bytes.NewBuffer(output).String()
	return result, err

}
