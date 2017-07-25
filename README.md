[![Build Status](https://travis-ci.org/XeLabs/benchyou.svg?branch=master)](https://travis-ci.org/XeLabs/benchyou)
## About

benchyou is a benchmark tool for MySQL, similar Sysbench.

In addition to real-time monitoring TPS, she also monitors vmstat/iostat via SSH tunnel.

The idea of stat per operation is inspired by Mark Callaghan, [Small Datum](http://smalldatum.blogspot.com)

## Screenshots
```
time            thds             tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[10s]        [r:4,w:4,u:4,d:4]  5372     3842    1530    0      0.00     0      0.00      0.00    0.00      0.00    0.00     0.00    0       0         3.14       2.61         45652

time            thds             tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[11s]        [r:4,w:4,u:4,d:4]  5325     3827    1498    0      0.00     0      0.00      0.00    0.00      0.00    0.00     0.00    0       0         3.12       2.66         50977

time            thds             tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[12s]        [r:4,w:4,u:4,d:4]  5342     3832    1510    0      0.00     0      0.00      0.00    0.00      0.00    0.00     0.00    0       0         3.11       2.65         56319

time            thds             tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[13s]        [r:4,w:4,u:4,d:4]  5335     3842    1493    0      0.00     0      0.00      0.00    0.00      0.00    0.00     0.00    0       0         3.13       2.68         61654

time            thds             tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[14s]        [r:4,w:4,u:4,d:4]  5321     3840    1481    0      0.00     0      0.00      0.00    0.00      0.00    0.00     0.00    0       0         3.09       2.70         66975

time            thds             tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[15s]        [r:4,w:4,u:4,d:4]  5365     3865    1500    0      0.00     0      0.00      0.00    0.00      0.00    0.00     0.00    0       0         3.10       2.67         72340

time            thds             tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[16s]        [r:4,w:4,u:4,d:4]  5352     3843    1509    0      0.00     0      0.00      0.00    0.00      0.00    0.00     0.00    0       0         3.13       2.65         77692

time            thds             tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[17s]        [r:4,w:4,u:4,d:4]  5330     3816    1514    0      0.00     0      0.00      0.00    0.00      0.00    0.00     0.00    0       0         3.14       2.64         83022

----------------------------------------------------------------------------------------------avg---------------------------------------------------------------------------------------------
time          tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op            w-rsp(ms)                       r-rsp(ms)              total-number
[17s]        5026     3601    1424    0      0.00     0      0.00      0.00    0.00      0.00    0.00     0.00    [avg:0.18,min:0.00,max:165.72]  [avg:0.16,min:0.00,max:24.81]      85455
```

the columns:
```
time:         benchmark uptime
thds:         read threads and write(insert/update/delete) threads
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
      --read-threads int            number of read threads to use(Default 32) (default 32)
      --write-threads int           number of write threads to use(Default 32) (default 32)
      --update-threads int          number of update threads to use(Default 0)
      --delete-threads int          number of delete threads to use(Default 0)
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
      --rows-per-insert int         #rows per insert(Default 1) (default 1)
      --batch-per-commit int        #rows per transaction(Default 1) (default 1)
      --ssh-host string             SSH server host(Default NULL, same as mysql-host)
      --ssh-password string         SSH server password(Default benchyou) (default "benchyou")
      --ssh-port int                SSH server port(Default 22) (default 22)
      --ssh-user string             SSH server user(Default benchyou) (default "benchyou")
```

## Examples

```
prepare 64 tables:
./bin/benchyou  --mysql-host=192.168.0.3 --mysql-user=benchyou --mysql-password=benchyou  --oltp-tables-count=64 prepare

cleanup 64 tables:
./bin/benchyou  --mysql-host=192.168.0.3 --mysql-user=benchyou --mysql-password=benchyou  --oltp-tables-count=64 cleanup

random insert(Write/Read Ratio=128:8):
 ./bin/benchyou  --mysql-host=192.168.0.3 --mysql-user=benchyou --mysql-password=benchyou --ssh-user=benchyou --ssh-password=benchyou --oltp-tables-count=64 --write-threads=128 --read-threads=8 --max-time=3600 random

sequential insert(Write/Read Ratio=128:8):
 ./bin/benchyou  --mysql-host=192.168.0.3 --mysql-user=benchyou --mysql-password=benchyou --ssh-user=benchyou --ssh-password=benchyou --oltp-tables-count=64 --write-threads=128 --read-threads=8 --max-time=3600 seq

mix(Write/Read/Update/Delete Ratio=4:4:4:4):
 ./bin/benchyou  --mysql-host=192.168.0.3 --mysql-user=benchyou --mysql-password=benchyou --ssh-user=benchyou --ssh-password=benchyou --oltp-tables-count=64 --write-threads=4 --read-threads=4 --update-threads=4 --delete-threads=4 --max-time=3600 random

insert multiple rows(10 rows per insert):
 ./bin/benchyou  --mysql-host=192.168.0.3 --mysql-user=benchyou --mysql-password=benchyou --ssh-user=benchyou --ssh-password=benchyou --oltp-tables-count=64 --write-threads=4 --rows-per-insert=10 --max-time=3600 random

batch update(10 rows per transaction):
 ./bin/benchyou  --mysql-host=192.168.0.3 --mysql-user=benchyou --mysql-password=benchyou --ssh-user=benchyou --ssh-password=benchyou --oltp-tables-count=64 --update-threads=4 --batch-per-commit=10 --max-time=3600 random

query-range(Write/Read Ratio=128:8):
 ./bin/benchyou  --mysql-host=192.168.0.3 --mysql-user=benchyou --mysql-password=benchyou --ssh-user=benchyou --ssh-password=benchyou --oltp-tables-count=64 --write-threads=128 --read-threads=8 --max-time=3600 --mysql-range-order=DESC range
```
