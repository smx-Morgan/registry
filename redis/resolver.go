// Copyright 2022 CloudWeGo Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redis

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app/client/discovery"

	cwRedis "github.com/cloudwego-contrib/cwgo-pkg/registry/redis/redishertz"
)

var _ discovery.Resolver = (*redisResolver)(nil)

type redisResolver struct {
	resolver discovery.Resolver
}

// NewRedisResolver creates a redis resolver
func NewRedisResolver(addr string, opts ...Option) discovery.Resolver {
	o := Options{}

	for _, opt := range opts {
		opt(&o)
	}

	return cwRedis.NewRedisResolver(addr, o.cfgs...)
}

func (r *redisResolver) Target(ctx context.Context, target *discovery.TargetInfo) string {
	return r.resolver.Target(ctx, target)
}

func (r *redisResolver) Resolve(ctx context.Context, desc string) (discovery.Result, error) {
	return r.resolver.Resolve(ctx, desc)
}

func (r *redisResolver) Name() string {
	return r.resolver.Name()
}
