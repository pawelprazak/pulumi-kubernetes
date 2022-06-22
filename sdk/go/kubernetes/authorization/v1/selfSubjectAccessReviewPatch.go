// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SelfSubjectAccessReview checks whether or the current user can perform an action.  Not filling in a spec.namespace means "in all namespaces".  Self is a special case, because users should always be able to check whether they can perform an action
type SelfSubjectAccessReviewPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec holds information about the request being evaluated.  user and groups must be empty
	Spec SelfSubjectAccessReviewSpecOutput `pulumi:"spec"`
	// Status is filled in by the server and indicates whether the request is allowed or not
	Status SubjectAccessReviewStatusPtrOutput `pulumi:"status"`
}

// NewSelfSubjectAccessReviewPatch registers a new resource with the given unique name, arguments, and options.
func NewSelfSubjectAccessReviewPatch(ctx *pulumi.Context,
	name string, args *SelfSubjectAccessReviewPatchArgs, opts ...pulumi.ResourceOption) (*SelfSubjectAccessReviewPatch, error) {
	if args == nil {
		args = &SelfSubjectAccessReviewPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("authorization.k8s.io/v1")
	args.Kind = pulumi.StringPtr("SelfSubjectAccessReview")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:authorization.k8s.io/v1beta1:SelfSubjectAccessReview"),
		},
	})
	opts = append(opts, aliases)
	var resource SelfSubjectAccessReviewPatch
	err := ctx.RegisterResource("kubernetes:authorization.k8s.io/v1:SelfSubjectAccessReviewPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSelfSubjectAccessReviewPatch gets an existing SelfSubjectAccessReviewPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSelfSubjectAccessReviewPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SelfSubjectAccessReviewPatchState, opts ...pulumi.ResourceOption) (*SelfSubjectAccessReviewPatch, error) {
	var resource SelfSubjectAccessReviewPatch
	err := ctx.ReadResource("kubernetes:authorization.k8s.io/v1:SelfSubjectAccessReviewPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SelfSubjectAccessReviewPatch resources.
type selfSubjectAccessReviewPatchState struct {
}

type SelfSubjectAccessReviewPatchState struct {
}

func (SelfSubjectAccessReviewPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*selfSubjectAccessReviewPatchState)(nil)).Elem()
}

type selfSubjectAccessReviewPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec holds information about the request being evaluated.  user and groups must be empty
	Spec *SelfSubjectAccessReviewSpec `pulumi:"spec"`
}

// The set of arguments for constructing a SelfSubjectAccessReviewPatch resource.
type SelfSubjectAccessReviewPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Spec holds information about the request being evaluated.  user and groups must be empty
	Spec SelfSubjectAccessReviewSpecPtrInput
}

func (SelfSubjectAccessReviewPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*selfSubjectAccessReviewPatchArgs)(nil)).Elem()
}

type SelfSubjectAccessReviewPatchInput interface {
	pulumi.Input

	ToSelfSubjectAccessReviewPatchOutput() SelfSubjectAccessReviewPatchOutput
	ToSelfSubjectAccessReviewPatchOutputWithContext(ctx context.Context) SelfSubjectAccessReviewPatchOutput
}

func (*SelfSubjectAccessReviewPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**SelfSubjectAccessReviewPatch)(nil)).Elem()
}

func (i *SelfSubjectAccessReviewPatch) ToSelfSubjectAccessReviewPatchOutput() SelfSubjectAccessReviewPatchOutput {
	return i.ToSelfSubjectAccessReviewPatchOutputWithContext(context.Background())
}

func (i *SelfSubjectAccessReviewPatch) ToSelfSubjectAccessReviewPatchOutputWithContext(ctx context.Context) SelfSubjectAccessReviewPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectAccessReviewPatchOutput)
}

// SelfSubjectAccessReviewPatchArrayInput is an input type that accepts SelfSubjectAccessReviewPatchArray and SelfSubjectAccessReviewPatchArrayOutput values.
// You can construct a concrete instance of `SelfSubjectAccessReviewPatchArrayInput` via:
//
//          SelfSubjectAccessReviewPatchArray{ SelfSubjectAccessReviewPatchArgs{...} }
type SelfSubjectAccessReviewPatchArrayInput interface {
	pulumi.Input

	ToSelfSubjectAccessReviewPatchArrayOutput() SelfSubjectAccessReviewPatchArrayOutput
	ToSelfSubjectAccessReviewPatchArrayOutputWithContext(context.Context) SelfSubjectAccessReviewPatchArrayOutput
}

type SelfSubjectAccessReviewPatchArray []SelfSubjectAccessReviewPatchInput

func (SelfSubjectAccessReviewPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SelfSubjectAccessReviewPatch)(nil)).Elem()
}

func (i SelfSubjectAccessReviewPatchArray) ToSelfSubjectAccessReviewPatchArrayOutput() SelfSubjectAccessReviewPatchArrayOutput {
	return i.ToSelfSubjectAccessReviewPatchArrayOutputWithContext(context.Background())
}

func (i SelfSubjectAccessReviewPatchArray) ToSelfSubjectAccessReviewPatchArrayOutputWithContext(ctx context.Context) SelfSubjectAccessReviewPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectAccessReviewPatchArrayOutput)
}

