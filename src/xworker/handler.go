/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xworker

type Handler interface {
	Run()
	Stop()
	Rows() uint64
}
