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

package companyctx

import (
	"context"
	"slices"
	"testing"
)

func TestGetRoleIDFromCtx(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "test role empty ctx",
			args:    args{ctx: context.Background()},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test role ctx",
			args:    args{ctx: context.WithValue(context.Background(), "companyId", 1)},
			want:    []string{"001", "002"},
			wantErr: false,
		},
		//{
		//	name: "test meta context",
		//	args: args{ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
		//		enum.RoleIdRpcCtxKey: "001,002",
		//	}))},
		//	want:    []string{"001", "002"},
		//	wantErr: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCompanyIDFromCtx(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCompanyIDFromCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("GetCompanyIDFromCtx() got = %v, want %v", got, tt.want)
			}
		})
	}
}
