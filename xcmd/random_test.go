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

func TestXcmdRandom(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	cmd := NewRandomCommand()
	MockInitFlags(cmd, mysql.Addr())
	randomCommandFn(cmd, nil)
}
