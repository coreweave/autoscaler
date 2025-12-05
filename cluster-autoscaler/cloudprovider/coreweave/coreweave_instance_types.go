/*
Copyright 2025 The Kubernetes Authors.

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

package coreweave

import "fmt"

// InstanceType represents the resource specifications for a CoreWeave instance type.
// Units are chosen to match the Kubernetes Node resource representation.
type InstanceType struct {
	// VCPU is the number of virtual CPU cores
	VCPU int64
	// MemoryKi is the amount of memory in kibibytes (1 Ki = 1024 bytes)
	MemoryKi int64
	// GPU is the number of GPUs
	GPU int64
	// EphemeralStorageMi is the amount of ephemeral storage in mebibytes (1 Mi = 1024 Ki)
	EphemeralStorageMi int64
	// Architecture is the CPU architecture (e.g., "amd64", "arm64")
	Architecture string
	// MaxPods is the maximum number of pods that can run on this instance type
	MaxPods int64
}

// InstanceTypes is a map of CoreWeave instance type names to their specifications.
// This map should be populated with the actual instance types supported by CoreWeave.
var InstanceTypes = map[string]*InstanceType{
	"turin-gp-l": {
		VCPU:               192,
		MemoryKi:           1583282428,
		GPU:                0,
		EphemeralStorageMi: 29299982,
		Architecture:       "amd64",
		MaxPods:            110,
	},
}

// GetInstanceType returns the InstanceType for the given instance type name.
// It returns an error if the instance type is not found in the InstanceTypes map.
func GetInstanceType(instanceTypeName string) (*InstanceType, error) {
	instanceType, exists := InstanceTypes[instanceTypeName]
	if !exists {
		return nil, fmt.Errorf("unknown instance type: %s", instanceTypeName)
	}
	return instanceType, nil
}
