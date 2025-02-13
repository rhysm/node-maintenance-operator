// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeMaintenance) DeepCopyInto(out *NodeMaintenance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMaintenance.
func (in *NodeMaintenance) DeepCopy() *NodeMaintenance {
	if in == nil {
		return nil
	}
	out := new(NodeMaintenance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeMaintenance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeMaintenanceList) DeepCopyInto(out *NodeMaintenanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeMaintenance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMaintenanceList.
func (in *NodeMaintenanceList) DeepCopy() *NodeMaintenanceList {
	if in == nil {
		return nil
	}
	out := new(NodeMaintenanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeMaintenanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeMaintenanceSpec) DeepCopyInto(out *NodeMaintenanceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMaintenanceSpec.
func (in *NodeMaintenanceSpec) DeepCopy() *NodeMaintenanceSpec {
	if in == nil {
		return nil
	}
	out := new(NodeMaintenanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeMaintenanceStatus) DeepCopyInto(out *NodeMaintenanceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMaintenanceStatus.
func (in *NodeMaintenanceStatus) DeepCopy() *NodeMaintenanceStatus {
	if in == nil {
		return nil
	}
	out := new(NodeMaintenanceStatus)
	in.DeepCopyInto(out)
	return out
}
