// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Status is a return value for calls that don't return other objects.
type StatusPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Suggested HTTP return code for this status, 0 if not set.
	Code pulumi.IntPtrOutput `pulumi:"code"`
	// Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.
	Details StatusDetailsPtrOutput `pulumi:"details"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// A human-readable description of the status of this operation.
	Message pulumi.StringPtrOutput `pulumi:"message"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata ListMetaPtrOutput `pulumi:"metadata"`
	// A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
	Reason pulumi.StringPtrOutput `pulumi:"reason"`
	// Status of the operation. One of: "Success" or "Failure". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status pulumi.StringPtrOutput `pulumi:"status"`
}

// NewStatusPatch registers a new resource with the given unique name, arguments, and options.
func NewStatusPatch(ctx *pulumi.Context,
	name string, args *StatusPatchArgs, opts ...pulumi.ResourceOption) (*StatusPatch, error) {
	if args == nil {
		args = &StatusPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("Status")
	var resource StatusPatch
	err := ctx.RegisterResource("kubernetes:meta/v1:StatusPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStatusPatch gets an existing StatusPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStatusPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StatusPatchState, opts ...pulumi.ResourceOption) (*StatusPatch, error) {
	var resource StatusPatch
	err := ctx.ReadResource("kubernetes:meta/v1:StatusPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StatusPatch resources.
type statusPatchState struct {
}

type StatusPatchState struct {
}

func (StatusPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*statusPatchState)(nil)).Elem()
}

type statusPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Suggested HTTP return code for this status, 0 if not set.
	Code *int `pulumi:"code"`
	// Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.
	Details *StatusDetails `pulumi:"details"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// A human-readable description of the status of this operation.
	Message *string `pulumi:"message"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *ListMeta `pulumi:"metadata"`
	// A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
	Reason *string `pulumi:"reason"`
}

// The set of arguments for constructing a StatusPatch resource.
type StatusPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Suggested HTTP return code for this status, 0 if not set.
	Code pulumi.IntPtrInput
	// Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.
	Details StatusDetailsPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// A human-readable description of the status of this operation.
	Message pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata ListMetaPtrInput
	// A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
	Reason pulumi.StringPtrInput
}

func (StatusPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*statusPatchArgs)(nil)).Elem()
}

type StatusPatchInput interface {
	pulumi.Input

	ToStatusPatchOutput() StatusPatchOutput
	ToStatusPatchOutputWithContext(ctx context.Context) StatusPatchOutput
}

func (*StatusPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**StatusPatch)(nil)).Elem()
}

func (i *StatusPatch) ToStatusPatchOutput() StatusPatchOutput {
	return i.ToStatusPatchOutputWithContext(context.Background())
}

func (i *StatusPatch) ToStatusPatchOutputWithContext(ctx context.Context) StatusPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusPatchOutput)
}

// StatusPatchArrayInput is an input type that accepts StatusPatchArray and StatusPatchArrayOutput values.
// You can construct a concrete instance of `StatusPatchArrayInput` via:
//
//          StatusPatchArray{ StatusPatchArgs{...} }
type StatusPatchArrayInput interface {
	pulumi.Input

	ToStatusPatchArrayOutput() StatusPatchArrayOutput
	ToStatusPatchArrayOutputWithContext(context.Context) StatusPatchArrayOutput
}

type StatusPatchArray []StatusPatchInput

func (StatusPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StatusPatch)(nil)).Elem()
}

func (i StatusPatchArray) ToStatusPatchArrayOutput() StatusPatchArrayOutput {
	return i.ToStatusPatchArrayOutputWithContext(context.Background())
}

func (i StatusPatchArray) ToStatusPatchArrayOutputWithContext(ctx context.Context) StatusPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusPatchArrayOutput)
}

