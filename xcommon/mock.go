package xcommon

import (
	"net"
	"strconv"

	"github.com/xelabs/go-mysqlstack/driver"
	"github.com/xelabs/go-mysqlstack/sqlparser/depends/sqltypes"
	"github.com/xelabs/go-mysqlstack/xlog"
)

// MockConf mocks conf.
func MockConf(addr string) *Conf {
	host, sport, _ := net.SplitHostPort(addr)
	port, _ := strconv.Atoi(sport)

	return &Conf{
		MysqlHost:       host,
		MysqlPort:       port,
		MysqlUser:       "mock",
		MaxRequest:      16,
		OltpTablesCount: 1,
		RowsPerInsert:   1,
		BatchPerCommit:  1,
	}
}

// MockMySQL mocks mysql.
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
