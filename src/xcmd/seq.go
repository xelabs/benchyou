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
	"xcommon"
	"xworker"
)

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

	// workers
	var insert xworker.InsertHandler
	var query xworker.QueryHandler
	iworker := workers[:wthds]
	qworker := workers[wthds:]
	benchConf := &xcommon.BenchConf{
		Random:          false,
		XA:              conf.XA,
		Rows_per_commit: conf.Rows_per_commit,
	}
	switch conf.Bench_mode {
	case "sysbench":
		insert = sysbench.NewInsert(benchConf, iworker)
		query = sysbench.NewQuery(benchConf, qworker)

	case "iibench":
		insert = iibench.NewInsert(benchConf, iworker)
		query = iibench.NewQuery(benchConf, qworker)
	}

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
