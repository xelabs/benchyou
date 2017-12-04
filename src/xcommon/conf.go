/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xcommon

// Conf tuple.
type Conf struct {
	WriteThreads     int
	ReadThreads      int
	DeleteThreads    int
	UpdateThreads    int
	SSHHost          string
	SSHUser          string
	SSHPassword      string
	SSHPort          int
	MysqlHost        string
	MysqlUser        string
	MysqlPassword    string
	MysqlPort        int
	MysqlDb          string
	MysqlTableEngine string
	MysqlRangeOrder  string
	RowsPerInsert    int
	BatchPerCommit   int
	MaxTime          int
	MaxRequest       uint64
	OltpTablesCount  int
	XA               bool
	Random           bool
}
