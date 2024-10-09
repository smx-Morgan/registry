/*
 * Copyright 2021 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package polaris

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app/client/discovery"

	cwPolaris "github.com/cloudwego-contrib/cwgo-pkg/registry/polaris/polarishertz"
)

// Resolver is extension interface of Hertz discovery.Resolver.
type Resolver interface {
	discovery.Resolver
}

// polarisResolver is a resolver using polaris.
type polarisResolver struct {
	resolver cwPolaris.Resolver
}

// NewPolarisResolver creates a polaris based resolver.
func NewPolarisResolver(configFile ...string) (Resolver, error) {
	return cwPolaris.NewPolarisResolver(configFile...)
}

// Target implements the Resolver interface.
func (polaris *polarisResolver) Target(ctx context.Context, target *discovery.TargetInfo) string {
	return polaris.resolver.Target(ctx, target)
}

// Resolve implements the Resolver interface.
func (polaris *polarisResolver) Resolve(ctx context.Context, desc string) (discovery.Result, error) {
	return polaris.resolver.Resolve(ctx, desc)
}

// Name implements the Resolver interface.
func (polaris *polarisResolver) Name() string {
	return polaris.resolver.Name()
}
