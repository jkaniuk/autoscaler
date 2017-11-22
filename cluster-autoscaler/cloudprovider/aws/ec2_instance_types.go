/*
Copyright 2017 The Kubernetes Authors.

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

// This file was generated by go generate; DO NOT EDIT

package aws

type instanceType struct {
	InstanceType string
	VCPU         int64
	MemoryMb     int64
	GPU          int64
}

// InstanceTypes is a map of ec2 resources
var InstanceTypes = map[string]*instanceType{
	"c1.medium": {
		InstanceType: "c1.medium",
		VCPU:         2,
		MemoryMb:     1740,
		GPU:          0,
	},
	"c1.xlarge": {
		InstanceType: "c1.xlarge",
		VCPU:         8,
		MemoryMb:     7168,
		GPU:          0,
	},
	"c3": {
		InstanceType: "c3",
		VCPU:         32,
		MemoryMb:     0,
		GPU:          0,
	},
	"c3.2xlarge": {
		InstanceType: "c3.2xlarge",
		VCPU:         8,
		MemoryMb:     15360,
		GPU:          0,
	},
	"c3.4xlarge": {
		InstanceType: "c3.4xlarge",
		VCPU:         16,
		MemoryMb:     30720,
		GPU:          0,
	},
	"c3.8xlarge": {
		InstanceType: "c3.8xlarge",
		VCPU:         32,
		MemoryMb:     61440,
		GPU:          0,
	},
	"c3.large": {
		InstanceType: "c3.large",
		VCPU:         2,
		MemoryMb:     3840,
		GPU:          0,
	},
	"c3.xlarge": {
		InstanceType: "c3.xlarge",
		VCPU:         4,
		MemoryMb:     7680,
		GPU:          0,
	},
	"c4": {
		InstanceType: "c4",
		VCPU:         36,
		MemoryMb:     0,
		GPU:          0,
	},
	"c4.2xlarge": {
		InstanceType: "c4.2xlarge",
		VCPU:         8,
		MemoryMb:     15360,
		GPU:          0,
	},
	"c4.4xlarge": {
		InstanceType: "c4.4xlarge",
		VCPU:         16,
		MemoryMb:     30720,
		GPU:          0,
	},
	"c4.8xlarge": {
		InstanceType: "c4.8xlarge",
		VCPU:         36,
		MemoryMb:     61440,
		GPU:          0,
	},
	"c4.large": {
		InstanceType: "c4.large",
		VCPU:         2,
		MemoryMb:     3840,
		GPU:          0,
	},
	"c4.xlarge": {
		InstanceType: "c4.xlarge",
		VCPU:         4,
		MemoryMb:     7680,
		GPU:          0,
	},
	"c5": {
		InstanceType: "c5",
		VCPU:         72,
		MemoryMb:     0,
		GPU:          0,
	},
	"c5.2xlarge": {
		InstanceType: "c5.2xlarge",
		VCPU:         8,
		MemoryMb:     16384,
		GPU:          0,
	},
	"c5.4xlarge": {
		InstanceType: "c5.4xlarge",
		VCPU:         16,
		MemoryMb:     32768,
		GPU:          0,
	},
	"c5.9xlarge": {
		InstanceType: "c5.9xlarge",
		VCPU:         36,
		MemoryMb:     73728,
		GPU:          0,
	},
	"c5.18xlarge": {
		InstanceType: "c5.18xlarge",
		VCPU:         72,
		MemoryMb:     147456,
		GPU:          0,
	},
	"c5.large": {
		InstanceType: "c5.large",
		VCPU:         2,
		MemoryMb:     4096,
		GPU:          0,
	},
	"c5.xlarge": {
		InstanceType: "c5.xlarge",
		VCPU:         4,
		MemoryMb:     8192,
		GPU:          0,
	},
	"cc1.4xlarge": {
		InstanceType: "cc1.4xlarge",
		VCPU:         16,
		MemoryMb:     23552,
		GPU:          0,
	},
	"cc2.8xlarge": {
		InstanceType: "cc2.8xlarge",
		VCPU:         32,
		MemoryMb:     61952,
		GPU:          0,
	},
	"cg1.4xlarge": {
		InstanceType: "cg1.4xlarge",
		VCPU:         16,
		MemoryMb:     23040,
		GPU:          0,
	},
	"cr1.8xlarge": {
		InstanceType: "cr1.8xlarge",
		VCPU:         32,
		MemoryMb:     249856,
		GPU:          0,
	},
	"d2": {
		InstanceType: "d2",
		VCPU:         36,
		MemoryMb:     0,
		GPU:          0,
	},
	"d2.2xlarge": {
		InstanceType: "d2.2xlarge",
		VCPU:         8,
		MemoryMb:     62464,
		GPU:          0,
	},
	"d2.4xlarge": {
		InstanceType: "d2.4xlarge",
		VCPU:         16,
		MemoryMb:     124928,
		GPU:          0,
	},
	"d2.8xlarge": {
		InstanceType: "d2.8xlarge",
		VCPU:         36,
		MemoryMb:     249856,
		GPU:          0,
	},
	"d2.xlarge": {
		InstanceType: "d2.xlarge",
		VCPU:         4,
		MemoryMb:     31232,
		GPU:          0,
	},
	"f1": {
		InstanceType: "f1",
		VCPU:         64,
		MemoryMb:     0,
		GPU:          0,
	},
	"f1.16xlarge": {
		InstanceType: "f1.16xlarge",
		VCPU:         64,
		MemoryMb:     999424,
		GPU:          0,
	},
	"f1.2xlarge": {
		InstanceType: "f1.2xlarge",
		VCPU:         8,
		MemoryMb:     124928,
		GPU:          0,
	},
	"g2": {
		InstanceType: "g2",
		VCPU:         32,
		MemoryMb:     0,
		GPU:          4,
	},
	"g2.2xlarge": {
		InstanceType: "g2.2xlarge",
		VCPU:         8,
		MemoryMb:     15360,
		GPU:          1,
	},
	"g2.8xlarge": {
		InstanceType: "g2.8xlarge",
		VCPU:         32,
		MemoryMb:     61440,
		GPU:          4,
	},
	"g3": {
		InstanceType: "g3",
		VCPU:         64,
		MemoryMb:     0,
		GPU:          4,
	},
	"g3.16xlarge": {
		InstanceType: "g3.16xlarge",
		VCPU:         64,
		MemoryMb:     499712,
		GPU:          4,
	},
	"g3.4xlarge": {
		InstanceType: "g3.4xlarge",
		VCPU:         16,
		MemoryMb:     124928,
		GPU:          1,
	},
	"g3.8xlarge": {
		InstanceType: "g3.8xlarge",
		VCPU:         32,
		MemoryMb:     249856,
		GPU:          2,
	},
	"hi1.4xlarge": {
		InstanceType: "hi1.4xlarge",
		VCPU:         16,
		MemoryMb:     61952,
		GPU:          0,
	},
	"hs1.8xlarge": {
		InstanceType: "hs1.8xlarge",
		VCPU:         17,
		MemoryMb:     119808,
		GPU:          0,
	},
	"i2": {
		InstanceType: "i2",
		VCPU:         32,
		MemoryMb:     0,
		GPU:          0,
	},
	"i2.2xlarge": {
		InstanceType: "i2.2xlarge",
		VCPU:         8,
		MemoryMb:     62464,
		GPU:          0,
	},
	"i2.4xlarge": {
		InstanceType: "i2.4xlarge",
		VCPU:         16,
		MemoryMb:     124928,
		GPU:          0,
	},
	"i2.8xlarge": {
		InstanceType: "i2.8xlarge",
		VCPU:         32,
		MemoryMb:     249856,
		GPU:          0,
	},
	"i2.xlarge": {
		InstanceType: "i2.xlarge",
		VCPU:         4,
		MemoryMb:     31232,
		GPU:          0,
	},
	"i3": {
		InstanceType: "i3",
		VCPU:         64,
		MemoryMb:     0,
		GPU:          0,
	},
	"i3.16xlarge": {
		InstanceType: "i3.16xlarge",
		VCPU:         64,
		MemoryMb:     499712,
		GPU:          0,
	},
	"i3.2xlarge": {
		InstanceType: "i3.2xlarge",
		VCPU:         8,
		MemoryMb:     62464,
		GPU:          0,
	},
	"i3.4xlarge": {
		InstanceType: "i3.4xlarge",
		VCPU:         16,
		MemoryMb:     124928,
		GPU:          0,
	},
	"i3.8xlarge": {
		InstanceType: "i3.8xlarge",
		VCPU:         32,
		MemoryMb:     249856,
		GPU:          0,
	},
	"i3.large": {
		InstanceType: "i3.large",
		VCPU:         2,
		MemoryMb:     15616,
		GPU:          0,
	},
	"i3.xlarge": {
		InstanceType: "i3.xlarge",
		VCPU:         4,
		MemoryMb:     31232,
		GPU:          0,
	},
	"m1.large": {
		InstanceType: "m1.large",
		VCPU:         2,
		MemoryMb:     7680,
		GPU:          0,
	},
	"m1.medium": {
		InstanceType: "m1.medium",
		VCPU:         1,
		MemoryMb:     3840,
		GPU:          0,
	},
	"m1.small": {
		InstanceType: "m1.small",
		VCPU:         1,
		MemoryMb:     1740,
		GPU:          0,
	},
	"m1.xlarge": {
		InstanceType: "m1.xlarge",
		VCPU:         4,
		MemoryMb:     15360,
		GPU:          0,
	},
	"m2.2xlarge": {
		InstanceType: "m2.2xlarge",
		VCPU:         4,
		MemoryMb:     35020,
		GPU:          0,
	},
	"m2.4xlarge": {
		InstanceType: "m2.4xlarge",
		VCPU:         8,
		MemoryMb:     70041,
		GPU:          0,
	},
	"m2.xlarge": {
		InstanceType: "m2.xlarge",
		VCPU:         2,
		MemoryMb:     17510,
		GPU:          0,
	},
	"m3": {
		InstanceType: "m3",
		VCPU:         8,
		MemoryMb:     0,
		GPU:          0,
	},
	"m3.2xlarge": {
		InstanceType: "m3.2xlarge",
		VCPU:         8,
		MemoryMb:     30720,
		GPU:          0,
	},
	"m3.large": {
		InstanceType: "m3.large",
		VCPU:         2,
		MemoryMb:     7680,
		GPU:          0,
	},
	"m3.medium": {
		InstanceType: "m3.medium",
		VCPU:         1,
		MemoryMb:     3840,
		GPU:          0,
	},
	"m3.xlarge": {
		InstanceType: "m3.xlarge",
		VCPU:         4,
		MemoryMb:     15360,
		GPU:          0,
	},
	"m4": {
		InstanceType: "m4",
		VCPU:         40,
		MemoryMb:     0,
		GPU:          0,
	},
	"m4.10xlarge": {
		InstanceType: "m4.10xlarge",
		VCPU:         40,
		MemoryMb:     163840,
		GPU:          0,
	},
	"m4.16xlarge": {
		InstanceType: "m4.16xlarge",
		VCPU:         64,
		MemoryMb:     262144,
		GPU:          0,
	},
	"m4.2xlarge": {
		InstanceType: "m4.2xlarge",
		VCPU:         8,
		MemoryMb:     32768,
		GPU:          0,
	},
	"m4.4xlarge": {
		InstanceType: "m4.4xlarge",
		VCPU:         16,
		MemoryMb:     65536,
		GPU:          0,
	},
	"m4.large": {
		InstanceType: "m4.large",
		VCPU:         2,
		MemoryMb:     8192,
		GPU:          0,
	},
	"m4.xlarge": {
		InstanceType: "m4.xlarge",
		VCPU:         4,
		MemoryMb:     16384,
		GPU:          0,
	},
	"p2": {
		InstanceType: "p2",
		VCPU:         64,
		MemoryMb:     0,
		GPU:          16,
	},
	"p2.16xlarge": {
		InstanceType: "p2.16xlarge",
		VCPU:         64,
		MemoryMb:     786432,
		GPU:          16,
	},
	"p2.8xlarge": {
		InstanceType: "p2.8xlarge",
		VCPU:         32,
		MemoryMb:     499712,
		GPU:          8,
	},
	"p2.xlarge": {
		InstanceType: "p2.xlarge",
		VCPU:         4,
		MemoryMb:     62464,
		GPU:          1,
	},
	"p3": {
		InstanceType: "p3",
		VCPU:         64,
		MemoryMb:     499712,
		GPU:          8,
	},
	"p3.16xlarge": {
		InstanceType: "p3.16xlarge",
		VCPU:         64,
		MemoryMb:     499712,
		GPU:          8,
	},
	"p3.2xlarge": {
		InstanceType: "p3.2xlarge",
		VCPU:         8,
		MemoryMb:     62464,
		GPU:          1,
	},
	"p3.8xlarge": {
		InstanceType: "p3.8xlarge",
		VCPU:         32,
		MemoryMb:     249856,
		GPU:          4,
	},
	"r3": {
		InstanceType: "r3",
		VCPU:         32,
		MemoryMb:     0,
		GPU:          0,
	},
	"r3.2xlarge": {
		InstanceType: "r3.2xlarge",
		VCPU:         8,
		MemoryMb:     62464,
		GPU:          0,
	},
	"r3.4xlarge": {
		InstanceType: "r3.4xlarge",
		VCPU:         16,
		MemoryMb:     124928,
		GPU:          0,
	},
	"r3.8xlarge": {
		InstanceType: "r3.8xlarge",
		VCPU:         32,
		MemoryMb:     249856,
		GPU:          0,
	},
	"r3.large": {
		InstanceType: "r3.large",
		VCPU:         2,
		MemoryMb:     15616,
		GPU:          0,
	},
	"r3.xlarge": {
		InstanceType: "r3.xlarge",
		VCPU:         4,
		MemoryMb:     31232,
		GPU:          0,
	},
	"r4": {
		InstanceType: "r4",
		VCPU:         64,
		MemoryMb:     0,
		GPU:          0,
	},
	"r4.16xlarge": {
		InstanceType: "r4.16xlarge",
		VCPU:         64,
		MemoryMb:     499712,
		GPU:          0,
	},
	"r4.2xlarge": {
		InstanceType: "r4.2xlarge",
		VCPU:         8,
		MemoryMb:     62464,
		GPU:          0,
	},
	"r4.4xlarge": {
		InstanceType: "r4.4xlarge",
		VCPU:         16,
		MemoryMb:     124928,
		GPU:          0,
	},
	"r4.8xlarge": {
		InstanceType: "r4.8xlarge",
		VCPU:         32,
		MemoryMb:     249856,
		GPU:          0,
	},
	"r4.large": {
		InstanceType: "r4.large",
		VCPU:         2,
		MemoryMb:     15616,
		GPU:          0,
	},
	"r4.xlarge": {
		InstanceType: "r4.xlarge",
		VCPU:         4,
		MemoryMb:     31232,
		GPU:          0,
	},
	"t1.micro": {
		InstanceType: "t1.micro",
		VCPU:         1,
		MemoryMb:     627,
		GPU:          0,
	},
	"t2.2xlarge": {
		InstanceType: "t2.2xlarge",
		VCPU:         8,
		MemoryMb:     32768,
		GPU:          0,
	},
	"t2.large": {
		InstanceType: "t2.large",
		VCPU:         2,
		MemoryMb:     8192,
		GPU:          0,
	},
	"t2.medium": {
		InstanceType: "t2.medium",
		VCPU:         2,
		MemoryMb:     4096,
		GPU:          0,
	},
	"t2.micro": {
		InstanceType: "t2.micro",
		VCPU:         1,
		MemoryMb:     1024,
		GPU:          0,
	},
	"t2.nano": {
		InstanceType: "t2.nano",
		VCPU:         1,
		MemoryMb:     512,
		GPU:          0,
	},
	"t2.small": {
		InstanceType: "t2.small",
		VCPU:         1,
		MemoryMb:     2048,
		GPU:          0,
	},
	"t2.xlarge": {
		InstanceType: "t2.xlarge",
		VCPU:         4,
		MemoryMb:     16384,
		GPU:          0,
	},
	"x1": {
		InstanceType: "x1",
		VCPU:         128,
		MemoryMb:     0,
		GPU:          0,
	},
	"x1.16xlarge": {
		InstanceType: "x1.16xlarge",
		VCPU:         64,
		MemoryMb:     999424,
		GPU:          0,
	},
	"x1.32xlarge": {
		InstanceType: "x1.32xlarge",
		VCPU:         128,
		MemoryMb:     1998848,
		GPU:          0,
	},
	"x1e": {
		InstanceType: "x1e",
		VCPU:         128,
		MemoryMb:     0,
		GPU:          0,
	},
	"x1e.32xlarge": {
		InstanceType: "x1e.32xlarge",
		VCPU:         128,
		MemoryMb:     3997696,
		GPU:          0,
	},
}
