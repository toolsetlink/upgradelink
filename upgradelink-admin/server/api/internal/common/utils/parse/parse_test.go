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

package parse

import (
	"reflect"
	"testing"

	"golang.org/x/text/language"
)

func TestParseTags(t *testing.T) {
	type args struct {
		lang string
	}
	tests := []struct {
		name string
		args args
		want []language.Tag
	}{
		{
			name: "testChinese",
			args: args{lang: "zh"},
			want: []language.Tag{language.Chinese},
		},
		{
			name: "testEn",
			args: args{lang: "en"},
			want: []language.Tag{language.English},
		},
		{
			name: "testWrong",
			args: args{lang: "one two"},
			want: []language.Tag{language.Chinese},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseTags(tt.args.lang); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
