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
	"math/rand"
	"time"

	"github.com/xelabs/go-mysqlstack/sqlparser/depends/common"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// RandInt64 rands the int65 between min and max.
func RandInt64(min int64, max int64) int64 {
	return min + int64(rand.Int63n(int64(max-min)))
}

// RandString rands the strings format by template.
func RandString(template string) string {
	nums := "0123456789"
	alpha := "abcdefghijklmnopqrstuvwxyz"
	nLen := len(nums)
	aLen := len(alpha)

	buf := common.NewBuffer(128)
	for i := 0; i < len(template); i++ {
		if template[i] == '#' {
			buf.WriteU8(nums[rand.Int31n(int32(nLen))])
		} else if template[i] == '@' {
			buf.WriteU8(alpha[rand.Int31n(int32(aLen))])
		} else {
			buf.WriteU8(template[i])
		}
	}
	return common.BytesToString(buf.Datas())
}
