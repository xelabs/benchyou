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
	MemFree  uint64
	MemCache uint64
	SwapSi   uint64
	SwapSo   uint64
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
21  0      0   5621    155   4353    0    0     6     4    0    0  0  0  0  0 100
23  0      0   5607    155   4364    0    0     0  7456 81544 95061  0  0  0  0  0
*/
func (v *VMS) parse(outs string) (err error) {
	lines := strings.Split(outs, "\n")
	l := lines[3]
	cols := splitColumns(l)

	if v.Stat.MemFree, err = strconv.ParseUint(cols[3], 10, 64); err != nil {
		return
	}

	if v.Stat.MemCache, err = strconv.ParseUint(cols[5], 10, 64); err != nil {
		return
	}

	if v.Stat.SwapSi, err = strconv.ParseUint(cols[6], 10, 64); err != nil {
		return
	}

	if v.Stat.SwapSo, err = strconv.ParseUint(cols[7], 10, 64); err != nil {
		return
	}

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
			log.Printf("vm timer startt\n")
			if err := v.fetch(); err != nil {
				log.Printf("vmstat.fetch.error[%v]\n", err)
			}
			log.Printf("vm timer end\n")
		}
	}()
}

func (v *VMS) Stop() {
	v.t.Stop()
}
