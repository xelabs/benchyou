/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xstat

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"xcommon"
)

type VMStat struct {
	SystemCS uint64
	IdleCPU  uint64
}

type VMS struct {
	conf *xcommon.Conf
	cmd  string
	Stat *VMStat
	All  *VMStat
	t    *time.Ticker
}

func NewVMS(conf *xcommon.Conf) *VMS {
	return &VMS{
		conf: conf,
		cmd:  "vmstat -SM 1 2",
		Stat: &VMStat{},
		All:  &VMStat{},
		t:    time.NewTicker(time.Second),
	}
}

func (v *VMS) args() []string {
	args := fmt.Sprintf("sshpass -p %s ssh -o 'StrictHostKeyChecking=no' %s@%s -p %d \"%s\"",
		v.conf.Ssh_password,
		v.conf.Ssh_user,
		v.conf.Ssh_host,
		v.conf.Ssh_port,
		v.cmd)

	return []string{
		"-c",
		args,
	}
}

/*
procs -----------memory---------- ---swap-- -----io---- -system-- ------cpu-----
 r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st
 0  0      0   1737    233   9408    0    0     2     1    0    0  0  0  0  0 100
 0  0      0   1737    233   9408    0    0     0     0  326  460  0  0  0  0  0
*/
func (v *VMS) parse(outs string) (err error) {
	lines := strings.Split(outs, "\n")
	l := lines[3]
	cols := splitColumns(l)

	if v.Stat.SystemCS, err = strconv.ParseUint(cols[11], 10, 64); err != nil {
		return
	}
	v.All.SystemCS += v.Stat.SystemCS

	if v.Stat.IdleCPU, err = strconv.ParseUint(cols[14], 10, 64); err != nil {
		return
	}
	v.All.IdleCPU += v.Stat.IdleCPU

	return
}

func (v *VMS) fetch() error {
	outs, err := xcommon.RunCommand("bash", v.args())
	if err != nil {
		return err
	}

	return v.parse(string(outs))
}

func (v *VMS) Start() {
	go func() {
		for _ = range v.t.C {
			if err := v.fetch(); err != nil {
				log.Printf("vmstat.fetch.error[%v]\n", err)
			}
		}
	}()
}

func (v *VMS) Stop() {
	v.t.Stop()
}
