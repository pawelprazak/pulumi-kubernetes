// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// StatefulSet represents a set of pods with consistent identities. Identities are defined as:
//  - Network: A single stable DNS and hostname.
//  - Storage: As many VolumeClaims as requested.
//    The StatefulSet guarantees that a given network identity will always map to the same storage identity.
//
// This resource waits until its status is ready before registering success
// for create/update, and populating output properties from the current state of the resource.
// The following conditions are used to determine whether the resource creation has
// succeeded or failed:
//
// 1. The value of 'spec.replicas' matches '.status.replicas', '.status.currentReplicas',
//    and '.status.readyReplicas'.
// 2. The value of '.status.updateRevision' matches '.status.currentRevision'.
//
// If the StatefulSet has not reached a Ready state after 10 minutes, it will
// time out and mark the resource update as Failed. You can override the default timeout value
// by setting the 'customTimeouts' option on the resource.
type StatefulSetPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Spec defines the desired identities of pods in this set.
	Spec StatefulSetSpecPatchPtrOutput `pulumi:"spec"`
	// Status is the current status of Pods in this StatefulSet. This data may be out of date by some window of time.
	Status StatefulSetStatusPatchPtrOutput `pulumi:"status"`
}

// NewStatefulSetPatch registers a new resource with the given unique name, arguments, and options.
func NewStatefulSetPatch(ctx *pulumi.Context,
	name string, args *StatefulSetPatchArgs, opts ...pulumi.ResourceOption) (*StatefulSetPatch, error) {
	if args == nil {
		args = &StatefulSetPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("apps/v1")
	args.Kind = pulumi.StringPtr("StatefulSet")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:apps/v1beta1:StatefulSetPatch"),
		},
		{
			Type: pulumi.String("kubernetes:apps/v1beta2:StatefulSetPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource StatefulSetPatch
	err := ctx.RegisterResource("kubernetes:apps/v1:StatefulSetPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStatefulSetPatch gets an existing StatefulSetPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStatefulSetPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StatefulSetPatchState, opts ...pulumi.ResourceOption) (*StatefulSetPatch, error) {
	var resource StatefulSetPatch
	err := ctx.ReadResource("kubernetes:apps/v1:StatefulSetPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StatefulSetPatch resources.
type statefulSetPatchState struct {
}

type StatefulSetPatchState struct {
}

func (StatefulSetPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*statefulSetPatchState)(nil)).Elem()
}

type statefulSetPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Spec defines the desired identities of pods in this set.
	Spec *StatefulSetSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a StatefulSetPatch resource.
type StatefulSetPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// Spec defines the desired identities of pods in this set.
	Spec StatefulSetSpecPatchPtrInput
}

func (StatefulSetPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*statefulSetPatchArgs)(nil)).Elem()
}

type StatefulSetPatchInput interface {
	pulumi.Input

	ToStatefulSetPatchOutput() StatefulSetPatchOutput
	ToStatefulSetPatchOutputWithContext(ctx context.Context) StatefulSetPatchOutput
}

func (*StatefulSetPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**StatefulSetPatch)(nil)).Elem()
}

func (i *StatefulSetPatch) ToStatefulSetPatchOutput() StatefulSetPatchOutput {
	return i.ToStatefulSetPatchOutputWithContext(context.Background())
}

func (i *StatefulSetPatch) ToStatefulSetPatchOutputWithContext(ctx context.Context) StatefulSetPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatefulSetPatchOutput)
}

// StatefulSetPatchArrayInput is an input type that accepts StatefulSetPatchArray and StatefulSetPatchArrayOutput values.
// You can construct a concrete instance of `StatefulSetPatchArrayInput` via:
//
//          StatefulSetPatchArray{ StatefulSetPatchArgs{...} }
type StatefulSetPatchArrayInput interface {
	pulumi.Input

	ToStatefulSetPatchArrayOutput() StatefulSetPatchArrayOutput
	ToStatefulSetPatchArrayOutputWithContext(context.Context) StatefulSetPatchArrayOutput
}

type StatefulSetPatchArray []StatefulSetPatchInput

func (StatefulSetPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StatefulSetPatch)(nil)).Elem()
}

func (i StatefulSetPatchArray) ToStatefulSetPatchArrayOutput() StatefulSetPatchArrayOutput {
	return i.ToStatefulSetPatchArrayOutputWithContext(context.Background())
}

