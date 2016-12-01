/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package iibench

import (
	"fmt"
	"log"
	"xworker"
)

type Table struct {
	workers []xworker.Worker
}

func NewTable(workers []xworker.Worker) *Table {
	return &Table{workers}
}

func (t *Table) Prepare() {
	session := t.workers[0].S
	count := t.workers[0].N
	engine := t.workers[0].E
	for i := 0; i < count; i++ {
		sql := fmt.Sprintf(`create table purchases_index%d (
							transactionid bigint(20) unsigned not null auto_increment,
							dateandtime datetime,
							cashregisterid int not null,
							customerid int not null,
							productid int not null,
							price float not null,
							data varchar(4000),
							primary key (transactionid),
							index pdc (price, dateandtime, customerid)
							) engine=%s`, i, engine)

		_, err := session.Exec(sql)
		if err != nil {
			log.Panicf("creata.table.error[%v]", err)
		}
		log.Printf("create table purchases_index%d(engine=%v) finished...\n", i, engine)
	}
}

func (t *Table) Cleanup() {
	session := t.workers[0].S
	count := t.workers[0].N
	for i := 0; i < count; i++ {
		sql := fmt.Sprintf(`drop table purchases_index%d;`, i)

		_, err := session.Exec(sql)
		if err != nil {
			log.Panicf("drop.table.error[%v]", err)
		}
		log.Printf("drop table purchases_index%d finished...\n", i)
	}
}
