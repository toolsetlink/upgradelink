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

package i18n

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslator(t *testing.T) {
	c := Conf{
		Dir: "",
	}

	trans := NewTranslator(c, LocaleFS)

	res := trans.Trans(context.WithValue(context.Background(), "lang", "zh"), "common.success")
	assert.Equal(t, "成功", res)
	fmt.Println("res: ", res)

	res = trans.Trans(context.WithValue(context.Background(), "lang", "ja"), "common.success")
	assert.Equal(t, "操作が成功しました", res)
	fmt.Println("res: ", res)

	res = trans.Trans(context.WithValue(context.Background(), "lang", "en"), "common.success")
	assert.Equal(t, "Successfully", res)
	fmt.Println("res: ", res)

}
