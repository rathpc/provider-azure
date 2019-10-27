// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlServer) DeepCopyInto(out *MysqlServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlServer.
func (in *MysqlServer) DeepCopy() *MysqlServer {
	if in == nil {
		return nil
	}
	out := new(MysqlServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MysqlServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlServerList) DeepCopyInto(out *MysqlServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MysqlServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlServerList.
func (in *MysqlServerList) DeepCopy() *MysqlServerList {
	if in == nil {
		return nil
	}
	out := new(MysqlServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MysqlServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlServerNameReferencer) DeepCopyInto(out *MysqlServerNameReferencer) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlServerNameReferencer.
func (in *MysqlServerNameReferencer) DeepCopy() *MysqlServerNameReferencer {
	if in == nil {
		return nil
	}
	out := new(MysqlServerNameReferencer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlServerVirtualNetworkRule) DeepCopyInto(out *MysqlServerVirtualNetworkRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlServerVirtualNetworkRule.
func (in *MysqlServerVirtualNetworkRule) DeepCopy() *MysqlServerVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(MysqlServerVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MysqlServerVirtualNetworkRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlServerVirtualNetworkRuleList) DeepCopyInto(out *MysqlServerVirtualNetworkRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MysqlServerVirtualNetworkRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlServerVirtualNetworkRuleList.
func (in *MysqlServerVirtualNetworkRuleList) DeepCopy() *MysqlServerVirtualNetworkRuleList {
	if in == nil {
		return nil
	}
	out := new(MysqlServerVirtualNetworkRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MysqlServerVirtualNetworkRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlVirtualNetworkRuleSpec) DeepCopyInto(out *MysqlVirtualNetworkRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	if in.ServerNameRef != nil {
		in, out := &in.ServerNameRef, &out.ServerNameRef
		*out = new(MysqlServerNameReferencer)
		**out = **in
	}
	out.VirtualNetworkRuleProperties = in.VirtualNetworkRuleProperties
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlVirtualNetworkRuleSpec.
func (in *MysqlVirtualNetworkRuleSpec) DeepCopy() *MysqlVirtualNetworkRuleSpec {
	if in == nil {
		return nil
	}
	out := new(MysqlVirtualNetworkRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlServer) DeepCopyInto(out *PostgresqlServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlServer.
func (in *PostgresqlServer) DeepCopy() *PostgresqlServer {
	if in == nil {
		return nil
	}
	out := new(PostgresqlServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlServerList) DeepCopyInto(out *PostgresqlServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgresqlServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlServerList.
func (in *PostgresqlServerList) DeepCopy() *PostgresqlServerList {
	if in == nil {
		return nil
	}
	out := new(PostgresqlServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlServerNameReferencer) DeepCopyInto(out *PostgresqlServerNameReferencer) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlServerNameReferencer.
func (in *PostgresqlServerNameReferencer) DeepCopy() *PostgresqlServerNameReferencer {
	if in == nil {
		return nil
	}
	out := new(PostgresqlServerNameReferencer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlServerVirtualNetworkRule) DeepCopyInto(out *PostgresqlServerVirtualNetworkRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlServerVirtualNetworkRule.
func (in *PostgresqlServerVirtualNetworkRule) DeepCopy() *PostgresqlServerVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(PostgresqlServerVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlServerVirtualNetworkRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlServerVirtualNetworkRuleList) DeepCopyInto(out *PostgresqlServerVirtualNetworkRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgresqlServerVirtualNetworkRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlServerVirtualNetworkRuleList.
func (in *PostgresqlServerVirtualNetworkRuleList) DeepCopy() *PostgresqlServerVirtualNetworkRuleList {
	if in == nil {
		return nil
	}
	out := new(PostgresqlServerVirtualNetworkRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlServerVirtualNetworkRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlVirtualNetworkRuleSpec) DeepCopyInto(out *PostgresqlVirtualNetworkRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	if in.ServerNameRef != nil {
		in, out := &in.ServerNameRef, &out.ServerNameRef
		*out = new(PostgresqlServerNameReferencer)
		**out = **in
	}
	out.VirtualNetworkRuleProperties = in.VirtualNetworkRuleProperties
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlVirtualNetworkRuleSpec.
func (in *PostgresqlVirtualNetworkRuleSpec) DeepCopy() *PostgresqlVirtualNetworkRuleSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresqlVirtualNetworkRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PricingTierSpec) DeepCopyInto(out *PricingTierSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PricingTierSpec.
func (in *PricingTierSpec) DeepCopy() *PricingTierSpec {
	if in == nil {
		return nil
	}
	out := new(PricingTierSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupNameReferencerForSQLServer) DeepCopyInto(out *ResourceGroupNameReferencerForSQLServer) {
	*out = *in
	out.ResourceGroupNameReferencer = in.ResourceGroupNameReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupNameReferencerForSQLServer.
func (in *ResourceGroupNameReferencerForSQLServer) DeepCopy() *ResourceGroupNameReferencerForSQLServer {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupNameReferencerForSQLServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLServerClass) DeepCopyInto(out *SQLServerClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.SpecTemplate.DeepCopyInto(&out.SpecTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLServerClass.
func (in *SQLServerClass) DeepCopy() *SQLServerClass {
	if in == nil {
		return nil
	}
	out := new(SQLServerClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SQLServerClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLServerClassList) DeepCopyInto(out *SQLServerClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SQLServerClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLServerClassList.
func (in *SQLServerClassList) DeepCopy() *SQLServerClassList {
	if in == nil {
		return nil
	}
	out := new(SQLServerClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SQLServerClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLServerClassSpecTemplate) DeepCopyInto(out *SQLServerClassSpecTemplate) {
	*out = *in
	in.ClassSpecTemplate.DeepCopyInto(&out.ClassSpecTemplate)
	in.SQLServerParameters.DeepCopyInto(&out.SQLServerParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLServerClassSpecTemplate.
func (in *SQLServerClassSpecTemplate) DeepCopy() *SQLServerClassSpecTemplate {
	if in == nil {
		return nil
	}
	out := new(SQLServerClassSpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLServerParameters) DeepCopyInto(out *SQLServerParameters) {
	*out = *in
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(ResourceGroupNameReferencerForSQLServer)
		**out = **in
	}
	out.PricingTier = in.PricingTier
	out.StorageProfile = in.StorageProfile
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLServerParameters.
func (in *SQLServerParameters) DeepCopy() *SQLServerParameters {
	if in == nil {
		return nil
	}
	out := new(SQLServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLServerSpec) DeepCopyInto(out *SQLServerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.SQLServerParameters.DeepCopyInto(&out.SQLServerParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLServerSpec.
func (in *SQLServerSpec) DeepCopy() *SQLServerSpec {
	if in == nil {
		return nil
	}
	out := new(SQLServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLServerStatus) DeepCopyInto(out *SQLServerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLServerStatus.
func (in *SQLServerStatus) DeepCopy() *SQLServerStatus {
	if in == nil {
		return nil
	}
	out := new(SQLServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerNameReferencerForMysqlServerVirtualNetworkRule) DeepCopyInto(out *ServerNameReferencerForMysqlServerVirtualNetworkRule) {
	*out = *in
	out.MysqlServerNameReferencer = in.MysqlServerNameReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerNameReferencerForMysqlServerVirtualNetworkRule.
func (in *ServerNameReferencerForMysqlServerVirtualNetworkRule) DeepCopy() *ServerNameReferencerForMysqlServerVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(ServerNameReferencerForMysqlServerVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerNameReferencerForPostgresqlServerVirtualNetworkRule) DeepCopyInto(out *ServerNameReferencerForPostgresqlServerVirtualNetworkRule) {
	*out = *in
	out.PostgresqlServerNameReferencer = in.PostgresqlServerNameReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerNameReferencerForPostgresqlServerVirtualNetworkRule.
func (in *ServerNameReferencerForPostgresqlServerVirtualNetworkRule) DeepCopy() *ServerNameReferencerForPostgresqlServerVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(ServerNameReferencerForPostgresqlServerVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageProfileSpec) DeepCopyInto(out *StorageProfileSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProfileSpec.
func (in *StorageProfileSpec) DeepCopy() *StorageProfileSpec {
	if in == nil {
		return nil
	}
	out := new(StorageProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRuleProperties) DeepCopyInto(out *VirtualNetworkRuleProperties) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRuleProperties.
func (in *VirtualNetworkRuleProperties) DeepCopy() *VirtualNetworkRuleProperties {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRuleProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRuleStatus) DeepCopyInto(out *VirtualNetworkRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRuleStatus.
func (in *VirtualNetworkRuleStatus) DeepCopy() *VirtualNetworkRuleStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRuleStatus)
	in.DeepCopyInto(out)
	return out
}
