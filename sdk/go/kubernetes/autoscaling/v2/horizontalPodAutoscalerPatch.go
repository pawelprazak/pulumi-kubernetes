// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler, which automatically manages the replica count of any resource implementing the scale subresource based on the metrics specified.
type HorizontalPodAutoscalerPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec HorizontalPodAutoscalerSpecPtrOutput `pulumi:"spec"`
	// status is the current information about the autoscaler.
	Status HorizontalPodAutoscalerStatusPtrOutput `pulumi:"status"`
}

// NewHorizontalPodAutoscalerPatch registers a new resource with the given unique name, arguments, and options.
func NewHorizontalPodAutoscalerPatch(ctx *pulumi.Context,
	name string, args *HorizontalPodAutoscalerPatchArgs, opts ...pulumi.ResourceOption) (*HorizontalPodAutoscalerPatch, error) {
	if args == nil {
		args = &HorizontalPodAutoscalerPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("autoscaling/v2")
	args.Kind = pulumi.StringPtr("HorizontalPodAutoscaler")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:autoscaling/v1:HorizontalPodAutoscaler"),
		},
		{
			Type: pulumi.String("kubernetes:autoscaling/v2beta1:HorizontalPodAutoscaler"),
		},
		{
			Type: pulumi.String("kubernetes:autoscaling/v2beta2:HorizontalPodAutoscaler"),
		},
	})
	opts = append(opts, aliases)
	var resource HorizontalPodAutoscalerPatch
	err := ctx.RegisterResource("kubernetes:autoscaling/v2:HorizontalPodAutoscalerPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHorizontalPodAutoscalerPatch gets an existing HorizontalPodAutoscalerPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHorizontalPodAutoscalerPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HorizontalPodAutoscalerPatchState, opts ...pulumi.ResourceOption) (*HorizontalPodAutoscalerPatch, error) {
	var resource HorizontalPodAutoscalerPatch
	err := ctx.ReadResource("kubernetes:autoscaling/v2:HorizontalPodAutoscalerPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HorizontalPodAutoscalerPatch resources.
type horizontalPodAutoscalerPatchState struct {
}

type HorizontalPodAutoscalerPatchState struct {
}

func (HorizontalPodAutoscalerPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*horizontalPodAutoscalerPatchState)(nil)).Elem()
}

type horizontalPodAutoscalerPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec *HorizontalPodAutoscalerSpec `pulumi:"spec"`
}

// The set of arguments for constructing a HorizontalPodAutoscalerPatch resource.
type HorizontalPodAutoscalerPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec HorizontalPodAutoscalerSpecPtrInput
}

func (HorizontalPodAutoscalerPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*horizontalPodAutoscalerPatchArgs)(nil)).Elem()
}

type HorizontalPodAutoscalerPatchInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerPatchOutput() HorizontalPodAutoscalerPatchOutput
	ToHorizontalPodAutoscalerPatchOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPatchOutput
}

func (*HorizontalPodAutoscalerPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**HorizontalPodAutoscalerPatch)(nil)).Elem()
}

func (i *HorizontalPodAutoscalerPatch) ToHorizontalPodAutoscalerPatchOutput() HorizontalPodAutoscalerPatchOutput {
	return i.ToHorizontalPodAutoscalerPatchOutputWithContext(context.Background())
}

func (i *HorizontalPodAutoscalerPatch) ToHorizontalPodAutoscalerPatchOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerPatchOutput)
}

// HorizontalPodAutoscalerPatchArrayInput is an input type that accepts HorizontalPodAutoscalerPatchArray and HorizontalPodAutoscalerPatchArrayOutput values.
// You can construct a concrete instance of `HorizontalPodAutoscalerPatchArrayInput` via:
//
//          HorizontalPodAutoscalerPatchArray{ HorizontalPodAutoscalerPatchArgs{...} }
type HorizontalPodAutoscalerPatchArrayInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerPatchArrayOutput() HorizontalPodAutoscalerPatchArrayOutput
	ToHorizontalPodAutoscalerPatchArrayOutputWithContext(context.Context) HorizontalPodAutoscalerPatchArrayOutput
}

type HorizontalPodAutoscalerPatchArray []HorizontalPodAutoscalerPatchInput

func (HorizontalPodAutoscalerPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HorizontalPodAutoscalerPatch)(nil)).Elem()
}

func (i HorizontalPodAutoscalerPatchArray) ToHorizontalPodAutoscalerPatchArrayOutput() HorizontalPodAutoscalerPatchArrayOutput {
	return i.ToHorizontalPodAutoscalerPatchArrayOutputWithContext(context.Background())
}

func (i HorizontalPodAutoscalerPatchArray) ToHorizontalPodAutoscalerPatchArrayOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerPatchArrayOutput)
}

