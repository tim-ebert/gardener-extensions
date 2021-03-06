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

package helper_test

import (
	"github.com/gardener/gardener-extensions/controllers/provider-gcp/pkg/apis/gcp"
	. "github.com/gardener/gardener-extensions/controllers/provider-gcp/pkg/apis/gcp/helper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Helper", func() {
	var (
		purpose      gcp.SubnetPurpose = "foo"
		purposeWrong gcp.SubnetPurpose = "baz"
	)

	DescribeTable("#FindSubnetByPurpose",
		func(subnets []gcp.Subnet, purpose gcp.SubnetPurpose, expectedSubnet *gcp.Subnet, expectErr bool) {
			subnet, err := FindSubnetByPurpose(subnets, purpose)
			expectResults(subnet, expectedSubnet, err, expectErr)
		},

		Entry("list is nil", nil, purpose, nil, true),
		Entry("empty list", []gcp.Subnet{}, purpose, nil, true),
		Entry("entry not found", []gcp.Subnet{{Name: "bar", Purpose: purposeWrong}}, purpose, nil, true),
		Entry("entry exists", []gcp.Subnet{{Name: "bar", Purpose: purpose}}, purpose, &gcp.Subnet{Name: "bar", Purpose: purpose}, false),
	)

	DescribeTable("#FindMachineImage",
		func(machineImages []gcp.MachineImage, name, version string, expectedMachineImage *gcp.MachineImage, expectErr bool) {
			machineImage, err := FindMachineImage(machineImages, name, version)
			expectResults(machineImage, expectedMachineImage, err, expectErr)
		},

		Entry("list is nil", nil, "foo", "1.2.3", nil, true),
		Entry("empty list", []gcp.MachineImage{}, "foo", "1.2.3", nil, true),
		Entry("entry not found (no name)", []gcp.MachineImage{{Name: "bar", Version: "1.2.3", Image: "image123"}}, "foo", "1.2.3", nil, true),
		Entry("entry not found (no version)", []gcp.MachineImage{{Name: "bar", Version: "1.2.3", Image: "image123"}}, "foo", "1.2.4", nil, true),
		Entry("entry exists", []gcp.MachineImage{{Name: "bar", Version: "1.2.3", Image: "image123"}}, "bar", "1.2.3", &gcp.MachineImage{Name: "bar", Version: "1.2.3", Image: "image123"}, false),
	)
})

func expectResults(result, expected interface{}, err error, expectErr bool) {
	if !expectErr {
		Expect(result).To(Equal(expected))
		Expect(err).NotTo(HaveOccurred())
	} else {
		Expect(result).To(BeNil())
		Expect(err).To(HaveOccurred())
	}
}
