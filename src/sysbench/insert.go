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
)

type Insert struct {
	workers []Worker
	stop    bool
	lock    sync.WaitGroup
	random  bool
}

func NewInsert(workers []Worker, random bool) *Insert {
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

func (insert *Insert) Insert(worker *Worker, num int, id int) {
	var rid, rk int64
	session := worker.S
	bs := int64(math.MaxInt64) / int64(num)
	lo := bs * int64(id)
	hi := bs * int64(id+1)

	for !insert.stop {
		if insert.random {
			rid = xcommon.RandInt64(lo, hi)
			rk = xcommon.RandInt64(lo, hi)
		} else {
			rid = lo
			rk = rid
			lo++
		}

		pad := xcommon.RandString(padtemplate)
		c := xcommon.RandString(ctemplate)
		table := rand.Int31n(int32(worker.N))
		sql := fmt.Sprintf("insert into benchyou%d(id,k,c,pad) values(%v,%v,'%s', '%s')",
			table, rid, rk, c, pad)

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
