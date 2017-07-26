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
	"xcommon"
	"xworker"

	"github.com/stretchr/testify/assert"
)

func TestSysbenchTable(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	conf := xcommon.MockConf(mysql.Addr())

	workers := xworker.CreateWorkers(conf, 2)
	assert.NotNil(t, workers)

	job := NewTable(workers)
	job.Prepare()
	job.Cleanup()
}
