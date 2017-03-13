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
	"time"
	"xcommon"
	"xworker"
)

type Query struct {
	stop    bool
	conf    *xcommon.BenchConf
	workers []xworker.Worker
	lock    sync.WaitGroup
}

func NewQuery(conf *xcommon.BenchConf, workers []xworker.Worker) xworker.QueryHandler {
	return &Query{
		conf:    conf,
		workers: workers,
	}
}

func (q *Query) Run() {
	threads := len(q.workers)
	for i := 0; i < threads; i++ {
		q.lock.Add(1)
		go q.Query(&q.workers[i], threads, i)
	}
}

func (q *Query) Stop() {
	q.stop = true
	q.lock.Wait()
}

func (q *Query) Query(worker *xworker.Worker, num int, id int) {
	var rid int64
	session := worker.S
	bs := int64(math.MaxInt64) / int64(num)
	lo := bs * int64(id)
	hi := bs * int64(id+1)

	for !q.stop {
		if q.conf.Random {
			rid = xcommon.RandInt64(lo, hi)
		} else {
			rid = lo
			lo++
		}

		table := rand.Int31n(int32(worker.N))
		sql := fmt.Sprintf("SELECT * FROM benchyou%d WHERE id=%v", table, rid)
		t := time.Now()
		if err := session.Exec(sql); err != nil {
			log.Panicf("query.error[%v]", err)
		}
		elapsed := time.Since(t)

		// stats
		nsec := uint64(elapsed.Nanoseconds())
		worker.M.QCosts += nsec
		if worker.M.QMax == 0 && worker.M.QMin == 0 {
			worker.M.QMax = nsec
			worker.M.QMin = nsec
		}
		if nsec > worker.M.QMax {
			worker.M.QMax = nsec
		}
		if nsec < worker.M.QMin {
			worker.M.QMin = nsec
		}
		worker.M.QNums++
	}
	q.lock.Done()
}
