// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ReplicationController represents the configuration of a replication controller.
type ReplicationControllerPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// If the Labels of a ReplicationController are empty, they are defaulted to be the same as the Pod(s) that the replication controller manages. Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Spec defines the specification of the desired behavior of the replication controller. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec ReplicationControllerSpecPatchPtrOutput `pulumi:"spec"`
	// Status is the most recently observed status of the replication controller. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status ReplicationControllerStatusPatchPtrOutput `pulumi:"status"`
}

// NewReplicationControllerPatch registers a new resource with the given unique name, arguments, and options.
func NewReplicationControllerPatch(ctx *pulumi.Context,
	name string, args *ReplicationControllerPatchArgs, opts ...pulumi.ResourceOption) (*ReplicationControllerPatch, error) {
	if args == nil {
		args = &ReplicationControllerPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("ReplicationController")
	var resource ReplicationControllerPatch
	err := ctx.RegisterResource("kubernetes:core/v1:ReplicationControllerPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationControllerPatch gets an existing ReplicationControllerPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationControllerPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationControllerPatchState, opts ...pulumi.ResourceOption) (*ReplicationControllerPatch, error) {
	var resource ReplicationControllerPatch
	err := ctx.ReadResource("kubernetes:core/v1:ReplicationControllerPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationControllerPatch resources.
type replicationControllerPatchState struct {
}

type ReplicationControllerPatchState struct {
}

func (ReplicationControllerPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationControllerPatchState)(nil)).Elem()
}

type replicationControllerPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// If the Labels of a ReplicationController are empty, they are defaulted to be the same as the Pod(s) that the replication controller manages. Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Spec defines the specification of the desired behavior of the replication controller. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *ReplicationControllerSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a ReplicationControllerPatch resource.
type ReplicationControllerPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// If the Labels of a ReplicationController are empty, they are defaulted to be the same as the Pod(s) that the replication controller manages. Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// Spec defines the specification of the desired behavior of the replication controller. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec ReplicationControllerSpecPatchPtrInput
}

func (ReplicationControllerPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationControllerPatchArgs)(nil)).Elem()
}

type ReplicationControllerPatchInput interface {
	pulumi.Input

	ToReplicationControllerPatchOutput() ReplicationControllerPatchOutput
	ToReplicationControllerPatchOutputWithContext(ctx context.Context) ReplicationControllerPatchOutput
}

func (*ReplicationControllerPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationControllerPatch)(nil)).Elem()
}

func (i *ReplicationControllerPatch) ToReplicationControllerPatchOutput() ReplicationControllerPatchOutput {
	return i.ToReplicationControllerPatchOutputWithContext(context.Background())
}

func (i *ReplicationControllerPatch) ToReplicationControllerPatchOutputWithContext(ctx context.Context) ReplicationControllerPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationControllerPatchOutput)
}

// ReplicationControllerPatchArrayInput is an input type that accepts ReplicationControllerPatchArray and ReplicationControllerPatchArrayOutput values.
// You can construct a concrete instance of `ReplicationControllerPatchArrayInput` via:
//
//          ReplicationControllerPatchArray{ ReplicationControllerPatchArgs{...} }
type ReplicationControllerPatchArrayInput interface {
	pulumi.Input

	ToReplicationControllerPatchArrayOutput() ReplicationControllerPatchArrayOutput
	ToReplicationControllerPatchArrayOutputWithContext(context.Context) ReplicationControllerPatchArrayOutput
}

type ReplicationControllerPatchArray []ReplicationControllerPatchInput

func (ReplicationControllerPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationControllerPatch)(nil)).Elem()
}

func (i ReplicationControllerPatchArray) ToReplicationControllerPatchArrayOutput() ReplicationControllerPatchArrayOutput {
	return i.ToReplicationControllerPatchArrayOutputWithContext(context.Background())
}

func (i ReplicationControllerPatchArray) ToReplicationControllerPatchArrayOutputWithContext(ctx context.Context) ReplicationControllerPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationControllerPatchArrayOutput)
}

// ReplicationControllerPatchMapInput is an input type that accepts ReplicationControllerPatchMap and ReplicationControllerPatchMapOutput values.
// You can construct a concrete instance of `ReplicationControllerPatchMapInput` via:
//
//          ReplicationControllerPatchMap{ "key": ReplicationControllerPatchArgs{...} }
type ReplicationControllerPatchMapInput interface {
	pulumi.Input

	ToReplicationControllerPatchMapOutput() ReplicationControllerPatchMapOutput
	ToReplicationControllerPatchMapOutputWithContext(context.Context) ReplicationControllerPatchMapOutput
}

