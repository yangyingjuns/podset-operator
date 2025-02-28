/*
Copyright 2021 The Pixiu Authors.

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

package types

const (
	// PodSetFailure is added in a pod set when one of its pods fails to be created
	// due to insufficient quota, limit ranges, pod security policy, node selectors, etc. or deleted
	// due to kubelet being down or finalizers are failing.
	PodSetFailure string = "PodSetFailure"

	PodSetSuccess string = "PodSetSuccess"

	// MinimumReplicasAvailable is added in a podSet when it has its minimum replicas required available.
	MinimumReplicasAvailable = "MinimumReplicasAvailable"

	MinimumReplicasUnavailable = "MinimumReplicasUnavailable"
)
