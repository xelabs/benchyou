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

func parseConf(cmd *cobra.Command) (conf *xcommon.Conf, err error) {
	conf = &xcommon.Conf{}

	if conf.WriteThreads, err = cmd.Flags().GetInt("write-threads"); err != nil {
		return
	}

	if conf.UpdateThreads, err = cmd.Flags().GetInt("update-threads"); err != nil {
		return
	}

	if conf.DeleteThreads, err = cmd.Flags().GetInt("delete-threads"); err != nil {
		return
	}

	if conf.ReadThreads, err = cmd.Flags().GetInt("read-threads"); err != nil {
		return
	}

	if conf.MysqlHost, err = cmd.Flags().GetString("mysql-host"); err != nil {
		return
	}

	if conf.SSHHost, err = cmd.Flags().GetString("ssh-host"); err != nil {
		return
	}
	if conf.SSHHost == "" {
		conf.SSHHost = conf.MysqlHost
	}

	if conf.SSHUser, err = cmd.Flags().GetString("ssh-user"); err != nil {
		return
	}

	if conf.SSHPassword, err = cmd.Flags().GetString("ssh-password"); err != nil {
		return
	}

	if conf.SSHPort, err = cmd.Flags().GetInt("ssh-port"); err != nil {
		return
	}

	if conf.MysqlUser, err = cmd.Flags().GetString("mysql-user"); err != nil {
		return
	}

	if conf.MysqlPassword, err = cmd.Flags().GetString("mysql-password"); err != nil {
		return
	}

	if conf.MysqlPort, err = cmd.Flags().GetInt("mysql-port"); err != nil {
		return
	}

	if conf.MysqlDb, err = cmd.Flags().GetString("mysql-db"); err != nil {
		return
	}

	if conf.MysqlTableEngine, err = cmd.Flags().GetString("mysql-table-engine"); err != nil {
		return
	}

	if conf.OltpTablesCount, err = cmd.Flags().GetInt("oltp-tables-count"); err != nil {
		return
	}

	if conf.RowsPerInsert, err = cmd.Flags().GetInt("rows-per-insert"); err != nil {
		return
	}

	if conf.BatchPerCommit, err = cmd.Flags().GetInt("batch-per-commit"); err != nil {
		return
	}

	if conf.MaxTime, err = cmd.Flags().GetInt("max-time"); err != nil {
		return
	}

	if conf.MaxRequest, err = cmd.Flags().GetUint64("max-request"); err != nil {
		return
	}

	if conf.MysqlRangeOrder, err = cmd.Flags().GetString("mysql-range-order"); err != nil {
		return
	}

	xa := 0
	if xa, err = cmd.Flags().GetInt("mysql-enable-xa"); err != nil {
		return
	}
	if xa > 0 {
		conf.XA = true
	}
	return
}

func start(conf *xcommon.Conf) {
	// worker
	var workers []xworker.Worker
	wthds := conf.WriteThreads
	rthds := conf.ReadThreads
	uthds := conf.UpdateThreads
	dthds := conf.DeleteThreads

	// workers
	iworkers := xworker.CreateWorkers(conf, wthds)
	insert := sysbench.NewInsert(conf, iworkers)
	workers = append(workers, iworkers...)

	qworkers := xworker.CreateWorkers(conf, rthds)
	query := sysbench.NewQuery(conf, qworkers)
	workers = append(workers, qworkers...)

	uworkers := xworker.CreateWorkers(conf, uthds)
	update := sysbench.NewUpdate(conf, uworkers)
	workers = append(workers, uworkers...)

	dworkers := xworker.CreateWorkers(conf, dthds)
	delete := sysbench.NewDelete(conf, dworkers)
	workers = append(workers, dworkers...)

	// monitor
	monitor := NewMonitor(conf, workers)

	// start
	insert.Run()
	query.Run()
	update.Run()
	delete.Run()
	monitor.Start()

	done := make(chan bool)
	go func(i xworker.Handler, q xworker.Handler, u xworker.Handler, d xworker.Handler, max uint64) {
		if max == 0 {
			return
		}

		for {
			time.Sleep(time.Millisecond * 10)
			all := i.Rows() + q.Rows() + u.Rows() + d.Rows()
			if all >= max {
				done <- true
			}
		}
	}(insert, query, update, delete, conf.MaxRequest)

	select {
	case <-time.After(time.Duration(conf.MaxTime) * time.Second):
	case <-done:
	}

	// stop
	insert.Stop()
	query.Stop()
	update.Stop()
	delete.Stop()
	monitor.Stop()
}
