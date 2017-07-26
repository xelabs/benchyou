/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xcommon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXCommonRandomInt(t *testing.T) {
	min := int64(100)
	max := int64(111)
	r := RandInt64(min, max)
	assert.True(t, min <= r)
	assert.True(t, r <= max)
}

func TestXCommonRandomString(t *testing.T) {
	r := RandString(Ctemplate)
	assert.Equal(t, len(Ctemplate), len(r))
}
