// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generic

import (
	"github.com/ligato/cn-infra/core"
	"github.com/ligato/cn-infra/db/keyval/etcdv3"
	"github.com/ligato/cn-infra/logging/logrus"
	"github.com/ligato/cn-infra/messaging/kafka"
)

// Flavour is set of common used generic plugins. This flavour can be used as a base
// for different flavours. The plugins are initialized in the same order as they appear
// in the structure.
type Flavour struct {
	injected bool

	Logrus logrus.Plugin
	Etcd   etcdv3.Plugin
	Kafka  kafka.Plugin
}

// Inject interconnects plugins - injects the dependencies. If it has been called
// already it is no op.
func (g *Flavour) Inject() error {
	if g.injected {
		return nil
	}
	g.injected = true

	g.Etcd.LogFactory = &g.Logrus
	g.Kafka.LogFactory = &g.Logrus
	return nil
}

// Plugins returns all plugins from the flavour. The set of plugins is supposed
// to be passed to the agent constructor. The method calls inject to make sure that
// dependencies have been injected.
func (g *Flavour) Plugins() []*core.NamedPlugin {
	g.Inject()
	return core.ListPluginsInFlavor(g)
}
