package xcommon

import (
	"net"
	"strconv"

	"github.com/XeLabs/go-mysqlstack/driver"
	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
	"github.com/XeLabs/go-mysqlstack/xlog"
)

func MockConf(addr string) *Conf {
	host, sport, _ := net.SplitHostPort(addr)
	port, _ := strconv.Atoi(sport)

	return &Conf{
		Mysql_host:        host,
		Mysql_port:        port,
		Mysql_user:        "mock",
		Max_request:       16,
		Oltp_tables_count: 1,
		Rows_per_insert:   1,
		Batch_per_commit:  1,
	}
}

func MockMySQL() (*driver.Listener, func()) {
	result1 := &sqltypes.Result{}

	log := xlog.NewStdLog(xlog.Level(xlog.ERROR))
	th := driver.NewTestHandler(log)
	svr, err := driver.MockMysqlServer(log, th)
	if err != nil {
		log.Panicf("mock.mysql.error:%+v", err)
	}

	// Querys.
	th.AddQueryPattern("insert .*", result1)
	th.AddQueryPattern("delete .*", result1)
	th.AddQueryPattern("update .*", result1)
	th.AddQueryPattern("select .*", result1)

	th.AddQueryPattern("create .*", result1)
	th.AddQueryPattern("drop .*", result1)

	th.AddQueryPattern("xa .*", result1)
	th.AddQueryPattern("begin", result1)
	th.AddQueryPattern("commit", result1)

	return svr, func() {
		svr.Close()
	}
}
