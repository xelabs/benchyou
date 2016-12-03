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
	"time"
	"xworker"
)

func NewRangeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "range",
		Run: rangeCommandFn,
	}

	return cmd
}

func rangeCommandFn(cmd *cobra.Command, args []string) {
	conf, err := parseConf(cmd)
	if err != nil {
		panic(err)
	}

	// worker
	wthds := conf.Write_threads
	rthds := conf.Read_threads
	thds := wthds + rthds
	workers, err := xworker.CreateWorkers(conf, thds)
	if err != nil {
		log.Panicf("create.workers.error:[%+v]", err)
	}

	// monitor
	monitor := NewMonitor(conf, workers)

	// insert
	var insert xworker.InsertHandler
	var query xworker.QueryHandler
	iworker := workers[:wthds]
	qworker := workers[wthds:]

	switch conf.Bench_mode {
	case "sysbench":
		insert = sysbench.NewInsert(iworker, conf.Rows_per_commit, true)
		query = sysbench.NewRange(qworker, conf.Mysql_range_order)

	case "iibench":
		insert = iibench.NewInsert(iworker, conf.Rows_per_commit, true)
		query = iibench.NewRange(qworker, conf.Mysql_range_order)
	}

	// start
	insert.Run()
	query.Run()
	monitor.Start()

	// wait
	time.Sleep(time.Duration(conf.Max_time) * time.Second)

	// stop
	insert.Stop()
	query.Stop()
	monitor.Stop()
}
