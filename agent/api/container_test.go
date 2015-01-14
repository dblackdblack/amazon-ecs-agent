// Copyright 2014-2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package api

import (
	"reflect"
	"testing"
)

func TestOverridden(t *testing.T) {
	container := &Container{
		Name:          "name",
		Image:         "image",
		Command:       []string{"foo", "bar"},
		Cpu:           1,
		Memory:        1,
		Links:         []string{},
		Ports:         []PortBinding{PortBinding{10, 10, ""}},
		Overrides:     ContainerOverrides{},
		DesiredStatus: ContainerRunning,
		AppliedStatus: ContainerRunning,
		KnownStatus:   ContainerRunning,
	}

	overridden := container.Overridden()
	// No overrides, should be identity
	if !reflect.DeepEqual(container, overridden) {
		t.Error("Were not equal")
	}
	if container == overridden {
		t.Error("Were pointer equal")
	}
	overridden.Name = "mutated"

	if container.Name != "name" {
		t.Error("Should make a copy")
	}
}
