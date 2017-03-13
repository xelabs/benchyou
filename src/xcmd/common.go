/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xcmd

import (
	"github.com/spf13/cobra"
	"xcommon"
)

func parseConf(cmd *cobra.Command) (conf *xcommon.Conf, err error) {
	conf = &xcommon.Conf{}

	if conf.Write_threads, err = cmd.Flags().GetInt("write-threads"); err != nil {
		return
	}

	if conf.Read_threads, err = cmd.Flags().GetInt("read-threads"); err != nil {
		return
	}

	if conf.Mysql_host, err = cmd.Flags().GetString("mysql-host"); err != nil {
		return
	}

	if conf.Ssh_host, err = cmd.Flags().GetString("ssh-host"); err != nil {
		return
	}
	if conf.Ssh_host == "" {
		conf.Ssh_host = conf.Mysql_host
	}

	if conf.Ssh_user, err = cmd.Flags().GetString("ssh-user"); err != nil {
		return
	}

	if conf.Ssh_password, err = cmd.Flags().GetString("ssh-password"); err != nil {
		return
	}

	if conf.Ssh_port, err = cmd.Flags().GetInt("ssh-port"); err != nil {
		return
	}

	if conf.Mysql_user, err = cmd.Flags().GetString("mysql-user"); err != nil {
		return
	}

	if conf.Mysql_password, err = cmd.Flags().GetString("mysql-password"); err != nil {
		return
	}

	if conf.Mysql_port, err = cmd.Flags().GetInt("mysql-port"); err != nil {
		return
	}

	if conf.Mysql_db, err = cmd.Flags().GetString("mysql-db"); err != nil {
		return
	}

	if conf.Mysql_table_engine, err = cmd.Flags().GetString("mysql-table-engine"); err != nil {
		return
	}

	if conf.Rows_per_commit, err = cmd.Flags().GetInt("rows-per-commit"); err != nil {
		return
	}

	if conf.Oltp_tables_count, err = cmd.Flags().GetInt("oltp-tables-count"); err != nil {
		return
	}

	if conf.Max_time, err = cmd.Flags().GetInt("max-time"); err != nil {
		return
	}

	if conf.Mysql_range_order, err = cmd.Flags().GetString("mysql-range-order"); err != nil {
		return
	}

	xa := 0
	if xa, err = cmd.Flags().GetInt("mysql-enable-xa"); err != nil {
		return
	}
	if xa > 0 {
		conf.XA = true
	}

	if conf.Bench_mode, err = cmd.Flags().GetString("bench-mode"); err != nil {
		return
	}

	return
}
