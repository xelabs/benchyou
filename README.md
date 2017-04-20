[![Build Status](https://travis-ci.org/XeLabs/benchyou.svg?branch=master)](https://travis-ci.org/XeLabs/benchyou)
## About

benchyou is a benchmark tool for MySQL, similar Sysbench.

In addition to real-time monitoring TPS, she also monitors vmstat/iostat via SSH tunnel.

benchyou supports two modes for benchmark: sysbench(defaults) and iibench(--bench-mode=iibench).

The idea of stat per operation is inspired by Mark Callaghan, [Small Datum](http://smalldatum.blogspot.com)

## Screenshots
```
time            thds       tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[1s]         [r:32,w:32]  51037    4020    47017   0      0.00     0      0.00      0.00    0.00      0.00    0.00     0.00    0       0         7.26       0.63         51037

time            thds       tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[2s]         [r:32,w:32]  49047    4576    44471   0      0.00     0      0.00      0.00    0.00      0.00    0.00     0.00    0       0         6.79       0.72         100084

time            thds       tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[3s]         [r:32,w:32]  57432    5157    52275   0      0.00     2716   0.05      0.00    0.00      23.59   0.42     2.47    3092    3970      6.08       0.61         157516

time            thds       tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[4s]         [r:32,w:32]  49988    3919    46069   0      0.00     2886   0.06      0.00    0.00      24.89   0.51     3.26    3082    3973      8.01       0.69         207504

time            thds       tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[5s]         [r:32,w:32]  46725    3518    43207   0      0.00     2370   0.05      0.00    0.00      20.39   0.45     2.94    3077    3977      8.98       0.74         254229

time            thds       tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[6s]         [r:32,w:32]  57784    4582    53202   0      0.00     1965   0.03      0.00    0.00      17.10   0.30     2.07    3071    3980      6.85       0.60         312013

time            thds       tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[7s]         [r:32,w:32]  56314    4754    51560   0      0.00     2765   0.05      0.00    0.00      24.09   0.44     2.86    3066    3983      6.53       0.62         368327

time            thds       tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[8s]         [r:32,w:32]  61745    5163    56582   0      0.00     2741   0.04      0.00    0.00      23.66   0.39     2.51    3059    3988      6.01       0.56         430072

time            thds       tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[9s]         [r:32,w:32]  63196    4961    58235   82     0.00     2994   0.05      2.70    0.04      26.46   0.43     2.93    3050    3992      6.30       0.55         493268

time            thds       tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[10s]        [r:32,w:32]  59810    4531    55279   160    0.00     2795   0.05      5.47    0.09      24.45   0.42     3.04    3031    3996      6.90       0.58         553078

----------------------------------------------------------------------------------------------avg---------------------------------------------------------------------------------------------
time          tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op            w-rsp(ms)                      r-rsp(ms)              total-number
[10s]        55680    4552    51128   16     0.00     279    0.00      0.55    0.00      2.44    0.00     0.03    [avg:0.69,min:0.00,max:67.53]  [avg:0.06,min:0.00,max:74.82]      556805

```

the columns:
```
time:         benchmark uptime
thds:         read threads and write threads
tps:          transaction per second, including write and read
wtps:         write tps
rtps:         read tps
rio:          read io numbers per second
rio/op:       rio per operation
wio:          write io numbers per second
wio/op:       wio per operation
rMB:          amount data read from the device(megabytes) per second
rKB/op:       rKB per operation
wMB:          amount data written to the device(megabytes) per second
wKB/op:       wKB per operation
cpu/op:       CPU usecs per operation, measured by vmstat
freeMB:       the amount of idle memory(megabytes)
cacheMB:      the amount of memory(megabytes) used as cache
w-rsp:        the response time of one write operation,  in millisecond
r-rsp:        the response time of one read  operation,  in millisecond
total-number: the total number events
```

## Build

```
$git clone https://github.com/XeLabs/benchyou
$cd benchyou
$make build
$./bin/benchyou -h
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
      --max-request uint            limit for total requests, including write and read(Default 0, means no limits)
      --max-time int                limit for total execution time in seconds(Default 3600) (default 3600)
      --mysql-db string             MySQL database name(Default sbtest) (default "sbtest")
      --mysql-enable-xa int         enable MySQL xa transaction for insertion {0|1} (Default 0)
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
./bin/benchyou  --mysql-host=192.168.0.3 --mysql-user=benchyou --mysql-password=benchyou  --oltp-tables-count=64 prepare

cleanup 64 tables:
./bin/benchyou  --mysql-host=192.168.0.3 --mysql-user=benchyou --mysql-password=benchyou  --oltp-tables-count=64 cleanup

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
