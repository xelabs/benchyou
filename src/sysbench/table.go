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
	"log"
)

type Table struct {
	workers []Worker
}

func NewTable(workers []Worker) *Table {
	return &Table{workers}
}

func (t *Table) Prepare() {
	session := t.workers[0].S
	count := t.workers[0].N
	engine := t.workers[0].E
	for i := 0; i < count; i++ {
		sql := fmt.Sprintf(`CREATE TABLE benchyou%d (
							id bigint(20) unsigned NOT NULL,
							k bigint(20) unsigned NOT NULL DEFAULT '0',
							c char(120) NOT NULL DEFAULT '',
							pad char(60) NOT NULL DEFAULT '',
							PRIMARY KEY (id),
							KEY k_1 (k)
							) ENGINE=%s`, i, engine)

		_, err := session.Exec(sql)
		if err != nil {
			log.Panicf("creata.table.error[%v]", err)
		}
		log.Printf("create table benchyou%d(engine=%v) finished...\n", i, engine)
	}
}

func (t *Table) Cleanup() {
	session := t.workers[0].S
	count := t.workers[0].N
	for i := 0; i < count; i++ {
		sql := fmt.Sprintf(`DROP TABLE benchyou%d;`, i)

		_, err := session.Exec(sql)
		if err != nil {
			log.Panicf("drop.table.error[%v]", err)
		}
		log.Printf("drop table benchyou%d finished...\n", i)
	}
}
