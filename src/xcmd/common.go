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

	if conf.Write_threads, err = cmd.Flags().GetInt("write-threads"); err != nil {
		return
	}

	if conf.Update_threads, err = cmd.Flags().GetInt("update-threads"); err != nil {
		return
	}

	if conf.Delete_threads, err = cmd.Flags().GetInt("delete-threads"); err != nil {
		return
	}

	if conf.Read_threads, err = cmd.Flags().GetInt("read-threads"); err != nil {
		return
	}

	if conf.Mysql_host, err = cmd.Flags().GetString("mysql-host"); err != nil {
		return
	}

	if conf.Ssh_host, err = cmd.Flags().GetString("ssh-host"); err != nil {
		return
	}
	if conf.Ssh_host == "" {
		conf.Ssh_host = conf.Mysql_host
	}

	if conf.Ssh_user, err = cmd.Flags().GetString("ssh-user"); err != nil {
		return
	}

	if conf.Ssh_password, err = cmd.Flags().GetString("ssh-password"); err != nil {
		return
	}

	if conf.Ssh_port, err = cmd.Flags().GetInt("ssh-port"); err != nil {
		return
	}

	if conf.Mysql_user, err = cmd.Flags().GetString("mysql-user"); err != nil {
		return
	}

	if conf.Mysql_password, err = cmd.Flags().GetString("mysql-password"); err != nil {
		return
	}

	if conf.Mysql_port, err = cmd.Flags().GetInt("mysql-port"); err != nil {
		return
	}

	if conf.Mysql_db, err = cmd.Flags().GetString("mysql-db"); err != nil {
		return
	}

	if conf.Mysql_table_engine, err = cmd.Flags().GetString("mysql-table-engine"); err != nil {
		return
	}

	if conf.Oltp_tables_count, err = cmd.Flags().GetInt("oltp-tables-count"); err != nil {
		return
	}

	if conf.Rows_per_insert, err = cmd.Flags().GetInt("rows-per-insert"); err != nil {
		return
	}

	if conf.Batch_per_commit, err = cmd.Flags().GetInt("batch-per-commit"); err != nil {
		return
	}

	if conf.Max_time, err = cmd.Flags().GetInt("max-time"); err != nil {
		return
	}

	if conf.Max_request, err = cmd.Flags().GetUint64("max-request"); err != nil {
		return
	}

	if conf.Mysql_range_order, err = cmd.Flags().GetString("mysql-range-order"); err != nil {
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

func start(conf *xcommon.Conf, benchConf *xcommon.BenchConf) {
	// worker
	var workers []xworker.Worker
	wthds := conf.Write_threads
	rthds := conf.Read_threads
	uthds := conf.Update_threads
	dthds := conf.Delete_threads

	// workers
	iworkers := xworker.CreateWorkers(conf, wthds)
	insert := sysbench.NewInsert(benchConf, iworkers)
	workers = append(workers, iworkers...)

	qworkers := xworker.CreateWorkers(conf, rthds)
	query := sysbench.NewQuery(benchConf, qworkers)
	workers = append(workers, qworkers...)

	uworkers := xworker.CreateWorkers(conf, uthds)
	update := sysbench.NewUpdate(benchConf, uworkers)
	workers = append(workers, uworkers...)

	dworkers := xworker.CreateWorkers(conf, dthds)
	delete := sysbench.NewDelete(benchConf, dworkers)
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
	go func(i xworker.InsertHandler, q xworker.QueryHandler, u xworker.UpdateHandler, d xworker.DeleteHandler, max uint64) {
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
	}(insert, query, update, delete, conf.Max_request)

	select {
	case <-time.After(time.Duration(conf.Max_time) * time.Second):
	case <-done:
	}

	// stop
	insert.Stop()
	query.Stop()
	update.Stop()
	delete.Stop()
	monitor.Stop()
}
