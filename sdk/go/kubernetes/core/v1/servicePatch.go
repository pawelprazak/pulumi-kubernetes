// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Service is a named abstraction of software service (for example, mysql) consisting of local port (for example 3306) that the proxy listens on, and the selector that determines which pods will answer requests sent through the proxy.
//
// This resource waits until its status is ready before registering success
// for create/update, and populating output properties from the current state of the resource.
// The following conditions are used to determine whether the resource creation has
// succeeded or failed:
//
// 1. Service object exists.
// 2. Related Endpoint objects are created. Each time we get an update, wait 10 seconds
//    for any stragglers.
// 3. The endpoints objects target some number of living objects (unless the Service is
//    an "empty headless" Service [1] or a Service with '.spec.type: ExternalName').
// 4. External IP address is allocated (if Service has '.spec.type: LoadBalancer').
//
// Known limitations:
// Services targeting ReplicaSets (and, by extension, Deployments,
// StatefulSets, etc.) with '.spec.replicas' set to 0 are not handled, and will time
// out. To work around this limitation, set 'pulumi.com/skipAwait: "true"' on
// '.metadata.annotations' for the Service. Work to handle this case is in progress [2].
//
// [1] https://kubernetes.io/docs/concepts/services-networking/service/#headless-services
// [2] https://github.com/pulumi/pulumi-kubernetes/pull/703
//
// If the Service has not reached a Ready state after 10 minutes, it will
// time out and mark the resource update as Failed. You can override the default timeout value
// by setting the 'customTimeouts' option on the resource.
type ServicePatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec defines the behavior of a service. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec ServiceSpecPtrOutput `pulumi:"spec"`
	// Most recently observed status of the service. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status ServiceStatusPtrOutput `pulumi:"status"`
}

// NewServicePatch registers a new resource with the given unique name, arguments, and options.
func NewServicePatch(ctx *pulumi.Context,
	name string, args *ServicePatchArgs, opts ...pulumi.ResourceOption) (*ServicePatch, error) {
	if args == nil {
		args = &ServicePatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("Service")
	var resource ServicePatch
	err := ctx.RegisterResource("kubernetes:core/v1:ServicePatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePatch gets an existing ServicePatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePatchState, opts ...pulumi.ResourceOption) (*ServicePatch, error) {
	var resource ServicePatch
	err := ctx.ReadResource("kubernetes:core/v1:ServicePatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePatch resources.
type servicePatchState struct {
}

type ServicePatchState struct {
}

func (ServicePatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePatchState)(nil)).Elem()
}

type servicePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec defines the behavior of a service. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *ServiceSpec `pulumi:"spec"`
}

// The set of arguments for constructing a ServicePatch resource.
type ServicePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Spec defines the behavior of a service. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec ServiceSpecPtrInput
}

func (ServicePatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePatchArgs)(nil)).Elem()
}

type ServicePatchInput interface {
	pulumi.Input

	ToServicePatchOutput() ServicePatchOutput
	ToServicePatchOutputWithContext(ctx context.Context) ServicePatchOutput
}

func (*ServicePatch) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePatch)(nil)).Elem()
}

func (i *ServicePatch) ToServicePatchOutput() ServicePatchOutput {
	return i.ToServicePatchOutputWithContext(context.Background())
}

func (i *ServicePatch) ToServicePatchOutputWithContext(ctx context.Context) ServicePatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePatchOutput)
}

// ServicePatchArrayInput is an input type that accepts ServicePatchArray and ServicePatchArrayOutput values.
// You can construct a concrete instance of `ServicePatchArrayInput` via:
//
//          ServicePatchArray{ ServicePatchArgs{...} }
type ServicePatchArrayInput interface {
	pulumi.Input

	ToServicePatchArrayOutput() ServicePatchArrayOutput
	ToServicePatchArrayOutputWithContext(context.Context) ServicePatchArrayOutput
}