// SelfSubjectAccessReviewPatchMapInput is an input type that accepts SelfSubjectAccessReviewPatchMap and SelfSubjectAccessReviewPatchMapOutput values.
// You can construct a concrete instance of `SelfSubjectAccessReviewPatchMapInput` via:
//
//          SelfSubjectAccessReviewPatchMap{ "key": SelfSubjectAccessReviewPatchArgs{...} }
type SelfSubjectAccessReviewPatchMapInput interface {
	pulumi.Input

	ToSelfSubjectAccessReviewPatchMapOutput() SelfSubjectAccessReviewPatchMapOutput
	ToSelfSubjectAccessReviewPatchMapOutputWithContext(context.Context) SelfSubjectAccessReviewPatchMapOutput
}

type SelfSubjectAccessReviewPatchMap map[string]SelfSubjectAccessReviewPatchInput

func (SelfSubjectAccessReviewPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SelfSubjectAccessReviewPatch)(nil)).Elem()
}

func (i SelfSubjectAccessReviewPatchMap) ToSelfSubjectAccessReviewPatchMapOutput() SelfSubjectAccessReviewPatchMapOutput {
	return i.ToSelfSubjectAccessReviewPatchMapOutputWithContext(context.Background())
}

func (i SelfSubjectAccessReviewPatchMap) ToSelfSubjectAccessReviewPatchMapOutputWithContext(ctx context.Context) SelfSubjectAccessReviewPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectAccessReviewPatchMapOutput)
}

type SelfSubjectAccessReviewPatchOutput struct{ *pulumi.OutputState }

func (SelfSubjectAccessReviewPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SelfSubjectAccessReviewPatch)(nil)).Elem()
}

func (o SelfSubjectAccessReviewPatchOutput) ToSelfSubjectAccessReviewPatchOutput() SelfSubjectAccessReviewPatchOutput {
	return o
}

func (o SelfSubjectAccessReviewPatchOutput) ToSelfSubjectAccessReviewPatchOutputWithContext(ctx context.Context) SelfSubjectAccessReviewPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o SelfSubjectAccessReviewPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SelfSubjectAccessReviewPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o SelfSubjectAccessReviewPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SelfSubjectAccessReviewPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o SelfSubjectAccessReviewPatchOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *SelfSubjectAccessReviewPatch) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Spec holds information about the request being evaluated.  user and groups must be empty
func (o SelfSubjectAccessReviewPatchOutput) Spec() SelfSubjectAccessReviewSpecOutput {
	return o.ApplyT(func(v *SelfSubjectAccessReviewPatch) SelfSubjectAccessReviewSpecOutput { return v.Spec }).(SelfSubjectAccessReviewSpecOutput)
}

// Status is filled in by the server and indicates whether the request is allowed or not
func (o SelfSubjectAccessReviewPatchOutput) Status() SubjectAccessReviewStatusPtrOutput {
	return o.ApplyT(func(v *SelfSubjectAccessReviewPatch) SubjectAccessReviewStatusPtrOutput { return v.Status }).(SubjectAccessReviewStatusPtrOutput)
}

type SelfSubjectAccessReviewPatchArrayOutput struct{ *pulumi.OutputState }

func (SelfSubjectAccessReviewPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SelfSubjectAccessReviewPatch)(nil)).Elem()
}

func (o SelfSubjectAccessReviewPatchArrayOutput) ToSelfSubjectAccessReviewPatchArrayOutput() SelfSubjectAccessReviewPatchArrayOutput {
	return o
}

func (o SelfSubjectAccessReviewPatchArrayOutput) ToSelfSubjectAccessReviewPatchArrayOutputWithContext(ctx context.Context) SelfSubjectAccessReviewPatchArrayOutput {
	return o
}

func (o SelfSubjectAccessReviewPatchArrayOutput) Index(i pulumi.IntInput) SelfSubjectAccessReviewPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SelfSubjectAccessReviewPatch {
		return vs[0].([]*SelfSubjectAccessReviewPatch)[vs[1].(int)]
	}).(SelfSubjectAccessReviewPatchOutput)
}

type SelfSubjectAccessReviewPatchMapOutput struct{ *pulumi.OutputState }

func (SelfSubjectAccessReviewPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SelfSubjectAccessReviewPatch)(nil)).Elem()
}

func (o SelfSubjectAccessReviewPatchMapOutput) ToSelfSubjectAccessReviewPatchMapOutput() SelfSubjectAccessReviewPatchMapOutput {
	return o
}

func (o SelfSubjectAccessReviewPatchMapOutput) ToSelfSubjectAccessReviewPatchMapOutputWithContext(ctx context.Context) SelfSubjectAccessReviewPatchMapOutput {
	return o
}

func (o SelfSubjectAccessReviewPatchMapOutput) MapIndex(k pulumi.StringInput) SelfSubjectAccessReviewPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SelfSubjectAccessReviewPatch {
		return vs[0].(map[string]*SelfSubjectAccessReviewPatch)[vs[1].(string)]
	}).(SelfSubjectAccessReviewPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSubjectAccessReviewPatchInput)(nil)).Elem(), &SelfSubjectAccessReviewPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSubjectAccessReviewPatchArrayInput)(nil)).Elem(), SelfSubjectAccessReviewPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSubjectAccessReviewPatchMapInput)(nil)).Elem(), SelfSubjectAccessReviewPatchMap{})
	pulumi.RegisterOutputType(SelfSubjectAccessReviewPatchOutput{})
	pulumi.RegisterOutputType(SelfSubjectAccessReviewPatchArrayOutput{})
	pulumi.RegisterOutputType(SelfSubjectAccessReviewPatchMapOutput{})
}
