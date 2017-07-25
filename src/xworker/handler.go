/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xworker

type InsertHandler interface {
	Run()
	Stop()
	Rows() uint64
}

type QueryHandler interface {
	Run()
	Stop()
	Rows() uint64
}

type DeleteHandler interface {
	Run()
	Stop()
	Rows() uint64
}

type UpdateHandler interface {
	Run()
	Stop()
	Rows() uint64
}
