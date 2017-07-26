/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xworker

import (
	"testing"
	"xcommon"

	"github.com/stretchr/testify/assert"
)

func TestXWorker(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	conf := xcommon.MockConf(mysql.Addr())

	workers := CreateWorkers(conf, 2)
	assert.NotNil(t, workers)
	w1 := workers[0]
	w1.M.WMax = 2
	w1.M.WMin = 1
	w1.M.QNums = 100
	w1.M.QMax = 10
	w1.M.QMin = 1

	metric := AllWorkersMetric(workers)
	assert.NotNil(t, metric)
	StopWorkers(workers)
}
