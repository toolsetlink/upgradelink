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

package uuidx

import (
	"github.com/gofrs/uuid/v5"
	"github.com/zeromicro/go-zero/core/logx"
)

// NewUUID returns a new UUID.
func NewUUID() uuid.UUID {
	id, err := uuid.NewV7()
	if err != nil {
		logx.Errorw("fail to generate UUID", logx.Field("detail", err))
		return uuid.UUID{}
	}
	return id
}

// ParseUUIDSlice parses the UUID string slice to UUID slice.
func ParseUUIDSlice(ids []string) []uuid.UUID {
	var result []uuid.UUID
	for _, v := range ids {
		p, err := uuid.FromString(v)
		if err != nil {
			logx.Errorw("fail to parse string to UUID", logx.Field("detail", err))
			return nil
		}
		result = append(result, p)
	}
	return result
}

// ParseUUIDString parses UUID string to UUID type.
func ParseUUIDString(id string) uuid.UUID {
	result, err := uuid.FromString(id)
	if err != nil {
		logx.Errorw("fail to parse string to UUID", logx.Field("detail", err))
		return uuid.UUID{}
	}
	return result
}

// ParseUUIDSliceToPointer parses the UUID string slice to UUID pointer slice.
func ParseUUIDSliceToPointer(ids []string) []*uuid.UUID {
	var result []*uuid.UUID
	for _, v := range ids {
		p, err := uuid.FromString(v)
		if err != nil {
			logx.Errorw("fail to parse string to UUID", logx.Field("detail", err))
			return nil
		}
		result = append(result, &p)
	}
	return result
}

// ParseUUIDStringToPointer parses UUID string to UUID pointer.
func ParseUUIDStringToPointer(id *string) *uuid.UUID {
	if id == nil {
		return nil
	}

	result, err := uuid.FromString(*id)
	if err != nil {
		logx.Errorw("fail to parse string to UUID", logx.Field("detail", err))
		return nil
	}
	return &result
}
