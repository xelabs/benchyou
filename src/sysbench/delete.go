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
)

type Delete struct {
	stop     bool
	requests uint64
	conf     *xcommon.BenchConf
	workers  []xworker.Worker
	lock     sync.WaitGroup
}

func NewDelete(conf *xcommon.BenchConf, workers []xworker.Worker) xworker.DeleteHandler {
	return &Delete{
		conf:    conf,
		workers: workers,
	}
}

func (delete *Delete) Run() {
	threads := len(delete.workers)
	for i := 0; i < threads; i++ {
		delete.lock.Add(1)
		go delete.Delete(&delete.workers[i], threads, i)
	}
}

func (delete *Delete) Stop() {
	delete.stop = true
	delete.lock.Wait()
}

func (delete *Delete) Rows() uint64 {
	return atomic.LoadUint64(&delete.requests)
}

func (delete *Delete) Delete(worker *xworker.Worker, num int, id int) {
	session := worker.S
	bs := int64(math.MaxInt64) / int64(num)
	lo := bs * int64(id)
	hi := bs * int64(id+1)

	for !delete.stop {
		var sql string
		var id int64

		if delete.conf.Random {
			id = xcommon.RandInt64(lo, hi)
		} else {
			id = lo
			lo++
		}
		table := rand.Int31n(int32(worker.N))
		sql = fmt.Sprintf("delete from benchyou%d where id=%v", table, id)

		t := time.Now()
		// Txn start.
		mod := worker.M.WNums % uint64(delete.conf.Batch_per_commit)
		if delete.conf.Batch_per_commit > 1 {
			if mod == 0 {
				if err := session.Exec("begin"); err != nil {
					log.Panicf("delete.error[%v]", err)
				}
			}
		}
		// XA start.
		if delete.conf.XA {
			xaStart(worker, hi, lo)
		}
		if err := session.Exec(sql); err != nil {
			log.Panicf("delete.error[%v]", err)
		}
		// XA end.
		if delete.conf.XA {
			xaEnd(worker)
		}
		// Txn end.
		if delete.conf.Batch_per_commit > 1 {
			if mod == uint64(delete.conf.Batch_per_commit-1) {
				if err := session.Exec("commit"); err != nil {
					log.Panicf("delete.error[%v]", err)
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
		atomic.AddUint64(&delete.requests, 1)
	}
	delete.lock.Done()
}
