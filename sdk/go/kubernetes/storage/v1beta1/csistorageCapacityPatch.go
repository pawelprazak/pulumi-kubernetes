// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CSIStorageCapacity stores the result of one CSI GetCapacity call. For a given StorageClass, this describes the available capacity in a particular topology segment.  This can be used when considering where to instantiate new PersistentVolumes.
//
// For example this can express things like: - StorageClass "standard" has "1234 GiB" available in "topology.kubernetes.io/zone=us-east1" - StorageClass "localssd" has "10 GiB" available in "kubernetes.io/hostname=knode-abc123"
//
// The following three cases all imply that no capacity is available for a certain combination: - no object exists with suitable topology and storage class name - such an object exists, but the capacity is unset - such an object exists, but the capacity is zero
//
// The producer of these objects can decide which approach is more suitable.
//
// They are consumed by the kube-scheduler when a CSI driver opts into capacity-aware scheduling with CSIDriverSpec.StorageCapacity. The scheduler compares the MaximumVolumeSize against the requested size of pending volumes to filter out unsuitable nodes. If MaximumVolumeSize is unset, it falls back to a comparison against the less precise Capacity. If that is also unset, the scheduler assumes that capacity is insufficient and tries some other node.
type CSIStorageCapacityPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Capacity is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// The semantic is currently (CSI spec 1.2) defined as: The available capacity, in bytes, of the storage that can be used to provision volumes. If not set, that information is currently unavailable.
	Capacity pulumi.StringPtrOutput `pulumi:"capacity"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// MaximumVolumeSize is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// This is defined since CSI spec 1.4.0 as the largest size that may be used in a CreateVolumeRequest.capacity_range.required_bytes field to create a volume with the same parameters as those in GetCapacityRequest. The corresponding value in the Kubernetes API is ResourceRequirements.Requests in a volume claim.
	MaximumVolumeSize pulumi.StringPtrOutput `pulumi:"maximumVolumeSize"`
	// Standard object's metadata. The name has no particular meaning. It must be be a DNS subdomain (dots allowed, 253 characters). To ensure that there are no conflicts with other CSI drivers on the cluster, the recommendation is to use csisc-<uuid>, a generated name, or a reverse-domain name which ends with the unique CSI driver name.
	//
	// Objects are namespaced.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// NodeTopology defines which nodes have access to the storage for which capacity was reported. If not set, the storage is not accessible from any node in the cluster. If empty, the storage is accessible from all nodes. This field is immutable.
	NodeTopology metav1.LabelSelectorPatchPtrOutput `pulumi:"nodeTopology"`
	// The name of the StorageClass that the reported capacity applies to. It must meet the same requirements as the name of a StorageClass object (non-empty, DNS subdomain). If that object no longer exists, the CSIStorageCapacity object is obsolete and should be removed by its creator. This field is immutable.
	StorageClassName pulumi.StringPtrOutput `pulumi:"storageClassName"`
}

// NewCSIStorageCapacityPatch registers a new resource with the given unique name, arguments, and options.
func NewCSIStorageCapacityPatch(ctx *pulumi.Context,
	name string, args *CSIStorageCapacityPatchArgs, opts ...pulumi.ResourceOption) (*CSIStorageCapacityPatch, error) {
	if args == nil {
		args = &CSIStorageCapacityPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("CSIStorageCapacity")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:storage.k8s.io/v1:CSIStorageCapacityPatch"),
		},
		{
			Type: pulumi.String("kubernetes:storage.k8s.io/v1alpha1:CSIStorageCapacityPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource CSIStorageCapacityPatch
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1beta1:CSIStorageCapacityPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCSIStorageCapacityPatch gets an existing CSIStorageCapacityPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCSIStorageCapacityPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CSIStorageCapacityPatchState, opts ...pulumi.ResourceOption) (*CSIStorageCapacityPatch, error) {
	var resource CSIStorageCapacityPatch
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1beta1:CSIStorageCapacityPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CSIStorageCapacityPatch resources.
type csistorageCapacityPatchState struct {
}

type CSIStorageCapacityPatchState struct {
}

func (CSIStorageCapacityPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*csistorageCapacityPatchState)(nil)).Elem()
}

type csistorageCapacityPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Capacity is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// The semantic is currently (CSI spec 1.2) defined as: The available capacity, in bytes, of the storage that can be used to provision volumes. If not set, that information is currently unavailable.
	Capacity *string `pulumi:"capacity"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// MaximumVolumeSize is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// This is defined since CSI spec 1.4.0 as the largest size that may be used in a CreateVolumeRequest.capacity_range.required_bytes field to create a volume with the same parameters as those in GetCapacityRequest. The corresponding value in the Kubernetes API is ResourceRequirements.Requests in a volume claim.
	MaximumVolumeSize *string `pulumi:"maximumVolumeSize"`
	// Standard object's metadata. The name has no particular meaning. It must be be a DNS subdomain (dots allowed, 253 characters). To ensure that there are no conflicts with other CSI drivers on the cluster, the recommendation is to use csisc-<uuid>, a generated name, or a reverse-domain name which ends with the unique CSI driver name.
	//
	// Objects are namespaced.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// NodeTopology defines which nodes have access to the storage for which capacity was reported. If not set, the storage is not accessible from any node in the cluster. If empty, the storage is accessible from all nodes. This field is immutable.
	NodeTopology *metav1.LabelSelectorPatch `pulumi:"nodeTopology"`
	// The name of the StorageClass that the reported capacity applies to. It must meet the same requirements as the name of a StorageClass object (non-empty, DNS subdomain). If that object no longer exists, the CSIStorageCapacity object is obsolete and should be removed by its creator. This field is immutable.
	StorageClassName *string `pulumi:"storageClassName"`
}

