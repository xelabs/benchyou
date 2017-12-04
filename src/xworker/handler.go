/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xworker

// Handler interface.
type Handler interface {
	Run()
	Stop()
	Rows() uint64
}
