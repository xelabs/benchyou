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
	"math/rand"
	"sync"
	"time"
	"xworker"
)

type Query struct {
	stop    bool
	random  bool
	workers []xworker.Worker
	lock    sync.WaitGroup
}

func NewQuery(workers []xworker.Worker, random bool) xworker.QueryHandler {
	return &Query{
		workers: workers,
		random:  random,
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
	session := worker.S
	for !q.stop {
		table := rand.Int31n(int32(worker.N))
		sql := fmt.Sprintf("select price,dateandtime,customerid from purchases_index%d force index (pdc) where (price>=%.2f) order by price,dateandtime,customerid limit 1",
			table,
			float32(rand.Int31n(10000))/100)

		t := time.Now()
		_, err := session.Exec(sql)
		if err != nil {
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