type ServicePatchArray []ServicePatchInput

func (ServicePatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePatch)(nil)).Elem()
}

func (i ServicePatchArray) ToServicePatchArrayOutput() ServicePatchArrayOutput {
	return i.ToServicePatchArrayOutputWithContext(context.Background())
}

func (i ServicePatchArray) ToServicePatchArrayOutputWithContext(ctx context.Context) ServicePatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePatchArrayOutput)
}

// ServicePatchMapInput is an input type that accepts ServicePatchMap and ServicePatchMapOutput values.
// You can construct a concrete instance of `ServicePatchMapInput` via:
//
//          ServicePatchMap{ "key": ServicePatchArgs{...} }
type ServicePatchMapInput interface {
	pulumi.Input

	ToServicePatchMapOutput() ServicePatchMapOutput
	ToServicePatchMapOutputWithContext(context.Context) ServicePatchMapOutput
}

type ServicePatchMap map[string]ServicePatchInput

func (ServicePatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePatch)(nil)).Elem()
}

func (i ServicePatchMap) ToServicePatchMapOutput() ServicePatchMapOutput {
	return i.ToServicePatchMapOutputWithContext(context.Background())
}

func (i ServicePatchMap) ToServicePatchMapOutputWithContext(ctx context.Context) ServicePatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePatchMapOutput)
}

type ServicePatchOutput struct{ *pulumi.OutputState }

func (ServicePatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePatch)(nil)).Elem()
}

func (o ServicePatchOutput) ToServicePatchOutput() ServicePatchOutput {
	return o
}

func (o ServicePatchOutput) ToServicePatchOutputWithContext(ctx context.Context) ServicePatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ServicePatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ServicePatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ServicePatchOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *ServicePatch) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Spec defines the behavior of a service. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o ServicePatchOutput) Spec() ServiceSpecPtrOutput {
	return o.ApplyT(func(v *ServicePatch) ServiceSpecPtrOutput { return v.Spec }).(ServiceSpecPtrOutput)
}

// Most recently observed status of the service. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o ServicePatchOutput) Status() ServiceStatusPtrOutput {
	return o.ApplyT(func(v *ServicePatch) ServiceStatusPtrOutput { return v.Status }).(ServiceStatusPtrOutput)
}

type ServicePatchArrayOutput struct{ *pulumi.OutputState }

func (ServicePatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePatch)(nil)).Elem()
}

func (o ServicePatchArrayOutput) ToServicePatchArrayOutput() ServicePatchArrayOutput {
	return o
}

func (o ServicePatchArrayOutput) ToServicePatchArrayOutputWithContext(ctx context.Context) ServicePatchArrayOutput {
	return o
}

func (o ServicePatchArrayOutput) Index(i pulumi.IntInput) ServicePatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServicePatch {
		return vs[0].([]*ServicePatch)[vs[1].(int)]
	}).(ServicePatchOutput)
}

type ServicePatchMapOutput struct{ *pulumi.OutputState }

func (ServicePatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePatch)(nil)).Elem()
}

func (o ServicePatchMapOutput) ToServicePatchMapOutput() ServicePatchMapOutput {
	return o
}

func (o ServicePatchMapOutput) ToServicePatchMapOutputWithContext(ctx context.Context) ServicePatchMapOutput {
	return o
}

func (o ServicePatchMapOutput) MapIndex(k pulumi.StringInput) ServicePatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServicePatch {
		return vs[0].(map[string]*ServicePatch)[vs[1].(string)]
	}).(ServicePatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePatchInput)(nil)).Elem(), &ServicePatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePatchArrayInput)(nil)).Elem(), ServicePatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePatchMapInput)(nil)).Elem(), ServicePatchMap{})
	pulumi.RegisterOutputType(ServicePatchOutput{})
	pulumi.RegisterOutputType(ServicePatchArrayOutput{})
	pulumi.RegisterOutputType(ServicePatchMapOutput{})
}
