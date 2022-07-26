//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Actions) DeepCopyInto(out *Actions) {
	*out = *in
	if in.RequestHeaderConfigurations != nil {
		in, out := &in.RequestHeaderConfigurations, &out.RequestHeaderConfigurations
		*out = make([]HeaderConfiguration, len(*in))
		copy(*out, *in)
	}
	if in.ResponseHeaderConfigurations != nil {
		in, out := &in.ResponseHeaderConfigurations, &out.ResponseHeaderConfigurations
		*out = make([]HeaderConfiguration, len(*in))
		copy(*out, *in)
	}
	out.UrlConfiguration = in.UrlConfiguration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Actions.
func (in *Actions) DeepCopy() *Actions {
	if in == nil {
		return nil
	}
	out := new(Actions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureApplicationGatewayRewrite) DeepCopyInto(out *AzureApplicationGatewayRewrite) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureApplicationGatewayRewrite.
func (in *AzureApplicationGatewayRewrite) DeepCopy() *AzureApplicationGatewayRewrite {
	if in == nil {
		return nil
	}
	out := new(AzureApplicationGatewayRewrite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureApplicationGatewayRewrite) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureApplicationGatewayRewriteList) DeepCopyInto(out *AzureApplicationGatewayRewriteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureApplicationGatewayRewrite, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureApplicationGatewayRewriteList.
func (in *AzureApplicationGatewayRewriteList) DeepCopy() *AzureApplicationGatewayRewriteList {
	if in == nil {
		return nil
	}
	out := new(AzureApplicationGatewayRewriteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureApplicationGatewayRewriteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureApplicationGatewayRewriteSpec) DeepCopyInto(out *AzureApplicationGatewayRewriteSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.RewriteRules != nil {
		in, out := &in.RewriteRules, &out.RewriteRules
		*out = make([]RewriteRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureApplicationGatewayRewriteSpec.
func (in *AzureApplicationGatewayRewriteSpec) DeepCopy() *AzureApplicationGatewayRewriteSpec {
	if in == nil {
		return nil
	}
	out := new(AzureApplicationGatewayRewriteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	return
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
func (in *HeaderConfiguration) DeepCopyInto(out *HeaderConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderConfiguration.
func (in *HeaderConfiguration) DeepCopy() *HeaderConfiguration {
	if in == nil {
		return nil
	}
	out := new(HeaderConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RewriteRule) DeepCopyInto(out *RewriteRule) {
	*out = *in
	in.Actions.DeepCopyInto(&out.Actions)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RewriteRule.
func (in *RewriteRule) DeepCopy() *RewriteRule {
	if in == nil {
		return nil
	}
	out := new(RewriteRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UrlConfiguration) DeepCopyInto(out *UrlConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UrlConfiguration.
func (in *UrlConfiguration) DeepCopy() *UrlConfiguration {
	if in == nil {
		return nil
	}
	out := new(UrlConfiguration)
	in.DeepCopyInto(out)
	return out
}