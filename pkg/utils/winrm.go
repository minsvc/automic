package utils

import (
	"automic/global"
	"bytes"
	"context"
	"fmt"
	"github.com/masterzen/winrm"
	"os"
)

func BatCommand(data interface{}, ip string, port int, user string, password string) (string, error) {

	ctx, cancel := context.WithCancel(context.Background())
	buf := new(bytes.Buffer)
	cmd := fmt.Sprintf("%v", data)
	endpoint := winrm.NewEndpoint(ip, port, false, false, nil, nil, nil, 0)
	println("-----------------")
	client, err := winrm.NewClient(endpoint, user, password)
	if err != nil {
		global.Logger.Errorf(ctx, "winrm client err: %s", err.Error())
		return "windows connection fail", err
	}
	println("----------##########")
	defer cancel()

	println("11111111")
	println(cmd)
	result, err := client.RunWithContext(ctx, cmd, buf, os.Stderr)
	if err != nil {
		global.Logger.Errorf(ctx, "winrm exec err: %d %s", result, err.Error())
		return "windows command fail", err
	}
	print(buf.String())

	return buf.String(), err

}
