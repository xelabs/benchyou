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
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"xcommon"
	"xworker"

	"github.com/xelabs/go-mysqlstack/sqlparser/depends/common"
)

// Insert tuple.
type Insert struct {
	stop     bool
	requests uint64
	conf     *xcommon.Conf
	workers  []xworker.Worker
	lock     sync.WaitGroup
}

// NewInsert creates the new insert handler.
func NewInsert(conf *xcommon.Conf, workers []xworker.Worker) xworker.Handler {
	return &Insert{
		conf:    conf,
		workers: workers,
	}
}

// Run used to start the worker.
func (insert *Insert) Run() {
	threads := len(insert.workers)
	for i := 0; i < threads; i++ {
		insert.lock.Add(1)
		go insert.Insert(&insert.workers[i], threads, i)
	}
}

// Stop used to stop the worker.
func (insert *Insert) Stop() {
	insert.stop = true
	insert.lock.Wait()
}

// Rows returns the row numbers.
func (insert *Insert) Rows() uint64 {
	return atomic.LoadUint64(&insert.requests)
}

// Insert used to execute the insert query.
func (insert *Insert) Insert(worker *xworker.Worker, num int, id int) {
	session := worker.S
	bs := int64(math.MaxInt64) / int64(num)
	lo := bs * int64(id)
	hi := bs * int64(id+1)
	columns1 := "k,c,pad"
	columns2 := "k,c,pad,id"
	valfmt1 := "(%v,'%s', '%s'),"
	valfmt2 := "(%v,'%s', '%s', %v),"

	for !insert.stop {
		var sql, value string
		buf := common.NewBuffer(256)

		table := rand.Int31n(int32(worker.N))
		if insert.conf.Random {
			sql = fmt.Sprintf("insert into benchyou%d(%s) values", table, columns2)
		} else {
			sql = fmt.Sprintf("insert into benchyou%d(%s) values", table, columns1)
		}

		// pack requests
		for n := 0; n < insert.conf.RowsPerInsert; n++ {
			pad := xcommon.RandString(xcommon.Padtemplate)
			c := xcommon.RandString(xcommon.Ctemplate)

			if insert.conf.Random {
				value = fmt.Sprintf(valfmt2,
					xcommon.RandInt64(lo, hi),
					c,
					pad,
					xcommon.RandInt64(lo, hi),
				)
			} else {
				value = fmt.Sprintf(valfmt1,
					xcommon.RandInt64(lo, hi),
					c,
					pad,
				)
			}
			buf.WriteString(value)
		}
		// -1 to trim right ','
		vals, err := buf.ReadString(buf.Length() - 1)
		if err != nil {
			log.Panicf("insert.error[%v]", err)
		}
		sql += vals

		t := time.Now()
		// Txn start.
		mod := worker.M.WNums % uint64(insert.conf.BatchPerCommit)
		if insert.conf.BatchPerCommit > 1 {
			if mod == 0 {
				if err := session.Exec("begin"); err != nil {
					log.Panicf("insert.error[%v]", err)
				}
			}
		}
		// XA start.
		if insert.conf.XA {
			xaStart(worker, hi, lo)
		}
		if err := session.Exec(sql); err != nil {
			log.Panicf("insert.error[%v]", err)
		}
		// XA end.
		if insert.conf.XA {
			xaEnd(worker)
		}
		// Txn end.
		if insert.conf.BatchPerCommit > 1 {
			if mod == uint64(insert.conf.BatchPerCommit-1) {
				if err := session.Exec("commit"); err != nil {
					log.Panicf("insert.error[%v]", err)
				}
			}
		}
		elapsed := time.Since(t)

		// stats
		nsec := uint64(elapsed.Nanoseconds())
		worker.M.WCosts += nsec
		if worker.M.WMax == 0 && worker.M.WMin == 0 {
			worker.M.WMax = nsec
			worker.M.WMin = nsec
		}

		if nsec > worker.M.WMax {
			worker.M.WMax = nsec
		}
		if nsec < worker.M.WMin {
			worker.M.WMin = nsec
		}
		worker.M.WNums++
		atomic.AddUint64(&insert.requests, 1)
	}
	insert.lock.Done()
}

func xaStart(worker *xworker.Worker, hi int64, lo int64) {
	session := worker.S
	worker.XID = fmt.Sprintf("BXID-%v-%v", time.Now().Format("20060102150405"), (rand.Int63n(hi-lo) + lo))
	start := fmt.Sprintf("xa start '%s'", worker.XID)
	if err := session.Exec(start); err != nil {
		log.Panicf("xa.start..error[%v]", err)
	}
}

func xaEnd(worker *xworker.Worker) {
	session := worker.S
	end := fmt.Sprintf("xa end '%s'", worker.XID)
	if err := session.Exec(end); err != nil {
		log.Panicf("xa.end.error[%v]", err)
	}
	prepare := fmt.Sprintf("xa prepare '%s'", worker.XID)
	if err := session.Exec(prepare); err != nil {
		log.Panicf("xa.prepare.error[%v]", err)
	}
	commit := fmt.Sprintf("xa commit '%s'", worker.XID)
	if err := session.Exec(commit); err != nil {
		log.Panicf("xa.commit.error[%v]", err)
	}
}
