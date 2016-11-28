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
	"log"
	"sysbench"
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
	workers, err := sysbench.CreateWorkers(conf, 1)
	if err != nil {
		log.Panicf("prepare.error:[%+v]", err)
	}

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
	workers, err := sysbench.CreateWorkers(conf, 1)
	if err != nil {
		log.Panicf("cleanup.error:[%+v]", err)
	}

	table := sysbench.NewTable(workers)
	table.Cleanup()
}
