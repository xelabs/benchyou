/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xcommon

type Conf struct {
	Write_threads      int
	Read_threads       int
	Ssh_host           string
	Ssh_user           string
	Ssh_password       string
	Ssh_port           int
	Mysql_host         string
	Mysql_user         string
	Mysql_password     string
	Mysql_port         int
	Mysql_db           string
	Mysql_table_engine string
	Mysql_range_order  string
	Rows_per_commit    int
	Max_time           int
	Oltp_tables_count  int
	Bench_mode         string
}
