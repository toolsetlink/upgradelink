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
	"crypto/tls"
	"errors"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"
	"github.com/zeromicro/go-zero/core/logx"
)

// A RedisConf is a redis config.
type RedisConf struct {
	Host         string `json:",env=REDIS_HOST"`
	Db           int    `json:",default=0,env=REDIS_DB"`
	Username     string `json:",optional,env=REDIS_USERNAME"`
	Pass         string `json:",optional,env=REDIS_PASSWORD"`
	Tls          bool   `json:",optional,env=REDIS_TLS"`
	Master       string `json:",optional,env=REDIS_MASTER"`
	PoolSize     int    `json:",optional,default=48,env=REDIS_POOL_SIZE"`
	MaxIdleConns int    `json:",optional,default=12,env=REDIS_MAX_IDLE_CONNS"`
}

func (r RedisConf) Validate() error {
	if len(r.Host) == 0 {
		return errors.New("host cannot be empty")
	}
	return nil
}

func (r RedisConf) NewUniversalRedis() (redis.UniversalClient, error) {
	err := r.Validate()
	if err != nil {
		return nil, err
	}

	opt := &redis.UniversalOptions{
		Addrs:        strings.Split(r.Host, ","),
		DB:           r.Db,
		Password:     r.Pass,
		Username:     r.Username,
		PoolSize:     r.PoolSize,
		MaxIdleConns: r.MaxIdleConns,
		MaintNotificationsConfig: &maintnotifications.Config{
			Mode: maintnotifications.ModeDisabled,
		},
	}

	if r.Master != "" {
		opt.MasterName = r.Master
	}

	if r.Tls {
		opt.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
	}

	rds := redis.NewUniversalClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = rds.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return rds, nil
}

func (r RedisConf) MustNewUniversalRedis() redis.UniversalClient {
	rds, err := r.NewUniversalRedis()
	logx.Must(err)

	return rds
}
