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
	"sync/atomic"
	"time"
	"xworker"
)

type Range struct {
	stop     bool
	requests uint64
	order    string
	workers  []xworker.Worker
	lock     sync.WaitGroup
}

func NewRange(workers []xworker.Worker, order string) xworker.QueryHandler {
	return &Range{
		workers: workers,
		order:   order,
	}
}

func (r *Range) Run() {
	threads := len(r.workers)
	for i := 0; i < threads; i++ {
		r.lock.Add(1)
		go r.Query(&r.workers[i], threads, i)
	}
}

func (r *Range) Stop() {
	r.stop = true
	r.lock.Wait()
}

func (r *Range) Rows() uint64 {
	return atomic.LoadUint64(&r.requests)
}

func (r *Range) Query(worker *xworker.Worker, num int, id int) {
	session := worker.S
	for !r.stop {
		table := rand.Int31n(int32(worker.N))
		sql := fmt.Sprintf("select price,dateandtime,customerid from purchases_index%d force index (pdc) where (price>=%.2f) order by price,dateandtime,customerid %s limit 10",
			table,
			float32(rand.Int31n(10000))/100,
			r.order)

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
		atomic.AddUint64(&r.requests, 1)
	}
	r.lock.Done()
}
