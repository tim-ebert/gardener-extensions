// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package calico

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAM) DeepCopyInto(out *IPAM) {
	*out = *in
	if in.CIDR != nil {
		in, out := &in.CIDR, &out.CIDR
		*out = new(CIDR)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAM.
func (in *IPAM) DeepCopy() *IPAM {
	if in == nil {
		return nil
	}
	out := new(IPAM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPv4) DeepCopyInto(out *IPv4) {
	*out = *in
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = new(IPv4Pool)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(IPv4PoolMode)
		**out = **in
	}
	if in.AutoDetectionMethod != nil {
		in, out := &in.AutoDetectionMethod, &out.AutoDetectionMethod
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPv4.
func (in *IPv4) DeepCopy() *IPv4 {
	if in == nil {
		return nil
	}
	out := new(IPv4)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfig) DeepCopyInto(out *NetworkConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(Backend)
		**out = **in
	}
	if in.IPAM != nil {
		in, out := &in.IPAM, &out.IPAM
		*out = new(IPAM)
		(*in).DeepCopyInto(*out)
	}
	if in.IPv4 != nil {
		in, out := &in.IPv4, &out.IPv4
		*out = new(IPv4)
		(*in).DeepCopyInto(*out)
	}
	if in.Typha != nil {
		in, out := &in.Typha, &out.Typha
		*out = new(Typha)
		**out = **in
	}
	if in.IPIP != nil {
		in, out := &in.IPIP, &out.IPIP
		*out = new(IPv4PoolMode)
		**out = **in
	}
	if in.IPAutoDetectionMethod != nil {
		in, out := &in.IPAutoDetectionMethod, &out.IPAutoDetectionMethod
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfig.
func (in *NetworkConfig) DeepCopy() *NetworkConfig {
	if in == nil {
		return nil
	}
	out := new(NetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkStatus) DeepCopyInto(out *NetworkStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkStatus.
func (in *NetworkStatus) DeepCopy() *NetworkStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Typha) DeepCopyInto(out *Typha) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Typha.
func (in *Typha) DeepCopy() *Typha {
	if in == nil {
		return nil
	}
	out := new(Typha)
	in.DeepCopyInto(out)
	return out
}
