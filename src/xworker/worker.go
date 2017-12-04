/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xworker

import (
	"fmt"
	"log"
	"xcommon"

	"github.com/XeLabs/go-mysqlstack/driver"
)

// Metric tuple.
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

// Worker tuple.
type Worker struct {
	// session
	S driver.Conn

	// mertric
	M *Metric

	// engine
	E string

	// xid
	XID string

	// table number
	N int
}

// CreateWorkers creates the new workers.
func CreateWorkers(conf *xcommon.Conf, threads int) []Worker {
	var workers []Worker
	var conn driver.Conn
	var err error

	// Check database is exists or not.
	utf8 := "utf8"
	dsn := fmt.Sprintf("%s:%d", conf.MysqlHost, conf.MysqlPort)
	if conn, err = driver.NewConn(conf.MysqlUser, conf.MysqlPassword, dsn, "", utf8); err != nil {
		log.Panicf("create.worker.check.database.error:%+v", err)
	}
	sql := fmt.Sprintf("create database if not exists `%s`", conf.MysqlDb)
	if err := conn.Exec(sql); err != nil {
		log.Panicf("create.worker.check.database.exec[%s].error:%+v", sql, err)
	}

	for i := 0; i < threads; i++ {
		if conn, err = driver.NewConn(conf.MysqlUser, conf.MysqlPassword, dsn, conf.MysqlDb, utf8); err != nil {
			log.Panicf("create.worker.error:%v", err)
		}
		workers = append(workers, Worker{
			S: conn,
			M: &Metric{},
			E: conf.MysqlTableEngine,
			N: conf.OltpTablesCount,
		},
		)
	}
	return workers
}

// AllWorkersMetric returns all the worker's metric.
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

// StopWorkers used to stop all the worker.
func StopWorkers(workers []Worker) {
	for _, worker := range workers {
		worker.S.Close()
	}
}
