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

type Range struct {
	workers []xworker.Worker
	stop    bool
	lock    sync.WaitGroup
	order   string
}

func NewRange(workers []xworker.Worker, order string) *Range {
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

func (r *Range) Query(worker *xworker.Worker, num int, id int) {
	session := worker.S
	bs := int64(math.MaxInt64) / int64(num)
	lo := bs * int64(id)
	hi := bs * int64(id+1)

	for !r.stop {
		lower := xcommon.RandInt64(lo, hi)
		upper := xcommon.RandInt64(lower, hi)

		table := rand.Int31n(int32(worker.N))
		sql := fmt.Sprintf("SELECT * FROM benchyou%d WHERE id BETWEEN %d AND %d ORDER BY id %v LIMIT 100",
			table, lower, upper, r.order)
		t := time.Now()
		if err := session.Exec(sql); err != nil {
			log.Panicf("range.error[%v]", err)
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
	r.lock.Done()
}
