// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package helper

import (
	"fmt"

	"github.com/gardener/gardener-extensions/controllers/provider-azure/pkg/apis/azure"
)

// FindSubnetByPurpose takes a list of subnets and tries to find the first entry
// whose purpose matches with the given purpose. If no such entry is found then an error will be
// returned.
func FindSubnetByPurpose(subnets []azure.Subnet, purpose azure.Purpose) (*azure.Subnet, error) {
	for _, subnet := range subnets {
		if subnet.Purpose == purpose {
			return &subnet, nil
		}
	}
	return nil, fmt.Errorf("cannot find subnet with purpose %q", purpose)
}

// FindSecurityGroupByPurpose takes a list of security groups and tries to find the first entry
// whose purpose matches with the given purpose. If no such entry is found then an error will be
// returned.
func FindSecurityGroupByPurpose(securityGroups []azure.SecurityGroup, purpose azure.Purpose) (*azure.SecurityGroup, error) {
	for _, securityGroup := range securityGroups {
		if securityGroup.Purpose == purpose {
			return &securityGroup, nil
		}
	}
	return nil, fmt.Errorf("cannot find security group with purpose %q", purpose)
}

// FindRouteTableByPurpose takes a list of route tables and tries to find the first entry
// whose purpose matches with the given purpose. If no such entry is found then an error will be
// returned.
func FindRouteTableByPurpose(routeTables []azure.RouteTable, purpose azure.Purpose) (*azure.RouteTable, error) {
	for _, routeTable := range routeTables {
		if routeTable.Purpose == purpose {
			return &routeTable, nil
		}
	}
	return nil, fmt.Errorf("cannot find route table with purpose %q", purpose)
}

// FindAvailabilitySetByPurpose takes a list of availability sets and tries to find the first entry
// whose purpose matches with the given purpose. If no such entry is found then an error will be
// returned.
func FindAvailabilitySetByPurpose(availabilitySets []azure.AvailabilitySet, purpose azure.Purpose) (*azure.AvailabilitySet, error) {
	for _, availabilitySet := range availabilitySets {
		if availabilitySet.Purpose == purpose {
			return &availabilitySet, nil
		}
	}
	return nil, fmt.Errorf("cannot find availability set with purpose %q", purpose)
}

// FindMachineImage takes a list of machine images and tries to find the first entry
// whose name, version, and zone matches with the given name, version, and zone. If no such entry is
// found then an error will be returned.
func FindMachineImage(machineImages []azure.MachineImage, name, version string) (*azure.MachineImage, error) {
	for _, machineImage := range machineImages {
		if machineImage.Name == name && machineImage.Version == version {
			return &machineImage, nil
		}
	}
	return nil, fmt.Errorf("no machine image with name %q, version %q found", name, version)
}
