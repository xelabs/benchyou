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
	"time"
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
	workers, err := sysbench.CreateWorkers(conf, thds)
	if err != nil {
		log.Panicf("create.workers.error:[%+v]", err)
	}

	// monitor
	monitor := NewMonitor(conf, workers)

	// insert
	iworker := workers[:wthds]
	insert := sysbench.NewInsert(iworker, true)
	insert.Run()

	// range query
	qworker := workers[wthds:]
	query := sysbench.NewRange(qworker, conf.Mysql_range_order)
	query.Run()

	monitor.Start()
	// wait
	time.Sleep(time.Duration(conf.Max_time) * time.Second)

	// stop
	insert.Stop()
	query.Stop()
	monitor.Stop()
}
