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
	switch conf.Bench_mode {
	case "sysbench":
		insert = sysbench.NewInsert(iworker, true)
		query = sysbench.NewQuery(qworker, true)

	case "iibench":
		insert = iibench.NewInsert(iworker, true)
		query = iibench.NewQuery(qworker, true)
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
