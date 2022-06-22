// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// APIService represents a server for a particular GroupVersion. Name must be "version.group".
type APIServicePatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec contains information for locating and communicating with a server
	Spec APIServiceSpecPtrOutput `pulumi:"spec"`
	// Status contains derived information about an API server
	Status APIServiceStatusPtrOutput `pulumi:"status"`
}

// NewAPIServicePatch registers a new resource with the given unique name, arguments, and options.
func NewAPIServicePatch(ctx *pulumi.Context,
	name string, args *APIServicePatchArgs, opts ...pulumi.ResourceOption) (*APIServicePatch, error) {
	if args == nil {
		args = &APIServicePatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("apiregistration.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("APIService")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:apiregistration.k8s.io/v1:APIService"),
		},
		{
			Type: pulumi.String("kubernetes:apiregistration/v1:APIService"),
		},
		{
			Type: pulumi.String("kubernetes:apiregistration/v1beta1:APIService"),
		},
	})
	opts = append(opts, aliases)
	var resource APIServicePatch
	err := ctx.RegisterResource("kubernetes:apiregistration.k8s.io/v1beta1:APIServicePatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAPIServicePatch gets an existing APIServicePatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAPIServicePatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *APIServicePatchState, opts ...pulumi.ResourceOption) (*APIServicePatch, error) {
	var resource APIServicePatch
	err := ctx.ReadResource("kubernetes:apiregistration.k8s.io/v1beta1:APIServicePatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering APIServicePatch resources.
type apiservicePatchState struct {
}

type APIServicePatchState struct {
}

func (APIServicePatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiservicePatchState)(nil)).Elem()
}

type apiservicePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec contains information for locating and communicating with a server
	Spec *APIServiceSpec `pulumi:"spec"`
}

// The set of arguments for constructing a APIServicePatch resource.
type APIServicePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec contains information for locating and communicating with a server
	Spec APIServiceSpecPtrInput
}

func (APIServicePatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiservicePatchArgs)(nil)).Elem()
}

type APIServicePatchInput interface {
	pulumi.Input

	ToAPIServicePatchOutput() APIServicePatchOutput
	ToAPIServicePatchOutputWithContext(ctx context.Context) APIServicePatchOutput
}

func (*APIServicePatch) ElementType() reflect.Type {
	return reflect.TypeOf((**APIServicePatch)(nil)).Elem()
}

func (i *APIServicePatch) ToAPIServicePatchOutput() APIServicePatchOutput {
	return i.ToAPIServicePatchOutputWithContext(context.Background())
}

func (i *APIServicePatch) ToAPIServicePatchOutputWithContext(ctx context.Context) APIServicePatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIServicePatchOutput)
}

// APIServicePatchArrayInput is an input type that accepts APIServicePatchArray and APIServicePatchArrayOutput values.
// You can construct a concrete instance of `APIServicePatchArrayInput` via:
//
//          APIServicePatchArray{ APIServicePatchArgs{...} }
type APIServicePatchArrayInput interface {
	pulumi.Input

	ToAPIServicePatchArrayOutput() APIServicePatchArrayOutput
	ToAPIServicePatchArrayOutputWithContext(context.Context) APIServicePatchArrayOutput
}

type APIServicePatchArray []APIServicePatchInput

func (APIServicePatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*APIServicePatch)(nil)).Elem()
}

func (i APIServicePatchArray) ToAPIServicePatchArrayOutput() APIServicePatchArrayOutput {
	return i.ToAPIServicePatchArrayOutputWithContext(context.Background())
}

func (i APIServicePatchArray) ToAPIServicePatchArrayOutputWithContext(ctx context.Context) APIServicePatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIServicePatchArrayOutput)
}

// APIServicePatchMapInput is an input type that accepts APIServicePatchMap and APIServicePatchMapOutput values.
// You can construct a concrete instance of `APIServicePatchMapInput` via:
//
//          APIServicePatchMap{ "key": APIServicePatchArgs{...} }
type APIServicePatchMapInput interface {
	pulumi.Input

	ToAPIServicePatchMapOutput() APIServicePatchMapOutput
	ToAPIServicePatchMapOutputWithContext(context.Context) APIServicePatchMapOutput
}

type APIServicePatchMap map[string]APIServicePatchInput

func (APIServicePatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*APIServicePatch)(nil)).Elem()
}

func (i APIServicePatchMap) ToAPIServicePatchMapOutput() APIServicePatchMapOutput {
	return i.ToAPIServicePatchMapOutputWithContext(context.Background())
}

func (i APIServicePatchMap) ToAPIServicePatchMapOutputWithContext(ctx context.Context) APIServicePatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIServicePatchMapOutput)
}

type APIServicePatchOutput struct{ *pulumi.OutputState }

func (APIServicePatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**APIServicePatch)(nil)).Elem()
}

func (o APIServicePatchOutput) ToAPIServicePatchOutput() APIServicePatchOutput {
	return o
}

func (o APIServicePatchOutput) ToAPIServicePatchOutputWithContext(ctx context.Context) APIServicePatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o APIServicePatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APIServicePatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o APIServicePatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APIServicePatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o APIServicePatchOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *APIServicePatch) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Spec contains information for locating and communicating with a server
func (o APIServicePatchOutput) Spec() APIServiceSpecPtrOutput {
	return o.ApplyT(func(v *APIServicePatch) APIServiceSpecPtrOutput { return v.Spec }).(APIServiceSpecPtrOutput)
}

// Status contains derived information about an API server
func (o APIServicePatchOutput) Status() APIServiceStatusPtrOutput {
	return o.ApplyT(func(v *APIServicePatch) APIServiceStatusPtrOutput { return v.Status }).(APIServiceStatusPtrOutput)
}

type APIServicePatchArrayOutput struct{ *pulumi.OutputState }

func (APIServicePatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*APIServicePatch)(nil)).Elem()
}

func (o APIServicePatchArrayOutput) ToAPIServicePatchArrayOutput() APIServicePatchArrayOutput {
	return o
}

func (o APIServicePatchArrayOutput) ToAPIServicePatchArrayOutputWithContext(ctx context.Context) APIServicePatchArrayOutput {
	return o
}

func (o APIServicePatchArrayOutput) Index(i pulumi.IntInput) APIServicePatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *APIServicePatch {
		return vs[0].([]*APIServicePatch)[vs[1].(int)]
	}).(APIServicePatchOutput)
}

type APIServicePatchMapOutput struct{ *pulumi.OutputState }

func (APIServicePatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*APIServicePatch)(nil)).Elem()
}

func (o APIServicePatchMapOutput) ToAPIServicePatchMapOutput() APIServicePatchMapOutput {
	return o
}

func (o APIServicePatchMapOutput) ToAPIServicePatchMapOutputWithContext(ctx context.Context) APIServicePatchMapOutput {
	return o
}

func (o APIServicePatchMapOutput) MapIndex(k pulumi.StringInput) APIServicePatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *APIServicePatch {
		return vs[0].(map[string]*APIServicePatch)[vs[1].(string)]
	}).(APIServicePatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*APIServicePatchInput)(nil)).Elem(), &APIServicePatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*APIServicePatchArrayInput)(nil)).Elem(), APIServicePatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*APIServicePatchMapInput)(nil)).Elem(), APIServicePatchMap{})
	pulumi.RegisterOutputType(APIServicePatchOutput{})
	pulumi.RegisterOutputType(APIServicePatchArrayOutput{})
	pulumi.RegisterOutputType(APIServicePatchMapOutput{})
}
