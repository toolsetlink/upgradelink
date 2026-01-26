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

func TestStripBearerPrefixFromToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{token: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQyMjExMzksImlhdCI6MTY5Mzk2MTkzOSwicm9sZUlkIjoiMDAxIiwidXNlcklkIjoiMDE4YTY3ZmYtODgxOS03NmI4LWE5MzAtMWIxMzRjZjFjMWFmIn0.wURcTLPsO1EO-ok_-Uv2o0t3qq6iEDWbc7v7WmL--HQ"},
			want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQyMjExMzksImlhdCI6MTY5Mzk2MTkzOSwicm9sZUlkIjoiMDAxIiwidXNlcklkIjoiMDE4YTY3ZmYtODgxOS03NmI4LWE5MzAtMWIxMzRjZjFjMWFmIn0.wURcTLPsO1EO-ok_-Uv2o0t3qq6iEDWbc7v7WmL--HQ",
		},
		{
			name: "test2",
			args: args{token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQyMjExMzksImlhdCI6MTY5Mzk2MTkzOSwicm9sZUlkIjoiMDAxIiwidXNlcklkIjoiMDE4YTY3ZmYtODgxOS03NmI4LWE5MzAtMWIxMzRjZjFjMWFmIn0.wURcTLPsO1EO-ok_-Uv2o0t3qq6iEDWbc7v7WmL--HQ"},
			want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQyMjExMzksImlhdCI6MTY5Mzk2MTkzOSwicm9sZUlkIjoiMDAxIiwidXNlcklkIjoiMDE4YTY3ZmYtODgxOS03NmI4LWE5MzAtMWIxMzRjZjFjMWFmIn0.wURcTLPsO1EO-ok_-Uv2o0t3qq6iEDWbc7v7WmL--HQ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripBearerPrefixFromToken(tt.args.token); got != tt.want {
				t.Errorf("StripBearerPrefixFromToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
