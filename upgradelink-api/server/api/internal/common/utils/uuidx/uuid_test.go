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
	"reflect"
	"testing"

	"github.com/gofrs/uuid/v5"
)

func TestParseUUIDSlice(t *testing.T) {
	type args struct {
		ids []string
	}
	tests := []struct {
		name string
		args args
		want []uuid.UUID
	}{
		{
			name: "test1",
			args: args{ids: []string{"123"}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseUUIDSlice(tt.args.ids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUUIDSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseUUIDString(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want uuid.UUID
	}{
		{
			name: "test1",
			args: args{id: "123456"},
			want: uuid.UUID{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseUUIDString(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUUIDString() = %v, want %v", got, tt.want)
			}
		})
	}
}
