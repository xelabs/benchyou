/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"xcmd"
)

var (
	writeThreads      int
	readThreads       int
	updateThreads     int
	deleteThreads     int
	mysqlHost         string
	mysqlPort         int
	mysqlUser         string
	mysqlPassword     string
	mysqlDb           string
	mysqlTableEngine string
	mysqlRangeOrder  string
	mysqlEnableXa    int
	rowsPerInsert    int
	batchPerCommit   int
	maxTime           int
	maxRequest        uint64
	oltpTablesCount  int
	sshHost           string
	sshUser           string
	sshPassword       string
	sshPort           int
)

var (
	rootCmd = &cobra.Command{
		Use:        "benchyou",
		Short:      "",
		SuggestFor: []string{"benchyou"},
	}
)

func init() {
	cobra.EnableCommandSorting = false
	rootCmd.PersistentFlags().IntVar(&writeThreads, "write-threads", 32, "number of write threads to use(Default 32)")
	rootCmd.PersistentFlags().IntVar(&readThreads, "read-threads", 32, "number of read threads to use(Default 32)")
	rootCmd.PersistentFlags().IntVar(&updateThreads, "update-threads", 0, "number of update threads to use(Default 0)")
	rootCmd.PersistentFlags().IntVar(&deleteThreads, "delete-threads", 0, "number of delete threads to use(Default 0)")
	rootCmd.PersistentFlags().StringVar(&mysqlHost, "mysql-host", "", "MySQL server host(Default NULL)")
	rootCmd.PersistentFlags().IntVar(&mysqlPort, "mysql-port", 3306, "MySQL server port(Default 3306)")
	rootCmd.PersistentFlags().StringVar(&mysqlUser, "mysql-user", "benchyou", "MySQL user(Default benchyou)")
	rootCmd.PersistentFlags().StringVar(&mysqlPassword, "mysql-password", "benchyou", "MySQL password(Default benchyou)")
	rootCmd.PersistentFlags().StringVar(&mysqlDb, "mysql-db", "sbtest", "MySQL database name(Default sbtest)")
	rootCmd.PersistentFlags().StringVar(&mysqlTableEngine, "mysql-table-engine", "tokudb", "storage engine to use for the test table {tokudb,innodb,...}(Default tokudb)")
	rootCmd.PersistentFlags().StringVar(&mysqlRangeOrder, "mysql-range-order", "ASC", "range query sort the result-set in {ASC|DESC} (Default ASC)")
	rootCmd.PersistentFlags().IntVar(&mysqlEnableXa, "mysql-enable-xa", 0, "enable MySQL xa transaction for insertion {0|1} (Default 0)")
	rootCmd.PersistentFlags().IntVar(&rowsPerInsert, "rows-per-insert", 1, "#rows per insert(Default 1)")
	rootCmd.PersistentFlags().IntVar(&batchPerCommit, "batch-per-commit", 1, "#rows per transaction(Default 1)")
	rootCmd.PersistentFlags().IntVar(&maxTime, "max-time", 3600, "limit for total execution time in seconds(Default 3600)")
	rootCmd.PersistentFlags().Uint64Var(&maxRequest, "max-request", 0, "limit for total requests, including write and read(Default 0, means no limits)")
	rootCmd.PersistentFlags().IntVar(&oltpTablesCount, "oltp-tables-count", 8, "number of tables to create(Default 8)")
	rootCmd.PersistentFlags().StringVar(&sshHost, "ssh-host", "", "SSH server host(Default NULL, same as mysql-host)")
	rootCmd.PersistentFlags().StringVar(&sshUser, "ssh-user", "benchyou", "SSH server user(Default benchyou)")
	rootCmd.PersistentFlags().StringVar(&sshPassword, "ssh-password", "benchyou", "SSH server password(Default benchyou)")
	rootCmd.PersistentFlags().IntVar(&sshPort, "ssh-port", 22, "SSH server port(Default 22)")

	rootCmd.AddCommand(xcmd.NewPrepareCommand())
	rootCmd.AddCommand(xcmd.NewCleanupCommand())
	rootCmd.AddCommand(xcmd.NewRandomCommand())
	rootCmd.AddCommand(xcmd.NewSeqCommand())
	rootCmd.AddCommand(xcmd.NewRangeCommand())
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
