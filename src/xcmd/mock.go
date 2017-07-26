package xcmd

import (
	"net"
	"strconv"

	"github.com/spf13/cobra"
)

func MockInitFlags(cmd *cobra.Command, addr string) {
	host, sport, _ := net.SplitHostPort(addr)
	port, _ := strconv.Atoi(sport)

	cmd.Flags().Int("write-threads", 32, "")
	cmd.Flags().Int("read-threads", 32, "number of read threads to use(Default 32)")
	cmd.Flags().Int("update-threads", 0, "number of update threads to use(Default 0)")
	cmd.Flags().Int("delete-threads", 0, "number of delete threads to use(Default 0)")
	cmd.Flags().String("mysql-host", host, "MySQL server host(Default NULL)")
	cmd.Flags().Int("mysql-port", port, "MySQL server port(Default 3306)")
	cmd.Flags().String("mysql-user", "mock", "MySQL user(Default benchyou)")
	cmd.Flags().String("mysql-password", "", "MySQL password(Default benchyou)")
	cmd.Flags().String("mysql-db", "sbtest", "MySQL database name(Default sbtest)")
	cmd.Flags().String("mysql-table-engine", "tokudb", "storage engine to use for the test table {tokudb,innodb,...}(Default tokudb)")
	cmd.Flags().String("mysql-range-order", "ASC", "range query sort the result-set in {ASC|DESC} (Default ASC)")
	cmd.Flags().Int("mysql-enable-xa", 0, "enable MySQL xa transaction for insertion {0|1} (Default 0)")
	cmd.Flags().Int("rows-per-insert", 1, "#rows per insert(Default 1)")
	cmd.Flags().Int("batch-per-commit", 1, "#rows per transaction(Default 1)")
	cmd.Flags().Int("max-time", 3600, "limit for total execution time in seconds(Default 3600)")
	cmd.Flags().Uint64("max-request", 10000, "limit for total requests, including write and read(Default 0, means no limits)")
	cmd.Flags().Int("oltp-tables-count", 8, "number of tables to create(Default 8)")
	cmd.Flags().String("ssh-host", "", "SSH server host(Default NULL, same as mysql-host)")
	cmd.Flags().String("ssh-user", "benchyou", "SSH server user(Default benchyou)")
	cmd.Flags().String("ssh-password", "benchyou", "SSH server password(Default benchyou)")
	cmd.Flags().Int("ssh-port", 22, "SSH server port(Default 22)")
}
