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
	"benchyou/xcommon"
	"benchyou/xworker"
)

// Range tuple.
type Range struct {
	stop     bool
	requests uint64
	order    string
	workers  []xworker.Worker
	lock     sync.WaitGroup
}

// NewRange creates the new range handler.
func NewRange(conf *xcommon.Conf, workers []xworker.Worker, order string) *Range {
	return &Range{
		workers: workers,
		order:   order,
	}
}

// Run used to start the worker.
func (r *Range) Run() {
	threads := len(r.workers)
	for i := 0; i < threads; i++ {
		r.lock.Add(1)
		go r.Query(&r.workers[i], threads, i)
	}
}

// Stop used to stop the worker.
func (r *Range) Stop() {
	r.stop = true
	r.lock.Wait()
}

// Rows returns the row numbers.
func (r *Range) Rows() uint64 {
	return atomic.LoadUint64(&r.requests)
}

// Query used to execute the range query.
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
		atomic.AddUint64(&r.requests, 1)
	}
	r.lock.Done()
}