// The set of arguments for constructing a CSIStorageCapacityPatch resource.
type CSIStorageCapacityPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Capacity is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// The semantic is currently (CSI spec 1.2) defined as: The available capacity, in bytes, of the storage that can be used to provision volumes. If not set, that information is currently unavailable.
	Capacity pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// MaximumVolumeSize is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// This is defined since CSI spec 1.4.0 as the largest size that may be used in a CreateVolumeRequest.capacity_range.required_bytes field to create a volume with the same parameters as those in GetCapacityRequest. The corresponding value in the Kubernetes API is ResourceRequirements.Requests in a volume claim.
	MaximumVolumeSize pulumi.StringPtrInput
	// Standard object's metadata. The name has no particular meaning. It must be be a DNS subdomain (dots allowed, 253 characters). To ensure that there are no conflicts with other CSI drivers on the cluster, the recommendation is to use csisc-<uuid>, a generated name, or a reverse-domain name which ends with the unique CSI driver name.
	//
	// Objects are namespaced.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// NodeTopology defines which nodes have access to the storage for which capacity was reported. If not set, the storage is not accessible from any node in the cluster. If empty, the storage is accessible from all nodes. This field is immutable.
	NodeTopology metav1.LabelSelectorPatchPtrInput
	// The name of the StorageClass that the reported capacity applies to. It must meet the same requirements as the name of a StorageClass object (non-empty, DNS subdomain). If that object no longer exists, the CSIStorageCapacity object is obsolete and should be removed by its creator. This field is immutable.
	StorageClassName pulumi.StringPtrInput
}

func (CSIStorageCapacityPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*csistorageCapacityPatchArgs)(nil)).Elem()
}

type CSIStorageCapacityPatchInput interface {
	pulumi.Input

	ToCSIStorageCapacityPatchOutput() CSIStorageCapacityPatchOutput
	ToCSIStorageCapacityPatchOutputWithContext(ctx context.Context) CSIStorageCapacityPatchOutput
}

func (*CSIStorageCapacityPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**CSIStorageCapacityPatch)(nil)).Elem()
}

func (i *CSIStorageCapacityPatch) ToCSIStorageCapacityPatchOutput() CSIStorageCapacityPatchOutput {
	return i.ToCSIStorageCapacityPatchOutputWithContext(context.Background())
}

