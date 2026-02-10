// Copyright 2023 The Ryan SU Authors (https://github.com/suyuan32). All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pointy

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetPointer(t *testing.T) {
	var dataIntOrigin int = 1
	dataInt := GetPointer(dataIntOrigin)
	assert.Equal(t, dataIntOrigin, *dataInt)

	var dataUintOrigin uint = 1
	dataUint := GetPointer(dataUintOrigin)
	assert.Equal(t, dataUintOrigin, *dataUint)

	var dataStringOrigin = "123"
	dataString := GetPointer(dataStringOrigin)
	assert.Equal(t, dataStringOrigin, *dataString)
}

func TestGetSlicePointer(t *testing.T) {
	nums := []int{1, 2, 3}
	numsPointer := GetSlicePointer(nums)
	assert.Equal(t, nums[0], *numsPointer[0])
	assert.Equal(t, nums[1], *numsPointer[1])
	assert.Equal(t, nums[2], *numsPointer[2])

	strs := []string{"a", "b", "c"}
	strsPointer := GetSlicePointer(strs)
	assert.Equal(t, strs[0], *strsPointer[0])
	assert.Equal(t, strs[1], *strsPointer[1])
	assert.Equal(t, strs[2], *strsPointer[2])
}

func TestGetUnixMilliPointer(t *testing.T) {
	var zeroTime time.Time
	testTime := GetUnixMilliPointer(zeroTime.UnixMilli())
	fmt.Println(testTime)
	if testTime != nil {
		t.Error("TestGetUnixMilliPointer: convert failed")
	}
}

func TestGetTimePointer(t *testing.T) {
	nowTime := time.Now().Unix()

	timestampNow := GetTimePointer(&nowTime, 0)

	assert.Equal(t, nowTime, timestampNow.Unix())
}

func TestGetTimeMilliPointer(t *testing.T) {
	nowTime := time.Now().UnixMilli()

	timestampNow := GetTimeMilliPointer(&nowTime)

	assert.Equal(t, nowTime, timestampNow.UnixMilli())
}
