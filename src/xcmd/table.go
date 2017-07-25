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
	"sysbench"
	"xworker"
)

func NewPrepareCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "prepare",
		Run: prepareCommandFn,
	}

	return cmd
}

func prepareCommandFn(cmd *cobra.Command, args []string) {
	conf, err := parseConf(cmd)
	if err != nil {
		panic(err)
	}

	// worker
	workers := xworker.CreateWorkers(conf, 1)
	table := sysbench.NewTable(workers)
	table.Prepare()
}

func NewCleanupCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "cleanup",
		Run: cleanupCommandFn,
	}

	return cmd
}

func cleanupCommandFn(cmd *cobra.Command, args []string) {
	conf, err := parseConf(cmd)
	if err != nil {
		panic(err)
	}

	// worker
	workers := xworker.CreateWorkers(conf, 1)
	table := sysbench.NewTable(workers)
	table.Cleanup()
}