func (i *CSIStorageCapacityPatch) ToCSIStorageCapacityPatchOutputWithContext(ctx context.Context) CSIStorageCapacityPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIStorageCapacityPatchOutput)
}

// CSIStorageCapacityPatchArrayInput is an input type that accepts CSIStorageCapacityPatchArray and CSIStorageCapacityPatchArrayOutput values.
// You can construct a concrete instance of `CSIStorageCapacityPatchArrayInput` via:
//
//          CSIStorageCapacityPatchArray{ CSIStorageCapacityPatchArgs{...} }
type CSIStorageCapacityPatchArrayInput interface {
	pulumi.Input

	ToCSIStorageCapacityPatchArrayOutput() CSIStorageCapacityPatchArrayOutput
	ToCSIStorageCapacityPatchArrayOutputWithContext(context.Context) CSIStorageCapacityPatchArrayOutput
}

type CSIStorageCapacityPatchArray []CSIStorageCapacityPatchInput

func (CSIStorageCapacityPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSIStorageCapacityPatch)(nil)).Elem()
}

func (i CSIStorageCapacityPatchArray) ToCSIStorageCapacityPatchArrayOutput() CSIStorageCapacityPatchArrayOutput {
	return i.ToCSIStorageCapacityPatchArrayOutputWithContext(context.Background())
}

func (i CSIStorageCapacityPatchArray) ToCSIStorageCapacityPatchArrayOutputWithContext(ctx context.Context) CSIStorageCapacityPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIStorageCapacityPatchArrayOutput)
}

// CSIStorageCapacityPatchMapInput is an input type that accepts CSIStorageCapacityPatchMap and CSIStorageCapacityPatchMapOutput values.
// You can construct a concrete instance of `CSIStorageCapacityPatchMapInput` via:
//
//          CSIStorageCapacityPatchMap{ "key": CSIStorageCapacityPatchArgs{...} }
type CSIStorageCapacityPatchMapInput interface {
	pulumi.Input

	ToCSIStorageCapacityPatchMapOutput() CSIStorageCapacityPatchMapOutput
	ToCSIStorageCapacityPatchMapOutputWithContext(context.Context) CSIStorageCapacityPatchMapOutput
}

type CSIStorageCapacityPatchMap map[string]CSIStorageCapacityPatchInput

func (CSIStorageCapacityPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSIStorageCapacityPatch)(nil)).Elem()
}

func (i CSIStorageCapacityPatchMap) ToCSIStorageCapacityPatchMapOutput() CSIStorageCapacityPatchMapOutput {
	return i.ToCSIStorageCapacityPatchMapOutputWithContext(context.Background())
}

func (i CSIStorageCapacityPatchMap) ToCSIStorageCapacityPatchMapOutputWithContext(ctx context.Context) CSIStorageCapacityPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIStorageCapacityPatchMapOutput)
}

type CSIStorageCapacityPatchOutput struct{ *pulumi.OutputState }

func (CSIStorageCapacityPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CSIStorageCapacityPatch)(nil)).Elem()
}

func (o CSIStorageCapacityPatchOutput) ToCSIStorageCapacityPatchOutput() CSIStorageCapacityPatchOutput {
	return o
}

