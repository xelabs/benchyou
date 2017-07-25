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
	"time"
	"xcommon"
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
	workers := xworker.CreateWorkers(conf, thds)

	// monitor
	monitor := NewMonitor(conf, workers)

	// insert
	var insert xworker.InsertHandler
	var query xworker.QueryHandler
	iworker := workers[:wthds]
	qworker := workers[wthds:]
	benchConf := &xcommon.BenchConf{
		Random:          true,
		XA:              conf.XA,
		Rows_per_insert: conf.Rows_per_insert,
	}
	insert = sysbench.NewInsert(benchConf, iworker)
	query = sysbench.NewRange(qworker, conf.Mysql_range_order)

	// start
	insert.Run()
	query.Run()
	monitor.Start()

	done := make(chan bool)
	go func(ins xworker.InsertHandler, q xworker.QueryHandler, max uint64) {
		if max == 0 {
			return
		}

		for {
			time.Sleep(time.Second)
			all := ins.Rows() + q.Rows()
			if all >= max {
				done <- true
			}
		}
	}(insert, query, conf.Max_request)

	select {
	case <-time.After(time.Duration(conf.Max_time) * time.Second):
	case <-done:
	}

	// stop
	insert.Stop()
	query.Stop()
	monitor.Stop()
}
