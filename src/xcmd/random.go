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

	qworker := workers[wthds:]
	query := sysbench.NewQuery(qworker, true)
	query.Run()

	monitor.Start()
	// wait
	time.Sleep(time.Duration(conf.Max_time) * time.Second)

	// stop
	insert.Stop()
	query.Stop()
	monitor.Stop()
}
