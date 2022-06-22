// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SelfSubjectRulesReview enumerates the set of actions the current user can perform within a namespace. The returned list of actions may be incomplete depending on the server's authorization mode, and any errors experienced during the evaluation. SelfSubjectRulesReview should be used by UIs to show/hide actions, or to quickly let an end user reason about their permissions. It should NOT Be used by external systems to drive authorization decisions as this raises confused deputy, cache lifetime/revocation, and correctness concerns. SubjectAccessReview, and LocalAccessReview are the correct way to defer authorization decisions to the API server.
type SelfSubjectRulesReviewPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec holds information about the request being evaluated.
	Spec SelfSubjectRulesReviewSpecOutput `pulumi:"spec"`
	// Status is filled in by the server and indicates the set of actions a user can perform.
	Status SubjectRulesReviewStatusPtrOutput `pulumi:"status"`
}

// NewSelfSubjectRulesReviewPatch registers a new resource with the given unique name, arguments, and options.
func NewSelfSubjectRulesReviewPatch(ctx *pulumi.Context,
	name string, args *SelfSubjectRulesReviewPatchArgs, opts ...pulumi.ResourceOption) (*SelfSubjectRulesReviewPatch, error) {
	if args == nil {
		args = &SelfSubjectRulesReviewPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("authorization.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("SelfSubjectRulesReview")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:authorization.k8s.io/v1:SelfSubjectRulesReview"),
		},
	})
	opts = append(opts, aliases)
	var resource SelfSubjectRulesReviewPatch
	err := ctx.RegisterResource("kubernetes:authorization.k8s.io/v1beta1:SelfSubjectRulesReviewPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSelfSubjectRulesReviewPatch gets an existing SelfSubjectRulesReviewPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSelfSubjectRulesReviewPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SelfSubjectRulesReviewPatchState, opts ...pulumi.ResourceOption) (*SelfSubjectRulesReviewPatch, error) {
	var resource SelfSubjectRulesReviewPatch
	err := ctx.ReadResource("kubernetes:authorization.k8s.io/v1beta1:SelfSubjectRulesReviewPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SelfSubjectRulesReviewPatch resources.
type selfSubjectRulesReviewPatchState struct {
}

type SelfSubjectRulesReviewPatchState struct {
}

func (SelfSubjectRulesReviewPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*selfSubjectRulesReviewPatchState)(nil)).Elem()
}

type selfSubjectRulesReviewPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec holds information about the request being evaluated.
	Spec *SelfSubjectRulesReviewSpec `pulumi:"spec"`
}

// The set of arguments for constructing a SelfSubjectRulesReviewPatch resource.
type SelfSubjectRulesReviewPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec holds information about the request being evaluated.
	Spec SelfSubjectRulesReviewSpecPtrInput
}

func (SelfSubjectRulesReviewPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*selfSubjectRulesReviewPatchArgs)(nil)).Elem()
}

type SelfSubjectRulesReviewPatchInput interface {
	pulumi.Input

	ToSelfSubjectRulesReviewPatchOutput() SelfSubjectRulesReviewPatchOutput
	ToSelfSubjectRulesReviewPatchOutputWithContext(ctx context.Context) SelfSubjectRulesReviewPatchOutput
}

func (*SelfSubjectRulesReviewPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**SelfSubjectRulesReviewPatch)(nil)).Elem()
}

func (i *SelfSubjectRulesReviewPatch) ToSelfSubjectRulesReviewPatchOutput() SelfSubjectRulesReviewPatchOutput {
	return i.ToSelfSubjectRulesReviewPatchOutputWithContext(context.Background())
}

func (i *SelfSubjectRulesReviewPatch) ToSelfSubjectRulesReviewPatchOutputWithContext(ctx context.Context) SelfSubjectRulesReviewPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectRulesReviewPatchOutput)
}

// SelfSubjectRulesReviewPatchArrayInput is an input type that accepts SelfSubjectRulesReviewPatchArray and SelfSubjectRulesReviewPatchArrayOutput values.
// You can construct a concrete instance of `SelfSubjectRulesReviewPatchArrayInput` via:
//
//          SelfSubjectRulesReviewPatchArray{ SelfSubjectRulesReviewPatchArgs{...} }
type SelfSubjectRulesReviewPatchArrayInput interface {
	pulumi.Input

	ToSelfSubjectRulesReviewPatchArrayOutput() SelfSubjectRulesReviewPatchArrayOutput
	ToSelfSubjectRulesReviewPatchArrayOutputWithContext(context.Context) SelfSubjectRulesReviewPatchArrayOutput
}

type SelfSubjectRulesReviewPatchArray []SelfSubjectRulesReviewPatchInput

func (SelfSubjectRulesReviewPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SelfSubjectRulesReviewPatch)(nil)).Elem()
}

func (i SelfSubjectRulesReviewPatchArray) ToSelfSubjectRulesReviewPatchArrayOutput() SelfSubjectRulesReviewPatchArrayOutput {
	return i.ToSelfSubjectRulesReviewPatchArrayOutputWithContext(context.Background())
}

func (i SelfSubjectRulesReviewPatchArray) ToSelfSubjectRulesReviewPatchArrayOutputWithContext(ctx context.Context) SelfSubjectRulesReviewPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectRulesReviewPatchArrayOutput)
}

