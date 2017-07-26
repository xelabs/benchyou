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
	"testing"
	"time"
	"xcommon"
	"xworker"

	"github.com/stretchr/testify/assert"
)

func TestSysbenchInsert(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	conf := xcommon.MockConf(mysql.Addr())

	workers := xworker.CreateWorkers(conf, 2)
	assert.NotNil(t, workers)

	job := NewInsert(conf, workers)
	job.Run()
	time.Sleep(time.Millisecond * 100)
	job.Stop()
	assert.True(t, job.Rows() > 0)
}

func TestSysbenchBatchInsert(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	conf := xcommon.MockConf(mysql.Addr())
	conf.Batch_per_commit = 10

	workers := xworker.CreateWorkers(conf, 2)
	assert.NotNil(t, workers)
	job := NewInsert(conf, workers)
	job.Run()
	time.Sleep(time.Millisecond * 100)
	job.Stop()
	assert.True(t, job.Rows() > 0)
}

func TestSysbenchRandomInsert(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	conf := xcommon.MockConf(mysql.Addr())
	conf.Random = true

	workers := xworker.CreateWorkers(conf, 2)
	assert.NotNil(t, workers)
	job := NewInsert(conf, workers)
	job.Run()
	time.Sleep(time.Millisecond * 100)
	job.Stop()
	assert.True(t, job.Rows() > 0)
}

func TestSysbenchXAInsert(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	conf := xcommon.MockConf(mysql.Addr())
	conf.XA = true

	workers := xworker.CreateWorkers(conf, 2)
	assert.NotNil(t, workers)
	job := NewInsert(conf, workers)
	job.Run()
	time.Sleep(time.Millisecond * 100)
	job.Stop()
	assert.True(t, job.Rows() > 0)
}
