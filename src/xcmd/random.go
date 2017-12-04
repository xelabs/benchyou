/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xcmd

import (
	"github.com/spf13/cobra"
)

// NewRandomCommand creates the new cmd.
func NewRandomCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "random",
		Run: randomCommandFn,
	}

	return cmd
}

func randomCommandFn(cmd *cobra.Command, args []string) {
	conf, err := parseConf(cmd)
	if err != nil {
		panic(err)
	}
	conf.Random = true
	start(conf)
}
