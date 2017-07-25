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
	"xcommon"
)

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
	benchConf := &xcommon.BenchConf{
		Random:           true,
		XA:               conf.XA,
		Rows_per_insert:  conf.Rows_per_insert,
		Batch_per_commit: conf.Batch_per_commit,
	}
	start(conf, benchConf)
}
