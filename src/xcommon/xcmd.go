/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xcommon

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunCommand(cmds string, args []string) (string, error) {
	cmd := exec.Command(cmds, args...)
	if outs, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("\ncmd:[%v]\nerr:[%v]",
			strings.Join(args, ""), err)
	} else {
		return string(outs), err
	}
}
