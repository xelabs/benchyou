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
	"github.com/XeLabs/go-mysqlstack/common"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandInt64(min int64, max int64) int64 {
	return min + int64(rand.Int63n(int64(max-min)))
}

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
