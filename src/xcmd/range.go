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
	"xworker"
)

// NewRangeCommand creates the new cmd.
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
	conf.Random = true

	// worker
	wthds := conf.WriteThreads
	rthds := conf.ReadThreads
	thds := wthds + rthds
	workers := xworker.CreateWorkers(conf, thds)

	// monitor
	monitor := NewMonitor(conf, workers)

	// insert
	iworker := workers[:wthds]
	qworker := workers[wthds:]
	insert := sysbench.NewInsert(conf, iworker)
	query := sysbench.NewRange(conf, qworker, conf.MysqlRangeOrder)

	// start
	insert.Run()
	query.Run()
	monitor.Start()

	done := make(chan bool)
	go func(i xworker.Handler, q xworker.Handler, max uint64) {
		if max == 0 {
			return
		}

		for {
			time.Sleep(time.Second)
			all := i.Rows() + q.Rows()
			if all >= max {
				done <- true
			}
		}
	}(insert, query, conf.MaxRequest)

	select {
	case <-time.After(time.Duration(conf.MaxTime) * time.Second):
	case <-done:
	}

	// stop
	insert.Stop()
	query.Stop()
	monitor.Stop()
}