// HorizontalPodAutoscalerPatchMapInput is an input type that accepts HorizontalPodAutoscalerPatchMap and HorizontalPodAutoscalerPatchMapOutput values.
// You can construct a concrete instance of `HorizontalPodAutoscalerPatchMapInput` via:
//
//          HorizontalPodAutoscalerPatchMap{ "key": HorizontalPodAutoscalerPatchArgs{...} }
type HorizontalPodAutoscalerPatchMapInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerPatchMapOutput() HorizontalPodAutoscalerPatchMapOutput
	ToHorizontalPodAutoscalerPatchMapOutputWithContext(context.Context) HorizontalPodAutoscalerPatchMapOutput
}

type HorizontalPodAutoscalerPatchMap map[string]HorizontalPodAutoscalerPatchInput

func (HorizontalPodAutoscalerPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HorizontalPodAutoscalerPatch)(nil)).Elem()
}

func (i HorizontalPodAutoscalerPatchMap) ToHorizontalPodAutoscalerPatchMapOutput() HorizontalPodAutoscalerPatchMapOutput {
	return i.ToHorizontalPodAutoscalerPatchMapOutputWithContext(context.Background())
}

func (i HorizontalPodAutoscalerPatchMap) ToHorizontalPodAutoscalerPatchMapOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerPatchMapOutput)
}

type HorizontalPodAutoscalerPatchOutput struct{ *pulumi.OutputState }

func (HorizontalPodAutoscalerPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HorizontalPodAutoscalerPatch)(nil)).Elem()
}

func (o HorizontalPodAutoscalerPatchOutput) ToHorizontalPodAutoscalerPatchOutput() HorizontalPodAutoscalerPatchOutput {
	return o
}

func (o HorizontalPodAutoscalerPatchOutput) ToHorizontalPodAutoscalerPatchOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o HorizontalPodAutoscalerPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HorizontalPodAutoscalerPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o HorizontalPodAutoscalerPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HorizontalPodAutoscalerPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o HorizontalPodAutoscalerPatchOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *HorizontalPodAutoscalerPatch) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
func (o HorizontalPodAutoscalerPatchOutput) Spec() HorizontalPodAutoscalerSpecPtrOutput {
	return o.ApplyT(func(v *HorizontalPodAutoscalerPatch) HorizontalPodAutoscalerSpecPtrOutput { return v.Spec }).(HorizontalPodAutoscalerSpecPtrOutput)
}

// status is the current information about the autoscaler.
func (o HorizontalPodAutoscalerPatchOutput) Status() HorizontalPodAutoscalerStatusPtrOutput {
	return o.ApplyT(func(v *HorizontalPodAutoscalerPatch) HorizontalPodAutoscalerStatusPtrOutput { return v.Status }).(HorizontalPodAutoscalerStatusPtrOutput)
}

type HorizontalPodAutoscalerPatchArrayOutput struct{ *pulumi.OutputState }

func (HorizontalPodAutoscalerPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HorizontalPodAutoscalerPatch)(nil)).Elem()
}

func (o HorizontalPodAutoscalerPatchArrayOutput) ToHorizontalPodAutoscalerPatchArrayOutput() HorizontalPodAutoscalerPatchArrayOutput {
	return o
}

func (o HorizontalPodAutoscalerPatchArrayOutput) ToHorizontalPodAutoscalerPatchArrayOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPatchArrayOutput {
	return o
}

func (o HorizontalPodAutoscalerPatchArrayOutput) Index(i pulumi.IntInput) HorizontalPodAutoscalerPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HorizontalPodAutoscalerPatch {
		return vs[0].([]*HorizontalPodAutoscalerPatch)[vs[1].(int)]
	}).(HorizontalPodAutoscalerPatchOutput)
}

type HorizontalPodAutoscalerPatchMapOutput struct{ *pulumi.OutputState }

func (HorizontalPodAutoscalerPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HorizontalPodAutoscalerPatch)(nil)).Elem()
}

func (o HorizontalPodAutoscalerPatchMapOutput) ToHorizontalPodAutoscalerPatchMapOutput() HorizontalPodAutoscalerPatchMapOutput {
	return o
}

func (o HorizontalPodAutoscalerPatchMapOutput) ToHorizontalPodAutoscalerPatchMapOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPatchMapOutput {
	return o
}

func (o HorizontalPodAutoscalerPatchMapOutput) MapIndex(k pulumi.StringInput) HorizontalPodAutoscalerPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HorizontalPodAutoscalerPatch {
		return vs[0].(map[string]*HorizontalPodAutoscalerPatch)[vs[1].(string)]
	}).(HorizontalPodAutoscalerPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HorizontalPodAutoscalerPatchInput)(nil)).Elem(), &HorizontalPodAutoscalerPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*HorizontalPodAutoscalerPatchArrayInput)(nil)).Elem(), HorizontalPodAutoscalerPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HorizontalPodAutoscalerPatchMapInput)(nil)).Elem(), HorizontalPodAutoscalerPatchMap{})
	pulumi.RegisterOutputType(HorizontalPodAutoscalerPatchOutput{})
	pulumi.RegisterOutputType(HorizontalPodAutoscalerPatchArrayOutput{})
	pulumi.RegisterOutputType(HorizontalPodAutoscalerPatchMapOutput{})
}
