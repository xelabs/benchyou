[![Build Status](https://travis-ci.org/xelabs/benchyou.svg?branch=master)](https://travis-ci.org/xelabs/benchyou) [![Go Report Card](https://goreportcard.com/badge/github.com/xelabs/benchyou)](https://goreportcard.com/report/github.com/xelabs/benchyou)  [![codecov.io](https://codecov.io/gh/xelabs/benchyou/graphs/badge.svg)](https://codecov.io/gh/xelabs/benchyou/branch/master)

## benchyou

benchyou is a benchmark tool for MySQL, similar Sysbench.

In addition to real-time monitoring TPS, she also monitors vmstat/iostat via SSH tunnel.

The idea of stat per operation is inspired by Mark Callaghan, [Small Datum](http://smalldatum.blogspot.com)

## Screenshots
```
time            thds               tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[1s]         [r:4,w:128,u:4,d:4]  33508    24056   9452    0      0.00     0      0.00      0.00    0.00      0.00    0.00     0.00    0       0         5.05       0.39         33508

time            thds               tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[2s]         [r:4,w:128,u:4,d:4]  29929    21287   8642    0      0.00     0      0.00      0.00    0.00      0.00    0.00     0.00    0       0         6.12       0.45         63437

time            thds               tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[3s]         [r:4,w:128,u:4,d:4]  27967    20215   7752    0      0.00     2472   0.09      0.00    0.00      25.57   0.94     6.22    6185    4570      6.51       0.51         91404

time            thds               tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[4s]         [r:4,w:128,u:4,d:4]  30072    21560   8512    0      0.00     2235   0.07      0.00    0.00      23.55   0.80     5.60    6174    4577      5.74       0.45         121476

time            thds               tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[5s]         [r:4,w:128,u:4,d:4]  32182    23609   8573    0      0.00     2810   0.09      0.00    0.00      29.55   0.94     5.91    6165    4584      5.45       0.46         153658

time            thds               tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[6s]         [r:4,w:128,u:4,d:4]  34548    24771   9777    0      0.00     2823   0.08      0.00    0.00      29.28   0.87     5.80    6156    4590      5.14       0.40         188206

time            thds               tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[7s]         [r:4,w:128,u:4,d:4]  35185    24844   10341   0      0.00     2553   0.07      0.00    0.00      26.40   0.77     5.74    6145    4596      5.20       0.38         223391

time            thds               tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[8s]         [r:4,w:128,u:4,d:4]  36266    26030   10236   0      0.00     2880   0.08      0.00    0.00      29.84   0.84     5.86    6137    4603      4.95       0.38         259657

time            thds               tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[9s]         [r:4,w:128,u:4,d:4]  37414    26834   10580   0      0.00     3234   0.09      0.00    0.00      34.07   0.93     6.06    6125    4611      4.81       0.37         297071

time            thds               tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op  freeMB  cacheMB   w-rsp(ms)  r-rsp(ms)    total-number
[10s]        [r:4,w:128,u:4,d:4]  36158    25845   10313   0      0.00     3329   0.09      0.00    0.00      35.53   1.01     6.43    6113    4619      4.98       0.38         333229

----------------------------------------------------------------------------------------------avg---------------------------------------------------------------------------------------------
time          tps     wtps    rtps    rio    rio/op   wio    wio/op    rMB     rKB/op    wMB     wKB/op   cpu/op            w-rsp(ms)                       r-rsp(ms)              total-number
[10s]        33642    24132   9509    0      0.00     332    0.00      0.00    0.00      3.55    0.01     0.07    [avg:0.53,min:0.00,max:149.79]  [avg:0.04,min:0.00,max:27.63]      336420
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
$git clone https://github.com/xelabs/benchyou
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
      --mysql-table-engine string   storage engine to use for the test table {tokudb,innodb,...}(Default innodb) (default "innodb")
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

## License

benchyou is released under the GPLv3. See LICENSE
