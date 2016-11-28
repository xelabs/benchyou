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
	"strings"
)

func splitColumns(line string) []string {
	cols := make([]string, 0)
	for _, f := range strings.Split(line, " ") {
		if len(f) > 0 {
			cols = append(cols, f)
		}
	}
	return cols
}
