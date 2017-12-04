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

// NewSeqCommand creates the new cmd.
func NewSeqCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "seq",
		Run: seqCommandFn,
	}

	return cmd
}

func seqCommandFn(cmd *cobra.Command, args []string) {
	conf, err := parseConf(cmd)
	if err != nil {
		panic(err)
	}
	start(conf)
}
