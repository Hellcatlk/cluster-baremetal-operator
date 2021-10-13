// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudPrivateIPConfig) DeepCopyInto(out *CloudPrivateIPConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudPrivateIPConfig.
func (in *CloudPrivateIPConfig) DeepCopy() *CloudPrivateIPConfig {
	if in == nil {
		return nil
	}
	out := new(CloudPrivateIPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudPrivateIPConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudPrivateIPConfigList) DeepCopyInto(out *CloudPrivateIPConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudPrivateIPConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudPrivateIPConfigList.
func (in *CloudPrivateIPConfigList) DeepCopy() *CloudPrivateIPConfigList {
	if in == nil {
		return nil
	}
	out := new(CloudPrivateIPConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudPrivateIPConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudPrivateIPConfigSpec) DeepCopyInto(out *CloudPrivateIPConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudPrivateIPConfigSpec.
func (in *CloudPrivateIPConfigSpec) DeepCopy() *CloudPrivateIPConfigSpec {
	if in == nil {
		return nil
	}
	out := new(CloudPrivateIPConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudPrivateIPConfigStatus) DeepCopyInto(out *CloudPrivateIPConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudPrivateIPConfigStatus.
func (in *CloudPrivateIPConfigStatus) DeepCopy() *CloudPrivateIPConfigStatus {
	if in == nil {
		return nil
	}
	out := new(CloudPrivateIPConfigStatus)
	in.DeepCopyInto(out)
	return out
}
