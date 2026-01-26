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

package jwt

import "testing"

func TestNewJwtToken(t *testing.T) {
	type args struct {
		secretKey string
		iat       int64
		seconds   int64
		opt       []Option
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				secretKey: "jS6VKDtsJf3z1n2VKDtsJf3z1n2",
				iat:       1000,
				seconds:   10,
				opt: []Option{
					WithOption("userId", "abc"),
					WithOption("roleId", 1),
				},
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjEwMTAsImlhdCI6MTAwMCwicm9sZUlkIjoxLCJ1c2VySWQiOiJhYmMifQ.wWLNg-FLr0d6encPe0p8Dw17orN89oLK_KAV0VDVBLk",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewJwtToken(tt.args.secretKey, tt.args.iat, tt.args.seconds, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewJwtToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewJwtToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
