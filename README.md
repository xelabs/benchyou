[![Build Status](https://travis-ci.org/XeLabs/benchyou.png)](https://travis-ci.org/XeLabs/benchyou)
## About

benchyou is a benchmark tool for MySQL, similar Sysbench.

In addition to real-time monitoring TPS, she also monitors vmstat/iostat via SSH tunnel(need to install sshpass).

benchyou supports two modes for benchmark: sysbench(defaults) and iibench(--bench-mode=iibench).

The idea of stat per operation is inspired by Mark Callaghan, [Small Datum](http://smalldatum.blogspot.com)

## Screenshots
```
... ...
time            thds       tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op
[59s]        [r:8,w:128]  11244    10357   887     0      0.00     177    0.02      0.00    0.00      24.58   2.24     8.89

time            thds       tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op
[60s]        [r:8,w:128]  11482    10638   844     0      0.00     177    0.02      0.00    0.00      24.58   2.19     8.70

-----------------------------------------------------------------------------------avg---------------------------------------------------------------------------------------------
time          tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op            w-rsp(ms)                       r-rsp(ms)
[60s]        11233    10397   835     0      0.00     2      0.00      0.00    0.00      0.41    0.00     0.00    [avg:0.20,min:0.00,max:493.46]  [avg:0.16,min:0.00,max:360.64]

```

the columns:
```
time:    benchmark uptime
thds:    read threads and write threads
tps:     transaction per second, including write and read
wtps:    write tps
rtps:    read tps
rio:     read io numbers per second
rio/op:  rio per operation
wio:     write io numbers per second
wio/op:  wio per operation
rMB:     amount data read from the device(megabytes) per second
rKB/op:  rKB per operation
wMB:     amount data written to the device(megabytes) per second
wKB/op:  wKB per operation
cpu/op:  CPU usecs per operation, measured by vmstat
```

## Build

```
$git clone https://github.com/XeLabs/benchyou
$cd benchyou
$make build
$./bin/benchyou -h
```

## Required

ubuntu:
```
$sudo apt-get install sshpass
```

## Usage

```
$ ./bin/benchyou -h
Usage:
  benchyou [command]

Available Commands:
  prepare
  cleanup
  random
  seq
  range

Flags:
      --bench-mode string           benchmark mode, {sysbench|iibench}(Default sysbench) (default "sysbench")
      --max-time int                limit for total execution time in seconds(Default 3600) (default 3600)
      --mysql-db string             MySQL database name(Default sbtest) (default "sbtest")
      --mysql-host string           MySQL server host(Default NULL)
      --mysql-password string       MySQL password(Default benchyou) (default "benchyou")
      --mysql-port int              MySQL server port(Default 3306) (default 3306)
      --mysql-range-order string    range query sort the result-set in {ASC|DESC} (Default ASC) (default "ASC")
      --mysql-table-engine string   storage engine to use for the test table {tokudb,innodb,...}(Default tokudb) (default "tokudb")
      --mysql-user string           MySQL user(Default benchyou) (default "benchyou")
      --oltp-tables-count int       number of tables to create(Default 8) (default 8)
      --read-threads int            number of read threads to use(Default 32) (default 32)
      --rows-per-commit int         #rows per transaction(Default 1) (default 1)
      --ssh-host string             SSH server host(Default NULL, same as mysql-host)
      --ssh-password string         SSH server password(Default benchyou) (default "benchyou")
      --ssh-port int                SSH server port(Default 22) (default 22)
      --ssh-user string             SSH server user(Default benchyou) (default "benchyou")
      --write-threads int           number of write threads to use(Default 32) (default 32)
```

## Examples

sysbench:
```
prepare 64 tables:
./bin/benchyou  --mysql-host=192.168.0.3 --mysql-user=benchyou --mysql-password=benchyou --mysql-user=benchyou --mysql-password=benchyou  --oltp-tables-count=64 prepare

cleanup 64 tables:
./bin/benchyou  --mysql-host=192.168.0.3 --mysql-user=benchyou --mysql-password=benchyou --mysql-user=benchyou --mysql-password=benchyou  --oltp-tables-count=64 cleanup

random(Write/Read Ratio=128:8):
 ./bin/benchyou  --mysql-host=192.168.0.3 --mysql-user=benchyou --mysql-password=benchyou --ssh-user=benchyou --ssh-password=benchyou --oltp-tables-count=64 --write-threads=128 --read-threads=8 --max-time=3600 random

sequential(Write/Read Ratio=128:8):
 ./bin/benchyou  --mysql-host=192.168.0.3 --mysql-user=benchyou --mysql-password=benchyou --ssh-user=benchyou --ssh-password=benchyou --oltp-tables-count=64 --write-threads=128 --read-threads=8 --max-time=3600 seq

query-range(Write/Read Ratio=128:8):
 ./bin/benchyou  --mysql-host=192.168.0.3 --mysql-user=benchyou --mysql-password=benchyou --ssh-user=benchyou --ssh-password=benchyou --oltp-tables-count=64 --write-threads=128 --read-threads=8 --max-time=3600 --mysql-range-order=DESC range

```

iibench:
```
... --bench-mode=iibench
```