// SelfSubjectRulesReviewPatchMapInput is an input type that accepts SelfSubjectRulesReviewPatchMap and SelfSubjectRulesReviewPatchMapOutput values.
// You can construct a concrete instance of `SelfSubjectRulesReviewPatchMapInput` via:
//
//          SelfSubjectRulesReviewPatchMap{ "key": SelfSubjectRulesReviewPatchArgs{...} }
type SelfSubjectRulesReviewPatchMapInput interface {
	pulumi.Input

	ToSelfSubjectRulesReviewPatchMapOutput() SelfSubjectRulesReviewPatchMapOutput
	ToSelfSubjectRulesReviewPatchMapOutputWithContext(context.Context) SelfSubjectRulesReviewPatchMapOutput
}

type SelfSubjectRulesReviewPatchMap map[string]SelfSubjectRulesReviewPatchInput

func (SelfSubjectRulesReviewPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SelfSubjectRulesReviewPatch)(nil)).Elem()
}

func (i SelfSubjectRulesReviewPatchMap) ToSelfSubjectRulesReviewPatchMapOutput() SelfSubjectRulesReviewPatchMapOutput {
	return i.ToSelfSubjectRulesReviewPatchMapOutputWithContext(context.Background())
}

func (i SelfSubjectRulesReviewPatchMap) ToSelfSubjectRulesReviewPatchMapOutputWithContext(ctx context.Context) SelfSubjectRulesReviewPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectRulesReviewPatchMapOutput)
}

type SelfSubjectRulesReviewPatchOutput struct{ *pulumi.OutputState }

func (SelfSubjectRulesReviewPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SelfSubjectRulesReviewPatch)(nil)).Elem()
}

func (o SelfSubjectRulesReviewPatchOutput) ToSelfSubjectRulesReviewPatchOutput() SelfSubjectRulesReviewPatchOutput {
	return o
}

func (o SelfSubjectRulesReviewPatchOutput) ToSelfSubjectRulesReviewPatchOutputWithContext(ctx context.Context) SelfSubjectRulesReviewPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o SelfSubjectRulesReviewPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SelfSubjectRulesReviewPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o SelfSubjectRulesReviewPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SelfSubjectRulesReviewPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SelfSubjectRulesReviewPatchOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *SelfSubjectRulesReviewPatch) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Spec holds information about the request being evaluated.
func (o SelfSubjectRulesReviewPatchOutput) Spec() SelfSubjectRulesReviewSpecOutput {
	return o.ApplyT(func(v *SelfSubjectRulesReviewPatch) SelfSubjectRulesReviewSpecOutput { return v.Spec }).(SelfSubjectRulesReviewSpecOutput)
}

// Status is filled in by the server and indicates the set of actions a user can perform.
func (o SelfSubjectRulesReviewPatchOutput) Status() SubjectRulesReviewStatusPtrOutput {
	return o.ApplyT(func(v *SelfSubjectRulesReviewPatch) SubjectRulesReviewStatusPtrOutput { return v.Status }).(SubjectRulesReviewStatusPtrOutput)
}

type SelfSubjectRulesReviewPatchArrayOutput struct{ *pulumi.OutputState }

func (SelfSubjectRulesReviewPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SelfSubjectRulesReviewPatch)(nil)).Elem()
}

func (o SelfSubjectRulesReviewPatchArrayOutput) ToSelfSubjectRulesReviewPatchArrayOutput() SelfSubjectRulesReviewPatchArrayOutput {
	return o
}

func (o SelfSubjectRulesReviewPatchArrayOutput) ToSelfSubjectRulesReviewPatchArrayOutputWithContext(ctx context.Context) SelfSubjectRulesReviewPatchArrayOutput {
	return o
}

func (o SelfSubjectRulesReviewPatchArrayOutput) Index(i pulumi.IntInput) SelfSubjectRulesReviewPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SelfSubjectRulesReviewPatch {
		return vs[0].([]*SelfSubjectRulesReviewPatch)[vs[1].(int)]
	}).(SelfSubjectRulesReviewPatchOutput)
}

type SelfSubjectRulesReviewPatchMapOutput struct{ *pulumi.OutputState }

func (SelfSubjectRulesReviewPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SelfSubjectRulesReviewPatch)(nil)).Elem()
}

func (o SelfSubjectRulesReviewPatchMapOutput) ToSelfSubjectRulesReviewPatchMapOutput() SelfSubjectRulesReviewPatchMapOutput {
	return o
}

func (o SelfSubjectRulesReviewPatchMapOutput) ToSelfSubjectRulesReviewPatchMapOutputWithContext(ctx context.Context) SelfSubjectRulesReviewPatchMapOutput {
	return o
}

func (o SelfSubjectRulesReviewPatchMapOutput) MapIndex(k pulumi.StringInput) SelfSubjectRulesReviewPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SelfSubjectRulesReviewPatch {
		return vs[0].(map[string]*SelfSubjectRulesReviewPatch)[vs[1].(string)]
	}).(SelfSubjectRulesReviewPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSubjectRulesReviewPatchInput)(nil)).Elem(), &SelfSubjectRulesReviewPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSubjectRulesReviewPatchArrayInput)(nil)).Elem(), SelfSubjectRulesReviewPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSubjectRulesReviewPatchMapInput)(nil)).Elem(), SelfSubjectRulesReviewPatchMap{})
	pulumi.RegisterOutputType(SelfSubjectRulesReviewPatchOutput{})
	pulumi.RegisterOutputType(SelfSubjectRulesReviewPatchArrayOutput{})
	pulumi.RegisterOutputType(SelfSubjectRulesReviewPatchMapOutput{})
}
