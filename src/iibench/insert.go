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
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"xcommon"
	"xworker"

	"github.com/XeLabs/go-mysqlstack/common"
)

type Insert struct {
	stop     bool
	requests uint64
	conf     *xcommon.BenchConf
	workers  []xworker.Worker
	lock     sync.WaitGroup
}

func NewInsert(conf *xcommon.BenchConf, workers []xworker.Worker) xworker.InsertHandler {
	return &Insert{
		conf:    conf,
		workers: workers,
	}
}

func (insert *Insert) Run() {
	threads := len(insert.workers)
	for i := 0; i < threads; i++ {
		insert.lock.Add(1)
		go insert.Insert(&insert.workers[i], threads, i)
	}
}

func (insert *Insert) Stop() {
	insert.stop = true
	insert.lock.Wait()
}

func (insert *Insert) Rows() uint64 {
	return atomic.LoadUint64(&insert.requests)
}

func (insert *Insert) Insert(worker *xworker.Worker, num int, id int) {
	session := worker.S
	bs := int64(math.MaxInt64) / int64(num)
	lo := bs * int64(id)
	hi := bs * int64(id+1)
	columns1 := "dateandtime,cashregisterid,customerid,productid,price,data"
	columns2 := "dateandtime,cashregisterid,customerid,productid,price,data, transactionid"
	valfmt1 := "('%v',%v,%v,%v,%.2f,'%s'),"
	valfmt2 := "('%v',%v,%v,%v,%.2f,'%s', %v),"

	for !insert.stop {
		var sql, value string
		buf := common.NewBuffer(256)

		table := rand.Int31n(int32(worker.N))
		if insert.conf.Random {
			sql = fmt.Sprintf("insert into purchases_index%d(%s) values", table, columns2)
		} else {
			sql = fmt.Sprintf("insert into purchases_index%d(%s) values", table, columns1)
		}

		// pack requests
		for n := 0; n < insert.conf.Rows_per_commit; n++ {
			c := xcommon.RandString(xcommon.Ctemplate)
			if insert.conf.Random {
				value = fmt.Sprintf(valfmt2,
					time.Now().Format("2006-01-02 15:04:05"),
					rand.Int31n(10000),
					rand.Int31n(1000000),
					rand.Int31n(1000000),
					float32(rand.Int31n(10000))/100,
					c,
					xcommon.RandInt64(lo, hi),
				)
			} else {
				value = fmt.Sprintf(valfmt1,
					time.Now().Format("2006-01-02 15:04:05"),
					rand.Int31n(10000),
					rand.Int31n(1000000),
					rand.Int31n(1000000),
					float32(rand.Int31n(10000))/100,
					c,
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

		if insert.conf.XA {
			worker.XID = fmt.Sprintf("BXID-%v-%v", time.Now().Format("20060102150405"), (rand.Int63n(hi-lo) + lo))
			start := fmt.Sprintf("xa start '%s'", worker.XID)
			if err = session.Exec(start); err != nil {
				log.Panicf("insert.error[%v]", err)
			}
		}
		if err = session.Exec(sql); err != nil {
			log.Panicf("insert.error[%v]", err)
		}
		if insert.conf.XA {
			end := fmt.Sprintf("xa end '%s'", worker.XID)
			if err = session.Exec(end); err != nil {
				log.Panicf("insert.error[%v]", err)
			}
			prepare := fmt.Sprintf("xa prepare '%s'", worker.XID)
			if err = session.Exec(prepare); err != nil {
				log.Panicf("insert.error[%v]", err)
			}
			commit := fmt.Sprintf("xa commit '%s'", worker.XID)
			if err = session.Exec(commit); err != nil {
				log.Panicf("insert.error[%v]", err)
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
