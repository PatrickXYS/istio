// Copyright 2018 Istio Authors
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

package kube

import (
	v1 "k8s.io/api/core/v1"
)

func hasProxyIP(addresses []v1.EndpointAddress, proxyIP string) bool {
	for _, addr := range addresses {
		if addr.IP == proxyIP {
			return true
		}
	}
	return false
}

func getLabelValue(node *v1.Node, label string, fallBackLabel string) string {
	val := node.Labels[label]
	if val != "" {
		return val
	}

	return node.Labels[fallBackLabel]
}