type ReplicationControllerPatchMap map[string]ReplicationControllerPatchInput

func (ReplicationControllerPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationControllerPatch)(nil)).Elem()
}

func (i ReplicationControllerPatchMap) ToReplicationControllerPatchMapOutput() ReplicationControllerPatchMapOutput {
	return i.ToReplicationControllerPatchMapOutputWithContext(context.Background())
}

func (i ReplicationControllerPatchMap) ToReplicationControllerPatchMapOutputWithContext(ctx context.Context) ReplicationControllerPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationControllerPatchMapOutput)
}

type ReplicationControllerPatchOutput struct{ *pulumi.OutputState }

func (ReplicationControllerPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationControllerPatch)(nil)).Elem()
}

func (o ReplicationControllerPatchOutput) ToReplicationControllerPatchOutput() ReplicationControllerPatchOutput {
	return o
}

func (o ReplicationControllerPatchOutput) ToReplicationControllerPatchOutputWithContext(ctx context.Context) ReplicationControllerPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ReplicationControllerPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationControllerPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ReplicationControllerPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationControllerPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// If the Labels of a ReplicationController are empty, they are defaulted to be the same as the Pod(s) that the replication controller manages. Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ReplicationControllerPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *ReplicationControllerPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Spec defines the specification of the desired behavior of the replication controller. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o ReplicationControllerPatchOutput) Spec() ReplicationControllerSpecPatchPtrOutput {
	return o.ApplyT(func(v *ReplicationControllerPatch) ReplicationControllerSpecPatchPtrOutput { return v.Spec }).(ReplicationControllerSpecPatchPtrOutput)
}

// Status is the most recently observed status of the replication controller. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o ReplicationControllerPatchOutput) Status() ReplicationControllerStatusPatchPtrOutput {
	return o.ApplyT(func(v *ReplicationControllerPatch) ReplicationControllerStatusPatchPtrOutput { return v.Status }).(ReplicationControllerStatusPatchPtrOutput)
}

type ReplicationControllerPatchArrayOutput struct{ *pulumi.OutputState }

func (ReplicationControllerPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationControllerPatch)(nil)).Elem()
}

func (o ReplicationControllerPatchArrayOutput) ToReplicationControllerPatchArrayOutput() ReplicationControllerPatchArrayOutput {
	return o
}

func (o ReplicationControllerPatchArrayOutput) ToReplicationControllerPatchArrayOutputWithContext(ctx context.Context) ReplicationControllerPatchArrayOutput {
	return o
}

func (o ReplicationControllerPatchArrayOutput) Index(i pulumi.IntInput) ReplicationControllerPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReplicationControllerPatch {
		return vs[0].([]*ReplicationControllerPatch)[vs[1].(int)]
	}).(ReplicationControllerPatchOutput)
}

type ReplicationControllerPatchMapOutput struct{ *pulumi.OutputState }

func (ReplicationControllerPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationControllerPatch)(nil)).Elem()
}

func (o ReplicationControllerPatchMapOutput) ToReplicationControllerPatchMapOutput() ReplicationControllerPatchMapOutput {
	return o
}

func (o ReplicationControllerPatchMapOutput) ToReplicationControllerPatchMapOutputWithContext(ctx context.Context) ReplicationControllerPatchMapOutput {
	return o
}

func (o ReplicationControllerPatchMapOutput) MapIndex(k pulumi.StringInput) ReplicationControllerPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReplicationControllerPatch {
		return vs[0].(map[string]*ReplicationControllerPatch)[vs[1].(string)]
	}).(ReplicationControllerPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationControllerPatchInput)(nil)).Elem(), &ReplicationControllerPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationControllerPatchArrayInput)(nil)).Elem(), ReplicationControllerPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationControllerPatchMapInput)(nil)).Elem(), ReplicationControllerPatchMap{})
	pulumi.RegisterOutputType(ReplicationControllerPatchOutput{})
	pulumi.RegisterOutputType(ReplicationControllerPatchArrayOutput{})
	pulumi.RegisterOutputType(ReplicationControllerPatchMapOutput{})
}