func (o CSIStorageCapacityPatchOutput) ToCSIStorageCapacityPatchOutputWithContext(ctx context.Context) CSIStorageCapacityPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o CSIStorageCapacityPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CSIStorageCapacityPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Capacity is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
//
// The semantic is currently (CSI spec 1.2) defined as: The available capacity, in bytes, of the storage that can be used to provision volumes. If not set, that information is currently unavailable.
func (o CSIStorageCapacityPatchOutput) Capacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CSIStorageCapacityPatch) pulumi.StringPtrOutput { return v.Capacity }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o CSIStorageCapacityPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CSIStorageCapacityPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// MaximumVolumeSize is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
//
// This is defined since CSI spec 1.4.0 as the largest size that may be used in a CreateVolumeRequest.capacity_range.required_bytes field to create a volume with the same parameters as those in GetCapacityRequest. The corresponding value in the Kubernetes API is ResourceRequirements.Requests in a volume claim.
func (o CSIStorageCapacityPatchOutput) MaximumVolumeSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CSIStorageCapacityPatch) pulumi.StringPtrOutput { return v.MaximumVolumeSize }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. The name has no particular meaning. It must be be a DNS subdomain (dots allowed, 253 characters). To ensure that there are no conflicts with other CSI drivers on the cluster, the recommendation is to use csisc-<uuid>, a generated name, or a reverse-domain name which ends with the unique CSI driver name.
//
// Objects are namespaced.
//
// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o CSIStorageCapacityPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *CSIStorageCapacityPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// NodeTopology defines which nodes have access to the storage for which capacity was reported. If not set, the storage is not accessible from any node in the cluster. If empty, the storage is accessible from all nodes. This field is immutable.
func (o CSIStorageCapacityPatchOutput) NodeTopology() metav1.LabelSelectorPatchPtrOutput {
	return o.ApplyT(func(v *CSIStorageCapacityPatch) metav1.LabelSelectorPatchPtrOutput { return v.NodeTopology }).(metav1.LabelSelectorPatchPtrOutput)
}

// The name of the StorageClass that the reported capacity applies to. It must meet the same requirements as the name of a StorageClass object (non-empty, DNS subdomain). If that object no longer exists, the CSIStorageCapacity object is obsolete and should be removed by its creator. This field is immutable.
func (o CSIStorageCapacityPatchOutput) StorageClassName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CSIStorageCapacityPatch) pulumi.StringPtrOutput { return v.StorageClassName }).(pulumi.StringPtrOutput)
}

type CSIStorageCapacityPatchArrayOutput struct{ *pulumi.OutputState }

func (CSIStorageCapacityPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSIStorageCapacityPatch)(nil)).Elem()
}

func (o CSIStorageCapacityPatchArrayOutput) ToCSIStorageCapacityPatchArrayOutput() CSIStorageCapacityPatchArrayOutput {
	return o
}

func (o CSIStorageCapacityPatchArrayOutput) ToCSIStorageCapacityPatchArrayOutputWithContext(ctx context.Context) CSIStorageCapacityPatchArrayOutput {
	return o
}

func (o CSIStorageCapacityPatchArrayOutput) Index(i pulumi.IntInput) CSIStorageCapacityPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CSIStorageCapacityPatch {
		return vs[0].([]*CSIStorageCapacityPatch)[vs[1].(int)]
	}).(CSIStorageCapacityPatchOutput)
}

type CSIStorageCapacityPatchMapOutput struct{ *pulumi.OutputState }

func (CSIStorageCapacityPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSIStorageCapacityPatch)(nil)).Elem()
}

func (o CSIStorageCapacityPatchMapOutput) ToCSIStorageCapacityPatchMapOutput() CSIStorageCapacityPatchMapOutput {
	return o
}

func (o CSIStorageCapacityPatchMapOutput) ToCSIStorageCapacityPatchMapOutputWithContext(ctx context.Context) CSIStorageCapacityPatchMapOutput {
	return o
}

func (o CSIStorageCapacityPatchMapOutput) MapIndex(k pulumi.StringInput) CSIStorageCapacityPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CSIStorageCapacityPatch {
		return vs[0].(map[string]*CSIStorageCapacityPatch)[vs[1].(string)]
	}).(CSIStorageCapacityPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CSIStorageCapacityPatchInput)(nil)).Elem(), &CSIStorageCapacityPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSIStorageCapacityPatchArrayInput)(nil)).Elem(), CSIStorageCapacityPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSIStorageCapacityPatchMapInput)(nil)).Elem(), CSIStorageCapacityPatchMap{})
	pulumi.RegisterOutputType(CSIStorageCapacityPatchOutput{})
	pulumi.RegisterOutputType(CSIStorageCapacityPatchArrayOutput{})
	pulumi.RegisterOutputType(CSIStorageCapacityPatchMapOutput{})
}
