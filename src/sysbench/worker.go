/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package sysbench

import (
	"fmt"
	"github.com/XeLabs/go-mysqlstack/driver"
	"xcommon"
)

type Metric struct {
	WNums  uint64
	WCosts uint64
	WMax   uint64
	WMin   uint64
	QNums  uint64
	QCosts uint64
	QMax   uint64
	QMin   uint64
}

type Worker struct {
	// session
	S *driver.Conn

	// mertric
	M *Metric

	// engine
	E string

	// table number
	N int
}

func CreateWorkers(conf *xcommon.Conf, threads int) ([]Worker, error) {
	var workers []Worker
	for i := 0; i < threads; i++ {
		conn, err := driver.NewConn(
			conf.Mysql_user,
			conf.Mysql_password,
			"tcp",
			fmt.Sprintf("%s:%d", conf.Mysql_host, conf.Mysql_port),
			conf.Mysql_db)
		if err != nil {
			return nil, err
		}
		workers = append(workers, Worker{
			S: conn,
			M: &Metric{},
			E: conf.Mysql_table_engine,
			N: conf.Oltp_tables_count,
		},
		)
	}

	return workers, nil
}

func AllWorkersMetric(workers []Worker) *Metric {
	all := &Metric{}
	for _, worker := range workers {
		all.WNums += worker.M.WNums
		all.WCosts += worker.M.WCosts

		if all.WMax < worker.M.WMax {
			all.WMax = worker.M.WMax
		}

		if all.WMin > worker.M.WMin {
			all.WMin = worker.M.WMin
		}

		all.QNums += worker.M.QNums
		all.QCosts += worker.M.QCosts

		if all.QMax < worker.M.QMax {
			all.QMax = worker.M.QMax
		}

		if all.QMin > worker.M.QMin {
			all.QMin = worker.M.QMin
		}
	}

	return all
}

func StopWorkers(workers []Worker) {
	for _, worker := range workers {
		worker.S.Close()
	}
}
