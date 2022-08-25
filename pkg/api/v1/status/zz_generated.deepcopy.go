//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (C) MongoDB, Inc. 2020-present.

Licensed under the Apache License, Version 2.0 (the "License"); you may
not use this file except in compliance with the License. You may obtain
a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by controller-gen. DO NOT EDIT.

package status

import (
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/authmode"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/project"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtlasDatabaseUserStatus) DeepCopyInto(out *AtlasDatabaseUserStatus) {
	*out = *in
	in.Common.DeepCopyInto(&out.Common)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtlasDatabaseUserStatus.
func (in *AtlasDatabaseUserStatus) DeepCopy() *AtlasDatabaseUserStatus {
	if in == nil {
		return nil
	}
	out := new(AtlasDatabaseUserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtlasDeploymentStatus) DeepCopyInto(out *AtlasDeploymentStatus) {
	*out = *in
	in.Common.DeepCopyInto(&out.Common)
	if in.ConnectionStrings != nil {
		in, out := &in.ConnectionStrings, &out.ConnectionStrings
		*out = new(ConnectionStrings)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtlasDeploymentStatus.
func (in *AtlasDeploymentStatus) DeepCopy() *AtlasDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(AtlasDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtlasNetworkPeer) DeepCopyInto(out *AtlasNetworkPeer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtlasNetworkPeer.
func (in *AtlasNetworkPeer) DeepCopy() *AtlasNetworkPeer {
	if in == nil {
		return nil
	}
	out := new(AtlasNetworkPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtlasProjectStatus) DeepCopyInto(out *AtlasProjectStatus) {
	*out = *in
	in.Common.DeepCopyInto(&out.Common)
	if in.ExpiredIPAccessList != nil {
		in, out := &in.ExpiredIPAccessList, &out.ExpiredIPAccessList
		*out = make([]project.IPAccessList, len(*in))
		copy(*out, *in)
	}
	if in.PrivateEndpoints != nil {
		in, out := &in.PrivateEndpoints, &out.PrivateEndpoints
		*out = make([]ProjectPrivateEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkPeers != nil {
		in, out := &in.NetworkPeers, &out.NetworkPeers
		*out = make([]AtlasNetworkPeer, len(*in))
		copy(*out, *in)
	}
	if in.AuthModes != nil {
		in, out := &in.AuthModes, &out.AuthModes
		*out = make(authmode.AuthModes, len(*in))
		copy(*out, *in)
	}
	if in.CloudProviderAccessRoles != nil {
		in, out := &in.CloudProviderAccessRoles, &out.CloudProviderAccessRoles
		*out = make([]CloudProviderAccessRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(Prometheus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtlasProjectStatus.
func (in *AtlasProjectStatus) DeepCopy() *AtlasProjectStatus {
	if in == nil {
		return nil
	}
	out := new(AtlasProjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudProviderAccessRole) DeepCopyInto(out *CloudProviderAccessRole) {
	*out = *in
	if in.FeatureUsages != nil {
		in, out := &in.FeatureUsages, &out.FeatureUsages
		*out = make([]FeatureUsage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudProviderAccessRole.
func (in *CloudProviderAccessRole) DeepCopy() *CloudProviderAccessRole {
	if in == nil {
		return nil
	}
	out := new(CloudProviderAccessRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Common) DeepCopyInto(out *Common) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Common.
func (in *Common) DeepCopy() *Common {
	if in == nil {
		return nil
	}
	out := new(Common)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionStrings) DeepCopyInto(out *ConnectionStrings) {
	*out = *in
	if in.PrivateEndpoint != nil {
		in, out := &in.PrivateEndpoint, &out.PrivateEndpoint
		*out = make([]PrivateEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionStrings.
func (in *ConnectionStrings) DeepCopy() *ConnectionStrings {
	if in == nil {
		return nil
	}
	out := new(ConnectionStrings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureUsage) DeepCopyInto(out *FeatureUsage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureUsage.
func (in *FeatureUsage) DeepCopy() *FeatureUsage {
	if in == nil {
		return nil
	}
	out := new(FeatureUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPEndpoint) DeepCopyInto(out *GCPEndpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPEndpoint.
func (in *GCPEndpoint) DeepCopy() *GCPEndpoint {
	if in == nil {
		return nil
	}
	out := new(GCPEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpoint) DeepCopyInto(out *PrivateEndpoint) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]Endpoint, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpoint.
func (in *PrivateEndpoint) DeepCopy() *PrivateEndpoint {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectPrivateEndpoint) DeepCopyInto(out *ProjectPrivateEndpoint) {
	*out = *in
	if in.ServiceAttachmentNames != nil {
		in, out := &in.ServiceAttachmentNames, &out.ServiceAttachmentNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]GCPEndpoint, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectPrivateEndpoint.
func (in *ProjectPrivateEndpoint) DeepCopy() *ProjectPrivateEndpoint {
	if in == nil {
		return nil
	}
	out := new(ProjectPrivateEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Prometheus) DeepCopyInto(out *Prometheus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prometheus.
func (in *Prometheus) DeepCopy() *Prometheus {
	if in == nil {
		return nil
	}
	out := new(Prometheus)
	in.DeepCopyInto(out)
	return out
}
