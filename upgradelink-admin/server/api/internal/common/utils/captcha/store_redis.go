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

package captcha

import (
	"context"
	"time"
	"upgradelink-admin/server/api/internal/common/enum"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// NewRedisStore returns a redis store for captcha.
func NewRedisStore(r *redis.Redis) *RedisStore {
	return &RedisStore{
		Expiration: time.Minute * 5,
		PreKey:     enum.RedisCaptchaPrefix,
		Redis:      r,
	}
}

// RedisStore stores captcha data.
type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
	Redis      *redis.Redis
}

// UseWithCtx add context for captcha.
func (r *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	r.Context = ctx
	return r
}

// Set sets the captcha KV to redis.
func (r *RedisStore) Set(id string, value string) error {
	err := r.Redis.Setex(r.PreKey+id, value, int(r.Expiration.Seconds()))
	if err != nil {
		logx.Errorw("error occurs when captcha key sets to redis", logx.Field("detail", err))
		return err
	}
	return nil
}

// Get gets the captcha KV from redis.
func (r *RedisStore) Get(key string, clear bool) string {
	val, err := r.Redis.Get(key)
	if err != nil {
		logx.Errorw("error occurs when captcha key gets from redis", logx.Field("detail", err))
		return ""
	}
	if clear {
		_, err := r.Redis.Del(key)
		if err != nil {
			logx.Errorw("error occurs when captcha key deletes from redis", logx.Field("detail", err))
			return ""
		}
	}
	return val
}

// Verify verifies the captcha whether it is correct.
func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	key := r.PreKey + id
	v := r.Get(key, clear)
	return v == answer
}
