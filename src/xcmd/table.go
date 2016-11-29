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
	"iibench"
	"log"
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
	switch conf.Bench_mode {
	case "sysbench":
		workers, err := xworker.CreateWorkers(conf, 1)
		if err != nil {
			log.Panicf("prepare.error:[%+v]", err)
		}
		table := sysbench.NewTable(workers)
		table.Prepare()

	case "iibench":
		workers, err := xworker.CreateWorkers(conf, 1)
		if err != nil {
			log.Panicf("prepare.error:[%+v]", err)
		}
		table := iibench.NewTable(workers)
		table.Prepare()

	default:
		panic("unknow.bench.mode")
	}
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
	switch conf.Bench_mode {
	case "sysbench":
		workers, err := xworker.CreateWorkers(conf, 1)
		if err != nil {
			log.Panicf("cleanup.error:[%+v]", err)
		}
		table := sysbench.NewTable(workers)
		table.Cleanup()

	case "iibench":
		workers, err := xworker.CreateWorkers(conf, 1)
		if err != nil {
			log.Panicf("cleanup.error:[%+v]", err)
		}
		table := iibench.NewTable(workers)
		table.Cleanup()
	}
}
