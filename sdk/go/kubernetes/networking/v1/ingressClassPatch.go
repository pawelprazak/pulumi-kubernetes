// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IngressClass represents the class of the Ingress, referenced by the Ingress Spec. The `ingressclass.kubernetes.io/is-default-class` annotation can be used to indicate that an IngressClass should be considered default. When a single IngressClass resource has this annotation set to true, new Ingress resources without a class specified will be assigned this default class.
type IngressClassPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec is the desired state of the IngressClass. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec IngressClassSpecPtrOutput `pulumi:"spec"`
}

// NewIngressClassPatch registers a new resource with the given unique name, arguments, and options.
func NewIngressClassPatch(ctx *pulumi.Context,
	name string, args *IngressClassPatchArgs, opts ...pulumi.ResourceOption) (*IngressClassPatch, error) {
	if args == nil {
		args = &IngressClassPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("networking.k8s.io/v1")
	args.Kind = pulumi.StringPtr("IngressClass")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:networking.k8s.io/v1beta1:IngressClass"),
		},
	})
	opts = append(opts, aliases)
	var resource IngressClassPatch
	err := ctx.RegisterResource("kubernetes:networking.k8s.io/v1:IngressClassPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIngressClassPatch gets an existing IngressClassPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIngressClassPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IngressClassPatchState, opts ...pulumi.ResourceOption) (*IngressClassPatch, error) {
	var resource IngressClassPatch
	err := ctx.ReadResource("kubernetes:networking.k8s.io/v1:IngressClassPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IngressClassPatch resources.
type ingressClassPatchState struct {
}

type IngressClassPatchState struct {
}

func (IngressClassPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*ingressClassPatchState)(nil)).Elem()
}

type ingressClassPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec is the desired state of the IngressClass. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *IngressClassSpec `pulumi:"spec"`
}

// The set of arguments for constructing a IngressClassPatch resource.
type IngressClassPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Spec is the desired state of the IngressClass. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec IngressClassSpecPtrInput
}

func (IngressClassPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ingressClassPatchArgs)(nil)).Elem()
}

type IngressClassPatchInput interface {
	pulumi.Input

	ToIngressClassPatchOutput() IngressClassPatchOutput
	ToIngressClassPatchOutputWithContext(ctx context.Context) IngressClassPatchOutput
}

func (*IngressClassPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressClassPatch)(nil)).Elem()
}

func (i *IngressClassPatch) ToIngressClassPatchOutput() IngressClassPatchOutput {
	return i.ToIngressClassPatchOutputWithContext(context.Background())
}

func (i *IngressClassPatch) ToIngressClassPatchOutputWithContext(ctx context.Context) IngressClassPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressClassPatchOutput)
}

// IngressClassPatchArrayInput is an input type that accepts IngressClassPatchArray and IngressClassPatchArrayOutput values.
// You can construct a concrete instance of `IngressClassPatchArrayInput` via:
//
//          IngressClassPatchArray{ IngressClassPatchArgs{...} }
type IngressClassPatchArrayInput interface {
	pulumi.Input

	ToIngressClassPatchArrayOutput() IngressClassPatchArrayOutput
	ToIngressClassPatchArrayOutputWithContext(context.Context) IngressClassPatchArrayOutput
}

type IngressClassPatchArray []IngressClassPatchInput

func (IngressClassPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IngressClassPatch)(nil)).Elem()
}

func (i IngressClassPatchArray) ToIngressClassPatchArrayOutput() IngressClassPatchArrayOutput {
	return i.ToIngressClassPatchArrayOutputWithContext(context.Background())
}

func (i IngressClassPatchArray) ToIngressClassPatchArrayOutputWithContext(ctx context.Context) IngressClassPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressClassPatchArrayOutput)
}

// IngressClassPatchMapInput is an input type that accepts IngressClassPatchMap and IngressClassPatchMapOutput values.
// You can construct a concrete instance of `IngressClassPatchMapInput` via:
//
//          IngressClassPatchMap{ "key": IngressClassPatchArgs{...} }
type IngressClassPatchMapInput interface {
	pulumi.Input

	ToIngressClassPatchMapOutput() IngressClassPatchMapOutput
	ToIngressClassPatchMapOutputWithContext(context.Context) IngressClassPatchMapOutput
}

type IngressClassPatchMap map[string]IngressClassPatchInput

func (IngressClassPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IngressClassPatch)(nil)).Elem()
}

func (i IngressClassPatchMap) ToIngressClassPatchMapOutput() IngressClassPatchMapOutput {
	return i.ToIngressClassPatchMapOutputWithContext(context.Background())
}

func (i IngressClassPatchMap) ToIngressClassPatchMapOutputWithContext(ctx context.Context) IngressClassPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressClassPatchMapOutput)
}

type IngressClassPatchOutput struct{ *pulumi.OutputState }

func (IngressClassPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressClassPatch)(nil)).Elem()
}

func (o IngressClassPatchOutput) ToIngressClassPatchOutput() IngressClassPatchOutput {
	return o
}

func (o IngressClassPatchOutput) ToIngressClassPatchOutputWithContext(ctx context.Context) IngressClassPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o IngressClassPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IngressClassPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o IngressClassPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IngressClassPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o IngressClassPatchOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *IngressClassPatch) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Spec is the desired state of the IngressClass. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o IngressClassPatchOutput) Spec() IngressClassSpecPtrOutput {
	return o.ApplyT(func(v *IngressClassPatch) IngressClassSpecPtrOutput { return v.Spec }).(IngressClassSpecPtrOutput)
}

type IngressClassPatchArrayOutput struct{ *pulumi.OutputState }

func (IngressClassPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IngressClassPatch)(nil)).Elem()
}

func (o IngressClassPatchArrayOutput) ToIngressClassPatchArrayOutput() IngressClassPatchArrayOutput {
	return o
}

func (o IngressClassPatchArrayOutput) ToIngressClassPatchArrayOutputWithContext(ctx context.Context) IngressClassPatchArrayOutput {
	return o
}

func (o IngressClassPatchArrayOutput) Index(i pulumi.IntInput) IngressClassPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IngressClassPatch {
		return vs[0].([]*IngressClassPatch)[vs[1].(int)]
	}).(IngressClassPatchOutput)
}

type IngressClassPatchMapOutput struct{ *pulumi.OutputState }

func (IngressClassPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IngressClassPatch)(nil)).Elem()
}

func (o IngressClassPatchMapOutput) ToIngressClassPatchMapOutput() IngressClassPatchMapOutput {
	return o
}

func (o IngressClassPatchMapOutput) ToIngressClassPatchMapOutputWithContext(ctx context.Context) IngressClassPatchMapOutput {
	return o
}

func (o IngressClassPatchMapOutput) MapIndex(k pulumi.StringInput) IngressClassPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IngressClassPatch {
		return vs[0].(map[string]*IngressClassPatch)[vs[1].(string)]
	}).(IngressClassPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IngressClassPatchInput)(nil)).Elem(), &IngressClassPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*IngressClassPatchArrayInput)(nil)).Elem(), IngressClassPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IngressClassPatchMapInput)(nil)).Elem(), IngressClassPatchMap{})
	pulumi.RegisterOutputType(IngressClassPatchOutput{})
	pulumi.RegisterOutputType(IngressClassPatchArrayOutput{})
	pulumi.RegisterOutputType(IngressClassPatchMapOutput{})
}
