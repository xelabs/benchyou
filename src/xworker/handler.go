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
}

type QueryHandler interface {
	Run()
	Stop()
}
