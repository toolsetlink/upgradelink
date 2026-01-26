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

package encrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		origin string
	}{
		{
			origin: "123456",
		},
		{
			origin: "123456789..",
		},
	}

	for _, v := range tests {
		// test encrypt
		encryptedData := BcryptEncrypt(v.origin)
		result := BcryptCheck(v.origin, encryptedData)
		assert.Equal(t, result, true)
	}
}

func TestBcryptCheck(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				password: "simple-admin",
				hash:     "$2a$10$RGY8FVLUSKNMdKQr/y2oi.kh4r/ns6hbpJc.0RP56jd3gazeOJa42",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, BcryptCheck(tt.args.password, tt.args.hash), "BcryptCheck(%v, %v)", tt.args.password, tt.args.hash)
		})
	}
}
