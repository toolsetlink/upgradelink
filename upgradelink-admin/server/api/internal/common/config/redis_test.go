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

package config

import (
	"context"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	t.Skip("skip test")

	// Test must redis
	conf := &RedisConf{
		Host:     "localhost:6379",
		Db:       0,
		Username: "",
		Pass:     "",
		Tls:      false,
	}

	rds := conf.MustNewUniversalRedis()

	rds.Set(context.Background(), "testKeyDb0", "testVal", 2*time.Minute)

	conf1 := &RedisConf{
		Host:     "localhost:6379",
		Db:       1,
		Username: "",
		Pass:     "",
		Tls:      false,
	}

	rds1 := conf1.MustNewUniversalRedis()

	rds1.Set(context.Background(), "testKeyDb1", "testVal", 2*time.Minute)

	conf2 := &RedisConf{
		Host:     "localhost:6379,localhost:6380",
		Db:       1,
		Username: "",
		Pass:     "",
		Tls:      false,
	}

	rds2 := conf2.MustNewUniversalRedis()

	rds2.Set(context.Background(), "testCluster", "testCluster", 2*time.Minute)
}
