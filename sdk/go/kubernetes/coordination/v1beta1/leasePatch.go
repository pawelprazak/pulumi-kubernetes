// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lease defines a lease concept.
type LeasePatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec LeaseSpecPatchPtrOutput `pulumi:"spec"`
}

// NewLeasePatch registers a new resource with the given unique name, arguments, and options.
func NewLeasePatch(ctx *pulumi.Context,
	name string, args *LeasePatchArgs, opts ...pulumi.ResourceOption) (*LeasePatch, error) {
	if args == nil {
		args = &LeasePatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("coordination.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("Lease")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:coordination.k8s.io/v1:LeasePatch"),
		},
	})
	opts = append(opts, aliases)
	var resource LeasePatch
	err := ctx.RegisterResource("kubernetes:coordination.k8s.io/v1beta1:LeasePatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLeasePatch gets an existing LeasePatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLeasePatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LeasePatchState, opts ...pulumi.ResourceOption) (*LeasePatch, error) {
	var resource LeasePatch
	err := ctx.ReadResource("kubernetes:coordination.k8s.io/v1beta1:LeasePatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LeasePatch resources.
type leasePatchState struct {
}

type LeasePatchState struct {
}

func (LeasePatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*leasePatchState)(nil)).Elem()
}

type leasePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *LeaseSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a LeasePatch resource.
type LeasePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// Specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec LeaseSpecPatchPtrInput
}

func (LeasePatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*leasePatchArgs)(nil)).Elem()
}

type LeasePatchInput interface {
	pulumi.Input

	ToLeasePatchOutput() LeasePatchOutput
	ToLeasePatchOutputWithContext(ctx context.Context) LeasePatchOutput
}

func (*LeasePatch) ElementType() reflect.Type {
	return reflect.TypeOf((**LeasePatch)(nil)).Elem()
}

func (i *LeasePatch) ToLeasePatchOutput() LeasePatchOutput {
	return i.ToLeasePatchOutputWithContext(context.Background())
}

func (i *LeasePatch) ToLeasePatchOutputWithContext(ctx context.Context) LeasePatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeasePatchOutput)
}

// LeasePatchArrayInput is an input type that accepts LeasePatchArray and LeasePatchArrayOutput values.
// You can construct a concrete instance of `LeasePatchArrayInput` via:
//
//          LeasePatchArray{ LeasePatchArgs{...} }
type LeasePatchArrayInput interface {
	pulumi.Input

	ToLeasePatchArrayOutput() LeasePatchArrayOutput
	ToLeasePatchArrayOutputWithContext(context.Context) LeasePatchArrayOutput
}

type LeasePatchArray []LeasePatchInput

func (LeasePatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LeasePatch)(nil)).Elem()
}

func (i LeasePatchArray) ToLeasePatchArrayOutput() LeasePatchArrayOutput {
	return i.ToLeasePatchArrayOutputWithContext(context.Background())
}

func (i LeasePatchArray) ToLeasePatchArrayOutputWithContext(ctx context.Context) LeasePatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeasePatchArrayOutput)
}

// LeasePatchMapInput is an input type that accepts LeasePatchMap and LeasePatchMapOutput values.
// You can construct a concrete instance of `LeasePatchMapInput` via:
//
//          LeasePatchMap{ "key": LeasePatchArgs{...} }
type LeasePatchMapInput interface {
	pulumi.Input

	ToLeasePatchMapOutput() LeasePatchMapOutput
	ToLeasePatchMapOutputWithContext(context.Context) LeasePatchMapOutput
}

type LeasePatchMap map[string]LeasePatchInput

func (LeasePatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LeasePatch)(nil)).Elem()
}

func (i LeasePatchMap) ToLeasePatchMapOutput() LeasePatchMapOutput {
	return i.ToLeasePatchMapOutputWithContext(context.Background())
}

func (i LeasePatchMap) ToLeasePatchMapOutputWithContext(ctx context.Context) LeasePatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeasePatchMapOutput)
}

type LeasePatchOutput struct{ *pulumi.OutputState }

func (LeasePatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LeasePatch)(nil)).Elem()
}

func (o LeasePatchOutput) ToLeasePatchOutput() LeasePatchOutput {
	return o
}

func (o LeasePatchOutput) ToLeasePatchOutputWithContext(ctx context.Context) LeasePatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o LeasePatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LeasePatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o LeasePatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LeasePatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o LeasePatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *LeasePatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o LeasePatchOutput) Spec() LeaseSpecPatchPtrOutput {
	return o.ApplyT(func(v *LeasePatch) LeaseSpecPatchPtrOutput { return v.Spec }).(LeaseSpecPatchPtrOutput)
}

type LeasePatchArrayOutput struct{ *pulumi.OutputState }

func (LeasePatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LeasePatch)(nil)).Elem()
}

func (o LeasePatchArrayOutput) ToLeasePatchArrayOutput() LeasePatchArrayOutput {
	return o
}

func (o LeasePatchArrayOutput) ToLeasePatchArrayOutputWithContext(ctx context.Context) LeasePatchArrayOutput {
	return o
}

func (o LeasePatchArrayOutput) Index(i pulumi.IntInput) LeasePatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LeasePatch {
		return vs[0].([]*LeasePatch)[vs[1].(int)]
	}).(LeasePatchOutput)
}

type LeasePatchMapOutput struct{ *pulumi.OutputState }

func (LeasePatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LeasePatch)(nil)).Elem()
}

func (o LeasePatchMapOutput) ToLeasePatchMapOutput() LeasePatchMapOutput {
	return o
}

func (o LeasePatchMapOutput) ToLeasePatchMapOutputWithContext(ctx context.Context) LeasePatchMapOutput {
	return o
}

func (o LeasePatchMapOutput) MapIndex(k pulumi.StringInput) LeasePatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LeasePatch {
		return vs[0].(map[string]*LeasePatch)[vs[1].(string)]
	}).(LeasePatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LeasePatchInput)(nil)).Elem(), &LeasePatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*LeasePatchArrayInput)(nil)).Elem(), LeasePatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LeasePatchMapInput)(nil)).Elem(), LeasePatchMap{})
	pulumi.RegisterOutputType(LeasePatchOutput{})
	pulumi.RegisterOutputType(LeasePatchArrayOutput{})
	pulumi.RegisterOutputType(LeasePatchMapOutput{})
}
