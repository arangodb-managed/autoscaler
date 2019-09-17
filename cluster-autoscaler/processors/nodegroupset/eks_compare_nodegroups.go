/*
Copyright 2018 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nodegroupset

import (
	schedulernodeinfo "k8s.io/kubernetes/pkg/scheduler/nodeinfo"
)

// EksInstanceTypeLabel is a label specifying EKS instance type of a particular node.
const EksInstanceTypeLabel = "beta.kubernetes.io/instance-type"

func nodesHaveSameInstanceType(n1, n2 *schedulernodeinfo.NodeInfo) bool {
	n1InstanceType := n1.Node().Labels[EksInstanceTypeLabel]
	n2InstanceType := n2.Node().Labels[EksInstanceTypeLabel]
	return n1InstanceType != "" && n1InstanceType == n2InstanceType
}

// IsEksNodeInfoSimilar compares if two nodes should be considered part of the
// same NodeGroupSet. This is true if they either have the same instance type
// or match usual conditions checked by IsNodeInfoSimilar.
func IsEksNodeInfoSimilar(n1, n2 *schedulernodeinfo.NodeInfo) bool {
	if nodesHaveSameInstanceType(n1, n2) {
		return true
	}
	return IsNodeInfoSimilar(n1, n2)
}