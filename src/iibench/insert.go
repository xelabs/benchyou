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
	"time"
	"xcommon"
	"xworker"
)

type Insert struct {
	stop    bool
	random  bool
	workers []xworker.Worker
	lock    sync.WaitGroup
}

func NewInsert(workers []xworker.Worker, random bool) xworker.InsertHandler {
	return &Insert{
		workers: workers,
		random:  random,
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

func (insert *Insert) Insert(worker *xworker.Worker, num int, id int) {
	var rid int64
	session := worker.S
	bs := int64(math.MaxInt64) / int64(num)
	lo := bs * int64(id)
	hi := bs * int64(id+1)

	for !insert.stop {
		c := xcommon.RandString(xcommon.Ctemplate)
		table := rand.Int31n(int32(worker.N))
		columns := "dateandtime,cashregisterid,customerid,productid,price,data"
		values := fmt.Sprintf("'%v',%v,%v,%v,%.2f,'%s'",
			time.Now().Format("2006-01-02 15:04:05"),
			rand.Int31n(10000),
			rand.Int31n(1000000),
			rand.Int31n(1000000),
			float32(rand.Int31n(10000))/100,
			c)

		if insert.random {
			rid = xcommon.RandInt64(lo, hi)
			columns += ",transactionid"
			values += fmt.Sprintf(",%v", rid)
		}

		sql := fmt.Sprintf(`insert into purchases_index%d(%s) values(%s)`,
			table,
			columns,
			values,
		)

		t := time.Now()
		_, err := session.Exec(sql)
		if err != nil {
			log.Panicf("insert.error[%v]", err)
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
	}
	insert.lock.Done()
}