func (i StatefulSetPatchArray) ToStatefulSetPatchArrayOutputWithContext(ctx context.Context) StatefulSetPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatefulSetPatchArrayOutput)
}

// StatefulSetPatchMapInput is an input type that accepts StatefulSetPatchMap and StatefulSetPatchMapOutput values.
// You can construct a concrete instance of `StatefulSetPatchMapInput` via:
//
//          StatefulSetPatchMap{ "key": StatefulSetPatchArgs{...} }
type StatefulSetPatchMapInput interface {
	pulumi.Input

	ToStatefulSetPatchMapOutput() StatefulSetPatchMapOutput
	ToStatefulSetPatchMapOutputWithContext(context.Context) StatefulSetPatchMapOutput
}

type StatefulSetPatchMap map[string]StatefulSetPatchInput

func (StatefulSetPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StatefulSetPatch)(nil)).Elem()
}

func (i StatefulSetPatchMap) ToStatefulSetPatchMapOutput() StatefulSetPatchMapOutput {
	return i.ToStatefulSetPatchMapOutputWithContext(context.Background())
}

func (i StatefulSetPatchMap) ToStatefulSetPatchMapOutputWithContext(ctx context.Context) StatefulSetPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatefulSetPatchMapOutput)
}

type StatefulSetPatchOutput struct{ *pulumi.OutputState }

func (StatefulSetPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StatefulSetPatch)(nil)).Elem()
}

func (o StatefulSetPatchOutput) ToStatefulSetPatchOutput() StatefulSetPatchOutput {
	return o
}

func (o StatefulSetPatchOutput) ToStatefulSetPatchOutputWithContext(ctx context.Context) StatefulSetPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o StatefulSetPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StatefulSetPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o StatefulSetPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StatefulSetPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o StatefulSetPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *StatefulSetPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Spec defines the desired identities of pods in this set.
func (o StatefulSetPatchOutput) Spec() StatefulSetSpecPatchPtrOutput {
	return o.ApplyT(func(v *StatefulSetPatch) StatefulSetSpecPatchPtrOutput { return v.Spec }).(StatefulSetSpecPatchPtrOutput)
}

// Status is the current status of Pods in this StatefulSet. This data may be out of date by some window of time.
func (o StatefulSetPatchOutput) Status() StatefulSetStatusPatchPtrOutput {
	return o.ApplyT(func(v *StatefulSetPatch) StatefulSetStatusPatchPtrOutput { return v.Status }).(StatefulSetStatusPatchPtrOutput)
}

type StatefulSetPatchArrayOutput struct{ *pulumi.OutputState }

func (StatefulSetPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StatefulSetPatch)(nil)).Elem()
}

func (o StatefulSetPatchArrayOutput) ToStatefulSetPatchArrayOutput() StatefulSetPatchArrayOutput {
	return o
}

func (o StatefulSetPatchArrayOutput) ToStatefulSetPatchArrayOutputWithContext(ctx context.Context) StatefulSetPatchArrayOutput {
	return o
}

func (o StatefulSetPatchArrayOutput) Index(i pulumi.IntInput) StatefulSetPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StatefulSetPatch {
		return vs[0].([]*StatefulSetPatch)[vs[1].(int)]
	}).(StatefulSetPatchOutput)
}

type StatefulSetPatchMapOutput struct{ *pulumi.OutputState }

func (StatefulSetPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StatefulSetPatch)(nil)).Elem()
}

func (o StatefulSetPatchMapOutput) ToStatefulSetPatchMapOutput() StatefulSetPatchMapOutput {
	return o
}

func (o StatefulSetPatchMapOutput) ToStatefulSetPatchMapOutputWithContext(ctx context.Context) StatefulSetPatchMapOutput {
	return o
}

func (o StatefulSetPatchMapOutput) MapIndex(k pulumi.StringInput) StatefulSetPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StatefulSetPatch {
		return vs[0].(map[string]*StatefulSetPatch)[vs[1].(string)]
	}).(StatefulSetPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StatefulSetPatchInput)(nil)).Elem(), &StatefulSetPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatefulSetPatchArrayInput)(nil)).Elem(), StatefulSetPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatefulSetPatchMapInput)(nil)).Elem(), StatefulSetPatchMap{})
	pulumi.RegisterOutputType(StatefulSetPatchOutput{})
	pulumi.RegisterOutputType(StatefulSetPatchArrayOutput{})
	pulumi.RegisterOutputType(StatefulSetPatchMapOutput{})
}
