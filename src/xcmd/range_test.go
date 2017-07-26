/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xcmd

import (
	"testing"
	"xcommon"
)

func TestXcmdRange(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	cmd := NewRangeCommand()
	MockInitFlags(cmd, mysql.Addr())
	rangeCommandFn(cmd, nil)
}
