// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// PodPreset is a policy resource that defines additional runtime requirements for a Pod.
type PodPresetPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	Spec     PodPresetSpecPtrOutput     `pulumi:"spec"`
}

// NewPodPresetPatch registers a new resource with the given unique name, arguments, and options.
func NewPodPresetPatch(ctx *pulumi.Context,
	name string, args *PodPresetPatchArgs, opts ...pulumi.ResourceOption) (*PodPresetPatch, error) {
	if args == nil {
		args = &PodPresetPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("settings.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("PodPreset")
	var resource PodPresetPatch
	err := ctx.RegisterResource("kubernetes:settings.k8s.io/v1alpha1:PodPresetPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPodPresetPatch gets an existing PodPresetPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPodPresetPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PodPresetPatchState, opts ...pulumi.ResourceOption) (*PodPresetPatch, error) {
	var resource PodPresetPatch
	err := ctx.ReadResource("kubernetes:settings.k8s.io/v1alpha1:PodPresetPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PodPresetPatch resources.
type podPresetPatchState struct {
}

type PodPresetPatchState struct {
}

func (PodPresetPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*podPresetPatchState)(nil)).Elem()
}

type podPresetPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     *PodPresetSpec     `pulumi:"spec"`
}

// The set of arguments for constructing a PodPresetPatch resource.
type PodPresetPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	Spec     PodPresetSpecPtrInput
}

func (PodPresetPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*podPresetPatchArgs)(nil)).Elem()
}

type PodPresetPatchInput interface {
	pulumi.Input

	ToPodPresetPatchOutput() PodPresetPatchOutput
	ToPodPresetPatchOutputWithContext(ctx context.Context) PodPresetPatchOutput
}

func (*PodPresetPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**PodPresetPatch)(nil)).Elem()
}

func (i *PodPresetPatch) ToPodPresetPatchOutput() PodPresetPatchOutput {
	return i.ToPodPresetPatchOutputWithContext(context.Background())
}

func (i *PodPresetPatch) ToPodPresetPatchOutputWithContext(ctx context.Context) PodPresetPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodPresetPatchOutput)
}

// PodPresetPatchArrayInput is an input type that accepts PodPresetPatchArray and PodPresetPatchArrayOutput values.
// You can construct a concrete instance of `PodPresetPatchArrayInput` via:
//
//          PodPresetPatchArray{ PodPresetPatchArgs{...} }
type PodPresetPatchArrayInput interface {
	pulumi.Input

	ToPodPresetPatchArrayOutput() PodPresetPatchArrayOutput
	ToPodPresetPatchArrayOutputWithContext(context.Context) PodPresetPatchArrayOutput
}

type PodPresetPatchArray []PodPresetPatchInput

func (PodPresetPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PodPresetPatch)(nil)).Elem()
}

func (i PodPresetPatchArray) ToPodPresetPatchArrayOutput() PodPresetPatchArrayOutput {
	return i.ToPodPresetPatchArrayOutputWithContext(context.Background())
}

func (i PodPresetPatchArray) ToPodPresetPatchArrayOutputWithContext(ctx context.Context) PodPresetPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodPresetPatchArrayOutput)
}

// PodPresetPatchMapInput is an input type that accepts PodPresetPatchMap and PodPresetPatchMapOutput values.
// You can construct a concrete instance of `PodPresetPatchMapInput` via:
//
//          PodPresetPatchMap{ "key": PodPresetPatchArgs{...} }
type PodPresetPatchMapInput interface {
	pulumi.Input

	ToPodPresetPatchMapOutput() PodPresetPatchMapOutput
	ToPodPresetPatchMapOutputWithContext(context.Context) PodPresetPatchMapOutput
}

type PodPresetPatchMap map[string]PodPresetPatchInput

func (PodPresetPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PodPresetPatch)(nil)).Elem()
}

func (i PodPresetPatchMap) ToPodPresetPatchMapOutput() PodPresetPatchMapOutput {
	return i.ToPodPresetPatchMapOutputWithContext(context.Background())
}

func (i PodPresetPatchMap) ToPodPresetPatchMapOutputWithContext(ctx context.Context) PodPresetPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodPresetPatchMapOutput)
}

type PodPresetPatchOutput struct{ *pulumi.OutputState }

func (PodPresetPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PodPresetPatch)(nil)).Elem()
}

func (o PodPresetPatchOutput) ToPodPresetPatchOutput() PodPresetPatchOutput {
	return o
}

func (o PodPresetPatchOutput) ToPodPresetPatchOutputWithContext(ctx context.Context) PodPresetPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PodPresetPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PodPresetPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PodPresetPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PodPresetPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o PodPresetPatchOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *PodPresetPatch) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

func (o PodPresetPatchOutput) Spec() PodPresetSpecPtrOutput {
	return o.ApplyT(func(v *PodPresetPatch) PodPresetSpecPtrOutput { return v.Spec }).(PodPresetSpecPtrOutput)
}

type PodPresetPatchArrayOutput struct{ *pulumi.OutputState }

func (PodPresetPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PodPresetPatch)(nil)).Elem()
}

func (o PodPresetPatchArrayOutput) ToPodPresetPatchArrayOutput() PodPresetPatchArrayOutput {
	return o
}

func (o PodPresetPatchArrayOutput) ToPodPresetPatchArrayOutputWithContext(ctx context.Context) PodPresetPatchArrayOutput {
	return o
}

func (o PodPresetPatchArrayOutput) Index(i pulumi.IntInput) PodPresetPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PodPresetPatch {
		return vs[0].([]*PodPresetPatch)[vs[1].(int)]
	}).(PodPresetPatchOutput)
}

type PodPresetPatchMapOutput struct{ *pulumi.OutputState }

func (PodPresetPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PodPresetPatch)(nil)).Elem()
}

func (o PodPresetPatchMapOutput) ToPodPresetPatchMapOutput() PodPresetPatchMapOutput {
	return o
}

func (o PodPresetPatchMapOutput) ToPodPresetPatchMapOutputWithContext(ctx context.Context) PodPresetPatchMapOutput {
	return o
}

func (o PodPresetPatchMapOutput) MapIndex(k pulumi.StringInput) PodPresetPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PodPresetPatch {
		return vs[0].(map[string]*PodPresetPatch)[vs[1].(string)]
	}).(PodPresetPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PodPresetPatchInput)(nil)).Elem(), &PodPresetPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodPresetPatchArrayInput)(nil)).Elem(), PodPresetPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodPresetPatchMapInput)(nil)).Elem(), PodPresetPatchMap{})
	pulumi.RegisterOutputType(PodPresetPatchOutput{})
	pulumi.RegisterOutputType(PodPresetPatchArrayOutput{})
	pulumi.RegisterOutputType(PodPresetPatchMapOutput{})
}
