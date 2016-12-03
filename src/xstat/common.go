/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xstat

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

func splitColumns(line string) []string {
	cols := make([]string, 0)
	for _, f := range strings.Split(line, " ") {
		if len(f) > 0 {
			cols = append(cols, f)
		}
	}
	return cols
}

func sshConnect(user, password, host string, port int) (client *ssh.Client, err error) {
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
	}

	dsn := fmt.Sprintf("%s:%d", host, port)
	if client, err = ssh.Dial("tcp", dsn, sshConfig); err != nil {
		return
	}

	/*
		if session, err = client.NewSession(); err != nil {
			client.Close()
			return
		}
	*/

	return
}