// StatusPatchMapInput is an input type that accepts StatusPatchMap and StatusPatchMapOutput values.
// You can construct a concrete instance of `StatusPatchMapInput` via:
//
//          StatusPatchMap{ "key": StatusPatchArgs{...} }
type StatusPatchMapInput interface {
	pulumi.Input

	ToStatusPatchMapOutput() StatusPatchMapOutput
	ToStatusPatchMapOutputWithContext(context.Context) StatusPatchMapOutput
}

type StatusPatchMap map[string]StatusPatchInput

func (StatusPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StatusPatch)(nil)).Elem()
}

func (i StatusPatchMap) ToStatusPatchMapOutput() StatusPatchMapOutput {
	return i.ToStatusPatchMapOutputWithContext(context.Background())
}

func (i StatusPatchMap) ToStatusPatchMapOutputWithContext(ctx context.Context) StatusPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusPatchMapOutput)
}

type StatusPatchOutput struct{ *pulumi.OutputState }

func (StatusPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StatusPatch)(nil)).Elem()
}

func (o StatusPatchOutput) ToStatusPatchOutput() StatusPatchOutput {
	return o
}

func (o StatusPatchOutput) ToStatusPatchOutputWithContext(ctx context.Context) StatusPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o StatusPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StatusPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Suggested HTTP return code for this status, 0 if not set.
func (o StatusPatchOutput) Code() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StatusPatch) pulumi.IntPtrOutput { return v.Code }).(pulumi.IntPtrOutput)
}

// Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.
func (o StatusPatchOutput) Details() StatusDetailsPtrOutput {
	return o.ApplyT(func(v *StatusPatch) StatusDetailsPtrOutput { return v.Details }).(StatusDetailsPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o StatusPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StatusPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// A human-readable description of the status of this operation.
func (o StatusPatchOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StatusPatch) pulumi.StringPtrOutput { return v.Message }).(pulumi.StringPtrOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o StatusPatchOutput) Metadata() ListMetaPtrOutput {
	return o.ApplyT(func(v *StatusPatch) ListMetaPtrOutput { return v.Metadata }).(ListMetaPtrOutput)
}

// A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
func (o StatusPatchOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StatusPatch) pulumi.StringPtrOutput { return v.Reason }).(pulumi.StringPtrOutput)
}

// Status of the operation. One of: "Success" or "Failure". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o StatusPatchOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StatusPatch) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

type StatusPatchArrayOutput struct{ *pulumi.OutputState }

func (StatusPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StatusPatch)(nil)).Elem()
}

func (o StatusPatchArrayOutput) ToStatusPatchArrayOutput() StatusPatchArrayOutput {
	return o
}

func (o StatusPatchArrayOutput) ToStatusPatchArrayOutputWithContext(ctx context.Context) StatusPatchArrayOutput {
	return o
}

func (o StatusPatchArrayOutput) Index(i pulumi.IntInput) StatusPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StatusPatch {
		return vs[0].([]*StatusPatch)[vs[1].(int)]
	}).(StatusPatchOutput)
}

type StatusPatchMapOutput struct{ *pulumi.OutputState }

func (StatusPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StatusPatch)(nil)).Elem()
}

func (o StatusPatchMapOutput) ToStatusPatchMapOutput() StatusPatchMapOutput {
	return o
}

func (o StatusPatchMapOutput) ToStatusPatchMapOutputWithContext(ctx context.Context) StatusPatchMapOutput {
	return o
}

func (o StatusPatchMapOutput) MapIndex(k pulumi.StringInput) StatusPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StatusPatch {
		return vs[0].(map[string]*StatusPatch)[vs[1].(string)]
	}).(StatusPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StatusPatchInput)(nil)).Elem(), &StatusPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatusPatchArrayInput)(nil)).Elem(), StatusPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatusPatchMapInput)(nil)).Elem(), StatusPatchMap{})
	pulumi.RegisterOutputType(StatusPatchOutput{})
	pulumi.RegisterOutputType(StatusPatchArrayOutput{})
	pulumi.RegisterOutputType(StatusPatchMapOutput{})
}
